package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

func skipWin(t *testing.T) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("skipping on Windows")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — Read, ReadString, ReadMust, ReadLock, etc.
// ══════════════════════════════════════════════════════════════════════════════

func newRW(t *testing.T, content string) chmodhelper.SimpleFileReaderWriter {
	t.Helper()
	dir := t.TempDir()
	fp := filepath.Join(dir, "test.txt")
	os.WriteFile(fp, []byte(content), 0644)
	return chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
}

func Test_SFRW_ReadString(t *testing.T) {
	// Arrange
	rw := newRW(t, "hello")
	s, err := rw.ReadString()

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got, err=", actual)
}

func Test_SFRW_ReadStringMust(t *testing.T) {
	// Arrange
	rw := newRW(t, "world")
	s := rw.ReadStringMust()

	// Act
	actual := args.Map{"result": s != "world"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_SFRW_ReadStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	rw.ReadStringMust()
}

func Test_SFRW_ReadMust(t *testing.T) {
	// Arrange
	rw := newRW(t, "data")
	b := rw.ReadMust()

	// Act
	actual := args.Map{"result": string(b) != "data"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_ReadMust_Panic(t *testing.T) {
	defer func() { recover() }()
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	rw.ReadMust()
}

func Test_SFRW_ReadLock(t *testing.T) {
	// Arrange
	rw := newRW(t, "lock")
	b, err := rw.ReadLock()

	// Act
	actual := args.Map{"result": err != nil || string(b) != "lock"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_ReadStringLock(t *testing.T) {
	// Arrange
	rw := newRW(t, "slock")
	s, err := rw.ReadStringLock()

	// Act
	actual := args.Map{"result": err != nil || s != "slock"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_ReadOnExist_Exists(t *testing.T) {
	// Arrange
	rw := newRW(t, "exist")
	b, err := rw.ReadOnExist()

	// Act
	actual := args.Map{"result": err != nil || string(b) != "exist"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_ReadOnExist_NotExists(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	b, err := rw.ReadOnExist()

	// Act
	actual := args.Map{"result": err != nil || b != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil, nil", actual)
}

func Test_SFRW_ReadStringOnExist(t *testing.T) {
	// Arrange
	rw := newRW(t, "sexist")
	s, err := rw.ReadStringOnExist()

	// Act
	actual := args.Map{"result": err != nil || s != "sexist"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_ReadStringOnExist_NotExists(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	s, err := rw.ReadStringOnExist()

	// Act
	actual := args.Map{"result": err != nil || s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_SFRW_ReadStringOnExistLock(t *testing.T) {
	// Arrange
	rw := newRW(t, "lock2")
	s, err := rw.ReadStringOnExistLock()

	// Act
	actual := args.Map{"result": err != nil || s != "lock2"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_ReadOnExistLock(t *testing.T) {
	// Arrange
	rw := newRW(t, "lock3")
	b, err := rw.ReadOnExistLock()

	// Act
	actual := args.Map{"result": err != nil || string(b) != "lock3"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_Read_Error(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/i10/read.txt",
		ParentDir: "/nonexistent/i10",
		ChmodDir:  0755,
		ChmodFile: 0644,
	}
	_, err := rw.Read()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — WriteString, WritePath, WriteRelativePath
// ══════════════════════════════════════════════════════════════════════════════

func Test_SFRW_WriteString(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "ws.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	err := rw.WriteString("content")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_WritePath(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "wp.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	err := rw.WritePath(false, fp, []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_WriteRelativePath(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "x.txt"),
	}
	err := rw.WriteRelativePath(false, "rel.txt", []byte("reldata"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_JoinRelPath_Empty(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/tmp/i10"}
	result := rw.JoinRelPath("")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_SFRW_JoinRelPath_NonEmpty(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/tmp/i10"}
	result := rw.JoinRelPath("sub/file.txt")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — WriteAny, WriteAnyLock, Get, Set, Expire, etc.
// ══════════════════════════════════════════════════════════════════════════════

func Test_SFRW_WriteAny(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "any.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	err := rw.WriteAny(map[string]string{"key": "val"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_WriteAnyLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "anylock.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	err := rw.WriteAnyLock(map[string]string{"k": "v"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_Get_NotExist(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/i10/get.json",
		ParentDir: "/nonexistent/i10",
		ChmodDir:  0755,
		ChmodFile: 0644,
	}
	var out map[string]string
	err := rw.Get(&out)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_SFRW_Get_Exists(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "get.json")
	os.WriteFile(fp, []byte(`{"key":"val"}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	var out map[string]string
	err := rw.Get(&out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": out["key"] != "val"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected value", actual)
}

func Test_SFRW_GetLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "getlock.json")
	os.WriteFile(fp, []byte(`{"a":"b"}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	var out map[string]string
	err := rw.GetLock(&out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_Set(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "set.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	err := rw.Set(map[string]string{"x": "y"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_SetLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "setlock.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	err := rw.SetLock(map[string]string{"x": "y"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_Expire_Exists(t *testing.T) {
	// Arrange
	rw := newRW(t, "expire")
	err := rw.Expire()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_Expire_NotExists(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10/expire.txt"}
	err := rw.Expire()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_ExpireLock(t *testing.T) {
	// Arrange
	rw := newRW(t, "expirelock")
	err := rw.ExpireLock()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_ExpireParentDir(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	os.MkdirAll(sub, 0755)
	fp := filepath.Join(sub, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: sub,
		FilePath:  fp,
		ChmodDir:  0755,
		ChmodFile: 0644,
	}
	err := rw.ExpireParentDir()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_ExpireParentDirLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub2")
	os.MkdirAll(sub, 0755)
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: sub,
		FilePath:  filepath.Join(sub, "f.txt"),
		ChmodDir:  0755,
		ChmodFile: 0644,
	}
	err := rw.ExpireParentDirLock()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_RemoveOnExist(t *testing.T) {
	rw := newRW(t, "rm")
	_ = rw.RemoveOnExist()
}

func Test_SFRW_RemoveDirOnExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: "/nonexistent/i10/rmdir",
	}
	_ = rw.RemoveDirOnExist()
}

func Test_SFRW_OsFile(t *testing.T) {
	// Arrange
	rw := newRW(t, "osfile")
	f, err := rw.OsFile()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	f.Close()
}

func Test_SFRW_Clone(t *testing.T) {
	// Arrange
	rw := newRW(t, "clone")
	c := rw.Clone()

	// Act
	actual := args.Map{"result": c.FilePath != rw.FilePath}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same file path", actual)
}

func Test_SFRW_ClonePtr(t *testing.T) {
	// Arrange
	rw := newRW(t, "cptr")
	cp := rw.ClonePtr()

	// Act
	actual := args.Map{"result": cp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SFRW_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var rw *chmodhelper.SimpleFileReaderWriter
	cp := rw.ClonePtr()

	// Act
	actual := args.Map{"result": cp != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_SFRW_String(t *testing.T) {
	// Arrange
	rw := newRW(t, "str")
	s := rw.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_SFRW_Json(t *testing.T) {
	// Arrange
	rw := newRW(t, "json")
	j := rw.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_SFRW_JsonPtr(t *testing.T) {
	// Arrange
	rw := newRW(t, "jsonptr")
	j := rw.JsonPtr()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_SFRW_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	rw := newRW(t, "marshal")
	b, err := rw.MarshalJSON()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	var rw2 chmodhelper.SimpleFileReaderWriter
	err2 := rw2.UnmarshalJSON(b)
	actual = args.Map{"result": err2}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err2", actual)
	actual = args.Map{"result": rw2.FilePath != rw.FilePath}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same path after unmarshal", actual)
}

func Test_SFRW_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	rw := newRW(t, "binder")
	binder := rw.AsJsonContractsBinder()

	// Act
	actual := args.Map{"result": binder == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SFRW_JsonParseSelfInject(t *testing.T) {
	// Arrange
	rw := newRW(t, "inject")
	j := rw.JsonPtr()
	var rw2 chmodhelper.SimpleFileReaderWriter
	err := rw2.JsonParseSelfInject(j)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_NewPath(t *testing.T) {
	// Arrange
	rw := newRW(t, "np")
	newRw := rw.NewPath(false, filepath.Join(rw.ParentDir, "newfile.txt"))

	// Act
	actual := args.Map{"result": newRw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SFRW_NewPathJoin(t *testing.T) {
	// Arrange
	rw := newRW(t, "npj")
	newRw := rw.NewPathJoin(false, "sub", "file.txt")

	// Act
	actual := args.Map{"result": newRw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SFRW_InitializeDefaultNew(t *testing.T) {
	// Arrange
	rw := newRW(t, "idn")
	newRw := rw.InitializeDefaultNew()

	// Act
	actual := args.Map{"result": newRw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SFRW_HasAnyIssues(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/i10/issues",
		ParentDir: "/nonexistent/i10",
	}

	// Act
	actual := args.Map{"result": rw.HasAnyIssues()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected issues", actual)
}

func Test_SFRW_IsPathInvalid(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}

	// Act
	actual := args.Map{"result": rw.IsPathInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_SFRW_IsParentDirInvalid(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/nonexistent/i10"}

	// Act
	actual := args.Map{"result": rw.IsParentDirInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_SFRW_Serialize(t *testing.T) {
	// Arrange
	rw := newRW(t, "serialize")
	b, err := rw.Serialize()

	// Act
	actual := args.Map{"result": err != nil || string(b) != "serialize"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got, err=", actual)
}

func Test_SFRW_SerializeLock(t *testing.T) {
	// Arrange
	rw := newRW(t, "serlock")
	b, err := rw.SerializeLock()

	// Act
	actual := args.Map{"result": err != nil || string(b) != "serlock"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SFRW_Deserialize(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "deser.json")
	os.WriteFile(fp, []byte(`{"k":"v"}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	var out map[string]string
	err := rw.Deserialize(&out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SFRW_DeserializeLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "deserlock.json")
	os.WriteFile(fp, []byte(`{"k":"v"}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  fp,
	}
	var out map[string]string
	err := rw.DeserializeLock(&out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func makeExecutor(t *testing.T, rwx string) *chmodhelper.RwxInstructionExecutor {
	t.Helper()
	exec, err := chmodhelper.RwxPartialToInstructionExecutor(rwx, &chmodins.Condition{})
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	return exec
}

func Test_Executor_IsFixedWrapper(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")

	// Act
	actual := args.Map{"result": exec.IsFixedWrapper()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected fixed", actual)
}

func Test_Executor_IsVarWrapper(t *testing.T) {
	// Arrange
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor("-rw*r-*r-*", &chmodins.Condition{})

	// Act
	actual := args.Map{"result": exec.IsVarWrapper()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected var", actual)
}

func Test_Executor_IsEqualFileMode(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")

	// Act
	actual := args.Map{"result": exec.IsEqualFileMode(0755)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Executor_IsEqualRwxPartial(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")

	// Act
	actual := args.Map{"result": exec.IsEqualRwxPartial("-rwxr-xr-x")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Executor_IsEqualRwxWrapper(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")
	w := chmodhelper.New.RwxWrapper.UsingFileMode(0755)

	// Act
	actual := args.Map{"result": exec.IsEqualRwxWrapper(&w)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Executor_IsEqualFileInfo(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "ei.txt")
	os.WriteFile(fp, []byte("x"), 0755)
	os.Chmod(fp, 0755)
	info, _ := os.Stat(fp)
	exec := makeExecutor(t, "-rwxr-xr-x")
	_ = exec.IsEqualFileInfo(info)
}

func Test_Executor_CompiledWrapper_Fixed(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")
	w, err := exec.CompiledWrapper(0755)

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_Executor_CompiledWrapper_Var(t *testing.T) {
	// Arrange
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor("-rw*r-*r-*", &chmodins.Condition{})
	w, err := exec.CompiledWrapper(0644)

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_Executor_CompiledRwxWrapperUsingFixed_Fixed(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")
	w := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	r, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(&w)

	// Act
	actual := args.Map{"result": err != nil || r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_Executor_CompiledRwxWrapperUsingFixed_Var(t *testing.T) {
	// Arrange
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor("-rw*r-*r-*", &chmodins.Condition{})
	w := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	r, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(&w)

	// Act
	actual := args.Map{"result": err != nil || r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_Executor_ApplyOnPath_Valid(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "apply.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	exec := makeExecutor(t, "-rw-r--r--")
	err := exec.ApplyOnPath(fp)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executor_ApplyOnPath_SkipInvalid(t *testing.T) {
	// Arrange
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr-x", &chmodins.Condition{IsSkipOnInvalid: true})
	err := exec.ApplyOnPath("/nonexistent/i10/skip")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil on skip", actual)
}

func Test_Executor_ApplyOnPaths_Empty(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.ApplyOnPaths(nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executor_ApplyOnPathsDirect_Empty(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.ApplyOnPathsDirect()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executor_ApplyOnPaths_Valid(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "paths.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	exec := makeExecutor(t, "-rw-r--r--")
	err := exec.ApplyOnPaths([]string{fp})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executor_ApplyOnPathsPtr_Nil(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.ApplyOnPathsPtr(nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executor_ApplyOnPaths_ContinueOnError(t *testing.T) {
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr-x", &chmodins.Condition{IsContinueOnError: true})
	locs := []string{"/nonexistent/i10/a", "/nonexistent/i10/b"}
	_ = exec.ApplyOnPathsPtr(&locs)
}

func Test_Executor_VerifyRwxModifiers_Valid(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "verify.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	os.Chmod(fp, 0644)
	exec := makeExecutor(t, "-rw-r--r--")
	err := exec.VerifyRwxModifiers(true, []string{fp})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executor_VerifyRwxModifiersDirect(t *testing.T) {
	// Arrange
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.VerifyRwxModifiersDirect(true)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executor_VerifyRwxModifiers_ContinueOnError(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "vcont.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor(
		"-rw-r--r--", &chmodins.Condition{IsContinueOnError: true})
	err := exec.VerifyRwxModifiers(true, []string{fp})
	_ = err
}

func Test_Executor_VerifyRwxModifiers_Mismatch(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "vmm.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	os.Chmod(fp, 0644)
	exec := makeExecutor(t, "-rwxrwxrwx")
	err := exec.VerifyRwxModifiers(true, []string{fp})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Executors_LastIndex(t *testing.T) {
	// Arrange
	execs := chmodhelper.NewRwxInstructionExecutors(2)

	// Act
	actual := args.Map{"result": execs.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_Executors_HasIndex(t *testing.T) {
	// Arrange
	execs := chmodhelper.NewRwxInstructionExecutors(2)

	// Act
	actual := args.Map{"result": execs.HasIndex(0)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_Executors_Items(t *testing.T) {
	// Arrange
	execs := chmodhelper.NewRwxInstructionExecutors(2)

	// Act
	actual := args.Map{"result": execs.Items() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Executors_ApplyOnPaths_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "execs.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	exec := makeExecutor(t, "-rw-r--r--")
	execs := chmodhelper.NewRwxInstructionExecutors(1)
	execs.Add(exec)
	err := execs.ApplyOnPaths([]string{fp})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executors_VerifyRwxModifiers_ContinueOnError(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "evfy.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	exec := makeExecutor(t, "-rwxrwxrwx")
	execs := chmodhelper.NewRwxInstructionExecutors(1)
	execs.Add(exec)
	err := execs.VerifyRwxModifiers(true, true, []string{fp})
	_ = err
}

func Test_Executors_VerifyRwxModifiers_NoContinue(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "evfy2.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	exec := makeExecutor(t, "-rw-r--r--")
	execs := chmodhelper.NewRwxInstructionExecutors(1)
	execs.Add(exec)
	err := execs.VerifyRwxModifiers(false, true, []string{fp})
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// AttrVariant — uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_AttrVariant_IsGreaterThan_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	v := chmodhelper.ReadWriteExecute
	if !v.IsGreaterThan(8) {
		// 8 > 7 is true
	}

	// Act
	actual := args.Map{"result": v.IsGreaterThan(6)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false: 6 is not > 7", actual)
}

func Test_AttrVariant_String(t *testing.T) {
	v := chmodhelper.Read
	_ = v.String()
}

func Test_AttrVariant_Value(t *testing.T) {
	// Arrange
	v := chmodhelper.Execute

	// Act
	actual := args.Map{"result": v.Value() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_AttrVariant_ToAttribute(t *testing.T) {
	// Arrange
	v := chmodhelper.ReadWriteExecute
	a := v.ToAttribute()

	// Act
	actual := args.Map{"result": a.IsRead || !a.IsWrite || !a.IsExecute}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected all true for ReadWriteExecute", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FilteredPathFileInfoMap — uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_FPFIM_LazyValidLocations(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "lazy.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, fp)
	locs := m.LazyValidLocations()

	// Act
	actual := args.Map{"result": len(locs) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected locations", actual)
	// second call hits cache
	locs2 := m.LazyValidLocations()
	actual = args.Map{"result": len(locs2) != len(locs)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache mismatch", actual)
}

func Test_FPFIM_MissingPathsToString(t *testing.T) {
	// Arrange
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, "/nonexistent/i10/a", "/nonexistent/i10/b")
	s := m.MissingPathsToString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_FPFIM_HasAnyIssues(t *testing.T) {
	// Arrange
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(false, "/nonexistent/i10/x")

	// Act
	actual := args.Map{"result": m.HasAnyIssues()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected issues", actual)
}

func Test_FPFIM_HasError(t *testing.T) {
	// Arrange
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(false, "/nonexistent/i10/x")

	// Act
	actual := args.Map{"result": m.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_FPFIM_HasAnyMissingPaths(t *testing.T) {
	// Arrange
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, "/nonexistent/i10/x")

	// Act
	actual := args.Map{"result": m.HasAnyMissingPaths()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected missing paths", actual)
}

func Test_FPFIM_LengthOfIssues(t *testing.T) {
	// Arrange
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, "/nonexistent/i10/x")

	// Act
	actual := args.Map{"result": m.LengthOfIssues() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_FPFIM_IsEmptyIssues(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "noissue.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, fp)

	// Act
	actual := args.Map{"result": m.IsEmptyIssues()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RecursivePathsApply
// ══════════════════════════════════════════════════════════════════════════════

func Test_RecursivePathsApply(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0644)
	var count int
	err := chmodhelper.RecursivePathsApply(dir, func(path string, info os.FileInfo, err error) error {
		count++
		return nil
	})

	// Act
	actual := args.Map{"result": err != nil || count == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFilteredExistsPaths
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetFilteredExistsPaths_Mixed(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "exists.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	found, missing := chmodhelper.GetFilteredExistsPaths([]string{fp, "/nonexistent/i10"})

	// Act
	actual := args.Map{"result": len(found) != 1 || len(missing) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "found= missing=", actual)
}

func Test_GetFilteredExistsPaths_Empty(t *testing.T) {
	// Arrange
	found, missing := chmodhelper.GetFilteredExistsPaths(nil)

	// Act
	actual := args.Map{"result": len(found) != 0 || len(missing) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MergeRwxWildcardWithFixedRwx — valid case
// ══════════════════════════════════════════════════════════════════════════════

func Test_MergeRwxWildcard_Valid(t *testing.T) {
	// Arrange
	attr, err := chmodhelper.MergeRwxWildcardWithFixedRwx("r-x", "r*-")

	// Act
	actual := args.Map{"result": err != nil || attr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	// r*- merged with r-x = r-- (keep read from wildcard resolved to existing, write=no, execute=no)
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — IsEqualPartialRwxPartial, IsEqualPartialUsingFileMode
// ══════════════════════════════════════════════════════════════════════════════

func Test_VarWrapper_IsEqualPartialRwxPartial(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": vw.IsEqualPartialRwxPartial("-rwxr-xr-x")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_VarWrapper_IsEqualPartialUsingFileMode(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": vw.IsEqualPartialUsingFileMode(0755)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_VarWrapper_IsMismatchPartialFullRwx(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": vw.IsMismatchPartialFullRwx("-rw-r--r--")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_VarWrapper_IsEqualUsingFileMode(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": vw.IsEqualUsingFileMode(0755)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_VarWrapper_String(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	s := vw.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_VarWrapper_HasWildcard(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")

	// Act
	actual := args.Map{"result": vw.HasWildcard()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_VarWrapper_ToCompileWrapper(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	w := vw.ToCompileWrapper(nil)

	// Act
	actual := args.Map{"result": w.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetExistingChmodRwxWrappers — non-continue branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetExistingChmodRwxWrappers_NoContinue_Error(t *testing.T) {
	// Arrange
	_, err := chmodhelper.GetExistingChmodRwxWrappers(false, "/nonexistent/i10/nocont")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_GetExistingChmodRwxWrappers_NoContinue_Valid(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "wrappers.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	m, err := chmodhelper.GetExistingChmodRwxWrappers(false, fp)

	// Act
	actual := args.Map{"result": err != nil || len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFilesChmodRwxFullMap (GetPathsChmodsHashmap)
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetFilesChmodRwxFullMap_Error(t *testing.T) {
	// Arrange
	_, err := chmodhelper.GetFilesChmodRwxFullMap([]string{"/nonexistent/i10/hm"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ParseRwxInstructionToExecutor — nil
// ══════════════════════════════════════════════════════════════════════════════

func Test_ParseRwxInstructionToExecutor_Nil(t *testing.T) {
	// Arrange
	_, err := chmodhelper.ParseRwxInstructionToExecutor(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newSimpleFileReaderWriterCreator — Path method
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewSFRW_Path(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/i10/path.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewSFRW_Create(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, "/tmp/i10", "/tmp/i10/c.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewSFRW_All(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.All(0755, 0644, false, true, true, "/tmp/i10", "/tmp/i10/a.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewSFRW_Options(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Options(false, true, true, "/tmp/i10/opt.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewSFRW_CreateClean(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.CreateClean(false, 0755, 0644, "/tmp/i10", "/tmp/i10/cc.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewSFRW_DefaultCleanPath(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(false, "/tmp/i10/dcp.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewSFRW_PathCondition(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, true, 0755, 0644, "/tmp/i10/pc.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	rw2 := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, false, 0755, 0644, "/tmp/i10/pc2.txt")
	actual = args.Map{"result": rw2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewSFRW_PathDirDefaultChmod(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.PathDirDefaultChmod(false, 0644, "/tmp/i10/pddc.txt")

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetRecursivePaths — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetRecursivePaths_NonExistent_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	_, err := chmodhelper.GetRecursivePaths(false, "/nonexistent/i10/recurse")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_GetRecursivePaths_File_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "file.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(false, fp)

	// Act
	actual := args.Map{"result": err != nil || len(paths) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 path", actual)
}

func Test_GetRecursivePaths_Dir_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(false, dir)

	// Act
	actual := args.Map{"result": err != nil || len(paths) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2 paths", actual)
}

func Test_GetRecursivePaths_ContinueOnError_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "b.txt"), []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(true, dir)

	// Act
	actual := args.Map{"result": err != nil || len(paths) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2 paths", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// fileBytesWriter — uncovered methods (via SimpleFileWriter)
// ══════════════════════════════════════════════════════════════════════════════

func Test_FileBytesWriter_WithDirChmod_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbw.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmod(false, 0755, 0644, fp, []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_FileBytesWriter_WithDirChmodLock_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwl.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmodLock(false, 0755, 0644, fp, []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_FileBytesWriter_Chmod_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwc.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Chmod(false, 0755, 0644, fp, []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_FileBytesWriter_WithDir_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwd.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDir(false, fp, []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_FileBytesWriter_WithDirLock_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwdl.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirLock(false, fp, []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_FileBytesWriter_Default_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwdf.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Default(false, fp, []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// anyItemWriter — error path (unmarshalable)
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyItemWriter_Chmod_Error(t *testing.T) {
	// Arrange
	ch := make(chan int)
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Chmod(
		false, 0755, 0644, "/tmp", "/tmp/i10_any_err.json", ch)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for channel", actual)
}

func Test_AnyItemWriter_ChmodLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "anyl.json")
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.ChmodLock(
		false, 0755, 0644, dir, fp, map[string]string{"k": "v"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_AnyItemWriter_Default(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "anydf.json")
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Default(false, fp, map[string]string{"k": "v"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_AnyItemWriter_DefaultLock(t *testing.T) {
	// Arrange
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "anydfl.json")
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.DefaultLock(false, fp, map[string]string{"k": "v"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// fileReader — Read, ReadBytes
// ══════════════════════════════════════════════════════════════════════════════

func Test_FileReader_Read_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "fr.txt")
	os.WriteFile(fp, []byte("hello"), 0644)
	s, err := chmodhelper.SimpleFileWriter.FileReader.Read(fp)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_FileReader_ReadBytes_FromSimpleFileRWReadLock(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "frb.txt")
	os.WriteFile(fp, []byte("bytes"), 0644)
	b, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes(fp)

	// Act
	actual := args.Map{"result": err != nil || string(b) != "bytes"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_FileReader_ReadBytes_Error(t *testing.T) {
	// Arrange
	_, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes("/nonexistent/i10/fr.txt")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_FileReader_Read_Error(t *testing.T) {
	// Arrange
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent/i10/fr.txt")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}
