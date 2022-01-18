package corepayload

type PayloadCreate struct {
	Name, Identifier string
	TaskTypeName     string
	EntityType       string // for any type no need to entity type it will be collected by reflection.
	CategoryName     string
	HasManyRecords   bool
	Payloads         interface{} // for any type no need to entity type it will be collected by reflection.
	Attributes       *Attributes
}
