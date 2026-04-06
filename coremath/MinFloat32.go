package coremath

//
//goland:noinspection ALL
func MinFloat32(left, right float32) float32 {
	if left > right {
		return right
	}

	return left
}
