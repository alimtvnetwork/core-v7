package stringutil

import "gitlab.com/auk-go/core/constants"

// SafeSubstringEnds
//
// content[:endingLen]
//
// -1 meaning upto the length
func SafeSubstringEnds(
	content string,
	endingLen int,
) string {
	length := len(content)

	if length == 0 {
		return constants.EmptyString
	}

	if length <= endingLen || endingLen == -1 {
		return content[:length]
	}

	return content[:endingLen]
}
