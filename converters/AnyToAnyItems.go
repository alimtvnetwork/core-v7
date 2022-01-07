package converters

import (
	"reflect"

	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

func AnyToAnyItems(
	isSkipOnNil bool,
	anyItem interface{},
) []interface{} {
	if isSkipOnNil && anyItem == nil {
		return []interface{}{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal)
}

func AnyToNonNullItems(
	isSkipOnNil bool,
	anyItem interface{},
) []interface{} {
	if isSkipOnNil && anyItem == nil || reflectinternal.IsNull(anyItem) {
		return []interface{}{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal)
}
