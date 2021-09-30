package stringslice

func AppendStringsWithMainSlice(
	isSkipEmpty bool,
	mainSlice []string,
	appendingItems ...string,
) []string {
	if len(appendingItems) == 0 {
		return mainSlice
	}

	for _, item := range appendingItems {
		if isSkipEmpty && item == "" {
			continue
		}

		mainSlice = append(
			mainSlice,
			item)
	}

	return mainSlice
}
