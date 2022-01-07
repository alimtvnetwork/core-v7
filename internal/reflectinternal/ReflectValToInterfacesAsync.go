package reflectinternal

import (
	"reflect"
	"sync"
)

func ReflectValToInterfacesAsync(
	reflectVal reflect.Value,
) []interface{} {
	if reflectVal.Kind() == reflect.Ptr {
		return ReflectValToInterfacesAsync(
			reflect.Indirect(reflect.ValueOf(reflectVal)))
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []interface{}{}
	}

	length := reflectVal.Len()
	slice := make([]interface{}, length)

	if length == 0 {
		return slice
	}

	wg := sync.WaitGroup{}
	setterIndexFunc := func(index int) {
		value := reflectVal.Index(index)

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		valueInf := value.Interface()
		slice[index] = valueInf

		wg.Done()
	}

	wg.Add(length)
	for i := 0; i < length; i++ {
		go setterIndexFunc(i)
	}

	wg.Wait()

	return slice
}
