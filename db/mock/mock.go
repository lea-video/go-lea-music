package mock

import (
	"github.com/lea-video/go-lea-music/db"
)

type LEAMockDB struct{}

func InitMockDB() (db.GenericDB, error) {
	return &LEAMockDB{}, nil
}
