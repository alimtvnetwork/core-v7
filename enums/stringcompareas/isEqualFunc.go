package stringcompareas

import "strings"

var isEqualFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	if isIgnoreCase {
		return strings.EqualFold(
			searchComparingLine,
			contentLine)
	}

	return contentLine ==
		searchComparingLine
}
