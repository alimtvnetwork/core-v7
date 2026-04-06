package coremath

// Deprecated: Use the built-in min() function (Go 1.21+).
//
//goland:noinspection ALL
func MinByte(left, right byte) byte {
	if left > right {
		return right
	}

	return left
}
