package stringcompareas

import "gitlab.com/evatix-go/core/coreutils/stringutil"

// isNotEndsWithFunc tided with NotEndsWith
var isNotEndsWithFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	return !stringutil.IsEndsWith(
		contentLine,
		searchComparingLine,
		isIgnoreCase)
}
