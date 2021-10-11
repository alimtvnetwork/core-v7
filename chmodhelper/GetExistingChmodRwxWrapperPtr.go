package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/errcore"
)

func GetExistingChmodRwxWrapperPtr(
	filePath string,
) (*RwxWrapper, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return nil, errcore.PathErrorMessage.
			Error(err.Error(), ", file:"+filePath)
	}

	return NewUsingFileModePtr(fileInfo.Mode()), err
}
