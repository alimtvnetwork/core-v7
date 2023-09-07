package converters

func BytesToString(
	rawBytes []byte,
) string {
	if len(rawBytes) == 0 {
		return ""
	}

	return string(rawBytes)
}
