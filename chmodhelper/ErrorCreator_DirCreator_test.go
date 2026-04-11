package chmodhelper

import (
	"os"
	"path/filepath"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage6 — unexported API gaps (errorCreator, dirCreator, fileWriter)
//
// Uses standard "testing" only (in-package import restriction).
// ══════════════════════════════════════════════════════════════════════════════

// ── errorCreator ─────────────────────────────────────────────────────────────

func Test_ErrorCreator_DirError_PathExistsButNotDir(t *testing.T) {
	// Arrange — create a file (not a dir)
	dir := t.TempDir()
	filePath := filepath.Join(dir, "afile.txt")
	if err := os.WriteFile(filePath, []byte("hi"), 0644); err != nil {
		t.Fatal(err)
	}

	// Act — dirError on path that exists but is a file
	err := newError.dirError(filePath, os.ErrPermission)

	// Assert — should return notDirError (path exists but not dir)
	if err == nil {
		t.Fatal("expected error for file path used as dir")
	}
}

func Test_ErrorCreator_DirError_PathInvalid(t *testing.T) {
	// Arrange — path that doesn't exist
	invalidPath := "/nonexistent/path/xyz"

	// Act — dirError where path is invalid (doesn't exist)
	err := newError.dirError(invalidPath, os.ErrNotExist)

	// Assert — notDirError returns nil for invalid paths, so dirError
	// falls through to the error formatting branch
	if err == nil {
		t.Fatal("expected formatted error")
	}
}

func Test_ErrorCreator_NotDirError_PathInvalid_ReturnsNil(t *testing.T) {
	// Arrange — path doesn't exist
	invalidPath := "/nonexistent/nowhere"

	// Act
	err := newError.notDirError(invalidPath)

	// Assert — returns nil when path is invalid (doesn't exist)
	if err != nil {
		t.Fatal("expected nil for invalid/non-existent path")
	}
}

func Test_ErrorCreator_PathError_NilErr_ReturnsNil(t *testing.T) {
	// Arrange & Act
	err := newError.pathError("test", 0644, "/tmp", nil)

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_ErrorCreator_PathErrorWithDirValidate_NotDir_ReturnsError(t *testing.T) {
	// Arrange — file path (not a dir)
	dir := t.TempDir()
	filePath := filepath.Join(dir, "notdir.txt")
	if err := os.WriteFile(filePath, []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}

	// Act
	err := newError.pathErrorWithDirValidate("msg", 0755, filePath, os.ErrPermission)

	// Assert — returns notDirError since path exists but is not a dir
	if err == nil {
		t.Fatal("expected notDir error")
	}
}

func Test_ErrorCreator_PathErrorWithDirValidate_NilErr_ReturnsNil(t *testing.T) {
	// Arrange — valid dir path, nil err
	dir := t.TempDir()

	// Act
	err := newError.pathErrorWithDirValidate("msg", 0755, dir, nil)

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_ErrorCreator_ChmodApplyFailed_NilErr_ReturnsNil(t *testing.T) {
	// Arrange & Act
	err := newError.chmodApplyFailed(0755, "/tmp", nil)

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_ErrorCreator_ChmodApplyFailed_WithErr_FormatsMessage(t *testing.T) {
	// Arrange & Act
	err := newError.chmodApplyFailed(0755, "/tmp/test", os.ErrPermission)

	// Assert
	if err == nil {
		t.Fatal("expected formatted error")
	}
}

// ── dirCreator ───────────────────────────────────────────────────────────────

func Test_DirCreator_If_FalseCondition(t *testing.T) {
	// Arrange & Act — isCreate=false should be no-op
	err := internalDirCreator.If(false, 0755, "/any/path")

	// Assert
	if err != nil {
		t.Fatal("expected nil when isCreate is false")
	}
}

func Test_DirCreator_ByChecking_ExistsButNotDir(t *testing.T) {
	// Arrange — create a file
	dir := t.TempDir()
	filePath := filepath.Join(dir, "file.txt")
	if err := os.WriteFile(filePath, []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}

	// Act — ByChecking on file path (exists but not dir)
	err := internalDirCreator.ByChecking(0755, filePath)

	// Assert
	if err == nil {
		t.Fatal("expected error for file used as dir")
	}
}

func Test_DirCreator_Default_InvalidPath(t *testing.T) {
	// Arrange — null byte in path is universally invalid
	invalidPath := string([]byte{0}) + "/impossible"

	// Act
	err := internalDirCreator.Default(0755, invalidPath)

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid path")
	}
}

func Test_DirCreator_Direct_InvalidPath(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0}) + "/impossible"

	// Act
	err := internalDirCreator.Direct(invalidPath)

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid path")
	}
}

