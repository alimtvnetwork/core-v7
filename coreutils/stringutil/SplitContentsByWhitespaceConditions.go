package stringutil

import (
	"sort"
	"strings"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

func SplitContentsByWhitespaceConditions(
	input string,
	isTrimEachLine,
	isNonEmptyWhitespace,
	isSort bool,
	isUnique bool,
	isLowerCase bool,
) []string {
	if isLowerCase || isUnique {
		input = strings.ToLower(input)
	}

	compiledStringSplits := strings.Fields(input)

	if isNonEmptyWhitespace && isTrimEachLine {
		compiledStringSplits = stringslice.NonWhitespaceTrimSlice(
			compiledStringSplits)
	} else if isNonEmptyWhitespace && !isTrimEachLine {
		compiledStringSplits = stringslice.NonWhitespaceSlice(
			compiledStringSplits)
	}

	if isUnique {
		hashset := corestr.New.Hashset.StringsPtr(&compiledStringSplits)
		compiledStringSplits = hashset.List()
	}

	if isSort {
		sort.Strings(compiledStringSplits)
	}

	return compiledStringSplits
}
