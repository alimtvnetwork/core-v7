package coremath

//
//goland:noinspection ALL
func MaxByte(left, right byte) byte {
	if left < right {
		return right
	}

	return left
}
