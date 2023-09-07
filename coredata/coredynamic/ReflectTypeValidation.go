package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/errcore"
)

func ReflectTypeValidation(
	isNotNullExpected bool,
	expectedType reflect.Type,
	anyItem interface{},
) error {
	if isNotNullExpected && anyItem == nil {
		return errcore.ExpectingErrorSimpleNoType(
			"ReflectTypeValidation: cannot be nil but got nil.",
			"not nil",
			"<nil>")
	}

	actualType := reflect.TypeOf(anyItem)

	if actualType == expectedType {
		return nil
	}

	return errcore.ExpectingErrorSimpleNoType(
		"ReflectTypeValidation: reflect type validation failed",
		expectedType,
		actualType)
}
