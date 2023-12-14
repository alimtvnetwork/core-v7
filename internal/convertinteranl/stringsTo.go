package convertinteranl

import (
	"fmt"
	"strings"
)

type stringsTo struct{}

// WithSpaces
//
// Generates lines with space given in the count.
func (it stringsTo) WithSpaces(spaceCount int, lines ...string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	newLines := make([]string, len(lines))
	prefix := strings.Repeat(
		" ",
		spaceCount,
	)

	for i, line := range lines {
		newLines[i] = fmt.Sprintf(
			"%s%s",
			prefix,
			line,
		)
	}

	return newLines
}
