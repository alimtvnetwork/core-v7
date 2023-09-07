package coreapi

import (
	"gitlab.com/auk-go/core/reqtype"
)

type ResponseAttribute struct {
	ResponseOfRequestType reqtype.Request
	Count                 int
	HasAnyRecord          bool
	NextPageRequestUrl    string
	StepsPerformed        *[]string `json:"StepsPerformed,omitempty"`
	DebugInfos            *[]string `json:"DebugInfos,omitempty"`
	HttpCode              int
	HttpMethod            reqtype.Request
	IsValid               bool
	Message               string `json:"Message,omitempty"`
}

func (receiver *ResponseAttribute) Clone() *ResponseAttribute {
	if receiver == nil {
		return nil
	}

	return &ResponseAttribute{
		ResponseOfRequestType: receiver.ResponseOfRequestType,
		Count:                 receiver.Count,
		HasAnyRecord:          receiver.HasAnyRecord,
		NextPageRequestUrl:    receiver.NextPageRequestUrl,
		StepsPerformed:        receiver.StepsPerformed,
		DebugInfos:            receiver.DebugInfos,
		HttpCode:              receiver.HttpCode,
		HttpMethod:            receiver.HttpMethod,
		IsValid:               receiver.IsValid,
		Message:               receiver.Message,
	}
}
