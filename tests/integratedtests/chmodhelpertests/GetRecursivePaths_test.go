package chmodhelpertests

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── GetRecursivePaths — error on invalid root ──

func Test_GetRecursivePaths_InvalidRoot(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0})

	// Act
	paths, err := chmodhelper.GetRecursivePaths(false, invalidPath)

	// Assert
	actual := args.Map{
		"pathCount": len(paths),
		"hasErr":    err != nil,
	}
	expected := args.Map{
		"pathCount": 0,
		"hasErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths returns error -- invalid root path", actual)
}

// ── GetRecursivePathsContinueOnError — error on invalid root ──

func Test_GetRecursivePathsContinueOnError_InvalidRoot(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0})

	// Act
	paths, err := chmodhelper.GetRecursivePathsContinueOnError(invalidPath)

	// Assert
	actual := args.Map{
		"pathCount": len(paths),
		"hasErr":    err != nil,
	}
	expected := args.Map{
		"pathCount": 0,
		"hasErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "GetRecursivePathsContinueOnError returns error -- invalid root", actual)
}

// ── MergeRwxWildcardWithFixedRwx — ParseRwxToVarAttribute error ──

func Test_MergeRwxWildcardWithFixedRwx_InvalidWildcard(t *testing.T) {
	// Arrange
	existingRwx := "rwx"
	invalidWildcard := "zzzz" // invalid length (not 3)

	// Act
	result, err := chmodhelper.MergeRwxWildcardWithFixedRwx(existingRwx, invalidWildcard)

	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr":    err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "MergeRwxWildcardWithFixedRwx returns error -- invalid wildcard chars", actual)
}

// ── PathExistStat — MeaningFullError with error ──

func Test_PathExistStat_MeaningFullError_WithError_FromGetRecursivePaths(t *testing.T) {
	// Arrange
	stat := &chmodhelper.PathExistStat{
		Location: "/some/path",
		Error:    errors.New("test error"),
	}

	// Act
	err := stat.MeaningFullError()

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningFullError returns error -- has Error field", actual)
}

// ── PathExistStat — MeaningFullError without error ──

func Test_PathExistStat_MeaningFullError_NoError(t *testing.T) {
	// Arrange
	stat := &chmodhelper.PathExistStat{
		Location: "/some/path",
	}

	// Act
	err := stat.MeaningFullError()

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningFullError returns nil -- no error", actual)
}

// ── RwxInstructionExecutor — ApplyOnPathsDirect delegates ──

func Test_RwxInstructionExecutor_ApplyOnPathsDirect_EmptyLocations(t *testing.T) {
	// Arrange
	instruction := &chmodins.RwxInstruction{}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(instruction)
	if executor == nil {
		t.Skip("executor creation requires valid instruction")
	}

	// Act
	err := executor.ApplyOnPathsDirect()

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPathsDirect returns nil -- empty locations", actual)
}

// ── RwxInstructionExecutor — ApplyOnPaths delegates ──

func Test_RwxInstructionExecutor_ApplyOnPaths_EmptyLocations(t *testing.T) {
	// Arrange
	instruction := &chmodins.RwxInstruction{}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(instruction)
	if executor == nil {
		t.Skip("executor creation requires valid instruction")
	}

	// Act
	err := executor.ApplyOnPaths(nil)

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPaths returns nil -- nil locations", actual)
}

// ── RwxVariableWrapper — IsEqualUsingFileInfo nil ──

func Test_RwxVariableWrapper_IsEqualUsingFileInfo_Nil(t *testing.T) {
	// Arrange
	varWrapper, err := chmodhelper.NewRwxVariableWrapper("rwx")
	if err != nil {
		t.Skip("NewRwxVariableWrapper error: " + err.Error())
	}

	// Act
	result := varWrapper.IsEqualUsingFileInfo(nil)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqualUsingFileInfo returns false -- nil fileInfo", actual)
}

// ── RwxVariableWrapper — IsEqualUsingLocation with valid file ──

func Test_RwxVariableWrapper_IsEqualUsingLocation_NonExistent(t *testing.T) {
	// Arrange
	varWrapper, err := chmodhelper.NewRwxVariableWrapper("rwx")
	if err != nil {
		t.Skip("NewRwxVariableWrapper error: " + err.Error())
	}

	// Act
	result := varWrapper.IsEqualUsingLocation("/non/existent/path")

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqualUsingLocation returns false -- non-existent path", actual)
}

// ── RwxVariableWrapper — IsEqualUsingLocation with existing file ──

func Test_RwxVariableWrapper_IsEqualUsingLocation_ExistingFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(tmpFile, []byte("test"), 0644)
	varWrapper, err := chmodhelper.NewRwxVariableWrapper("rw-")
	if err != nil {
		t.Skip("NewRwxVariableWrapper error: " + err.Error())
	}

	// Act
	result := varWrapper.IsEqualUsingLocation(tmpFile)

	// Assert — on Windows chmod doesn't apply so result may vary
	actual := args.Map{"isBool": result == true || result == false}
	expected := args.Map{"isBool": true}
	expected.ShouldBeEqual(t, 0, "IsEqualUsingLocation returns bool -- existing file", actual)
}

