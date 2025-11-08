package storage

import "pixelated/internal/models"

type Storage interface {
	GetChunkByID(chunkID string) (models.Chunk, error)
	GetAllChunks() ([][]models.Chunk, error)
	UpdateChunk(chunkID string, changes models.Chunk) error
	UpdateAllChunks(changes map[string]models.Chunk) error
}
