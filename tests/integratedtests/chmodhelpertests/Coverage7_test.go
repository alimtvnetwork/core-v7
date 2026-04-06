package chmodhelpertests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ParseRwxOwnerGroupOtherToFileModeMust ──

func Test_Cov7_ParseRwxOwnerGroupOtherToFileModeMust(t *testing.T) {
	rwx := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r-x",
	}
	mode := chmodhelper.ParseRwxOwnerGroupOtherToFileModeMust(rwx)
	actual := args.Map{"result": mode == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero mode", actual)
}

// ── ParseBaseRwxInstructionsToExecutors ──

func Test_Cov7_ParseBaseRwxInstructionsToExecutors_Nil(t *testing.T) {
	_, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Cov7_ParseBaseRwxInstructionsToExecutors_Valid(t *testing.T) {
	base := &chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "rwx",
					Group: "r-x",
					Other: "r-x",
				},
				Condition: chmodins.Condition{},
			},
		},
	}
	executors, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(base)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": executors == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ── GetFilesChmodRwxFullMap ──

func Test_Cov7_GetFilesChmodRwxFullMap_Empty(t *testing.T) {
	hm, err := chmodhelper.GetFilesChmodRwxFullMap(nil)
	actual := args.Map{"result": err != nil || hm == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_GetFilesChmodRwxFullMap_Valid(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "test.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	hm, err := chmodhelper.GetFilesChmodRwxFullMap([]string{f})
	actual := args.Map{"result": err != nil || hm == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_GetFilesChmodRwxFullMap_Invalid(t *testing.T) {
	hm, err := chmodhelper.GetFilesChmodRwxFullMap([]string{"/nonexistent/path/xyz123"})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	_ = hm
}

// ── SimpleFileReaderWriter additional methods ──

func Test_Cov7_SimpleFileReaderWriter_InitializeDefault(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "init.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  f,
	}
	initialized := rw.InitializeDefault(true)
	actual := args.Map{"result": initialized.ParentDir == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected parent dir", actual)
}

func Test_Cov7_SimpleFileReaderWriter_InitializeDefaultApplyChmod(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "init2.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  f,
	}
	initialized := rw.InitializeDefaultApplyChmod()
	actual := args.Map{"result": initialized == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov7_SimpleFileReaderWriter_IsExistAndParent(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "exist.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	actual := args.Map{"result": rw.IsExist()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected exist", actual)
	actual := args.Map{"result": rw.IsParentExist()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected parent exist", actual)
	actual := args.Map{"result": rw.HasPathIssues()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
	actual := args.Map{"result": rw.IsPathInvalid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual := args.Map{"result": rw.IsParentDirInvalid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid parent", actual)
	actual := args.Map{"result": rw.HasAnyIssues()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
}

func Test_Cov7_SimpleFileReaderWriter_WriteAndRead(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "wr.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.Write([]byte("hello"))
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	content, err := rw.Read()
	actual := args.Map{"result": err != nil || string(content) != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_WriteString(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "ws.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.WriteString("world")
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	content, err := rw.ReadString()
	actual := args.Map{"result": err != nil || content != "world"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ReadOnExist(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "roe.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	// File doesn't exist yet
	bytes, err := rw.ReadOnExist()
	actual := args.Map{"result": err != nil || bytes != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil nil", actual)
	content, err := rw.ReadStringOnExist()
	actual := args.Map{"result": err != nil || content != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Cov7_SimpleFileReaderWriter_WritePath(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "wp.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.WritePath(false, f, []byte("test"))
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Cov7_SimpleFileReaderWriter_WriteRelativePath(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "dummy.txt"),
	}
	err := rw.WriteRelativePath(false, "rel.txt", []byte("data"))
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Cov7_SimpleFileReaderWriter_JoinRelPath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: "/tmp/base",
	}
	p := rw.JoinRelPath("sub/file.txt")
	actual := args.Map{"result": p == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected path", actual)
	p2 := rw.JoinRelPath("")
	actual := args.Map{"result": p2 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected path", actual)
}

func Test_Cov7_SimpleFileReaderWriter_WriteAny(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "any.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ Name string }
	err := rw.WriteAny(&data{Name: "test"})
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Cov7_SimpleFileReaderWriter_WriteAnyLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "anylock.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ Val int }
	err := rw.WriteAnyLock(&data{Val: 42})
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ReadLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "rl.txt")
	_ = os.WriteFile(f, []byte("locked"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.ReadLock()
	actual := args.Map{"result": err != nil || string(b) != "locked"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ReadStringLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "rsl.txt")
	_ = os.WriteFile(f, []byte("locked"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	s, err := rw.ReadStringLock()
	actual := args.Map{"result": err != nil || s != "locked"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ReadOnExistLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "roel.txt")
	_ = os.WriteFile(f, []byte("exists"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.ReadOnExistLock()
	actual := args.Map{"result": err != nil || string(b) != "exists"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ReadStringOnExistLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "rsoel.txt")
	_ = os.WriteFile(f, []byte("exists"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	s, err := rw.ReadStringOnExistLock()
	actual := args.Map{"result": err != nil || s != "exists"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_String(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}
	s := rw.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_Cov7_SimpleFileReaderWriter_StringFilePath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}
	s := rw.StringFilePath("/other/path.txt")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ChmodApplier(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "ca.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	applier := rw.ChmodApplier()
	_ = applier
}

func Test_Cov7_SimpleFileReaderWriter_ChmodVerifier(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "cv.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	verifier := rw.ChmodVerifier()
	_ = verifier
}

func Test_Cov7_SimpleFileReaderWriter_NewPath(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "orig.txt"),
	}
	newRw := rw.NewPath(false, filepath.Join(dir, "new.txt"))
	actual := args.Map{"result": newRw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov7_SimpleFileReaderWriter_NewPathJoin(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "orig.txt"),
	}
	newRw := rw.NewPathJoin(false, "sub", "file.txt")
	actual := args.Map{"result": newRw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov7_SimpleFileReaderWriter_InitializeDefaultNew(t *testing.T) {
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "idn.txt"),
	}
	newRw := rw.InitializeDefaultNew()
	actual := args.Map{"result": newRw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov7_SimpleFileReaderWriter_Set(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "set.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	err := rw.Set(&data{X: 1})
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Cov7_SimpleFileReaderWriter_SetLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "setlock.json")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	err := rw.SetLock(&data{X: 2})
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Cov7_SimpleFileReaderWriter_Get(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "get.json")
	_ = os.WriteFile(f, []byte(`{"X":42}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.Get(result)
	actual := args.Map{"result": err != nil || result.X != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_GetLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "getlock.json")
	_ = os.WriteFile(f, []byte(`{"X":99}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.GetLock(result)
	actual := args.Map{"result": err != nil || result.X != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_Expire(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "expire.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	rw.Expire()
	actual := args.Map{"result": rw.IsExist()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected removed", actual)
}

func Test_Cov7_SimpleFileReaderWriter_Serialize(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "ser.txt")
	_ = os.WriteFile(f, []byte("data"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.Serialize()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_SerializeLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "serlock.txt")
	_ = os.WriteFile(f, []byte("data"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b, err := rw.SerializeLock()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_Deserialize(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "deser.json")
	_ = os.WriteFile(f, []byte(`{"X":10}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.Deserialize(result)
	actual := args.Map{"result": err != nil || result.X != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_DeserializeLock(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "deserlock.json")
	_ = os.WriteFile(f, []byte(`{"X":20}`), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	type data struct{ X int }
	result := &data{}
	err := rw.DeserializeLock(result)
	actual := args.Map{"result": err != nil || result.X != 20}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ReadMust(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "readmust.txt")
	_ = os.WriteFile(f, []byte("must"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	b := rw.ReadMust()
	actual := args.Map{"result": string(b) != "must"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Cov7_SimpleFileReaderWriter_ReadStringMust(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "readstrmust.txt")
	_ = os.WriteFile(f, []byte("strmust"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	s := rw.ReadStringMust()
	actual := args.Map{"result": s != "strmust"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}
