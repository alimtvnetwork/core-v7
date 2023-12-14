package isany

import "reflect"

// Null
//
// # Returns true for any nil given
//
// Reference : https://stackoverflow.com/a/43896204
func Null(item interface{}) bool {
	if item == nil {
		return true
	}

	rv := reflect.ValueOf(item)

	switch rv.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.UnsafePointer,
		reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}
