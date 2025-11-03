package utils

import (
	"fmt"
	"strconv"

	"pixelated/internal/models"
)

func GetChunkName(row int, column int) string {
	chunkName := "c." + strconv.Itoa(row) + "." + strconv.Itoa(column)
	return chunkName
}

func printPixel(pixel models.Pixel) {
	if pixel.ColorID < 0 || pixel.ColorID > 7 {
		pixel.ColorID = 9
	}

	fmt.Printf("\x1b[%dm  \x1b[0m", 40+pixel.ColorID)
}

func PrintChunk(chunk models.Chunk) {
	for i := 0; i < len(chunk.Pixels); i++ {
		for j := 0; j < len(chunk.Pixels[i]); j++ {
			printPixel(chunk.Pixels[i][j])
		}
		fmt.Println()
	}
}

func HexToRgb(hexColor string) ([3]int64, error) {
	var rgbColor [3]int64

	for i := 0; i < 3; i++ {
		val, err := strconv.ParseInt(hexColor[1+i*2:3+i*2], 16, 32)
		if err != nil {
			return [3]int64{}, err
		}
		rgbColor[i] = val
	}

	return rgbColor, nil
}
