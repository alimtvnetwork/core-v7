package conditional

// ErrorFunc returns func based on condition
func ErrorFunc(
	isTrue bool,
	trueValueFunction func() error,
	falseValueFunction func() error,
) func() error {
	if isTrue {
		return trueValueFunction
	}

	return falseValueFunction
}
