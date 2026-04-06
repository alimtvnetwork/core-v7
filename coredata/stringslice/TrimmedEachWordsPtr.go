package stringslice

// Deprecated: Use TrimmedEachWords instead.
func TrimmedEachWordsPtr(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return TrimmedEachWords(slice)
}
