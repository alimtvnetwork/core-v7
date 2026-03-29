package coredynamic

import "reflect"

func PointerOrNonPointerUsingReflectValue(
	isPointerOutput bool,
	rv reflect.Value,
) (output any, finalReflectVal reflect.Value) {
	k := rv.Kind()

	if !isPointerOutput && (k == reflect.Ptr || k == reflect.Interface) {
		finalReflectVal = rv.Elem()

		return finalReflectVal.Interface(), finalReflectVal
	}

	if isPointerOutput && k != reflect.Ptr {
		val := rv.Interface()

		return &val, rv
	}

	// struct or value
	return rv.Interface(), rv
}
