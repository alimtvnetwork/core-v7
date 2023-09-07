package coreapi

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
)

type GenericResponseResult struct {
	Attribute *ResponseAttribute        `json:"Attribute,omitempty"`
	Response  *coredynamic.SimpleResult `json:"Response,omitempty"`
}

func (it GenericResponseResult) Clone() GenericResponseResult {
	return GenericResponseResult{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response.ClonePtr(),
	}
}

func (it *GenericResponseResult) ClonePtr() *GenericResponseResult {
	if it == nil {
		return nil
	}

	return &GenericResponseResult{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response.ClonePtr(),
	}
}
