package issetter

func NewBool(isResult bool) Value {
	if isResult {
		return True
	}

	return False
}
