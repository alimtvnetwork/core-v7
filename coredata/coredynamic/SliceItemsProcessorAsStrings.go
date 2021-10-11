package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/errcore"
)

func SliceItemsProcessorAsStrings(
	reflectVal reflect.Value,
	processor func(index int, item interface{}) (result string, isTake, isBreak bool),
) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return SliceItemsAsStrings(
			reflect.Indirect(reflect.ValueOf(reflectVal)))
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			errcore.TypeMismatch.Error("Reflection is not Slice or Array", reflectVal)
	}

	length := reflectVal.Len()
	slice := make([]string, 0, length)

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)
		result, isTake, isBreak := processor(
			i, value)

		if isTake {
			slice = append(
				slice,
				result)
		}

		if isBreak {
			return slice, nil
		}
	}

	return slice, nil
}
