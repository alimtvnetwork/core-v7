package coreindexes

func Of(indexes []int, seekingValue int) int {
	for i, indexValue := range indexes {
		if indexValue == seekingValue {
			return i
		}
	}

	return -1
}
