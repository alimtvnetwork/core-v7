package stringslice

func AllElemLengthSlices(slices ...*[]string) int {
	if slices == nil || len(slices) == 0 {
		return 0
	}

	countOfAll := 0

	for _, slice := range slices {
		if slice == nil {
			continue
		}

		countOfAll += len(*slice)
	}

	return countOfAll
}
