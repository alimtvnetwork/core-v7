package conditional

// VoidFunctions execute void functions based on condition
func VoidFunctions(
	isTrue bool,
	trueValueFunctions []func(),
	falseValueFunctions []func(),
) {
	if isTrue {
		executeVoidFunctions(trueValueFunctions)
	}

	executeVoidFunctions(falseValueFunctions)
}
