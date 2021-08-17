package reflectinternal

import "reflect"

func ReflectInterfaceVal(any interface{}) interface{} {
	rVal := reflect.ValueOf(any)

	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
	}

	return rVal.Interface()
}
