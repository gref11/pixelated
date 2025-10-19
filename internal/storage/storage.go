package storage

type Storage interface {
	GetChunkByID(chunkID string) (string, error)
	GetAllChunks() ([]string, error)
	UpdateChunk(chunkID string, changes string) error
}