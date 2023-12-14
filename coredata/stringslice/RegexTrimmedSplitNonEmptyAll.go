package stringslice

import (
	"regexp"

	"gitlab.com/auk-go/core/constants"
)

func RegexTrimmedSplitNonEmptyAll(
	regexp *regexp.Regexp,
	content string,
) []string {
	items := regexp.Split(
		content,
		constants.TakeAllMinusOne)

	if len(items) == 0 {
		return []string{}
	}

	return TrimmedEachWords(items)
}
