package corecompare

func IntArray(array, other []int) bool {
	if array == nil && other == nil {
		return true
	}

	if array == nil || other == nil {
		return false
	}

	arrayPtr := &array
	otherPtr := &other

	return IntArrayPtr(arrayPtr, otherPtr)
}
