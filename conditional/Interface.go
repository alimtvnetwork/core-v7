package conditional

func Interface(
	isTrue bool,
	value, defaultVal interface{},
) interface{} {
	if isTrue {
		return value
	}

	return defaultVal
}
