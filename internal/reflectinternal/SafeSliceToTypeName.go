package reflectinternal

import "reflect"

// SafeSliceToTypeName
//
// Gets slice element type name, reduce ptr slice as well.
func SafeSliceToTypeName(slice interface{}) string {
	rt := reflect.TypeOf(slice)

	if IsNull(rt) {
		return ""
	}

	if rt.Kind() == reflect.Ptr || rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt.Elem().String()
}
