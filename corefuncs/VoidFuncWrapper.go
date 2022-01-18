package corefuncs

type NamedVoidFuncWrapper struct {
	Name   string
	Action NamedVoidFunc
}

func (it NamedVoidFuncWrapper) Exec() {
	it.Action(it.Name)
}

func (it NamedVoidFuncWrapper) Next(next *NamedVoidFuncWrapper) {
	it.Exec()
	next.Exec()
}

type ReturnErrFuncWrapper struct {
	Name   string
	Action ReturnErrFunc
}

type IsSuccessFuncWrapper struct {
	Name   string
	Action IsSuccessFunc
}

type ResultDelegatingFuncWrapper struct {
	Name   string
	Action ResultDelegatingFunc
}
