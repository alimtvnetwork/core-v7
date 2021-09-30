package converters

import (
	"reflect"

	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

func AnyToStrings(
	isSkipOnNil bool,
	anyItem interface{},
) []string {
	if isSkipOnNil && anyItem == nil {
		return []string{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	anyItems := reflectinternal.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal)

	return AnyItemsToStringsSkipOnNil(anyItems)
}
