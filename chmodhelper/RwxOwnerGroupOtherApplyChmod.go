package chmodhelper

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/errcore"
)

// RwxOwnerGroupOtherApplyChmod rwxFullString 10 chars "-rwxrwxrwx"
func RwxOwnerGroupOtherApplyChmod(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if rwxOwnerGroupOther == nil {
		return errcore.CannotBeNilOrEmptyMessage.
			ErrorNoRefs("rwxOwnerGroupOther")
	}

	if condition == nil {
		return errcore.CannotBeNilOrEmptyMessage.
			ErrorNoRefs("condition")
	}

	rwxWrapper, err := NewUsingHyphenedRwxFullString(
		rwxOwnerGroupOther.String())

	if err != nil {
		return err
	}

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}
