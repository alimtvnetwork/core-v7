package stringutil

import "gitlab.com/auk-go/core/coredata/corestr"

func SplitLeftRightsTrims(separator string, lines ...string) []*corestr.LeftRight {
	length := len(lines)
	slice := make([]*corestr.LeftRight, length)

	if length == 0 {
		return slice
	}

	for i, line := range lines {
		slice[i] = SplitLeftRightTypeTrimmed(
			separator,
			line)
	}

	return slice
}
