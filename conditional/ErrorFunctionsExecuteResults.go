package conditional

// ErrorFunctionsExecuteResults execute error functions based on condition and returns final error
func ErrorFunctionsExecuteResults(
	isTrue bool,
	trueValueFunctions []func() error,
	falseValueFunctions []func() error,
) error {
	if isTrue {
		return executeErrorFunctions(trueValueFunctions)
	}

	return executeErrorFunctions(falseValueFunctions)
}
