package chmodhelper

import (
	"os"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

type FileWithContent struct {
	RelativePath string
	FileMode     os.FileMode // default for file fileDefaultChmod
	Content      []string
}

func (it FileWithContent) AbsPath(parentDir string) string {
	return pathinternal.Join(parentDir, it.RelativePath)
}

func (it FileWithContent) ContentToString() string {
	return strings.Join(it.Content, constants.NewLineUnix)
}

func (it FileWithContent) ContentToBytes() []byte {
	return []byte(it.ContentToString())
}

func (it FileWithContent) Read(parentDir string) ([]byte, error) {
	return SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              it.FileMode,
		ParentDir:              parentDir,
		FilePath:               it.AbsPath(parentDir),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}.Read()
}

func (it FileWithContent) ReadString(parentDir string) (string, error) {
	return SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              it.FileMode,
		ParentDir:              parentDir,
		FilePath:               it.AbsPath(parentDir),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}.ReadString()
}

func (it FileWithContent) ReadLines(parentDir string) ([]string, error) {
	toStr, err := it.ReadString(parentDir)

	if err != nil {
		return []string{}, err
	}

	// safe

	return strings.Split(toStr, constants.NewLineUnix), nil
}

func (it FileWithContent) Write(parentDir string) error {
	return SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              it.FileMode,
		ParentDir:              parentDir,
		FilePath:               it.AbsPath(parentDir),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}.Write(it.ContentToBytes())
}
