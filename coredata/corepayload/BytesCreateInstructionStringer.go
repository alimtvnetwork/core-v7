package corepayload

import "fmt"

type BytesCreateInstructionStringer struct {
	Name, Identifier string
	TaskTypeName     fmt.Stringer
	EntityType       string // for any type no need to entity type it will be collected by reflection.
	CategoryName     fmt.Stringer
	HasManyRecords   bool
	Payloads         []byte
	Attributes       *Attributes
}
