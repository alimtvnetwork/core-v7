package isany

// AnyZero
//
//	returns true if any values is null or zero.
//
// Reference:
//   - Stackoverflow Example : https://stackoverflow.com/a/23555352
func AnyZero(anyItems ...interface{}) bool {
	if len(anyItems) == 0 {
		return true
	}

	for _, item := range anyItems {
		if Zero(item) {
			return true
		}
	}

	return false
}
