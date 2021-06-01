package bytetype

func GetSetVariant(
	isCondition bool,
	trueValue byte,
	falseValue byte,
) Variant {
	if isCondition {
		return Variant(trueValue)
	}

	return Variant(falseValue)
}
