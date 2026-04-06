package stringslice

// Deprecated: Use IsEmpty instead.
func IsEmptyPtr(slice []string) bool {
	return len(slice) == 0
}
