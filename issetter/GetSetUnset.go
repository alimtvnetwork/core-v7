package issetter

func GetSetUnset(
	isCondition bool,
) Value {
	if isCondition {
		return Set
	}

	return Unset
}
