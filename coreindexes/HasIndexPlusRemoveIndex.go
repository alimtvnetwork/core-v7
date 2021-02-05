package coreindexes

func HasIndexPlusRemoveIndex(indexes *[]int, currentIndex int) bool {
	for i, indexValue := range *indexes {
		if indexValue == currentIndex {
			*indexes = append(
				(*indexes)[:i],
				(*indexes)[i+1:]...)

			return true
		}
	}

	return false
}
