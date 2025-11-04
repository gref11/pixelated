package regexps

import (
	"regexp"
)

var (
	ChunkNameRegexp = regexp.MustCompile("c.(-?[0-9]+).(-?[0-9]+)")
)
