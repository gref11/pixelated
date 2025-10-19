package storage

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type FileStorage struct {
	folderPath string
}

func NewFileStorage(folderPath string) (FileStorage, error) {
	if _ ,err := os.Stat(folderPath); err != nil {
		if errors.Is(err, fs.ErrNotExist){
			fmt.Errorf("File storage folder not found error")
		}
		return FileStorage{}, err
	}

	fileStorage := FileStorage{folderPath: folderPath}
	
}