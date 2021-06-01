package coreinterface

import "reflect"

type ReflectTypeGetter interface {
	ReflectType() reflect.Type
	GetErrorOnTypeMismatch(
		typeMatch reflect.Type,
		isIncludeInvalidMessage bool,
	) error
}
