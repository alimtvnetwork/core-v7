package corecompare

func IntArrayPtr(array, other *[]int) bool {
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

	for i := 0; i < length; i++ {
		if (*array)[i] != (*other)[i] {
			return false
		}
	}

	return true
}
