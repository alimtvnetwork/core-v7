package codestack

func NewStacksDefaultCountSkip1() TraceCollection {
	return NewStacks(
		true,
		true,
		Skip1+defaultInternalSkip,
		DefaultStackCount)
}
