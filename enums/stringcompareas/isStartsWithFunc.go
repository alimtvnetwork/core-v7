package stringcompareas

import "gitlab.com/auk-go/core/coreutils/stringutil"

var isStartsWithFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	return stringutil.IsStartsWith(
		contentLine,
		searchComparingLine,
		isIgnoreCase)
}
