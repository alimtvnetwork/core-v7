package codestack

import "gitlab.com/auk-go/core/constants"

type newStacksCreator struct{}

func (it newStacksCreator) All(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex, // should start from 1
	stackCount int,
) TraceCollection {
	traces := New.traces.Cap(stackCount + constants.Capacity2)

	return *traces.AddsUsingSkip(
		isSkipInvalid,
		isBreakOnceInvalid,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}

func (it newStacksCreator) Default(
	startSkipIndex,
	stackCount int,
) TraceCollection {
	return it.All(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}

func (it newStacksCreator) DefaultCount(
	startSkipIndex int,
) TraceCollection {
	return it.All(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		DefaultStackCount,
	)
}

func (it newStacksCreator) SkipOne() TraceCollection {
	return it.All(
		true,
		true,
		Skip1+defaultInternalSkip,
		DefaultStackCount,
	)
}

func (it newStacksCreator) SkipNone() TraceCollection {
	return it.All(
		true,
		true,
		defaultInternalSkip,
		DefaultStackCount,
	)
}
