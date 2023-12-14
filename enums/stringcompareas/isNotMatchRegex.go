package stringcompareas

import "gitlab.com/auk-go/core/regexnew"

// NotMatchRegex no use of isCaseSensitive
//
// # Tided with NotMatchRegex, invert of isRegexFunc
//
// isCaseSensitive is kept for consistency and calling ability
var isNotMatchRegex = func(
	contentLine,
	regexStringSearching string,
	isIgnoreCase bool,
) bool {
	return !regexnew.IsMatchLock(
		regexStringSearching,
		contentLine)
}
