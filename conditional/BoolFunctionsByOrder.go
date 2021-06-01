package conditional

// BoolFunctionsByOrder returns the first func which is true and don't execute others.
func BoolFunctionsByOrder(
	boolFunctions ...func() bool,
) bool {
	for _, f := range boolFunctions {
		if f() {
			return true
		}
	}

	return false
}
