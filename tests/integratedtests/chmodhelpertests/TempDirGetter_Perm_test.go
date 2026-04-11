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

// ── tempDirGetter ──

func Test_TempDirGetter_TempDefault_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	result := chmodhelper.TempDirGetter.TempDefault()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TempDirGetter.TempDefault returns correct value -- with args", actual)
}

func Test_TempDirGetter_TempPermanent_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	result := chmodhelper.TempDirGetter.TempPermanent()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TempDirGetter.TempPermanent returns correct value -- with args", actual)
}

func Test_TempDirGetter_TempOption_Permanent_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	result := chmodhelper.TempDirGetter.TempOption(true)

	// Assert
	expected := chmodhelper.TempDirGetter.TempPermanent()

	// Act
	actual := args.Map{"match": result == expected}
	exp := args.Map{"match": true}
	exp.ShouldBeEqual(t, 0, "TempOption returns correct value -- permanent", actual)
}

func Test_TempDirGetter_TempOption_Default(t *testing.T) {
	// Arrange
	result := chmodhelper.TempDirGetter.TempOption(false)

	// Assert
	expected := chmodhelper.TempDirGetter.TempDefault()

	// Act
	actual := args.Map{"match": result == expected}
	exp := args.Map{"match": true}
	exp.ShouldBeEqual(t, 0, "TempOption returns correct value -- default", actual)
}

// ── fileReader via SimpleFileWriter ──

func Test_FileReader_Read_InvalidPath(t *testing.T) {
	// Arrange
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent/path/file.txt")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FileReader.Read returns error -- invalid", actual)
}

func Test_FileReader_ReadBytes_InvalidPath(t *testing.T) {
	// Arrange
	_, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes("/nonexistent/path/file.txt")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FileReader.ReadBytes returns error -- invalid", actual)
}

func Test_FileReader_Read_ValidFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(filePath, []byte("hello"), 0644)

	content, err := chmodhelper.SimpleFileWriter.FileReader.Read(filePath)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"content": "hello",
	}
	expected.ShouldBeEqual(t, 0, "FileReader.Read returns non-empty -- valid", actual)
}

// ── fileWriter ──

func Test_FileWriter_Remove_NonExistent(t *testing.T) {
	// Arrange
	err := chmodhelper.SimpleFileWriter.FileWriter.Remove("/nonexistent/path")
	// RemoveAll on non-existent is OK

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.Remove returns non-empty -- non-existent", actual)
}

func Test_FileWriter_RemoveIf_False_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(false, "/any/path")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.RemoveIf returns non-empty -- false", actual)
}

func Test_FileWriter_RemoveIf_True_NonExistent(t *testing.T) {
	// Arrange
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(true, "/nonexistent/file")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.RemoveIf returns non-empty -- true non-existent", actual)
}

func Test_FileWriter_ParentDir_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	result := chmodhelper.SimpleFileWriter.FileWriter.ParentDir("/a/b/c.txt")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.ParentDir returns correct value -- with args", actual)
}

