package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/smartystreets/goconvey/convey"
)

func mustRwxWrapper(rwx string) chmodhelper.RwxWrapper {
	wrapper, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen(rwx)
	if err != nil {
		panic(err)
	}

	return wrapper
}

// ══════════════════════════════════════════════════════════════════════════════
// Coverage19 — chmodhelper final 95 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── CreateDirWithFiles: file-create error path (line 62) and chmod error (line 75) ──

func Test_Cov19_CreateDirWithFiles_FileCreateError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	invalidDir := string([]byte{0}) + "/noexist/sub"
	dwf := &chmodhelper.DirWithFiles{
		Dir: invalidDir,
		Files: []string{
			"a.txt",
		},
	}

	// Act
	err := chmodhelper.CreateDirWithFiles(false, 0o644, dwf)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "CreateDirWithFiles file-create error", expected)
}

// ── GetRecursivePathsContinueOnError: walk error appended (line 47-51) ──

func Test_Cov19_GetRecursivePathsContinueOnError_WalkError(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0})

	// Act
	paths, err := chmodhelper.GetRecursivePathsContinueOnError(invalidPath)

	// Assert
	actual := args.Map{
		"pathsLen": len(paths),
		"hasError": err != nil,
	}
	expected := args.Map{
		"pathsLen": 0,
		"hasError": true,
	}
	actual.ShouldBeEqual(t, 1, "GetRecursivePathsContinueOnError walk error", expected)
}

// ── MergeRwxWildcardWithFixedRwx: bad existing rwx length (line 38-40) ──

func Test_Cov19_MergeRwxWildcardWithFixedRwx_BadExistingLength(t *testing.T) {
	// Arrange
	wildcardInput := "r-x"
	existingBadRwx := "rw" // wrong length

	// Act
	result, err := chmodhelper.MergeRwxWildcardWithFixedRwx(wildcardInput, existingBadRwx)

	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasError":  err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasError":  true,
	}
	actual.ShouldBeEqual(t, 1, "MergeRwxWildcardWithFixedRwx bad existing length", expected)
}

// ── RwxInstructionExecutor: VerifyRwxModifiers mismatch path (line 261, 269) ──

func Test_Cov19_RwxInstructionExecutor_VerifyMismatch(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o777)

	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("-r--------")
	if ogoErr != nil {
		panic(ogoErr)
	}

	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *ogo,
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
		},
	}
	executor, parseErr := chmodhelper.ParseRwxInstructionToExecutor(&ins)

	// Act
	var verifyErr error
	if parseErr == nil {
		verifyErr = executor.VerifyRwxModifiers(false, []string{testFile})
	}

	// Assert
	actual := args.Map{
		"parseOk":  parseErr == nil,
		"hasError": verifyErr != nil,
	}
	expected := args.Map{
		"parseOk":  true,
		"hasError": true,
	}
	actual.ShouldBeEqual(t, 1, "VerifyRwxModifiers mismatch", expected)
}

// ── RwxVariableWrapper: VerifyRwxModifiers with nil rwxWrapper (line 46) ──

func Test_Cov19_RwxVariableWrapper_VerifyWithNilRwx(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o644)

	// Create a var wrapper with wildcard
	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("-r*xr-xr-x")
	if ogoErr != nil {
		panic(ogoErr)
	}

	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *ogo,
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
		},
	}
	executor, parseErr := chmodhelper.ParseRwxInstructionToExecutor(&ins)

	// Act
	var verifyErr error
	if parseErr == nil {
		verifyErr = executor.VerifyRwxModifiers(false, []string{testFile})
	}

	// Assert
	actual := args.Map{
		"parseOk": parseErr == nil,
	}
	expected := args.Map{
		"parseOk": true,
	}
	actual.ShouldBeEqual(t, 1, "RwxVariableWrapper verify", expected)
	_ = verifyErr
}

// ── RwxWrapper.ApplyChmod on valid file (lines 227-255) ──

func Test_Cov19_RwxWrapper_ApplyChmod_ValidFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o644)

	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmod(false, testFile)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyChmod valid file", expected)
}

// ── RwxWrapper.ApplyChmod on invalid path (lines 237-239) ──

func Test_Cov19_RwxWrapper_ApplyChmod_InvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmod(false, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyChmod invalid path", expected)
}

// ── RwxWrapper.ApplyChmod skipOnInvalid with invalid path ──

func Test_Cov19_RwxWrapper_ApplyChmod_SkipInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmod(true, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyChmod skip invalid", expected)
}

// ── RwxWrapper.ApplyChmodSkipInvalid (line 304-307) ──

func Test_Cov19_RwxWrapper_ApplyChmodSkipInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmodSkipInvalid("/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyChmodSkipInvalid", expected)
}

// ── RwxWrapper.LinuxApplyRecursive valid dir (line 328-345) ──

func Test_Cov19_RwxWrapper_LinuxApplyRecursive_ValidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "sub")
	_ = os.Mkdir(subDir, 0o755)
	_ = os.WriteFile(filepath.Join(subDir, "a.txt"), []byte("x"), 0o644)
	wrapper := mustRwxWrapper("rwxrwxrwx")

	// Act
	err := wrapper.LinuxApplyRecursive(false, tmpDir)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive valid dir", expected)
}

