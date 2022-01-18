package converters

func BytesPointerToBytes(fromBytesPointer *[]byte) []byte {
	if fromBytesPointer == nil || *fromBytesPointer == nil {
		return []byte{}
	}

	return *fromBytesPointer
}
