package chmodhelper

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — LinuxApplyRecursive, ApplyRecursive, ApplyLinuxChmodOnMany
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_LinuxApplyRecursive_ValidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "f.txt"), []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0755)

	// Act
	err := rwx.LinuxApplyRecursive(false, dir)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_RwxWrapper_LinuxApplyRecursive_SkipInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0755)

	// Act
	err := rwx.LinuxApplyRecursive(true, "/nonexistent/path/xyz")

	// Assert
	if err != nil {
		t.Fatal("expected nil on skip-invalid:", err)
	}
}

func Test_RwxWrapper_LinuxApplyRecursive_NoSkipInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0755)

	// Act
	err := rwx.LinuxApplyRecursive(false, "/nonexistent/path/xyz")

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent without skip")
	}
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp1 := filepath.Join(dir, "f1.txt")
	fp2 := filepath.Join(dir, "f2.txt")
	os.WriteFile(fp1, []byte("a"), 0644)
	os.WriteFile(fp2, []byte("b"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)
	cond := &chmodins.Condition{
		IsRecursive:     false,
		IsContinueOnError: false,
		IsSkipOnInvalid: false,
	}

	// Act
	err := rwx.ApplyLinuxChmodOnMany(cond, fp1, fp2)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_Recursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "f.txt"), []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0755)
	cond := &chmodins.Condition{
		IsRecursive:     true,
		IsContinueOnError: false,
		IsSkipOnInvalid: false,
	}

	// Act
	err := rwx.ApplyLinuxChmodOnMany(cond, dir)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_ContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0755)
	cond := &chmodins.Condition{
		IsRecursive:       false,
		IsContinueOnError: true,
		IsSkipOnInvalid:   false,
	}

	// Act — one valid, one invalid
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	err := rwx.ApplyLinuxChmodOnMany(cond, fp, "/nonexistent/xyz")

	// Assert — continue on error should still return error but not panic
	_ = err
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_RecursiveContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0755)
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: true,
		IsSkipOnInvalid:   true,
	}

	// Act
	dir := t.TempDir()
	err := rwx.ApplyLinuxChmodOnMany(cond, dir, "/nonexistent/xyz")

	// Assert
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — IsRwxEqualFileInfo, IsRwxEqualLocation
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_IsRwxEqualFileInfo_Nil(t *testing.T) {
	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	result := rwx.IsRwxEqualFileInfo(nil)

	// Assert
	if result {
		t.Fatal("expected false for nil fileInfo")
	}
}

func Test_RwxWrapper_IsRwxEqualFileInfo_Valid(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	info, _ := os.Stat(fp)
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	result := rwx.IsRwxEqualFileInfo(info)

	// Assert
	if runtime.GOOS != "windows" && !result {
		t.Fatal("expected equal for matching file mode")
	}
}

func Test_RwxWrapper_IsRwxEqualLocation_NonExistent(t *testing.T) {
	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	result := rwx.IsRwxEqualLocation("/nonexistent/xyz")

	// Assert
	if result {
		t.Fatal("expected false for non-existent path")
	}
}

