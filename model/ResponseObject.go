package model

type ResponseObject struct {
	Order []int `json:"order,omitempty"`

	Artists map[int]*OneOfArtist `json:"artists,omitempty"`

	Media      map[int]*Media      `json:"media,omitempty"`
	MediaTrack map[int]*MediaTrack `json:"mediatrack,omitempty"`
	TMPFile    map[int]*TMPFile    `json:"tmpfile,omitempty"`

	Playlists   map[int]*Playlist    `json:"playlists,omitempty"`
	Clip        map[int]*Clip        `json:"clips,omitempty"`
	Song        map[int]*Song        `json:"songs,omitempty"`
	Performance map[int]*Performance `json:"performances,omitempty"`
}
