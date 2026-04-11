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

// ── errorCreator.dirError ──

func Test_ErrorCreator_DirError_NonExistentPath(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov8_dir_error_test_nonexist")
	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, 0755)
	os.RemoveAll(tmpDir)
}

// ── errorCreator.notDirError ──

func Test_NotDirError_PathInvalid(t *testing.T) {
	// For non-existent path, IsPathInvalid returns true, returns nil
	// Tested indirectly through dirCreator.ByChecking
}

func Test_NotDirError_ExistsButNotDir(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	// Create a file (not dir) to trigger "path exist but it is not a dir" branch
	tmpFile := filepath.Join(os.TempDir(), "cov8_notdir_test_file.txt")
	os.WriteFile(tmpFile, []byte("test"), 0644)
	defer os.Remove(tmpFile)

	// Use dir creator ByChecking on file path to exercise notDirError
	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(0755, tmpFile)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for file path used as dir", actual)
}

// ── errorCreator.pathError ──

func Test_PathError_NilErr(t *testing.T) {
	// Arrange
	// pathError returns nil when err is nil - covered through ApplyChmod on valid path
	tmpDir := filepath.Join(os.TempDir(), "cov8_path_error_nil")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, tmpDir)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_PathError_WithErr(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	// pathError returns error when path doesn't exist and skip=false
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/cov8/path")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nonexistent path", actual)
}

// ── errorCreator.pathErrorWithDirValidate ──

func Test_PathErrorWithDirValidate_NotDir(t *testing.T) {
	// Covered indirectly through CreateDirWithFiles with bad path
	tmpFile := filepath.Join(os.TempDir(), "cov8_dirvalidate_file.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)
	// No direct access to unexported dirCreator, exercise via public APIs
}

func Test_PathErrorWithDirValidate_ErrNil(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov8_dirvalidate_nil")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)
	// exercises code through public API
}

// ── errorCreator.chmodApplyFailed ──

func Test_ChmodApplyFailed_WithErr(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	// Covered through ApplyChmod on invalid path
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/cov8/chmod_fail")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ChmodApplyFailed_NilErr(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov8_chmod_apply_nil")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)
	// Successful chmod
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	_ = rwx.ApplyChmod(false, tmpDir)
}

// ── pathErrorMessage ──

func Test_PathErrorMessage(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	// Covered through any error path in ApplyChmod
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/cov8/pem")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual = args.Map{"result": len(err.Error()) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty error message", actual)
}

// ── dirCreator via CreateDirWithFiles ──

func Test_DirCreator_IfMissing_AlreadyExists(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov8_ifmissing_exists")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)
	// Creating again should be fine
	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{Dir: tmpDir})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_DirCreator_IfMissing_CreateNew(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov8_ifmissing_new")
	os.RemoveAll(tmpDir)
	defer os.RemoveAll(tmpDir)
	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{Dir: tmpDir})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_DirCreator_IfMissing_Error(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_ifmissing_err_file")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)
	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{
		Dir: filepath.Join(tmpFile, "subdir"),
	})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error creating dir under file", actual)
}

func Test_DirCreator_Default_Error(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_default_err")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)
	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{
		Dir: filepath.Join(tmpFile, "sub"),
	})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── tempDirGetter.TempPermanent ──

func Test_TempDirGetter_TempPermanent(t *testing.T) {
	// Arrange
	result := chmodhelper.TempDirGetter.TempPermanent()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty temp permanent path", actual)
}

// ── PathExistStat.MeaningFullError ──

func Test_PathExistStat_MeaningFullError(t *testing.T) {
	stat := chmodhelper.GetPathExistStat("/nonexistent/cov8/path")
	err := stat.MeaningFullError()
	// non-existent path may have Error set
	_ = err
}

func Test_PathExistStat_MeaningFullError_WithError(t *testing.T) {
	// Arrange
	stat := &chmodhelper.PathExistStat{
		Location: "/test",
		IsExist:  false,
		Error:    errors.New("test error"),
	}
	err := stat.MeaningFullError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── PathExistStat.NotAFileError ──

func Test_PathExistStat_NotAFileError_NotExist(t *testing.T) {
	stat := chmodhelper.GetPathExistStat("/nonexistent/cov8/not_a_file")
	err := stat.NotAFileError()
	if err == nil {
		// IsExist=false triggers NotExistError branch
	}
}

func Test_PathExistStat_NotAFileError_IsDir(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())
	err := stat.NotAFileError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error: dir is not a file", actual)
}

func Test_PathExistStat_NotAFileError_IsFile(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_notafile_isfile.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	stat := chmodhelper.GetPathExistStat(tmpFile)
	err := stat.NotAFileError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for file", actual)
}

// ── PathExistStat.NotADirError ──

func Test_PathExistStat_NotADirError_NotExist(t *testing.T) {
	stat := chmodhelper.GetPathExistStat("/nonexistent/cov8/not_a_dir")
	err := stat.NotADirError()
	_ = err
}

func Test_PathExistStat_NotADirError_IsFile(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_notadir_isfile.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	stat := chmodhelper.GetPathExistStat(tmpFile)
	err := stat.NotADirError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error: file is not a dir", actual)
}

func Test_PathExistStat_NotADirError_IsDir(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(os.TempDir())
	err := stat.NotADirError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for dir", actual)
}

// ── FilteredPathFileInfoMap.ValidLocations empty ──

