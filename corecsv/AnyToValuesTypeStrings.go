package corecsv

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// AnyToValuesTypeStrings
//
// Output : []{ 'value - type', 'value - type', ... }
func AnyToValuesTypeStrings(
	references ...interface{},
) []string {
	if len(references) == 0 {
		return []string{}
	}

	slice := make([]string, len(references))

	for i, item := range references {
		finalString := toString(item)

		if finalString == "" {
			slice[i] = fmt.Sprintf(
				constants.TypeWithSingleQuoteFormat,
				item)

			continue
		}

		slice[i] = fmt.Sprintf(
			constants.StringTypeFormat,
			toString(item),
			item)
	}

	return slice
}
