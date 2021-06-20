package conditional

func AnyFunctionsExecuteResults(
	isTrue bool,
	trueValueFunctions, falseValueFunctions []func() (
	result interface{},
	isTake,
	isBreak bool,
),
) []interface{} {
	if isTrue {
		return executeAnyFunctions(trueValueFunctions)
	}

	return executeAnyFunctions(falseValueFunctions)
}
