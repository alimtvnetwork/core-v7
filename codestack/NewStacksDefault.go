package codestack

func NewStacksDefault(
	startSkipIndex,
	stackCount int,
) TraceCollection {
	return NewStacks(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		stackCount)
}
