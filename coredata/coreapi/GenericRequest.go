package coreapi

import "gitlab.com/evatix-go/core/coredata/coredynamic"

type GenericRequestIn struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   interface{}       `json:"Request,omitempty"`
}

func (receiver *GenericRequestIn) SimpleGenericRequest(
	isValid bool,
	invalidMessage string,
) *SimpleGenericRequest {
	return &SimpleGenericRequest{
		Attribute: receiver.Attribute,
		Request:   coredynamic.NewSimpleRequest(receiver.Request, isValid, invalidMessage),
	}
}
