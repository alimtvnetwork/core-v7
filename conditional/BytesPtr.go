package conditional

func BytesPtr(
	isTrue bool,
	value, defaultVal *[]byte,
) *[]byte {
	if isTrue {
		return value
	}

	return defaultVal
}
