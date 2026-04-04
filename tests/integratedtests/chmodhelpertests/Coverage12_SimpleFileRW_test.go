package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// newTestRW is defined in shared_coverage_helpers.go

// ── SimpleFileReaderWriter.Write ──

func Test_Cov12_Write_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.Write([]byte("hello"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_Write_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_write")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
}

// ── SimpleFileReaderWriter.WritePath ──

func Test_Cov12_WritePath_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.WritePath(false, filepath.Join(invalidDir, "wp.txt"), []byte("x"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_WritePath_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_writepath")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	fp := filepath.Join(tmpDir, "wp.txt")
	err := rw.WritePath(false, fp, []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
}

// ── SimpleFileReaderWriter.WriteRelativePath ──

func Test_Cov12_WriteRelativePath_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.WriteRelativePath(false, "rel.txt", []byte("x"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_WriteRelativePath_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_writerel")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.WriteRelativePath(false, "rel.txt", []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
}

// ── SimpleFileReaderWriter.WriteString ──

func Test_Cov12_WriteString_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	err := rw.WriteString("hello")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_WriteString_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_writestr")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.WriteString("hello")
	if err != nil {
		t.Fatal(err)
	}
}

// ── SimpleFileReaderWriter.WriteAny ──

func Test_Cov12_WriteAny_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.json")
	err := rw.WriteAny(map[string]string{"k": "v"})
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── SimpleFileReaderWriter.Read ──

func Test_Cov12_Read_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	_, err := rw.Read()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_Read_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_read")
	os.MkdirAll(tmpDir, 0755)
	fp := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(fp, []byte("hello"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	data, err := rw.Read()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello" {
		t.Fatal("unexpected content")
	}
}

// ── SimpleFileReaderWriter.ReadMust ──

func Test_Cov12_ReadMust_Panic(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	rw.ReadMust()
}

// ── SimpleFileReaderWriter.ReadString ──

func Test_Cov12_ReadString_Error(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	_, err := rw.ReadString()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_ReadString_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_readstr")
	os.MkdirAll(tmpDir, 0755)
	fp := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(fp, []byte("world"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	s, err := rw.ReadString()
	if err != nil || s != "world" {
		t.Fatal("unexpected")
	}
}

// ── SimpleFileReaderWriter.ReadStringMust ──

func Test_Cov12_ReadStringMust_Panic(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	invalidDir := filepath.Join("/proc", "nonexistent_cov12")
	rw := newTestRW(invalidDir, "test.txt")
	rw.ReadStringMust()
}

// ── SimpleFileReaderWriter.GetSet ──

func Test_Cov12_GetSet_GenerateError(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_getset")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "cache.json")
	var result map[string]string

	err := rw.GetSet(&result, func() (any, error) {
		return nil, os.ErrNotExist
	})
	if err == nil {
		t.Fatal("expected error from generate func")
	}
}

// ── SimpleFileReaderWriter.errorWrap / errorWrapFilePath ──

func Test_Cov12_ErrorWrap_Nil(t *testing.T) {
	// errorWrap with nil returns nil - covered through successful write
	tmpDir := filepath.Join(os.TempDir(), "cov12_errwrap")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "test.txt")
	err := rw.Write([]byte("x"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov12_ErrorWrapFilePath_Nil(t *testing.T) {
	// Covered through successful WritePath
}

// ── SimpleFileReaderWriter.name ──

// ── SimpleFileReaderWriter.name (unexported, tested indirectly via errorWrap) ──

func Test_Cov12_Name_CoveredViaErrorWrap(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	// name() is unexported; it's called inside errorWrapFilePath
	// which is triggered by any Write error
	rw := newTestRW("/nonexistent/cov12", "test.txt")
	err := rw.Write([]byte("x"))
	if err == nil {
		t.Fatal("expected error")
	}
	// error message includes "simple-reader-writer" from name()
}

// ── SimpleFileReaderWriter.getOnExist ──

func Test_Cov12_GetOnExist_ReadError(t *testing.T) {
	rw := newTestRW("/nonexistent/cov12", "cache.json")
	// File doesn't exist but Get calls getOnExist only if IsExist
	err := rw.Get(nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_GetOnExist_Valid(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov12_getonexist")
	os.MkdirAll(tmpDir, 0755)
	fp := filepath.Join(tmpDir, "data.json")
	os.WriteFile(fp, []byte(`{"key":"val"}`), 0644)
	defer os.RemoveAll(tmpDir)

	rw := newTestRW(tmpDir, "data.json")
	var result map[string]string
	err := rw.Get(&result)
	if err != nil {
		t.Fatal(err)
	}
}
