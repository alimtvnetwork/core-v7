package chmodhelper

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/evatix-go/core/constants"
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

func (it *PathExistStat) LastModifiedDate() *time.Time {
	if it.IsInvalid() {
		return nil
	}

	lastModifiedTime := it.FileInfo.ModTime()

	return &lastModifiedTime
}

func (it *PathExistStat) FileMode() *os.FileMode {
	if it.IsInvalid() {
		return nil
	}

	fileMode := it.FileInfo.Mode()

	return &fileMode
}

func (it *PathExistStat) Size() *int64 {
	if it.IsInvalid() {
		return nil
	}

	size := it.FileInfo.Size()

	return &size
}

func (it *PathExistStat) Split() (dir, filename string) {
	if it.IsInvalid() || it.FileInfo.IsDir() {
		return "", ""
	}

	return filepath.Split(it.Location)
}

func (it *PathExistStat) FileName() (filename string) {
	_, fileName := it.Split()

	return fileName
}

func (it *PathExistStat) ParentDir() (parentDir string) {
	parentDir, _ = it.Split()

	return parentDir
}

func (it *PathExistStat) Parent() *PathExistStat {
	parentDir, _ := it.Split()

	return GetPathExistStat(parentDir)
}

func (it *PathExistStat) ParentWithNewPath(additionalPaths ...string) string {
	parentDir, _ := it.Split()
	slice := append([]string{parentDir}, additionalPaths...)

	return filepath.Join(slice...)
}

func (it *PathExistStat) ParentWithGlobPatternFiles(globPatterns ...string) ([]string, error) {
	filePath := it.ParentWithNewPath(globPatterns...)

	return filepath.Glob(filePath)
}

func (it *PathExistStat) ParentWith(additionalPaths ...string) *PathExistStat {
	return GetPathExistStat(it.ParentWithNewPath(additionalPaths...))
}

func (it *PathExistStat) CombineWithNewPath(additionalPaths ...string) string {
	slice := append([]string{it.Location}, additionalPaths...)

	return filepath.Join(slice...)
}

func (it *PathExistStat) CombineWith(additionalPaths ...string) *PathExistStat {
	return GetPathExistStat(it.CombineWithNewPath(additionalPaths...))
}

func (it *PathExistStat) DotExt() (dotExt string) {
	_, fileName := it.Split()

	return filepath.Ext(fileName)
}

func (it *PathExistStat) Dispose() {
	if it == nil {
		return
	}

	it.Location = constants.EmptyString
	it.IsExist = false
	it.Error = nil
	it.FileInfo = nil
}

func (it *PathExistStat) IsInvalid() bool {
	return it == nil ||
		it.FileInfo == nil ||
		it.Error != nil
}

func (it *PathExistStat) HasAnyIssues() bool {
	return it == nil ||
		it.FileInfo == nil ||
		it.Error != nil
}

func (it *PathExistStat) MeaningFullError() error {
	if it == nil {
		return nil
	}

	if it.IsEmptyError() {
		return nil
	}

	newErrMsg := it.Error.Error() +
		", location :" +
		it.Location

	newErr := errors.New(newErrMsg)
	meaningFulErr := errcore.MeaningfulError(
		errcore.PathInvalidErrorType,
		"PathExistStat",
		newErr,
	)

	return meaningFulErr
}
