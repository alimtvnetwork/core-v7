package reflectinternal

import "reflect"

func ReflectValToInterfacesUsingProcessor(
	isSkipOnNil bool,
	processorFunc func(item interface{}) (result interface{}, isTake, isBreak bool),
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
		valueInf := value.Interface()

		if isSkipOnNil && valueInf == nil {
			continue
		}

		rs, isTake, isBreak :=
			processorFunc(valueInf)

		if isTake {
			slice = append(slice, rs)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}
