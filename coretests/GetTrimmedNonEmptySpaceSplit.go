package coretests

import (
	"sort"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/testconsts"
)

func GetTrimmedNonEmptySpaceSplit(
	message string,
	isSort bool,
) []string {
	trimmed := strings.TrimSpace(message)

	if trimmed == "" {
		return []string{}
	}

	items := testconsts.WhitespaceOrPipeFinderRegex.Split(
		message,
		constants.TakeAllMinusOne)

	items = stringslice.NonWhitespaceTrimSlice(items)

	if isSort {
		sort.Strings(items)
	}

	return items
}
