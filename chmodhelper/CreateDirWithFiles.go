package chmodhelper

import (
	"os"
	"path"

	"gitlab.com/evatix-go/core/errcore"
)

func CreateDirWithFiles(
	isRemoveAllDirBeforeCreate bool,
	fileChmod os.FileMode,
	dirWithFile *DirWithFiles,
) error {
	const funcName = "CreateDirWithFiles"
	dir := dirWithFile.Dir
	var removeErr error

	if isRemoveAllDirBeforeCreate && IsPathExists(dir) {
		removeErr = os.RemoveAll(dir)
	}

	if removeErr != nil {
		return errcore.PathMeaningFulError(
			errcore.PathCreateFailedMessage,
			funcName,
			removeErr,
			dir)
	}

	mkDirErr := os.MkdirAll(
		dir, fileChmod)

	if mkDirErr != nil {
		return errcore.PathMeaningFulError(
			errcore.PathCreateFailedMessage,
			funcName,
			mkDirErr,
			dir)
	}

	var fileManipulateErr error

	if len(dirWithFile.Files) == 0 {
		return nil
	}

	for _, filePath := range dirWithFile.Files {
		compiledPath := path.Join(dir, filePath)
		osFile, err := os.Create(compiledPath)

		if err != nil {
			return errcore.PathMeaningFulError(
				errcore.PathCreateFailedMessage,
				funcName,
				err,
				dir)
		}

		if osFile != nil {
			fileManipulateErr = osFile.Close()
		}

		if fileManipulateErr != nil {
			return errcore.PathMeaningFulError(
				errcore.FileCloseFailedMessage,
				funcName,
				fileManipulateErr,
				compiledPath)
		}

		chmodErr := os.Chmod(
			compiledPath,
			fileChmod)

		if chmodErr != nil {
			return errcore.PathMeaningFulError(
				errcore.PathChmodApplyMessage,
				funcName,
				chmodErr,
				compiledPath)
		}
	}

	return nil
}
