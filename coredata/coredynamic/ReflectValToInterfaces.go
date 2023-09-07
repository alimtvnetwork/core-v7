package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func ReflectValToInterfaces(
	isSkipOnNil bool,
	reflectVal reflect.Value,
) []interface{} {
	return reflectinternal.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal)
}
