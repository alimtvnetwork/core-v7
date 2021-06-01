package reflectinternal

import "reflect"

// IsPrimitive function returns true if the kind passed to it is one of the
// primitive types (boolean, int, uint, float, string)
func IsPrimitive(kind reflect.Kind) bool {
	switch kind {
	case
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}
