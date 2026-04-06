package stringslice

// Deprecated: Use FirstLastDefaultStatus instead.
func FirstLastDefaultStatusPtr(slice []string) *FirstLastStatus {
	if IsEmptyPtr(slice) {
		return InvalidFirstLastStatus()
	}

	return FirstLastDefaultStatus(slice)
}
