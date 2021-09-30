package corecomparator

func MinLength(left, right int) (min int) {
	if left < right {
		return left
	}

	return right
}
