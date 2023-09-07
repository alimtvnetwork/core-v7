package corefuncs

import "gitlab.com/auk-go/core/errcore"

type IsSuccessFuncWrapper struct {
	Name   string
	Action IsSuccessFunc
}

func (it IsSuccessFuncWrapper) Exec() (isSuccess bool) {
	return it.Action()
}

func (it IsSuccessFuncWrapper) AsActionFunc() ActionFunc {
	return func() {
		it.Action()
	}
}

func (it IsSuccessFuncWrapper) AsActionReturnsErrorFunc() ActionReturnsErrorFunc {
	return func() error {
		isSuccess := it.Action()

		if !isSuccess {
			return errcore.
				FailedToExecuteType.
				ErrorNoRefs("Function Failed to execute :" + it.Name)
		}

		return nil
	}
}
