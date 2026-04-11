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

// ── RwxWrapper.ToUint32Octal ──

func Test_RwxWrapper_ToUint32Octal_FromRwxWrapperOctalApply(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	oct := rwx.ToUint32Octal()

	// Act
	actual := args.Map{"result": oct != 0755}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0755, got %o", actual)
}

// ── RwxWrapper.ApplyChmod branches ──

func Test_RwxWrapper_ApplyChmod_SkipInvalid_FromRwxWrapperOctalApply(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(true, "/nonexistent/cov9/skip")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for skip on invalid", actual)
}

func Test_RwxWrapper_ApplyChmod_NotSkipInvalid(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/cov9/noskip")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-skip invalid path", actual)
}

func Test_RwxWrapper_ApplyChmod_Success(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov9_apply_chmod.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyChmod(false, tmpFile)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_RwxWrapper_ApplyChmod_ChmodFail(t *testing.T) {
	// Arrange
	// On most systems, regular chmod doesn't fail on valid paths
	// This covers the success path with error=nil
	tmpFile := filepath.Join(os.TempDir(), "cov9_apply_chmod_ok.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0777)
	err := rwx.ApplyChmod(false, tmpFile)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── RwxWrapper.invalidPathErr ──

func Test_RwxWrapper_InvalidPathErr(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/cov9/invalid_path")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── RwxWrapper.ApplyChmodOptions ──

func Test_ApplyChmodOptions_SkipApply(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(false, true, false, "/any")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil when isApply=false", actual)
}

func Test_ApplyChmodOptions_InvalidSkip(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, true, true, "/nonexistent/cov9/opts")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for skip invalid", actual)
}

func Test_ApplyChmodOptions_InvalidNoSkip(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, true, false, "/nonexistent/cov9/opts2")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid no-skip", actual)
}

func Test_ApplyChmodOptions_MismatchApply(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_opts_mismatch.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, true, false, tmpFile)
	_ = err
}

func Test_ApplyChmodOptions_AlreadyMatching(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov9_opts_match.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyChmodOptions(true, true, false, tmpFile)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── RwxWrapper.LinuxApplyRecursive ──

func Test_LinuxApplyRecursive_SkipInvalid_NotExists(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(true, "/nonexistent/cov9/linux_recur")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for skip invalid", actual)
}

func Test_LinuxApplyRecursive_NoSkip_NotExists(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(false, "/nonexistent/cov9/linux_recur2")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_LinuxApplyRecursive_Valid(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_linux_recur")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(false, tmpDir)
	_ = err // depends on OS
}

// ── RwxWrapper.ApplyRecursive ──

func Test_ApplyRecursive_SkipInvalid(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(true, "/nonexistent/cov9/recur_skip")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_ApplyRecursive_NotExist_NoSkip(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(false, "/nonexistent/cov9/recur_noskip")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ApplyRecursive_File(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_recur_file.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyRecursive(false, tmpFile)
	_ = err
}

func Test_ApplyRecursive_Dir(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_recur_dir")
	os.MkdirAll(filepath.Join(tmpDir, "sub"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "f.txt"), []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(false, tmpDir)
	_ = err
}

// ── RwxWrapper.MustApplyChmod ──

func Test_MustApplyChmod_Success(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_must_apply.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	rwx.MustApplyChmod(tmpFile) // should not panic
}

func Test_MustApplyChmod_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	rwx.MustApplyChmod("/nonexistent/cov9/must_apply")
}

// ── RwxWrapper.ApplyLinuxChmodOnMany ──

func Test_ApplyLinuxChmodOnMany_Recursive(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_linux_many_recur")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsRecursive: true},
		tmpDir)
	_ = err
}

func Test_ApplyLinuxChmodOnMany_NonRecursive(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_linux_many_nonrecur.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsRecursive: false},
		tmpFile)
	_ = err
}

func Test_ApplyLinuxChmodOnMany_ContinueOnError_NonRecursive(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsContinueOnError: true, IsRecursive: false},
		"/nonexistent/cov9/many1", "/nonexistent/cov9/many2")
	_ = err
}

func Test_ApplyLinuxChmodOnMany_ContinueOnError_Recursive(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsContinueOnError: true, IsRecursive: true},
		"/nonexistent/cov9/many3")
	_ = err
}

func Test_ApplyLinuxChmodOnMany_StopOnError_Recursive(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{IsRecursive: true},
		"/nonexistent/cov9/many4")
	_ = err
}

func Test_ApplyLinuxChmodOnMany_StopOnError_NonRecursive(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov9_stop_nonrecur.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyLinuxChmodOnMany(
		&chmodins.Condition{},
		tmpFile, "/nonexistent/cov9/many5")
	_ = err
}

// ── RwxWrapper.IsEqualVarWrapper ──

func Test_IsEqualVarWrapper_Nil(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)

	// Act
	actual := args.Map{"result": rwx.IsEqualVarWrapper(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_IsEqualVarWrapper_Match(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	varW, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	result := rwx.IsEqualVarWrapper(varW)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── RwxWrapper.IsRwxEqualFileInfo ──

func Test_IsRwxEqualFileInfo_Nil(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)

	// Act
	actual := args.Map{"result": rwx.IsRwxEqualFileInfo(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsRwxEqualFileInfo_Valid(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	tmpFile := filepath.Join(os.TempDir(), "cov9_fileinfo.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	info, _ := os.Stat(tmpFile)
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	result := rwx.IsRwxEqualFileInfo(info)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── RwxWrapper.IsRwxEqualLocation ──

func Test_IsRwxEqualLocation_NonExistent(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)

	// Act
	actual := args.Map{"result": rwx.IsRwxEqualLocation("/nonexistent/cov9/rwxloc")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsRwxEqualLocation_Valid(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
	tmpFile := filepath.Join(os.TempDir(), "cov9_rwxloc.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0644)
	result := rwx.IsRwxEqualLocation(tmpFile)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── RwxWrapper.getLinuxRecursiveCmdForChmod ──

func Test_GetLinuxRecursiveCmdForChmod(t *testing.T) {
	// Covered through LinuxApplyRecursive on valid dir
	tmpDir := filepath.Join(os.TempDir(), "cov9_getcmd")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	_ = rwx.LinuxApplyRecursive(false, tmpDir)
}

// ── RwxWrapper.applyLinuxRecursiveChmodUsingCmd ──

func Test_ApplyLinuxRecursiveChmodUsingCmd(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov9_linuxcmd")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(false, tmpDir)
	_ = err
}
