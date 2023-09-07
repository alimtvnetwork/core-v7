package corefuncs

import "gitlab.com/auk-go/core/errcore"

type ActionReturnsErrorFuncWrapper struct {
	Name   string
	Action ActionReturnsErrorFunc
}

func (it ActionReturnsErrorFuncWrapper) Exec() error {
	return it.Action()
}

func (it ActionReturnsErrorFuncWrapper) AsActionFunc() ActionFunc {
	return func() {
		errcore.HandleErr(it.AsActionReturnsErrorFunc()())
	}
}

func (it ActionReturnsErrorFuncWrapper) AsActionReturnsErrorFunc() ActionReturnsErrorFunc {
	return func() error {
		err := it.Action()

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return nil
	}
}
