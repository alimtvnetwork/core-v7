package conditional

func BoolsPtr(
	isTrue bool,
	value, defaultVal *[]bool,
) *[]bool {
	if isTrue {
		return value
	}

	return defaultVal
}
