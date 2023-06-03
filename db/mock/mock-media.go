package mock

import (
	"github.com/lea-video/go-lea-music/model"
)

func (db *LEAMockDB) GetMedia() (map[int]*model.Media, error) {
	return map[int]*model.Media{
		3: {ID: 3, Tracks: []int{5, 6, 7}, Source: ""},
		4: {ID: 4, Tracks: []int{6, 7}, Source: ""},
	}, nil
}

func (db *LEAMockDB) GetMediaByID([]int) (map[int]*model.Media, error) {
	return db.GetMedia()
}

func (db *LEAMockDB) CreateMedia([]*model.Media) (map[int]*model.Media, error) {
	return db.GetMedia()
}

func (db *LEAMockDB) GetMediaTracks([]int) (map[int]*model.MediaTrack, error) {
	return map[int]*model.MediaTrack{
		5: {ID: 5, HasAudio: false, HasVideo: false},
		6: {ID: 6, HasAudio: false, HasVideo: true},
		7: {ID: 7, HasAudio: true, HasVideo: false},
	}, nil
}

func (db *LEAMockDB) GetMediaTracksByID([]int) (map[int]*model.MediaTrack, error) {
	return db.GetMediaTracks([]int{})
}

func (db *LEAMockDB) CreateTMPFiles([]*model.TMPFile) (map[int]*model.TMPFile, error) {
	return db.GetTMPFileByID([]int{})
}

func (db *LEAMockDB) GetTMPFileByID([]int) (map[int]*model.TMPFile, error) {
	return map[int]*model.TMPFile{
		8: {ID: 8, ParentMedia: 4, Location: "fa1234", AccessToken: "asdf", MaxAge: 1337},
	}, nil
}
