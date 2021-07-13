package coretests

import (
	"sort"
	"strings"

	"gitlab.com/evatix-go/core/coredata/stringslice"
)

func GetTrimmedNonEmptySpaceSplit(
	message string,
	isSort bool,
) []string {
	trimmed := strings.TrimSpace(message)

	if trimmed == "" {
		return []string{}
	}

	items := strings.Fields(message)
	items = stringslice.NonWhitespaceTrimSlice(items)

	if isSort {
		sort.Strings(items)
	}

	return items
}
