package stringslice

func CloneIf(
	isClone bool,
	additionalCap int,
	slice []string,
) (newSlice []string) {
	if slice == nil && !isClone {
		return []string{}
	}

	if !isClone {
		return slice
	}

	return CloneUsingCap(additionalCap, slice)
}
