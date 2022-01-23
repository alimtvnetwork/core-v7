package corejson

func BytesDeepClone(
	inputBytes []byte,
) []byte {
	if len(inputBytes) == 0 {
		return []byte{}
	}

	newBytes := make([]byte, 0, len(inputBytes))
	copy(newBytes, inputBytes)

	return newBytes
}
