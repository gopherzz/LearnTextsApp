package worderaser

import (
	"math/rand"
	"strings"
)

// █

func EraseWord(str string) string {
	str = strings.Repeat("█", len(str))
	return str
}

func EraseRandomWordsByLevel(line []string, level uint) []string {
	var chance uint
	var res []string = line
	lineLen := uint(len(res))
	if level > lineLen {
		chance = uint(lineLen)
	}
	chance = lineLen / level
	for i := 0; uint(i) < chance; i++ {
		wordIndex := rand.Intn(int(lineLen))
		res[wordIndex] = EraseWord(strings.TrimSpace(res[wordIndex]))
	}
	return line
}
