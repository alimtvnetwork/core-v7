package stringcompareas

import "gitlab.com/evatix-go/core/coreutils/stringutil"

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
