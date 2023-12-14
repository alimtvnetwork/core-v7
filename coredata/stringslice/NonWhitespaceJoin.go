package stringslice

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// NonWhitespaceJoin
//
// Don't include line which is empty or whitespace.
// Reference : NonWhitespace
//
// Join this converted slice to a single string using strings.Join
func NonWhitespaceJoin(slice []string, joiner string) string {
	if slice == nil {
		return constants.EmptyString
	}

	length := len(slice)

	if length == 0 {
		return constants.EmptyString
	}

	return strings.Join(NonWhitespace(slice), joiner)
}
