package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

type SimpleGenericRequest struct {
	Attribute *RequestAttribute          `json:"Attribute,omitempty"`
	Request   *coredynamic.SimpleRequest `json:"Request,omitempty"`
}