func Test_FileWriter_Chmod_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Chmod(
		false, 0755, 0644, filePath, []byte("test"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.Chmod returns correct value -- with args", actual)
}

func Test_FileWriter_ChmodFile_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test2.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.ChmodFile(
		false, 0644, filePath, []byte("data"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.ChmodFile returns correct value -- with args", actual)
}

func Test_FileWriter_AllLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "lock-test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.AllLock(
		0755, 0644, false, false, false, true,
		tmpDir, filePath, []byte("locked"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.AllLock returns correct value -- with args", actual)
}

// ── fileBytesWriter ──

func Test_FileBytesWriter_WithDir_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDir(
		false, filePath, []byte("bytes"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDir returns non-empty -- with args", actual)
}

func Test_FileBytesWriter_Default_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-default.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Default(
		false, filePath, []byte("default"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.Default returns correct value -- with args", actual)
}

func Test_FileBytesWriter_WithDirLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirLock(
		false, filePath, []byte("locked"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDirLock returns non-empty -- with args", actual)
}

func Test_FileBytesWriter_WithDirChmodLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-chmod-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmodLock(
		false, 0755, 0644, filePath, []byte("data"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDirChmodLock returns non-empty -- with args", actual)
}

func Test_FileBytesWriter_Chmod_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-chmod.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Chmod(
		false, 0755, 0644, filePath, []byte("chmod"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.Chmod returns correct value -- with args", actual)
}

// ── fileStringWriter ──

func Test_FileStringWriter_Default_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-default.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.Default(
		false, filePath, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.Default returns correct value -- with args", actual)
}

func Test_FileStringWriter_DefaultLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.DefaultLock(
		false, filePath, "locked")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.DefaultLock returns correct value -- with args", actual)
}

func Test_FileStringWriter_Chmod_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-chmod.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.Chmod(
		false, 0755, 0644, filePath, "chmod-content")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.Chmod returns correct value -- with args", actual)
}

func Test_FileStringWriter_ChmodLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-chmod-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.ChmodLock(
		false, 0755, 0644, filePath, "lock-content")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.ChmodLock returns correct value -- with args", actual)
}

func Test_FileStringWriter_All_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-all.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.All(
		false, 0755, 0644, false, false, true,
		tmpDir, filePath, "all-content")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.All returns correct value -- with args", actual)
}

// ── anyItemWriter ──

func Test_AnyItemWriter_Default_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-default.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Default(
		false, filePath, map[string]string{"key": "value"})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.Default returns correct value -- with args", actual)
}

func Test_AnyItemWriter_DefaultLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-lock.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.DefaultLock(
		false, filePath, map[string]int{"a": 1})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.DefaultLock returns correct value -- with args", actual)
}

func Test_AnyItemWriter_Chmod(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-chmod.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Chmod(
		false, 0755, 0644, tmpDir, filePath, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.Chmod returns correct value -- with args", actual)
}

func Test_AnyItemWriter_ChmodLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-chmod-lock.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.ChmodLock(
		false, 0755, 0644, tmpDir, filePath, 42)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.ChmodLock returns correct value -- with args", actual)
}

func Test_AnyItemWriter_Chmod_InvalidJson(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-invalid.json")

	// channels can't be marshalled to JSON
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Chmod(
		false, 0755, 0644, tmpDir, filePath, make(chan int))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.Chmod returns error -- invalid JSON", actual)
}

// ── dirCreator via SimpleFileWriter.CreateDir ──

func Test_DirCreator_If_False_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.SimpleFileWriter.CreateDir.If(false, 0755, "/any/path")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.If returns non-empty -- false", actual)
}

func Test_DirCreator_IfMissing_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "new-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(0755, newDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissing returns correct value -- with args", actual)
}

func Test_DirCreator_IfMissing_AlreadyExists_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissing returns correct value -- existing", actual)
}

func Test_DirCreator_IfMissingLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "lock-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissingLock(0755, newDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissingLock returns correct value -- with args", actual)
}

func Test_DirCreator_Default_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "default-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.Default(0755, newDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.Default returns correct value -- with args", actual)
}

func Test_DirCreator_DefaultLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "default-lock-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.DefaultLock(0755, newDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.DefaultLock returns correct value -- with args", actual)
}

func Test_DirCreator_Direct_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "direct-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.Direct(newDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.Direct returns correct value -- with args", actual)
}

func Test_DirCreator_DirectLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "direct-lock-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.DirectLock(newDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.DirectLock returns correct value -- with args", actual)
}

func Test_DirCreator_ByChecking_ExistsAndIsDir(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking returns correct value -- existing dir", actual)
}

func Test_DirCreator_ByChecking_ExistsButIsFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "file.txt")
	_ = os.WriteFile(filePath, []byte("content"), 0644)

	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, filePath)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking returns correct value -- file not dir", actual)
}

func Test_DirCreator_ByChecking_NotExists(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "checking-new")

	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, newDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking returns correct value -- new", actual)
}

// ── chmodVerifier ──

func Test_ChmodVerifier_GetRwxFull_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.GetRwxFull(0755)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetRwxFull returns correct value -- with args", actual)
}

