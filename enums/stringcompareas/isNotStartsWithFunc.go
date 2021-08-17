package stringcompareas

import "gitlab.com/evatix-go/core/coreutils/stringutil"

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
