package chmodhelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/osconstsinternal"
)

type simpleFileWriter struct{}

func (it simpleFileWriter) Lock() {
	globalMutex.Lock()
}

func (it simpleFileWriter) Unlock() {
	globalMutex.Unlock()
}

func (it simpleFileWriter) CreateDirOn(
	isCreate bool,
	chmod os.FileMode,
	dirPath string,
) error {
	if !isCreate {
		return nil
	}

	return it.CreateDirOnRequired(
		chmod, dirPath)
}

func (it simpleFileWriter) CreateDirOnRequiredLock(
	applyChmod os.FileMode,
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.CreateDirOnRequired(
		applyChmod,
		dirPath)
}

func (it simpleFileWriter) CreateDirOnRequired(
	applyChmod os.FileMode,
	dirPath string,
) error {
	if IsPathExists(dirPath) {
		return nil
	}

	err := os.MkdirAll(
		dirPath,
		applyChmod)

	if err == nil {
		return nil
	}

	// has err
	return pathError(
		"dir creation failed",
		applyChmod,
		dirPath,
		err)
}

func (it simpleFileWriter) CreateDirLock(
	applyChmod os.FileMode,
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.CreateDir(
		applyChmod,
		dirPath)
}

func (it simpleFileWriter) CreateDir(
	applyChmod os.FileMode,
	dirPath string,
) error {
	err := os.MkdirAll(
		dirPath,
		applyChmod)

	if err == nil {
		return nil
	}

	// has err
	return errors.New(
		"dir : " + dirPath +
			", applyChmod :" + applyChmod.String() +
			", " + err.Error())
}

func (it simpleFileWriter) CreateDirDefaultLock(
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.CreateDirDefault(
		dirPath)
}

// CreateDirDefault
//
// Dir default chmod 0755
func (it simpleFileWriter) CreateDirDefault(
	dirPath string,
) error {
	err := os.MkdirAll(
		dirPath,
		dirDefaultChmod)

	if err == nil {
		return nil
	}

	// has err
	return errors.New(
		"dir : " + dirPath +
			", applyChmod :" + dirDefaultChmod.String() +
			", " + err.Error())
}

// WriteFileLock
//
//  Writes contents to file system.
//
// parentDirPath:
//  - is a full path to the parent dir for checking
//    if parent dir exist if not then created
//
// writingFilePath:
//  - is a full path to the actual file where to write contents
func (it simpleFileWriter) WriteFileLock(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteFile(
		chmodDir,
		chmodFile,
		isApplyChmodMust,
		isApplyChmodOnMismatch,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		contentsBytes)
}

// WriteFile
//
//  Writes contents to file system.
//
// parentDirPath:
//  - is a full path to the parent dir for checking
//    if parent dir exist if not then created
//
// writingFilePath:
//  - is a full path to the actual file where to write contents
//
// Warning:
//  - Chmod will NOT be applied to dir if already created.
//    This may harm other files.
func (it simpleFileWriter) WriteFile(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	contentsBytes []byte,
) error {
	dirErr := it.CreateDirOn(
		isCreateDirOnRequired,
		chmodDir,
		parentDirPath)

	if dirErr != nil {
		return dirErr
	}

	err := ioutil.WriteFile(
		writingFilePath,
		contentsBytes,
		chmodFile)

	if err != nil {
		return errors.New(
			"file writing failed" +
				"filePath : " + writingFilePath +
				"contents : " + corejson.BytesToString(contentsBytes) +
				", chmod file :" + chmodFile.String() + ", " +
				", chmod dir :" + chmodDir.String() + ", " +
				err.Error())
	}

	isNotApplyChmod := !isApplyChmodMust

	if isNotApplyChmod || osconstsinternal.IsWindows {
		return nil
	}

	// unix, must chmod
	if isApplyChmodOnMismatch && ChmodVerify.IsEqual(writingFilePath, chmodFile) {
		return nil
	}

	// not equal or apply anyway
	return ChmodApply.Default(chmodFile, writingFilePath)
}

