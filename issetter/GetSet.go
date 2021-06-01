package issetter

func GetSet(
	isCondition bool,
	trueValue Value,
	falseValue Value,
) Value {
	if isCondition {
		return trueValue
	}

	return falseValue
}
