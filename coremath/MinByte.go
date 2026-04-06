package coremath

//
//goland:noinspection ALL
func MinByte(left, right byte) byte {
	if left > right {
		return right
	}

	return left
}
