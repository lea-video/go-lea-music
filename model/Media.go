package model

import "github.com/lea-video/go-lea-music/utility"

type MediaTrack struct {
	ID          int
	ParentMedia int
	HasAudio    bool
	HasVideo    bool
	HasPicture  bool
}

type Media struct {
	ID     int
	Tracks []int
	Source string
}

type TMPFile struct {
	ID          int
	ParentMedia int
	AccessToken string
	Location    string `json:"-"`
	MaxAge      int
}

func NewTMPFile(parentID int) (*TMPFile, error) {
	at, err := utility.GenerateRandomStringURLSafe(64)
	if err != nil {
		return nil, err
	}
	loc, err := utility.GenerateRandomStringURLSafe(64)
	if err != nil {
		return nil, err
	}
	return &TMPFile{
		ParentMedia: parentID,
		AccessToken: at,
		Location:    loc,
		MaxAge:      1337,
	}, nil
}

// NOTE: unused
type Clip struct {
	ID    int
	Media int
	Start int
	End   int
}
