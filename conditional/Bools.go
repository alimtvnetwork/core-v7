package conditional

func Bools(
	isTrue bool,
	value, defaultVal []bool,
) []bool {
	if isTrue {
		return value
	}

	return defaultVal
}
