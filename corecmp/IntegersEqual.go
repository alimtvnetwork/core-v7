package corecmp

func IntegersEqual(array, other []int) bool {
	if array == nil && other == nil {
		return true
	}

	if array == nil || other == nil {
		return false
	}

	arrayPtr := &array
	otherPtr := &other

	return IntegersEqualPtr(arrayPtr, otherPtr)
}
