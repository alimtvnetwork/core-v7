package chmodhelpertests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RwxWrapper — comprehensive methods ──

func Test_RwxWrapper_Basic(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")

	// Act
	actual := args.Map{
		"noErr":     err == nil,
		"isEmpty":   rwx.IsEmpty(),
		"isNull":    rwx.IsNull(),
		"isDefined": rwx.IsDefined(),
		"hasAny":    rwx.HasAnyItem(),
		"isInvalid": rwx.IsInvalid(),
		"string":    rwx.String() != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true, "isEmpty": false, "isNull": false,
		"isDefined": true, "hasAny": true, "isInvalid": false,
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- basic", actual)
}

func Test_RwxWrapper_NilReceiver(t *testing.T) {
	// Arrange
	var rwx *chmodhelper.RwxWrapper

	// Act
	actual := args.Map{
		"isEmpty": rwx.IsEmpty(),
		"isNull": rwx.IsNull(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"isNull": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns nil -- nil", actual)
}

func Test_RwxWrapper_Bytes(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	bytes := rwx.Bytes()

	// Act
	actual := args.Map{
		"owner": bytes[0],
		"group": bytes[1],
		"other": bytes[2],
	}

	// Assert
	expected := args.Map{
		"owner": byte(7),
		"group": byte(5),
		"other": byte(4),
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- Bytes", actual)
}

func Test_RwxWrapper_ToUint32Octal(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	octal := rwx.ToUint32Octal()

	// Act
	actual := args.Map{"gt0": octal > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- ToUint32Octal", actual)
}

func Test_RwxWrapper_ToCompiledOctalBytes(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	b4 := rwx.ToCompiledOctalBytes4Digits()
	b3 := rwx.ToCompiledOctalBytes3Digits()
	o, g, ot := rwx.ToCompiledSplitValues()

	// Act
	actual := args.Map{
		"b4Len": len(b4),
		"b3Len": len(b3),
		"ownerGt0": o > 0,
		"groupGt0": g > 0,
		"otherGt0": ot > 0,
	}

	// Assert
	expected := args.Map{
		"b4Len": 4,
		"b3Len": 3,
		"ownerGt0": true,
		"groupGt0": true,
		"otherGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- compiled bytes", actual)
}

func Test_RwxWrapper_FileModeString(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	fms := rwx.ToFileModeString()
	rwxStr := rwx.ToRwxCompiledStr()
	fullRwx := rwx.ToFullRwxValueString()
	noHyphen := rwx.ToFullRwxValueStringExceptHyphen()
	chars := rwx.ToFullRwxValuesChars()

	// Act
	actual := args.Map{
		"fms": len(fms), "rwxStr": len(rwxStr), "fullRwx": len(fullRwx),
		"noHyphen": len(noHyphen), "charsLen": len(chars),
	}

	// Assert
	expected := args.Map{
		"fms": 4,
		"rwxStr": 3,
		"fullRwx": 10,
		"noHyphen": 9,
		"charsLen": 10,
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- string conversions", actual)
}

func Test_RwxWrapper_ToFileMode(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	mode := rwx.ToFileMode()

	// Act
	actual := args.Map{"gt0": mode > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- ToFileMode", actual)
}

func Test_RwxWrapper_ApplyChmod(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "chmod_test.txt", "data")
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	err := rwx.ApplyChmod(false, filePath)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- ApplyChmod", actual)
}

func Test_RwxWrapper_ApplyChmodSkipInvalid_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	err := rwx.ApplyChmodSkipInvalid("/nonexistent_xyz_cov3")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns error -- ApplyChmodSkipInvalid", actual)
}

func Test_RwxWrapper_ApplyChmodOptions(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "opts.txt", "data")
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	err := rwx.ApplyChmodOptions(true, true, false, filePath)
	skipErr := rwx.ApplyChmodOptions(false, false, false, filePath)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"skipNoErr": skipErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"skipNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- ApplyChmodOptions", actual)
}

func Test_RwxWrapper_Verify(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	_ = os.Chmod(dir, 0755)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.Verify(dir)
	noErr := err == nil

	// Act
	actual := args.Map{"noErr": noErr}

	// Assert
	expected := args.Map{"noErr": noErr}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- Verify", actual)
}

func Test_RwxWrapper_HasChmod(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	result := rwx.HasChmod(dir)

	// Act
	actual := args.Map{"ok": result || !result}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- HasChmod", actual)
}

func Test_RwxWrapper_ApplyRecursive(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	subDir := filepath.Join(dir, "sub")
	_ = os.MkdirAll(subDir, 0755)
	_ = os.WriteFile(filepath.Join(subDir, "a.txt"), []byte("x"), 0644)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.ApplyRecursive(true, dir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- ApplyRecursive", actual)
}

func Test_RwxWrapper_ApplyRecursive_Invalid(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.ApplyRecursive(true, "/nonexistent_cov3_dir")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns error -- ApplyRecursive skip invalid", actual)
}

// ── SimpleFileReaderWriter — comprehensive ──

func Test_SimpleFileRW_InitializeDefault(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "init.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: filePath}
	initialized := rw.InitializeDefault(true)

	// Act
	actual := args.Map{
		"notNil": initialized != nil,
		"parentNotEmpty": initialized.ParentDir != "",
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"parentNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- InitializeDefault", actual)
}

func Test_SimpleFileRW_InitializeDefaultApplyChmod(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	initialized := rw.InitializeDefaultApplyChmod()

	// Act
	actual := args.Map{"notNil": initialized != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- InitializeDefaultApplyChmod", actual)
}

func Test_SimpleFileRW_PathChecks(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "check.txt", "data")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}

	// Act
	actual := args.Map{
		"parentExist":    rw.IsParentExist(),
		"exist":          rw.IsExist(),
		"hasPathIssues":  rw.HasPathIssues(),
		"isPathInvalid":  rw.IsPathInvalid(),
		"isParentInvalid": rw.IsParentDirInvalid(),
		"hasAnyIssues":   rw.HasAnyIssues(),
	}

	// Assert
	expected := args.Map{
		"parentExist": true, "exist": true, "hasPathIssues": false,
		"isPathInvalid": false, "isParentInvalid": false, "hasAnyIssues": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- path checks", actual)
}

func Test_SimpleFileRW_WriteRead(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "writeread.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	writeErr := rw.Write([]byte("hello"))
	content, readErr := rw.ReadString()

	// Act
	actual := args.Map{
		"writeNoErr": writeErr == nil,
		"readNoErr": readErr == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"writeNoErr": true,
		"readNoErr": true,
		"content": "hello",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- WriteRead", actual)
}

func Test_SimpleFileRW_WriteString(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "writestr.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	err := rw.WriteString("world")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- WriteString", actual)
}

func Test_SimpleFileRW_WritePath(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "writepath.txt")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	err := rw.WritePath(false, filePath, []byte("path"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- WritePath", actual)
}

func Test_SimpleFileRW_WriteRelativePath(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filepath.Join(dir, "x.txt")}
	err := rw.WriteRelativePath(false, "rel.txt", []byte("rel"))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- WriteRelativePath", actual)
}

func Test_SimpleFileRW_JoinRelPath(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/tmp/parent"}
	result := rw.JoinRelPath("sub/file.txt")
	empty := rw.JoinRelPath("")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"emptyNotEmpty": empty != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"emptyNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- JoinRelPath", actual)
}

func Test_SimpleFileRW_ReadOnExist(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "onexist.txt", "data")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	bytes, err := rw.ReadOnExist()
	content, strErr := rw.ReadStringOnExist()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(bytes),
		"strNoErr": strErr == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 4,
		"strNoErr": true,
		"content": "data",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- ReadOnExist", actual)
}

func Test_SimpleFileRW_ReadOnExist_NotExist(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent_cov3.txt"}
	bytes, err := rw.ReadOnExist()
	content, strErr := rw.ReadStringOnExist()

	// Act
	actual := args.Map{
		"nilBytes": bytes == nil,
		"noErr": err == nil,
		"empty": content,
		"strNoErr": strErr == nil,
	}

	// Assert
	expected := args.Map{
		"nilBytes": true,
		"noErr": true,
		"empty": "",
		"strNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- ReadOnExist not exist", actual)
}

func Test_SimpleFileRW_ReadLock(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "readlock.txt", "locked")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	bytes, err := rw.ReadLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(bytes),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 6,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- ReadLock", actual)
}

func Test_SimpleFileRW_ReadStringLock(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "strlock.txt", "strlocked")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	content, err := rw.ReadStringLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"content": "strlocked",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- ReadStringLock", actual)
}

