package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

//goland:noinspection GoVarAndConstTypeMayBeOmitted
var (
	emptyBytesType               reflect.Type = reflect.TypeOf([]byte{})
	emptyBytesPointerType        reflect.Type = reflect.TypeOf(&[]byte{})
	getTypeNameFunc                           = reflectinternal.ReflectType.Name
	getTypeNamesUsingReflectFunc              = reflectinternal.ReflectType.NamesStringUsingReflectType
)
