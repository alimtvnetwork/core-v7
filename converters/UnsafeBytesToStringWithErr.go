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
