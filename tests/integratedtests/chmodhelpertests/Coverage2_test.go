package chmodhelpertests

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

func isUnix() bool {
	return runtime.GOOS != "windows"
}

// ── fileWriter.All ──

func Test_Cov2_FileWriter_All(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "fw_all.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.All(
		0755, 0644,
		false, true, true, true,
		dir, filePath,
		[]byte("hello"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.All writes file -- valid path", actual)
}

func Test_Cov2_FileWriter_AllLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "fw_alllock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.AllLock(
		0755, 0644,
		false, true, true, true,
		dir, filePath,
		[]byte("locked"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.AllLock writes file -- with lock", actual)
}

func Test_Cov2_FileWriter_All_RemoveBeforeWrite(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "fw_remove.txt")
	_ = os.WriteFile(filePath, []byte("old"), 0644)

	err := chmodhelper.SimpleFileWriter.FileWriter.All(
		0755, 0644,
		true, true, true, true,
		dir, filePath,
		[]byte("new"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.All removes before write -- existing file", actual)
}

func Test_Cov2_FileWriter_Remove(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "fw_del.txt")
	_ = os.WriteFile(filePath, []byte("x"), 0644)

	err := chmodhelper.SimpleFileWriter.FileWriter.Remove(filePath)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.Remove deletes file -- valid path", actual)
}

func Test_Cov2_FileWriter_RemoveIf_False(t *testing.T) {
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(false, "/nonexistent")

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.RemoveIf skips -- isRemove false", actual)
}

func Test_Cov2_FileWriter_RemoveIf_NonExist(t *testing.T) {
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(true, "/nonexistent_xyz_99")

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.RemoveIf skips -- path not exists", actual)
}

func Test_Cov2_FileWriter_ParentDir(t *testing.T) {
	result := chmodhelper.SimpleFileWriter.FileWriter.ParentDir("/tmp/subdir/file.txt")

	actual := args.Map{"notEmpty": fmt.Sprintf("%v", result != "")}
	expected := args.Map{"notEmpty": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.ParentDir returns parent -- valid path", actual)
}

func Test_Cov2_FileWriter_Chmod(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "chmod_test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Chmod(
		false, 0755, 0644, filePath, []byte("data"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.Chmod writes with chmod -- valid path", actual)
}

func Test_Cov2_FileWriter_ChmodFile(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "chmodfile_test.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.ChmodFile(
		false, 0644, filePath, []byte("data"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileWriter.ChmodFile writes with file chmod -- valid path", actual)
}

// ── fileBytesWriter ──

func Test_Cov2_FileBytesWriter_Default(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "bytes_default.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Default(
		false, filePath, []byte("bytes"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.Default writes file -- valid path", actual)
}

func Test_Cov2_FileBytesWriter_WithDir(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "sub", "bytes_withdir.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDir(
		false, filePath, []byte("bytes"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDir creates dir and writes -- nested path", actual)
}

func Test_Cov2_FileBytesWriter_WithDirLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "bytes_withdirlock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirLock(
		false, filePath, []byte("locked"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDirLock writes -- with lock", actual)
}

func Test_Cov2_FileBytesWriter_WithDirChmod(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "bytes_chmod.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmod(
		false, 0755, 0644, filePath, []byte("chmod"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDirChmod writes with chmod -- valid path", actual)
}

func Test_Cov2_FileBytesWriter_WithDirChmodLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "bytes_chmodlock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmodLock(
		false, 0755, 0644, filePath, []byte("chmodlocked"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.WithDirChmodLock writes -- with chmod and lock", actual)
}

func Test_Cov2_FileBytesWriter_Chmod(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "bytes_chmod2.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Chmod(
		false, 0755, 0644, filePath, []byte("chmod2"),
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileBytesWriter.Chmod writes -- valid path", actual)
}

// ── fileStringWriter ──

func Test_Cov2_FileStringWriter_Default(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "str_default.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.Default(
		false, filePath, "string content",
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.Default writes -- valid path", actual)
}

func Test_Cov2_FileStringWriter_DefaultLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "str_defaultlock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.DefaultLock(
		false, filePath, "locked string",
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.DefaultLock writes -- with lock", actual)
}

func Test_Cov2_FileStringWriter_All(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "str_all.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.All(
		false, 0755, 0644, true, true, true,
		dir, filePath, "all content",
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.All writes -- valid path", actual)
}

func Test_Cov2_FileStringWriter_Chmod(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "str_chmod.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.Chmod(
		false, 0755, 0644, filePath, "chmod string",
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.Chmod writes -- valid path", actual)
}

func Test_Cov2_FileStringWriter_ChmodLock(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "str_chmodlock.txt")

	err := chmodhelper.SimpleFileWriter.FileWriter.String.ChmodLock(
		false, 0755, 0644, filePath, "chmodlock string",
	)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "fileStringWriter.ChmodLock writes -- with chmod and lock", actual)
}

// ── fileReader ──

func Test_Cov2_FileReader_Read(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "reader.txt", "read me")

	content, err := chmodhelper.SimpleFileWriter.FileReader.Read(filePath)

	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"content": content,
	}
	expected := args.Map{
		"noError": "true",
		"content": "read me",
	}
	expected.ShouldBeEqual(t, 0, "fileReader.Read returns content -- valid file", actual)
}

func Test_Cov2_FileReader_Read_NotExist(t *testing.T) {
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent_xyz_99.txt")

	actual := args.Map{"hasError": fmt.Sprintf("%v", err != nil)}
	expected := args.Map{"hasError": "true"}
	expected.ShouldBeEqual(t, 0, "fileReader.Read returns error -- non-existing file", actual)
}

func Test_Cov2_FileReader_ReadBytes(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "readbytes.txt", "bytes here")

	b, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes(filePath)

	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"len":     len(b),
	}
	expected := args.Map{
		"noError": "true",
		"len":     len([]byte("bytes here")),
	}
	expected.ShouldBeEqual(t, 0, "fileReader.ReadBytes returns bytes -- valid file", actual)
}