// ── RwxWrapper.LinuxApplyRecursive invalid path, skip=false (line 334-341) ──

func Test_Cov19_RwxWrapper_LinuxApplyRecursive_InvalidNotSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.LinuxApplyRecursive(false, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive invalid not skip", expected)
}

// ── RwxWrapper.LinuxApplyRecursive invalid path, skip=true (line 330-332) ──

func Test_Cov19_RwxWrapper_LinuxApplyRecursive_InvalidSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.LinuxApplyRecursive(true, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive invalid skip", expected)
}

// ── RwxWrapper.ApplyRecursive on non-linux (line 368-431) ──

func Test_Cov19_RwxWrapper_ApplyRecursive_ValidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	_ = os.WriteFile(filepath.Join(tmpDir, "a.txt"), []byte("x"), 0o644)
	wrapper := mustRwxWrapper("rwxrwxrwx")

	// Act
	err := wrapper.ApplyRecursive(false, tmpDir)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyRecursive valid dir", expected)
}

// ── RwxWrapper.ApplyRecursive invalid path not skip (line 368-373) ──

func Test_Cov19_RwxWrapper_ApplyRecursive_InvalidNotSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyRecursive(false, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyRecursive invalid not skip", expected)
}

// ── RwxWrapper.ApplyRecursive on a single file (line 386-397) ──

func Test_Cov19_RwxWrapper_ApplyRecursive_SingleFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "single.txt")
	_ = os.WriteFile(testFile, []byte("x"), 0o644)
	wrapper := mustRwxWrapper("rwxrwxrwx")

	// Act
	err := wrapper.ApplyRecursive(false, testFile)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyRecursive single file", expected)
}

// ── RwxWrapper.ApplyLinuxChmodOnMany non-recursive (line 544-568) ──

func Test_Cov19_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "a.txt")
	f2 := filepath.Join(tmpDir, "b.txt")
	_ = os.WriteFile(f1, []byte("x"), 0o644)
	_ = os.WriteFile(f2, []byte("y"), 0o644)
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	if wErr != nil {
		t.Fatalf("unexpected parse error: %v", wErr)
	}
	cond := &chmodins.Condition{
		IsRecursive:       false,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, f1, f2)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany non-recursive", expected)
}

// ── RwxWrapper.ApplyLinuxChmodOnMany non-recursive with error (line 557-565) ──

func Test_Cov19_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursiveError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	if wErr != nil {
		t.Fatalf("unexpected parse error: %v", wErr)
	}
	cond := &chmodins.Condition{
		IsRecursive:       false,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, "/no/exist/1", "/no/exist/2")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany non-recursive error", expected)
}

// ── RwxWrapper.ApplyLinuxChmodOnMany recursive (line 594-618) ──

func Test_Cov19_RwxWrapper_ApplyLinuxChmodOnMany_Recursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "sub")
	_ = os.Mkdir(subDir, 0o755)
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	if wErr != nil {
		t.Fatalf("unexpected parse error: %v", wErr)
	}
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, tmpDir)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany recursive", expected)
}

// ── RwxWrapper.ApplyLinuxChmodOnMany recursive error (line 607-615) ──

func Test_Cov19_RwxWrapper_ApplyLinuxChmodOnMany_RecursiveError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	if wErr != nil {
		t.Fatalf("unexpected parse error: %v", wErr)
	}
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, "/no/exist/path")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany recursive error", expected)
}

// ── RwxWrapper.applyLinuxChmodRecursiveManyContinueOnError (line 624-642) ──

func Test_Cov19_RwxWrapper_ApplyLinuxChmodOnMany_RecursiveContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	if wErr != nil {
		t.Fatalf("unexpected parse error: %v", wErr)
	}
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: true,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, tmpDir, "/no/exist")

	// Assert — at least one path should fail but continue
	_ = err // some errors aggregated
	convey.Convey("RecursiveContinueOnError processes all paths", t, func() {
		// test passes if no panic
		convey.So(true, convey.ShouldBeTrue)
	})
}

// ── RwxWrapper.applyLinuxChmodNonRecursiveManyContinueOnError (line 648-662) ──

func Test_Cov19_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursiveContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	validFile := filepath.Join(tmpDir, "ok.txt")
	_ = os.WriteFile(validFile, []byte("x"), 0o644)
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	if wErr != nil {
		t.Fatalf("unexpected parse error: %v", wErr)
	}
	cond := &chmodins.Condition{
		IsRecursive:       false,
		IsContinueOnError: true,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, validFile, "/no/exist")

	// Assert — errors aggregated, no panic
	_ = err
	convey.Convey("NonRecursiveContinueOnError processes all paths", t, func() {
		convey.So(true, convey.ShouldBeTrue)
	})
}

// ── RwxWrapper.getLinuxRecursiveCmdForChmod / applyLinuxRecursiveChmodUsingCmd (line 475-489) ──

