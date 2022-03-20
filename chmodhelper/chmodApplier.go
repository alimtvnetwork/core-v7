package chmodhelper

import "os"

type chmodApplier struct {
	rw *SimpleFileReaderWriter
}

func (it chmodApplier) OnParent() error {
	return it.OnDir(it.rw.ParentDir)
}

func (it chmodApplier) OnDir(dir string) error {
	return os.Chmod(dir, it.rw.ChmodDir)
}

func (it chmodApplier) OnFile() error {
	return os.Chmod(it.rw.FilePath, it.rw.ChmodFile)
}

func (it chmodApplier) OnDiffFile(filePath string) error {
	return os.Chmod(filePath, it.rw.ChmodFile)
}

func (it chmodApplier) ApplyAll() error {
	err := it.OnParent()

	if err != nil {
		return err
	}

	return it.OnFile()
}

func (it chmodApplier) DirRecursive(
	isSkipOnInvalid bool,
	dir string,
) error {
	rwx := NewUsingFileMode(it.rw.ChmodDir)

	return rwx.ApplyChmod(isSkipOnInvalid, dir)
}

func (it chmodApplier) OnParentRecursive() error {
	return it.DirRecursive(
		false,
		it.rw.ParentDir)
}
