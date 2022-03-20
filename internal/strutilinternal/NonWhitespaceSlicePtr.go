package strutilinternal

func NonWhitespaceSlicePtr(
	slice *[]string,
) *[]string {
	if slice == nil || *slice == nil {
		return &[]string{}
	}

	nonPtrSlice := NonWhitespaceSlice(*slice)

	return &nonPtrSlice
}
