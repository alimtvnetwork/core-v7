package conditional

func BytesTrueFunc(
	isTrue bool,
	trueValueFunc func() []byte,
) []byte {
	if !isTrue {
		return []byte{}
	}

	return trueValueFunc()
}
