package conditional

func Bytes(
	isTrue bool,
	value, defaultVal []byte,
) []byte {
	if isTrue {
		return value
	}

	return defaultVal
}