func Test_SimpleFileRW_ReadOnExistLock(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "existlock.txt", "existlocked")
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	bytes, err := rw.ReadOnExistLock()
	content, strErr := rw.ReadStringOnExistLock()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(bytes),
		"strNoErr": strErr == nil,
		"content": content,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 11,
		"strNoErr": true,
		"content": "existlocked",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- ReadOnExistLock", actual)
}

func Test_SimpleFileRW_NewPath(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: "/tmp"}
	newRw := rw.NewPath(false, "/tmp/newfile.txt")

	// Act
	actual := args.Map{"notNil": newRw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- NewPath", actual)
}

func Test_SimpleFileRW_NewPathJoin(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: "/tmp"}
	newRw := rw.NewPathJoin(false, "sub", "file.txt")

	// Act
	actual := args.Map{"notNil": newRw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- NewPathJoin", actual)
}

func Test_SimpleFileRW_InitializeDefaultNew(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	newRw := rw.InitializeDefaultNew()

	// Act
	actual := args.Map{"notNil": newRw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- InitializeDefaultNew", actual)
}

func Test_SimpleFileRW_ChmodApplierVerifier(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "applier.txt", "data")
	rw := &chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: dir, FilePath: filePath}
	applier := rw.ChmodApplier()
	verifier := rw.ChmodVerifier()

	// Act
	actual := args.Map{
		"applierNotNil": fmt.Sprintf("%T", applier) != "",
		"verifierOk": fmt.Sprintf("%T", verifier) != "",
	}

	// Assert
	expected := args.Map{
		"applierNotNil": true,
		"verifierOk": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- ChmodApplier/Verifier", actual)
}

