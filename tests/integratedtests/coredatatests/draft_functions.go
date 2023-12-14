package coredatatests

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func someFunctionV1(arg1, arg2, arg3 string) string {
	return fmt.Sprintf(
		"%s => called with (%s, %s, %s) - some new stuff",
		reflectinternal.GetFunc.Name(someFunctionV1),
		arg1,
		arg2,
		arg3,
	)
}

func someFunctionV2(arg1, arg2 string) (string, error) {
	leftPart := fmt.Sprintf(
		"%s => called with (%s, %s) - (string, error)",
		reflectinternal.GetFunc.Name(someFunctionV2),
		arg1,
		arg2,
	)

	return leftPart, errors.New("some err v2")
}

func someFunctionV3(arg1, arg2 string) (int, string, error) {
	middlePart := fmt.Sprintf(
		"%s => called with (%s, %s) - (int, string, error)",
		reflectinternal.GetFunc.Name(someFunctionV3),
		arg1,
		arg2,
	)

	return 5, middlePart, errors.New("some err of v3")
}
