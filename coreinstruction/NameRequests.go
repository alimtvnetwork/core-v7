package coreinstruction

import "gitlab.com/auk-go/core/coredata/corestr"

type NameRequests struct {
	Name     string               `json:"Name,omitempty"`
	Requests *corestr.SimpleSlice `json:"Requests,omitempty"`
}