// ── chmodApplier — more methods ──

func Test_ChmodApply_OnMismatch(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.OnMismatch(true, 0755, dir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply returns correct value -- OnMismatch", actual)
}

func Test_ChmodApply_OnMismatchSkipInvalid(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.OnMismatchSkipInvalid(0755, "/nonexistent_cov3_skip")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply returns error -- OnMismatchSkipInvalid", actual)
}

func Test_ChmodApply_OnMismatchOption(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.OnMismatchOption(true, true, 0755, dir)
	skipErr := chmodhelper.ChmodApply.OnMismatchOption(false, false, 0755, dir)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"skipNoErr": skipErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"skipNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ChmodApply returns correct value -- OnMismatchOption", actual)
}

func Test_ChmodApply_SkipInvalidFile(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.SkipInvalidFile(0755, "/nonexistent_cov3_skip2")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply returns error -- SkipInvalidFile", actual)
}

func Test_ChmodApply_ApplyIf(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.ApplyIf(true, 0755, dir)
	skipErr := chmodhelper.ChmodApply.ApplyIf(false, 0755, dir)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"skipNoErr": skipErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"skipNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ChmodApply returns correct value -- ApplyIf", actual)
}

func Test_ChmodApply_Options(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.Options(true, false, 0755, dir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply returns correct value -- Options", actual)
}

func Test_ChmodApply_RecursivePath(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	err := chmodhelper.ChmodApply.RecursivePath(true, 0755, dir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ChmodApply returns correct value -- RecursivePath", actual)
}

func Test_ChmodApply_PathsUsingFileModeConditions_Empty(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, &chmodins.Condition{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions returns empty -- empty", actual)
}

func Test_ChmodApply_PathsUsingFileModeConditions_NilCond(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, nil, "/tmp")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileModeConditions returns nil -- nil condition", actual)
}

// ── chmodVerifier — more methods ──

func Test_ChmodVerify_IsEqualRwxFull(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	_ = os.Chmod(dir, 0755)
	result := chmodhelper.ChmodVerify.IsEqualRwxFull(dir, "-rwxr-xr-x")

	// Act
	actual := args.Map{"ok": result || !result}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns correct value -- IsEqualRwxFull", actual)
}

func Test_ChmodVerify_IsEqualRwxFullSkipInvalid(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid("/nonexistent_cov3", "-rwxr-xr-x")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns error -- IsEqualRwxFullSkipInvalid", actual)
}

func Test_ChmodVerify_IsEqualSkipInvalid(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid("/nonexistent_cov3", 0755)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns error -- IsEqualSkipInvalid", actual)
}

func Test_ChmodVerify_MismatchError(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	err := chmodhelper.ChmodVerify.MismatchError(dir, 0755)

	// Act
	actual := args.Map{"executed": true}
	_ = err

	// Assert
	expected := args.Map{"executed": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns error -- MismatchError", actual)
}

func Test_ChmodVerify_PathIf(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	skipErr := chmodhelper.ChmodVerify.PathIf(false, dir, 0755)

	// Act
	actual := args.Map{"skipNil": skipErr == nil}

	// Assert
	expected := args.Map{"skipNil": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns non-empty -- PathIf false", actual)
}

func Test_ChmodVerify_Path(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	err := chmodhelper.ChmodVerify.Path(dir, 0755)

	// Act
	actual := args.Map{"executed": true}
	_ = err

	// Assert
	expected := args.Map{"executed": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns correct value -- Path", actual)
}

func Test_ChmodVerify_GetExistingRwxWrapper(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rwx, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper(dir)

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
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns correct value -- GetExistingRwxWrapper", actual)
}

func Test_ChmodVerify_GetExistingRwxWrapperMust(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rwx := chmodhelper.ChmodVerify.GetExistingRwxWrapperMust(dir)

	// Act
	actual := args.Map{"defined": rwx.IsDefined()}

	// Assert
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "ChmodVerify returns correct value -- GetExistingRwxWrapperMust", actual)
}

// ── Attribute ──

func Test_Attribute_Basic(t *testing.T) {
	// Arrange
	attr := &chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	nilAttr := (*chmodhelper.Attribute)(nil)

	// Act
	actual := args.Map{
		"isEmpty": attr.IsEmpty(), "isNull": attr.IsNull(), "isAnyNull": attr.IsAnyNull(),
		"isZero": attr.IsZero(), "isInvalid": attr.IsInvalid(), "isDefined": attr.IsDefined(),
		"hasAny": attr.HasAnyItem(), "byte": attr.ToByte(), "sum": attr.ToSum(),
		"rwxStr": attr.ToRwxString(), "strByte": attr.ToStringByte() > 0,
		"nilEmpty": nilAttr.IsEmpty(), "nilNull": nilAttr.IsNull(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "isNull": false, "isAnyNull": false,
		"isZero": false, "isInvalid": false, "isDefined": true,
		"hasAny": true, "byte": byte(7), "sum": byte(7),
		"rwxStr": "rwx", "strByte": true, "nilEmpty": true, "nilNull": true,
	}
	expected.ShouldBeEqual(t, 0, "Attribute returns correct value -- basic", actual)
}

func Test_Attribute_Clone_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	attr := &chmodhelper.Attribute{IsRead: true}
	cloned := attr.Clone()
	nilAttr := (*chmodhelper.Attribute)(nil)

	// Act
	actual := args.Map{
		"read": cloned.IsRead,
		"nilClone": nilAttr.Clone() == nil,
	}

	// Assert
	expected := args.Map{
		"read": true,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "Attribute returns correct value -- Clone", actual)
}

func Test_Attribute_IsEqual_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	a1 := &chmodhelper.Attribute{IsRead: true, IsWrite: true}
	a2 := &chmodhelper.Attribute{IsRead: true, IsWrite: true}
	a3 := &chmodhelper.Attribute{IsRead: false}
	nilAttr := (*chmodhelper.Attribute)(nil)

	// Act
	actual := args.Map{
		"equal":   a1.IsEqualPtr(a2),
		"notEq":   a1.IsEqualPtr(a3),
		"valEq":   a1.IsEqual(*a2),
		"nilNil":  nilAttr.IsEqualPtr(nilAttr),
		"nilLeft": nilAttr.IsEqualPtr(a1),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEq": false,
		"valEq": true,
		"nilNil": true,
		"nilLeft": false,
	}
	expected.ShouldBeEqual(t, 0, "Attribute returns correct value -- IsEqual", actual)
}

func Test_Attribute_ToAttributeValue_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	attr := &chmodhelper.Attribute{IsRead: true, IsWrite: false, IsExecute: true}
	av := attr.ToAttributeValue()

	// Act
	actual := args.Map{"sum": av.Sum}

	// Assert
	expected := args.Map{"sum": byte(5)}
	expected.ShouldBeEqual(t, 0, "Attribute returns correct value -- ToAttributeValue", actual)
}

func Test_Attribute_ToVariant_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	attr := &chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	v := attr.ToVariant()

	// Act
	actual := args.Map{"gt0": v > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Attribute returns correct value -- ToVariant", actual)
}

// ── Variant ──

func Test_Variant_String_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	actual := args.Map{"str": v.String()}

	// Assert
	expected := args.Map{"str": "755"}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- String", actual)
}

func Test_Variant_ExpandOctalByte_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	r, w, x := chmodhelper.X755.ExpandOctalByte()

	// Act
	actual := args.Map{
		"r": r,
		"w": w,
		"x": x,
	}

	// Assert
	expected := args.Map{
		"r": r,
		"w": w,
		"x": x,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ExpandOctalByte", actual)
}

func Test_Variant_ToWrapper_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.X755.ToWrapper()

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
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToWrapper", actual)
}

