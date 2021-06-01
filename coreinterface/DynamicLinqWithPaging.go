package coreinterface

type DynamicLinqWithPaging interface {
	DynamicLinq
	SlicePager
	DynamicPageItemsGetter
}