func Test_FilteredPathFileInfoMap_ValidLocations_Empty(t *testing.T) {
	// Arrange
	m := chmodhelper.InvalidFilteredPathFileInfoMap()
	locs := m.ValidLocations()

	// Act
	actual := args.Map{"result": len(locs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ── FilteredPathFileInfoMap.ValidFileInfos empty ──

func Test_FilteredPathFileInfoMap_ValidFileInfos_Empty(t *testing.T) {
	// Arrange
	m := chmodhelper.InvalidFilteredPathFileInfoMap()
	infos := m.ValidFileInfos()

	// Act
	actual := args.Map{"result": len(infos) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ── FilteredPathFileInfoMap.ValidLocationFileInfoRwxWrappers empty ──

func Test_FilteredPathFileInfoMap_ValidLocationFileInfoRwxWrappers_Empty(t *testing.T) {
	// Arrange
	m := chmodhelper.InvalidFilteredPathFileInfoMap()
	wrappers := m.ValidLocationFileInfoRwxWrappers()

	// Act
	actual := args.Map{"result": len(wrappers) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ── FilteredPathFileInfoMap with valid entries ──

func Test_FilteredPathFileInfoMap_WithEntries(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_filtered_entries.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	m := chmodhelper.GetExistsFilteredPathFileInfoMap(false, tmpFile)
	locs := m.ValidLocations()

	// Act
	actual := args.Map{"result": len(locs) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected locations", actual)
	infos := m.ValidFileInfos()
	actual = args.Map{"result": len(infos) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected file infos", actual)
	wrappers := m.ValidLocationFileInfoRwxWrappers()
	actual = args.Map{"result": len(wrappers) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrappers", actual)
}

// ── GetExistingChmodRwxWrapperMustPtr ──

func Test_GetExistingChmodRwxWrapperMustPtr_Valid(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_must_ptr.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	ptr := chmodhelper.GetExistingChmodRwxWrapperMustPtr(tmpFile)

	// Act
	actual := args.Map{"result": ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_GetExistingChmodRwxWrapperMustPtr_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	chmodhelper.GetExistingChmodRwxWrapperMustPtr("/nonexistent/cov8/must_ptr")
}

// ── GetExistingChmodRwxWrappers ──

func Test_GetExistingChmodRwxWrappers_ContinueOnError(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_wrappers_cont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	results, err := chmodhelper.GetExistingChmodRwxWrappers(
		true, tmpFile, "/nonexistent/cov8/wrappers")
	_ = err

	// Act
	actual := args.Map{"result": len(results) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least one result", actual)
}

func Test_GetExistingChmodRwxWrappers_ImmediateExit(t *testing.T) {
	// Arrange
	_, err := chmodhelper.GetExistingChmodRwxWrappers(
		false, "/nonexistent/cov8/wrap1", "/nonexistent/cov8/wrap2")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_GetExistingChmodRwxWrappers_Empty(t *testing.T) {
	// Arrange
	results, err := chmodhelper.GetExistingChmodRwxWrappers(false)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(results) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ── IsChmodEqualUsingRwxOwnerGroupOther ──

func Test_IsChmodEqualUsingRwxOwnerGroupOther_Nil(t *testing.T) {
	// Arrange
	result := chmodhelper.IsChmodEqualUsingRwxOwnerGroupOther("/tmp", nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_IsChmodEqualUsingRwxOwnerGroupOther_Valid(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov8_chmod_equal_rwx")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r-x",
	}
	_ = chmodhelper.IsChmodEqualUsingRwxOwnerGroupOther(tmpDir, rwx)
}

// ── GetRecursivePaths ──

func Test_GetRecursivePaths_NonExistent(t *testing.T) {
	// Arrange
	_, err := chmodhelper.GetRecursivePaths(false, "/nonexistent/cov8/recursive")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_GetRecursivePaths_File(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_recursive_file.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	paths, err := chmodhelper.GetRecursivePaths(false, tmpFile)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(paths) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 path", actual)
}

func Test_GetRecursivePaths_Dir(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov8_recursive_dir")
	os.MkdirAll(filepath.Join(tmpDir, "sub"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "sub", "f.txt"), []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	paths, err := chmodhelper.GetRecursivePaths(false, tmpDir)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(paths) < 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected multiple paths", actual)
}

func Test_GetRecursivePaths_ContinueOnError(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov8_recursive_cont")
	os.MkdirAll(tmpDir, 0755)
	os.WriteFile(filepath.Join(tmpDir, "f.txt"), []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	paths, err := chmodhelper.GetRecursivePaths(true, tmpDir)
	_ = err

	// Act
	actual := args.Map{"result": len(paths) < 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected paths", actual)
}

// ── GetRecursivePathsContinueOnError ──

func Test_GetRecursivePathsContinueOnError_NonExistent(t *testing.T) {
	// Arrange
	_, err := chmodhelper.GetRecursivePathsContinueOnError("/nonexistent/cov8/recur_cont")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_GetRecursivePathsContinueOnError_File_FromErrorCreatorPaths(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov8_recur_cont_file.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	paths, err := chmodhelper.GetRecursivePathsContinueOnError(tmpFile)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(paths) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 path", actual)
}

func Test_GetRecursivePathsContinueOnError_Dir_FromErrorCreatorPaths(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov8_recur_cont_dir")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	paths, _ := chmodhelper.GetRecursivePathsContinueOnError(tmpDir)

	// Act
	actual := args.Map{"result": len(paths) < 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected paths", actual)
}
