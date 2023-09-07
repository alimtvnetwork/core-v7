package stringcompareas

import "gitlab.com/auk-go/core/coreutils/stringutil"

var isNotStartsWithFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	return !stringutil.IsStartsWith(
		contentLine,
		searchComparingLine,
		isIgnoreCase)
}
