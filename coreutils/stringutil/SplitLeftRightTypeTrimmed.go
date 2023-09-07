package stringutil

import (
	"gitlab.com/auk-go/core/coredata/corestr"
)

func SplitLeftRightTypeTrimmed(s, separator string) *corestr.LeftRight {
	left, right := SplitLeftRightTrimmed(s, separator)

	return &corestr.LeftRight{
		Left:    left,
		Right:   right,
		IsValid: right != "" && left != "",
		Message: "",
	}
}
