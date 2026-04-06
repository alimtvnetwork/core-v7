package coremath

// Deprecated: Use the built-in max() function (Go 1.21+).
//
//goland:noinspection ALL
func MaxByte(left, right byte) byte {
	if left < right {
		return right
	}

	return left
}
