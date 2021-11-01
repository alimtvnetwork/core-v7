package coreinstruction

type BaseContinueOnError struct {
	IsContinueOnError bool
}

func (it *BaseContinueOnError) IsExitOnError() bool {
	return it != nil && !it.IsContinueOnError
}
