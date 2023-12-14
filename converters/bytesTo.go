package converters

type bytesTo struct{}

func (it bytesTo) PtrString(
	rawBytes *[]byte,
) string {
	if rawBytes == nil || len(*rawBytes) == 0 {
		return ""
	}

	return string(*rawBytes)
}

func (it bytesTo) String(
	rawBytes []byte,
) string {
	if len(rawBytes) == 0 {
		return ""
	}

	return string(rawBytes)
}

func (it bytesTo) PointerToBytes(fromBytesPointer *[]byte) []byte {
	if fromBytesPointer == nil || *fromBytesPointer == nil {
		return []byte{}
	}

	return *fromBytesPointer
}
