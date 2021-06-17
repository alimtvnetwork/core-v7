package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
)

func FileModeApplyChmod(
	fileMode os.FileMode,
	condition *chmodins.Condition,
	locations ...string,
) error {
	rwxWrapper := NewUsingFileMode(fileMode)

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}
