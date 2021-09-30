package stringslice

func NonEmptyStrings(slice []string) []string {
	if slice == nil {
		return []string{}
	}

	length := len(slice)

	if length == 0 {
		return []string{}
	}

	sliceNew := make([]string, 0, length)

	for _, s := range slice {
		if s == "" {
			continue
		}

		sliceNew = append(sliceNew, s)
	}

	return sliceNew
}
