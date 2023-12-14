package stringslice

import (
	"strings"
)

func SplitTrimmedNonEmptyAll(
	content,
	splitter string,
) []string {
	items := strings.Split(
		content,
		splitter)

	if len(items) == 0 {
		return []string{}
	}

	return TrimmedEachWords(items)
}
