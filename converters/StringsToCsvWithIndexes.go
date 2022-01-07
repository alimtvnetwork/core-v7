package converters

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/simplewrap"
)

func StringsToCsvWithIndexes(stringsSlice []string) string {
	csvLines := simplewrap.DoubleQuoteWrapElementsWithIndexes(
		stringsSlice...,
	)

	return strings.Join(csvLines, constants.Comma)
}
