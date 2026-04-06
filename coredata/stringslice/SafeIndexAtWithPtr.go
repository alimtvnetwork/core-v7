package stringslice

// Deprecated: Use SafeIndexAtWith instead.
func SafeIndexAtWithPtr(
	slice []string,
	index int,
	defaultVal string,
) string {
	if IsEmptyPtr(slice) || index < 0 || len(slice)-1 < index {
		return defaultVal
	}

	return slice[index]
}
