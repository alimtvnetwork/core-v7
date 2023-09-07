package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/errcore"
)

func NotAcceptedTypesErr(
	input interface{},
	acceptedTypes ...reflect.Type,
) error {
	currentRv := reflect.TypeOf(input)
	isNotMatchingAcceptedType := !IsAnyTypesOf(
		currentRv,
		acceptedTypes...)

	if isNotMatchingAcceptedType {
		return errcore.
			ExpectingSimpleNoTypeError(
				"type doesn't match, accepting types are listed.",
				acceptedTypes,
				currentRv)
	}

	return nil
}
