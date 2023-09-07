package corefuncs

import "gitlab.com/auk-go/core/errcore"

type InOutErrFuncWrapper struct {
	Name   string
	Action InOutErrFunc
}

func (it InOutErrFuncWrapper) Exec(
	input interface{},
) (output interface{}, err error) {
	return it.Action(input)
}

func (it InOutErrFuncWrapper) AsActionFunc(input interface{}) ActionFunc {
	return func() {
		errcore.MustBeEmpty(
			it.AsActionReturnsErrorFunc(input)())
	}
}

func (it InOutErrFuncWrapper) AsActionReturnsErrorFunc(
	input interface{},
) ActionReturnsErrorFunc {
	return func() error {
		_, err := it.Action(input)

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return err
	}
}
