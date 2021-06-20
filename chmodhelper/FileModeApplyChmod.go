package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/msgtype"
)

func FileModeApplyChmod(
	fileMode os.FileMode,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if condition == nil {
		return msgtype.CannotBeNilOrEmptyMessage.
			ErrorNoRefs("condition")
	}

	rwxWrapper := NewUsingFileMode(fileMode)

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}
