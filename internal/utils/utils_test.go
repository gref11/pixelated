package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingChunkName(t *testing.T) {
	row := -123
	column := 456
	expected := "c.-123.456"

	actual := GetChunkName(row, column)

	assert.Equal(t, expected, actual)
}

func TestGettingChunkCoords(t *testing.T) {
	chunkName := "c.-123.456"

	expectedRow := -123
	expectedColumn := 456

	actualRow, actualColumn, err := GetChunkCoords(chunkName)

	assert.Nil(t, err)

	assert.Equal(t, expectedRow, actualRow)
	assert.Equal(t, expectedColumn, actualColumn)
}

func TestGettingChunkCoordsRowError(t *testing.T) {
	chunkName := "c.999999999999.456"

	expectedRow := 0
	expectedColumn := 0

	actualRow, actualColumn, err := GetChunkCoords(chunkName)

	assert.NotNil(t, err)

	assert.Equal(t, expectedRow, actualRow)
	assert.Equal(t, expectedColumn, actualColumn)
}

func TestGettingChunkCoordsColumnError(t *testing.T) {
	chunkName := "c.-123.999999999999"

	expectedRow := 0
	expectedColumn := 0

	actualRow, actualColumn, err := GetChunkCoords(chunkName)

	assert.NotNil(t, err)

	assert.Equal(t, expectedRow, actualRow)
	assert.Equal(t, expectedColumn, actualColumn)
}

func TestGettingChunkCoordsFormatError(t *testing.T) {
	chunkName := "c..456"

	expectedRow := 0
	expectedColumn := 0

	actualRow, actualColumn, err := GetChunkCoords(chunkName)

	assert.NotNil(t, err)

	assert.Equal(t, expectedRow, actualRow)
	assert.Equal(t, expectedColumn, actualColumn)
}
