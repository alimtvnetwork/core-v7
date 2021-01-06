package conditional

func Int(
	isTrue bool,
	value, defaultVal int,
) int {
	if isTrue {
		return value
	}

	return defaultVal
}
