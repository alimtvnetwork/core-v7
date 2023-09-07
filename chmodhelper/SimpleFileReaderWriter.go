package chmodhelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
)

type SimpleFileReaderWriter struct {
	ChmodDir, ChmodFile    os.FileMode
	ParentDir              string // full path to the parent dir
	FilePath               string // full path to the actual file to write to or read from.
	IsMustChmodApplyOnFile bool
	IsApplyChmodOnMismatch bool
}

func (it SimpleFileReaderWriter) InitializeDefault(
	isMustApplyChmod bool,
) *SimpleFileReaderWriter {
	filePath := path.Clean(it.FilePath)
	parentDir := it.ParentDir

	if parentDir == "" {
		parentDir = filepath.Dir(filePath)
	}

	return &SimpleFileReaderWriter{
		ChmodDir:               it.ChmodDir,
		ChmodFile:              it.ChmodFile,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsMustChmodApplyOnFile: isMustApplyChmod,
		IsApplyChmodOnMismatch: true,
	}
}

func (it SimpleFileReaderWriter) InitializeDefaultApplyChmod() *SimpleFileReaderWriter {
	return it.InitializeDefault(true)
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

func (it *SimpleFileReaderWriter) ChmodApplier() fwChmodApplier {
	return fwChmodApplier{
		rw: it,
	}
}

func (it *SimpleFileReaderWriter) ChmodVerifier() fwChmodVerifier {
	return fwChmodVerifier{
		rw: it,
	}
}

func (it SimpleFileReaderWriter) Write(allBytes []byte) error {
	err := SimpleFileWriter.WriteFile(
		it.ChmodDir,
		it.ChmodFile,
		it.IsMustChmodApplyOnFile,
		it.IsApplyChmodOnMismatch,
		true,
		it.ParentDir,
		it.FilePath,
		allBytes)

	if err == nil {
		return nil
	}

	return it.errorWrap(err)
}

func (it SimpleFileReaderWriter) InitializeDefaultNew() (newRw *SimpleFileReaderWriter) {
	return New.SimpleFileReaderWriter.Default(
		it.FilePath)
}

func (it SimpleFileReaderWriter) NewPath(
	newLocation string,
) (newRw *SimpleFileReaderWriter) {
	return New.SimpleFileReaderWriter.Path(
		it.ChmodDir,
		it.ChmodFile,
		newLocation)
}

func (it SimpleFileReaderWriter) NewPathJoin(
	newLocationsFromParentDir ...string,
) (newRw *SimpleFileReaderWriter) {
	joined := strings.Join(
		newLocationsFromParentDir,
		constants.ForwardSlash)
	newLocation := filepath.Join(
		it.ParentDir,
		joined)

	return New.SimpleFileReaderWriter.Path(
		it.ChmodDir,
		it.ChmodFile,
		newLocation)
}

func (it SimpleFileReaderWriter) WriteString(content string) error {
	err := SimpleFileWriter.WriteFileString(
		it.ChmodDir,
		it.ChmodFile,
		it.IsMustChmodApplyOnFile,
		it.IsApplyChmodOnMismatch,
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
		"err: %s simple-reader-writer: %s",
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

func (it SimpleFileReaderWriter) WriteAnyLock(
	anyItem interface{},
) error {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.WriteAny(anyItem)
}

func (it SimpleFileReaderWriter) ReadMust() []byte {
	allBytes, err := it.Read()

	if err != nil {
		panic(err)
	}

	return allBytes
}

func (it SimpleFileReaderWriter) ReadString() (content string, err error) {
	allBytes, err := it.Read()

	if len(allBytes) > 0 {
		return string(allBytes), err
	}

	return "", err
}

func (it SimpleFileReaderWriter) ReadStringMust() (content string) {
	content, err := it.ReadString()

	if err != nil {
		panic(err)
	}

	return content
}

func (it SimpleFileReaderWriter) ReadStringLock() (content string, err error) {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.ReadString()
}

func (it SimpleFileReaderWriter) Read() ([]byte, error) {
	allBytes, err := ioutil.ReadFile(it.FilePath)

	if err == nil {
		return allBytes, err
	}

	message := fmt.Sprintf(
		"cannot read file: %q, err: %s simple-reader-writer: %s",
		it.FilePath,
		err.Error(),
		it.String(),
	)

	return allBytes, errors.New(message)
}

func (it SimpleFileReaderWriter) ReadLock() ([]byte, error) {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.Read()
}

func (it SimpleFileReaderWriter) ReadOnExist() ([]byte, error) {
	if it.IsExist() {
		return it.Read()
	}

	return nil, nil
}

func (it SimpleFileReaderWriter) ReadStringOnExist() (content string, err error) {
	if it.IsExist() {
		return it.ReadString()
	}

	return "", nil
}

func (it SimpleFileReaderWriter) ReadStringOnExistLock() (content string, err error) {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.ReadStringOnExist()
}

func (it SimpleFileReaderWriter) ReadOnExistLock() ([]byte, error) {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.ReadOnExist()
}

func (it SimpleFileReaderWriter) Get(toPtr interface{}) error {
	if it.IsExist() {
		return it.getOnExist(toPtr)
	}

	return it.errorWrap(errors.New("cannot read cache, save first, file not exist: " + it.FilePath))
}

func (it SimpleFileReaderWriter) GetLock(toPtr interface{}) error {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.Get(toPtr)
}

func (it SimpleFileReaderWriter) ReadWrite(
	readToPtr interface{},
	onInvalidGenerateFunc func() (interface{}, error),
) error {
	return it.GetSet(
		readToPtr,
		onInvalidGenerateFunc)
}

func (it SimpleFileReaderWriter) ReadWriteLock(
	readToPtr interface{},
	onInvalidGenerateFunc func() (interface{}, error),
) error {
	return it.GetSetLock(
		readToPtr,
		onInvalidGenerateFunc)
}

func (it SimpleFileReaderWriter) GetSetLock(
	toPtr interface{},
	onInvalidGenerateFunc func() (interface{}, error),
) error {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.GetSet(
		toPtr, onInvalidGenerateFunc)
}

func (it SimpleFileReaderWriter) GetSet(
	toPtr interface{},
	onInvalidGenerateFunc func() (interface{}, error),
) error {
	readErr := it.Get(toPtr)
	if readErr != nil {
		// remove file
		// fine swallow error
		it.Expire()
	}

	newAnyItem, err := onInvalidGenerateFunc()

	if err == nil {
		// if things are all right
		reflect.ValueOf(toPtr).Elem().Set(reflect.ValueOf(newAnyItem).Elem())

		// save
		return it.Set(newAnyItem)
	}

	return it.errorWrap(errors.New("read cache failed + cannot generate: " + err.Error()))
}

func (it SimpleFileReaderWriter) CacheGetSet(
	toPtr interface{},
	onInvalidGenerateFunc func() (interface{}, error),
) error {
	return it.GetSet(
		toPtr,
		onInvalidGenerateFunc)
}

func (it SimpleFileReaderWriter) CacheGetSetLock(
	toPtr interface{},
	onInvalidGenerateFunc func() (interface{}, error),
) error {
	return it.GetSetLock(
		toPtr,
		onInvalidGenerateFunc)
}

// Deserialize
//
//  alias for Get
func (it SimpleFileReaderWriter) Deserialize(
	toPtr interface{},
) error {
	return it.Get(toPtr)
}

// DeserializeLock
//
//  alias for Get
func (it SimpleFileReaderWriter) DeserializeLock(
	toPtr interface{},
) error {
	return it.GetLock(toPtr)
}

// Serialize
//
//  alias for ReadOnExist
func (it SimpleFileReaderWriter) Serialize() ([]byte, error) {
	return it.ReadOnExist()
}

// SerializeLock
//
//  alias for ReadOnExist
func (it SimpleFileReaderWriter) SerializeLock() ([]byte, error) {
	return it.ReadOnExistLock()
}

// Set
//
//  alias for WriteAny
func (it SimpleFileReaderWriter) Set(toPtr interface{}) error {
	return it.WriteAny(toPtr)
}

func (it SimpleFileReaderWriter) SetLock(toPtr interface{}) error {
	return it.WriteAnyLock(toPtr)
}

// Expire
//
//  Removes file on exist only
//  alias for RemoveOnExist
func (it SimpleFileReaderWriter) Expire() error {
	if it.IsExist() {
		return os.RemoveAll(it.FilePath)
	}

	return nil
}

// ExpireLock
//
//  Removes file on exist only
//  alias for RemoveOnExist
func (it SimpleFileReaderWriter) ExpireLock() error {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.Expire()
}

// OsFile
//
//  Os open files must be closed
func (it SimpleFileReaderWriter) OsFile() (*os.File, error) {
	return os.Open(it.FilePath)
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

// ExpireParentDirLock
//
//  warning: recursive process remove all files in it, undoable.
func (it SimpleFileReaderWriter) ExpireParentDirLock() error {
	SimpleFileWriter.Lock()
	defer SimpleFileWriter.Unlock()

	return it.ExpireParentDir()
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

func (it SimpleFileReaderWriter) Clone() SimpleFileReaderWriter {
	return SimpleFileReaderWriter{
		ChmodDir:  it.ChmodDir,
		ChmodFile: it.ChmodFile,
		ParentDir: it.ParentDir,
		FilePath:  it.FilePath,
	}
}

func (it *SimpleFileReaderWriter) ClonePtr() *SimpleFileReaderWriter {
	if it == nil {
		return nil
	}

	return &SimpleFileReaderWriter{
		ChmodDir:  it.ChmodDir,
		ChmodFile: it.ChmodFile,
		ParentDir: it.ParentDir,
		FilePath:  it.FilePath,
	}
}

func (it SimpleFileReaderWriter) MarshalJSON() ([]byte, error) {
	model := simpleFileReaderWriterModel{
		ChmodDir:  New.RwxWrapper.UsingFileMode(it.ChmodDir),
		ChmodFile: New.RwxWrapper.UsingFileMode(it.ChmodFile),
		ParentDir: it.ParentDir,
		FilePath:  it.FilePath,
	}

	return corejson.Serialize.Raw(model)
}

func (it *SimpleFileReaderWriter) UnmarshalJSON(jsonBytes []byte) error {
	var model simpleFileReaderWriterModel
	err := corejson.Deserialize.UsingBytes(
		jsonBytes, &model)

	if err == nil {
		// success
		it.ChmodDir = model.ChmodDir.ToFileMode()
		it.ChmodFile = model.ChmodFile.ToFileMode()
		it.ParentDir = model.ParentDir
		it.FilePath = model.FilePath
	}

	return err
}

func (it SimpleFileReaderWriter) Json() corejson.Result {
	return corejson.New(it)
}

func (it SimpleFileReaderWriter) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *SimpleFileReaderWriter) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	return jsonResult.Deserialize(it)
}

func (it SimpleFileReaderWriter) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
