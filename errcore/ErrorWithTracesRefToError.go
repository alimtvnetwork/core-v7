package errcore

func ErrorWithTracesRefToError(
	err error,
	traces []string,
	reference interface{},
) error {
	if err == nil {
		return nil
	}

	if len(traces) == 0 {
		return ErrorWithRefToError(err, reference)
	}

	tracesCompiled := StackTracesCompiled(traces)

	return ErrorWithCompiledTraceRefToError(
		err,
		tracesCompiled,
		reference)
}
