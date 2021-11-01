package coreindexes

func IsWithinIndexRange(index, length int) bool {
	return length-1 >= index
}
