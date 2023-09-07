package strutilinternal

func NonWhitespaceTrimSlicePtr(slice *[]string) *[]string {
	if slice == nil || *slice == nil {
		return &[]string{}
	}

	results := NonWhitespaceTrimSlice(*slice)

	return &results
}
