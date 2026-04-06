package strutilinternal

func NonWhitespaceTrimSlicePtr(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return NonWhitespaceTrimSlice(slice)
}
