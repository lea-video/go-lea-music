package gorm

import (
	"errors"
	"github.com/lea-video/go-lea-music/model"
)

func (mg *myGORM) ListSongVariations() ([]model.SongVariation, error) {
	return nil, errors.New("not yet implemented")
}
func (mg *myGORM) FetchSongVariations(ids []int) ([]model.SongVariation, error) {
	return nil, errors.New("not yet implemented")
}
func (mg *myGORM) DeleteSongVariations(ids []int) error {
	return errors.New("not yet implemented")
}
func (mg *myGORM) UpdateVariations([]model.SongVariation) ([]model.SongVariation, error) {
	return nil, errors.New("not yet implemented")
}
func (mg *myGORM) CreateVariations(variations []model.SongVariation) ([]model.SongVariation, error) {
	return nil, errors.New("not yet implemented")
}
