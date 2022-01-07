package coreinterface

type ValueReflectSetter interface {
	ValueReflectSet(setterPtr interface{}) error
}
