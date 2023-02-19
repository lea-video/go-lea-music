package model

type Artist struct {
	ID             string           `gorm:"primaryKey"`
	Name           string           `gorm:"not null"`
	SongVariations []*SongVariation `gorm:"many2many:link_artist_to_song_variation;"`
	Tags           []*Tag           `gorm:"many2many:link_tags_to_artist;"`
}
