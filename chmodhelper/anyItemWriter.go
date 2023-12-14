package chmodhelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"gitlab.com/auk-go/core/constants"
)

type anyItemWriter struct{}

func (it anyItemWriter) ChmodLock(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	parentDir,
	writingFilePath string,
	anyItem interface{},
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Chmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		parentDir,
		writingFilePath,
		anyItem,
	)
}

// Chmod
//
//	Writes contents to file system by serializing using JSON.
//
// parentDirPath:
//   - is a full path to the parent dir for checking
//     if parent dir exist if not then created
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
func (it anyItemWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	parentDir,
	writingFilePath string,
	anyItem interface{},
) error {
	jsonBytes, err := json.Marshal(anyItem)

	if err == nil {
		return fileWriter{}.All(
			chmodDir,
			chmodFile,
			isRemoveBeforeWrite,
			true,
			true,
			true,
			parentDir,
			writingFilePath,
			jsonBytes,
		)
	}

	var typeName, anyString string
	if anyItem != nil {
		// fine if var type not detected as nil
		// we want to avoid interface nil only
		typeName = reflect.TypeOf(anyItem).String()
		anyString = fmt.Sprintf(
			constants.SprintValueFormat,
			anyItem,
		)
	}

	// has err
	return errors.New(
		"json convert failed," +
			", filePath : " + writingFilePath +
			", AnyType : " + typeName +
			", AnyItem(String) : " + anyString +
			", chmodFile :" + chmodFile.String() + ", " +
			", chmodDir :" + chmodDir.String() + ", " +
			err.Error(),
	)
}

// DefaultLock
//
//	Writes contents to file system by serializing using JSON.
//	Applies default chmod (for dir - 0755, for file - 0644)
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
func (it anyItemWriter) DefaultLock(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	anyItem interface{},
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Default(
		isRemoveBeforeWrite,
		writingFilePath,
		anyItem,
	)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it anyItemWriter) Default(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	anyItem interface{},
) error {
	parentDir := filepath.Dir(writingFilePath)

	return it.Chmod(
		isRemoveBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		parentDir,
		writingFilePath,
		anyItem,
	)
}
