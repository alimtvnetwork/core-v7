package pagingutil

// PagingRequest
//
// - PageIndex : Starts from one.
type PagingRequest struct {
	// PageIndex usually starts from one.
	Length, PageIndex, EachPageSize int
}
