package worderaser

import (
	"strings"
)

const (
	substitude = "█"
)

// Replace String to substitude string
func EraseWord(str string) string {
	str = strings.Repeat(substitude, len(str))
	return str
}

// Blurs words by level(count words)
func EraseWordsByLevel(line []string, level uint) []string {
	var res []string = line
	if level >= uint(len(line)) {
		level = uint(len(line))
	}
	for i := 0; uint(i) < level; i++ {
		res[i] = EraseWord(strings.TrimSpace(res[i]))
	}
	return line
}
