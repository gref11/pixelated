package test

import (
	"pixelated/internal/storage"
	"pixelated/internal/utils"
)

func main() {
	fileStorage, _ := storage.NewFileStorage("C:\\Users\\MI\\OneDrive\\кубгу\\go\\pixelated\\data", 1, 1)
	var testImg = [10][10]int64{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 1, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fileStorage.Chunks[0][0].Pixels[i][j].ColorID = testImg[i][j]
		}
	}

	utils.PrintChunk(fileStorage.Chunks[0][0])
}
