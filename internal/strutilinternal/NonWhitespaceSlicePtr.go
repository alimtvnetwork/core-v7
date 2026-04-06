package strutilinternal

func NonWhitespaceSlicePtr(
	slice []string,
) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return NonWhitespaceSlice(slice)
}
