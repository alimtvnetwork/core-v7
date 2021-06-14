package coreapi

import "gitlab.com/evatix-go/core/coredata/coredynamic"

type SimpleGenericRequest struct {
	Attribute *RequestAttribute          `json:"Attribute,omitempty"`
	Request   *coredynamic.SimpleRequest `json:"Request,omitempty"`
}
