package converters

import (
	"unsafe"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/defaulterr"
)

func UnsafeBytesToStringWithErr(unsafeBytes []byte) (string, error) {
	if unsafeBytes == nil {
		return constants.EmptyString, defaulterr.CannotProcessNilOrEmpty
	}

	return *(*string)(unsafe.Pointer(&unsafeBytes)), nil
}

func UnsafeBytesToStringPtr(unsafeBytes []byte) *string {
	if unsafeBytes == nil {
		return nil
	}

	return (*string)(unsafe.Pointer(&unsafeBytes))
}

func UnsafeBytesToString(unsafeBytes []byte) string {
	if unsafeBytes == nil {
		return constants.EmptyString
	}

	return *(*string)(unsafe.Pointer(&unsafeBytes))
}

// UnsafeBytesPtrToStringPtr Returns string from unsafe bytes
//
// May panic on conversion if the bytes were not in unsafe pointer.
//
// Expressions:
// - return (*string)(unsafe.Pointer(allBytes))
func UnsafeBytesPtrToStringPtr(unsafeBytes []byte) *string {
	if unsafeBytes == nil {
		return nil
	}

	return (*string)(unsafe.Pointer(&unsafeBytes))
}
