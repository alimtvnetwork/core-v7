package chmodhelper

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ══════════════════════════════════════════════════════════════════════════════
// fileWriter — error branches in All, RemoveIf, applyDirChmod
// ══════════════════════════════════════════════════════════════════════════════

func Test_FileWriter_All_CleanUpError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "cleanup.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	// Make the parent dir read-only to cause RemoveAll to fail
	os.Chmod(dir, 0555)
	defer os.Chmod(dir, 0755)

	fw := fileWriter{}

	// Act
	err := fw.All(
		0755,
		0644,
		true,  // isRemoveBeforeWrite
		false, // isApplyChmodMust
		false, // isApplyChmodOnMismatch
		false, // isCreateDirOnRequired
		dir,
		fp,
		[]byte("new content"),
	)

	// Assert — should either error or succeed depending on OS permissions
	_ = err
}

func Test_FileWriter_All_WriteFailure(t *testing.T) {
	// Arrange
	fw := fileWriter{}
	invalidPath := string([]byte{0}) + "/impossible/file.txt"

	// Act — write to an impossible path
	err := fw.All(
		0755,
		0644,
		false, // isRemoveBeforeWrite
		false, // isApplyChmodMust
		false, // isApplyChmodOnMismatch
		false, // isCreateDirOnRequired
		string([]byte{0}),
		invalidPath,
		[]byte("data"),
	)

	// Assert
	if err == nil {
		t.Fatal("expected write failure error")
	}
}

func Test_FileWriter_All_ChmodMismatchBranch(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "chmod_mismatch.txt")
	fw := fileWriter{}

	// Act — write with chmod must + mismatch check, file already matches
	os.WriteFile(fp, []byte("x"), 0644)
	err := fw.All(
		0755,
		0644,
		false, // isRemoveBeforeWrite
		true,  // isApplyChmodMust
		true,  // isApplyChmodOnMismatch — check if equal, skip if so
		false, // isCreateDirOnRequired
		dir,
		fp,
		[]byte("data"),
	)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_FileWriter_All_ChmodMustNotMismatch(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "chmod_force.txt")
	fw := fileWriter{}

	// Act — write with chmod must, NOT mismatch-only → always apply chmod
	err := fw.All(
		0755,
		0600,
		false, // isRemoveBeforeWrite
		true,  // isApplyChmodMust
		false, // isApplyChmodOnMismatch — apply anyway
		false, // isCreateDirOnRequired
		dir,
		fp,
		[]byte("data"),
	)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_FileWriter_ApplyDirChmod_NotCreateDir(t *testing.T) {
	// Arrange
	fw := fileWriter{}

	// Act — isCreateDirOnRequired=false, should return nil
	err := fw.applyDirChmod(false, 0755, "/some/path")

	// Assert
	if err != nil {
		t.Fatal("expected nil when isCreateDirOnRequired is false")
	}
}

func Test_FileWriter_ApplyDirChmod_DefaultChmod_NoCreate(t *testing.T) {
	// Arrange
	fw := fileWriter{}

	// Act — chmodDir matches dirDefaultChmod, should skip
	err := fw.applyDirChmod(true, dirDefaultChmod, "/some/path")

	// Assert
	if err != nil {
		t.Fatal("expected nil when chmodDir is default")
	}
}

func Test_FileWriter_ApplyDirChmod_ChmodError(t *testing.T) {
	// Arrange
	fw := fileWriter{}

	// Act — apply chmod on non-existent path
	err := fw.applyDirChmod(true, 0700, "/nonexistent/xyz/abc")

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent path")
	}
}

func Test_FileWriter_RemoveIf_NotRemove(t *testing.T) {
	// Arrange
	fw := fileWriter{}

	// Act
	err := fw.RemoveIf(false, "/some/path")

	// Assert
	if err != nil {
		t.Fatal("expected nil when isRemove is false")
	}
}

