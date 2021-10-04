package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/msgtype"
)

func ReflectTypeValidation(
	isNotNullExpected bool,
	expectedType reflect.Type,
	anyItem interface{},
) error {
	if isNotNullExpected && anyItem == nil {
		return msgtype.ExpectingErrorSimpleNoType(
			"ReflectTypeValidation: cannot be nil but got nil.",
			"not nil",
			"<nil>")
	}

	actualType := reflect.TypeOf(anyItem)

	if actualType == expectedType {
		return nil
	}

	return msgtype.ExpectingErrorSimpleNoType(
		"ReflectTypeValidation: reflect type validation failed",
		expectedType,
		actualType)
}