// WriteFileString
//
//  Writes contents to file system.
//
// parentDirPath:
//  - is a full path to the parent dir for checking
//    if parent dir exist if not then created
//
// writingFilePath:
//  - is a full path to the actual file where to write contents
func (it simpleFileWriter) WriteFileString(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	content string,
) error {
	return it.WriteFile(
		chmodDir,
		chmodFile,
		isApplyChmodMust,
		isApplyChmodOnMismatch,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		[]byte(content))
}

func (it simpleFileWriter) WriteStringLock(
	writingFilePath string,
	content string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteString(
		writingFilePath,
		content)
}

// WriteString
//
//  Applies default chmod (for dir - 0755, for file - 0644)
func (it simpleFileWriter) WriteString(
	writingFilePath string,
	content string,
) error {
	return it.WriteFileStringCreateDirDefault(
		writingFilePath,
		content)
}

func (it simpleFileWriter) WriteFileStringCreateDir(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	content string,
) error {
	return it.WriteFileCreateDir(
		chmodDir,
		chmodFile,
		writingFilePath,
		[]byte(content))
}

func (it simpleFileWriter) WriteFileCreateDirLock(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteFileCreateDirLock(
		chmodDir,
		chmodFile,
		writingFilePath,
		contentsBytes)
}

func (it simpleFileWriter) WriteFileCreateDir(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := filepath.Clean(
		filepath.Dir(writingFilePath))

	return it.WriteFile(
		chmodDir,
		chmodFile,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes)
}

// WriteFileStringCreateDirDefault
//
//  Applies default chmod (for dir - 0755, for file - 0644)
func (it simpleFileWriter) WriteFileStringCreateDirDefault(
	writingFilePath string,
	content string,
) error {
	return it.WriteFileCreateDir(
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		[]byte(content))
}

func (it simpleFileWriter) WriteFileCreateDirDefaultLock(
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteFileCreateDirDefault(
		writingFilePath,
		contentsBytes)
}

// WriteFileCreateDirDefault
//
//  Applies default chmod (for dir - 0755, for file - 0644)
func (it simpleFileWriter) WriteFileCreateDirDefault(
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WriteFileCreateDir(
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes)
}

func (it simpleFileWriter) WriteAnyItemLock(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	parentDir,
	writingFilePath string,
	anyItem interface{},
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteAnyItem(
		chmodDir,
		chmodFile,
		parentDir,
		writingFilePath,
		anyItem)
}

// WriteAnyItem
//
//  Writes contents to file system.
//
// parentDirPath:
//  - is a full path to the parent dir for checking
//    if parent dir exist if not then created
//
// writingFilePath:
//  - is a full path to the actual file where to write contents
func (it simpleFileWriter) WriteAnyItem(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	parentDir,
	writingFilePath string,
	anyItem interface{},
) error {
	jsonBytes, err := json.Marshal(anyItem)

	if err == nil {
		return it.WriteFile(
			chmodDir,
			chmodFile,
			true,
			true,
			true,
			parentDir,
			writingFilePath,
			jsonBytes)
	}

	var typeName, anyString string
	if anyItem != nil {
		// fine if var type not detected as nil
		// we want to avoid interface nil only
		typeName = reflect.TypeOf(anyItem).String()
		anyString = fmt.Sprintf(
			constants.SprintValueFormat,
			anyItem)
	}

	// has err
	return errors.New(
		"json convert failed," +
			"filePath : " + writingFilePath +
			"AnyType : " + typeName +
			"AnyItem(String) : " + anyString +
			", chmodFile :" + chmodFile.String() + ", " +
			", chmodDir :" + chmodDir.String() + ", " +
			err.Error())
}

func (it simpleFileWriter) WriteAnyItemDefaultLock(
	writingFilePath string,
	anyItem interface{},
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteAnyItemDefault(
		writingFilePath, anyItem)
}

// WriteAnyItemDefault
//
//  Applies default chmod (for dir - 0755, for file - 0644)
func (it simpleFileWriter) WriteAnyItemDefault(
	writingFilePath string,
	anyItem interface{},
) error {
	parentDir := filepath.Dir(writingFilePath)

	return it.WriteAnyItem(
		dirDefaultChmod,
		fileDefaultChmod,
		parentDir,
		writingFilePath,
		anyItem,
	)
}
