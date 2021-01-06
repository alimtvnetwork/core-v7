package conditional

func Bool(
	isTrue bool,
	value, defaultVal bool,
) bool {
	if isTrue {
		return value
	}

	return defaultVal
}
