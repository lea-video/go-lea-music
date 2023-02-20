package model

type Song struct {
	ID   int
	Name string
	// Tags []*Tag `gorm:"many2many:link_tags_to_song;"`
}
