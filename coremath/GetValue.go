package coremath

func GetValue(isTrue bool, value byte) byte {
	if isTrue {
		return value
	}

	return 0
}
