package model

type Tag struct {
	ID                     int    `gorm:"primaryKey"`
	Name                   string `gorm:"not null"`
	AutomaticallyGenerated bool   `gorm:"not null"`
	ParentID               int
	Parent                 *Tag             `gorm:"references:ParentID;foreignKey:ID"`
	Artists                []*Artist        `gorm:"many2many:link_tags_to_artist;"`
	Songs                  []*Song          `gorm:"many2many:link_tags_to_song;"`
	SongVariations         []*SongVariation `gorm:"many2many:link_tags_to_song_variations;"`
}
