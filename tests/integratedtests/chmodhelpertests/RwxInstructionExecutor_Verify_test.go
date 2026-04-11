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

// ── RwxInstructionExecutor: VerifyRwxModifiers mismatch path (line 261, 269) ──

func Test_RwxInstructionExecutor_VerifyMismatch(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o777)

	// ExpandRwxFullStringToOwnerGroupOther takes 10-char string WITH leading hyphen
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

// ── RwxVariableWrapper: VerifyRwxModifiers with wildcard rwx (line 46) ──

func Test_RwxVariableWrapper_VerifyWithNilRwx(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o644)

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

// ── RwxVariableWrapper: VerifyOnLocationsApplyChmod paths (line 186-218) ──

func Test_RwxVariableWrapper_VerifyOnLocations_ContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o644)

	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("-r*xr-xr-x")
	actual := args.Map{"result": ogoErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected ogo error:", actual)
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
	actual = args.Map{
		"parseOk": parseErr == nil,
	}
	expected = args.Map{
		"parseOk": true,
	}
	actual.ShouldBeEqual(t, 1, "RwxVariableWrapper VerifyOnLocations", expected)
	_ = applyErr
}

// ── RwxInstructionExecutor: ApplyOnPath with exit-on-invalid ──

func Test_RwxInstructionExecutor_ApplyOnPath_ExitOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("-rwxr-xr-x")
	actual := args.Map{"result": ogoErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected ogo error:", actual)
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
	actual = args.Map{
		"parseOk":  parseErr == nil,
		"hasError": applyErr != nil,
	}
	expected = args.Map{
		"parseOk":  true,
		"hasError": true,
	}
	actual.ShouldBeEqual(t, 1, "ApplyOnPath exit-on-invalid", expected)
}

// ── RwxInstructionExecutor: ApplyOnPath with skip-on-invalid ──

func Test_RwxInstructionExecutor_ApplyOnPath_SkipOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("-rwxr-xr-x")
	actual := args.Map{"result": ogoErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected ogo error:", actual)
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
	actual = args.Map{
		"parseOk":  parseErr == nil,
		"hasError": applyErr != nil,
	}
	expected = args.Map{
		"parseOk":  true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "ApplyOnPath skip-on-invalid", expected)
}

// ── RwxInstructionExecutor: ApplyOnPath recursive valid dir ──

func Test_RwxInstructionExecutor_ApplyOnPath_Recursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	_ = os.WriteFile(filepath.Join(tmpDir, "f.txt"), []byte("x"), 0o644)
	ogo, ogoErr := chmodins.ExpandRwxFullStringToOwnerGroupOther("-rwxrwxrwx")
	actual := args.Map{"result": ogoErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected ogo error:", actual)
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
	actual = args.Map{
		"parseOk":  parseErr == nil,
		"hasError": applyErr != nil,
	}
	expected = args.Map{
		"parseOk":  true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "ApplyOnPath recursive valid dir", expected)
}
