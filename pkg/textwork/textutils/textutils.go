package textutils

import "strings"

func LineToString(line []string) string {
	return strings.Join(line, " ")
}
