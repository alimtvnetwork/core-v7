package coreinterface

type ReflectSetter interface {
	ReflectSet(toPointer interface{}) error
}
