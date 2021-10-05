package conditional

func StringTrueFunc(
	isTrue bool,
	trueValueFunc func() string,
) string {
	if !isTrue {
		return ""
	}

	return trueValueFunc()
}
