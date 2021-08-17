package stringcompareas

import "strings"

// isNotEqualFunc tided with NotEqual
var isNotEqualFunc = func(
	contentLine,
	notEqualText string,
	isIgnoreCase bool,
) bool {
	if isIgnoreCase {
		return !strings.EqualFold(
			notEqualText,
			contentLine)
	}

	return contentLine != notEqualText
}
