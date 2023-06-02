package filebackends

import (
	"github.com/lea-video/go-lea-music/db"
)

type LocalFileDB struct {
	rootDir string
}

func InitLocalFileDB(rootDir string) (db.GenericFileDB, error) {
	return &LocalFileDB{
		rootDir: rootDir,
	}, nil
}
