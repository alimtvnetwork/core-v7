package corecomparator

func Integer64(left, right int64) Compare {
	if left == right {
		return Equal
	} else if left < right {
		return LeftLess
	} else if left > right {
		return LeftGreater
	}

	return NotEqual
}
