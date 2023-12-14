package chmodhelper

import "os"

type fileStringWriter struct{}

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
func (it fileStringWriter) All(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	content string,
) error {
	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		isApplyChmodMust,
		isApplyChmodOnMismatch,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		[]byte(content),
	)
}

func (it fileStringWriter) DefaultLock(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	content string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Default(
		isRemoveBeforeWrite,
		writingFilePath,
		content,
	)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileStringWriter) Default(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	content string,
) error {
	return fileWriter{}.Bytes.Default(
		isRemoveBeforeWrite,
		writingFilePath,
		[]byte(content),
	)
}

func (it fileStringWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	content string,
) error {
	return fileWriter{}.Bytes.WithDirChmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		writingFilePath,
		[]byte(content),
	)
}

func (it fileStringWriter) ChmodLock(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	content string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return fileWriter{}.Bytes.WithDirChmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		writingFilePath,
		[]byte(content),
	)
}
