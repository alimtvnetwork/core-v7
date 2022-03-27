package reflectinternal

import "reflect"

// IsZero
//
//  returns true if the current value is null
//  or reflect value is zero
//
// Reference:
//  - Stackoverflow Example : https://stackoverflow.com/a/23555352
func IsZero(anyItem interface{}) bool {
	if IsNull(anyItem) {
		return true
	}

	return IsZeroReflectValue(reflect.ValueOf(anyItem))
}
