package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/lea-video/go-lea-music/model"
	"github.com/lea-video/go-lea-music/utility"
)

func (db *SQLiteDB) GetMedia() (map[int]*model.Media, error) {
	rows, err := db.db.Query("SELECT id, source, tracks FROM view_media_tracks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allMedia := make(map[int]*model.Media)
	for rows.Next() {
		var media model.Media
		var track_list sql.NullString
		err := rows.Scan(&media.ID, &media.Source, &track_list)
		if err != nil {
			return nil, err
		}
		media.Tracks, err = utility.SplitList(track_list)
		if err != nil {
			return nil, err
		}
		allMedia[media.ID] = &media
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allMedia, nil
}

func (db *SQLiteDB) GetMediaByID(mediaIDs []int) (map[int]*model.Media, error) {
	in_placeholder, in_args := build_in_args(mediaIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, source, tracks FROM view_media_tracks WHERE parentMedia IN (%s)", in_placeholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, in_args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allMedia := make(map[int]*model.Media)
	for rows.Next() {
		var media model.Media
		var track_list sql.NullString
		err := rows.Scan(&media.ID, &media.Source, &track_list)
		if err != nil {
			return nil, err
		}
		media.Tracks, err = utility.SplitList(track_list)
		if err != nil {
			return nil, err
		}
		allMedia[media.ID] = &media
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allMedia, nil
}

func (db *SQLiteDB) CreateMedia(media []*model.Media) (map[int]*model.Media, error) {
	insertStmt := "INSERT INTO media (source) VALUES (?);"

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

	media_w_id := make(map[int]*model.Media)
	for _, m := range media {
		result, err := stmt.Exec(m.Source)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		// TODO: insert tracks

		m.ID = int(id)
		media_w_id[m.ID] = m
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return media_w_id, nil
}

func (db *SQLiteDB) GetMediaTracks(mediaIDs []int) (map[int]*model.MediaTrack, error) {
	in_placeholder, in_args := build_in_args(mediaIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, parentMedia, has_audio, has_video, has_picture FROM mediatrack WHERE parentMedia IN (%s)", in_placeholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, in_args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allTracks := make(map[int]*model.MediaTrack)
	for rows.Next() {
		var track model.MediaTrack
		err := rows.Scan(&track.ID, &track.ParentMedia, &track.HasAudio, &track.HasVideo, &track.HasPicture)
		if err != nil {
			return nil, err
		}
		allTracks[track.ID] = &track
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allTracks, nil
}

func (db *SQLiteDB) GetMediaTracksByID(trackIDs []int) (map[int]*model.MediaTrack, error) {
	in_placeholder, in_args := build_in_args(trackIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, parentMedia, has_audio, has_video, has_picture FROM mediatrack WHERE id IN (%s)", in_placeholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, in_args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allTracks := make(map[int]*model.MediaTrack)
	for rows.Next() {
		var track model.MediaTrack
		err := rows.Scan(&track.ID, &track.ParentMedia, &track.HasAudio, &track.HasVideo, &track.HasPicture)
		if err != nil {
			return nil, err
		}
		allTracks[track.ID] = &track
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allTracks, nil
}

func (db *SQLiteDB) CreateTMPFiles(files []*model.TMPFile) (map[int]*model.TMPFile, error) {
	insertStmt := "INSERT INTO tmpfile (parentMedia, accessToken, location, maxage) VALUES (?, ?, ?, ?);"

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

	files_w_id := make(map[int]*model.TMPFile)
	for _, file := range files {
		result, err := stmt.Exec(file.ParentMedia, file.AccessToken, file.Location, file.MaxAge)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		file.ID = int(id)
		files_w_id[file.ID] = file
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return files_w_id, nil
}

func (db *SQLiteDB) GetTMPFileByID([]int) (map[int]*model.TMPFile, error) {
	rows, err := db.db.Query("SELECT id, parentMedia, accessToken, location, maxage FROM tmpfile")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allFiles := make(map[int]*model.TMPFile)
	for rows.Next() {
		var tmpfile model.TMPFile
		err := rows.Scan(&tmpfile.ID, &tmpfile.ParentMedia, &tmpfile.AccessToken, &tmpfile.Location, &tmpfile.MaxAge)
		if err != nil {
			return nil, err
		}
		allFiles[tmpfile.ID] = &tmpfile
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allFiles, nil
}
