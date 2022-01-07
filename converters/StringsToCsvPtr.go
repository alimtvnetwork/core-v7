package converters

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/simplewrap"
)

func StringsToCsvPtr(isSkipQuoteOnlyOnExistence bool, stringsSlice *[]string) string {
	if stringsSlice == nil {
		return ""
	}

	csvLines := simplewrap.DoubleQuoteWrapElements(
		isSkipQuoteOnlyOnExistence,
		*stringsSlice...,
	)

	return strings.Join(csvLines, constants.Comma)
}
