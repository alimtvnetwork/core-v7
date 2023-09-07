package stringutil

import "gitlab.com/auk-go/core/constants"

// SafeSubstring content[startAt:endAt]
func SafeSubstring(
	content string,
	startAt, endingLength int,
) string {
	if startAt == -1 && endingLength == -1 {
		return content
	}

	length := len(content)

	if length == 0 {
		return content
	}

	if startAt == -1 {
		return SafeSubstringEnds(content, endingLength)
	}

	if endingLength == -1 {
		return SafeSubstringStarts(content, startAt)
	}

	if startAt < length && endingLength <= length && startAt <= endingLength {
		return content[startAt:endingLength]
	}

	return constants.EmptyString
}