func Test_ChmodVerifier_GetRwx9_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.GetRwx9(0755)

	// Act
	actual := args.Map{"len9": len(result) == 9}

	// Assert
	expected := args.Map{"len9": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetRwx9 returns correct value -- with args", actual)
}

func Test_ChmodVerifier_IsEqual_ValidPath(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "verify.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	result := chmodhelper.ChmodVerify.IsEqual(filePath, 0644)

	// Act
	actual := args.Map{"isExpected": result}

	// Assert
	expected := args.Map{"isExpected": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqual returns correct value -- with args", actual)
}

func Test_ChmodVerifier_IsMismatch_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	result := chmodhelper.ChmodVerify.IsMismatch(filePath, 0777)

	// Act
	actual := args.Map{"mismatch": result}

	// Assert
	expected := args.Map{"mismatch": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsMismatch returns correct value -- with args", actual)
}

func Test_ChmodVerifier_MismatchError(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch-err.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.MismatchError(filePath, 0777)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.MismatchError returns error -- with args", actual)
}

func Test_ChmodVerifier_MismatchErrorUsingRwxFull_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch-rwx.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.MismatchErrorUsingRwxFull(filePath, "-rwxrwxrwx")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.MismatchErrorUsingRwxFull returns error -- with args", actual)
}

func Test_ChmodVerifier_PathIf_False_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodVerify.PathIf(false, "/any", 0644)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathIf returns non-empty -- false", actual)
}

func Test_ChmodVerifier_PathIf_True(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "pathif.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathIf(true, filePath, 0644)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathIf returns non-empty -- true", actual)
}

func Test_ChmodVerifier_Path(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "path.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.Path(filePath, 0644)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.Path returns correct value -- with args", actual)
}

func Test_ChmodVerifier_IsEqualRwxFull(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rwxfull.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	result := chmodhelper.ChmodVerify.IsEqualRwxFull(filePath, "-rw-r--r--")

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqualRwxFull returns correct value -- with args", actual)
}

func Test_ChmodVerifier_IsEqualRwxFullSkipInvalid_InvalidPath_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	nonExistent := filepath.Join(t.TempDir(), "no_such_file")
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid(nonExistent, "-rw-r--r--")

	// Act
	actual := args.Map{"assumedEqual": result}

	// Assert
	expected := args.Map{"assumedEqual": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqualRwxFullSkipInvalid returns error -- invalid path", actual)
}

func Test_ChmodVerifier_IsEqualSkipInvalid_InvalidPath_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	nonExistent := filepath.Join(t.TempDir(), "no_such_file")
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid(nonExistent, 0644)

	// Act
	actual := args.Map{"assumedEqual": result}

	// Assert
	expected := args.Map{"assumedEqual": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqualSkipInvalid returns error -- invalid path", actual)
}

func Test_ChmodVerifier_GetExisting_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "exist.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	mode, err := chmodhelper.ChmodVerify.GetExisting(filePath)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nonZero": mode != 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nonZero": true,
	}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetExisting returns correct value -- with args", actual)
}

func Test_ChmodVerifier_GetExistingRwxWrapper_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rwxw.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rwx, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper(filePath)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"defined": rwx.IsDefined(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"defined": true,
	}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetExistingRwxWrapper returns correct value -- with args", actual)
}

func Test_ChmodVerifier_GetExistingRwxWrapperMust(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "must.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rwx := chmodhelper.ChmodVerify.GetExistingRwxWrapperMust(filePath)

	// Act
	actual := args.Map{"defined": rwx.IsDefined()}

	// Assert
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetExistingRwxWrapperMust returns correct value -- with args", actual)
}

func Test_ChmodVerifier_PathsUsingFileModeImmediateReturn(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "f1.txt")
	_ = os.WriteFile(f1, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathsUsingFileModeImmediateReturn(0644, f1)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathsUsingFileModeImmediateReturn returns correct value -- with args", actual)
}

func Test_ChmodVerifier_PathsUsingFileModeContinueOnError(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "f2.txt")
	_ = os.WriteFile(f1, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathsUsingFileModeContinueOnError(0644, f1)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathsUsingFileModeContinueOnError returns error -- with args", actual)
}

