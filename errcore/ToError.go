package errcore

import "errors"

func ToError(errMessage string) error {
	if errMessage == "" {
		return nil
	}

	return errors.New(errMessage)
}
