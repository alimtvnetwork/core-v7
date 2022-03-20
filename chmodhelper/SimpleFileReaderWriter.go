package chmodhelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"gitlab.com/evatix-go/core/coredata/corejson"
)

type SimpleFileReaderWriter struct {
	ChmodDir, ChmodFile os.FileMode
	ParentDir           string
	FilePath            string
}

func NewSimpleFileReaderWriter(
	chmodDir,
	chmodFile os.FileMode,
	parentDir,
	filePath string,
) *SimpleFileReaderWriter {
	return &SimpleFileReaderWriter{
		ChmodDir:  chmodDir,
		ChmodFile: chmodFile,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}

// NewSimpleFileReaderWriterDefault
//
//  Default chmod dir - 0755, file - 0644
func NewSimpleFileReaderWriterDefault(
	filePath string,
) *SimpleFileReaderWriter {
	parentDir := filepath.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:  dirDefaultChmod,
		ChmodFile: fileDefaultChmod,
		ParentDir: parentDir,
		FilePath:  filePath,
	}

}

func (it SimpleFileReaderWriter) IsParentExist() bool {
	return IsPathExists(it.ParentDir)
}

func (it SimpleFileReaderWriter) IsExist() bool {
	return IsPathExists(it.FilePath)
}

func (it SimpleFileReaderWriter) HasPathIssues() bool {
	return !IsPathExists(it.FilePath)
}

func (it SimpleFileReaderWriter) IsPathInvalid() bool {
	return !IsPathExists(it.FilePath)
}

func (it SimpleFileReaderWriter) IsParentDirInvalid() bool {
	return !IsPathExists(it.ParentDir)
}

// HasAnyIssues
//
//  it.IsPathInvalid() || it.IsParentDirInvalid()
func (it SimpleFileReaderWriter) HasAnyIssues() bool {
	return it.IsPathInvalid() || it.IsParentDirInvalid()
}

func (it *SimpleFileReaderWriter) ChmodApplier() chmodApplier {
	return chmodApplier{
		rw: it,
	}
}

func (it SimpleFileReaderWriter) Write(allBytes []byte) error {
	err := SimpleFileWriter.WriteFile(
		it.ChmodDir,
		it.ChmodFile,
		true,
		it.ParentDir,
		it.FilePath,
		allBytes)

	if err == nil {
		return nil
	}

	return it.errorWrap(err)
}

func (it SimpleFileReaderWriter) WriteString(content string) error {
	err := SimpleFileWriter.WriteFileString(
		it.ChmodDir,
		it.ChmodFile,
		true,
		it.ParentDir,
		it.FilePath,
		content)

	if err == nil {
		return nil
	}

	return it.errorWrap(err)
}

func (it SimpleFileReaderWriter) errorWrap(err error) error {
	if err == nil {
		return nil
	}

	message := fmt.Sprintf(
		"err: %s, simple-reader-writer: %s",
		err.Error(),
		it.String(),
	)

	return errors.New(message)
}

func (it SimpleFileReaderWriter) WriteAny(
	anyItem interface{},
) error {
	err := SimpleFileWriter.WriteAnyItem(
		it.ChmodDir,
		it.ChmodFile,
		it.ParentDir,
		it.FilePath,
		anyItem)

	if err == nil {
		return nil
	}

	return it.errorWrap(err)
}

func (it SimpleFileReaderWriter) Read() ([]byte, error) {
	allBytes, err := ioutil.ReadFile(it.FilePath)

	if err == nil {
		return allBytes, err
	}

	message := fmt.Sprintf(
		"cannot read file : %q, err: %s, simple-reader-writer: %s",
		it.FilePath,
		err.Error(),
		it.String(),
	)

	return allBytes, errors.New(message)
}

func (it SimpleFileReaderWriter) ReadOnExist() ([]byte, error) {
	if it.IsExist() {
		return it.Read()
	}

	return nil, nil
}

func (it SimpleFileReaderWriter) Get(toPtr interface{}) error {
	if it.IsExist() {
		return it.getOnExist(toPtr)
	}

	return it.errorWrap(errors.New("cannot read cache, save first, file not exist : " + it.FilePath))
}

func (it SimpleFileReaderWriter) GetSet(
	toPtr interface{},
	onInvalidGenerateFunc func() (interface{}, error),
) error {
	if it.IsExist() {
		return it.getOnExist(toPtr)
	}

	newAnyItem, err := onInvalidGenerateFunc()

	if err == nil {
		// if things are all right
		reflect.ValueOf(toPtr).Set(reflect.ValueOf(newAnyItem))

		return nil
	}

	return it.errorWrap(errors.New("read cache failed + cannot generate: " + err.Error()))
}

// Deserialize
//
//  alias for Get
func (it SimpleFileReaderWriter) Deserialize(toPtr interface{}) error {
	return it.Get(toPtr)
}

// Serialize
//
//  alias for ReadOnExist
func (it SimpleFileReaderWriter) Serialize() ([]byte, error) {
	return it.ReadOnExist()
}

// Set
//
//  alias for WriteAny
func (it SimpleFileReaderWriter) Set(toPtr interface{}) error {
	return it.WriteAny(toPtr)
}

func (it SimpleFileReaderWriter) Expire() error {
	if it.IsExist() {
		return os.RemoveAll(it.FilePath)
	}

	return nil
}

// ExpireParentDir
//
//  warning: recursive process remove all files in it, undoable.
func (it SimpleFileReaderWriter) ExpireParentDir() error {
	if it.IsParentExist() {
		return os.RemoveAll(it.ParentDir)
	}

	return nil
}

func (it SimpleFileReaderWriter) RemoveOnExist() error {
	return it.Expire()
}

// RemoveDirOnExist
//
//  alias for ExpireParentDir
//  warning: recursive process remove all files in it, undoable.
func (it SimpleFileReaderWriter) RemoveDirOnExist() error {
	return it.ExpireParentDir()
}

func (it SimpleFileReaderWriter) getOnExist(toPtr interface{}) error {
	allBytes, err := it.Read()

	if err != nil {
		return err
	}

	return corejson.Deserialize.UsingBytes(
		allBytes,
		toPtr)
}

func (it SimpleFileReaderWriter) String() string {
	jsonString, err := json.Marshal(it)

	if err != nil {
		return err.Error()
	}

	return string(jsonString)
}