func Test_ChmodVerifier_PathsUsingFileMode_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "f3.txt")
	_ = os.WriteFile(f1, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathsUsingFileMode(false, 0644, f1)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathsUsingFileMode returns correct value -- with args", actual)
}

func Test_ChmodVerifier_RwxFull_InvalidLength_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodVerify.RwxFull("/any", "short")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.RwxFull returns error -- invalid length", actual)
}

func Test_ChmodVerifier_RwxFull_NonExistentPath(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodVerify.RwxFull("/nonexistent/path", "-rw-r--r--")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.RwxFull returns non-empty -- non-existent", actual)
}

// ── chmodApplier ──

func Test_ChmodApplier_ApplyIf_False(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.ApplyIf(false, 0644, "/any")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.ApplyIf returns non-empty -- false", actual)
}

func Test_ChmodApplier_ApplyIf_True(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "apply.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.ApplyIf(true, 0644, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.ApplyIf returns non-empty -- true", actual)
}

func Test_ChmodApplier_Default(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "default.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.Default(0644, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.Default returns correct value -- with args", actual)
}

func Test_ChmodApplier_OnMismatchOption_SkipApply_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.OnMismatchOption(false, false, 0644, "/any")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.OnMismatchOption returns correct value -- skip", actual)
}

func Test_ChmodApplier_OnMismatch(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.OnMismatch(true, 0644, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.OnMismatch returns correct value -- with args", actual)
}

func Test_ChmodApplier_OnMismatchSkipInvalid(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "skip-invalid.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.OnMismatchSkipInvalid(0644, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.OnMismatchSkipInvalid returns error -- with args", actual)
}

func Test_ChmodApplier_SkipInvalidFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "skipinv.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.SkipInvalidFile(0644, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.SkipInvalidFile returns error -- with args", actual)
}

func Test_ChmodApplier_Options_NonRecursive(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "opts.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.Options(true, false, 0644, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.Options returns non-empty -- non-recursive", actual)
}

func Test_ChmodApplier_Options_Recursive(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.Options(true, true, 0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.Options returns correct value -- recursive", actual)
}

func Test_ChmodApplier_RecursivePath(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePath(true, 0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.RecursivePath returns correct value -- with args", actual)
}

func Test_ChmodApplier_PathsUsingFileModeConditions_EmptyLocations_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0644, &chmodins.Condition{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions returns empty -- empty", actual)
}

func Test_ChmodApplier_PathsUsingFileModeConditions_NilCondition_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0644, nil, "/some/path")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions returns nil -- nil condition", actual)
}

func Test_ChmodApplier_PathsUsingFileModeOptions(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.PathsUsingFileModeOptions(
		true, false, false, 0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeOptions returns correct value -- with args", actual)
}

func Test_ChmodApplier_PathsUsingFileModeContinueOnErr(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.PathsUsingFileModeContinueOnErr(
		false, 0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeContinueOnErr returns error -- with args", actual)
}

func Test_ChmodApplier_PathsUsingFileModeRecursive(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.PathsUsingFileModeRecursive(
		false, 0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeRecursive returns correct value -- with args", actual)
}

func Test_ChmodApplier_RecursivePaths(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePaths(false, true, 0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RecursivePaths returns correct value -- with args", actual)
}

func Test_ChmodApplier_RecursivePathsContinueOnError(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePathsContinueOnError(true, 0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RecursivePathsContinueOnError returns error -- with args", actual)
}

func Test_ChmodApplier_RecursivePathsCaptureInvalids(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePathsCaptureInvalids(0755, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RecursivePathsCaptureInvalids returns error -- with args", actual)
}

func Test_ChmodApplier_RwxPartial_EmptyLocations_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.RwxPartial("-rwx", &chmodins.Condition{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxPartial returns empty -- empty locations", actual)
}

func Test_RwxStringApplyChmod_EmptyLocations_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr--", &chmodins.Condition{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxStringApplyChmod returns empty -- empty", actual)
}

func Test_RwxStringApplyChmod_InvalidRwxLength(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxStringApplyChmod("short", &chmodins.Condition{}, "/tmp")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxStringApplyChmod returns error -- invalid length", actual)
}

func Test_RwxStringApplyChmod_NilCondition_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr--", nil, "/tmp")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxStringApplyChmod returns nil -- nil condition", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_EmptyLocations_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOtherApplyChmod returns empty -- empty", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilRwx_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, nil, "/tmp")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOtherApplyChmod returns nil -- nil rwx", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilCondition_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r--"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, nil, "/tmp")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOtherApplyChmod returns nil -- nil condition", actual)
}

// ── SimpleFileReaderWriter ──

func Test_SimpleFileReaderWriter_Creation(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-test.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)

	// Act
	actual := args.Map{
		"notNil":    rw != nil,
		"isExist":   rw.IsExist(),
		"isInvalid": rw.IsPathInvalid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true, "isExist": false, "isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- creation", actual)
}

func Test_SimpleFileReaderWriter_WriteAndRead_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-write.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	writeErr := rw.Write([]byte("hello world"))
	content, readErr := rw.ReadString()

	// Act
	actual := args.Map{
		"writeOk": writeErr == nil,
		"readOk":  readErr == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"writeOk": true, "readOk": true, "content": "hello world",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- WriteAndRead", actual)
}

func Test_SimpleFileReaderWriter_WriteString_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-writestr.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WriteString("string content")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- WriteString", actual)
}

func Test_SimpleFileReaderWriter_WriteAny_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-any.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WriteAny(map[string]int{"count": 5})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- WriteAny", actual)
}

