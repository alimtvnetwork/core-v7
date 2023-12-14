package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/internal/pathinternal"
)

type fileBytesWriter struct{}

// WithDirChmodLock
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmodLock(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDirChmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		writingFilePath,
		contentsBytes,
	)
}

// WithDirChmod
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := pathinternal.ParentDir(writingFilePath)

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes,
	)
}

// Chmod
//
// Create dir safely if required.
func (it fileBytesWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := pathinternal.ParentDir(writingFilePath)

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes,
	)
}

func (it fileBytesWriter) WithDirLock(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDir(
		isRemoveBeforeWrite,
		writingFilePath,
		contentsBytes,
	)
}

// WithDir
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) WithDir(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		isRemoveBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes,
	)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) Default(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		isRemoveBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes,
	)
}
