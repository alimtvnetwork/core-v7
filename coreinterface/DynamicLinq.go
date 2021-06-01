package coreinterface

type DynamicLinq interface {
	CountGetter
	LengthGetter
	FirstDynamic() interface{}
	LastDynamic() interface{}
	FirstOrDefaultDynamic() interface{}
	LastOrDefaultDynamic() interface{}
	SkipDynamic(skippingItemsCount int) interface{}
	TakeDynamic(takeDynamicItems int) interface{}
	// LimitDynamic alias for TakeDynamic
	LimitDynamic(limit int) interface{}
}