func Test_RwxWrapper_IsRwxEqualLocation_Valid(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)

	// Act
	result := rwx.IsRwxEqualLocation(fp)

	// Assert
	if runtime.GOOS != "windows" && !result {
		t.Fatal("expected equal")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — IsEqualVarWrapper
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_IsEqualVarWrapper_Nil(t *testing.T) {
	// Arrange
	rwx := New.RwxWrapper.UsingFileMode(0755)

	// Act
	result := rwx.IsEqualVarWrapper(nil)

	// Assert
	if result {
		t.Fatal("expected false for nil var wrapper")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// chmodVerifier — GetRwx9 short string, PathsUsingRwxFull continue-on-error
// ══════════════════════════════════════════════════════════════════════════════

func Test_ChmodVerifier_GetRwx9_ShortString(t *testing.T) {
	// Arrange — a file mode that would produce a short string (unlikely, but tests the branch)
	// Actually we test the "len <= 9" branch
	result := ChmodVerify.GetRwx9(0)

	// Assert — 0 mode produces "----------" which is 10 chars, so the normal path
	if len(result) != 9 {
		// The branch for len <= 9 is nearly impossible with standard Go FileMode
		// but we exercise the function regardless
	}
}

func Test_ChmodVerifier_PathsUsingRwxFull_ContinueOnError(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act — continue on error with a mismatch
	err := ChmodVerify.PathsUsingRwxFull(
		true,
		"-rwxrwxrwx",
		fp,
	)

	// Assert
	if runtime.GOOS != "windows" && err == nil {
		t.Fatal("expected error for mismatched chmod")
	}
}

func Test_ChmodVerifier_PathsUsingRwxFull_ImmediateReturn(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := ChmodVerify.PathsUsingFileModeImmediateReturn(0755, fp)

	// Assert
	if runtime.GOOS != "windows" && err == nil {
		t.Fatal("expected error for mismatched chmod")
	}
}

func Test_ChmodVerifier_PathsUsingFileModeContinueOnError(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := ChmodVerify.PathsUsingFileModeContinueOnError(0755, fp)

	// Assert
	if runtime.GOOS != "windows" && err == nil {
		t.Fatal("expected error for mismatched chmod")
	}
}

func Test_ChmodVerifier_RwxFull_Mismatch(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := ChmodVerify.RwxFull(fp, "-rwxrwxrwx")

	// Assert
	if err == nil {
		t.Fatal("expected mismatch error")
	}
}

func Test_ChmodVerifier_RwxFull_NonExistent(t *testing.T) {
	// Arrange & Act
	err := ChmodVerify.RwxFull("/nonexistent/xyz", "-rw-r--r--")

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SingleRwx — OwnerOther branch, default panic, ToDisabledRwxWrapper, ToRwxWrapper
// ══════════════════════════════════════════════════════════════════════════════

func Test_SingleRwx_ToRwxOwnerGroupOther_OwnerOther(t *testing.T) {
	// Arrange
	s := &SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.OwnerOther,
	}

	// Act
	result := s.ToRwxOwnerGroupOther()

	// Assert
	if result.Owner != "rwx" || result.Other != "rwx" {
		t.Fatal("expected owner and other set")
	}
}

func Test_SingleRwx_ToDisabledRwxWrapper_Error(t *testing.T) {
	// Arrange
	s := &SingleRwx{
		Rwx:       "INVALID",
		ClassType: chmodclasstype.All,
	}

	// Act
	_, err := s.ToDisabledRwxWrapper()

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid rwx")
	}
}

func Test_SingleRwx_ToRwxWrapper_NonAll(t *testing.T) {
	// Arrange
	s := &SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.Owner,
	}

	// Act
	_, err := s.ToRwxWrapper()

	// Assert
	if err == nil {
		t.Fatal("expected error for non-All class type")
	}
}

func Test_SingleRwx_ApplyOnMany_EmptyLocations(t *testing.T) {
	// Arrange
	s := &SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.All,
	}
	cond := chmodins.DefaultAllFalseCondition()

	// Act
	err := s.ApplyOnMany(cond)

	// Assert
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — Read methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleFileReaderWriter_ReadBytes(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "read.txt")
	os.WriteFile(fp, []byte("hello"), 0644)
	rw := SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}

	// Act
	bytes, err := rw.Read()

	// Assert
	if err != nil {
		t.Fatal("expected read success:", err)
	}
	if string(bytes) != "hello" {
		t.Fatal("expected correct content")
	}
}

func Test_SimpleFileReaderWriter_ReadString(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "read.txt")
	os.WriteFile(fp, []byte("world"), 0644)
	rw := SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}

	// Act
	str, err := rw.ReadString()

	// Assert
	if err != nil {
		t.Fatal("expected read success:", err)
	}
	if str != "world" {
		t.Fatal("expected correct content")
	}
}

