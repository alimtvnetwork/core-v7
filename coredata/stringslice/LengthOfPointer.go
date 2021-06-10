package stringslice

func LengthOfPointer(slices *[]string) int {
	if slices == nil {
		return 0
	}

	return len(*slices)
}
