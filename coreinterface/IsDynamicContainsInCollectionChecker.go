package coreinterface

type IsDynamicContainsInCollectionChecker interface {
	IsDynamicContainsInCollection(collection, item interface{}) bool
}
