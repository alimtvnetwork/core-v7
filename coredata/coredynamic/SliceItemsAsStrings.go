package coredynamic

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func SliceItemsAsStrings(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return SliceItemsAsStrings(
			reflect.Indirect(reflect.ValueOf(reflectVal)))
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			errcore.TypeMismatchType.Error("Reflection is not Slice or Array", reflectVal)
	}

	length := reflectVal.Len()
	slice := make([]string, length)

	if length == 0 {
		return slice, nil
	}

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)
		toString := fmt.Sprintf(
			constants.SprintValueFormat,
			value.Interface())
		slice[i] = toString
	}

	return slice, nil
}
