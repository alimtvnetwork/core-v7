package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/msgtype"
)

func GetExistingChmodWrapper(filePath string) (Wrapper, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return Wrapper{}, msgtype.
			FileErrorMessage.
			Error(err.Error(), ", file:"+filePath)
	}

	return NewUsingFileMode(fileInfo.Mode()), err
}
