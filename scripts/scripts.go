package scripts

import (
	"strconv"
)

func GetChunkName(row int, column int) string {
	chunkName := "c." + strconv.Itoa(row) + "." + strconv.Itoa(column)
	return chunkName
}
