package main

import (
	db_sqlite "github.com/lea-video/go-lea-music/db/sqlite"
	"github.com/lea-video/go-lea-music/filebackends"
)

func main() {
	// db, err := db_mock.InitMockDB()
	db, err := db_sqlite.InitSQLite("test.sqlite")
	panicOn(err)

	fileDB, err := filebackends.InitLocalFileDB("./files/")
	panicOn(err)

	// start rest server
	app := prepServer()
	addRoutes(app, db, fileDB)
	startServer(app)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
