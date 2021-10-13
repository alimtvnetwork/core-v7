package corecmp

func StringsEqualPtr(array, other *[]string) bool {
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

	return StringsEqual(*array, *other)
}
