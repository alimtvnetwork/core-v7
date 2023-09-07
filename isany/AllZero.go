package isany

// AllZero
//
//  returns true if all values is null or zero or given array is nil.
//
// Reference:
//  - Stackoverflow Example : https://stackoverflow.com/a/23555352
func AllZero(anyItems ...interface{}) bool {
	if len(anyItems) == 0 {
		return true
	}

	for _, item := range anyItems {
		if !Zero(item) {
			return false
		}
	}

	return true
}
