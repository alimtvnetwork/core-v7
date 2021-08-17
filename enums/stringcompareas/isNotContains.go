package stringcompareas

import "strings"

// isNotContainsFunc tided with NotContains
var isNotContainsFunc = func(
	contentLine,
	searchComparingLine string,
	isIgnoreCase bool,
) bool {
	if isIgnoreCase {
		return !strings.Contains(
			strings.ToLower(contentLine),
			strings.ToLower(searchComparingLine),
		)
	}

	return !strings.Contains(
		contentLine,
		searchComparingLine,
	)
}
