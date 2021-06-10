package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/msgtype"
)

func GetExistingChmodWrapper(
	filePath string,
) (RwxWrapper, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return RwxWrapper{}, msgtype.
			PathErrorMessage.
			Error(err.Error(), ", file:"+filePath)
	}

	return NewUsingFileMode(fileInfo.Mode()), err
}
