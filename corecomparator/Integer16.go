package corecomparator

func Integer16(left, right int16) Compare {
	if left == right {
		return Equal
	} else if left < right {
		return LeftLess
	} else if left > right {
		return LeftGreater
	}

	return NotEqual
}

func Integer8(left, right int8) Compare {
	if left == right {
		return Equal
	} else if left < right {
		return LeftLess
	} else if left > right {
		return LeftGreater
	}

	return NotEqual
}
