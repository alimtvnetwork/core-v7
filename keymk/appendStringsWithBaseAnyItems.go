package keymk

func appendStringsWithBaseAnyItems(
	isSkipOnEmpty bool,
	mainSlice []interface{},
	appendingItems []string,
) []interface{} {
	if len(appendingItems) == 0 {
		return mainSlice
	}

	for _, item := range appendingItems {
		if isSkipOnEmpty && item == "" {
			continue
		}

		mainSlice = append(
			mainSlice,
			item)
	}

	return mainSlice
}
