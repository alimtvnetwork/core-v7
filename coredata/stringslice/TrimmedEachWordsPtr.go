package stringslice

func TrimmedEachWordsPtr(slice *[]string) *[]string {
	if slice == nil || *slice == nil {
		return &[]string{}
	}

	results := TrimmedEachWords(*slice)

	return &results
}
