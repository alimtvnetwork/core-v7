package converters

import (
	"unsafe"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaulterr"
)

func UnsafeBytesToStringWithErr(unsafeBytes []byte) (string, error) {
	if unsafeBytes == nil {
		return constants.EmptyString, defaulterr.CannotProcessNilOrEmpty
	}

	return *(*string)(unsafe.Pointer(&unsafeBytes)), nil
}
