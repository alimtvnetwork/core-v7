package converters

func BytesPtrToString(
	rawBytes *[]byte,
) string {
	if rawBytes == nil || len(*rawBytes) == 0 {
		return ""
	}

	return string(*rawBytes)
}
