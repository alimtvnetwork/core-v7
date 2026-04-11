package chmodhelper

import (
	"os"
	"path"

	"github.com/alimtvnetwork/core/errcore"
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
		dir, dirDefaultChmod,
	)

	if mkDirErr != nil {
		return errcore.PathMeaningfulError(
			errcore.PathCreateFailedType,
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
				chmodErr,
				compiledPath,
			)
		}
	}

	return nil
}
