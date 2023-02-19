package main

import (
	"github.com/lea-video/go-lea-music/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	panicOn(err)

	// Migrate the schema
	err = model.AutoMigrate(db)
	panicOn(err)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
