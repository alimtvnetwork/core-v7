package converters

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/simplewrap"
)

func StringsToCsv(isSkipQuoteOnlyOnExistence bool, stringsSlice ...string) string {
	csvLines := simplewrap.DoubleQuoteWrapElements(
		isSkipQuoteOnlyOnExistence,
		stringsSlice...)

	return strings.Join(csvLines, constants.Comma)
}
