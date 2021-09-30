package corecomparator

func Integer(left, right int) Compare {
	if left == right {
		return Equal
	} else if left < right {
		return LeftLess
	} else if left > right {
		return LeftGreater
	}

	return NotEqual
}
