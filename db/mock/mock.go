package mock

import (
	"github.com/lea-video/go-lea-music/db"
)

type MockDB struct{}

func InitMockDB() (db.GenericDB, error) {
	return &MockDB{}, nil
}
