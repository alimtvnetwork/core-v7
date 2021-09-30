package corecomparator

func Integer32Ptr(left, right *int32) Compare {
	if left == nil && right == nil {
		return Equal
	}

	if left == nil || right == nil {
		return NotEqual
	}

	return Integer32(*left, *right)
}
