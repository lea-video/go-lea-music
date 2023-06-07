package db

import "github.com/lea-video/go-lea-music/model"

type GenericDB interface {
	GetArtists() (map[int]*model.OneOfArtist, error)
	GetArtistsByID([]int) (map[int]*model.OneOfArtist, error)
	GetArtistSolos() (map[int]*model.ArtistSolo, error)
	CreateArtistSolos([]*model.ArtistSolo) (map[int]*model.ArtistSolo, error)
	GetArtistGroups() (map[int]*model.ArtistGroup, error)
	CreateArtistGroups([]*model.ArtistGroup) (map[int]*model.ArtistGroup, error)
	AddArtistGroupMembers(int, []int) error

	GetMedia() (map[int]*model.Media, error)
	GetMediaByID([]int) (map[int]*model.Media, error)
	CreateMedia([]*model.Media) (map[int]*model.Media, error)
	GetMediaTracks([]int) (map[int]*model.MediaTrack, error)
	GetMediaTracksByID([]int) (map[int]*model.MediaTrack, error)

	CreateTMPFiles([]*model.TMPFile) (map[int]*model.TMPFile, error)
	GetTMPFileByID([]int) (map[int]*model.TMPFile, error)

	GetPlaylists() (map[int]*model.Playlist, error)
	CreatePlaylists([]*model.Playlist) (map[int]*model.Playlist, error)
	AppendPlaylistItems(int, []int) error
	ChangePlaylistItemPosition(int, int, int) error
	GetPlaylistsByID([]int) (map[int]*model.Playlist, error)
}
