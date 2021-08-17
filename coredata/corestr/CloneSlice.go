package corestr

func CloneSlice(items []string) []string {
	if len(items) == 0 {
		return []string{}
	}

	slice := make(
		[]string,
		0,
		len(items))

	slice = append(slice, items...)

	return slice
}
