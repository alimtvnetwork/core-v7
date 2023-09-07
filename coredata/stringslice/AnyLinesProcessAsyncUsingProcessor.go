package stringslice

import (
	"reflect"
	"sync"

	"gitlab.com/auk-go/core/constants"
)

func AnyLinesProcessAsyncUsingProcessor(
	lines interface{},
	lineProcessor func(index int, lineIn interface{}) (lineOut string),
) []string {
	if lines == nil {
		return []string{}
	}

	reflectType := reflect.TypeOf(lines)
	kind := reflectType.Kind()
	isArrayOrSlice := kind == reflect.Slice ||
		kind == reflect.Array

	if !isArrayOrSlice {
		return []string{}
	}

	reflectValue := reflect.ValueOf(lines)
	length := reflectValue.Len()

	if length == 0 {
		return []string{}
	}

	slice := Make(constants.Zero, length)
	wg := sync.WaitGroup{}

	wg.Add(length)

	asyncProcessor := func(index int, lineIn interface{}) {
		slice[index] = lineProcessor(index, lineIn)

		wg.Done()
	}

	for i := 0; i < length; i++ {
		itemAt := reflectValue.Index(i)
		go asyncProcessor(i, itemAt.Interface())
	}

	wg.Wait()

	return slice
}
