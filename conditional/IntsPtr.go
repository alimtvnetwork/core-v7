package conditional

func IntsPtr(
	isTrue bool,
	value, defaultVal *[]int,
) *[]int {
	if isTrue {
		return value
	}

	return defaultVal
}
