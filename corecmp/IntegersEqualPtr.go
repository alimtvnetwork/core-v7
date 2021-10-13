package corecmp

func IntegersEqualPtr(array, other *[]int) bool {
	if array == nil && other == nil {
		return true
	}

	if array == nil || other == nil {
		return false
	}

	length := len(*array)

	if length != len(*other) {
		return false
	}

	return IntegersEqual(*array, *other)
}
