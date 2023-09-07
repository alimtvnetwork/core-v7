package chmodhelper

type fwChmodVerifier struct {
	rw *SimpleFileReaderWriter
}

func (it fwChmodVerifier) HasMismatchFile() bool {
	return ChmodVerify.IsMismatch(
		it.rw.FilePath,
		it.rw.ChmodFile)
}

func (it fwChmodVerifier) IsEqualFile() bool {
	return ChmodVerify.IsEqual(
		it.rw.FilePath,
		it.rw.ChmodFile)
}

func (it fwChmodVerifier) HasMismatchParentDir() bool {
	return ChmodVerify.IsMismatch(
		it.rw.ParentDir,
		it.rw.ChmodDir)
}

func (it fwChmodVerifier) IsEqualParentDir() bool {
	return ChmodVerify.IsEqual(
		it.rw.ParentDir,
		it.rw.ChmodDir)
}

func (it fwChmodVerifier) MismatchErrorParentDir() error {
	return ChmodVerify.MismatchError(
		it.rw.ParentDir,
		it.rw.ChmodDir)
}

func (it fwChmodVerifier) MismatchErrorFile() error {
	return ChmodVerify.MismatchError(
		it.rw.FilePath,
		it.rw.ChmodFile)
}
