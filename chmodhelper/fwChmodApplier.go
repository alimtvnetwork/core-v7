package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/errcore"
)

type fwChmodApplier struct {
	rw *SimpleFileReaderWriter
}

func (it fwChmodApplier) OnParent() error {
	return it.OnDir(it.rw.ParentDir)
}

func (it fwChmodApplier) OnDir(dir string) error {
	return it.Apply(
		it.rw.ChmodDir,
		dir)
}

func (it fwChmodApplier) OnFile() error {
	return it.Apply(
		it.rw.ChmodFile,
		it.rw.FilePath)
}

func (it fwChmodApplier) Apply(
	fileMode os.FileMode,
	location string,
) error {
	err := os.Chmod(
		location,
		fileMode)

	if err == nil {
		return nil
	}

	// has error
	return pathError(
		"applying chmod failed",
		fileMode,
		location,
		err)
}

// OnDiffFile
//
//  apply chmod on file if file doesn't have the save chmod
func (it fwChmodApplier) OnDiffFile(
	isSkipOnInvalidFile bool,
	filePath string,
) error {
	if ChmodVerify.IsEqual(filePath, it.rw.ChmodFile) {
		return nil
	}

	if isSkipOnInvalidFile && IsPathInvalid(filePath) {
		return nil
	}

	return it.Apply(it.rw.ChmodFile, filePath)
}

// OnDiffDir
//
//  apply chmod on file if file doesn't have the save chmod
func (it fwChmodApplier) OnDiffDir(
	isSkipOnInvalidDir bool,
	dirPath string,
) error {
	if ChmodVerify.IsEqual(dirPath, it.rw.ChmodDir) {
		return nil
	}

	if isSkipOnInvalidDir && IsPathInvalid(dirPath) {
		return nil
	}

	return it.Apply(it.rw.ChmodDir, dirPath)
}

// OnAll
//
//  both file, parent dir
func (it fwChmodApplier) OnAll() error {
	err := it.OnParent()

	if err != nil {
		return err
	}

	return it.OnFile()
}

func (it fwChmodApplier) DirRecursive(
	isSkipOnInvalid bool,
	dir string,
) error {
	rwx := New.RwxWrapper.UsingFileMode(it.rw.ChmodDir)

	return rwx.ApplyRecursive(isSkipOnInvalid, dir)
}

func (it fwChmodApplier) OnParentRecursive() error {
	return it.DirRecursive(
		false,
		it.rw.ParentDir)
}

func (it fwChmodApplier) OnMismatch(
	isFile,
	isParentDir bool,
) error {
	if !isFile && !isParentDir {
		return nil
	}

	verifier := it.rw.ChmodVerifier()
	var fileErr, dirErr error

	if isFile && verifier.HasMismatchFile() {
		fileErr = it.OnFile()
	}

	if isParentDir && verifier.HasMismatchParentDir() {
		dirErr = it.OnParent()
	}

	return errcore.MergeErrors(fileErr, dirErr)
}
