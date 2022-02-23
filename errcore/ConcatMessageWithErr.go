package errcore

import "errors"

func ConcatMessageWithErr(
	errMessage string,
	err error,
) error {
	if err == nil {
		return nil
	}

	return errors.New(errMessage + err.Error())
}
