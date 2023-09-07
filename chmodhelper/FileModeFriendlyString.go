package chmodhelper

import (
	"os"
)

func FileModeFriendlyString(
	fileMode os.FileMode,
) string {
	rwxWrapper := New.RwxWrapper.UsingFileMode(
		fileMode)

	return rwxWrapper.FriendlyDisplay()
}
