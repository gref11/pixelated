package utils

import (
	"errors"
	"fmt"
	"strconv"

	"pixelated/internal/models"
	"pixelated/internal/regexps"
)

func GetChunkName(row int, column int) string {
	chunkName := "c." + strconv.Itoa(row) + "." + strconv.Itoa(column)
	return chunkName
}

func GetChunkCoords(chunkName string) (int, int, error) {
	if ischunkName := regexps.ChunkNameRegexp.MatchString(chunkName); !ischunkName {
		return 0, 0, errors.New("cannot get chunk coords: incorrect chunk name")
	}

	matches := regexps.ChunkNameRegexp.FindStringSubmatch(chunkName)

	row, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get chunk coords: %w", err)
	}
	column, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get chunk coords: %w", err)
	}

	return row, column, nil
}

func printPixel(pixel models.Pixel) {
	if pixel.ColorID < 0 || pixel.ColorID > 7 {
		pixel.ColorID = 9
	}

	fmt.Printf("\x1b[%dm  \x1b[0m", 40+pixel.ColorID)
}

func PrintChunk(chunk models.Chunk) {
	for i := range len(chunk.Pixels) {
		for j := range len(chunk.Pixels[i]) {
			printPixel(chunk.Pixels[i][j])
		}
		fmt.Println()
	}
}

func HexToRgb(hexColor string) ([3]int64, error) {
	var rgbColor [3]int64

	for i := range 3 {
		val, err := strconv.ParseInt(hexColor[1+i*2:3+i*2], 16, 32)
		if err != nil {
			return [3]int64{}, err
		}
		rgbColor[i] = val
	}

	return rgbColor, nil
}
