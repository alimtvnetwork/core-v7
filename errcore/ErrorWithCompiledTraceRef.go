package errcore

import "fmt"

func ErrorWithCompiledTraceRef(
	err error,
	compiledTraces string,
	reference interface{},
) string {
	if err == nil {
		return ""
	}

	if compiledTraces == "" {
		return ErrorWithRef(err, reference)
	}

	if reference == nil {
		return fmt.Sprintf(
			messageWithTracesWithoutRefFormat,
			err.Error(),
			compiledTraces,
		)
	}

	return fmt.Sprintf(
		messageWithTracesRefFormat,
		err.Error(),
		compiledTraces,
		reference)
}
