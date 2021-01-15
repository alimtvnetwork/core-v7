package conditional

func StringsPtr(
	isTrue bool,
	value, defaultVal *[]string,
) *[]string {
	if isTrue {
		return value
	}

	return defaultVal
}
