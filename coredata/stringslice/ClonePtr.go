package stringslice

func ClonePtr(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}

	newSlice := make([]string, len(slice))
	copy(newSlice, slice)

	return newSlice
}
