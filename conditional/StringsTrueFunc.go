package conditional

func StringsTrueFunc(
	isTrue bool,
	trueValueFunc func() []string,
) []string {
	if !isTrue {
		return []string{}
	}

	return trueValueFunc()
}
