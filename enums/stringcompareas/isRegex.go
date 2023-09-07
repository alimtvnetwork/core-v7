package stringcompareas

import "gitlab.com/auk-go/core/regexnew"

// isRegexFunc no use of isCaseSensitive
//
// isCaseSensitive is kept for consistency and calling ability
var isRegexFunc = func(
	contentLine,
	regexStringSearching string,
	isIgnoreCase bool,
) bool {
	return regexnew.IsMatchLock(
		regexStringSearching,
		contentLine)
}
