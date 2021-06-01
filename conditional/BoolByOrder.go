package conditional

// BoolByOrder returns the first boolean which is true
func BoolByOrder(
	booleans ...bool,
) bool {
	for _, boolean := range booleans {
		if boolean {
			return true
		}
	}

	return false
}
