package conditional

func Strings(
	isTrue bool,
	value, defaultVal []string,
) []string {
	if isTrue {
		return value
	}

	return defaultVal
}
