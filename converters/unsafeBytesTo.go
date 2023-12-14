package converters

import (
	"unsafe"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/defaulterr"
)

func UnsafeBytesToStringWithErr(unsafeBytes []byte) (string, error) {
	if unsafeBytes == nil {
		return constants.EmptyString, defaulterr.CannotProcessNilOrEmpty
	}

	return *(*string)(unsafe.Pointer(&unsafeBytes)), nil
}

// UnsafeBytesToStrings
//
// # Returns string arrays from unsafe bytes pointer
//
// May panic on conversion if the bytes were not in unsafe pointer.
//
// Expressions:
// - return (*[] string)(unsafe.Pointer(allBytes))
func UnsafeBytesToStrings(unsafeBytes *[]byte) *[]string {
	if unsafeBytes == nil || *unsafeBytes == nil {
		return nil
	}

	return (*[]string)(unsafe.Pointer(unsafeBytes))
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
