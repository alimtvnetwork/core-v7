package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"github.com/alimtvnetwork/core/coretests/args"
)

// newTestRW is defined in shared_coverage_helpers.go

// ── SimpleFileReaderWriter.Write ──

func Test_Write_Error(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.Write([]byte("hello"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Write_Success(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_write")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.Write([]byte("hello"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── SimpleFileReaderWriter.WritePath ──

func Test_WritePath_Error(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.WritePath(false, filepath.Join(invalidDir, "wp.txt"), []byte("x"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_WritePath_Success(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_writepath")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	fp := filepath.Join(tmpDir, "wp.txt")
	err := rw.WritePath(false, fp, []byte("hello"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── SimpleFileReaderWriter.WriteRelativePath ──

func Test_WriteRelativePath_Error(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.WriteRelativePath(false, "rel.txt", []byte("x"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_WriteRelativePath_Success(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_writerel")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.WriteRelativePath(false, "rel.txt", []byte("hello"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── SimpleFileReaderWriter.WriteString ──

func Test_WriteString_Error(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.WriteString("hello")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_WriteString_Success(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_writestr")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.WriteString("hello")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── SimpleFileReaderWriter.WriteAny ──

func Test_WriteAny_Error(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.json")
	err := rw.WriteAny(map[string]string{"k": "v"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── SimpleFileReaderWriter.Read ──

func Test_Read_Error(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	_, err := rw.Read()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Read_Success(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_read")
	os.MkdirAll(tmpDir, 0755)
	fp := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(fp, []byte("hello"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	data, err := rw.Read()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": string(data) != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected content", actual)
}

// ── SimpleFileReaderWriter.ReadMust ──

func Test_ReadMust_Panic(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	rw.ReadMust()
}

// ── SimpleFileReaderWriter.ReadString ──

func Test_ReadString_Error(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	_, err := rw.ReadString()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ReadString_Success(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_readstr")
	os.MkdirAll(tmpDir, 0755)
	fp := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(fp, []byte("world"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	s, err := rw.ReadString()

	// Act
	actual := args.Map{"result": err != nil || s != "world"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ── SimpleFileReaderWriter.ReadStringMust ──

func Test_ReadStringMust_Panic(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	rw.ReadStringMust()
}

// ── SimpleFileReaderWriter.GetSet ──

func Test_GetSet_GenerateError(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_getset")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "cache.json")
	var result map[string]string

	err := rw.GetSet(&result, func() (any, error) {
		return nil, os.ErrNotExist
	})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error from generate func", actual)
}

// ── SimpleFileReaderWriter.errorWrap / errorWrapFilePath ──

func Test_ErrorWrap_Nil(t *testing.T) {
	// Arrange
	// errorWrap with nil returns nil - covered through successful write
	tmpDir := filepath.Join(os.TempDir(), "cov12_errwrap")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.Write([]byte("x"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_ErrorWrapFilePath_Nil(t *testing.T) {
	// Covered through successful WritePath
}

// ── SimpleFileReaderWriter.name ──

// ── SimpleFileReaderWriter.name (unexported, tested indirectly via errorWrap) ──

func Test_Name_CoveredViaErrorWrap(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	// name() is unexported; it's called inside errorWrapFilePath
	// which is triggered by any Write error
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.Write([]byte("x"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// error message includes "simple-reader-writer" from name()
}

// ── SimpleFileReaderWriter.getOnExist ──

func Test_GetOnExist_ReadError(t *testing.T) {
	// Arrange
	rw := newTestRW("/nonexistent/cov12", "cache.json")
	// File doesn't exist but Get calls getOnExist only if IsExist
	err := rw.Get(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_GetOnExist_Valid(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov12_getonexist")
	os.MkdirAll(tmpDir, 0755)
	fp := filepath.Join(tmpDir, "data.json")
	os.WriteFile(fp, []byte(`{"key":"val"}`), 0644)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "data.json")
	var result map[string]string
	err := rw.Get(&result)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}
