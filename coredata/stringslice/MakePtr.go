package stringslice

// Deprecated: Use Make instead.
func MakePtr(length, capacity int) []string {
	return make([]string, length, capacity)
}
