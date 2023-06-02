package filebackends

import (
	"os"
	"path/filepath"
)

func (db *LocalFileDB) AppendFile(loc string, data []byte) error {
	filePath := filepath.Join(db.rootDir, loc[:2], loc[2:4], loc)

	// Create all necessary directories in the file path
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