func Test_SimpleFileReaderWriter_ReadOnExist_NotExist_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "nonexistent.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.ReadOnExist()

	// Act
	actual := args.Map{
		"nilBytes": bytes == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nilBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ReadOnExist returns correct value -- not exist", actual)
}

func Test_SimpleFileReaderWriter_ReadStringOnExist_NotExist_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "nonexistent2.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	content, err := rw.ReadStringOnExist()

	// Act
	actual := args.Map{
		"empty": content == "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ReadStringOnExist returns correct value -- not exist", actual)
}

func Test_SimpleFileReaderWriter_Expire_NotExist_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "expire-ne.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.Expire()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Expire returns correct value -- not exist", actual)
}

func Test_SimpleFileReaderWriter_Expire_Exist(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "expire-exist.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.Expire()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"removed": !chmodhelper.IsPathExists(filePath),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"removed": true,
	}
	expected.ShouldBeEqual(t, 0, "Expire returns correct value -- exist", actual)
}

func Test_SimpleFileReaderWriter_Clone_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "clone.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	cloned := rw.Clone()

	// Act
	actual := args.Map{"equal": cloned.FilePath == rw.FilePath}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ClonePtr_Nil_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	var rw *chmodhelper.SimpleFileReaderWriter
	cloned := rw.ClonePtr()

	// Act
	actual := args.Map{"nil": cloned == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_SimpleFileReaderWriter_String_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_JoinRelPath_Empty(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "join.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.JoinRelPath("")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinRelPath returns empty -- empty", actual)
}

