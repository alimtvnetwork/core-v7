package coreinterface

import "reflect"

type IsReflectionTypeChecker interface {
	IsManyReflectionOfType(typeOf reflect.Type, dynamicItems ...interface{}) bool
	IsReflectionOfType(dynamic interface{}, typeOf reflect.Type) bool
	IsReflectionOfTypeName(dynamic interface{}, typeOfName string) bool
}
