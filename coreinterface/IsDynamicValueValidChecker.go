package coreinterface

type IsDynamicValueValidChecker interface {
	IsDynamicValueValid(value interface{}) bool
}
