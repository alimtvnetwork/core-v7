package stringslice

import "strings"

// SplitTrimmedNonEmpty
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
func SplitTrimmedNonEmpty(
	content,
	splitter string,
	number int,
) []string {
	items := strings.SplitN(
		content,
		splitter,
		number)

	if len(items) == 0 {
		return []string{}
	}

	return TrimmedEachWords(items)
}
