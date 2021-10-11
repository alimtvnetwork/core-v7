package chmodhelper

import (
	"errors"
	"os"

	"gitlab.com/evatix-go/core/errcore"
)

type PathExistStat struct {
	Location string
	FileInfo os.FileInfo
	IsExist  bool
	Error    error
}

func (it *PathExistStat) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *PathExistStat) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *PathExistStat) HasFileInfo() bool {
	return it != nil && it.FileInfo != nil
}

func (it *PathExistStat) IsFile() bool {
	return it.HasFileInfo() && !it.FileInfo.IsDir()
}

func (it *PathExistStat) IsDir() bool {
	return it.HasFileInfo() && it.FileInfo.IsDir()
}

func (it *PathExistStat) MeaningFullError() error {
	if it.IsEmptyError() {
		return nil
	}

	newErrMsg := it.Error.Error() +
		", location :" +
		it.Location

	newErr := errors.New(newErrMsg)
	meaningFulErr := errcore.MeaningfulError(
		errcore.PathInvalidErrorMessage,
		"PathExistStat",
		newErr,
	)

	return meaningFulErr
}
