package conditional

func InterfacesPtr(
	isTrue bool,
	value, defaultVal *[]interface{},
) *[]interface{} {
	if isTrue {
		return value
	}

	return defaultVal
}