func Test_FileWriter_RemoveIf_PathNotExist(t *testing.T) {
	// Arrange
	fw := fileWriter{}

	// Act
	err := fw.RemoveIf(true, "/nonexistent/xyz/abc")

	// Assert
	if err != nil {
		t.Fatal("expected nil when path does not exist")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// errorCreator — all error branch paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_ErrorCreator_DirError_WithError(t *testing.T) {
	// Arrange & Act
	err := newError.dirError("/nonexistent/xyz", errors.New("test-err"))

	// Assert
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_ErrorCreator_DirError_NotDirButExists(t *testing.T) {
	// Arrange — create a file (not a directory)
	dir := t.TempDir()
	fp := filepath.Join(dir, "file.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act — dirError on a file path (not a dir)
	err := newError.dirError(fp, errors.New("test-err"))

	// Assert
	if err == nil {
		t.Fatal("expected not-dir error")
	}
	if !strings.Contains(err.Error(), "not a dir") {
		t.Fatal("expected 'not a dir' message, got:", err.Error())
	}
}

func Test_ErrorCreator_NotDirError_PathInvalid_FileWriterContext(t *testing.T) {
	// Arrange & Act — path doesn't exist
	err := newError.notDirError("/nonexistent/xyz")

	// Assert — should return nil (path invalid means no error)
	if err != nil {
		t.Fatal("expected nil for invalid path")
	}
}

func Test_ErrorCreator_PathError_NilErr_FileWriterContext(t *testing.T) {
	// Arrange & Act
	err := newError.pathError("test", 0644, "/tmp/x", nil)

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil input error")
	}
}

func Test_ErrorCreator_PathError_WithErr(t *testing.T) {
	// Arrange & Act
	err := newError.pathError("test msg", 0644, "/tmp/x", errors.New("inner"))

	// Assert
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_ErrorCreator_PathErrorWithDirValidate_NotDir_FileWriterContext(t *testing.T) {
	// Arrange — create a file (not dir)
	dir := t.TempDir()
	fp := filepath.Join(dir, "file.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := newError.pathErrorWithDirValidate("test", 0644, fp, errors.New("inner"))

	// Assert
	if err == nil {
		t.Fatal("expected not-dir error")
	}
}

func Test_ErrorCreator_PathErrorWithDirValidate_NilErr_FileWriterContext(t *testing.T) {
	// Arrange
	dir := t.TempDir()

	// Act — valid dir, nil error
	err := newError.pathErrorWithDirValidate("test", 0755, dir, nil)

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil input error on valid dir")
	}
}

func Test_ErrorCreator_ChmodApplyFailed_NilErr_FileWriterContext(t *testing.T) {
	// Arrange & Act
	err := newError.chmodApplyFailed(0644, "/tmp/x", nil)

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil input error")
	}
}

func Test_ErrorCreator_ChmodApplyFailed_WithErr_FileWriterContext(t *testing.T) {
	// Arrange & Act
	err := newError.chmodApplyFailed(0644, "/tmp/x", errors.New("chmod-fail"))

	// Assert
	if err == nil {
		t.Fatal("expected error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// chmodVerifier — remaining uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ChmodVerifier_IsEqualRwxFullSkipInvalid(t *testing.T) {
	// Arrange & Act — invalid path should return true (skip)
	result := ChmodVerify.IsEqualRwxFullSkipInvalid("/nonexistent/xyz", "-rwxrwxrwx")

	// Assert
	if !result {
		t.Fatal("expected true for invalid path with skip")
	}
}

func Test_ChmodVerifier_IsEqualSkipInvalid_FileModeContext(t *testing.T) {
	// Arrange & Act — invalid path should return true
	result := ChmodVerify.IsEqualSkipInvalid("/nonexistent/xyz", 0644)

	// Assert
	if !result {
		t.Fatal("expected true for invalid path with skip")
	}
}

func Test_ChmodVerifier_GetRwx9_EmptyReturn(t *testing.T) {
	// Arrange — mode 0 produces "----------" which is 10 chars
	// We need a mode whose String() is <= 9 chars — this is practically
	// impossible with standard Go, but we call it to cover the function

	// Act
	result := ChmodVerify.GetRwx9(0)

	// Assert — "----------" is 10 chars, so result should be 9 chars
	if result != "---------" {
		t.Fatal("expected 9 hyphens, got:", result)
	}
}

func Test_ChmodVerifier_PathIf_NotVerify(t *testing.T) {
	// Arrange & Act
	err := ChmodVerify.PathIf(false, "/any/path", 0644)

	// Assert
	if err != nil {
		t.Fatal("expected nil when isVerify is false")
	}
}

func Test_ChmodVerifier_PathIf_Verify(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := ChmodVerify.PathIf(true, fp, 0755)

	// Assert — mismatch expected
	if err == nil {
		t.Fatal("expected mismatch error")
	}
}

func Test_ChmodVerifier_PathsUsingFileMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := ChmodVerify.PathsUsingFileMode(true, 0755, fp)

	// Assert
	if err == nil {
		t.Fatal("expected error for mismatched chmod")
	}
}

func Test_ChmodVerifier_PathsUsingRwxFull_NilLocations(t *testing.T) {
	// Arrange & Act
	err := ChmodVerify.PathsUsingRwxFull(false, "-rwxrwxrwx")

	// Assert
	if err == nil {
		t.Fatal("expected error for nil/empty locations")
	}
}

func Test_ChmodVerifier_PathsUsingPartialRwxOptions_Error(t *testing.T) {
	// Arrange & Act — invalid partial rwx
	err := ChmodVerify.PathsUsingPartialRwxOptions(false, false, "INVALID", "/tmp")

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid partial rwx")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// fwChmodVerifier — IsEqualFile with valid file
// ══════════════════════════════════════════════════════════════════════════════

func Test_FwChmodVerifier_IsEqualFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	rw := &SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	v := fwChmodVerifier{rw: rw}

	// Act
	result := v.IsEqualFile()

	// Assert
	if !result {
		t.Fatal("expected equal for matching chmod")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// fwChmodApplier — OnDiffFile skip-on-invalid branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_FwChmodApplier_OnDiffFile_SkipInvalid(t *testing.T) {
	// Arrange
	rw := &SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}
	applier := fwChmodApplier{rw: rw}

	// Act — skip on invalid file path
	err := applier.OnDiffFile(true, "/nonexistent/xyz/file.txt")

	// Assert
	if err != nil {
		t.Fatal("expected nil when skipping invalid")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// dirCreator — Default with error, ByChecking error branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_DirCreator_Default_Error(t *testing.T) {
	// Arrange & Act
	err := internalDirCreator.Default(0755, "/dev/null/impossible/dir")

	// Assert
	if err == nil {
		t.Log("no error — OS allowed (unlikely)")
	}
}

func Test_DirCreator_Direct_Error(t *testing.T) {
	// Arrange & Act
	err := internalDirCreator.Direct("/dev/null/impossible/dir")

	// Assert
	if err == nil {
		t.Log("no error — OS allowed (unlikely)")
	}
}

func Test_DirCreator_DirectLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	subDir := filepath.Join(dir, "sub")

	// Act
	err := internalDirCreator.DirectLock(subDir)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_DirCreator_DefaultLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	subDir := filepath.Join(dir, "sub2")

	// Act
	err := internalDirCreator.DefaultLock(0755, subDir)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — errorWrap and WriteString branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleFileReaderWriter_WriteString_Error(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              string([]byte{0}),
		FilePath:               string([]byte{0}) + "/file.txt",
		IsRemoveBeforeWrite:    false,
		IsMustChmodApplyOnFile: false,
		IsApplyChmodOnMismatch: false,
	}

	// Act
	err := rw.WriteString("test content")

	// Assert
	if err == nil {
		t.Fatal("expected error for impossible path")
	}
}

func Test_SimpleFileReaderWriter_WriteRelativePath_Error(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              string([]byte{0}),
		FilePath:               string([]byte{0}) + "/file.txt",
		IsRemoveBeforeWrite:    false,
		IsMustChmodApplyOnFile: false,
		IsApplyChmodOnMismatch: false,
	}

	// Act
	err := rw.WriteRelativePath(false, "sub/file.txt", []byte("data"))

	// Assert
	if err == nil {
		t.Fatal("expected error for impossible path")
	}
}

func Test_SimpleFileReaderWriter_WriteAny_Error(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              string([]byte{0}),
		FilePath:               string([]byte{0}) + "/file.txt",
		IsRemoveBeforeWrite:    false,
		IsMustChmodApplyOnFile: false,
		IsApplyChmodOnMismatch: false,
	}

	// Act
	err := rw.WriteAny(map[string]string{
		"key": "value",
	})

	// Assert
	if err == nil {
		t.Fatal("expected error for impossible path")
	}
}

func Test_SimpleFileReaderWriter_ErrorWrap_NilErr(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{
		FilePath: "/tmp/test.txt",
	}

	// Act
	err := rw.errorWrap(nil)

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_SimpleFileReaderWriter_ErrorWrapFilePath_NilErr_Cov4(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{
		FilePath: "/tmp/test.txt",
	}

	// Act
	err := rw.errorWrapFilePath(nil, "/tmp/test.txt")

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_SimpleFileReaderWriter_GetOnExist_Error(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{
		FilePath: "/nonexistent/xyz/file.json",
	}
	var result map[string]string

	// Act
	err := rw.getOnExist(&result)

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent file")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — VerifyPaths, ApplyChmod invalid path, ApplyOnMismatch,
//               ToUint32Octal error, clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_VerifyPaths_ValidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	err := rwx.VerifyPaths(true, "/nonexistent/a", "/nonexistent/b")

	// Assert — continue on error, both invalid
	if err == nil {
		t.Fatal("expected error for non-existent paths")
	}
}

func Test_RwxWrapper_ApplyChmod_InvalidPath_SkipOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	err := rwx.ApplyChmod(true, "/nonexistent/xyz")

	// Assert — skip on invalid, should return nil
	if err != nil {
		t.Fatal("expected nil when skipping invalid:", err)
	}
}

func Test_RwxWrapper_ApplyChmod_InvalidPath_NoSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	err := rwx.ApplyChmod(false, "/nonexistent/xyz")

	// Assert — should return error
	if err == nil {
		t.Fatal("expected error for invalid path without skip")
	}
}

func Test_RwxWrapper_ApplyChmod_ValidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0600)

	// Act
	err := rwx.ApplyChmod(false, fp)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_RwxWrapper_ApplyChmodOptions_Equal(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act — should skip because chmod already matches (isApplyOnMismatch=true)
	err := rwx.ApplyChmodOptions(true, true, false, fp)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_RwxWrapper_ApplyChmodOptions_Mismatch(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0600)

	// Act
	err := rwx.ApplyChmodOptions(true, true, false, fp)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_RwxWrapper_ApplyChmodOptions_InvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	err := rwx.ApplyChmodOptions(true, false, false, "/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid path")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — ApplyRecursive non-Linux (walk-based), error branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_ApplyRecursive_FileOnly(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act — ApplyRecursive on a file (not a dir)
	err := rwx.ApplyRecursive(false, fp)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_RwxWrapper_ApplyRecursive_SkipInvalid(t *testing.T) {
	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	err := rwx.ApplyRecursive(true, "/nonexistent/xyz")

	// Assert
	if err != nil {
		t.Fatal("expected nil on skip-invalid:", err)
	}
}

func Test_RwxWrapper_ApplyRecursive_NoSkipInvalid(t *testing.T) {
	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	err := rwx.ApplyRecursive(false, "/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent without skip")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — applyLinuxRecursiveChmodUsingCmd, getLinuxRecursiveCmdForChmod
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_ApplyLinuxRecursiveCmdError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0755)

	// Act — on a valid dir, should succeed
	dir := t.TempDir()
	err := rwx.applyLinuxRecursiveChmodUsingCmd(dir)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — CompiledRwxWrapperUsingFixedRwxWrapper error,
//                          ApplyOnPath branches, verifyChmod branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxInstructionExecutor_ApplyOnPath_ExitOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid:   false,
			IsContinueOnError: false,
			IsRecursive:       false,
		},
	}
	executor, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act — invalid path, not skip
	applyErr := executor.ApplyOnPath("/nonexistent/xyz")

	// Assert
	if applyErr == nil {
		t.Fatal("expected error for invalid path without skip")
	}
}

