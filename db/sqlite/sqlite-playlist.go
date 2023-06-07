package sqlite

import (
	"fmt"
	"github.com/lea-video/go-lea-music/model"
	"strings"
)

func (db *LEASQLiteDB) GetPlaylists() (map[int]*model.Playlist, error) {
	rows, err := db.db.Query("SELECT id, name, items FROM view_playlists_with_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allLists := make(map[int]*model.Playlist)
	for rows.Next() {
		var tmpList model.Playlist
		err := rows.Scan(&tmpList.ID, &tmpList.Name, &tmpList.Items)
		if err != nil {
			return nil, err
		}
		allLists[tmpList.ID] = &tmpList
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allLists, nil
}

func (db *LEASQLiteDB) CreatePlaylists(plists []*model.Playlist) (map[int]*model.Playlist, error) {
	insertPlaylist := "INSERT INTO playlists (name) VALUES (?);"

	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}

	stmtArt, err := tx.Prepare(insertPlaylist)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmtArt.Close()

	playlistsWithID := make(map[int]*model.Playlist)
	for _, plist := range plists {
		result, err := stmtArt.Exec(plist.Name)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		plist.ID = int(id)
		playlistsWithID[plist.ID] = plist

		// insert the members
		err = db.AppendPlaylistItems(plist.ID, plist.Items)
		if err != nil {
			return nil, err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return playlistsWithID, nil
}

func (db *LEASQLiteDB) AppendPlaylistItems(playlistID int, items []int) error {
	// TODO: fetch the highest position once in advance instead of executing the sub-query every time
	// this utilises a nested query to set the position (of the item in the playlist) to the last (highest) element
	insertItem := strings.TrimSpace(`
INSERT INTO playlist_item (playlist, position, element)
SELECT ?, COALESCE(CEIL(MAX(position) + 1), 1), ?
FROM playlist_item
WHERE playlist = ?
GROUP BY playlist;
`)

	tx, err := db.db.Begin()
	if err != nil {
		return err
	}

	stmtMem, err := tx.Prepare(insertItem)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtMem.Close()

	for _, item := range items {
		_, err := stmtMem.Exec(playlistID, item, item)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (db *LEASQLiteDB) ChangePlaylistItemPosition(playlistID int, itemID int, positionAfter int) error {
	var query string
	args := make([]interface{}, 0)
	if positionAfter < 0 {
		// insert at the top
		query = "UPDATE playlist_item SET position = 0 WHERE playlist = ? AND element = ?;"
		args = append(args, playlistID, itemID)
	} else {
		// insert after marker element and before next (which should be position + 1)
		// if marker not found keeps current position
		query = `
UPDATE playlist_item
SET position = COALESCE(
    (SELECT position + 0.5
     FROM playlist_item
     WHERE playlist = ? AND element = ?),
    position
)
WHERE playlist = ? AND element = ?;
`
		args = append(args, playlistID, positionAfter, playlistID, itemID)
	}

	// execute query
	_, err := db.db.Exec(query, args...)
	if err != nil {
		return err
	}

	// normalize so that the IDs are again 1->len(items)
	err = db.normalizePlaylistItemPositions(playlistID)
	if err != nil {
		return err
	}

	return nil
}

func (db *LEASQLiteDB) normalizePlaylistItemPositions(playlistID int) error {
	query := `
UPDATE playlist_item
SET position = new_position
FROM (
    SELECT id, ROW_NUMBER() OVER (PARTITION BY playlist ORDER BY position) AS new_position
    FROM playlist_item WHERE playlist_item.playlist = ?
) AS subquery
WHERE playlist_item.id = subquery.id;
`
	_, err := db.db.Exec(query, playlistID, playlistID)
	if err != nil {
		return err
	}

	return nil
}

func (db *LEASQLiteDB) GetPlaylistsByID(playlistIDs []int) (map[int]*model.Playlist, error) {
	inPlaceholder, inArgs := buildINStatement(playlistIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, name, items FROM view_playlists_with_items WHERE id IN (%s)", inPlaceholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, inArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allLists := make(map[int]*model.Playlist)
	for rows.Next() {
		var tmpList model.Playlist
		err := rows.Scan(&tmpList.ID, &tmpList.Name, &tmpList.Items)
		if err != nil {
			return nil, err
		}
		allLists[tmpList.ID] = &tmpList
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allLists, nil
}
