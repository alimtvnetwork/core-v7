package codestack

func NewSkip1() Trace {
	return New(Skip2)
}
