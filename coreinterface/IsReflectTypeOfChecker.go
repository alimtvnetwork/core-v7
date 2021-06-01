package coreinterface

import "reflect"

type IsReflectTypeOfChecker interface {
	IsReflectTypeOf(typeRequest reflect.Type) bool
}
