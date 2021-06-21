package coreinstruction

type BaseSpecification struct {
	Specification *Specification `json:"Specification,omitempty"`
}

func NewBaseSpecification(
	id,
	display,
	typeName string,
	tags []string,
	isGlobal bool,
) *BaseSpecification {
	spec := NewSpecification(
		id, display,
		typeName,
		tags,
		isGlobal)

	return &BaseSpecification{
		Specification: spec,
	}
}

func (b *BaseSpecification) Identifier() BaseIdentifier {
	return b.Specification.BaseIdentifier
}

func (b *BaseSpecification) Display() BaseDisplay {
	return b.Specification.BaseDisplay
}

func (b *BaseSpecification) Type() BaseType {
	return b.Specification.BaseType
}

func (b *BaseSpecification) IsEmptySpec() bool {
	return !b.HasSpec()
}

func (b *BaseSpecification) HasSpec() bool {
	return b != nil && b.Specification != nil
}

func (b *BaseSpecification) Clone() *BaseSpecification {
	if b == nil {
		return nil
	}

	return &BaseSpecification{
		Specification: b.Specification.Clone(),
	}
}
