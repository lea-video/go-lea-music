package model

type SongVariation struct {
	ID     int
	SongID int
	Song   *Song
	FileID int
	File   *File
	// Artists []*Artist `gorm:"many2many:link_artist_to_song_variation"`
	// Tags    []*Tag    `gorm:"many2many:link_tags_to_song_variations;"`
}
