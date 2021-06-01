package coreinterface

type DynamicAdder interface {
	AddDynamic(addingItem interface{}) (isSuccess bool)
	AddDynamicMany(addingItem ...interface{}) (isSuccess bool)
}