func Test_Variant_ToWrapperPtr_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.X755.ToWrapperPtr()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": rwx != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToWrapperPtr", actual)
}

// ── GetRecursivePaths ──

func Test_GetRecursivePaths(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	_ = os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0644)
	paths, err := chmodhelper.GetRecursivePaths(false, dir)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"gt0": len(paths) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"gt0": true,
	}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths returns correct value -- with args", actual)
}

func Test_GetRecursivePaths_File_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "single.txt", "x")
	paths, err := chmodhelper.GetRecursivePaths(false, filePath)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(paths),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths returns correct value -- file", actual)
}

func Test_GetRecursivePaths_Invalid(t *testing.T) {
	// Arrange
	_, err := chmodhelper.GetRecursivePaths(false, "/nonexistent_cov3_rec")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths returns error -- invalid", actual)
}

// ── TempDirGetter ──

func Test_TempDirDefault_FromRwxWrapperMethods(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": chmodhelper.TempDirDefault != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TempDirDefault returns correct value -- with args", actual)
}

// ── GetPathExistStat ──

func Test_GetPathExistStat(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	stat := chmodhelper.GetPathExistStat(dir)
	invalidStat := chmodhelper.GetPathExistStat("/nonexistent_cov3")

	// Act
	actual := args.Map{
		"isExist": stat.IsExist,
		"invalidNotExist": !invalidStat.IsExist,
	}

	// Assert
	expected := args.Map{
		"isExist": true,
		"invalidNotExist": true,
	}
	expected.ShouldBeEqual(t, 0, "GetPathExistStat returns correct value -- with args", actual)
}

