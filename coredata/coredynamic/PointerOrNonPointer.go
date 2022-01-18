package coredynamic

import "reflect"

func PointerOrNonPointer(
	isPointerOutput bool,
	input interface{},
) (output interface{}, finalReflectVal reflect.Value) {
	return PointerOrNonPointerUsingReflectValue(
		isPointerOutput,
		reflect.ValueOf(input))
}
