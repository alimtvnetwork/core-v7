package stringcompareas

import "strings"

// isAnyCharsFunc reports whether any Unicode
// code points in chars are within contentLine.
//
// Tided with AnyChars
var isAnyCharsFunc = func(
	contentLine,
	charsFind string,
	isIgnoreCase bool,
) bool {
	if isIgnoreCase {
		return strings.ContainsAny(
			strings.ToLower(contentLine),
			strings.ToLower(charsFind),
		)
	}

	return strings.ContainsAny(
		contentLine,
		charsFind,
	)
}
