package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"pixelated/internal/models"
	"pixelated/internal/utils"
)

func TestNewFileStorage(t *testing.T) {
	absolutePath, _ := filepath.Abs("")
	folderPath := filepath.Join(absolutePath, "test_data")
	os.Mkdir(folderPath, 0777)
	defer os.RemoveAll(folderPath)

	rows := 3
	columns := 3
	chunks := make([][]models.Chunk, rows)

	for row := range rows {
		chunks[row] = make([]models.Chunk, columns)
	}

	actual, err := NewFileStorage(folderPath, rows, columns)

	assert.Nil(t, err)

	assert.Equal(t, folderPath, actual.FolderPath)
	assert.Equal(t, chunks, actual.Chunks)
	assert.Equal(t, rows, actual.RowChunks)
	assert.Equal(t, columns, actual.ColumnChunks)
}

func TestNewFileStorageNoFileExistsErr(t *testing.T) {
	folderPath := "not_existing_folder"
	rows := 1
	columns := 1

	expectedFolderPath := ""
	expectedRows := 0
	expectedColumns := 0
	expectedChunks := [][]models.Chunk(nil)

	actual, err := NewFileStorage(folderPath, rows, columns)

	assert.NotNil(t, err)

	assert.Equal(t, expectedFolderPath, actual.FolderPath)
	assert.Equal(t, expectedRows, actual.RowChunks)
	assert.Equal(t, expectedColumns, actual.ColumnChunks)
	assert.Equal(t, expectedChunks, actual.Chunks)
}

func TestGetChunkByID(t *testing.T) {
	absolutePath, _ := filepath.Abs("")
	folderPath := filepath.Join(absolutePath, "test_data")
	os.Mkdir(folderPath, 0777)
	defer os.RemoveAll(folderPath)

	rows := 1
	columns := 1
	chunkID := utils.GetChunkName(0, 0)
	expectedChunk := models.Chunk{}

	fileStorage, err := NewFileStorage(folderPath, rows, columns)

	assert.Nil(t, err)

	actual, err := fileStorage.GetChunkByID(chunkID)

	assert.Nil(t, err)

	assert.Equal(t, expectedChunk, actual)
}
