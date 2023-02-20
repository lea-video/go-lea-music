package main

import "github.com/lea-video/go-lea-music/gorm"

func main() {
	db, err := gorm.InitSQLite("test.sqlite")
	panicOn(err)

	// start rest server
	app := prepServer()
	addRoutes(app, db)
	startServer(app)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
