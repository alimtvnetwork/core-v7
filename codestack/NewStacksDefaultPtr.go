package codestack

func NewStacksDefaultPtr(
	startSkipIndex,
	stackCount int,
) *TraceCollection {
	return NewStacksPtr(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		stackCount)
}