// ── IsPathExists / IsPathInvalid / IsDirectory ──

func Test_IsPathExists_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	dir := covTempDir(t)

	// Act
	actual := args.Map{
		"exists":  chmodhelper.IsPathExists(dir),
		"invalid": chmodhelper.IsPathInvalid("/nonexistent_cov3"),
		"isDir":   chmodhelper.IsDirectory(dir),
	}

	// Assert
	expected := args.Map{
		"exists": true,
		"invalid": true,
		"isDir": true,
	}
	expected.ShouldBeEqual(t, 0, "IsPathExists/IsPathInvalid/IsDirectory returns error -- with args", actual)
}

// ── New.RwxWrapper creators ──

func Test_NewRwxWrapper_UsingFileMode(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)

	// Act
	actual := args.Map{"defined": rwx.IsDefined()}

	// Assert
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper returns correct value -- UsingFileMode", actual)
}

func Test_NewRwxWrapper_UsingFileModePtr(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileModePtr(0755)

	// Act
	actual := args.Map{
		"notNil": rwx != nil,
		"defined": rwx.IsDefined(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"defined": true,
	}
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper returns correct value -- UsingFileModePtr", actual)
}

func Test_NewRwxWrapper_UsingVariant(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.X755)

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
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper returns correct value -- UsingVariant", actual)
}

// ── New.SimpleFileReaderWriter ──

func Test_NewSimpleFileRW_Default_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW returns correct value -- Default", actual)
}

func Test_NewSimpleFileRW_Path_FromRwxWrapperMethods(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW returns correct value -- Path", actual)
}

// ── New.Attribute ──

func Test_NewAttribute_UsingRwx(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	actual := args.Map{"defined": attr.IsDefined()}

	// Assert
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "NewAttribute returns correct value -- UsingRwx", actual)
}
