package coreapi

import (
	"gitlab.com/auk-go/core/reqtype"
)

type RequestAttribute struct {
	Url           string `json:"Url,omitempty"`
	Host          string `json:"Host,omitempty"`
	ResourceName  string `json:"ResourceName,omitempty"`
	ActionName    string `json:"ActionName,omitempty"`
	Identifier    string `json:"Identifier,omitempty"`
	OptionalAuth  string `json:"OptionalAuth,omitempty"`
	ErrorJson     string `json:"ErrorJson,omitempty"`
	RequestType   reqtype.Request
	IsValid       bool
	HasError      bool
	SearchRequest *SearchRequest `json:"SearchRequest,omitempty"`
	PageRequest   *PageRequest   `json:"PageRequest,omitempty"`
}

func (it *RequestAttribute) HasSearchRequest() bool {
	return it != nil && it.SearchRequest != nil
}

func (it *RequestAttribute) HasPageRequest() bool {
	return it != nil && it.PageRequest != nil
}

func (it *RequestAttribute) IsEmpty() bool {
	return it == nil
}

func (it *RequestAttribute) IsAnyNull() bool {
	return it == nil
}

func (it *RequestAttribute) IsPageRequestEmpty() bool {
	return it == nil || it.PageRequest == nil
}

func (it *RequestAttribute) IsSearchRequestEmpty() bool {
	return it == nil || it.SearchRequest == nil
}

func (it *RequestAttribute) Clone() *RequestAttribute {
	if it == nil {
		return nil
	}

	return &RequestAttribute{
		Url:           it.Url,
		Host:          it.Host,
		ResourceName:  it.ResourceName,
		RequestType:   it.RequestType,
		IsValid:       it.IsValid,
		SearchRequest: it.SearchRequest.Clone(),
		PageRequest:   it.PageRequest.Clone(),
	}
}