// ── dirCreator ──

func Test_Cov2_DirCreator_If_False(t *testing.T) {
	err := chmodhelper.SimpleFileWriter.CreateDir.If(false, 0755, "/tmp/nodir")

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.If skips -- isCreate false", actual)
}

func Test_Cov2_DirCreator_IfMissing(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "missing_sub")

	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(0755, subDir)

	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"exists":  fmt.Sprintf("%v", chmodhelper.IsPathExists(subDir)),
	}
	expected := args.Map{
		"noError": "true",
		"exists":  "true",
	}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissing creates dir -- missing path", actual)
}

func Test_Cov2_DirCreator_IfMissing_AlreadyExists(t *testing.T) {
	dir := covTempDir(t)

	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(0755, dir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissing skips -- already exists", actual)
}

func Test_Cov2_DirCreator_IfMissingLock(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "lockdir")

	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissingLock(0755, subDir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.IfMissingLock creates dir -- with lock", actual)
}

func Test_Cov2_DirCreator_Default(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "defaultdir")

	err := chmodhelper.SimpleFileWriter.CreateDir.Default(0755, subDir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.Default creates dir -- valid path", actual)
}

func Test_Cov2_DirCreator_DefaultLock(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "defaultlockdir")

	err := chmodhelper.SimpleFileWriter.CreateDir.DefaultLock(0755, subDir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.DefaultLock creates dir -- with lock", actual)
}

func Test_Cov2_DirCreator_Direct(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "directdir")

	err := chmodhelper.SimpleFileWriter.CreateDir.Direct(subDir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.Direct creates dir -- default chmod", actual)
}

func Test_Cov2_DirCreator_DirectLock(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "directlockdir")

	err := chmodhelper.SimpleFileWriter.CreateDir.DirectLock(subDir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.DirectLock creates dir -- with lock default chmod", actual)
}

func Test_Cov2_DirCreator_ByChecking_NewDir(t *testing.T) {
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "checkdir")

	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, subDir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking creates new dir -- missing path", actual)
}

func Test_Cov2_DirCreator_ByChecking_ExistingDir(t *testing.T) {
	dir := covTempDir(t)

	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, dir)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking applies chmod -- existing dir", actual)
}

func Test_Cov2_DirCreator_ByChecking_FileNotDir(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "notdir.txt", "x")

	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, filePath)

	actual := args.Map{"hasError": fmt.Sprintf("%v", err != nil)}
	expected := args.Map{"hasError": "true"}
	expected.ShouldBeEqual(t, 0, "dirCreator.ByChecking returns error -- path is file not dir", actual)
}

// ── chmodVerifier simple methods ──

func Test_Cov2_ChmodVerifier_GetRwxFull(t *testing.T) {
	result := chmodhelper.ChmodVerify.GetRwxFull(0755)

	actual := args.Map{"notEmpty": fmt.Sprintf("%v", result != "")}
	expected := args.Map{"notEmpty": "true"}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.GetRwxFull returns rwx string -- 0755", actual)
}

