package msgcreator

import (
	"sort"
	"strings"

	"gitlab.com/auk-go/core/coredata/stringslice"
)

func SplitByEachWordTrimmedNoSpace(
	message string,
	isSort bool,
) []string {
	trimmed := strings.TrimSpace(message)

	if trimmed == "" {
		return []string{}
	}

	items := strings.Fields(message)
	items = stringslice.TrimmedEachWords(items)

	if isSort {
		sort.Strings(items)
	}

	return items
}
