package coreinterface

type BasicReflection interface {
	IsPointerChecker
	IsReflectTypeOfChecker
	IsReflectKindChecker
	ReflectValueGetter
	ReflectTypeGetter
}
