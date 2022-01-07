package codestack

type (
	FilterFunc func(trace *Trace) (isTake, isBreak bool)
	Formatter  func(trace *Trace) (output string)
)
