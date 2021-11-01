package codestack

func NewDefault() Trace {
	return New(defaultInternalSkip)
}
