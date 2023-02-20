package gorm

import (
	"github.com/lea-video/go-lea-music/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

type myGORM struct {
	db *gorm.DB
}

func (mg *myGORM) AutoMigrate() error {
	return mg.db.AutoMigrate(
		&DBArtist{},
	)
}

func InitSQLite(file string) (routes.GenericDB, error) {
	os.Remove(file)
	// create db
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	mg := &myGORM{db}

	// migrate the db schema
	err = mg.AutoMigrate()
	return mg, err
}