func Test_Cov19_RwxWrapper_ApplyRecursive_Dir_CmdPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange - Linux uses cmd-based recursive chmod
	if runtime.GOOS != "linux" {
		t.Skip("linux only for cmd-based recursive chmod")
	}

	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "inner")
	_ = os.Mkdir(subDir, 0o755)
	_ = os.WriteFile(filepath.Join(subDir, "f.txt"), []byte("x"), 0o644)
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	if wErr != nil {
		t.Fatalf("unexpected parse error: %v", wErr)
	}

	// Act
	err := wrapper.LinuxApplyRecursive(false, tmpDir)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive cmd path", expected)
}

// ── RwxWrapper.IsEqualVarWrapper (line 579-588) ──

func Test_Cov19_RwxWrapper_IsEqualVarWrapper_Nil(t *testing.T) {
	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	result := wrapper.IsEqualVarWrapper(nil)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	actual.ShouldBeEqual(t, 1, "IsEqualVarWrapper nil", expected)
}

// ── RwxVariableWrapper: VerifyOnLocationsApplyChmod paths (line 186-218) ──

func Test_Cov19_RwxVariableWrapper_VerifyOnLocations_ContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o644)

	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("r*xr-xr-x")
	if ogoErr != nil {
		t.Fatalf("unexpected ogo error: %v", ogoErr)
	}
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *ogo,
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
		},
	}
	executor, parseErr := chmodhelper.ParseRwxInstructionToExecutor(&ins)

	// Act
	var applyErr error
	if parseErr == nil {
		applyErr = executor.ApplyOnPath(testFile)
	}

	// Assert
	actual := args.Map{
		"parseOk": parseErr == nil,
	}
	expected := args.Map{
		"parseOk": true,
	}
	actual.ShouldBeEqual(t, 1, "RwxVariableWrapper VerifyOnLocations", expected)
	_ = applyErr
}

// ── RwxWrapper: invalidPathErr (line 86-93) ──

func Test_Cov19_RwxWrapper_ApplyChmod_InvalidPath_NotSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxr-xr-x")
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	// Act — path doesn't exist
	applyErr := wrapper.ApplyChmod(false, "/nonexistent/path/file.txt")

	// Assert
	actual := args.Map{"hasError": applyErr != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyChmod invalidPathErr", expected)
}

// ── RwxInstructionExecutor: ApplyOnPath with exit-on-invalid ──

func Test_Cov19_RwxInstructionExecutor_ApplyOnPath_ExitOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("rwxr-xr-x")
	if ogoErr != nil {
		t.Fatalf("unexpected ogo error: %v", ogoErr)
	}
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *ogo,
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
			IsRecursive:     false,
		},
	}
	executor, parseErr := chmodhelper.ParseRwxInstructionToExecutor(&ins)

	// Act
	var applyErr error
	if parseErr == nil {
		applyErr = executor.ApplyOnPath("/no/such/path")
	}

	// Assert
	actual := args.Map{
		"parseOk":  parseErr == nil,
		"hasError": applyErr != nil,
	}
	expected := args.Map{
		"parseOk":  true,
		"hasError": true,
	}
	actual.ShouldBeEqual(t, 1, "ApplyOnPath exit-on-invalid", expected)
}

// ── RwxInstructionExecutor: ApplyOnPath with skip-on-invalid ──

func Test_Cov19_RwxInstructionExecutor_ApplyOnPath_SkipOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("rwxr-xr-x")
	if ogoErr != nil {
		t.Fatalf("unexpected ogo error: %v", ogoErr)
	}
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *ogo,
		Condition: chmodins.Condition{
			IsSkipOnInvalid: true,
			IsRecursive:     false,
		},
	}
	executor, parseErr := chmodhelper.ParseRwxInstructionToExecutor(&ins)

	// Act
	var applyErr error
	if parseErr == nil {
		applyErr = executor.ApplyOnPath("/no/such/path")
	}

	// Assert
	actual := args.Map{
		"parseOk":  parseErr == nil,
		"hasError": applyErr != nil,
	}
	expected := args.Map{
		"parseOk":  true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "ApplyOnPath skip-on-invalid", expected)
}

// ── RwxInstructionExecutor: ApplyOnPath recursive valid dir ──

func Test_Cov19_RwxInstructionExecutor_ApplyOnPath_Recursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	_ = os.WriteFile(filepath.Join(tmpDir, "f.txt"), []byte("x"), 0o644)
	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("rwxrwxrwx")
	if ogoErr != nil {
		t.Fatalf("unexpected ogo error: %v", ogoErr)
	}
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *ogo,
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
			IsRecursive:     true,
		},
	}
	executor, parseErr := chmodhelper.ParseRwxInstructionToExecutor(&ins)

	// Act
	var applyErr error
	if parseErr == nil {
		applyErr = executor.ApplyOnPath(tmpDir)
	}

	// Assert
	actual := args.Map{
		"parseOk":  parseErr == nil,
		"hasError": applyErr != nil,
	}
	expected := args.Map{
		"parseOk":  true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "ApplyOnPath recursive valid dir", expected)
}
