package coreinterface

// DynamicPageItemsGetter Paging items related methods
type DynamicPageItemsGetter interface {
	// GetDynamicPagedItems returns items for the paged collection.
	GetDynamicPagedItems(perPageItems int, pageIndex int) interface{}
}
