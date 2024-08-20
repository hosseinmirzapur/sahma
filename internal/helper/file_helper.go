package helper

import (
	"fmt"
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

func SaveFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to the file: %w", err)
	}

	return nil
}

func Mkdir(path string) error {
	return os.Mkdir(path, 0777)
}
