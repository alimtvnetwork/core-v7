package isany

import "reflect"

func PositiveIntegerTypeRv(rv reflect.Value) bool {
	switch rv.Kind() {
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return true
	default:
		return false
	}
}

func PositiveIntegerType(anyItem interface{}) bool {
	return PositiveIntegerTypeRv(reflect.ValueOf(anyItem))
}
