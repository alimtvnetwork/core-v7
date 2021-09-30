package corestr

func CloneSliceIf(
	isClone bool,
	sourceItems ...string,
) []string {
	if len(sourceItems) == 0 {
		return []string{}
	}

	if !isClone {
		return sourceItems
	}

	destinationSlice := make(
		[]string,
		len(sourceItems))
	copy(destinationSlice, sourceItems)

	return destinationSlice
}
