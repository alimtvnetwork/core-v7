package stringcompareas

import "gitlab.com/evatix-go/core/coreutils/stringutil"

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
