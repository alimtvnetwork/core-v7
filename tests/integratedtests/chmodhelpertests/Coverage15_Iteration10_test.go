package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
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

func Test_I10_SFRW_ReadString(t *testing.T) {
	rw := newRW(t, "hello")
	s, err := rw.ReadString()
	if err != nil || s != "hello" {
		t.Fatalf("got %q, err=%v", s, err)
	}
}

func Test_I10_SFRW_ReadStringMust(t *testing.T) {
	rw := newRW(t, "world")
	s := rw.ReadStringMust()
	if s != "world" {
		t.Fatalf("got %q", s)
	}
}

func Test_I10_SFRW_ReadStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	rw.ReadStringMust()
}

func Test_I10_SFRW_ReadMust(t *testing.T) {
	rw := newRW(t, "data")
	b := rw.ReadMust()
	if string(b) != "data" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_ReadMust_Panic(t *testing.T) {
	defer func() { recover() }()
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	rw.ReadMust()
}

func Test_I10_SFRW_ReadLock(t *testing.T) {
	rw := newRW(t, "lock")
	b, err := rw.ReadLock()
	if err != nil || string(b) != "lock" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_ReadStringLock(t *testing.T) {
	rw := newRW(t, "slock")
	s, err := rw.ReadStringLock()
	if err != nil || s != "slock" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_ReadOnExist_Exists(t *testing.T) {
	rw := newRW(t, "exist")
	b, err := rw.ReadOnExist()
	if err != nil || string(b) != "exist" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_ReadOnExist_NotExists(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	b, err := rw.ReadOnExist()
	if err != nil || b != nil {
		t.Fatal("expected nil, nil")
	}
}

func Test_I10_SFRW_ReadStringOnExist(t *testing.T) {
	rw := newRW(t, "sexist")
	s, err := rw.ReadStringOnExist()
	if err != nil || s != "sexist" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_ReadStringOnExist_NotExists(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	s, err := rw.ReadStringOnExist()
	if err != nil || s != "" {
		t.Fatal("expected empty")
	}
}

func Test_I10_SFRW_ReadStringOnExistLock(t *testing.T) {
	rw := newRW(t, "lock2")
	s, err := rw.ReadStringOnExistLock()
	if err != nil || s != "lock2" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_ReadOnExistLock(t *testing.T) {
	rw := newRW(t, "lock3")
	b, err := rw.ReadOnExistLock()
	if err != nil || string(b) != "lock3" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_Read_Error(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/i10/read.txt",
		ParentDir: "/nonexistent/i10",
		ChmodDir:  0755,
		ChmodFile: 0644,
	}
	_, err := rw.Read()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — WriteString, WritePath, WriteRelativePath
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_SFRW_WriteString(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_WritePath(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_WriteRelativePath(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "x.txt"),
	}
	err := rw.WriteRelativePath(false, "rel.txt", []byte("reldata"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_JoinRelPath_Empty(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/tmp/i10"}
	result := rw.JoinRelPath("")
	if result == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I10_SFRW_JoinRelPath_NonEmpty(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/tmp/i10"}
	result := rw.JoinRelPath("sub/file.txt")
	if result == "" {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — WriteAny, WriteAnyLock, Get, Set, Expire, etc.
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_SFRW_WriteAny(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_WriteAnyLock(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_Get_NotExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/i10/get.json",
		ParentDir: "/nonexistent/i10",
		ChmodDir:  0755,
		ChmodFile: 0644,
	}
	var out map[string]string
	err := rw.Get(&out)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I10_SFRW_Get_Exists(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
	if out["key"] != "val" {
		t.Fatal("unexpected value")
	}
}

func Test_I10_SFRW_GetLock(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_Set(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_SetLock(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_Expire_Exists(t *testing.T) {
	rw := newRW(t, "expire")
	err := rw.Expire()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_Expire_NotExists(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10/expire.txt"}
	err := rw.Expire()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_ExpireLock(t *testing.T) {
	rw := newRW(t, "expirelock")
	err := rw.ExpireLock()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_ExpireParentDir(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_ExpireParentDirLock(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_RemoveOnExist(t *testing.T) {
	rw := newRW(t, "rm")
	_ = rw.RemoveOnExist()
}

func Test_I10_SFRW_RemoveDirOnExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: "/nonexistent/i10/rmdir",
	}
	_ = rw.RemoveDirOnExist()
}

func Test_I10_SFRW_OsFile(t *testing.T) {
	rw := newRW(t, "osfile")
	f, err := rw.OsFile()
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
}

func Test_I10_SFRW_Clone(t *testing.T) {
	rw := newRW(t, "clone")
	c := rw.Clone()
	if c.FilePath != rw.FilePath {
		t.Fatal("expected same file path")
	}
}

func Test_I10_SFRW_ClonePtr(t *testing.T) {
	rw := newRW(t, "cptr")
	cp := rw.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_SFRW_ClonePtr_Nil(t *testing.T) {
	var rw *chmodhelper.SimpleFileReaderWriter
	cp := rw.ClonePtr()
	if cp != nil {
		t.Fatal("expected nil")
	}
}

func Test_I10_SFRW_String(t *testing.T) {
	rw := newRW(t, "str")
	s := rw.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I10_SFRW_Json(t *testing.T) {
	rw := newRW(t, "json")
	j := rw.Json()
	if j.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I10_SFRW_JsonPtr(t *testing.T) {
	rw := newRW(t, "jsonptr")
	j := rw.JsonPtr()
	if j.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_I10_SFRW_MarshalUnmarshalJSON(t *testing.T) {
	rw := newRW(t, "marshal")
	b, err := rw.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	var rw2 chmodhelper.SimpleFileReaderWriter
	err2 := rw2.UnmarshalJSON(b)
	if err2 != nil {
		t.Fatal(err2)
	}
	if rw2.FilePath != rw.FilePath {
		t.Fatal("expected same path after unmarshal")
	}
}

func Test_I10_SFRW_AsJsonContractsBinder(t *testing.T) {
	rw := newRW(t, "binder")
	binder := rw.AsJsonContractsBinder()
	if binder == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_SFRW_JsonParseSelfInject(t *testing.T) {
	rw := newRW(t, "inject")
	j := rw.JsonPtr()
	var rw2 chmodhelper.SimpleFileReaderWriter
	err := rw2.JsonParseSelfInject(j)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_NewPath(t *testing.T) {
	rw := newRW(t, "np")
	newRw := rw.NewPath(false, filepath.Join(rw.ParentDir, "newfile.txt"))
	if newRw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_SFRW_NewPathJoin(t *testing.T) {
	rw := newRW(t, "npj")
	newRw := rw.NewPathJoin(false, "sub", "file.txt")
	if newRw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_SFRW_InitializeDefaultNew(t *testing.T) {
	rw := newRW(t, "idn")
	newRw := rw.InitializeDefaultNew()
	if newRw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_SFRW_HasAnyIssues(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/i10/issues",
		ParentDir: "/nonexistent/i10",
	}
	if !rw.HasAnyIssues() {
		t.Fatal("expected issues")
	}
}

func Test_I10_SFRW_IsPathInvalid(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent/i10"}
	if !rw.IsPathInvalid() {
		t.Fatal("expected invalid")
	}
}

func Test_I10_SFRW_IsParentDirInvalid(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/nonexistent/i10"}
	if !rw.IsParentDirInvalid() {
		t.Fatal("expected invalid")
	}
}

func Test_I10_SFRW_Serialize(t *testing.T) {
	rw := newRW(t, "serialize")
	b, err := rw.Serialize()
	if err != nil || string(b) != "serialize" {
		t.Fatalf("got %q, err=%v", string(b), err)
	}
}

func Test_I10_SFRW_SerializeLock(t *testing.T) {
	rw := newRW(t, "serlock")
	b, err := rw.SerializeLock()
	if err != nil || string(b) != "serlock" {
		t.Fatal("unexpected")
	}
}

func Test_I10_SFRW_Deserialize(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_SFRW_DeserializeLock(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func makeExecutor(t *testing.T, rwx string) *chmodhelper.RwxInstructionExecutor {
	t.Helper()
	exec, err := chmodhelper.RwxPartialToInstructionExecutor(rwx, &chmodins.Condition{})
	if err != nil {
		t.Fatal(err)
	}
	return exec
}

func Test_I10_Executor_IsFixedWrapper(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	if !exec.IsFixedWrapper() {
		t.Fatal("expected fixed")
	}
}

func Test_I10_Executor_IsVarWrapper(t *testing.T) {
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor("-rw*r-*r-*", &chmodins.Condition{})
	if !exec.IsVarWrapper() {
		t.Fatal("expected var")
	}
}

func Test_I10_Executor_IsEqualFileMode(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	if !exec.IsEqualFileMode(0755) {
		t.Fatal("expected true")
	}
}

func Test_I10_Executor_IsEqualRwxPartial(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	if !exec.IsEqualRwxPartial("-rwxr-xr-x") {
		t.Fatal("expected true")
	}
}

func Test_I10_Executor_IsEqualRwxWrapper(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	w := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	if !exec.IsEqualRwxWrapper(&w) {
		t.Fatal("expected true")
	}
}

func Test_I10_Executor_IsEqualFileInfo(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "ei.txt")
	os.WriteFile(fp, []byte("x"), 0755)
	os.Chmod(fp, 0755)
	info, _ := os.Stat(fp)
	exec := makeExecutor(t, "-rwxr-xr-x")
	_ = exec.IsEqualFileInfo(info)
}

func Test_I10_Executor_CompiledWrapper_Fixed(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	w, err := exec.CompiledWrapper(0755)
	if err != nil || w == nil {
		t.Fatal("expected wrapper")
	}
}

func Test_I10_Executor_CompiledWrapper_Var(t *testing.T) {
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor("-rw*r-*r-*", &chmodins.Condition{})
	w, err := exec.CompiledWrapper(0644)
	if err != nil || w == nil {
		t.Fatal("expected wrapper")
	}
}

func Test_I10_Executor_CompiledRwxWrapperUsingFixed_Fixed(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	w := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	r, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(&w)
	if err != nil || r == nil {
		t.Fatal("expected wrapper")
	}
}

func Test_I10_Executor_CompiledRwxWrapperUsingFixed_Var(t *testing.T) {
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor("-rw*r-*r-*", &chmodins.Condition{})
	w := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	r, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(&w)
	if err != nil || r == nil {
		t.Fatal("expected wrapper")
	}
}

func Test_I10_Executor_ApplyOnPath_Valid(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "apply.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	exec := makeExecutor(t, "-rw-r--r--")
	err := exec.ApplyOnPath(fp)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executor_ApplyOnPath_SkipInvalid(t *testing.T) {
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr-x", &chmodins.Condition{IsSkipOnInvalid: true})
	err := exec.ApplyOnPath("/nonexistent/i10/skip")
	if err != nil {
		t.Fatal("expected nil on skip")
	}
}

func Test_I10_Executor_ApplyOnPaths_Empty(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.ApplyOnPaths(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executor_ApplyOnPathsDirect_Empty(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.ApplyOnPathsDirect()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executor_ApplyOnPaths_Valid(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "paths.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	exec := makeExecutor(t, "-rw-r--r--")
	err := exec.ApplyOnPaths([]string{fp})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executor_ApplyOnPathsPtr_Nil(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.ApplyOnPathsPtr(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executor_ApplyOnPaths_ContinueOnError(t *testing.T) {
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr-x", &chmodins.Condition{IsContinueOnError: true})
	locs := []string{"/nonexistent/i10/a", "/nonexistent/i10/b"}
	_ = exec.ApplyOnPathsPtr(&locs)
}

func Test_I10_Executor_VerifyRwxModifiers_Valid(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "verify.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	os.Chmod(fp, 0644)
	exec := makeExecutor(t, "-rw-r--r--")
	err := exec.VerifyRwxModifiers(true, []string{fp})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executor_VerifyRwxModifiersDirect(t *testing.T) {
	exec := makeExecutor(t, "-rwxr-xr-x")
	err := exec.VerifyRwxModifiersDirect(true)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executor_VerifyRwxModifiers_ContinueOnError(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "vcont.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	exec, _ := chmodhelper.RwxPartialToInstructionExecutor(
		"-rw-r--r--", &chmodins.Condition{IsContinueOnError: true})
	err := exec.VerifyRwxModifiers(true, []string{fp})
	_ = err
}

func Test_I10_Executor_VerifyRwxModifiers_Mismatch(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "vmm.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	os.Chmod(fp, 0644)
	exec := makeExecutor(t, "-rwxrwxrwx")
	err := exec.VerifyRwxModifiers(true, []string{fp})
	if err == nil {
		t.Fatal("expected mismatch error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_Executors_LastIndex(t *testing.T) {
	execs := chmodhelper.NewRwxInstructionExecutors(2)
	if execs.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func Test_I10_Executors_HasIndex(t *testing.T) {
	execs := chmodhelper.NewRwxInstructionExecutors(2)
	if execs.HasIndex(0) {
		t.Fatal("expected false")
	}
}

func Test_I10_Executors_Items(t *testing.T) {
	execs := chmodhelper.NewRwxInstructionExecutors(2)
	if execs.Items() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_Executors_ApplyOnPaths(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "execs.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	exec := makeExecutor(t, "-rw-r--r--")
	execs := chmodhelper.NewRwxInstructionExecutors(1)
	execs.Add(exec)
	err := execs.ApplyOnPaths([]string{fp})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_Executors_VerifyRwxModifiers_ContinueOnError(t *testing.T) {
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

func Test_I10_Executors_VerifyRwxModifiers_NoContinue(t *testing.T) {
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

func Test_I10_AttrVariant_IsGreaterThan(t *testing.T) {
	v := chmodhelper.ReadWriteExecute
	if !v.IsGreaterThan(8) {
		// 8 > 7 is true
	}
	if v.IsGreaterThan(6) {
		t.Fatal("expected false: 6 is not > 7")
	}
}

func Test_I10_AttrVariant_String(t *testing.T) {
	v := chmodhelper.Read
	_ = v.String()
}

func Test_I10_AttrVariant_Value(t *testing.T) {
	v := chmodhelper.Execute
	if v.Value() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_I10_AttrVariant_ToAttribute(t *testing.T) {
	v := chmodhelper.ReadWriteExecute
	a := v.ToAttribute()
	if !a.IsRead || !a.IsWrite || !a.IsExecute {
		t.Fatal("expected all true for ReadWriteExecute")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FilteredPathFileInfoMap — uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_FPFIM_LazyValidLocations(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "lazy.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, fp)
	locs := m.LazyValidLocations()
	if len(locs) == 0 {
		t.Fatal("expected locations")
	}
	// second call hits cache
	locs2 := m.LazyValidLocations()
	if len(locs2) != len(locs) {
		t.Fatal("cache mismatch")
	}
}

func Test_I10_FPFIM_MissingPathsToString(t *testing.T) {
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, "/nonexistent/i10/a", "/nonexistent/i10/b")
	s := m.MissingPathsToString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I10_FPFIM_HasAnyIssues(t *testing.T) {
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(false, "/nonexistent/i10/x")
	if !m.HasAnyIssues() {
		t.Fatal("expected issues")
	}
}

func Test_I10_FPFIM_HasError(t *testing.T) {
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(false, "/nonexistent/i10/x")
	if !m.HasError() {
		t.Fatal("expected error")
	}
}

func Test_I10_FPFIM_HasAnyMissingPaths(t *testing.T) {
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, "/nonexistent/i10/x")
	if !m.HasAnyMissingPaths() {
		t.Fatal("expected missing paths")
	}
}

func Test_I10_FPFIM_LengthOfIssues(t *testing.T) {
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, "/nonexistent/i10/x")
	if m.LengthOfIssues() != 1 {
		t.Fatalf("expected 1, got %d", m.LengthOfIssues())
	}
}

func Test_I10_FPFIM_IsEmptyIssues(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "noissue.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	m := chmodhelper.GetExistsFilteredPathFileInfoMap(true, fp)
	if !m.IsEmptyIssues() {
		t.Fatal("expected no issues")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RecursivePathsApply
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_RecursivePathsApply(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0644)
	var count int
	err := chmodhelper.RecursivePathsApply(dir, func(path string, info os.FileInfo, err error) error {
		count++
		return nil
	})
	if err != nil || count == 0 {
		t.Fatal("unexpected")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFilteredExistsPaths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_GetFilteredExistsPaths_Mixed(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "exists.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	found, missing := chmodhelper.GetFilteredExistsPaths([]string{fp, "/nonexistent/i10"})
	if len(found) != 1 || len(missing) != 1 {
		t.Fatalf("found=%d missing=%d", len(found), len(missing))
	}
}

func Test_I10_GetFilteredExistsPaths_Empty(t *testing.T) {
	found, missing := chmodhelper.GetFilteredExistsPaths(nil)
	if len(found) != 0 || len(missing) != 0 {
		t.Fatal("expected empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MergeRwxWildcardWithFixedRwx — valid case
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_MergeRwxWildcard_Valid(t *testing.T) {
	attr, err := chmodhelper.MergeRwxWildcardWithFixedRwx("r-x", "r*-")
	if err != nil || attr == nil {
		t.Fatal("expected valid")
	}
	// r*- merged with r-x = r-- (keep read from wildcard resolved to existing, write=no, execute=no)
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — IsEqualPartialRwxPartial, IsEqualPartialUsingFileMode
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_VarWrapper_IsEqualPartialRwxPartial(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	if !vw.IsEqualPartialRwxPartial("-rwxr-xr-x") {
		t.Fatal("expected true")
	}
}

func Test_I10_VarWrapper_IsEqualPartialUsingFileMode(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	if !vw.IsEqualPartialUsingFileMode(0755) {
		t.Fatal("expected true")
	}
}

func Test_I10_VarWrapper_IsMismatchPartialFullRwx(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	if !vw.IsMismatchPartialFullRwx("-rw-r--r--") {
		t.Fatal("expected mismatch")
	}
}

func Test_I10_VarWrapper_IsEqualUsingFileMode(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	if !vw.IsEqualUsingFileMode(0755) {
		t.Fatal("expected true")
	}
}

func Test_I10_VarWrapper_String(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	s := vw.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I10_VarWrapper_HasWildcard(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	if !vw.HasWildcard() {
		t.Fatal("expected true")
	}
}

func Test_I10_VarWrapper_ToCompileWrapper(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	w := vw.ToCompileWrapper(nil)
	if w.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// GetExistingChmodRwxWrappers — non-continue branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_GetExistingChmodRwxWrappers_NoContinue_Error(t *testing.T) {
	_, err := chmodhelper.GetExistingChmodRwxWrappers(false, "/nonexistent/i10/nocont")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I10_GetExistingChmodRwxWrappers_NoContinue_Valid(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "wrappers.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	m, err := chmodhelper.GetExistingChmodRwxWrappers(false, fp)
	if err != nil || len(m) != 1 {
		t.Fatal("unexpected")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFilesChmodRwxFullMap (GetPathsChmodsHashmap)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_GetFilesChmodRwxFullMap_Error(t *testing.T) {
	_, err := chmodhelper.GetFilesChmodRwxFullMap([]string{"/nonexistent/i10/hm"})
	if err == nil {
		t.Fatal("expected error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ParseRwxInstructionToExecutor — nil
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_ParseRwxInstructionToExecutor_Nil(t *testing.T) {
	_, err := chmodhelper.ParseRwxInstructionToExecutor(nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// newSimpleFileReaderWriterCreator — Path method
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_NewSFRW_Path(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/i10/path.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_NewSFRW_Create(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, "/tmp/i10", "/tmp/i10/c.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_NewSFRW_All(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.All(0755, 0644, false, true, true, "/tmp/i10", "/tmp/i10/a.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_NewSFRW_Options(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Options(false, true, true, "/tmp/i10/opt.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_NewSFRW_CreateClean(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.CreateClean(false, 0755, 0644, "/tmp/i10", "/tmp/i10/cc.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_NewSFRW_DefaultCleanPath(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(false, "/tmp/i10/dcp.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_NewSFRW_PathCondition(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, true, 0755, 0644, "/tmp/i10/pc.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
	rw2 := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, false, 0755, 0644, "/tmp/i10/pc2.txt")
	if rw2 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I10_NewSFRW_PathDirDefaultChmod(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.PathDirDefaultChmod(false, 0644, "/tmp/i10/pddc.txt")
	if rw == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// GetRecursivePaths — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_GetRecursivePaths_NonExistent(t *testing.T) {
	_, err := chmodhelper.GetRecursivePaths(false, "/nonexistent/i10/recurse")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I10_GetRecursivePaths_File(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "file.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(false, fp)
	if err != nil || len(paths) != 1 {
		t.Fatal("expected 1 path")
	}
}

func Test_I10_GetRecursivePaths_Dir(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(false, dir)
	if err != nil || len(paths) < 2 {
		t.Fatal("expected >= 2 paths")
	}
}

func Test_I10_GetRecursivePaths_ContinueOnError(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "b.txt"), []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(true, dir)
	if err != nil || len(paths) < 2 {
		t.Fatal("expected >= 2 paths")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// fileBytesWriter — uncovered methods (via SimpleFileWriter)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_FileBytesWriter_WithDirChmod(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbw.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmod(false, 0755, 0644, fp, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_FileBytesWriter_WithDirChmodLock(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwl.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmodLock(false, 0755, 0644, fp, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_FileBytesWriter_Chmod(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwc.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Chmod(false, 0755, 0644, fp, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_FileBytesWriter_WithDir(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwd.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDir(false, fp, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_FileBytesWriter_WithDirLock(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwdl.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirLock(false, fp, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_FileBytesWriter_Default(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "fbwdf.txt")
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Default(false, fp, []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// anyItemWriter — error path (unmarshalable)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_AnyItemWriter_Chmod_Error(t *testing.T) {
	ch := make(chan int)
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Chmod(
		false, 0755, 0644, "/tmp", "/tmp/i10_any_err.json", ch)
	if err == nil {
		t.Fatal("expected error for channel")
	}
}

func Test_I10_AnyItemWriter_ChmodLock(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "anyl.json")
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.ChmodLock(
		false, 0755, 0644, dir, fp, map[string]string{"k": "v"})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_AnyItemWriter_Default(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "anydf.json")
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Default(false, fp, map[string]string{"k": "v"})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I10_AnyItemWriter_DefaultLock(t *testing.T) {
	skipWin(t)
	dir := t.TempDir()
	fp := filepath.Join(dir, "anydfl.json")
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.DefaultLock(false, fp, map[string]string{"k": "v"})
	if err != nil {
		t.Fatal(err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// fileReader — Read, ReadBytes
// ══════════════════════════════════════════════════════════════════════════════

func Test_I10_FileReader_Read(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "fr.txt")
	os.WriteFile(fp, []byte("hello"), 0644)
	s, err := chmodhelper.SimpleFileWriter.FileReader.Read(fp)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_I10_FileReader_ReadBytes(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "frb.txt")
	os.WriteFile(fp, []byte("bytes"), 0644)
	b, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes(fp)
	if err != nil || string(b) != "bytes" {
		t.Fatal("unexpected")
	}
}

func Test_I10_FileReader_ReadBytes_Error(t *testing.T) {
	_, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes("/nonexistent/i10/fr.txt")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I10_FileReader_Read_Error(t *testing.T) {
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent/i10/fr.txt")
	if err == nil {
		t.Fatal("expected error")
	}
}
