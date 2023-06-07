package mock

import (
	"github.com/lea-video/go-lea-music/model"
)

func (db *LEAMockDB) GetArtists() (map[int]*model.OneOfArtist, error) {
	return map[int]*model.OneOfArtist{
		1: model.ArtistGroup{ID: 1, Name: "G1", Members: []int{2}}.ToOneOf(),
		2: model.ArtistSolo{ID: 2, Name: "S1"}.ToOneOf(),
	}, nil
}

func (db *LEAMockDB) GetArtistsByID([]int) (map[int]*model.OneOfArtist, error) {
	return db.GetArtists()
}

func (db *LEAMockDB) GetArtistSolos() (map[int]*model.ArtistSolo, error) {
	return map[int]*model.ArtistSolo{
		2: {ID: 2, Name: "S1"},
	}, nil
}

func (db *LEAMockDB) CreateArtistSolos([]*model.ArtistSolo) (map[int]*model.ArtistSolo, error) {
	return db.GetArtistSolos()
}

func (db *LEAMockDB) GetArtistGroups() (map[int]*model.ArtistGroup, error) {
	return map[int]*model.ArtistGroup{
		1: {ID: 1, Name: "G1", Members: []int{2}},
	}, nil
}

func (db *LEAMockDB) CreateArtistGroups([]*model.ArtistGroup) (map[int]*model.ArtistGroup, error) {
	return db.GetArtistGroups()
}

func (db *LEAMockDB) AddArtistGroupMembers(int, []int) error {
	return nil
}
