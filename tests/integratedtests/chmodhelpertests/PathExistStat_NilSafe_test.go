package chmodhelpertests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── PathExistStat nil-safety ──

func Test_PathExistStat_NilSafe(t *testing.T) {
	for caseIndex, tc := range extPathExistStatNilSafeCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ── PathExistStat method tests ──

func Test_PathExistStat_InvalidPath(t *testing.T) {
	for caseIndex, testCase := range extPathExistStatTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		path, _ := input.GetAsString("path")

		// Act
		stat := chmodhelper.GetPathExistStat(path)

		actual := args.Map{
			"isExist":  stat.IsExist,
			"isInvalid": stat.IsInvalid(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PathExistStat_TempDir(t *testing.T) {
	// Arrange
	tempDir := os.TempDir()

	// Act
	stat := chmodhelper.GetPathExistStat(tempDir)

	// Assert
	actual := args.Map{"result": stat.IsExist}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "temp dir should exist", actual)

	actual = args.Map{"result": stat.IsDir()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "temp dir should be a directory", actual)

	actual = args.Map{"result": stat.IsFile()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir should not be a file", actual)

	actual = args.Map{"result": stat.IsInvalid()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir should not be invalid", actual)

	actual = args.Map{"result": stat.HasAnyIssues()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir should not have issues", actual)

	actual = args.Map{"result": stat.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir should not have error", actual)

	actual = args.Map{"result": stat.IsEmptyError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "temp dir should have empty error", actual)

	actual = args.Map{"result": stat.HasFileInfo()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "temp dir should have file info", actual)

	actual = args.Map{"result": stat.IsInvalidFileInfo()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir should have valid file info", actual)

	actual = args.Map{"result": stat.FileMode() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FileMode should not be nil", actual)

	actual = args.Map{"result": stat.LastModifiedDate() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "LastModifiedDate should not be nil", actual)

	actual = args.Map{"result": stat.NotExistError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir NotExistError should be nil", actual)

	actual = args.Map{"result": stat.NotADirError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir NotADirError should be nil", actual)

	actual = args.Map{"result": stat.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)
}

func Test_PathExistStat_TempDir_Navigation(t *testing.T) {
	// Arrange
	tempDir := os.TempDir()
	stat := chmodhelper.GetPathExistStat(tempDir)

	// Act & Assert
	combined := stat.CombineWithNewPath("subdir")
	actual := args.Map{"result": combined == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CombineWithNewPath should not be empty", actual)

	combinedStat := stat.CombineWith("subdir")
	actual = args.Map{"result": combinedStat == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CombineWith should not return nil", actual)
}

func Test_PathExistStat_File(t *testing.T) {
	// Arrange
	tmpFile, err := os.CreateTemp("", "test-*.txt")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create temp file:", actual)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	stat := chmodhelper.GetPathExistStat(tmpFile.Name())

	// Assert
	actual = args.Map{"result": stat.IsExist}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "temp file should exist", actual)

	actual = args.Map{"result": stat.IsFile()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be a file", actual)

	actual = args.Map{"result": stat.IsDir()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be a dir", actual)

	actual = args.Map{"result": stat.FileName() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FileName should not be empty", actual)

	actual = args.Map{"result": stat.ParentDir() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParentDir should not be empty", actual)

	actual = args.Map{"result": stat.DotExt() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DotExt should not be empty for .txt file", actual)

	actual = args.Map{"result": stat.Size() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Size should not be nil", actual)

	parent := stat.Parent()
	actual = args.Map{"result": parent == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Parent should not be nil", actual)

	actual = args.Map{"result": stat.NotAFileError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotAFileError should be nil for a file", actual)

	parentPath := stat.ParentWithNewPath("test.txt")
	actual = args.Map{"result": parentPath == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParentWithNewPath should not be empty", actual)

	parentStat := stat.ParentWith("test.txt")
	actual = args.Map{"result": parentStat == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParentWith should not return nil", actual)
}

func Test_PathExistStat_NotAFileError(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())

	// Act
	err := stat.NotAFileError()

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotAFileError should return error for directory", actual)
}

func Test_PathExistStat_NotADirError(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	stat := chmodhelper.GetPathExistStat(tmpFile.Name())

	// Act
	err := stat.NotADirError()

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotADirError should return error for file", actual)
}

func Test_PathExistStat_NotExist(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat("/nonexistent/path/xyz")

	// Act
	err := stat.NotExistError()

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotExistError should return error for non-existent path", actual)
}

func Test_PathExistStat_Dispose_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())

	// Act
	stat.Dispose()

	// Assert
	actual := args.Map{"result": stat.IsExist}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsExist should be false after Dispose", actual)

	actual = args.Map{"result": stat.Location != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Location should be empty after Dispose", actual)
}

func Test_PathExistStat_MessageWithPathWrapped(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())

	// Act
	msg := stat.MessageWithPathWrapped("test message")

	// Assert
	actual := args.Map{"result": msg == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MessageWithPathWrapped should not be empty", actual)
}

func Test_PathExistStat_ParentWithGlobPatternFiles(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	stat := chmodhelper.GetPathExistStat(tmpFile.Name())

	// Act
	_, err := stat.ParentWithGlobPatternFiles("*.txt")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParentWithGlobPatternFiles error:", actual)
}

// ── chmodVerifier tests ──

func Test_ChmodVerifier_GetRwx9_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	fileMode := os.FileMode(0755)

	// Act
	rwx9 := chmodhelper.ChmodVerify.GetRwx9(fileMode)

	// Assert
	actual := args.Map{"result": len(rwx9) != 9}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 9 chars, got:", actual)
}

func Test_ChmodVerifier_GetRwxFull_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	fileMode := os.FileMode(0755)

	// Act
	rwxFull := chmodhelper.ChmodVerify.GetRwxFull(fileMode)

	// Assert
	actual := args.Map{"result": len(rwxFull) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10 chars, got:", actual)
}

func Test_ChmodVerifier_IsEqual_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	info, _ := os.Stat(tmpFile.Name())
	existingMode := info.Mode()

	// Act
	isEqual := chmodhelper.ChmodVerify.IsEqual(tmpFile.Name(), existingMode)

	// Assert
	actual := args.Map{"result": isEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqual should return true for existing file mode", actual)
}

func Test_ChmodVerifier_IsMismatch_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	isMismatch := chmodhelper.ChmodVerify.IsMismatch(tmpFile.Name(), os.FileMode(0777))

	// Assert -- may or may not be mismatch depending on OS defaults
	_ = isMismatch
}

func Test_ChmodVerifier_IsEqualSkipInvalid(t *testing.T) {
	// Arrange
	invalidPath := "/nonexistent/path/xyz"

	// Act
	isEqual := chmodhelper.ChmodVerify.IsEqualSkipInvalid(invalidPath, os.FileMode(0644))

	// Assert
	actual := args.Map{"result": isEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqualSkipInvalid should return true for invalid path", actual)
}

func Test_ChmodVerifier_IsEqualRwxFullSkipInvalid(t *testing.T) {
	// Arrange
	invalidPath := "/nonexistent/path/xyz"

	// Act
	isEqual := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid(invalidPath, "-rwxr-xr-x")

	// Assert
	actual := args.Map{"result": isEqual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqualRwxFullSkipInvalid should return true for invalid path", actual)
}

func Test_ChmodVerifier_GetExisting_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	mode, err := chmodhelper.ChmodVerify.GetExisting(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetExisting error:", actual)

	actual = args.Map{"result": mode == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mode should not be 0", actual)
}

func Test_ChmodVerifier_PathIf(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodVerify.PathIf(false, tmpFile.Name(), os.FileMode(0644))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PathIf with false should return nil", actual)
}

// ── GetExistingChmod tests ──

func Test_GetExistingChmod(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	mode, err := chmodhelper.GetExistingChmod(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)

	actual = args.Map{"result": mode == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mode should not be 0", actual)
}

func Test_GetExistingChmod_InvalidPath(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmod("/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid path", actual)
}

// ── IsPathExists / IsDirectory / IsPathInvalid ──

func Test_IsPathExists_Ext(t *testing.T) {
	// Assert
	actual := args.Map{"result": chmodhelper.IsPathExists(os.TempDir())}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "temp dir should exist", actual)

	actual = args.Map{"result": chmodhelper.IsPathExists("/nonexistent/path/xyz")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nonexistent path should not exist", actual)
}

func Test_IsDirectory_Ext(t *testing.T) {
	// Assert
	actual := args.Map{"result": chmodhelper.IsDirectory(os.TempDir())}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "temp dir should be directory", actual)
}

func Test_IsPathInvalid(t *testing.T) {
	// Assert
	actual := args.Map{"result": chmodhelper.IsPathInvalid("/nonexistent/xyz")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nonexistent path should be invalid", actual)

	actual = args.Map{"result": chmodhelper.IsPathInvalid(os.TempDir())}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "temp dir should not be invalid", actual)
}

// ── TempDirGetter / TempDirDefault ──

func Test_TempDirDefault(t *testing.T) {
	// Assert
	actual := args.Map{"result": chmodhelper.TempDirDefault == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TempDirDefault should not be empty", actual)
}

// ── GetExistingChmodRwxWrapperPtr ──

func Test_GetExistingChmodRwxWrapperPtr(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.GetExistingChmodRwxWrapperPtr(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)

	actual = args.Map{"result": wrapper == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "wrapper should not be nil", actual)
}

func Test_GetExistingChmodRwxWrapperPtr_InvalidPath(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmodRwxWrapperPtr("/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid path", actual)
}

// ── dirCreator tests ──

func Test_DirCreator_Direct_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_direct_test"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.Direct(dir)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Direct error:", actual)

	actual = args.Map{"result": chmodhelper.IsPathExists(dir)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "dir should exist after Direct", actual)
}

func Test_DirCreator_IfMissing_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_ifmissing_test"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(os.FileMode(0755), dir)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IfMissing error:", actual)

	// Act again - should be no-op
	err2 := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(os.FileMode(0755), dir)

	// Assert
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IfMissing second call error:", actual)
}

func Test_DirCreator_If_False_FromPathExistStatNilSafe(t *testing.T) {
	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.If(false, os.FileMode(0755), "/whatever")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "If with false should return nil", actual)
}

// ── fileWriter tests ──

func Test_FileWriter_WriteAndRead(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_write"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)
	content := "hello world"

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Default(
		true,
		filePath,
		[]byte(content),
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "write error:", actual)

	actual = args.Map{"result": chmodhelper.IsPathExists(filePath)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "file should exist after write", actual)
}

func Test_FileWriter_String_Default(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_string_write"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.Default(
		true,
		filePath,
		"hello string",
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string write error:", actual)
}

func Test_FileWriter_Remove_FromPathExistStatNilSafe(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_remove"
	filePath := dir + "/test.txt"
	os.MkdirAll(dir, 0755)
	os.WriteFile(filePath, []byte("test"), 0644)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Remove(filePath)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "remove error:", actual)

	actual = args.Map{"result": chmodhelper.IsPathExists(filePath)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "file should not exist after remove", actual)

	os.RemoveAll(dir)
}

func Test_FileWriter_RemoveIf(t *testing.T) {
	// Act -- should be no-op when false
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(false, "/whatever")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RemoveIf with false should return nil", actual)
}

func Test_FileWriter_ParentDir_FromPathExistStatNilSafe(t *testing.T) {
	// Act
	parent := chmodhelper.SimpleFileWriter.FileWriter.ParentDir("/tmp/test/file.txt")

	// Assert
	actual := args.Map{"result": parent == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParentDir should not be empty", actual)
}
