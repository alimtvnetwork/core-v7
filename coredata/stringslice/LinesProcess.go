package stringslice

import (
	"gitlab.com/auk-go/core/constants"
)

// LinesProcess split text using constants.NewLineUnix
// then returns lines processed by lineProcessor
func LinesProcess(
	lines []string,
	lineProcessor func(index int, lineIn string) (lineOut string, isTake, isBreak bool),
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	slice := Make(constants.Zero, len(lines))

	for i, lineIn := range lines {
		lineOut, isTake, isBreak := lineProcessor(i, lineIn)

		if isTake {
			slice = append(slice, lineOut)
		}

		if isBreak {
			break
		}
	}

	return slice
}
