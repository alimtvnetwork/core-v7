package corestr

import "reflect"

func reflectInterfaceVal(any interface{}) interface{} {
	rVal := reflect.ValueOf(any)

	if rVal.Kind() != reflect.Ptr {
		return rVal.Interface()
	}

	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
	}

	return rVal.Interface()
}
