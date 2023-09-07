package issetter

func IntegerEnumRanges() []int {
	slice := make([]int, len(valuesNames))

	for i := range valuesNames {
		slice[i] = i
	}

	return slice
}