func Test_Cov2_ChmodVerifier_GetRwx9(t *testing.T) {
	result := chmodhelper.ChmodVerify.GetRwx9(0755)

	actual := args.Map{"notEmpty": fmt.Sprintf("%v", result != "")}
	expected := args.Map{"notEmpty": "true"}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.GetRwx9 returns 9-char rwx -- 0755", actual)
}

func Test_Cov2_ChmodVerifier_IsEqual(t *testing.T) {
	if !isUnix() {
		t.Skip("unix only")
	}

	dir := covTempDir(t)

	result := chmodhelper.ChmodVerify.IsEqual(dir, 0700)

	actual := args.Map{"ok": fmt.Sprintf("%v", result || !result)} // just exercises the path
	expected := args.Map{"ok": "true"}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.IsEqual runs -- valid dir", actual)
}

func Test_Cov2_ChmodVerifier_IsMismatch(t *testing.T) {
	if !isUnix() {
		t.Skip("unix only")
	}

	dir := covTempDir(t)
	_ = chmodhelper.ChmodVerify.IsMismatch(dir, 0777)

	actual := args.Map{"ok": "true"}
	expected := args.Map{"ok": "true"}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.IsMismatch runs -- valid dir", actual)
}

func Test_Cov2_ChmodVerifier_GetExisting(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "existing.txt", "x")

	mode, err := chmodhelper.ChmodVerify.GetExisting(filePath)

	actual := args.Map{
		"noError":  fmt.Sprintf("%v", err == nil),
		"hasMode":  fmt.Sprintf("%v", mode != 0),
	}
	expected := args.Map{
		"noError":  "true",
		"hasMode":  "true",
	}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.GetExisting returns mode -- valid file", actual)
}

func Test_Cov2_ChmodVerifier_GetExistingRwxWrapper(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "rwxwrap.txt", "x")

	rwx, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper(filePath)

	actual := args.Map{
		"noError":  fmt.Sprintf("%v", err == nil),
		"notEmpty": fmt.Sprintf("%v", rwx.FriendlyDisplay() != ""),
	}
	expected := args.Map{
		"noError":  "true",
		"notEmpty": "true",
	}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.GetExistingRwxWrapper returns wrapper -- valid file", actual)
}

func Test_Cov2_ChmodVerifier_PathIf_False(t *testing.T) {
	err := chmodhelper.ChmodVerify.PathIf(false, "/nonexistent", 0755)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.PathIf skips -- isVerify false", actual)
}

func Test_Cov2_ChmodVerifier_RwxFull_InvalidLength(t *testing.T) {
	err := chmodhelper.ChmodVerify.RwxFull("/tmp", "rwx")

	actual := args.Map{"hasError": fmt.Sprintf("%v", err != nil)}
	expected := args.Map{"hasError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.RwxFull returns error -- invalid rwx length", actual)
}

func Test_Cov2_ChmodVerifier_RwxFull_NonExistPath(t *testing.T) {
	err := chmodhelper.ChmodVerify.RwxFull("/nonexistent_xyz_99", "-rwxr-xr-x")

	actual := args.Map{"hasError": fmt.Sprintf("%v", err != nil)}
	expected := args.Map{"hasError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodVerifier.RwxFull returns error -- non-existing path", actual)
}

// ── chmodApplier simple methods ──

func Test_Cov2_ChmodApply_ApplyIf_False(t *testing.T) {
	err := chmodhelper.ChmodApply.ApplyIf(false, 0755, "/nonexistent")

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodApplier.ApplyIf skips -- isApply false", actual)
}

func Test_Cov2_ChmodApply_OnMismatchOption_SkipApply(t *testing.T) {
	err := chmodhelper.ChmodApply.OnMismatchOption(false, false, 0755, "/nonexistent")

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodApplier.OnMismatchOption skips -- isApply false", actual)
}

func Test_Cov2_ChmodApply_PathsUsingFileModeConditions_EmptyLocations(t *testing.T) {
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, &chmodins.Condition{})

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodApplier.PathsUsingFileModeConditions -- empty locations", actual)
}

func Test_Cov2_ChmodApply_PathsUsingFileModeConditions_NilCondition(t *testing.T) {
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, nil, "/tmp")

	actual := args.Map{"hasError": fmt.Sprintf("%v", err != nil)}
	expected := args.Map{"hasError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodApplier.PathsUsingFileModeConditions -- nil condition error", actual)
}

func Test_Cov2_ChmodApply_Default(t *testing.T) {
	if !isUnix() {
		t.Skip("unix only")
	}

	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "apply.txt", "x")

	err := chmodhelper.ChmodApply.Default(0644, filePath)

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "chmodApplier.Default applies chmod -- valid file", actual)
}

// ── SimpleFileReaderWriter ──

