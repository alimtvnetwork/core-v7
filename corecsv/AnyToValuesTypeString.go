package corecsv

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// AnyToValuesTypeString
//
// Output : 'value - type', 'value - type'...
func AnyToValuesTypeString(
	references ...interface{},
) string {
	if len(references) == 0 {
		return ""
	}

	toSlice := AnyToValuesTypeStrings(references...)

	return strings.Join(
		toSlice,
		constants.CommaSpace,
	)
}
