package reflectinternal

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
)

type sliceConverter struct{}

func (it sliceConverter) Length(i interface{}) int {
	if Is.Null(i) {
		return 0
	}

	reflectVal := reflect.ValueOf(i)

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array ||
		k == reflect.Map

	if isSliceOrArray {
		return reflectVal.Len()
	}

	return 0
}

func (it sliceConverter) ToStringsRv(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRv(
			reflect.Indirect(reflectVal),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			fmt.Errorf("reflection is not a slice nor array but %s", reflectVal.String())
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
			value.Interface(),
		)
		slice[i] = toString
	}

	return slice, nil
}

func (it sliceConverter) ToStrings(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	return it.ToStringsRv(reflectVal)
}

func (it sliceConverter) ToStringsMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	items, err := it.ToStringsRv(reflectVal)

	if err != nil {
		panic(err)
	}

	return items
}

func (it sliceConverter) ToStringsRvUsingProcessor(
	reflectVal reflect.Value,
	processor func(index int, item interface{}) (result string, isTake, isBreak bool),
) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRvUsingProcessor(
			reflect.Indirect(reflectVal),
			processor,
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			fmt.Errorf("reflection is not a slice nor array but %s", reflectVal.String())
	}

	length := reflectVal.Len()
	slice := make([]string, 0, length)

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)
		result, isTake, isBreak := processor(
			i, value,
		)

		if isTake {
			slice = append(
				slice,
				result,
			)
		}

		if isBreak {
			return slice, nil
		}
	}

	return slice, nil
}

func (it sliceConverter) ToStringsRvUsingSimpleProcessor(
	reflectVal reflect.Value,
	isSkipEmpty bool,
	processor func(index int, item interface{}) (result string),
) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRvUsingSimpleProcessor(
			reflect.Indirect(reflectVal),
			isSkipEmpty,
			processor,
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []string{},
			fmt.Errorf("reflection is not a slice nor array but %s", reflectVal.String())
	}

	length := reflectVal.Len()
	slice := make([]string, 0, length)

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)
		result := processor(
			i, value,
		)

		if isSkipEmpty && result == "" {
			continue
		}

		slice = append(
			slice,
			result,
		)
	}

	return slice, nil
}

func (it sliceConverter) ToAnyItemsAsync(
	slice interface{},
) []interface{} {
	if slice == nil {
		return []interface{}{}
	}

	return Converter.ReflectValToInterfacesAsync(
		reflect.ValueOf(slice),
	)
}
