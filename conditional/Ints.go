package conditional

func Ints(
	isTrue bool,
	value, defaultVal []int,
) []int {
	if isTrue {
		return value
	}

	return defaultVal
}
