package stringslice

// Clone on nil or empty makes new  &[]string{}
// else makes a copy of itself
func Clone(slice []string) (newSlice []string) {
	if len(slice) == 0 {
		return []string{}
	}

	return CloneUsingCap(0, slice)
}
