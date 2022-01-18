package coreapi

type PayloadsRequestIn struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   []byte            `json:"Request,omitempty"`
}
