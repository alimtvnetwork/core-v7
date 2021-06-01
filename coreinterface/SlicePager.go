package coreinterface

// SlicePager Paging count related methods
type SlicePager interface {
	GetPagesCount(perPageItems int) int
	// HasPageAt returns true if page count is less than the page count
	HasPageAt(perPageItems int, pageIndex int) bool
}
