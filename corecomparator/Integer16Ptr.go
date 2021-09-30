package corecomparator

func Integer16Ptr(left, right *int16) Compare {
	if left == nil && right == nil {
		return Equal
	}

	if left == nil || right == nil {
		return NotEqual
	}

	return Integer16(*left, *right)
}
