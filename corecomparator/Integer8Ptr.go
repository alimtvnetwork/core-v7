package corecomparator

func Integer8Ptr(left, right *int8) Compare {
	if left == nil && right == nil {
		return Equal
	}

	if left == nil || right == nil {
		return NotEqual
	}

	return Integer8(*left, *right)
}
