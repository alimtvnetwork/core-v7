package codestack

func NewStacksDefaultCount(
	startSkipIndex int,
) TraceCollection {
	return NewStacks(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		DefaultStackCount)
}
