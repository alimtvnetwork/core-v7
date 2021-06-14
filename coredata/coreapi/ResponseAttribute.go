package coreapi

import "gitlab.com/evatix-go/core/reqtype"

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
	InvalidMessage        string `json:"InvalidMessage,omitempty"`
}
