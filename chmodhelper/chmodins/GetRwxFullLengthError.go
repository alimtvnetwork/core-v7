package chmodins

import (
	"gitlab.com/evatix-go/core/msgtype"
)

// GetRwxFullLengthError must be 10 chars length
func GetRwxFullLengthError(rwxFull string) error {
	if len(rwxFull) != RwxFullLength {
		return msgtype.LengthShouldBeEqualToMessage.
			Error(
				"rwxFull length should be "+RwxFullLengthString,
				len(rwxFull))
	}

	return nil
}
