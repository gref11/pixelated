package storage

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"pixelated/internal/models"
	"pixelated/scripts"
)

type FileStorage struct {
	FolderPath   string
	Chunks       [][]models.Chunk
	RowChunks    int
	ColumnChunks int
}

func NewFileStorage(folderPath string, rowChunks int, columnChunks int) (FileStorage, error) {
	if _, err := os.Stat(folderPath); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Errorf("File storage folder not found error")
		}
		return FileStorage{}, err
	}

	fileStorage := FileStorage{
		FolderPath:   folderPath,
		Chunks:       make([][]models.Chunk, rowChunks),
		RowChunks:    rowChunks,
		ColumnChunks: columnChunks,
	}

	chunksOneDimArray := make([]models.Chunk, rowChunks*columnChunks) // fileStorage.Chunks[i][j] = chunksOneDimArray[i * columnChunks + j];

	for row := 0; row < rowChunks; row++ {
		fileStorage.Chunks[row] = chunksOneDimArray[row*columnChunks : (row+1)*columnChunks]

		for column := 0; column < columnChunks; column++ {
			chunkFilePath := filepath.Join(folderPath, scripts.GetChunkName(row, column)+".pix")
			if _, err := os.Stat(chunkFilePath); err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					chunkFile, err := os.Create(chunkFilePath)
					if err != nil {
						fmt.Printf("Err: cannot create chunk file: %v\n", err)
						continue
					}
					defer chunkFile.Close()

					err = binary.Write(chunkFile, binary.LittleEndian, fileStorage.Chunks[row][column])
					if err != nil {
						fmt.Printf("Err: cannot write chunk file: %v\n", err)
						continue
					}
				} else {
					fmt.Printf("Err: %v\n", err)
				}
			} else {
				chunkFile, err := os.Open(chunkFilePath)
				if err != nil {
					fmt.Printf("Err: cannot create chunk file: %v\n", err)
					continue
				}
				defer chunkFile.Close()

				err = binary.Read(chunkFile, binary.LittleEndian, &fileStorage.Chunks[row][column])
				if err != nil {
					fmt.Printf("Err: cannot read chunk file: %v\n", err)
					continue
				}
			}
		}
	}

	return fileStorage, nil
}
