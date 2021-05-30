package textutils

import "strings"

// Convert Line([]string) to string
func LineToString(line []string) string {
	return strings.Join(line, " ")
}
