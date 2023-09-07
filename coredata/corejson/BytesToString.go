package corejson

func BytesToString(
	jsonBytes []byte,
) string {
	if len(jsonBytes) == 0 {
		return ""
	}

	return string(jsonBytes)
}

func BytesToPrettyString(
	jsonBytes []byte,
) string {
	if len(jsonBytes) == 0 {
		return ""
	}

	rs := Result{
		Bytes: jsonBytes,
	}

	return rs.PrettyJsonString()
}
