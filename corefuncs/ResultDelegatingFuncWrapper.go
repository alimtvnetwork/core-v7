package corefuncs

import "gitlab.com/auk-go/core/errcore"

type ResultDelegatingFuncWrapper struct {
	Name   string
	Action ResultDelegatingFunc
}

func (it ResultDelegatingFuncWrapper) Exec(
	toPointer interface{},
) error {
	return it.Action(toPointer)
}

func (it ResultDelegatingFuncWrapper) AsActionFunc(toPointer interface{}) ActionFunc {
	return func() {
		actionReturnsErrFunc := it.AsActionReturnsErrorFunc(toPointer)
		errcore.HandleErr(actionReturnsErrFunc())
	}
}

func (it ResultDelegatingFuncWrapper) AsActionReturnsErrorFunc(
	toPointer interface{},
) ActionReturnsErrorFunc {
	return func() error {
		err := it.Action(toPointer)

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return nil
	}
}
