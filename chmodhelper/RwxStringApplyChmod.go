package chmodhelper

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/msgtype"
)

// RwxStringApplyChmod rwxFullString 10 chars "-rwxrwxrwx"
func RwxStringApplyChmod(
	rwxFullString string, // rwxFullString 10 chars "-rwxrwxrwx"
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	rwxFullStringErr := chmodins.GetRwxFullLengthError(rwxFullString)
	if rwxFullStringErr != nil {
		return rwxFullStringErr
	}

	if condition == nil {
		return msgtype.
			CannotBeNilOrEmptyMessage.
			ErrorNoRefs("condition")
	}

	rwxWrapper, err := NewUsingHyphenedRwxFullString(rwxFullString)

	if err != nil {
		return err
	}

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}
