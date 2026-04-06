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

// UnsafeBytesToStrings
//
// # Returns string arrays from unsafe bytes
//
// Deprecated: This function is fundamentally broken — it reinterprets
// []byte (1-byte elements) as []string (16-byte elements) via unsafe.Pointer,
// producing corrupted memory reads. Do not use.
// Use a proper loop conversion instead.
func UnsafeBytesToStrings(unsafeBytes []byte) []string {
	if unsafeBytes == nil {
		return nil
	}

	results := make([]string, len(unsafeBytes))
	for i, b := range unsafeBytes {
		results[i] = string(b)
	}

	return results
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
