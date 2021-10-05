package conditional

func BooleanTrueFunc(
	isTrue bool,
	trueValueFunc func() bool,
) bool {
	if !isTrue {
		return false
	}

	return trueValueFunc()
}
