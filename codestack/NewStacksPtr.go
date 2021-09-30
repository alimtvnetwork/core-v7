package codestack

import "gitlab.com/evatix-go/core/constants"

func NewStacksPtr(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex,
	stackCount int,
) *TraceCollection {
	traces := NewTraceCollection(stackCount + constants.Capacity2)

	return traces.AddsUsingSkip(
		isSkipInvalid,
		isBreakOnceInvalid,
		startSkipIndex+defaultInternalSkip,
		stackCount)
}
