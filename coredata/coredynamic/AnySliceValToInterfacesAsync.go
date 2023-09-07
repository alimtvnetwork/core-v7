package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func AnySliceValToInterfacesAsync(
	slice interface{},
) []interface{} {
	if slice == nil {
		return []interface{}{}
	}

	return reflectinternal.ReflectValToInterfacesAsync(
		reflect.ValueOf(slice))
}
