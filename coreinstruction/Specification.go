package coreinstruction

type Specification struct {
	Id, Display, Type string
	Tags              *[]string `json:"Tags,omitempty"`
	IsGlobal          bool      `json:"IsGlobal"`
}
