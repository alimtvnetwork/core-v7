package errcore

import "errors"

func CompiledError(mainErr error, additionalMessage string) error {
	if mainErr == nil {
		return nil
	}

	return errors.New(
		CompiledErrorString(mainErr, additionalMessage))
}
