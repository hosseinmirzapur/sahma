package helper

import (
	"os"
	"path/filepath"
)

// Checks for existence of a file
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err != nil || !os.IsNotExist(err)
}

func FileExtension(path string) string {
	if FileExists(path) {
		return filepath.Ext(path)[1:]
	}
	return ""
}

func FileName(path string) string {
	return filepath.Base(path)
}
