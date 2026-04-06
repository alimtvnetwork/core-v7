package stringslice

// SafeIndexesPtr Only indexes which are present values will be included.
//
// Warning : Not found indexes will not be included in the values.
//
// Deprecated: Use SafeIndexes instead.
func SafeIndexesPtr(slice []string, indexes ...int) (values []string) {
	if IsEmptyPtr(slice) {
		return MakeLenPtr(len(indexes))
	}

	return SafeIndexes(slice, indexes...)
}
