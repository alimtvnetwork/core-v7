package coreinstruction

import "gitlab.com/auk-go/core/coredata/stringslice"

type FlatSpecification struct {
	Id       string   `json:"Id"`
	Display  string   `json:"Display"`
	Type     string   `json:"Type"`
	IsGlobal bool     `json:"IsGlobal"`
	Tags     []string `json:"Tags,omitempty"`
	IsValid  bool     `json:"IsValid,omitempty"`
	spec     *Specification
}

func InvalidFlatSpecification() *FlatSpecification {
	return &FlatSpecification{
		Id:       "",
		Display:  "",
		Type:     "",
		IsGlobal: false,
		Tags:     []string{},
		IsValid:  false,
	}
}

func NewFlatSpecificationUsingSpec(spec *Specification, isValid bool) *FlatSpecification {
	return &FlatSpecification{
		Id:       spec.Id,
		Display:  spec.Display,
		Type:     spec.Type,
		IsGlobal: spec.IsGlobal,
		Tags:     spec.Tags,
		IsValid:  isValid,
		spec:     spec,
	}
}

func (it *FlatSpecification) BaseIdentifier() BaseIdentifier {
	return it.Spec().BaseIdentifier
}

func (it *FlatSpecification) BaseTags() BaseTags {
	return it.Spec().BaseTags
}

func (it *FlatSpecification) BaseIsGlobal() BaseIsGlobal {
	return it.Spec().BaseIsGlobal
}

func (it *FlatSpecification) BaseDisplay() BaseDisplay {
	return it.Spec().BaseDisplay
}

func (it *FlatSpecification) BaseType() BaseType {
	return it.Spec().BaseType
}

func (it *FlatSpecification) Spec() *Specification {
	if it == nil {
		return nil
	}

	if it.spec != nil {
		return it.spec
	}

	it.spec = &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: it.Id},
			BaseDisplay:    BaseDisplay{it.Display},
			BaseType:       BaseType{it.Type},
		},
		BaseTags: BaseTags{
			Tags: it.Tags,
		},
		BaseIsGlobal: BaseIsGlobal{IsGlobal: it.IsGlobal},
	}

	return it.spec
}

func (it *FlatSpecification) Clone() *FlatSpecification {
	if it == nil {
		return nil
	}

	return &FlatSpecification{
		Id:       it.Id,
		Display:  it.Display,
		Type:     it.Type,
		IsGlobal: it.IsGlobal,
		Tags:     stringslice.Clone(it.Tags),
		IsValid:  it.IsValid,
	}
}