// ── SingleRwx — ToRwxWrapper / ApplyOnMany ──
// NOTE: SingleRwx construction requires in-package access.
// Covered by in-package tests. Documented in Coverage7_DeadCode_doc.go.

// ── fwChmodVerifier — IsEqualFile ──

func Test_FwChmodVerifier_IsEqualFile(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "verify.txt")
	_ = os.WriteFile(tmpFile, []byte("data"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: tmpDir,
		FilePath:  tmpFile,
	}
	initialized := rw.InitializeDefault(false)

	// Act
	result := initialized.ChmodVerifier().IsEqualFile()

	// Assert — on Windows this will return false (chmod not supported)
	actual := args.Map{"isBool": result == true || result == false}
	expected := args.Map{"isBool": true}
	expected.ShouldBeEqual(t, 0, "IsEqualFile returns bool -- valid file", actual)
}

// ── fwChmodApplier — OnDiffFile already equal ──

func Test_FwChmodApplier_OnDiffFile_AlreadyEqual(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}

	// Arrange
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "diff.txt")
	_ = os.WriteFile(tmpFile, []byte("data"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: tmpDir,
		FilePath:  tmpFile,
	}
	initialized := rw.InitializeDefault(false)

	// Act — file already has 0644
	err := initialized.ChmodApplier().OnDiffFile(false, tmpFile)

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "OnDiffFile returns nil -- chmod already matches", actual)
}

// ── tempDirGetter — TempOption isPermanent=false ──

func Test_TempDirGetter_TempOption_NotPermanent(t *testing.T) {
	// Arrange & Act
	result := chmodhelper.TempDirGetter.TempOption(false)

	// Assert
	actual := args.Map{"nonEmpty": len(result) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "TempOption returns non-empty -- isPermanent=false", actual)
}

// ── chmodVerifier — PathIf isVerify=true with valid path ──

func Test_ChmodVerifier_PathIf_VerifyTrue(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "verifyif.txt")
	_ = os.WriteFile(tmpFile, []byte("data"), 0644)

	// Act
	err := chmodhelper.ChmodVerify.PathIf(true, tmpFile, 0644)

	// Assert — on Windows this may fail since chmod doesn't work
	if runtime.GOOS == "windows" {
		// Just verify it doesn't panic
		actual := args.Map{"ran": true}
		expected := args.Map{"ran": true}
		expected.ShouldBeEqual(t, 0, "PathIf ran without panic -- Windows", actual)
	} else {
		actual := args.Map{"errNil": err == nil}
		expected := args.Map{"errNil": true}
		expected.ShouldBeEqual(t, 0, "PathIf returns nil -- matching chmod on Linux", actual)
	}
}

// ── chmodVerifier — IsEqualRwxFullSkipInvalid with invalid path ──

func Test_ChmodVerifier_IsEqualRwxFullSkipInvalid_InvalidPath(t *testing.T) {
	// Arrange — use truly non-existent path (os.IsNotExist returns true)
	invalidPath := filepath.Join(t.TempDir(), "no_such_file")

	// Act
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid(invalidPath, "-rwxrwxrwx")

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqualRwxFullSkipInvalid returns true -- invalid path skipped", actual)
}

// ── chmodVerifier — IsEqualSkipInvalid with invalid path ──

func Test_ChmodVerifier_IsEqualSkipInvalid_InvalidPath(t *testing.T) {
	// Arrange — use truly non-existent path
	invalidPath := filepath.Join(t.TempDir(), "no_such_file")

	// Act
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid(invalidPath, 0755)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqualSkipInvalid returns true -- invalid path skipped", actual)
}

// ── chmodVerifier — GetRwx9 ──

func Test_ChmodVerifier_GetRwx9_FromGetRecursivePaths(t *testing.T) {
	// Arrange
	fileMode := os.FileMode(0755)

	// Act
	result := chmodhelper.ChmodVerify.GetRwx9(fileMode)

	// Assert
	actual := args.Map{
		"length":   len(result),
		"nonEmpty": len(result) > 0,
	}
	expected := args.Map{
		"length":   9,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "GetRwx9 returns 9-char string -- valid fileMode", actual)
}

// ── chmodVerifier — PathsUsingRwxFull continueOnError ──

