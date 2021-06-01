package issetter

func GetSetByte(
	isCondition bool,
	trueValue byte,
	falseValue byte,
) Value {
	if isCondition {
		return Value(trueValue)
	}

	return Value(falseValue)
}
