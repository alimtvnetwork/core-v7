package chmodhelper

import (
	"fmt"
	"os"
)

func pathErrorMessage(
	message string,
	applyChmod os.FileMode,
	location string,
	err error,
) string {
	return fmt.Sprintf(
		"%s, path : %q, chmod: %q, is-file-exist: %v, err: %s",
		message,
		location,
		FileModeFriendlyString(applyChmod),
		IsPathExists(location),
		err.Error())
}
