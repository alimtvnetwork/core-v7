package conditional

func Interfaces(
	isTrue bool,
	value, defaultVal []interface{},
) []interface{} {
	if isTrue {
		return value
	}

	return defaultVal
}
