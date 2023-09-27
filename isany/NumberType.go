package isany

import "reflect"

func NumberType(anyItem interface{}) bool {
	return NumberTypeRv(reflect.ValueOf(anyItem))
}
