package routes

import "github.com/lea-video/go-lea-music/db"

type RouteWrapper struct {
	DB     db.GenericDB
	FileDB db.GenericFileDB
}
