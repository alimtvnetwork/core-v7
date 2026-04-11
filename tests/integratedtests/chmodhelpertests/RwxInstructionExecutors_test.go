package chmodhelpertests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RwxInstructionExecutor.CompiledWrapper ──

func Test_CompiledWrapper_Fixed(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	exec, err := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	w, err := exec.CompiledWrapper(0755)
	actual = args.Map{"result": err != nil || w == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_CompiledWrapper_Var(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw*", Group: "r-*", Other: "r-*",
		},
	}
	exec, err := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	w, err := exec.CompiledWrapper(0755)
	actual = args.Map{"result": err != nil || w == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

// ── RwxInstructionExecutor.CompiledRwxWrapperUsingFixedRwxWrapper ──

func Test_CompiledRwxWrapperUsingFixed_Fixed(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	exec, err := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	existing := chmodhelper.New.RwxWrapper.UsingFileModePtr(0755)
	w, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(existing)
	actual = args.Map{"result": err != nil || w == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_CompiledRwxWrapperUsingFixed_Var(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw*", Group: "*-x", Other: "r-*",
		},
	}
	exec, err := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	existing := chmodhelper.New.RwxWrapper.UsingFileModePtr(0755)
	w, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(existing)
	actual = args.Map{"result": err != nil || w == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

// ── RwxInstructionExecutor.ApplyOnPath ──

func Test_ApplyOnPath_ExitOnInvalid_Error(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.ApplyOnPath("/nonexistent/cov10/exit_invalid")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ApplyOnPath_SkipOnInvalid(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid: true,
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.ApplyOnPath("/nonexistent/cov10/skip_invalid")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for skip on invalid", actual)
}

func Test_ApplyOnPath_Recursive(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov10_apply_recur")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsRecursive: true,
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.ApplyOnPath(tmpDir)
	_ = err
}

func Test_ApplyOnPath_NonRecursive(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov10_apply_nonrecur.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.ApplyOnPath(tmpFile)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── RwxInstructionExecutor.ApplyOnPaths / ApplyOnPathsDirect / ApplyOnPathsPtr ──

func Test_ApplyOnPaths_Empty(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.ApplyOnPaths([]string{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_ApplyOnPathsDirect_Empty(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.ApplyOnPathsDirect()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_ApplyOnPathsPtr_Nil(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.ApplyOnPathsPtr(nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_ApplyOnPathsPtr_ContinueOnError(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsContinueOnError: true,
			IsSkipOnInvalid:   true,
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	locs := []string{"/nonexistent/cov10/p1", "/nonexistent/cov10/p2"}
	err := exec.ApplyOnPathsPtr(&locs)
	// skip on invalid -> nil

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_ApplyOnPathsPtr_StopOnError(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov10_stop.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	locs := []string{tmpFile}
	err := exec.ApplyOnPathsPtr(&locs)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── RwxInstructionExecutor.VerifyRwxModifiers ──

func Test_VerifyRwxModifiers_Empty(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.VerifyRwxModifiers(true, []string{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_VerifyRwxModifiers_ContinueOnError(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov10_verify_cont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
		Condition: chmodins.Condition{
			IsContinueOnError: true,
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.VerifyRwxModifiers(true, []string{tmpFile, "/nonexistent/cov10/verify1"})
	_ = err
}

func Test_VerifyRwxModifiers_NoContinue(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov10_verify_nocont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.VerifyRwxModifiers(true, []string{tmpFile})
	_ = err
}

func Test_VerifyRwxModifiers_RecursiveNotSupported(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsRecursive: true,
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.VerifyRwxModifiers(false, []string{"/some/path"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for recursive verify", actual)
}

func Test_VerifyRwxModifiersDirect(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov10_verify_direct.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.VerifyRwxModifiersDirect(true, tmpFile)
	_ = err
}

// ── RwxInstructionExecutor.verifyChmodLocationsNoContinue branches ──

func Test_VerifyNoContinue_ErrorWithSkip(t *testing.T) {
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid: true,
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.VerifyRwxModifiers(true, []string{"/nonexistent/cov10/vn1"})
	_ = err
}

func Test_VerifyNoContinue_Mismatch(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov10_verify_mismatch.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "rwx", Other: "rwx",
		},
	}
	exec, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	err := exec.VerifyRwxModifiers(true, []string{tmpFile})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch error", actual)
}

// ── RwxInstructionExecutors ──

func Test_Executors_Adds(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(5)

	ins1 := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
	}
	ins2 := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	e1, _ := chmodhelper.ParseRwxInstructionToExecutor(ins1)
	e2, _ := chmodhelper.ParseRwxInstructionToExecutor(ins2)

	executors.Adds(e1, e2)

	// Act
	actual := args.Map{"result": executors.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Executors_Adds_Nil(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(5)
	executors.Adds(nil)
	// nil items get appended (not skipped in Adds)
}

func Test_Executors_Length_Empty(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(0)
	l := executors.Length()

	// Act
	actual := args.Map{"result": l != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Executors_ApplyOnPath(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov10_exec_apply.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	executors := chmodhelper.NewRwxInstructionExecutors(2)
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	e, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	executors.Add(e)

	err := executors.ApplyOnPath(tmpFile)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executors_ApplyOnPath_Error(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(2)
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "r-x", Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid: false,
		},
	}
	e, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	executors.Add(e)

	err := executors.ApplyOnPath("/nonexistent/cov10/exec_err")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Executors_ApplyOnPaths(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(1)
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	e, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	executors.Add(e)

	err := executors.ApplyOnPaths([]string{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executors_ApplyOnPathsPtr(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov10_exec_ptr.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	executors := chmodhelper.NewRwxInstructionExecutors(1)
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	e, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	executors.Add(e)

	locs := []string{tmpFile}
	err := executors.ApplyOnPathsPtr(locs)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Executors_VerifyRwxModifiers(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov10_exec_verify.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	executors := chmodhelper.NewRwxInstructionExecutors(1)
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rw-", Group: "r--", Other: "r--",
		},
	}
	e, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	executors.Add(e)

	err := executors.VerifyRwxModifiers(false, true, []string{tmpFile})
	_ = err
}

func Test_Executors_VerifyRwxModifiers_ContinueOnErr(t *testing.T) {
	executors := chmodhelper.NewRwxInstructionExecutors(1)
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx", Group: "rwx", Other: "rwx",
		},
	}
	e, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)
	executors.Add(e)

	err := executors.VerifyRwxModifiers(true, true, []string{"/nonexistent/cov10/v1"})
	_ = err
}

func Test_Executors_VerifyRwxModifiers_Empty(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(1)
	err := executors.VerifyRwxModifiers(false, true, []string{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}
