package codestack

func NewStacksDefaultCountSkipNone() TraceCollection {
	return NewStacks(
		true,
		true,
		defaultInternalSkip,
		DefaultStackCount)
}
