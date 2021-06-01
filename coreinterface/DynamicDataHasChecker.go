package coreinterface

type DynamicDataHasChecker interface {
	HasDynamic(searchItem interface{}) bool
	HasDynamicAll(searchTerms ...interface{}) bool
	HasDynamicAny(searchTerms ...interface{}) bool
}
