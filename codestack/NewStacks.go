package codestack

import "gitlab.com/evatix-go/core/constants"

func NewStacks(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex, // should start from 1
	stackCount int,
) TraceCollection {
	traces := NewTraceCollection(stackCount + constants.Capacity2)

	return *traces.AddsUsingSkip(
		isSkipInvalid,
		isBreakOnceInvalid,
		startSkipIndex+defaultInternalSkip,
		stackCount)
}
