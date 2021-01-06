package conditional

func StringPtr(
	isTrue bool,
	value, defaultVal *string,
) *string {
	if isTrue {
		return value
	}

	return defaultVal
}
