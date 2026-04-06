package stringslice

// Deprecated: Use SafeRangeItems instead.
func SafeRangeItemsPtr(
	slice []string,
	start, end int,
) []string {
	if len(slice) == 0 {
		return []string{}
	}

	return SafeRangeItems(slice, start, end)
}
