package coreapi

type PageRequest struct {
	PageSize  int
	PageIndex int
}

func (it *PageRequest) IsPageSizeEmpty() bool {
	return it == nil || it.PageSize <= 0
}

func (it *PageRequest) IsPageIndexEmpty() bool {
	return it == nil || it.PageIndex <= 0
}

func (it *PageRequest) HasPageSize() bool {
	return it != nil && it.PageSize > 0
}

func (it *PageRequest) HasPageIndex() bool {
	return it != nil && it.PageIndex > 0
}

func (it *PageRequest) Clone() *PageRequest {
	if it == nil {
		return nil
	}

	return &PageRequest{
		PageSize:  it.PageSize,
		PageIndex: it.PageIndex,
	}
}
