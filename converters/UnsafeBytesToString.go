package converters

import (
	"unsafe"

	"gitlab.com/auk-go/core/constants"
)

func UnsafeBytesToString(unsafeBytes []byte) string {
	if unsafeBytes == nil {
		return constants.EmptyString
	}

	return *(*string)(unsafe.Pointer(&unsafeBytes))
}
