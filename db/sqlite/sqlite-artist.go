package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/lea-video/go-lea-music/model"
	"github.com/lea-video/go-lea-music/utility"
)

func (db *SQLiteDB) GetArtists() (map[int]*model.OneOfArtist, error) {
	solos, err := db.GetArtistSolos()
	if err != nil {
		return nil, err
	}

	groups, err := db.GetArtistGroups()
	if err != nil {
		return nil, err
	}

	// typecast
	artists := make(map[int]*model.OneOfArtist)
	for _, s := range solos {
		artists[s.ID] = s.ToOneOf()
	}
	for _, g := range groups {
		artists[g.ID] = g.ToOneOf()
	}

	return artists, nil
}

func (db *SQLiteDB) GetArtistsByID(artistIDs []int) (map[int]*model.OneOfArtist, error) {
	solos, err := db.GetArtistSolosByID(artistIDs)
	if err != nil {
		return nil, err
	}

	groups, err := db.GetArtistGroupsByID(artistIDs)
	if err != nil {
		return nil, err
	}

	// typecast
	artists := make(map[int]*model.OneOfArtist)
	for _, s := range solos {
		artists[s.ID] = s.ToOneOf()
	}
	for _, g := range groups {
		artists[g.ID] = g.ToOneOf()
	}

	return artists, nil
}

func (db *SQLiteDB) GetArtistSolos() (map[int]*model.ArtistSolo, error) {
	rows, err := db.db.Query("SELECT id, name FROM artists WHERE is_group = 0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	artists := make(map[int]*model.ArtistSolo)
	for rows.Next() {
		var artistSolo model.ArtistSolo
		err := rows.Scan(&artistSolo.ID, &artistSolo.Name)
		if err != nil {
			return nil, err
		}
		artists[artistSolo.ID] = &artistSolo
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (db *SQLiteDB) GetArtistSolosByID(artistIDs []int) (map[int]*model.ArtistSolo, error) {
	in_placeholder, in_args := build_in_args(artistIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, name FROM artists WHERE is_group = 0 AND id IN (%s)", in_placeholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, in_args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	artists := make(map[int]*model.ArtistSolo)
	for rows.Next() {
		var artistSolo model.ArtistSolo
		err := rows.Scan(&artistSolo.ID, &artistSolo.Name)
		if err != nil {
			return nil, err
		}
		artists[artistSolo.ID] = &artistSolo
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (db *SQLiteDB) CreateArtistSolos(artists []*model.ArtistSolo) (map[int]*model.ArtistSolo, error) {
	insertStmt := "INSERT INTO artists (name, is_group) VALUES (?, ?);"

	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(insertStmt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	artists_w_id := make(map[int]*model.ArtistSolo)
	for _, artist := range artists {
		result, err := stmt.Exec(artist.Name, 0)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		artist.ID = int(id)
		artists_w_id[artist.ID] = artist
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return artists_w_id, nil
}

func (db *SQLiteDB) GetArtistGroups() (map[int]*model.ArtistGroup, error) {
	rows, err := db.db.Query("SELECT id, name, members FROM view_artist_groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	artists := make(map[int]*model.ArtistGroup)
	for rows.Next() {
		var artistGroups model.ArtistGroup
		var member_list sql.NullString
		err := rows.Scan(&artistGroups.ID, &artistGroups.Name, &member_list)
		if err != nil {
			return nil, err
		}
		artistGroups.Members, err = utility.SplitList(member_list)
		if err != nil {
			return nil, err
		}
		artists[artistGroups.ID] = &artistGroups
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (db *SQLiteDB) GetArtistGroupsByID(artistIDs []int) (map[int]*model.ArtistGroup, error) {
	in_placeholder, in_args := build_in_args(artistIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, name, members FROM view_artist_groups WHERE id IN (%s)", in_placeholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, in_args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	artists := make(map[int]*model.ArtistGroup)
	for rows.Next() {
		var artistGroups model.ArtistGroup
		var member_list sql.NullString
		err := rows.Scan(&artistGroups.ID, &artistGroups.Name, &member_list)
		if err != nil {
			return nil, err
		}
		artistGroups.Members, err = utility.SplitList(member_list)
		if err != nil {
			return nil, err
		}
		artists[artistGroups.ID] = &artistGroups
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (db *SQLiteDB) CreateArtistGroups(artists []*model.ArtistGroup) (map[int]*model.ArtistGroup, error) {
	insertArtist := "INSERT INTO artists (name, is_group) VALUES (?, ?);"
	insertMember := "INSERT INTO map_artist_group_members (artist_group, member) VALUES (?, ?);"

	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}

	stmtArt, err := tx.Prepare(insertArtist)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmtArt.Close()

	stmtMem, err := tx.Prepare(insertMember)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmtMem.Close()

	artists_w_id := make(map[int]*model.ArtistGroup)
	for _, artist := range artists {
		result, err := stmtArt.Exec(artist.Name, 1)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		artist.ID = int(id)
		artists_w_id[artist.ID] = artist

		// insert the members
		for _, m := range artist.Members {
			_, err := stmtMem.Exec(artist.ID, m)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return artists_w_id, nil
}
