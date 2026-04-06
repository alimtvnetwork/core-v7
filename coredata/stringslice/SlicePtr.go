package stringslice

func SlicePtr(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return slice
}
