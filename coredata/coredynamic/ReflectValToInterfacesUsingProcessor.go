package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func ReflectValToInterfacesUsingProcessor(
	isSkipOnNil bool,
	processorFunc func(item interface{}) (result interface{}, isTake, isBreak bool),
	reflectVal reflect.Value,
) []interface{} {
	return reflectinternal.ReflectValToInterfacesUsingProcessor(
		isSkipOnNil,
		processorFunc,
		reflectVal)
}
