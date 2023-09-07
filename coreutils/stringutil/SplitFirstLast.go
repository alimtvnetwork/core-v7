package stringutil

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreindexes"
)

func SplitFirstLast(s, separator string) (first, last string) {
	splits := strings.Split(
		s, separator,
	)

	length := len(splits)
	first = splits[coreindexes.First]

	if length > 1 {
		return first, splits[length-1]
	}

	return first, constants.EmptyString
}
