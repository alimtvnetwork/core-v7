package coreapi

import "gitlab.com/evatix-go/core/coredata/coredynamic"

type GenericResponseResult struct {
	Attribute *ResponseAttribute        `json:"Attribute,omitempty"`
	Response  *coredynamic.SimpleResult `json:"Response,omitempty"`
}
