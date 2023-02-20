package gorm

import (
	"errors"
	"github.com/lea-video/go-lea-music/model"
)

func (mg *myGORM) ListSongs() ([]model.Song, error) {
	return nil, errors.New("not yet implemented")
}
func (mg *myGORM) FetchSongs(ids []int) ([]model.Song, error) {
	return nil, errors.New("not yet implemented")
}
func (mg *myGORM) DeleteSongs(ids []int) error {
	return errors.New("not yet implemented")
}
func (mg *myGORM) UpdateSongs([]model.Song) ([]model.Song, error) {
	return nil, errors.New("not yet implemented")
}
func (mg *myGORM) CreateSong(songs []model.Song) ([]model.Song, error) {
	return nil, errors.New("not yet implemented")
}
