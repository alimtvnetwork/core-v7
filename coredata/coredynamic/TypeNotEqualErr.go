package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func TypeNotEqualErr(
	left, right interface{},
) error {
	leftRt := reflect.TypeOf(left)
	rightRt := reflect.TypeOf(right)

	if leftRt == rightRt {
		return nil
	}

	return errcore.
		TypeMismatchType.
		SrcDestinationErr(
			"left, right type doesn't match!",
			constants.LeftLower, leftRt.String(),
			constants.Right, rightRt.String(),
		)
}
