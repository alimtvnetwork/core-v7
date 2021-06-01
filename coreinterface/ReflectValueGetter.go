package coreinterface

import "reflect"

type ReflectValueGetter interface {
	ReflectValue() *reflect.Value
}
