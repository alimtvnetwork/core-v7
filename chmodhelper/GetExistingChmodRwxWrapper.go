package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/errcore"
)

func GetExistingChmodRwxWrapper(
	filePath string,
) (RwxWrapper, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return RwxWrapper{}, errcore.
			PathErrorMessage.
			Error(err.Error(), ", file:"+filePath)
	}

	return NewUsingFileMode(fileInfo.Mode()), err
}
