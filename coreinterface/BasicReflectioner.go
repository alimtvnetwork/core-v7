package coreinterface

type BasicReflectioner interface {
	IsPointerChecker
	IsReflectTypeOfChecker
	IsReflectKindChecker
	ReflectValueGetter
	ReflectTypeGetter
}
