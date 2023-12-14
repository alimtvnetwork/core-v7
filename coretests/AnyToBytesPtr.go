package coretests

func AnyToBytesPtr(anyItem interface{}) *[]byte {
	toBytes := AnyToBytes(anyItem)

	return &toBytes
}
