package isany

import "reflect"

func FloatingPointType(anyItem interface{}) bool {
	return FloatingPointTypeRv(reflect.ValueOf(anyItem))
}
