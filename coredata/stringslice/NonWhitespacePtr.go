package stringslice

// NonWhitespacePtr
//
// Don't include line which is empty or whitespace.
func NonWhitespacePtr(
	slice *[]string,
) *[]string {
	if slice == nil || *slice == nil {
		return &[]string{}
	}

	nonPtrSlice := NonWhitespace(*slice)

	return &nonPtrSlice
}
