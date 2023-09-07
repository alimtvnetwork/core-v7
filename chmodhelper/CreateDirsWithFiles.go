package chmodhelper

import "os"

func CreateDirsWithFiles(
	isRemoveAllDirBeforeCreate bool,
	fileChmod os.FileMode,
	dirsWithFiles ...DirWithFiles,
) error {
	if dirsWithFiles == nil || len(dirsWithFiles) == 0 {
		return nil
	}

	for _, dirWithFile := range dirsWithFiles {
		err := CreateDirWithFiles(
			isRemoveAllDirBeforeCreate,
			fileChmod,
			&dirWithFile)

		if err != nil {
			return err
		}
	}

	return nil
}
