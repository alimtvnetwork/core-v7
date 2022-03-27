package reflectinternal

import "reflect"

// IsZeroReflectValue
//
//  returns true if the current value is null
//  or reflect value is zero
//
// Reference:
//  - Stackoverflow Example : https://stackoverflow.com/a/23555352
func IsZeroReflectValue(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice, reflect.Ptr:
		return rv.IsNil()
	case reflect.Array:
		isAllZero := true
		for i := 0; i < rv.Len(); i++ {
			isAllZero = isAllZero && IsZeroReflectValue(rv.Index(i))
		}

		return isAllZero
	case reflect.Struct:
		isAllZero := true
		for i := 0; i < rv.NumField(); i++ {
			isAllZero = isAllZero && IsZeroReflectValue(rv.Field(i))
		}

		return isAllZero
	}

	// Compare other types directly:
	z := reflect.Zero(rv.Type())

	return rv.Interface() == z.Interface()
}