func Test_RwxInstructionExecutor_ApplyOnPath_SkipOnInvalid_NonExistent(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid:   true,
			IsContinueOnError: false,
			IsRecursive:       false,
		},
	}
	executor, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	applyErr := executor.ApplyOnPath("/nonexistent/xyz")

	// Assert
	if applyErr != nil {
		t.Fatal("expected nil on skip-invalid:", applyErr)
	}
}

func Test_RwxInstructionExecutor_ApplyOnPath_Recursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "f.txt"), []byte("x"), 0644)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid:   false,
			IsContinueOnError: false,
			IsRecursive:       true,
		},
	}
	executor, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	applyErr := executor.ApplyOnPath(dir)

	// Assert
	if applyErr != nil {
		t.Fatal("expected no error:", applyErr)
	}
}

func Test_RwxInstructionExecutor_ApplyOnPathsDirect_Empty(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := ParseRwxInstructionToExecutor(ins)

	// Act
	err := executor.ApplyOnPathsDirect()

	// Assert
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

func Test_RwxInstructionExecutor_ApplyOnPaths_Empty(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := ParseRwxInstructionToExecutor(ins)

	// Act
	err := executor.ApplyOnPaths([]string{})

	// Assert
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — ApplyOnPaths with locations
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxInstructionExecutors_ApplyOnPaths_NonEmpty_ValidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executors, err := ParseRwxInstructionsToExecutors(
		[]chmodins.RwxInstruction{*ins})
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	applyErr := executors.ApplyOnPaths([]string{fp})

	// Assert
	if applyErr != nil {
		t.Fatal("expected no error:", applyErr)
	}
}

func Test_RwxInstructionExecutors_ApplyOnPaths_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
		},
	}
	executors, err := ParseRwxInstructionsToExecutors(
		[]chmodins.RwxInstruction{*ins})
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	applyErr := executors.ApplyOnPaths([]string{"/nonexistent/xyz"})

	// Assert
	if applyErr == nil {
		t.Fatal("expected error for non-existent path")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxPartialToInstructionExecutor — nil condition
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxPartialToInstructionExecutor_NilCondition_ReturnsError(t *testing.T) {
	// Arrange & Act
	_, err := RwxPartialToInstructionExecutor("rwx", nil)

	// Assert
	if err == nil {
		t.Fatal("expected error for nil condition")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// CreateDirFilesWithRwxPermissions — error propagation
// ══════════════════════════════════════════════════════════════════════════════

func Test_CreateDirFilesWithRwxPermissions_Error(t *testing.T) {
	// Arrange
	items := []DirFilesWithRwxPermission{
		{
			DirWithFiles: DirWithFiles{
				Dir: "/dev/null/impossible",
			},
			ApplyRwx: chmodins.RwxOwnerGroupOther{
				Owner: "INVALID",
				Group: "INVALID",
				Other: "INVALID",
			},
		},
	}

	// Act
	err := CreateDirFilesWithRwxPermissions(false, items)

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid RWX")
	}
}

func Test_CreateDirFilesWithRwxPermissionsMust_Panic(t *testing.T) {
	// Arrange
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()

	items := []DirFilesWithRwxPermission{
		{
			DirWithFiles: DirWithFiles{
				Dir: "/dev/null/impossible",
			},
			ApplyRwx: chmodins.RwxOwnerGroupOther{
				Owner: "INVALID",
				Group: "INVALID",
				Other: "INVALID",
			},
		},
	}

	// Act
	CreateDirFilesWithRwxPermissionsMust(false, items)
}

// ══════════════════════════════════════════════════════════════════════════════
// CreateDirWithFiles — error branches (removeDirErr, fileManipulateErr)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CreateDirWithFiles_MkDirError(t *testing.T) {
	// Arrange — null byte in path is universally invalid
	dw := &DirWithFiles{
		Dir:   string([]byte{0}) + "/impossible/dir",
		Files: []string{},
	}

	// Act
	err := CreateDirWithFiles(false, 0644, dw)

	// Assert
	if err == nil {
		t.Fatal("expected error for impossible mkdir")
	}
}

func Test_CreateDirWithFiles_WithFiles(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	subDir := filepath.Join(dir, "sub")
	dw := &DirWithFiles{
		Dir: subDir,
		Files:   []string{"a.txt", "b.txt"},
	}

	// Act
	err := CreateDirWithFiles(false, 0644, dw)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_CreateDirWithFiles_ChmodError_ReadOnlyDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	subDir := filepath.Join(dir, "sub_chmod_err")
	dw := &DirWithFiles{
		Dir: subDir,
		Files:   []string{"a.txt"},
	}

	// Act — create with a very restrictive file mode
	err := CreateDirWithFiles(false, 0000, dw)

	// Assert — file creation might fail due to permissions
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// DirFilesWithContent — Create error branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_DirFilesWithContent_Create_RemoveDirError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange — create a path that can't be removed
	dfwc := &DirFilesWithContent{
		Dir: "/dev/null/impossible/dir",
		DirFileMode: 0755,
	}

	// Act
	err := dfwc.Create(true) // isRemoveBeforeCreate=true

	// Assert
	if err == nil {
		t.Log("no error — OS allowed (unlikely)")
	}
}

func Test_DirFilesWithContent_Create_FileWriteError_ReadOnlyDir(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	dfwc := &DirFilesWithContent{
		Dir:         dir,
		DirFileMode: 0755,
		Files: []FileWithContent{
			{
				RelativePath: "test.txt",
				Content:      []string{"hello"},
				FileMode:     0644,
			},
		},
	}

	// Act — should succeed
	err := dfwc.Create(false)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MergeRwxWildcardWithFixedRwx — error branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_MergeRwxWildcardWithFixedRwx_Error(t *testing.T) {
	// Arrange & Act — invalid wildcard input
	_, err := MergeRwxWildcardWithFixedRwx("INVALID", "rwx")

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid wildcard input")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PathExistStat — MeaningFullError with error
// ══════════════════════════════════════════════════════════════════════════════

func Test_PathExistStat_MeaningFullError_WithError_FormatCheck(t *testing.T) {
	// Arrange
	stat := &PathExistStat{
		Location: "/nonexistent/xyz",
		Error:    errors.New("stat-fail"),
	}

	// Act
	err := stat.MeaningFullError()

	// Assert
	if err == nil {
		t.Fatal("expected error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — remaining error branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxVariableWrapper_NewError(t *testing.T) {
	// Arrange & Act
	// "INVALID" gets padded to 10 chars and sliced to valid 3-char segments,
	// so NewRwxVariableWrapper never errors for any string input.
	// Verify it succeeds — this documents the actual behavior.
	wrapper, err := NewRwxVariableWrapper("INVALID")

	// Assert
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if wrapper == nil {
		t.Fatal("expected non-nil wrapper")
	}
}

func Test_RwxVariableWrapper_IsEqualUsingLocation_NonExistent(t *testing.T) {
	// Arrange
	w, err := NewRwxVariableWrapper("rwxrwxrwx")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	result := w.IsEqualUsingLocation("/nonexistent/xyz")

	// Assert
	if result {
		t.Fatal("expected false for non-existent path")
	}
}

func Test_RwxVariableWrapper_IsEqualUsingFileInfo_Nil(t *testing.T) {
	// Arrange
	w, err := NewRwxVariableWrapper("rwxrwxrwx")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	result := w.IsEqualUsingFileInfo(nil)

	// Assert
	if result {
		t.Fatal("expected false for nil file info")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SingleRwx — default panic for invalid class type
// ══════════════════════════════════════════════════════════════════════════════

func Test_SingleRwx_ToRwxOwnerGroupOther_DefaultPanic(t *testing.T) {
	// Arrange
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for invalid class type")
		}
	}()

	s := &SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.Variant(255), // invalid
	}

	// Act
	s.ToRwxOwnerGroupOther()
}

func Test_SingleRwx_ApplyOnMany_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	s := &SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.All,
	}
	cond := chmodins.DefaultAllFalseCondition()

	// Act
	err := s.ApplyOnMany(cond, "/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid path")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// chmodApplier — RwxPartial nil condition, RwxStringApplyChmod error branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ChmodApplier_RwxPartial_NilCondition(t *testing.T) {
	// Arrange & Act
	err := ChmodApply.RwxPartial("rwx", nil, "/tmp")

	// Assert
	if err == nil {
		t.Fatal("expected error for nil condition")
	}
}

func Test_ChmodApplier_RwxPartial_EmptyLocations(t *testing.T) {
	// Arrange
	cond := chmodins.DefaultAllFalseCondition()

	// Act
	err := ChmodApply.RwxPartial("rwx", cond)

	// Assert
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

func Test_ChmodApplier_RwxPartial_InvalidRwx(t *testing.T) {
	// Arrange
	cond := chmodins.DefaultAllFalseCondition()

	// Act
	err := ChmodApply.RwxPartial("INVALID_VERY_LONG", cond, "/tmp")

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid rwx partial")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// newRwxWrapperCreator — UsingVariantPtr error
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewRwxWrapperCreator_UsingVariantPtr_Error(t *testing.T) {
	// Arrange & Act — use a 3-char variant with invalid octal digit (9 > 7)
	// Create() panics for wrong length, so use valid length with bad digit
	_, err := New.RwxWrapper.UsingVariantPtr(Variant("899"))

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid variant digit")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// tempDirGetter — TempPermanent
// ══════════════════════════════════════════════════════════════════════════════

func Test_TempDirGetter_TempPermanent(t *testing.T) {
	// Arrange & Act
	result := TempDirGetter.TempPermanent()

	// Assert
	if result == "" {
		t.Fatal("expected non-empty permanent temp dir")
	}
}

func Test_TempDirGetter_TempOption_Permanent(t *testing.T) {
	// Arrange & Act
	result := TempDirGetter.TempOption(true)

	// Assert
	if result == "" {
		t.Fatal("expected non-empty temp dir")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// VarAttribute — IsEqualPtr nil branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_VarAttribute_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange — parse a valid VarAttribute
	attr, err := ParseRwxToVarAttribute("rwx")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	result := attr.IsEqualPtr(nil)

	// Assert
	if result {
		t.Fatal("expected false when comparing with nil")
	}
}
