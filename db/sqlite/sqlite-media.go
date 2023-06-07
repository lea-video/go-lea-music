package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/lea-video/go-lea-music/model"
	"github.com/lea-video/go-lea-music/utility"
)

func (db *LEASQLiteDB) GetMedia() (map[int]*model.Media, error) {
	rows, err := db.db.Query("SELECT id, source, tracks FROM view_media_tracks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allMedia := make(map[int]*model.Media)
	for rows.Next() {
		var media model.Media
		var trackList sql.NullString
		err := rows.Scan(&media.ID, &media.Source, &trackList)
		if err != nil {
			return nil, err
		}
		media.Tracks, err = utility.SplitList(trackList)
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

func (db *LEASQLiteDB) GetMediaByID(mediaIDs []int) (map[int]*model.Media, error) {
	inPlaceholder, inArgs := buildINStatement(mediaIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, source, tracks FROM view_media_tracks WHERE parentMedia IN (%s)", inPlaceholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, inArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allMedia := make(map[int]*model.Media)
	for rows.Next() {
		var media model.Media
		var trackList sql.NullString
		err := rows.Scan(&media.ID, &media.Source, &trackList)
		if err != nil {
			return nil, err
		}
		media.Tracks, err = utility.SplitList(trackList)
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

func (db *LEASQLiteDB) CreateMedia(media []*model.Media) (map[int]*model.Media, error) {
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

	mediaWithID := make(map[int]*model.Media)
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
		mediaWithID[m.ID] = m
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return mediaWithID, nil
}

func (db *LEASQLiteDB) GetMediaTracks(mediaIDs []int) (map[int]*model.MediaTrack, error) {
	inPlaceholder, inArgs := buildINStatement(mediaIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, parentMedia, has_audio, has_video, has_picture FROM media_track WHERE parentMedia IN (%s)", inPlaceholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, inArgs...)
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

func (db *LEASQLiteDB) GetMediaTracksByID(trackIDs []int) (map[int]*model.MediaTrack, error) {
	inPlaceholder, inArgs := buildINStatement(trackIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, parentMedia, has_audio, has_video, has_picture FROM media_track WHERE id IN (%s)", inPlaceholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, inArgs...)
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

func (db *LEASQLiteDB) CreateTMPFiles(files []*model.TMPFile) (map[int]*model.TMPFile, error) {
	insertStmt := "INSERT INTO tmp_file (parentMedia, accessToken, location, maxAge) VALUES (?, ?, ?, ?);"

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

	filesWithID := make(map[int]*model.TMPFile)
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
		filesWithID[file.ID] = file
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return filesWithID, nil
}

func (db *LEASQLiteDB) GetTMPFileByID(fileIDs []int) (map[int]*model.TMPFile, error) {
	inPlaceholder, inArgs := buildINStatement(fileIDs)

	// Construct the query with the placeholders
	query := fmt.Sprintf("SELECT id, parentMedia, accessToken, location, maxAge FROM tmp_file WHERE parentMedia IN (%s)", inPlaceholder)

	// Execute the query with the medias slice
	rows, err := db.db.Query(query, inArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allFiles := make(map[int]*model.TMPFile)
	for rows.Next() {
		var tmpFile model.TMPFile
		err := rows.Scan(&tmpFile.ID, &tmpFile.ParentMedia, &tmpFile.AccessToken, &tmpFile.Location, &tmpFile.MaxAge)
		if err != nil {
			return nil, err
		}
		allFiles[tmpFile.ID] = &tmpFile
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return allFiles, nil
}
