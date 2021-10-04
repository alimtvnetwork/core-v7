package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/msgtype"
)

func ReflectKindValidation(
	expectedKind reflect.Kind,
	anyItem interface{},
) error {
	actualKind := reflect.ValueOf(anyItem).Kind()

	if actualKind == expectedKind {
		return nil
	}

	return msgtype.ExpectingErrorSimpleNoType(
		"ReflectKindValidation: reflect kind validation failed",
		expectedKind,
		actualKind)
}
