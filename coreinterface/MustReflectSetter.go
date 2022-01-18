package coreinterface

type MustReflectSetter interface {
	ReflectSetMust(toPointer interface{})
}
