package corecomparator

func Integer64Ptr(left, right *int64) Compare {
	if left == nil && right == nil {
		return Equal
	}

	if left == nil || right == nil {
		return NotEqual
	}

	return Integer64(*left, *right)
}
