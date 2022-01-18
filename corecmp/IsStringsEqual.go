package corecmp

func IsStringsEqual(leftItems, rightItems []string) bool {
	if leftItems == nil && rightItems == nil {
		return true
	}

	if leftItems == nil || rightItems == nil {
		return false
	}

	length := len(leftItems)
	if length != len(rightItems) {
		return false
	}

	for i := 0; i < length; i++ {
		if leftItems[i] != rightItems[i] {
			return false
		}
	}

	return true
}
