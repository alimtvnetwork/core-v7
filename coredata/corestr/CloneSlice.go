package corestr

func CloneSlice(items []string) []string {
	if len(items) == 0 {
		return []string{}
	}

	slice := make(
		[]string,
		len(items))
	copy(slice, items)

	return slice
}
