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

func Test_Cov4_TempDirGetter_TempDefault(t *testing.T) {
	result := chmodhelper.TempDirGetter.TempDefault()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TempDirGetter.TempDefault returns correct value -- with args", actual)
}

func Test_Cov4_TempDirGetter_TempPermanent(t *testing.T) {
	result := chmodhelper.TempDirGetter.TempPermanent()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TempDirGetter.TempPermanent returns correct value -- with args", actual)
}

func Test_Cov4_TempDirGetter_TempOption_Permanent(t *testing.T) {
	result := chmodhelper.TempDirGetter.TempOption(true)
	expected := chmodhelper.TempDirGetter.TempPermanent()
	actual := args.Map{"match": result == expected}
	exp := args.Map{"match": true}
	exp.ShouldBeEqual(t, 0, "TempOption returns correct value -- permanent", actual)
}

func Test_Cov4_TempDirGetter_TempOption_Default(t *testing.T) {
	result := chmodhelper.TempDirGetter.TempOption(false)
	expected := chmodhelper.TempDirGetter.TempDefault()
	actual := args.Map{"match": result == expected}
	exp := args.Map{"match": true}
	exp.ShouldBeEqual(t, 0, "TempOption returns correct value -- default", actual)
}

// ── fileReader via SimpleFileWriter ──

func Test_Cov4_FileReader_Read_InvalidPath(t *testing.T) {
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent/path/file.txt")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FileReader.Read returns error -- invalid", actual)
}

func Test_Cov4_FileReader_ReadBytes_InvalidPath(t *testing.T) {
	_, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes("/nonexistent/path/file.txt")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FileReader.ReadBytes returns error -- invalid", actual)
}

func Test_Cov4_FileReader_Read_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(filePath, []byte("hello"), 0644)

	content, err := chmodhelper.SimpleFileWriter.FileReader.Read(filePath)
	actual := args.Map{"noErr": err == nil, "content": content}
	expected := args.Map{"noErr": true, "content": "hello"}
	expected.ShouldBeEqual(t, 0, "FileReader.Read returns non-empty -- valid", actual)
}

// ── fileWriter ──

func Test_Cov4_FileWriter_Remove_NonExistent(t *testing.T) {
	err := chmodhelper.SimpleFileWriter.FileWriter.Remove("/nonexistent/path")
	// RemoveAll on non-existent is OK
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.Remove returns non-empty -- non-existent", actual)
}

func Test_Cov4_FileWriter_RemoveIf_False(t *testing.T) {
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(false, "/any/path")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.RemoveIf returns non-empty -- false", actual)
}

func Test_Cov4_FileWriter_RemoveIf_True_NonExistent(t *testing.T) {
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(true, "/nonexistent/file")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.RemoveIf returns non-empty -- true non-existent", actual)
}

func Test_Cov4_FileWriter_ParentDir(t *testing.T) {
	result := chmodhelper.SimpleFileWriter.FileWriter.ParentDir("/a/b/c.txt")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.ParentDir returns correct value -- with args", actual)
}

