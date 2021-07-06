package stringslice

func SafeIndexAtWith(
	slice []string,
	index int,
	defaultVal string,
) string {
	if len(slice) == 0 || len(slice)-1 < index {
		return defaultVal
	}

	return slice[index]
}
