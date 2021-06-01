package strtype

func GetSet(
	isCondition bool,
	trueValue Variant,
	falseValue Variant,
) Variant {
	if isCondition {
		return trueValue
	}

	return falseValue
}
