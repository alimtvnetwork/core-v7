package corecomparator

func Integer32(left, right int32) Compare {
	if left == right {
		return Equal
	} else if left < right {
		return LeftLess
	} else if left > right {
		return LeftGreater
	}

	return NotEqual
}
