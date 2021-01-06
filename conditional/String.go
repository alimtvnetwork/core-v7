package conditional

func String(
	isTrue bool,
	value, defaultVal string,
) string {
	if isTrue {
		return value
	}

	return defaultVal
}
