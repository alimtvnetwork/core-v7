package coreinstruction

type IdentifierWithIsGlobal struct {
	BaseIdentifier
	IsGlobal bool `json:"IsGlobal"`
}

func NewIdentifierWithIsGlobal(
	id string,
	isGlobal bool,
) *IdentifierWithIsGlobal {
	return &IdentifierWithIsGlobal{
		BaseIdentifier: BaseIdentifier{
			Id: id,
		},
		IsGlobal: isGlobal,
	}
}

func (receiver *IdentifierWithIsGlobal) Clone() *IdentifierWithIsGlobal {
	return &IdentifierWithIsGlobal{
		BaseIdentifier: *receiver.BaseIdentifier.Clone(),
		IsGlobal:       receiver.IsGlobal,
	}
}
