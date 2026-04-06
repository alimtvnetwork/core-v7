package stringslice

// Deprecated: Use HasAnyItem instead.
func HasAnyItemPtr(slice []string) bool {
	return !IsEmptyPtr(slice)
}
