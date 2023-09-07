package corepayload

type payloadWrapperModel struct {
	Name           string      `json:"Name,omitempty"`
	Identifier     string      `json:"Identifier,omitempty"`
	TaskTypeName   string      `json:"TaskTypeName,omitempty"`
	EntityType     string      `json:"EntityType,omitempty"`
	CategoryName   string      `json:"CategoryName,omitempty"`
	HasManyRecords bool        `json:"HasManyRecords,omitempty"`
	Payloads       string      `json:"Payloads,omitempty"`
	Attributes     *Attributes `json:"AnyAttributes,omitempty"`
}
