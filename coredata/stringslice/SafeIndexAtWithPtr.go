package stringslice

func SafeIndexAtWithPtr(
	slice *[]string,
	index int,
	defaultVal string,
) string {
	if IsEmptyPtr(slice) || len(*slice)-1 < index {
		return defaultVal
	}

	return (*slice)[index]
}
