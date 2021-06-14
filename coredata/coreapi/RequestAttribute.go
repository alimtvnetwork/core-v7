package coreapi

import (
	"gitlab.com/evatix-go/core/reqtype"
)

type RequestAttribute struct {
	Url           string `json:"Url,omitempty"`
	Host          string `json:"Host,omitempty"`
	ResourceName  string `json:"ResourceName,omitempty"`
	RequestType   reqtype.Request
	SearchRequest *SearchRequest `json:"SearchRequest,omitempty"`
	PageRequest   *PageRequest   `json:"PageRequest,omitempty"`
}
