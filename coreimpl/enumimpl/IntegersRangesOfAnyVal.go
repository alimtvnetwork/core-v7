package enumimpl

import (
	"reflect"
	"sort"
)

func IntegersRangesOfAnyVal(anyValue interface{}) []int {
	reflectValues := reflect.ValueOf(anyValue)
	length := reflectValues.Len()
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		rfVal := reflectValues.Index(i)
		anyVal := rfVal.Interface()

		slice[i] = ConvEnumAnyValToInteger(anyVal)
	}

	sort.Ints(slice)

	return slice
}
