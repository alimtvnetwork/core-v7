package converters

import "unsafe"

func UnsafeBytesToStringPtr(unsafeBytes []byte) *string {
	if unsafeBytes == nil {
		return nil
	}

	return (*string)(unsafe.Pointer(&unsafeBytes))
}
