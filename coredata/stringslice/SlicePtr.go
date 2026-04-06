package stringslice

// Deprecated: No longer needed - slices are used directly.
func SlicePtr(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return slice
}
