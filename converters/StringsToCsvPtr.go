package converters

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/simplewrap"
)

func StringsToCsvPtr(isSkipQuoteOnlyOnExistence bool, stringsSlice *[]string) string {
	csvLines := simplewrap.DoubleQuoteWrapElements(
		stringsSlice,
		isSkipQuoteOnlyOnExistence)

	return strings.Join(*csvLines, constants.Comma)
}
