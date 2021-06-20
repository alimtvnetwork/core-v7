package conditional

// ErrorFunctionResult execute error function based on condition and returns final error
func ErrorFunctionResult(
	isTrue bool,
	trueValueFunction func() error,
	falseValueFunction func() error,
) error {
	if isTrue {
		return trueValueFunction()
	}

	return falseValueFunction()
}
