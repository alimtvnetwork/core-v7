package coreapi

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
)

type GenericResponse struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  interface{}        `json:"Response,omitempty"`
}

func InvalidGenericResponse(attr *ResponseAttribute) *GenericResponse {
	if attr == nil {
		return &GenericResponse{
			Attribute: InvalidResponseAttribute(constants.EmptyString),
			Response:  nil,
		}
	}

	return &GenericResponse{
		Attribute: attr,
		Response:  nil,
	}
}

func (it *GenericResponse) GenericResponseResult() *GenericResponseResult {
	return &GenericResponseResult{
		Attribute: it.Attribute,
		Response: coredynamic.NewSimpleResult(
			it,
			it.Attribute.IsValid,
			it.Attribute.Message),
	}
}

// Clone Cannot copy interface, just putting response in response field.
func (it *GenericResponse) Clone() *GenericResponse {
	if it == nil {
		return nil
	}

	return &GenericResponse{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response,
	}
}
