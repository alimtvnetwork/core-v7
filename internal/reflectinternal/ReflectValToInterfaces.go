package reflectinternal

import (
	"reflect"
)

func ReflectValToInterfaces(
	isSkipOnNil bool,
	reflectVal reflect.Value,
) []interface{} {
	if reflectVal.Kind() == reflect.Ptr {
		return ReflectValToInterfaces(
			isSkipOnNil,
			reflect.Indirect(reflect.ValueOf(reflectVal)))
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []interface{}{}
	}

	length := reflectVal.Len()
	slice := make([]interface{}, 0, length)

	if length == 0 {
		return slice
	}

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		valueInf := value.Interface()

		if isSkipOnNil && IsNull(value) {
			continue
		}

		slice = append(slice, valueInf)
	}

	return slice
}
