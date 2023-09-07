package stringutil

import (
	"gitlab.com/auk-go/core/constants"
)

// SafeSubstringStarts
//
// content[startAt:]
// -1 meaning get full text
func SafeSubstringStarts(
	content string,
	startAt int,
) string {
	length := len(content)

	if length == 0 || length <= startAt {
		return constants.EmptyString
	}

	if startAt == -1 {
		return content
	}

	return content[startAt:]
}
