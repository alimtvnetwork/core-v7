package conditional

func AnyFunctions(
	isTrue bool,
	trueValueFunctions, falseValueFunctions []func() (
	result interface{},
	isTake,
	isBreak bool,
),
) []func() (
	result interface{},
	isTake,
	isBreak bool,
) {
	if isTrue {
		return trueValueFunctions
	}

	return falseValueFunctions
}