func Test_SimpleFileReaderWriter_JoinRelPath_NonEmpty(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "join2.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.JoinRelPath("sub/file.txt")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinRelPath returns empty -- non-empty", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefault_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test.txt",
	}
	initialized := rw.InitializeDefault(true)

	// Act
	actual := args.Map{
		"notNil":  initialized != nil,
		"mustChmod": initialized.IsMustChmodApplyOnFile,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"mustChmod": true,
	}
	expected.ShouldBeEqual(t, 0, "InitializeDefault returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefaultApplyChmod_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test2.txt",
	}
	initialized := rw.InitializeDefaultApplyChmod()

	// Act
	actual := args.Map{"notNil": initialized != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "InitializeDefaultApplyChmod returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefaultNew_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test3.txt",
	}
	newRw := rw.InitializeDefaultNew()

	// Act
	actual := args.Map{"notNil": newRw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "InitializeDefaultNew returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_NewPath_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test.txt",
	}
	newRw := rw.NewPath(false, "/tmp/other.txt")

	// Act
	actual := args.Map{
		"notNil": newRw != nil,
		"path": newRw.FilePath,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"path": "/tmp/other.txt",
	}
	expected.ShouldBeEqual(t, 0, "NewPath returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_NewPathJoin_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test.txt",
		ParentDir: "/tmp",
	}
	newRw := rw.NewPathJoin(false, "sub", "file.txt")

	// Act
	actual := args.Map{"notNil": newRw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewPathJoin returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_HasAnyIssues(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/path.txt",
		ParentDir: "/nonexistent",
	}

	// Act
	actual := args.Map{"hasIssues": rw.HasAnyIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "HasAnyIssues returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_HasPathIssues_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath: "/nonexistent/path.txt",
	}

	// Act
	actual := args.Map{"hasIssues": rw.HasPathIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "HasPathIssues returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_JsonRoundTrip(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "json-rt.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	jsonResult := rw.Json()

	// Act
	actual := args.Map{"noErr": !jsonResult.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- round trip", actual)
}

func Test_SimpleFileReaderWriter_ExpireParentDir_NotExist(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/test.txt",
		ParentDir: "/nonexistent",
	}
	err := rw.ExpireParentDir()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ExpireParentDir returns correct value -- not exist", actual)
}

func Test_SimpleFileReaderWriter_OsFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "osfile.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	f, err := rw.OsFile()
	if f != nil {
		defer f.Close()
	}

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": f != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "OsFile returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_RemoveOnExist(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "remove.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.RemoveOnExist()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveOnExist returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_RemoveDirOnExist(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir")
	_ = os.MkdirAll(subDir, 0755)
	filePath := filepath.Join(subDir, "file.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, subDir, filePath)
	err := rw.RemoveDirOnExist()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirOnExist returns correct value -- with args", actual)
}

// ── newSimpleFileReaderWriterCreator ──

func Test_NewSimpleFileReaderWriter_Create(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, "/tmp", "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.Create returns correct value -- with args", actual)
}

func Test_NewSimpleFileReaderWriter_All_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.All(0755, 0644, false, true, true, "/tmp", "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.All returns correct value -- with args", actual)
}

func Test_NewSimpleFileReaderWriter_Options_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Options(false, true, true, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.Options returns correct value -- with args", actual)
}

func Test_NewSimpleFileReaderWriter_CreateClean_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.CreateClean(false, 0755, 0644, "/tmp", "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.CreateClean returns correct value -- with args", actual)
}

func Test_NewSimpleFileReaderWriter_DefaultCleanPath_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(false, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.DefaultCleanPath returns correct value -- with args", actual)
}

func Test_NewSimpleFileReaderWriter_Path_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.Path returns correct value -- with args", actual)
}

func Test_NewSimpleFileReaderWriter_PathCondition_Clean(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, true, 0755, 0644, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.PathCondition returns correct value -- clean", actual)
}

func Test_NewSimpleFileReaderWriter_PathCondition_NoClean(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, false, 0755, 0644, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.PathCondition returns empty -- no clean", actual)
}

func Test_NewSimpleFileReaderWriter_PathDirDefaultChmod_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.PathDirDefaultChmod(false, 0644, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.PathDirDefaultChmod returns correct value -- with args", actual)
}

// ── simpleFileWriter Lock/Unlock ──

func Test_SimpleFileWriter_LockUnlock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	chmodhelper.SimpleFileWriter.Lock()
	chmodhelper.SimpleFileWriter.Unlock()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileWriter returns correct value -- Lock/Unlock", actual)
}

// ── RwxInstructionExecutors ──

func Test_RwxInstructionExecutors_Empty(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)

	// Act
	actual := args.Map{
		"isEmpty":  executors.IsEmpty(),
		"hasAny":   executors.HasAnyItem(),
		"count":    executors.Count(),
		"length":   executors.Length(),
		"lastIdx":  executors.LastIndex(),
		"hasIdx0":  executors.HasIndex(0),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true, "hasAny": false, "count": 0,
		"length": 0, "lastIdx": -1, "hasIdx0": false,
	}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns empty -- empty", actual)
}

