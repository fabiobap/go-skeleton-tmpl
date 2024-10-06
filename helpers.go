package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

// getExecutableDirectory returns the directory of the executable as a listable URI
func GetExecutableDirectory() (fyne.ListableURI, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	absoluteDirPath := filepath.Dir(executablePath)
	return storage.ListerForURI(storage.NewFileURI(absoluteDirPath))
}
