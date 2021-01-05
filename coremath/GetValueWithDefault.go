package coremath

func GetValueWithDefault(isTrue bool, value, defaultVal byte) byte {
	if isTrue {
		return value
	}

	return defaultVal
}
