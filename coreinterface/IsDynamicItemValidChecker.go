package coreinterface

type IsDynamicItemValidChecker interface {
	IsDynamicItemValid(item interface{}) bool
	IsDynamicItemsValid(items ...interface{}) bool
}
