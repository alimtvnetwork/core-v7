package coredynamic

import "reflect"

func Type(any interface{}) reflect.Type {
	return reflect.TypeOf(any)
}
