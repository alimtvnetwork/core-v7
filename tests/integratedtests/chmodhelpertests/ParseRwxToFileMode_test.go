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

func Test_ParseRwxOwnerGroupOtherToFileModeMust(t *testing.T) {
	// Arrange
	rwx := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r-x",
	}
	mode := chmodhelper.ParseRwxOwnerGroupOtherToFileModeMust(rwx)

	// Act
	actual := args.Map{"result": mode == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero mode", actual)
}

// ── ParseBaseRwxInstructionsToExecutors ──

func Test_ParseBaseRwxInstructionsToExecutors_Nil(t *testing.T) {
	// Arrange
	_, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ParseBaseRwxInstructionsToExecutors_Valid(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": executors == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ── GetFilesChmodRwxFullMap ──

func Test_GetFilesChmodRwxFullMap_Empty(t *testing.T) {
	// Arrange
	hm, err := chmodhelper.GetFilesChmodRwxFullMap(nil)

	// Act
	actual := args.Map{"result": err != nil || hm == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_GetFilesChmodRwxFullMap_Valid(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	f := filepath.Join(dir, "test.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	hm, err := chmodhelper.GetFilesChmodRwxFullMap([]string{f})

	// Act
	actual := args.Map{"result": err != nil || hm == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_GetFilesChmodRwxFullMap_Invalid(t *testing.T) {
	// Arrange
	hm, err := chmodhelper.GetFilesChmodRwxFullMap([]string{"/nonexistent/path/xyz123"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	_ = hm
}

// ── SimpleFileReaderWriter additional methods ──

func Test_SimpleFileReaderWriter_InitializeDefault_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	f := filepath.Join(dir, "init.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  f,
	}
	initialized := rw.InitializeDefault(true)

	// Act
	actual := args.Map{"result": initialized.ParentDir == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected parent dir", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefaultApplyChmod_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	f := filepath.Join(dir, "init2.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  f,
	}
	initialized := rw.InitializeDefaultApplyChmod()

	// Act
	actual := args.Map{"result": initialized == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SimpleFileReaderWriter_IsExistAndParent(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	f := filepath.Join(dir, "exist.txt")
	_ = os.WriteFile(f, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}

	// Act
	actual := args.Map{"result": rw.IsExist()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected exist", actual)
	actual = args.Map{"result": rw.IsParentExist()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected parent exist", actual)
	actual = args.Map{"result": rw.HasPathIssues()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
	actual = args.Map{"result": rw.IsPathInvalid()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual = args.Map{"result": rw.IsParentDirInvalid()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid parent", actual)
	actual = args.Map{"result": rw.HasAnyIssues()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
}

func Test_SimpleFileReaderWriter_WriteAndRead_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	f := filepath.Join(dir, "wr.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.Write([]byte("hello"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	content, err := rw.Read()
	actual = args.Map{"result": err != nil || string(content) != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_WriteString_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	f := filepath.Join(dir, "ws.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.WriteString("world")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	content, err := rw.ReadString()
	actual = args.Map{"result": err != nil || content != "world"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_ReadOnExist(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || bytes != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil nil", actual)
	content, err := rw.ReadStringOnExist()
	actual = args.Map{"result": err != nil || content != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_SimpleFileReaderWriter_WritePath(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	f := filepath.Join(dir, "wp.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  f,
	}
	err := rw.WritePath(false, f, []byte("test"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SimpleFileReaderWriter_WriteRelativePath(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "dummy.txt"),
	}
	err := rw.WriteRelativePath(false, "rel.txt", []byte("data"))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SimpleFileReaderWriter_JoinRelPath_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: "/tmp/base",
	}
	p := rw.JoinRelPath("sub/file.txt")

	// Act
	actual := args.Map{"result": p == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected path", actual)
	p2 := rw.JoinRelPath("")
	actual = args.Map{"result": p2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected path", actual)
}

func Test_SimpleFileReaderWriter_WriteAny(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SimpleFileReaderWriter_WriteAnyLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SimpleFileReaderWriter_ReadLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || string(b) != "locked"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_ReadStringLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || s != "locked"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_ReadOnExistLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || string(b) != "exists"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_ReadStringOnExistLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || s != "exists"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_String_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}
	s := rw.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_SimpleFileReaderWriter_StringFilePath(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}
	s := rw.StringFilePath("/other/path.txt")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_SimpleFileReaderWriter_ChmodApplier(t *testing.T) {
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

func Test_SimpleFileReaderWriter_ChmodVerifier(t *testing.T) {
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

func Test_SimpleFileReaderWriter_NewPath_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "orig.txt"),
	}
	newRw := rw.NewPath(false, filepath.Join(dir, "new.txt"))

	// Act
	actual := args.Map{"result": newRw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SimpleFileReaderWriter_NewPathJoin(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "orig.txt"),
	}
	newRw := rw.NewPathJoin(false, "sub", "file.txt")

	// Act
	actual := args.Map{"result": newRw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefaultNew(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filepath.Join(dir, "idn.txt"),
	}
	newRw := rw.InitializeDefaultNew()

	// Act
	actual := args.Map{"result": newRw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_SimpleFileReaderWriter_Set(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SimpleFileReaderWriter_SetLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SimpleFileReaderWriter_Get(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || result.X != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_GetLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || result.X != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_Expire_FromParseRwxToFileMode(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": rw.IsExist()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected removed", actual)
}

func Test_SimpleFileReaderWriter_Serialize(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_SerializeLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_Deserialize(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || result.X != 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_DeserializeLock(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err != nil || result.X != 20}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_ReadMust(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": string(b) != "must"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_SimpleFileReaderWriter_ReadStringMust(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": s != "strmust"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}
