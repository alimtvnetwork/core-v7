package stringslice

// Deprecated: Use MakeLen instead.
func MakeLenPtr(length int) []string {
	return make([]string, length)
}
