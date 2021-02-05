package coreindexes

func Of(indexes *[]int, currentIndex int) int {
	for i, indexValue := range *indexes {
		if indexValue == currentIndex {
			return i
		}
	}

	return -1
}
