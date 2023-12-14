package chmodhelper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/namevalue"
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

func (it *PathExistStat) IsInvalidFileInfo() bool {
	return it == nil || it.FileInfo == nil
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
		!it.IsExist ||
		it.FileInfo == nil ||
		it.Error != nil
}

func (it *PathExistStat) HasAnyIssues() bool {
	return it == nil ||
		!it.IsExist ||
		it.FileInfo == nil ||
		it.Error != nil
}

func (it *PathExistStat) NotExistError() error {
	if it == nil {
		return nil
	}

	if !it.IsExist || it.FileInfo == nil {
		return it.MeaningFullError()
	}

	return nil
}

func (it *PathExistStat) MessageWithPathWrapped(
	message string,
) string {
	if it == nil {
		return ""
	}

	return fmt.Sprintf(
		messageWithPathWrappedFormat,
		message,
		it.Location)
}

// NotAFileError
//
// Get error on:
//   - Path has issues or not exist
//   - Expecting file, if not file then error
func (it *PathExistStat) NotAFileError() error {
	if it == nil {
		return nil
	}

	if !it.IsExist {
		return it.NotExistError()
	}

	if it.IsDir() {
		return errcore.ExpectingSimpleNoTypeError(
			it.MessageWithPathWrapped("Expecting file but received directory."),
			"File",
			"Directory")
	}

	return nil
}

// NotADirError
//
// Get error on:
//   - Path has issues or not exist
//   - Expecting dir, if not dir then error
func (it *PathExistStat) NotADirError() error {
	if it == nil {
		return nil
	}

	if !it.IsExist {
		return it.NotExistError()
	}

	if it.IsFile() {
		return errcore.ExpectingSimpleNoTypeError(
			it.MessageWithPathWrapped("Expecting directory but received file."),
			"Directory",
			"File")
	}

	return nil
}

func (it *PathExistStat) MeaningFullError() error {
	if it == nil {
		return nil
	}

	if it.IsEmptyError() {
		return nil
	}

	newErrMsg := it.Error.Error() +
		" Location :" +
		it.Location

	newErr := errors.New(newErrMsg)
	meaningFulErr := errcore.MeaningfulError(
		errcore.PathInvalidErrorType,
		"Function : {PathExistStat.MeaningFullError()}",
		newErr,
	)

	return meaningFulErr
}

func (it *PathExistStat) String() string {
	if it == nil {
		return ""
	}

	slice := errcore.VarNameValuesStrings(
		namevalue.Instance{
			Name:  "Location",
			Value: it.Location,
		},
		namevalue.Instance{
			Name:  "Name",
			Value: it.FileName(),
		},
		namevalue.Instance{
			Name:  "IsExist",
			Value: it.IsExist,
		},
		namevalue.Instance{
			Name:  "IsFile",
			Value: it.IsFile(),
		},
		namevalue.Instance{
			Name:  "IsDir",
			Value: it.IsDir(),
		},
		namevalue.Instance{
			Name:  "Chmod",
			Value: it.FileMode(),
		},
		namevalue.Instance{
			Name:  "Error",
			Value: it.Error,
		})

	return constants.NewLineSpaceHyphenSpace + strings.Join(
		slice,
		constants.IndentFileInfoEachLineJoiner)
}
