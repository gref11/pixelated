package storage

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"pixelated/internal/models"
	"pixelated/internal/utils"
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
			return FileStorage{}, fmt.Errorf("cannot init file storage: folder not found")
		}
		return FileStorage{}, fmt.Errorf("cannot init file storage: %w", err)
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
			chunkFilePath := filepath.Join(folderPath, utils.GetChunkName(row, column)+".pix")
			if _, err := os.Stat(chunkFilePath); err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					chunkFile, err := os.Create(chunkFilePath)
					if err != nil {
						return FileStorage{}, fmt.Errorf("err: cannot create chunk file: %w", err)
					}
					defer chunkFile.Close()

					err = binary.Write(chunkFile, binary.LittleEndian, fileStorage.Chunks[row][column])
					if err != nil {
						return FileStorage{}, fmt.Errorf("err: cannot write chunk file: %w", err)
					}
				} else {
					return FileStorage{}, fmt.Errorf("cannot open chunk file: %w", err)
				}
			} else {
				chunkFile, err := os.Open(chunkFilePath)
				if err != nil {
					return FileStorage{}, fmt.Errorf("cannot open chunk file: %w", err)
				}
				defer chunkFile.Close()

				err = binary.Read(chunkFile, binary.LittleEndian, &fileStorage.Chunks[row][column])
				if err != nil {
					return FileStorage{}, fmt.Errorf("cannot read chunk file: %w", err)
				}
			}
		}
	}

	return fileStorage, nil
}
