package mock

import "github.com/lea-video/go-lea-music/model"

func (db *LEAMockDB) GetPlaylists() (map[int]*model.Playlist, error) {
	return map[int]*model.Playlist{
		9:  {ID: 9, Name: "Test Playlist", Items: []int{3, 4}},
		10: {ID: 10, Name: "Another Test Playlist", Items: []int{3}},
	}, nil
}

func (db *LEAMockDB) CreatePlaylists([]*model.Playlist) (map[int]*model.Playlist, error) {
	return db.GetPlaylists()
}

func (db *LEAMockDB) AppendPlaylistItems(int, []int) error {
	return nil
}

func (db *LEAMockDB) ChangePlaylistItemPosition(int, int, int) error {
	return nil
}

func (db *LEAMockDB) GetPlaylistsByID([]int) (map[int]*model.Playlist, error) {
	return db.GetPlaylists()
}
