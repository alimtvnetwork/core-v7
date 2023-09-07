package coreinstruction

import "gitlab.com/auk-go/core/constants"

type BaseRequestIds struct {
	RequestIds []IdentifierWithIsGlobal `json:"RequestIds,omitempty"`
}

func NewBaseRequestIds(
	isGlobal bool,
	ids ...string,
) *BaseRequestIds {
	return &BaseRequestIds{
		RequestIds: *NewRequestIds(
			isGlobal,
			ids...),
	}
}

func NewRequestIds(
	isGlobal bool,
	ids ...string,
) *[]IdentifierWithIsGlobal {
	slice := make([]IdentifierWithIsGlobal, len(ids))
	if len(ids) == 0 {
		return &slice
	}

	for i, id := range ids {
		slice[i] = IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{Id: id},
			IsGlobal:       isGlobal,
		}
	}

	return &slice
}

func NewRequestId(
	isGlobal bool,
	id string,
) *IdentifierWithIsGlobal {
	return &IdentifierWithIsGlobal{
		BaseIdentifier: BaseIdentifier{Id: id},
		IsGlobal:       isGlobal,
	}
}

func (b *BaseRequestIds) RequestIdsLength() int {
	if b != nil && b.RequestIds != nil {
		return constants.Zero
	}

	return len(b.RequestIds)
}

func (b *BaseRequestIds) AddReqId(
	requestId IdentifierWithIsGlobal,
) *BaseRequestIds {
	b.RequestIds = append(
		b.RequestIds,
		requestId)

	return b
}

func (b *BaseRequestIds) AddIds(
	isGlobal bool,
	ids ...string,
) *BaseRequestIds {
	if len(ids) == 0 {
		return b
	}

	for _, id := range ids {
		b.RequestIds = append(b.RequestIds, IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{id},
			IsGlobal:       isGlobal,
		})
	}

	return b
}

func (b *BaseRequestIds) IsEmptyRequestIds() bool {
	return b.RequestIdsLength() == 0
}

func (b *BaseRequestIds) HasRequestIds() bool {
	return b != nil && b.RequestIds != nil && len(b.RequestIds) > 0
}

func (b *BaseRequestIds) Clone() *BaseRequestIds {
	if b == nil {
		return nil
	}

	length := b.RequestIdsLength()
	slice := make(
		[]IdentifierWithIsGlobal,
		length)

	if length == 0 {
		return &BaseRequestIds{
			RequestIds: slice,
		}
	}

	for i, reqId := range b.RequestIds {
		slice[i] = IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{Id: reqId.Id},
			IsGlobal:       reqId.IsGlobal,
		}
	}

	return &BaseRequestIds{
		RequestIds: slice,
	}
}
