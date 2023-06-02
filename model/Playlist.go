package model

// NOTE: unused
type Playlist struct {
	ID    int
	Name  string
	Items []int
}

// NOTE: unused
type PlaylistItem struct {
	ID       int
	Position int

	Media    int
	Clip     int
	Playlist int
}