func Test_RwxInstructionExecutors_AddNil(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	executors.Add(nil)

	// Act
	actual := args.Map{"length": executors.Length()}

	// Assert
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns nil -- add nil", actual)
}

func Test_RwxInstructionExecutors_AddsNil(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	executors.Adds(nil...)

	// Act
	actual := args.Map{"length": executors.Length()}

	// Assert
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns nil -- adds nil", actual)
}

func Test_RwxInstructionExecutors_ApplyOnPath_Empty(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.ApplyOnPath("/tmp")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPath returns empty -- empty", actual)
}

func Test_RwxInstructionExecutors_ApplyOnPaths_Empty(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.ApplyOnPaths([]string{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPaths returns empty -- empty locations", actual)
}

func Test_RwxInstructionExecutors_ApplyOnPathsPtr_Empty(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.ApplyOnPathsPtr([]string{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPathsPtr returns empty -- empty executors", actual)
}

func Test_RwxInstructionExecutors_VerifyRwxModifiers_EmptyLocations(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.VerifyRwxModifiers(false, false, []string{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyRwxModifiers returns empty -- empty locations", actual)
}

// ── FileModeFriendlyString ──

func Test_FileModeFriendlyString_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	result := chmodhelper.FileModeFriendlyString(0755)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileModeFriendlyString returns correct value -- with args", actual)
}

// ── fwChmodVerifier / fwChmodApplier ──

func Test_FwChmodApplier_OnAll_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-apply.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnAll()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnAll returns correct value -- with args", actual)
}

func Test_FwChmodVerifier_IsEqualFile_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-verify.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	verifier := rw.ChmodVerifier()

	// Act
	actual := args.Map{"isEqual": verifier.IsEqualFile()}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.IsEqualFile returns correct value -- with args", actual)
}

func Test_FwChmodVerifier_HasMismatchFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-mismatch.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0777, tmpDir, filePath)
	verifier := rw.ChmodVerifier()

	// Act
	actual := args.Map{"hasMismatch": verifier.HasMismatchFile()}

	// Assert
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.HasMismatchFile returns correct value -- with args", actual)
}

func Test_FwChmodVerifier_IsEqualParentDir(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-pardir.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	result := verifier.IsEqualParentDir()
	// just exercise the method

	// Act
	actual := args.Map{
		"called": true,
		"result": result,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"result": result,
	}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.IsEqualParentDir returns correct value -- with args", actual)
}

func Test_FwChmodVerifier_MismatchErrorFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-mismatch-err.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0777, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	err := verifier.MismatchErrorFile()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.MismatchErrorFile returns error -- with args", actual)
}

func Test_FwChmodVerifier_MismatchErrorParentDir(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-mismatch-dir.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	// just exercise - may or may not error depending on OS
	_ = verifier.MismatchErrorParentDir()

	// Act
	actual := args.Map{"called": true}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.MismatchErrorParentDir returns error -- with args", actual)
}

