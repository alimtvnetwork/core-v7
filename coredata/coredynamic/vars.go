package coredynamic

import "reflect"

//goland:noinspection GoVarAndConstTypeMayBeOmitted
var (
	emptyBytesType                 reflect.Type = reflect.TypeOf([]byte{})
	emptyBytesPointerType          reflect.Type = reflect.TypeOf(&[]byte{})
	ReflectGetter                               = reflectGetter{}
	ReflectGetterUsingReflectValue              = reflectGetUsingReflectValue{}
)
