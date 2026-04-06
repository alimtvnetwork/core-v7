package coretests

// Deprecated: Use AnyToBytes instead. Returns []byte directly.
func AnyToBytesPtr(anyItem any) []byte {
	return AnyToBytes(anyItem)
}