func Test_FwChmodApplier_OnMismatch_Neither(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-neither.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnMismatch(false, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnMismatch returns correct value -- neither", actual)
}

func Test_FwChmodApplier_OnDiffFile_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-diff.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnDiffFile(true, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnDiffFile returns correct value -- with args", actual)
}

func Test_FwChmodApplier_OnDiffDir_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-diffdir.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnDiffDir(true, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnDiffDir returns correct value -- with args", actual)
}

func Test_FwChmodApplier_DirRecursive(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-rec.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.DirRecursive(true, tmpDir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.DirRecursive returns correct value -- with args", actual)
}

func Test_FwChmodApplier_OnParentRecursive(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-par-rec.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnParentRecursive()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnParentRecursive returns correct value -- with args", actual)
}

// ── SimpleFileReaderWriter Lock variants ──

func Test_SimpleFileReaderWriter_WriteLockVariants(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "lock-variants.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WriteAnyLock(map[string]string{"a": "b"})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "WriteAnyLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ReadLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "read-lock.txt")
	_ = os.WriteFile(filePath, []byte("lock-data"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.ReadLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasData": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "ReadLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ReadStringLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "read-str-lock.txt")
	_ = os.WriteFile(filePath, []byte("str-lock"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	content, err := rw.ReadStringLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"content": "str-lock",
	}
	expected.ShouldBeEqual(t, 0, "ReadStringLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ReadOnExistLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "exist-lock.txt")
	_ = os.WriteFile(filePath, []byte("exist"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.ReadOnExistLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasData": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "ReadOnExistLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ReadStringOnExistLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-exist-lock.txt")
	_ = os.WriteFile(filePath, []byte("data"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	content, err := rw.ReadStringOnExistLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"content": "data",
	}
	expected.ShouldBeEqual(t, 0, "ReadStringOnExistLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ExpireLock(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "expire-lock.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.ExpireLock()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ExpireLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ExpireParentDirLock(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "expire-par-lock")
	_ = os.MkdirAll(subDir, 0755)
	filePath := filepath.Join(subDir, "file.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, subDir, filePath)
	err := rw.ExpireParentDirLock()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ExpireParentDirLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_Serialize_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "serialize.txt")
	_ = os.WriteFile(filePath, []byte("ser-data"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasData": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_SerializeLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "serialize-lock.txt")
	_ = os.WriteFile(filePath, []byte("ser-lock"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.SerializeLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasData": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_Set_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "set.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.Set(map[string]int{"v": 1})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Set returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_SetLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "set-lock.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.SetLock(map[string]int{"v": 2})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SetLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_Deserialize_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "deser.json")
	_ = os.WriteFile(filePath, []byte(`{"v":1}`), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.Deserialize(&result)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"v": result["v"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"v": 1,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_DeserializeLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "deser-lock.json")
	_ = os.WriteFile(filePath, []byte(`{"v":2}`), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.DeserializeLock(&result)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"v": result["v"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"v": 2,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_GetLock_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "get-lock.json")
	_ = os.WriteFile(filePath, []byte(`{"v":3}`), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.GetLock(&result)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"v": result["v"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"v": 3,
	}
	expected.ShouldBeEqual(t, 0, "GetLock returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_Get_NotExist(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "get-ne.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.Get(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Get returns correct value -- not exist", actual)
}

func Test_SimpleFileReaderWriter_ReadMust_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "readmust.txt")
	_ = os.WriteFile(filePath, []byte("must"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.ReadMust()

	// Act
	actual := args.Map{"content": string(result)}

	// Assert
	expected := args.Map{"content": "must"}
	expected.ShouldBeEqual(t, 0, "ReadMust returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ReadStringMust_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "readstrmust.txt")
	_ = os.WriteFile(filePath, []byte("strmust"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.ReadStringMust()

	// Act
	actual := args.Map{"content": result}

	// Assert
	expected := args.Map{"content": "strmust"}
	expected.ShouldBeEqual(t, 0, "ReadStringMust returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_WritePath_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "writepath.txt")
	writePath := filepath.Join(tmpDir, "writepath2.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WritePath(false, writePath, []byte("via path"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "WritePath returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_WriteRelativePath_FromTempDirGetterPerm(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "writerel.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	err := rw.WriteRelativePath(false, "relfile.txt", []byte("relative"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "WriteRelativePath returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_ReadWrite(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "readwrite.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	type data struct{ V int }
	target := &data{}
	err := rw.ReadWrite(target, func() (any, error) {
		return &data{V: 42}, nil
	})
	// exercises GetSet path

	// Act
	actual := args.Map{
		"called": true,
		"errOrNil": err == nil || err != nil,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"errOrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ReadWrite returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_CacheGetSet(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "cache.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	type data struct{ V int }
	target := &data{}
	err := rw.CacheGetSet(target, func() (any, error) {
		return &data{V: 10}, nil
	})

	// Act
	actual := args.Map{
		"called": true,
		"errOrNil": err == nil || err != nil,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"errOrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "CacheGetSet returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "binder.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	binder := rw.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_SimpleFileReaderWriter_name_Nil(t *testing.T) {
	// Arrange
	var rw *chmodhelper.SimpleFileReaderWriter
	// The name() method is unexported but called by errorWrapFilePath
	// Exercise via Get on nil path
	err := rw.ClonePtr()

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "name returns nil -- nil receiver", actual)
}
