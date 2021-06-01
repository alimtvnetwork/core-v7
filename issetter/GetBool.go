package issetter

func GetBool(
	isCondition bool,
) Value {
	if isCondition {
		return True
	}

	return False
}
