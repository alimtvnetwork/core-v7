package strutilinternal

func NonEmptySlicePtr(slice []string) []string {
	length := len(slice)

	if length == 0 {
		return []string{}
	}

	newSlice := make(
		[]string,
		0,
		length)

	for _, s := range slice {
		if s != "" {
			newSlice = append(newSlice, s)
		}
	}

	return newSlice
}
