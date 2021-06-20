package stringutil

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func SplitLeftRightTrimmed(s, separator string) (left, right string) {
	splits := strings.SplitN(
		s, separator,
		ExpectedLeftRightLength)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectedLeftRightLength {
		return strings.TrimSpace(first), strings.TrimSpace(splits[coreindexes.Second])
	}

	return strings.TrimSpace(first), constants.EmptyString
}