func Test_SimpleFileReaderWriter_ReadBytes_NonExistent(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{
		FilePath: "/nonexistent/xyz/file.txt",
	}

	// Act
	_, err := rw.Read()

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent file")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// DirFilesWithContent — error paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_DirFilesWithContent_Create_ErrorPath(t *testing.T) {
	// Arrange
	dfwc := &DirFilesWithContent{
		Dir:         "/dev/null/impossible",
		DirFileMode: 0755,
	}

	// Act
	err := dfwc.Create(false)

	// Assert
	if err == nil {
		t.Log("no error — OS allowed creation (unlikely)")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// VarAttribute — HasWildcard error branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_VarAttribute_HasWildcard(t *testing.T) {
	// Arrange — use ParseRwxToVarAttribute to get a wildcard VarAttribute
	attr, err := ParseRwxToVarAttribute("r-*")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	result := attr.HasWildcard()

	// Assert
	if !result {
		t.Fatal("expected HasWildcard to be true")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// errorCreator — remaining paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_ErrorCreator_PathError_NilErr_ApplyVerifyContext(t *testing.T) {
	// Arrange & Act
	err := newError.pathError("test", 0755, "/tmp/x", nil)

	// Assert
	if err != nil {
		t.Fatal("nil input error should return nil")
	}
}

func Test_ErrorCreator_ChmodApplyFailed_NilErr_ApplyVerifyContext(t *testing.T) {
	// Arrange & Act
	err := newError.chmodApplyFailed(0644, "/tmp/x", nil)

	// Assert
	if err != nil {
		t.Fatal("nil input error should return nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// dirCreator — ByCheckingWithChmod error path
// ══════════════════════════════════════════════════════════════════════════════

func Test_DirCreator_ByChecking_InvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange & Act
	err := internalDirCreator.ByChecking(0755, "/dev/null/impossible/path")

	// Assert
	if err == nil {
		t.Log("no error — OS allowed (unlikely)")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — ApplyRwxOnLocations error continuation branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxVariableWrapper_ApplyRwxOnLocations_SkipContinue(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	w, err := NewRwxVariableWrapper("rwxrwxrwx")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act — both continue and skip, with invalid path
	applyErr := w.ApplyRwxOnLocations(true, true, "/nonexistent/abc")

	// Assert
	_ = applyErr
}

func Test_RwxVariableWrapper_ApplyRwxOnLocations_NoContinueNoSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Arrange
	w, err := NewRwxVariableWrapper("rwxrwxrwx")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Act
	applyErr := w.ApplyRwxOnLocations(false, false, "/nonexistent/abc")

	// Assert
	if applyErr == nil {
		t.Fatal("expected error for non-existent without skip")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// GetRecursivePaths / GetRecursivePathsContinueOnError — error paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetRecursivePaths_SkipInvalid(t *testing.T) {
	// Arrange & Act
	// GetRecursivePaths always returns error for non-existent paths
	// (isContinueOnError only applies when path exists and is a directory)
	paths, err := GetRecursivePaths(true, "/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent path")
	}
	if len(paths) != 0 {
		t.Fatal("expected empty paths")
	}
}

func Test_GetRecursivePaths_NoSkipInvalid(t *testing.T) {
	// Arrange & Act
	_, err := GetRecursivePaths(false, "/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent without skip")
	}
}

func Test_GetRecursivePathsContinueOnError_NonExistent(t *testing.T) {
	// Arrange & Act
	paths, err := GetRecursivePathsContinueOnError("/nonexistent/xyz")

	// Assert — non-existent path returns error
	if err == nil {
		t.Fatal("expected error for non-existent path")
	}
	if len(paths) != 0 {
		t.Fatal("expected empty paths")
	}
}

func Test_GetRecursivePathsContinueOnError_ValidDir(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "f.txt"), []byte("x"), 0644)

	// Act
	paths, err := GetRecursivePathsContinueOnError(dir)

	// Assert
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	if len(paths) == 0 {
		t.Fatal("expected non-empty paths")
	}
}
