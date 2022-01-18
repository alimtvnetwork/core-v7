package coredynamic

import "reflect"

// ReflectInterfaceVal
//
// Reduce pointer to value (one step only)
func ReflectInterfaceVal(any interface{}) interface{} {
	rVal := reflect.ValueOf(any)
	k := rVal.Kind()

	if k != reflect.Ptr && k != reflect.Interface {
		return rVal.Interface()
	}

	if k == reflect.Ptr || k == reflect.Interface {
		return rVal.Elem().Interface()
	}

	return rVal.Interface()
}
