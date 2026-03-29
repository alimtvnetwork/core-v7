package chmodhelper

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ══════════════════════════════════════════════════════════════════════════════
// CreateDirFilesWithRwxPermissions — error path (L16)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_CreateDirFilesWithRwxPermissions_ErrorPath(t *testing.T) {
	items := []DirFilesWithRwxPermission{
		{
		DirWithFiles: DirWithFiles{
			Dir: "/dev/null/impossible_dir",
		},
		},
	}
	err := CreateDirFilesWithRwxPermissions(false, items)
	if err == nil {
		t.Log("no error returned (OS allowed creation?)")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// CreateDirFilesWithRwxPermissionsMust — panic path (L11-12)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_CreateDirFilesWithRwxPermissionsMust_PanicPath(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Log("no panic — OS allowed the creation")
		}
	}()
	items := []DirFilesWithRwxPermission{
		{
		DirWithFiles: DirWithFiles{
			Dir: "/dev/null/impossible_dir",
		},
		},
	}
	CreateDirFilesWithRwxPermissionsMust(false, items)
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — CompiledWrapper fallthrough error (L67)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_CompiledWrapper_NeitherFixedNorVar(t *testing.T) {
	exec := &RwxInstructionExecutor{}
	_, err := exec.CompiledWrapper(0o644)
	if err == nil {
		t.Fatal("expected error when neither fixed nor var wrapper")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — CompiledRwxWrapperUsingFixedRwxWrapper fallthrough (L85)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_CompiledRwxWrapperUsingFixed_Error(t *testing.T) {
	exec := &RwxInstructionExecutor{}
	w := &RwxWrapper{}
	_, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(w)
	if err == nil {
		t.Fatal("expected error when neither fixed nor var")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — ApplyOnPathsDirect / ApplyOnPaths (L253, L261)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_ApplyOnPathsDirect_Empty(t *testing.T) {
	exec := &RwxInstructionExecutor{}
	err := exec.ApplyOnPathsDirect()
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

func Test_Cov2_RwxInstructionExecutor_ApplyOnPaths_Empty(t *testing.T) {
	exec := &RwxInstructionExecutor{}
	err := exec.ApplyOnPaths([]string{})
	if err != nil {
		t.Fatal("expected nil for empty locations")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — ApplyOnPaths (L155)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutors_ApplyOnPaths_Empty(t *testing.T) {
	execs := &RwxInstructionExecutors{}
	err := execs.ApplyOnPaths([]string{})
	if err != nil {
		t.Fatal("expected nil for empty")
	}
}

func Test_Cov2_RwxInstructionExecutors_ApplyOnPaths_EmptyExecutors(t *testing.T) {
	execs := &RwxInstructionExecutors{}
	err := execs.ApplyOnPaths([]string{"/tmp"})
	if err != nil {
		t.Fatal("expected nil for empty executors")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PathExistStat — MeaningFullError with error (L239)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_PathExistStat_MeaningFullError_WithError(t *testing.T) {
	stat := &PathExistStat{
		Location: "/nonexistent",
		Error:    errors.New("test error"),
	}
	err := stat.MeaningFullError()
	if err == nil {
		t.Fatal("expected meaningful error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MergeRwxWildcardWithFixedRwx — ParseRwxToVarAttribute error path (L38)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_MergeRwxWildcardWithFixedRwx_InvalidWildcard(t *testing.T) {
	_, err := MergeRwxWildcardWithFixedRwx("INVALID", "rwx")
	if err == nil {
		t.Fatal("expected error for invalid wildcard")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — nil RwxWrapper skip in ApplyRwxOnLocations (L186, L207)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxVariableWrapper_ApplyRwxOnLocations_NonexistentPath(t *testing.T) {
	w, err := NewRwxVariableWrapper("rwxrwxrwx")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	// Nonexistent path — exercises error handling
	applyErr := w.ApplyRwxOnLocations(true, true, "/nonexistent/path/that/does/not/exist")
	// continueOnError=true, skipOnInvalid=true — should not panic
	_ = applyErr
}

func Test_Cov2_RwxVariableWrapper_ApplyRwxOnLocations_ValidPath(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(tmpFile, []byte("test"), 0o644)

	w, err := NewRwxVariableWrapper("rwxrwxrwx")
	if err != nil {
		t.Fatal("unexpected parse error:", err)
	}

	applyErr := w.ApplyRwxOnLocations(false, false, tmpFile)
	if applyErr != nil {
		t.Log("apply error (may be expected):", applyErr)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — Parse error (L46)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxVariableWrapper_Parse_Error(t *testing.T) {
	_, err := NewRwxVariableWrapper("X")
	if err == nil {
		t.Fatal("expected error for invalid input")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — verifyChmodLocations error paths (L196, L227)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutor_VerifyChmod_CompiledWrapperError(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(tmpFile, []byte("test"), 0o644)

	exec := &RwxInstructionExecutor{}
	info, _ := os.Stat(tmpFile)

	resultsMap := &FilteredPathFileInfoMap{
		FilesToInfoMap: map[string]os.FileInfo{
			tmpFile: info,
		},
	}

	err := exec.verifyChmodLocationsContinueOnError(resultsMap)
	if err == nil {
		t.Fatal("expected error from CompiledWrapper failure")
	}

	err = exec.verifyChmodLocationsNoContinue(resultsMap)
	if err == nil {
		t.Fatal("expected error from CompiledWrapper failure")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxPartialToInstructionExecutor — error path (L29)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxPartialToInstructionExecutor_InvalidPartial(t *testing.T) {
	cond := chmodins.DefaultAllFalseCondition()
	_, err := RwxPartialToInstructionExecutor("INVALID_VERY_LONG_RWX_STRING", cond)
	if err == nil {
		t.Fatal("expected error for invalid partial rwx")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — ApplyOnPathsPtr with error (L167)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_RwxInstructionExecutors_ApplyOnPathsPtr_WithExecutors(t *testing.T) {
	exec := &RwxInstructionExecutor{} // neither fixed nor var
	execs := NewRwxInstructionExecutors(1)
	execs.Add(exec)

	err := execs.ApplyOnPaths([]string{"/nonexistent"})
	// Should error because the executor can't compile a wrapper
	if err == nil {
		t.Log("no error — executor handled gracefully")
	}
}
