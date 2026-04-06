package coremath

// Deprecated: Use the built-in min() function (Go 1.21+).
//
//goland:noinspection ALL
func MinFloat32(left, right float32) float32 {
	if left > right {
		return right
	}

	return left
}
