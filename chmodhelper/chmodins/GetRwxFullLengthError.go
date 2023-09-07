package chmodins

import (
	"gitlab.com/auk-go/core/errcore"
)

// GetRwxFullLengthError must be 10 chars length
func GetRwxFullLengthError(rwxFull string) error {
	if len(rwxFull) != RwxFullLength {
		return errcore.LengthShouldBeEqualToType.
			Error(
				"rwxFull length should be "+RwxFullLengthString,
				len(rwxFull))
	}

	return nil
}
