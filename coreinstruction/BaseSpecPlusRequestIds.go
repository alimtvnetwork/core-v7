package coreinstruction

type BaseSpecPlusRequestIds struct {
	BaseSpecification
	BaseRequestIds
}

func NewBaseSpecPlusRequestIdsUsingSpecOnly(
	spec *Specification,
) *BaseSpecPlusRequestIds {
	return &BaseSpecPlusRequestIds{
		BaseSpecification: BaseSpecification{
			Specification: spec,
		},
		BaseRequestIds: BaseRequestIds{
			RequestIds: []IdentifierWithIsGlobal{},
		},
	}
}

func NewBaseSpecPlusRequestIds(
	spec *Specification,
	reqIds []IdentifierWithIsGlobal,
) *BaseSpecPlusRequestIds {
	return &BaseSpecPlusRequestIds{
		BaseSpecification: BaseSpecification{
			Specification: spec,
		},
		BaseRequestIds: BaseRequestIds{
			RequestIds: reqIds,
		},
	}
}

func (b *BaseSpecPlusRequestIds) Clone() *BaseSpecPlusRequestIds {
	if b == nil {
		return nil
	}

	return &BaseSpecPlusRequestIds{
		BaseSpecification: *b.BaseSpecification.Clone(),
		BaseRequestIds:    *b.BaseRequestIds.Clone(),
	}
}
