package stringslice

func NonWhitespacePtr(
	slice []string,
) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return NonWhitespace(slice)
}
