package conditional

func Byte(
	isTrue bool,
	value, defaultVal byte,
) byte {
	if isTrue {
		return value
	}

	return defaultVal
}
