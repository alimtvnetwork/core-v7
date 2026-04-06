package strutilinternal

// Deprecated: Use NonWhitespaceSlice instead.
func NonWhitespaceSlicePtr(
	slice []string,
) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return NonWhitespaceSlice(slice)
}