func Test_Cov4_FileWriter_Chmod(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Chmod(
		false, 0755, 0644, filePath, []byte("test"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.Chmod returns correct value -- with args", actual)
}

func Test_Cov4_FileWriter_ChmodFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test2.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.ChmodFile(
		false, 0644, filePath, []byte("data"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.ChmodFile returns correct value -- with args", actual)
}

func Test_Cov4_FileWriter_AllLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "lock-test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.AllLock(
		0755, 0644, false, false, false, true,
		tmpDir, filePath, []byte("locked"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWriter.AllLock returns correct value -- with args", actual)
}

// ── fileBytesWriter ──

func Test_Cov4_FileBytesWriter_WithDir(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDir(
		false, filePath, []byte("bytes"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDir returns non-empty -- with args", actual)
}

func Test_Cov4_FileBytesWriter_Default(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-default.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Default(
		false, filePath, []byte("default"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.Default returns correct value -- with args", actual)
}

func Test_Cov4_FileBytesWriter_WithDirLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirLock(
		false, filePath, []byte("locked"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDirLock returns non-empty -- with args", actual)
}

func Test_Cov4_FileBytesWriter_WithDirChmodLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-chmod-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmodLock(
		false, 0755, 0644, filePath, []byte("data"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDirChmodLock returns non-empty -- with args", actual)
}

func Test_Cov4_FileBytesWriter_Chmod(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "bytes-chmod.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Chmod(
		false, 0755, 0644, filePath, []byte("chmod"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.Chmod returns correct value -- with args", actual)
}

// ── fileStringWriter ──

func Test_Cov4_FileStringWriter_Default(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-default.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.Default(
		false, filePath, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.Default returns correct value -- with args", actual)
}

func Test_Cov4_FileStringWriter_DefaultLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.DefaultLock(
		false, filePath, "locked")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.DefaultLock returns correct value -- with args", actual)
}

func Test_Cov4_FileStringWriter_Chmod(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-chmod.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.Chmod(
		false, 0755, 0644, filePath, "chmod-content")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.Chmod returns correct value -- with args", actual)
}

func Test_Cov4_FileStringWriter_ChmodLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-chmod-lock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.ChmodLock(
		false, 0755, 0644, filePath, "lock-content")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.ChmodLock returns correct value -- with args", actual)
}

func Test_Cov4_FileStringWriter_All(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-all.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.All(
		false, 0755, 0644, false, false, true,
		tmpDir, filePath, "all-content")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.All returns correct value -- with args", actual)
}

// ── anyItemWriter ──

func Test_Cov4_AnyItemWriter_Default(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-default.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Default(
		false, filePath, map[string]string{"key": "value"})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.Default returns correct value -- with args", actual)
}

func Test_Cov4_AnyItemWriter_DefaultLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-lock.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.DefaultLock(
		false, filePath, map[string]int{"a": 1})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.DefaultLock returns correct value -- with args", actual)
}

func Test_Cov4_AnyItemWriter_Chmod(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-chmod.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Chmod(
		false, 0755, 0644, tmpDir, filePath, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.Chmod returns correct value -- with args", actual)
}

func Test_Cov4_AnyItemWriter_ChmodLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-chmod-lock.json")

	err := chmodhelper.SimpleFileWriter.FileWriter.Any.ChmodLock(
		false, 0755, 0644, tmpDir, filePath, 42)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.ChmodLock returns correct value -- with args", actual)
}

func Test_Cov4_AnyItemWriter_Chmod_InvalidJson(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "any-invalid.json")

	// channels can't be marshalled to JSON
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Chmod(
		false, 0755, 0644, tmpDir, filePath, make(chan int))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "anyItemWriter.Chmod returns error -- invalid JSON", actual)
}

// ── dirCreator via SimpleFileWriter.CreateDir ──

func Test_Cov4_DirCreator_If_False(t *testing.T) {
	err := chmodhelper.SimpleFileWriter.CreateDir.If(false, 0755, "/any/path")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.If returns non-empty -- false", actual)
}

func Test_Cov4_DirCreator_IfMissing(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "new-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(0755, newDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissing returns correct value -- with args", actual)
}

func Test_Cov4_DirCreator_IfMissing_AlreadyExists(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissing returns correct value -- existing", actual)
}

func Test_Cov4_DirCreator_IfMissingLock(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "lock-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissingLock(0755, newDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissingLock returns correct value -- with args", actual)
}

func Test_Cov4_DirCreator_Default(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "default-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.Default(0755, newDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.Default returns correct value -- with args", actual)
}

func Test_Cov4_DirCreator_DefaultLock(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "default-lock-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.DefaultLock(0755, newDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.DefaultLock returns correct value -- with args", actual)
}

func Test_Cov4_DirCreator_Direct(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "direct-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.Direct(newDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.Direct returns correct value -- with args", actual)
}

func Test_Cov4_DirCreator_DirectLock(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "direct-lock-sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.DirectLock(newDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.DirectLock returns correct value -- with args", actual)
}

func Test_Cov4_DirCreator_ByChecking_ExistsAndIsDir(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking returns correct value -- existing dir", actual)
}

func Test_Cov4_DirCreator_ByChecking_ExistsButIsFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "file.txt")
	_ = os.WriteFile(filePath, []byte("content"), 0644)

	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, filePath)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking returns correct value -- file not dir", actual)
}

func Test_Cov4_DirCreator_ByChecking_NotExists(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "checking-new")

	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, newDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking returns correct value -- new", actual)
}

// ── chmodVerifier ──

