package utility

import "os"

// FileExists returns true if the file with the given filename exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