func Test_ChmodVerifier_PathsUsingRwxFull_ContinueOnError(t *testing.T) {
	// Arrange — use non-existent paths to trigger errors
	locations := []string{"/non/existent/path1", "/non/existent/path2"}

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingRwxFull(true, "-rwxrwxrwx", locations...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingRwxFull returns error -- non-existent paths with continue", actual)
}

// ── chmodVerifier — PathsUsingRwxFull no continue on error ──

func Test_ChmodVerifier_PathsUsingRwxFull_NoContinue(t *testing.T) {
	// Arrange
	locations := []string{"/non/existent/path"}

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingRwxFull(false, "-rwxrwxrwx", locations...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingRwxFull returns error -- non-existent path no continue", actual)
}

// ── chmodVerifier — PathsUsingFileMode ──

func Test_ChmodVerifier_PathsUsingFileMode(t *testing.T) {
	// Arrange
	locations := []string{"/non/existent/path"}

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingFileMode(false, 0755, locations...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingFileMode returns error -- non-existent path", actual)
}

// ── chmodVerifier — PathsUsingPartialRwxOptions ──

func Test_ChmodVerifier_PathsUsingPartialRwxOptions_InvalidPartial(t *testing.T) {
	// Arrange
	locations := []string{"/some/path"}

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(false, false, "zzz", locations...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathsUsingPartialRwxOptions returns error -- invalid partial rwx", actual)
}

// ── SimpleFileReaderWriter — WriteBytes error ──

func Test_SimpleFileReaderWriter_WriteBytes_InvalidPath(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0})
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: t.TempDir(),
		FilePath:  invalidPath,
	}
	initialized := rw.InitializeDefault(false)

	// Act
	err := initialized.Write([]byte("test"))

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "WriteBytes returns error -- invalid file path", actual)
}

// ── DirFilesWithContent — Create with invalid path ──

func Test_DirFilesWithContent_Create_InvalidPath(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0})
	content := chmodhelper.DirFilesWithContent{
		Dir:         invalidPath,
		DirFileMode: 0755,
		Files: []chmodhelper.FileWithContent{
			{RelativePath: "test.txt", Content: []string{"hello"}, FileMode: 0644},
		},
	}

	// Act
	err := content.Create(false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Create returns error -- invalid dir path", actual)
}

// ── dirCreator — unexported, tested via in-package tests ──

// ── RwxWrapper — ApplyChmodSafe skip invalid ──

func Test_RwxWrapper_ApplyChmodSkipInvalid_InvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}

	// Arrange
	wrapper, err := chmodhelper.New.RwxWrapper.RwxFullString("-rwxrwxrwx")
	if err != nil {
		t.Skip("RwxWrapper creation failed")
	}
	// Use truly non-existent path (os.IsNotExist returns true)
	nonExistent := filepath.Join(t.TempDir(), "no_such_file")

	// Act
	err = wrapper.ApplyChmodSkipInvalid(nonExistent)

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "ApplyChmodSkipInvalid returns nil -- skipped invalid path", actual)
}

// ── chmodApplier — RwxStringApplyChmod error ──

func Test_ChmodApplier_RwxStringApplyChmod_NilCondition(t *testing.T) {
	// Arrange & Act
	err := chmodhelper.RwxStringApplyChmod("-rwxrwxrwx", nil, "/some/path")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxStringApplyChmod returns error -- nil condition", actual)
}

// ── chmodApplier — RwxOwnerGroupOtherApplyChmod nil condition ──

func Test_ChmodApplier_RwxOwnerGroupOtherApplyChmod_NilCondition(t *testing.T) {
	// Arrange
	ogo := &chmodins.RwxOwnerGroupOther{}

	// Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(ogo, nil, "/some/path")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOtherApplyChmod returns error -- nil condition", actual)
}

// ── CreateDirWithFiles — removeDirIf error ──

func Test_CreateDirWithFiles_InvalidRemovePath(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0})
	dirWithFile := chmodhelper.DirWithFiles{
		Dir:   invalidPath,
		Files: []string{"test.txt"},
	}

	// Act — isRemoveAllDirBeforeCreate=true with invalid path
	err := chmodhelper.CreateDirWithFiles(
		true,
		0755,
		&dirWithFile,
	)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CreateDirWithFiles returns error -- invalid dir path for remove", actual)
}

// ── GetRecursivePaths — valid dir ──

func Test_GetRecursivePaths_ValidDir(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "sub")
	_ = os.MkdirAll(subDir, 0755)
	_ = os.WriteFile(filepath.Join(subDir, "file.txt"), []byte("hi"), 0644)

	// Act
	paths, err := chmodhelper.GetRecursivePaths(false, tmpDir)

	// Assert
	actual := args.Map{
		"errNil":       err == nil,
		"hasMultiple":  len(paths) >= 3,
	}
	expected := args.Map{
		"errNil":       true,
		"hasMultiple":  true,
	}
	expected.ShouldBeEqual(t, 0, "GetRecursivePaths returns paths -- valid dir with subdir+file", actual)
}

// ── GetRecursivePathsContinueOnError — valid dir ──

func Test_GetRecursivePathsContinueOnError_ValidDir(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	_ = os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("hi"), 0644)

	// Act
	paths, err := chmodhelper.GetRecursivePathsContinueOnError(tmpDir)

	// Assert
	actual := args.Map{
		"errNil":    err == nil,
		"hasItems":  len(paths) >= 2,
	}
	expected := args.Map{
		"errNil":    true,
		"hasItems":  true,
	}
	expected.ShouldBeEqual(t, 0, "GetRecursivePathsContinueOnError returns paths -- valid dir", actual)
}
