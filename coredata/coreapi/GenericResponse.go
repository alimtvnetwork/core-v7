package coreapi

import "gitlab.com/evatix-go/core/coredata/coredynamic"

type GenericResponse struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  interface{}        `json:"Response,omitempty"`
}

func (receiver GenericResponse) GenericResponseResult() *GenericResponseResult {
	return &GenericResponseResult{
		Attribute: receiver.Attribute,
		Response: coredynamic.NewSimpleResult(
			receiver,
			receiver.Attribute.IsValid,
			receiver.Attribute.InvalidMessage),
	}
}
