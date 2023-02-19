package model

type SongVariation struct {
	ID      int       `gorm:"primaryKey"`
	Song    *Song     `gorm:"not null;foreignKey:ID"`
	Artists []*Artist `gorm:"many2many:link_artist_to_song_variation;"`
	Tags    []*Tag    `gorm:"many2many:link_tags_to_song_variations;"`
	FileID  int       `gorm:"not null"`
	File    *File     `gorm:"not null;references:FileID;foreignKey:ID"`
}
