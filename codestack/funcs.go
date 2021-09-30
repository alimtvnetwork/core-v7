package codestack

type (
	FilterFunc func(trace *Trace) (isTake, isBreak bool)
)
