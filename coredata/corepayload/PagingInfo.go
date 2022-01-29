package corepayload

type PagingInfo struct {
	CurrentPageIndex                     int // -- 1 based index
	TotalPages, PerPageItems, TotalItems int
}

func (it *PagingInfo) IsEmpty() bool {
	return it == nil ||
		it.TotalPages == 0 && it.TotalItems == 0
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

func (it *PagingInfo) HasTotalPages() bool {
	return it != nil && it.TotalPages > 0
}

func (it *PagingInfo) HasCurrentPageIndex() bool {
	return it != nil && it.CurrentPageIndex > 0
}

func (it *PagingInfo) HasPerPageItems() bool {
	return it != nil && it.PerPageItems > 0
}

func (it *PagingInfo) HasTotalItems() bool {
	return it != nil && it.TotalItems > 0
}

func (it *PagingInfo) IsInvalidTotalPages() bool {
	return it == nil || it.TotalPages <= 0
}

func (it *PagingInfo) IsInvalidCurrentPageIndex() bool {
	return it == nil || it.CurrentPageIndex <= 0
}

func (it *PagingInfo) IsInvalidPerPageItems() bool {
	return it == nil || it.PerPageItems <= 0
}

func (it *PagingInfo) IsInvalidTotalItems() bool {
	return it == nil || it.TotalItems <= 0
}

func (it PagingInfo) Clone() PagingInfo {
	return PagingInfo{
		TotalPages:       it.TotalPages,
		CurrentPageIndex: it.CurrentPageIndex,
		PerPageItems:     it.PerPageItems,
		TotalItems:       it.TotalItems,
	}
}

func (it *PagingInfo) ClonePtr() *PagingInfo {
	if it == nil {
		return nil
	}

	return &PagingInfo{
		TotalPages:       it.TotalPages,
		CurrentPageIndex: it.CurrentPageIndex,
		PerPageItems:     it.PerPageItems,
		TotalItems:       it.TotalItems,
	}
}
