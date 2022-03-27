package bytetype

func String(rawBytes []byte) string {
	if len(rawBytes) == 0 {
		return ""
	}

	return string(rawBytes)
}
