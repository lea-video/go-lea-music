package gorm

import (
	"errors"
	"github.com/lea-video/go-lea-music/model"
)

type DBArtist struct {
	ID   int
	Type string
	Name string
}

func (dba DBArtist) Conv() model.Artist {
	if dba.Type == model.TypeSolo {
		return model.SoloArtist{
			CoreArtist: model.CoreArtist{
				ID:   dba.ID,
				Type: dba.Type,
				Name: dba.Name,
			},
		}
	} else {
		return model.GroupArtist{
			CoreArtist: model.CoreArtist{
				ID:   dba.ID,
				Type: dba.Type,
				Name: dba.Name,
			},
		}
	}
}

func ConvArtist(in model.Artist) DBArtist {
	return DBArtist{
		ID:   in.GetCore().ID,
		Type: in.GetCore().Type,
		Name: in.GetCore().Name,
	}
}

func (mg *myGORM) ListArtists() ([]model.Artist, error) {
	var artists []DBArtist
	results := mg.db.Find(&artists)
	if results.Error != nil {
		return nil, results.Error
	}
	x := make([]model.Artist, len(artists))
	for idx, art := range artists {
		x[idx] = art.Conv()
	}
	return x, nil
}

func (mg *myGORM) FetchArtists(ids []int) ([]model.Artist, error) {
	return nil, errors.New("not yet implemented")
}

func (mg *myGORM) DeleteArtists(ids []int) error {
	return errors.New("not yet implemented")
}

func (mg *myGORM) UpdateArtists([]model.Artist) ([]model.Artist, error) {
	return nil, errors.New("not yet implemented")
}

func (mg *myGORM) CreateArtists(artists []model.Artist) ([]model.Artist, error) {
	x := make([]model.Artist, len(artists))

	for idx, art := range artists {
		artist := ConvArtist(art)
		results := mg.db.Create(&artist)
		if results.Error != nil {
			return nil, results.Error
		}
		x[idx] = artist.Conv()
	}

	return x, nil
}