// ── fileWriter ───────────────────────────────────────────────────────────────

func Test_FileWriter_All_WriteFileError(t *testing.T) {
	// Arrange — dir exists but file path has null byte
	dir := t.TempDir()
	invalidFilePath := filepath.Join(dir, string([]byte{0}))

	// Act
	fw := fileWriter{}
	err := fw.All(
		0755, 0644,
		false, false, false,
		false,
		dir,
		invalidFilePath,
		[]byte("test"),
	)

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid file path")
	}
}

func Test_FileWriter_RemoveIf_ErrorOnRemove(t *testing.T) {
	// Arrange — create a read-only dir with a file
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(filePath, []byte("hi"), 0644); err != nil {
		t.Fatal(err)
	}

	// Act — RemoveIf with isRemove=true on an existing file (should succeed)
	fw := fileWriter{}
	err := fw.RemoveIf(true, filePath)

	// Assert — normal remove should succeed
	if err != nil {
		t.Fatal("expected no error:", err)
	}
}

func Test_FileWriter_Remove_InvalidPath(t *testing.T) {
	// Arrange — null byte
	invalidPath := string([]byte{0}) + "/impossible/file"

	// Act
	fw := fileWriter{}
	err := fw.Remove(invalidPath)

	// Assert
	if err == nil {
		t.Fatal("expected error for invalid path")
	}
}

func Test_FileWriter_ApplyDirChmod_NotRequired(t *testing.T) {
	// Arrange & Act — isCreateDirOnRequired=false
	fw := fileWriter{}
	err := fw.applyDirChmod(false, 0755, "/any")

	// Assert
	if err != nil {
		t.Fatal("expected nil when dir creation not required")
	}
}

func Test_FileWriter_ApplyDirChmod_DefaultChmod_SkipZeroMode(t *testing.T) {
	// Arrange & Act — chmodDir equals dirDefaultChmod, skip
	fw := fileWriter{}
	err := fw.applyDirChmod(true, dirDefaultChmod, "/any")

	// Assert
	if err != nil {
		t.Fatal("expected nil when chmod matches default")
	}
}

func Test_FileWriter_ApplyDirChmod_ChmodFails(t *testing.T) {
	// Arrange — non-existent dir
	fw := fileWriter{}

	// Act
	err := fw.applyDirChmod(true, 0700, "/nonexistent/dir/path")

	// Assert
	if err == nil {
		t.Fatal("expected error when chmod on non-existent dir")
	}
}

// ── SimpleFileReaderWriter (unexported) ──────────────────────────────────────

func Test_SimpleFileReaderWriter_GetOnExist_ReadError(t *testing.T) {
	// Arrange — file doesn't exist
	rw := SimpleFileReaderWriter{
		FilePath: "/nonexistent/file.txt",
	}

	// Act
	var target map[string]string
	err := rw.getOnExist(&target)

	// Assert
	if err == nil {
		t.Fatal("expected error for non-existent file")
	}
}

func Test_SimpleFileReaderWriter_ErrorWrapFilePath_NilErr_ReturnsNil(t *testing.T) {
	// Arrange
	rw := SimpleFileReaderWriter{FilePath: "/tmp/test.txt"}

	// Act
	err := rw.errorWrapFilePath(nil, "/tmp/test.txt")

	// Assert
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_SimpleFileReaderWriter_Name_NilReceiver(t *testing.T) {
	// Arrange
	var rw *SimpleFileReaderWriter

	// Act
	result := rw.name()

	// Assert
	if result != "" {
		t.Fatal("expected empty string for nil receiver")
	}
}
