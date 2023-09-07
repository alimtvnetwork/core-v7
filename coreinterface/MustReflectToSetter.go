package coreinterface

type MustReflectToSetter interface {
	ReflectSetToMust(toPointer interface{})
}
