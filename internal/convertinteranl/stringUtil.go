package convertinteranl

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type stringUtil struct{}

func (it stringUtil) PrependWithSpaces(
	joiner string,
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) string {
	toSlice := Util.Strings.PrependWithSpaces(
		spaceCountLines,
		existingLines,
		prependingLinesSpaceCount,
		prependingLines...,
	)

	return strings.Join(toSlice, joiner)
}

func (it stringUtil) PrependWithSpacesDefault(
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) string {
	toSlice := Util.Strings.PrependWithSpaces(
		spaceCountLines,
		existingLines,
		prependingLinesSpaceCount,
		prependingLines...,
	)

	return strings.Join(toSlice, constants.NewLineUnix)
}

// IndexToPosition
//
// Index to position format
//
//   - Index 0 => Position => 1st
//   - Index 1 => Position => 2nd
//   - Index 3 => Position => 3rd
//   - Index Rest => Position => %dth
func (it stringUtil) IndexToPosition(
	index int,
) string {
	position := index + 1

	switch position {
	case 1:
		return "1st"
	case 2:
		return "2nd"
	case 3:
		return "3rd"
	default:
		return fmt.Sprintf(
			"%dth",
			position,
		)
	}
}

func (it stringUtil) PascalCase(
	name string,
) string {
	if len(name) == 0 {
		return ""
	}

	allRunes := []rune(name)
	firstChar := allRunes[0]
	firstCharStr := string(firstChar)
	firstCharUpper := strings.ToUpper(firstCharStr)

	if len(allRunes) == 1 {
		return firstCharUpper
	}

	return firstCharUpper + string(allRunes[1:])
}

func (it stringUtil) CamelCase(
	name string,
) string {
	if len(name) == 0 {
		return ""
	}

	allRunes := []rune(name)
	firstChar := allRunes[0]
	firstCharStr := string(firstChar)
	firstCharLower := strings.ToLower(firstCharStr)

	if len(allRunes) == 1 {
		return firstCharLower
	}

	return firstCharLower + string(allRunes[1:])
}
