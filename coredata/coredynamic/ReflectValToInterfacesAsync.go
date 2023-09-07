package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func ReflectValToInterfacesAsync(
	reflectVal reflect.Value,
) []interface{} {
	return reflectinternal.ReflectValToInterfacesAsync(
		reflectVal)
}
