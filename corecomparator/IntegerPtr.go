package corecomparator

func IntegerPtr(left, right *int) Compare {
	if left == nil && right == nil {
		return Equal
	}

	if left == nil || right == nil {
		return NotEqual
	}

	return Integer(*left, *right)
}
