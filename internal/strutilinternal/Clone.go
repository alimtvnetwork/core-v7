package strutilinternal

func Clone(
	currentSlice []string,
) []string {
	if len(currentSlice) == 0 {
		return []string{}
	}

	newList := make([]string, len(currentSlice))
	copy(newList, currentSlice)

	return newList
}
