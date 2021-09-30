package stringslice

func CloneUsingCap(newAdditionalCap int, slice []string) []string {
	newSlice := make(
		[]string,
		0,
		newAdditionalCap+len(slice))

	if len(slice) == 0 {
		return newSlice
	}

	newSlice = append(newSlice, slice...)

	return newSlice
}
