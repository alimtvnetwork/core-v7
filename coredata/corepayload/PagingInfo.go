package corepayload

type PagingInfo struct {
	TotalPages, CurrentPageIndex, PerPageItems, TotalItems int
}

func (it *PagingInfo) IsEqual(right *PagingInfo) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it.TotalPages != right.TotalPages {
		return false
	}

	if it.CurrentPageIndex != right.CurrentPageIndex {
		return false
	}

	if it.PerPageItems != right.PerPageItems {
		return false
	}

	return it.TotalItems == right.TotalItems
}
