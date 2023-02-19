package model

type Song struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Tags []*Tag `gorm:"many2many:link_tags_to_song;"`
}