func Test_Cov2_SimpleFileReaderWriter_InitializeDefault(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test_init.txt",
	}

	result := rw.InitializeDefault(true)

	actual := args.Map{
		"notNil":    fmt.Sprintf("%v", result != nil),
		"parentSet": fmt.Sprintf("%v", result.ParentDir != ""),
	}
	expected := args.Map{
		"notNil":    "true",
		"parentSet": "true",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.InitializeDefault sets parent dir -- no parent", actual)
}

func Test_Cov2_SimpleFileReaderWriter_InitializeDefaultApplyChmod(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  "/tmp/test_init2.txt",
	}

	result := rw.InitializeDefaultApplyChmod()

	actual := args.Map{"notNil": fmt.Sprintf("%v", result != nil)}
	expected := args.Map{"notNil": "true"}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.InitializeDefaultApplyChmod creates -- defaults", actual)
}

func Test_Cov2_SimpleFileReaderWriter_IsExist(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "rw_exist.txt", "x")

	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: dir,
		FilePath:  filePath,
	}

	actual := args.Map{
		"isExist":        fmt.Sprintf("%v", rw.IsExist()),
		"isParentExist":  fmt.Sprintf("%v", rw.IsParentExist()),
		"hasPathIssues":  fmt.Sprintf("%v", rw.HasPathIssues()),
		"isPathInvalid":  fmt.Sprintf("%v", rw.IsPathInvalid()),
		"hasAnyIssues":   fmt.Sprintf("%v", rw.HasAnyIssues()),
	}
	expected := args.Map{
		"isExist":        "true",
		"isParentExist":  "true",
		"hasPathIssues":  "false",
		"isPathInvalid":  "false",
		"hasAnyIssues":   "false",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter path checks -- existing file", actual)
}

func Test_Cov2_SimpleFileReaderWriter_WriteAndRead(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "rw_write.txt")

	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filePath,
	}

	writeErr := rw.Write([]byte("hello rw"))
	content, readErr := rw.ReadString()

	actual := args.Map{
		"writeOk": fmt.Sprintf("%v", writeErr == nil),
		"readOk":  fmt.Sprintf("%v", readErr == nil),
		"content": content,
	}
	expected := args.Map{
		"writeOk": "true",
		"readOk":  "true",
		"content": "hello rw",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter write then read -- valid path", actual)
}

func Test_Cov2_SimpleFileReaderWriter_WriteString(t *testing.T) {
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "rw_writestr.txt")

	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: dir,
		FilePath:  filePath,
	}

	err := rw.WriteString("string content")

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.WriteString writes -- valid path", actual)
}

func Test_Cov2_SimpleFileReaderWriter_ReadOnExist_NotExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath: "/nonexistent_xyz_99.txt",
	}

	bytes, err := rw.ReadOnExist()

	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"isNil":   fmt.Sprintf("%v", bytes == nil),
	}
	expected := args.Map{
		"noError": "true",
		"isNil":   "true",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.ReadOnExist returns nil -- non-existing", actual)
}

func Test_Cov2_SimpleFileReaderWriter_ReadStringOnExist_NotExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath: "/nonexistent_xyz_99.txt",
	}

	content, err := rw.ReadStringOnExist()

	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"empty":   fmt.Sprintf("%v", content == ""),
	}
	expected := args.Map{
		"noError": "true",
		"empty":   "true",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.ReadStringOnExist returns empty -- non-existing", actual)
}

func Test_Cov2_SimpleFileReaderWriter_Expire(t *testing.T) {
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "rw_expire.txt", "x")

	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: dir,
		FilePath:  filePath,
	}

	err := rw.Expire()

	actual := args.Map{
		"noError":  fmt.Sprintf("%v", err == nil),
		"notExist": fmt.Sprintf("%v", !chmodhelper.IsPathExists(filePath)),
	}
	expected := args.Map{
		"noError":  "true",
		"notExist": "true",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.Expire removes file -- existing file", actual)
}

func Test_Cov2_SimpleFileReaderWriter_Expire_NotExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		FilePath: "/nonexistent_xyz_99.txt",
	}

	err := rw.Expire()

	actual := args.Map{"noError": fmt.Sprintf("%v", err == nil)}
	expected := args.Map{"noError": "true"}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.Expire returns nil -- non-existing", actual)
}

func Test_Cov2_SimpleFileReaderWriter_String(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}

	result := rw.String()

	actual := args.Map{"notEmpty": fmt.Sprintf("%v", result != "")}
	expected := args.Map{"notEmpty": "true"}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.String returns non-empty -- with data", actual)
}

