package coremath

//
//goland:noinspection ALL
func MaxFloat32(left, right float32) float32 {
	if left < right {
		return right
	}

	return left
}
