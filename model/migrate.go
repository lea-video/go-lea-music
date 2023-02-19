package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&ArtistSolo{},
		&ArtistGroup{},
		&Song{},
		&SongVariation{},
		&File{},
		&Folder{},
		&Tag{},
	)
}