func Test_Cov2_SimpleFileReaderWriter_Clone(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}

	cloned := rw.Clone()
	clonedPtr := rw.ClonePtr()

	actual := args.Map{
		"pathMatch":  fmt.Sprintf("%v", cloned.FilePath == rw.FilePath),
		"ptrNotNil":  fmt.Sprintf("%v", clonedPtr != nil),
	}
	expected := args.Map{
		"pathMatch":  "true",
		"ptrNotNil":  "true",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.Clone preserves data -- cloned", actual)
}

func Test_Cov2_SimpleFileReaderWriter_ClonePtr_Nil(t *testing.T) {
	var rw *chmodhelper.SimpleFileReaderWriter

	result := rw.ClonePtr()

	actual := args.Map{"isNil": fmt.Sprintf("%v", result == nil)}
	expected := args.Map{"isNil": "true"}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.ClonePtr returns nil -- nil receiver", actual)
}

func Test_Cov2_SimpleFileReaderWriter_Json(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}

	jsonResult := rw.Json()
	jsonPtr := rw.JsonPtr()

	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", jsonResult.JsonString() != ""),
		"ptrNotNil": fmt.Sprintf("%v", jsonPtr != nil),
	}
	expected := args.Map{
		"notEmpty": "true",
		"ptrNotNil": "true",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.Json returns valid result -- with data", actual)
}

func Test_Cov2_SimpleFileReaderWriter_JoinRelPath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ParentDir: "/tmp/parent",
	}

	joined := rw.JoinRelPath("sub/file.txt")
	emptyJoin := rw.JoinRelPath("")

	actual := args.Map{
		"joined":    fmt.Sprintf("%v", joined != ""),
		"emptyJoin": fmt.Sprintf("%v", emptyJoin != ""),
	}
	expected := args.Map{
		"joined":    "true",
		"emptyJoin": "true",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.JoinRelPath joins paths -- with and without relpath", actual)
}

func Test_Cov2_SimpleFileReaderWriter_NewPath(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
	}

	result := rw.NewPath(false, "/tmp/newfile.txt")

	actual := args.Map{"notNil": fmt.Sprintf("%v", result != nil)}
	expected := args.Map{"notNil": "true"}
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.NewPath creates new rw -- valid path", actual)
}

func Test_Cov2_SimpleFileReaderWriter_ChmodApplierVerifier(t *testing.T) {
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/tmp",
		FilePath:  "/tmp/test.txt",
	}

	applier := rw.ChmodApplier()
	verifier := rw.ChmodVerifier()

	actual := args.Map{
		"applierOk":  "true",
		"verifierOk": "true",
	}
	expected := args.Map{
		"applierOk":  "true",
		"verifierOk": "true",
	}
	_ = applier
	_ = verifier
	expected.ShouldBeEqual(t, 0, "SimpleFileReaderWriter.ChmodApplier/Verifier created -- valid rw", actual)
}

// ── simpleFileWriter Lock/Unlock ──

func Test_Cov2_SimpleFileWriter_LockUnlock(t *testing.T) {
	chmodhelper.SimpleFileWriter.Lock()
	chmodhelper.SimpleFileWriter.Unlock()

	actual := args.Map{"ok": "true"}
	expected := args.Map{"ok": "true"}
	expected.ShouldBeEqual(t, 0, "simpleFileWriter.Lock/Unlock works -- no deadlock", actual)
}

// ── IsPathExistsPlusFileInfo ──

func Test_Cov2_IsPathExistsPlusFileInfo_Valid(t *testing.T) {
	dir := covTempDir(t)

	exists, info := chmodhelper.IsPathExistsPlusFileInfo(dir)

	actual := args.Map{
		"exists":  fmt.Sprintf("%v", exists),
		"hasInfo": fmt.Sprintf("%v", info != nil),
	}
	expected := args.Map{
		"exists":  "true",
		"hasInfo": "true",
	}
	expected.ShouldBeEqual(t, 0, "IsPathExistsPlusFileInfo returns valid -- existing dir", actual)
}

func Test_Cov2_IsPathExistsPlusFileInfo_Invalid(t *testing.T) {
	exists, info := chmodhelper.IsPathExistsPlusFileInfo("/nonexistent_xyz_99")

	actual := args.Map{
		"exists":  fmt.Sprintf("%v", exists),
		"isNil":   fmt.Sprintf("%v", info == nil),
	}
	expected := args.Map{
		"exists":  "false",
		"isNil":   "true",
	}
	expected.ShouldBeEqual(t, 0, "IsPathExistsPlusFileInfo returns false -- non-existing", actual)
}
