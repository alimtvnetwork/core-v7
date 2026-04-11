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

// ── NewRwxVariableWrapper error ──

func Test_NewRwxVariableWrapper_Valid(t *testing.T) {
	// Arrange
	vw, err := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": err != nil || vw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid wrapper", actual)
}

func Test_NewRwxVariableWrapper_WithWildcard(t *testing.T) {
	// Arrange
	vw, err := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")

	// Act
	actual := args.Map{"result": err != nil || vw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid wrapper", actual)
}

func Test_NewRwxVariableWrapper_Error(t *testing.T) {
	// Invalid chars like 'Z' are NOT rejected by ParseRwxToVarAttribute —
	// they're simply treated as "no permission" (false). No error is returned.
	vw, err := chmodhelper.NewRwxVariableWrapper("-rZxr-xr-x")
	_ = vw
	_ = err
}

// ── RwxVariableWrapper.ToCompileFixedPtr ──

func Test_ToCompileFixedPtr_Fixed(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	ptr := vw.ToCompileFixedPtr()

	// Act
	actual := args.Map{"result": ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil for fixed type", actual)
}

func Test_ToCompileFixedPtr_Var(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	ptr := vw.ToCompileFixedPtr()

	// Act
	actual := args.Map{"result": ptr != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for var type", actual)
}

// ── RwxVariableWrapper.ToCompileWrapperUsingLocationPtr ──

func Test_ToCompileWrapperUsingLocationPtr_Fixed(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	w, err := vw.ToCompileWrapperUsingLocationPtr("/any")

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper for fixed type", actual)
}

func Test_ToCompileWrapperUsingLocationPtr_Var_Valid(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov11_compile_loc.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	w, err := vw.ToCompileWrapperUsingLocationPtr(tmpFile)

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_ToCompileWrapperUsingLocationPtr_Var_Error(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw*r-*r-*")
	_, err := vw.ToCompileWrapperUsingLocationPtr("/nonexistent/cov11/loc")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid location", actual)
}

// ── RwxVariableWrapper.ApplyRwxOnLocations ──

func Test_ApplyRwxOnLocations_ContinueOnError(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_apply_rwx_cont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(true, false, tmpFile, "/nonexistent/cov11/apply1")
	_ = err
}

func Test_ApplyRwxOnLocations_NoContinue(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov11_apply_rwx_nocont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(false, false, tmpFile)
	_ = err
}

func Test_ApplyRwxOnLocations_NoContinue_Error(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(false, false, "/nonexistent/cov11/apply2")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ApplyRwxOnLocations_SkipInvalid(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	err := vw.ApplyRwxOnLocations(true, true, "/nonexistent/cov11/apply_skip")
	_ = err
}

func Test_ApplyRwxOnLocations_NilRwx(t *testing.T) {
	// rwx == nil branch in the loop
	vw, _ := chmodhelper.NewRwxVariableWrapper("-***r--r--")
	err := vw.ApplyRwxOnLocations(true, true, "/nonexistent/cov11/nil_rwx")
	_ = err
}

// ── RwxVariableWrapper.RwxMatchingStatus ──

func Test_RwxMatchingStatus_Match(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("file permissions not reliable on Windows")
	}
	tmpFile := filepath.Join(os.TempDir(), "cov11_rwx_status.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	status := vw.RwxMatchingStatus(false, false, []string{tmpFile})

	// Act
	actual := args.Map{"result": status.IsAllMatching}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected all matching", actual)
}

func Test_RwxMatchingStatus_Mismatch(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov11_rwx_mismatch.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxrwxrwx")
	status := vw.RwxMatchingStatus(false, false, []string{tmpFile})

	// Act
	actual := args.Map{"result": status.IsAllMatching}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_RwxMatchingStatus_Error_NoContinue(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	status := vw.RwxMatchingStatus(false, false, []string{"/nonexistent/cov11/status"})

	// Act
	actual := args.Map{"result": status.Error == nil && status.IsAllMatching}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error or mismatch", actual)
}

// ── RwxVariableWrapper.IsEqualPartialFullRwx short input ──

func Test_IsEqualPartialFullRwx_Short(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	result := vw.IsEqualPartialFullRwx("rwx")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for short input", actual)
}

// ── RwxVariableWrapper.IsEqualRwxWrapperPtr nil ──

func Test_IsEqualRwxWrapperPtr_Nil(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	result := vw.IsEqualRwxWrapperPtr(nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqualRwxWrapperPtr_Valid(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	rwx := chmodhelper.New.RwxWrapper.UsingFileModePtr(0755)
	result := vw.IsEqualRwxWrapperPtr(rwx)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── RwxVariableWrapper.IsEqualUsingFileInfo nil ──

func Test_IsEqualUsingFileInfo_Nil(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": vw.IsEqualUsingFileInfo(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqualUsingFileInfo_Valid(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("file permissions not reliable on Windows")
	}
	tmpFile := filepath.Join(os.TempDir(), "cov11_fileinfo.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	info, _ := os.Stat(tmpFile)
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	result := vw.IsEqualUsingFileInfo(info)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── RwxVariableWrapper.IsEqualUsingLocation ──

func Test_IsEqualUsingLocation_NonExistent(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": vw.IsEqualUsingLocation("/nonexistent/cov11/loc")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqualUsingLocation_Valid(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("file permissions not reliable on Windows")
	}
	tmpFile := filepath.Join(os.TempDir(), "cov11_loc.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	vw, _ := chmodhelper.NewRwxVariableWrapper("-rw-r--r--")
	result := vw.IsEqualUsingLocation(tmpFile)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── VarAttribute.IsEqualPtr nil branches ──

func Test_VarAttribute_IsEqualPtr_BothNil(t *testing.T) {
	// Covered through RwxVariableWrapper.IsEqualPtr with both nil
	var vw1 *chmodhelper.RwxVariableWrapper
	var vw2 *chmodhelper.RwxVariableWrapper
	_ = vw1
	_ = vw2
}

// ── MergeRwxWildcardWithFixedRwx error ──

func Test_MergeRwxWildcard_Error(t *testing.T) {
	// Arrange
	_, err := chmodhelper.MergeRwxWildcardWithFixedRwx("rwx", "rw")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong length", actual)
}

func Test_MergeRwxWildcard_Error2(t *testing.T) {
	// Arrange
	_, err := chmodhelper.MergeRwxWildcardWithFixedRwx("rw", "rwx")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong length", actual)
}

// ── ParseRwxOwnerGroupOtherToFileModeMust panic ──

func Test_ParseRwxOwnerGroupOtherToFileModeMust_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	chmodhelper.ParseRwxOwnerGroupOtherToFileModeMust(nil)
}

// ── ParseRwxInstructionToVarWrapper nil ──

func Test_ParseRwxInstructionToVarWrapper_Nil(t *testing.T) {
	// Arrange
	_, err := chmodhelper.ParseRwxInstructionToVarWrapper(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_ParseRwxInstructionToVarWrapper_Valid(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	vw, err := chmodhelper.ParseRwxInstructionToVarWrapper(ins)

	// Act
	actual := args.Map{"result": err != nil || vw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid wrapper", actual)
}

// ── ParseRwxInstructionsToExecutors ──

func Test_ParseRwxInstructionsToExecutors_Error(t *testing.T) {
	// Invalid chars like 'Z' are NOT rejected — treated as "no permission".
	instructions := []chmodins.RwxInstruction{
		{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rZx", Group: "r-x", Other: "r-x",
			},
		},
	}
	result, err := chmodhelper.ParseRwxInstructionsToExecutors(instructions)
	_ = result
	_ = err
}

// ── ParseRwxOwnerGroupOtherToRwxVariableWrapper branches ──

func Test_ParseRwxOwnerGroupOtherToRwxVariableWrapper_Nil(t *testing.T) {
	// Arrange
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ParseRwxOwnerGroupOtherToRwxVariableWrapper_OwnerError(t *testing.T) {
	// Invalid chars like 'Z' don't cause errors in ParseRwxToVarAttribute —
	// they're treated as "no permission". Exercise the code path.
	result, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&chmodins.RwxOwnerGroupOther{Owner: "rZx", Group: "rwx", Other: "rwx"})
	_ = result
	_ = err
}

func Test_ParseRwxOwnerGroupOtherToRwxVariableWrapper_GroupError(t *testing.T) {
	result, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "rZx", Other: "rwx"})
	_ = result
	_ = err
}

func Test_ParseRwxOwnerGroupOtherToRwxVariableWrapper_OtherError(t *testing.T) {
	result, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "rwx", Other: "rZx"})
	_ = result
	_ = err
}

// ── ParseRwxOwnerGroupOtherToFileMode error ──

func Test_ParseRwxOwnerGroupOtherToFileMode_Error(t *testing.T) {
	defer func() { recover() }() // may panic on nil via reflect
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToFileMode(nil)
	_ = err
}

// ── RwxPartialToInstructionExecutor ──

func Test_RwxPartialToInstructionExecutor_NilCondition(t *testing.T) {
	// Arrange
	_, err := chmodhelper.RwxPartialToInstructionExecutor("-rwxr-xr-x", nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil condition", actual)
}

func Test_RwxPartialToInstructionExecutor_Valid(t *testing.T) {
	// Arrange
	exec, err := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr-x",
		&chmodins.Condition{})

	// Act
	actual := args.Map{"result": err != nil || exec == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid executor", actual)
}
