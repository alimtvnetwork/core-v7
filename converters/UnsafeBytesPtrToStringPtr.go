package converters

import "unsafe"

// UnsafeBytesPtrToStringPtr Returns string from unsafe bytes pointer
//
// May panic on conversion if the bytes were not in unsafe pointer.
//
// Expressions:
// - return (*string)(unsafe.Pointer(allBytes))
func UnsafeBytesPtrToStringPtr(unsafeBytes *[]byte) *string {
	if unsafeBytes == nil || *unsafeBytes == nil {
		return nil
	}

	return (*string)(unsafe.Pointer(unsafeBytes))
}
