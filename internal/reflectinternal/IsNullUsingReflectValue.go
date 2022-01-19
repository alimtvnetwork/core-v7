package reflectinternal

import "reflect"

func IsNullUsingReflectValue(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}
