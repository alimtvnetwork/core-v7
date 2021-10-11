package errcore

import "errors"

func ErrorWithCompiledTraceRefToError(
	err error,
	compiledTraces string,
	reference interface{},
) error {
	if err == nil {
		return nil
	}

	compiledErr := ErrorWithCompiledTraceRef(
		err,
		compiledTraces,
		reference)

	return errors.New(compiledErr)
}
