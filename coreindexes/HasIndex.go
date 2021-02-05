package coreindexes

func HasIndex(indexes *[]int, currentIndex int) bool {
	for _, indexValue := range *indexes {
		if indexValue == currentIndex {
			return true
		}
	}

	return false
}
