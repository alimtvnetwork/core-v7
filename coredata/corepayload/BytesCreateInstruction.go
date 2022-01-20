package corepayload

type BytesCreateInstruction struct {
	Name, Identifier string
	TaskTypeName     string
	EntityType       string // for any type no need to entity type it will be collected by reflection.
	CategoryName     string
	HasManyRecords   bool
	Payloads         []byte
	Attributes       *Attributes
}