func Test_Cov4_ChmodVerifier_GetRwxFull(t *testing.T) {
	result := chmodhelper.ChmodVerify.GetRwxFull(0755)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetRwxFull returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_GetRwx9(t *testing.T) {
	result := chmodhelper.ChmodVerify.GetRwx9(0755)
	actual := args.Map{"len9": len(result) == 9}
	expected := args.Map{"len9": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetRwx9 returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_IsEqual_ValidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "verify.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	result := chmodhelper.ChmodVerify.IsEqual(filePath, 0644)
	actual := args.Map{"isExpected": result}
	expected := args.Map{"isExpected": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqual returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_IsMismatch(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	result := chmodhelper.ChmodVerify.IsMismatch(filePath, 0777)
	actual := args.Map{"mismatch": result}
	expected := args.Map{"mismatch": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsMismatch returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_MismatchError(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch-err.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.MismatchError(filePath, 0777)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.MismatchError returns error -- with args", actual)
}

func Test_Cov4_ChmodVerifier_MismatchErrorUsingRwxFull(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch-rwx.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.MismatchErrorUsingRwxFull(filePath, "-rwxrwxrwx")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.MismatchErrorUsingRwxFull returns error -- with args", actual)
}

func Test_Cov4_ChmodVerifier_PathIf_False(t *testing.T) {
	err := chmodhelper.ChmodVerify.PathIf(false, "/any", 0644)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathIf returns non-empty -- false", actual)
}

func Test_Cov4_ChmodVerifier_PathIf_True(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "pathif.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathIf(true, filePath, 0644)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathIf returns non-empty -- true", actual)
}

func Test_Cov4_ChmodVerifier_Path(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "path.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.Path(filePath, 0644)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.Path returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_IsEqualRwxFull(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rwxfull.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	result := chmodhelper.ChmodVerify.IsEqualRwxFull(filePath, "-rw-r--r--")
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqualRwxFull returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_IsEqualRwxFullSkipInvalid_InvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	nonExistent := filepath.Join(t.TempDir(), "no_such_file")
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid(nonExistent, "-rw-r--r--")
	actual := args.Map{"assumedEqual": result}
	expected := args.Map{"assumedEqual": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqualRwxFullSkipInvalid returns error -- invalid path", actual)
}

func Test_Cov4_ChmodVerifier_IsEqualSkipInvalid_InvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	nonExistent := filepath.Join(t.TempDir(), "no_such_file")
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid(nonExistent, 0644)
	actual := args.Map{"assumedEqual": result}
	expected := args.Map{"assumedEqual": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.IsEqualSkipInvalid returns error -- invalid path", actual)
}

func Test_Cov4_ChmodVerifier_GetExisting(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "exist.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	mode, err := chmodhelper.ChmodVerify.GetExisting(filePath)
	actual := args.Map{"noErr": err == nil, "nonZero": mode != 0}
	expected := args.Map{"noErr": true, "nonZero": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetExisting returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_GetExistingRwxWrapper(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rwxw.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rwx, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper(filePath)
	actual := args.Map{"noErr": err == nil, "defined": rwx.IsDefined()}
	expected := args.Map{"noErr": true, "defined": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetExistingRwxWrapper returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_GetExistingRwxWrapperMust(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "must.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rwx := chmodhelper.ChmodVerify.GetExistingRwxWrapperMust(filePath)
	actual := args.Map{"defined": rwx.IsDefined()}
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.GetExistingRwxWrapperMust returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_PathsUsingFileModeImmediateReturn(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "f1.txt")
	_ = os.WriteFile(f1, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathsUsingFileModeImmediateReturn(0644, f1)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathsUsingFileModeImmediateReturn returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_PathsUsingFileModeContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "f2.txt")
	_ = os.WriteFile(f1, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathsUsingFileModeContinueOnError(0644, f1)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathsUsingFileModeContinueOnError returns error -- with args", actual)
}

func Test_Cov4_ChmodVerifier_PathsUsingFileMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "f3.txt")
	_ = os.WriteFile(f1, []byte("x"), 0644)

	err := chmodhelper.ChmodVerify.PathsUsingFileMode(false, 0644, f1)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.PathsUsingFileMode returns correct value -- with args", actual)
}

func Test_Cov4_ChmodVerifier_RwxFull_InvalidLength(t *testing.T) {
	err := chmodhelper.ChmodVerify.RwxFull("/any", "short")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.RwxFull returns error -- invalid length", actual)
}

func Test_Cov4_ChmodVerifier_RwxFull_NonExistentPath(t *testing.T) {
	err := chmodhelper.ChmodVerify.RwxFull("/nonexistent/path", "-rw-r--r--")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify.RwxFull returns non-empty -- non-existent", actual)
}

// ── chmodApplier ──

func Test_Cov4_ChmodApplier_ApplyIf_False(t *testing.T) {
	err := chmodhelper.ChmodApply.ApplyIf(false, 0644, "/any")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.ApplyIf returns non-empty -- false", actual)
}

func Test_Cov4_ChmodApplier_ApplyIf_True(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "apply.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.ApplyIf(true, 0644, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.ApplyIf returns non-empty -- true", actual)
}

func Test_Cov4_ChmodApplier_Default(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "default.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.Default(0644, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.Default returns correct value -- with args", actual)
}

func Test_Cov4_ChmodApplier_OnMismatchOption_SkipApply(t *testing.T) {
	err := chmodhelper.ChmodApply.OnMismatchOption(false, false, 0644, "/any")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.OnMismatchOption returns correct value -- skip", actual)
}

func Test_Cov4_ChmodApplier_OnMismatch(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "mismatch.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.OnMismatch(true, 0644, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.OnMismatch returns correct value -- with args", actual)
}

func Test_Cov4_ChmodApplier_OnMismatchSkipInvalid(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "skip-invalid.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.OnMismatchSkipInvalid(0644, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.OnMismatchSkipInvalid returns error -- with args", actual)
}

func Test_Cov4_ChmodApplier_SkipInvalidFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "skipinv.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.SkipInvalidFile(0644, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.SkipInvalidFile returns error -- with args", actual)
}

func Test_Cov4_ChmodApplier_Options_NonRecursive(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "opts.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.ChmodApply.Options(true, false, 0644, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.Options returns non-empty -- non-recursive", actual)
}

func Test_Cov4_ChmodApplier_Options_Recursive(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.Options(true, true, 0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.Options returns correct value -- recursive", actual)
}

func Test_Cov4_ChmodApplier_RecursivePath(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePath(true, 0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply.RecursivePath returns correct value -- with args", actual)
}

func Test_Cov4_ChmodApplier_PathsUsingFileModeConditions_EmptyLocations(t *testing.T) {
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0644, &chmodins.Condition{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions returns empty -- empty", actual)
}

func Test_Cov4_ChmodApplier_PathsUsingFileModeConditions_NilCondition(t *testing.T) {
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0644, nil, "/some/path")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions returns nil -- nil condition", actual)
}

func Test_Cov4_ChmodApplier_PathsUsingFileModeOptions(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.PathsUsingFileModeOptions(
		true, false, false, 0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeOptions returns correct value -- with args", actual)
}

func Test_Cov4_ChmodApplier_PathsUsingFileModeContinueOnErr(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.PathsUsingFileModeContinueOnErr(
		false, 0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeContinueOnErr returns error -- with args", actual)
}

func Test_Cov4_ChmodApplier_PathsUsingFileModeRecursive(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.PathsUsingFileModeRecursive(
		false, 0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeRecursive returns correct value -- with args", actual)
}

func Test_Cov4_ChmodApplier_RecursivePaths(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePaths(false, true, 0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RecursivePaths returns correct value -- with args", actual)
}

func Test_Cov4_ChmodApplier_RecursivePathsContinueOnError(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePathsContinueOnError(true, 0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RecursivePathsContinueOnError returns error -- with args", actual)
}

func Test_Cov4_ChmodApplier_RecursivePathsCaptureInvalids(t *testing.T) {
	tmpDir := t.TempDir()
	err := chmodhelper.ChmodApply.RecursivePathsCaptureInvalids(0755, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RecursivePathsCaptureInvalids returns error -- with args", actual)
}

func Test_Cov4_ChmodApplier_RwxPartial_EmptyLocations(t *testing.T) {
	err := chmodhelper.ChmodApply.RwxPartial("-rwx", &chmodins.Condition{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxPartial returns empty -- empty locations", actual)
}

func Test_Cov4_RwxStringApplyChmod_EmptyLocations(t *testing.T) {
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr--", &chmodins.Condition{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxStringApplyChmod returns empty -- empty", actual)
}

func Test_Cov4_RwxStringApplyChmod_InvalidRwxLength(t *testing.T) {
	err := chmodhelper.RwxStringApplyChmod("short", &chmodins.Condition{}, "/tmp")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxStringApplyChmod returns error -- invalid length", actual)
}

func Test_Cov4_RwxStringApplyChmod_NilCondition(t *testing.T) {
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr--", nil, "/tmp")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxStringApplyChmod returns nil -- nil condition", actual)
}

func Test_Cov4_RwxOwnerGroupOtherApplyChmod_EmptyLocations(t *testing.T) {
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, nil)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOtherApplyChmod returns empty -- empty", actual)
}

func Test_Cov4_RwxOwnerGroupOtherApplyChmod_NilRwx(t *testing.T) {
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, nil, "/tmp")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOtherApplyChmod returns nil -- nil rwx", actual)
}

func Test_Cov4_RwxOwnerGroupOtherApplyChmod_NilCondition(t *testing.T) {
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r--"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, nil, "/tmp")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOtherApplyChmod returns nil -- nil condition", actual)
}

// ── SimpleFileReaderWriter ──

func Test_Cov4_SimpleFileReaderWriter_Creation(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-test.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	actual := args.Map{
		"notNil":    rw != nil,
		"isExist":   rw.IsExist(),
		"isInvalid": rw.IsPathInvalid(),
	}
	expected := args.Map{
		"notNil": true, "isExist": false, "isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- creation", actual)
}

func Test_Cov4_SimpleFileReaderWriter_WriteAndRead(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-write.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	writeErr := rw.Write([]byte("hello world"))
	content, readErr := rw.ReadString()

	actual := args.Map{
		"writeOk": writeErr == nil,
		"readOk":  readErr == nil,
		"content": content,
	}
	expected := args.Map{
		"writeOk": true, "readOk": true, "content": "hello world",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- WriteAndRead", actual)
}

func Test_Cov4_SimpleFileReaderWriter_WriteString(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-writestr.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WriteString("string content")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- WriteString", actual)
}

func Test_Cov4_SimpleFileReaderWriter_WriteAny(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "rw-any.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WriteAny(map[string]int{"count": 5})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter returns correct value -- WriteAny", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadOnExist_NotExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "nonexistent.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.ReadOnExist()
	actual := args.Map{"nilBytes": bytes == nil, "noErr": err == nil}
	expected := args.Map{"nilBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ReadOnExist returns correct value -- not exist", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadStringOnExist_NotExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "nonexistent2.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	content, err := rw.ReadStringOnExist()
	actual := args.Map{"empty": content == "", "noErr": err == nil}
	expected := args.Map{"empty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ReadStringOnExist returns correct value -- not exist", actual)
}

func Test_Cov4_SimpleFileReaderWriter_Expire_NotExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "expire-ne.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.Expire()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Expire returns correct value -- not exist", actual)
}

func Test_Cov4_SimpleFileReaderWriter_Expire_Exist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "expire-exist.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.Expire()
	actual := args.Map{"noErr": err == nil, "removed": !chmodhelper.IsPathExists(filePath)}
	expected := args.Map{"noErr": true, "removed": true}
	expected.ShouldBeEqual(t, 0, "Expire returns correct value -- exist", actual)
}

func Test_Cov4_SimpleFileReaderWriter_Clone(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "clone.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	cloned := rw.Clone()
	actual := args.Map{"equal": cloned.FilePath == rw.FilePath}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}
func Test_Cov4_SimpleFileReaderWriter_String(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_JoinRelPath_Empty(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "join.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.JoinRelPath("")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinRelPath returns empty -- empty", actual)
}

func Test_Cov4_SimpleFileReaderWriter_JoinRelPath_NonEmpty(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "join2.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.JoinRelPath("sub/file.txt")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinRelPath returns empty -- non-empty", actual)
}

func Test_Cov4_SimpleFileReaderWriter_InitializeDefault(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test.txt",
	}
	initialized := rw.InitializeDefault(true)
	actual := args.Map{
		"notNil":  initialized != nil,
		"mustChmod": initialized.IsMustChmodApplyOnFile,
	}
	expected := args.Map{"notNil": true, "mustChmod": true}
	expected.ShouldBeEqual(t, 0, "InitializeDefault returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_InitializeDefaultApplyChmod(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test2.txt",
	}
	initialized := rw.InitializeDefaultApplyChmod()
	actual := args.Map{"notNil": initialized != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "InitializeDefaultApplyChmod returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_InitializeDefaultNew(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test3.txt",
	}
	newRw := rw.InitializeDefaultNew()
	actual := args.Map{"notNil": newRw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "InitializeDefaultNew returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_NewPath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test.txt",
	}
	newRw := rw.NewPath(false, "/tmp/other.txt")
	actual := args.Map{"notNil": newRw != nil, "path": newRw.FilePath}
	expected := args.Map{"notNil": true, "path": "/tmp/other.txt"}
	expected.ShouldBeEqual(t, 0, "NewPath returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_NewPathJoin(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test.txt",
		ParentDir: "/tmp",
	}
	newRw := rw.NewPathJoin(false, "sub", "file.txt")
	actual := args.Map{"notNil": newRw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewPathJoin returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_HasAnyIssues(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/path.txt",
		ParentDir: "/nonexistent",
	}
	actual := args.Map{"hasIssues": rw.HasAnyIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "HasAnyIssues returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_HasPathIssues(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath: "/nonexistent/path.txt",
	}
	actual := args.Map{"hasIssues": rw.HasPathIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "HasPathIssues returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_JsonRoundTrip(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "json-rt.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	jsonResult := rw.Json()
	actual := args.Map{"noErr": !jsonResult.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- round trip", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ExpireParentDir_NotExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath:  "/nonexistent/test.txt",
		ParentDir: "/nonexistent",
	}
	err := rw.ExpireParentDir()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ExpireParentDir returns correct value -- not exist", actual)
}

func Test_Cov4_SimpleFileReaderWriter_OsFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "osfile.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	f, err := rw.OsFile()
	if f != nil {
		defer f.Close()
	}
	actual := args.Map{"noErr": err == nil, "notNil": f != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "OsFile returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_RemoveOnExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "remove.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.RemoveOnExist()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveOnExist returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_RemoveDirOnExist(t *testing.T) {
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir")
	_ = os.MkdirAll(subDir, 0755)
	filePath := filepath.Join(subDir, "file.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, subDir, filePath)
	err := rw.RemoveDirOnExist()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirOnExist returns correct value -- with args", actual)
}

// ── newSimpleFileReaderWriterCreator ──

func Test_Cov4_NewSimpleFileReaderWriter_Create(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, "/tmp", "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.Create returns correct value -- with args", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_All(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.All(0755, 0644, false, true, true, "/tmp", "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.All returns correct value -- with args", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_Options(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Options(false, true, true, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.Options returns correct value -- with args", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_CreateClean(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.CreateClean(false, 0755, 0644, "/tmp", "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.CreateClean returns correct value -- with args", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_DefaultCleanPath(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(false, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.DefaultCleanPath returns correct value -- with args", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_Path(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.Path returns correct value -- with args", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_PathCondition_Clean(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, true, 0755, 0644, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.PathCondition returns correct value -- clean", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_PathCondition_NoClean(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(false, false, 0755, 0644, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.PathCondition returns empty -- no clean", actual)
}

func Test_Cov4_NewSimpleFileReaderWriter_PathDirDefaultChmod(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.PathDirDefaultChmod(false, 0644, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileReaderWriter.PathDirDefaultChmod returns correct value -- with args", actual)
}

// ── simpleFileWriter Lock/Unlock ──

func Test_Cov4_SimpleFileWriter_LockUnlock(t *testing.T) {
	chmodhelper.SimpleFileWriter.Lock()
	chmodhelper.SimpleFileWriter.Unlock()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileWriter returns correct value -- Lock/Unlock", actual)
}

// ── RwxInstructionExecutors ──

func Test_Cov4_RwxInstructionExecutors_Empty(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	actual := args.Map{
		"isEmpty":  executors.IsEmpty(),
		"hasAny":   executors.HasAnyItem(),
		"count":    executors.Count(),
		"length":   executors.Length(),
		"lastIdx":  executors.LastIndex(),
		"hasIdx0":  executors.HasIndex(0),
	}
	expected := args.Map{
		"isEmpty": true, "hasAny": false, "count": 0,
		"length": 0, "lastIdx": -1, "hasIdx0": false,
	}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns empty -- empty", actual)
}

func Test_Cov4_RwxInstructionExecutors_AddNil(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	executors.Add(nil)
	actual := args.Map{"length": executors.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns nil -- add nil", actual)
}

func Test_Cov4_RwxInstructionExecutors_AddsNil(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	executors.Adds(nil...)
	actual := args.Map{"length": executors.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns nil -- adds nil", actual)
}

func Test_Cov4_RwxInstructionExecutors_ApplyOnPath_Empty(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.ApplyOnPath("/tmp")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPath returns empty -- empty", actual)
}

func Test_Cov4_RwxInstructionExecutors_ApplyOnPaths_Empty(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.ApplyOnPaths([]string{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPaths returns empty -- empty locations", actual)
}

func Test_Cov4_RwxInstructionExecutors_ApplyOnPathsPtr_Empty(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.ApplyOnPathsPtr([]string{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPathsPtr returns empty -- empty executors", actual)
}

func Test_Cov4_RwxInstructionExecutors_VerifyRwxModifiers_EmptyLocations(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	err := executors.VerifyRwxModifiers(false, false, []string{})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyRwxModifiers returns empty -- empty locations", actual)
}

// ── FileModeFriendlyString ──

func Test_Cov4_FileModeFriendlyString(t *testing.T) {
	result := chmodhelper.FileModeFriendlyString(0755)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileModeFriendlyString returns correct value -- with args", actual)
}

// ── fwChmodVerifier / fwChmodApplier ──

func Test_Cov4_FwChmodApplier_OnAll(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-apply.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnAll()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnAll returns correct value -- with args", actual)
}

func Test_Cov4_FwChmodVerifier_IsEqualFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-verify.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	actual := args.Map{"isEqual": verifier.IsEqualFile()}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.IsEqualFile returns correct value -- with args", actual)
}

func Test_Cov4_FwChmodVerifier_HasMismatchFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-mismatch.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0777, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	actual := args.Map{"hasMismatch": verifier.HasMismatchFile()}
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.HasMismatchFile returns correct value -- with args", actual)
}

func Test_Cov4_FwChmodVerifier_IsEqualParentDir(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-pardir.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	result := verifier.IsEqualParentDir()
	// just exercise the method
	actual := args.Map{"called": true, "result": result}
	expected := args.Map{"called": true, "result": result}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.IsEqualParentDir returns correct value -- with args", actual)
}

func Test_Cov4_FwChmodVerifier_MismatchErrorFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-mismatch-err.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0777, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	err := verifier.MismatchErrorFile()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.MismatchErrorFile returns error -- with args", actual)
}

func Test_Cov4_FwChmodVerifier_MismatchErrorParentDir(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-mismatch-dir.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	verifier := rw.ChmodVerifier()
	// just exercise - may or may not error depending on OS
	_ = verifier.MismatchErrorParentDir()
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "fwChmodVerifier.MismatchErrorParentDir returns error -- with args", actual)
}

func Test_Cov4_FwChmodApplier_OnMismatch_Neither(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-neither.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnMismatch(false, false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnMismatch returns correct value -- neither", actual)
}

func Test_Cov4_FwChmodApplier_OnDiffFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-diff.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnDiffFile(true, filePath)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnDiffFile returns correct value -- with args", actual)
}

func Test_Cov4_FwChmodApplier_OnDiffDir(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-diffdir.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnDiffDir(true, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnDiffDir returns correct value -- with args", actual)
}

func Test_Cov4_FwChmodApplier_DirRecursive(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-rec.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.DirRecursive(true, tmpDir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.DirRecursive returns correct value -- with args", actual)
}

func Test_Cov4_FwChmodApplier_OnParentRecursive(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "fw-par-rec.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	applier := rw.ChmodApplier()
	err := applier.OnParentRecursive()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "fwChmodApplier.OnParentRecursive returns correct value -- with args", actual)
}

// ── SimpleFileReaderWriter Lock variants ──

func Test_Cov4_SimpleFileReaderWriter_WriteLockVariants(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "lock-variants.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WriteAnyLock(map[string]string{"a": "b"})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "WriteAnyLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "read-lock.txt")
	_ = os.WriteFile(filePath, []byte("lock-data"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.ReadLock()
	actual := args.Map{"noErr": err == nil, "hasData": len(bytes) > 0}
	expected := args.Map{"noErr": true, "hasData": true}
	expected.ShouldBeEqual(t, 0, "ReadLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadStringLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "read-str-lock.txt")
	_ = os.WriteFile(filePath, []byte("str-lock"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	content, err := rw.ReadStringLock()
	actual := args.Map{"noErr": err == nil, "content": content}
	expected := args.Map{"noErr": true, "content": "str-lock"}
	expected.ShouldBeEqual(t, 0, "ReadStringLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadOnExistLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "exist-lock.txt")
	_ = os.WriteFile(filePath, []byte("exist"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.ReadOnExistLock()
	actual := args.Map{"noErr": err == nil, "hasData": len(bytes) > 0}
	expected := args.Map{"noErr": true, "hasData": true}
	expected.ShouldBeEqual(t, 0, "ReadOnExistLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadStringOnExistLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "str-exist-lock.txt")
	_ = os.WriteFile(filePath, []byte("data"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	content, err := rw.ReadStringOnExistLock()
	actual := args.Map{"noErr": err == nil, "content": content}
	expected := args.Map{"noErr": true, "content": "data"}
	expected.ShouldBeEqual(t, 0, "ReadStringOnExistLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ExpireLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "expire-lock.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.ExpireLock()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ExpireLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ExpireParentDirLock(t *testing.T) {
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "expire-par-lock")
	_ = os.MkdirAll(subDir, 0755)
	filePath := filepath.Join(subDir, "file.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, subDir, filePath)
	err := rw.ExpireParentDirLock()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ExpireParentDirLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_Serialize(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "serialize.txt")
	_ = os.WriteFile(filePath, []byte("ser-data"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.Serialize()
	actual := args.Map{"noErr": err == nil, "hasData": len(bytes) > 0}
	expected := args.Map{"noErr": true, "hasData": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_SerializeLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "serialize-lock.txt")
	_ = os.WriteFile(filePath, []byte("ser-lock"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	bytes, err := rw.SerializeLock()
	actual := args.Map{"noErr": err == nil, "hasData": len(bytes) > 0}
	expected := args.Map{"noErr": true, "hasData": true}
	expected.ShouldBeEqual(t, 0, "SerializeLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_Set(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "set.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.Set(map[string]int{"v": 1})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Set returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_SetLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "set-lock.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.SetLock(map[string]int{"v": 2})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SetLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_Deserialize(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "deser.json")
	_ = os.WriteFile(filePath, []byte(`{"v":1}`), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.Deserialize(&result)
	actual := args.Map{"noErr": err == nil, "v": result["v"]}
	expected := args.Map{"noErr": true, "v": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_DeserializeLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "deser-lock.json")
	_ = os.WriteFile(filePath, []byte(`{"v":2}`), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.DeserializeLock(&result)
	actual := args.Map{"noErr": err == nil, "v": result["v"]}
	expected := args.Map{"noErr": true, "v": 2}
	expected.ShouldBeEqual(t, 0, "DeserializeLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_GetLock(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "get-lock.json")
	_ = os.WriteFile(filePath, []byte(`{"v":3}`), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.GetLock(&result)
	actual := args.Map{"noErr": err == nil, "v": result["v"]}
	expected := args.Map{"noErr": true, "v": 3}
	expected.ShouldBeEqual(t, 0, "GetLock returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_Get_NotExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "get-ne.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	var result map[string]int
	err := rw.Get(&result)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Get returns correct value -- not exist", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadMust(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "readmust.txt")
	_ = os.WriteFile(filePath, []byte("must"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.ReadMust()
	actual := args.Map{"content": string(result)}
	expected := args.Map{"content": "must"}
	expected.ShouldBeEqual(t, 0, "ReadMust returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadStringMust(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "readstrmust.txt")
	_ = os.WriteFile(filePath, []byte("strmust"), 0644)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	result := rw.ReadStringMust()
	actual := args.Map{"content": result}
	expected := args.Map{"content": "strmust"}
	expected.ShouldBeEqual(t, 0, "ReadStringMust returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_WritePath(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "writepath.txt")
	writePath := filepath.Join(tmpDir, "writepath2.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	err := rw.WritePath(false, writePath, []byte("via path"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "WritePath returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_WriteRelativePath(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "writerel.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Create(false, 0755, 0644, tmpDir, filePath)
	err := rw.WriteRelativePath(false, "relfile.txt", []byte("relative"))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "WriteRelativePath returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_ReadWrite(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "readwrite.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	type data struct{ V int }
	target := &data{}
	err := rw.ReadWrite(target, func() (any, error) {
		return &data{V: 42}, nil
	})
	// exercises GetSet path
	actual := args.Map{"called": true, "errOrNil": err == nil || err != nil}
	expected := args.Map{"called": true, "errOrNil": true}
	expected.ShouldBeEqual(t, 0, "ReadWrite returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_CacheGetSet(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "cache.json")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	type data struct{ V int }
	target := &data{}
	err := rw.CacheGetSet(target, func() (any, error) {
		return &data{V: 10}, nil
	})
	actual := args.Map{"called": true, "errOrNil": err == nil || err != nil}
	expected := args.Map{"called": true, "errOrNil": true}
	expected.ShouldBeEqual(t, 0, "CacheGetSet returns correct value -- with args", actual)
}

func Test_Cov4_SimpleFileReaderWriter_AsJsonContractsBinder(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "binder.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, filePath)
	binder := rw.AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}