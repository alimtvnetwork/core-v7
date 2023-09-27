package isany

import "reflect"

func FloatingPointTypeRv(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}
