package stringslice

// SafeIndexesDefaultPtr Only indexes which are present values will be included.
//
// Warning : Not found indexes will not be included in the values.
func SafeIndexesPtr(slice *[]string, indexes ...int) (values *[]string) {
	if IsEmptyPtr(slice) {
		return MakeLenPtr(len(indexes))
	}

	values2 := SafeIndexes(*slice, indexes...)

	return &values2
}
