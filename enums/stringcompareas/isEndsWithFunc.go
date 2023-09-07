package stringcompareas

import "gitlab.com/auk-go/core/coreutils/stringutil"

var isEndsWithFunc = func(
	contentLine,
	equalChecking string,
	isIgnoreCase bool,
) bool {
	return stringutil.IsEndsWith(
		contentLine,
		equalChecking,
		isIgnoreCase)
}
