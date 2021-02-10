package coremath

//goland:noinspection ALL
func MinByte(v1, v2 byte) byte {
	if v1 > v2 {
		return v2
	}

	return v1
}
