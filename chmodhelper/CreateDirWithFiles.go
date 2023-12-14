package chmodhelper

import (
	"os"
	"path"

	"gitlab.com/auk-go/core/errcore"
)

func CreateDirWithFiles(
	isRemoveAllDirBeforeCreate bool,
	fileChmod os.FileMode,
	dirWithFile *DirWithFiles,
) error {
	const funcName = "CreateDirWithFiles"
	dir := dirWithFile.Dir

	removeDirErr := removeDirIf(
		isRemoveAllDirBeforeCreate,
		dir,
		funcName,
	)

	if removeDirErr != nil {
		return removeDirErr
	}

	mkDirErr := os.MkdirAll(
		dir, fileChmod,
	)

	if mkDirErr != nil {
		return errcore.PathMeaningfulError(
			errcore.PathCreateFailedType,
			funcName,
			mkDirErr,
			dir,
		)
	}

	var fileManipulateErr error

	if len(dirWithFile.Files) == 0 {
		return nil
	}

	for _, filePath := range dirWithFile.Files {
		compiledPath := path.Join(dir, filePath)
		osFile, err := os.Create(compiledPath)

		if err != nil {
			return errcore.PathMeaningfulError(
				errcore.PathCreateFailedType,
				funcName,
				err,
				dir,
			)
		}

		if osFile != nil {
			fileManipulateErr = osFile.Close()
		}

		if fileManipulateErr != nil {
			return errcore.PathMeaningfulError(
				errcore.FileCloseFailedType,
				funcName,
				fileManipulateErr,
				compiledPath,
			)
		}

		chmodErr := os.Chmod(
			compiledPath,
			fileChmod,
		)

		if chmodErr != nil {
			return errcore.PathMeaningfulError(
				errcore.PathChmodApplyType,
				funcName,
				chmodErr,
				compiledPath,
			)
		}
	}

	return nil
}
