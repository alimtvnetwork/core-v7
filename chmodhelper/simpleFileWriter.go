package chmodhelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"gitlab.com/evatix-go/core/constants"
)

type simpleFileWriter struct{}

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
	return errors.New(
		"dir : " + dirPath +
			", applyChmod :" + applyChmod.String() +
			", " + err.Error())
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

func (it simpleFileWriter) WriteFile(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
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

	if err == nil {
		return nil
	}

	var contentsString string

	if len(contentsBytes) > 0 {
		contentsString = string(contentsBytes)
	}

	// has err
	return errors.New(
		"file writing failed" +
			"filePath : " + writingFilePath +
			"contents : " + contentsString +
			", chmod file :" + chmodFile.String() + ", " +
			", chmod dir :" + chmodDir.String() + ", " +
			err.Error())
}

func (it simpleFileWriter) WriteFileString(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	content string,
) error {
	return it.WriteFile(
		chmodDir,
		chmodFile,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		[]byte(content))
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
