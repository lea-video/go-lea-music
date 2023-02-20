package routes

import (
	"github.com/lea-video/go-lea-music/model"
)

type GenericDB interface {
	ListArtists() ([]model.Artist, error)
	FetchArtists(ids []int) ([]model.Artist, error)
	DeleteArtists(ids []int) error
	UpdateArtists([]model.Artist) ([]model.Artist, error)
	CreateArtists(artists []model.Artist) ([]model.Artist, error)

	ListSongs() ([]model.Song, error)
	FetchSongs(ids []int) ([]model.Song, error)
	DeleteSongs(ids []int) error
	UpdateSongs([]model.Song) ([]model.Song, error)
	CreateSong(songs []model.Song) ([]model.Song, error)

	ListSongVariations() ([]model.SongVariation, error)
	FetchSongVariations(ids []int) ([]model.SongVariation, error)
	DeleteSongVariations(ids []int) error
	UpdateVariations([]model.SongVariation) ([]model.SongVariation, error)
	CreateVariations(variations []model.SongVariation) ([]model.SongVariation, error)
}

type RouteWrapper struct {
	DB GenericDB
}
