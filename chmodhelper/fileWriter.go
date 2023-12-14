package chmodhelper

import (
	"errors"
	"os"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/osconstsinternal"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

type fileWriter struct {
	Bytes  fileBytesWriter
	String fileStringWriter
	Any    anyItemWriter // writes any item using JSON
}

// AllLock
//
//	Writes contents to file system.
//
// parentDirPath:
//   - is a full path to the parent dir for checking
//     if parent dir exist if not then created
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
func (it fileWriter) AllLock(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isRemoveBeforeWrite,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		isApplyChmodMust,
		isApplyChmodOnMismatch,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		contentsBytes,
	)
}

// All
//
//	Writes contents to file system.
//
// parentDirPath:
//   - is a full path to the parent dir for checking
//     if parent dir exist if not then created
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
//
// Warning:
//   - Chmod will NOT be applied to dir if already created.
//     This may harm other files.
func (it fileWriter) All(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isRemoveBeforeWrite,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	contentsBytes []byte,
) error {
	dirErr := dirCreator{}.If(
		isCreateDirOnRequired,
		chmodDir,
		parentDirPath,
	)

	if dirErr != nil {
		return dirErr
	}

	cleanUpErr := it.RemoveIf(
		isRemoveBeforeWrite,
		writingFilePath,
	)

	if cleanUpErr != nil {
		return cleanUpErr
	}

	err := os.WriteFile(
		writingFilePath,
		contentsBytes,
		chmodFile,
	)

	if err != nil {
		return errors.New(
			"writing failed " +
				"filePath : " + writingFilePath +
				", contents : " + corejson.BytesToString(contentsBytes) +
				", chmod file :" + chmodFile.String() + ", " +
				", chmod dir :" + chmodDir.String() + ", " +
				err.Error(),
		)
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

func (it fileWriter) Remove(removePath string) error {
	err := os.RemoveAll(removePath)

	if err != nil {
		return errors.New(
			"clean up or remove failed " +
				"filePath : " + removePath +
				", " +
				err.Error(),
		)
	}

	return nil
}

func (it fileWriter) RemoveIf(
	isRemove bool,
	removePath string,
) error {
	if !isRemove {
		return nil
	}

	if !pathinternal.IsPathExists(removePath) {
		return nil
	}

	return it.Remove(removePath)
}

func (it fileWriter) ParentDir(filePath string) string {
	return pathinternal.ParentDir(filePath)
}

func (it fileWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	filePath string,
	contentsBytes []byte,
) error {
	parentDir := it.ParentDir(filePath)

	return it.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		filePath,
		contentsBytes,
	)
}

func (it fileWriter) ChmodFile(
	isRemoveBeforeWrite bool,
	chmodFile os.FileMode,
	filePath string,
	contentsBytes []byte,
) error {
	parentDir := it.ParentDir(filePath)

	return it.All(
		dirDefaultChmod,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		filePath,
		contentsBytes,
	)
}
