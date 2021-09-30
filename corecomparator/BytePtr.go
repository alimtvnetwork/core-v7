package corecomparator

func BytePtr(left, right *byte) Compare {
	if left == nil && right == nil {
		return Equal
	}

	if left == nil || right == nil {
		return NotEqual
	}

	return Byte(*left, *right)
}
