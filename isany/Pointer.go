package isany

import "reflect"

func Pointer(anyItem interface{}) (isPtr bool) {
	rv := reflect.ValueOf(anyItem)

	return rv.Kind() == reflect.Ptr
}
