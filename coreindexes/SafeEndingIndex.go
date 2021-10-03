package coreindexes

func SafeEndingIndex(length, lastTakingIndex int) int {
	lastIndex := length - 1

	if lastIndex < lastTakingIndex {
		return lastIndex
	}

	return lastTakingIndex
}
