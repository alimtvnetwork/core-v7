package corecmp

func IsIntegersEqual(array, other []int) bool {
	if array == nil && other == nil {
		return true
	}

	if array == nil || other == nil {
		return false
	}

	arrayPtr := &array
	otherPtr := &other

	return IsIntegersEqualPtr(arrayPtr, otherPtr)
}
