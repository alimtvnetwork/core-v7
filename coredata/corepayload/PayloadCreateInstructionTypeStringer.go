package corepayload

import "fmt"

type PayloadCreateInstructionTypeStringer struct {
	Name, Identifier     string
	TaskTypeNameStringer fmt.Stringer
	CategoryNameStringer fmt.Stringer
	HasManyRecords       bool
	Payloads             interface{} // for any type no need to entity type it will be collected by reflection.
	Attributes           *Attributes
}

func (it PayloadCreateInstructionTypeStringer) PayloadCreateInstruction() *PayloadCreateInstruction {
	return &PayloadCreateInstruction{
		Name:           it.Name,
		Identifier:     it.Identifier,
		TaskTypeName:   it.TaskTypeNameStringer.String(),
		CategoryName:   it.CategoryNameStringer.String(),
		HasManyRecords: it.HasManyRecords,
		Payloads:       it.Payloads,
		Attributes:     it.Attributes,
	}
}
