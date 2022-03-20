package strutilinternal

func NonWhitespaceSlice(
	slice []string,
) []string {
	if slice == nil {
		return []string{}
	}

	length := len(slice)

	if length == 0 {
		return []string{}
	}

	newSlice := make(
		[]string,
		0,
		length)

	for _, s := range slice {
		if IsEmptyOrWhitespace(s) {
			continue
		}

		newSlice = append(newSlice, s)
	}

	return newSlice
}
