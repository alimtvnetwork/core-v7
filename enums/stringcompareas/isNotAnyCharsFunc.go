package stringcompareas

import "strings"

// isNotAnyCharsFunc reports whether any
// Unicode code points in chars are NOT within contentLine.
//
// Tided with NotAnyChars
var isNotAnyCharsFunc = func(
	contentLine,
	charsFind string,
	isIgnoreCase bool,
) bool {
	if isIgnoreCase {
		return !strings.ContainsAny(
			strings.ToLower(contentLine),
			strings.ToLower(charsFind),
		)
	}

	return !strings.ContainsAny(
		contentLine,
		charsFind,
	)
}
