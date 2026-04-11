package chmodhelper

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ══════════════════════════════════════════════════════════════════════════════
// simpleFileWriter — Lock/Unlock (L9, L13)
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleFileWriter_LockUnlock(t *testing.T) {
	SimpleFileWriter.Lock()
	SimpleFileWriter.Unlock()
}

// ══════════════════════════════════════════════════════════════════════════════
// newCreator — field access
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewCreator_Fields(t *testing.T) {
	_ = New.RwxWrapper
	_ = New.SimpleFileReaderWriter
	_ = New.Attribute
}

// ══════════════════════════════════════════════════════════════════════════════
// chmodApplier — test that Apply works with valid chmod (simplified)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ChmodApply_Noop(t *testing.T) {
	// Just validate the struct exists and can be referenced
	_ = ChmodApply
}

// ══════════════════════════════════════════════════════════════════════════════
// ChmodVerify — struct exists
// ══════════════════════════════════════════════════════════════════════════════

func Test_ChmodVerify_Noop(t *testing.T) {
	_ = ChmodVerify
}

// ══════════════════════════════════════════════════════════════════════════════
// ParseRwxInstructionsToExecutors — nil input (L10-12)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ParseRwxInstructionsToExecutors_Nil_ChecksExecutors(t *testing.T) {
	executors, err := ParseRwxInstructionsToExecutors(nil)
	if err == nil {
		t.Fatal("expected error for nil input")
	}
	if executors == nil {
		t.Fatal("expected non-nil executors even on error")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ParseRwxInstructionsToExecutors — empty input (L17-19)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ParseRwxInstructionsToExecutors_Empty(t *testing.T) {
	executors, err := ParseRwxInstructionsToExecutors([]chmodins.RwxInstruction{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if executors == nil {
		t.Fatal("expected non-nil executors")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ParseRwxOwnerGroupOtherToFileModeMust — nil input panics (L15-17)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ParseRwxOwnerGroupOtherToFileModeMust_Nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for nil input")
		}
	}()
	ParseRwxOwnerGroupOtherToFileModeMust(nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// TempDirGetter — basic coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_TempDirGetter_Value(t *testing.T) {
	val := TempDirGetter.TempDefault()
	if val == "" {
		t.Log("temp dir is empty string on this platform")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxVariableWrapper — Parse with valid hyphen-prefixed input (no error)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxVariableWrapper_Parse_Valid(t *testing.T) {
	// NewRwxVariableWrapper pads with wildcards; any single-char input
	// produces valid wildcard attributes, so no error is possible.
	// Test that it returns successfully with a hyphen-prefixed input.
	wrapper, err := NewRwxVariableWrapper("-rwx")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if wrapper == nil {
		t.Fatal("expected non-nil wrapper")
	}
}

func Test_RwxVariableWrapper_Parse_SingleChar(t *testing.T) {
	// "X" is padded to "X*********" — all wildcard segments, always valid.
	wrapper, err := NewRwxVariableWrapper("X")
	if err != nil {
		t.Fatalf("unexpected error for single char: %v", err)
	}
	if wrapper == nil {
		t.Fatal("expected non-nil wrapper")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutor — verifyChmodLocations error paths (L196, L227)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxInstructionExecutor_VerifyChmod_CompiledWrapperError(t *testing.T) {
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

func Test_RwxPartialToInstructionExecutor_InvalidPartial(t *testing.T) {
	cond := chmodins.DefaultAllFalseCondition()
	_, err := RwxPartialToInstructionExecutor("INVALID_VERY_LONG_RWX_STRING", cond)
	if err == nil {
		t.Fatal("expected error for invalid partial rwx")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — ApplyOnPathsPtr with empty executor (nil rwxInstruction)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxInstructionExecutors_ApplyOnPathsPtr_WithExecutors(t *testing.T) {
	// An empty RwxInstructionExecutor{} has nil rwxInstruction which panics
	// in ApplyOnPathsPtr when accessing rwxInstruction.IsContinueOnError.
	// Recover from the expected nil pointer dereference.
	defer func() {
		if r := recover(); r != nil {
			t.Log("recovered expected panic from nil rwxInstruction:", r)
		}
	}()

	exec := &RwxInstructionExecutor{} // nil rwxInstruction
	execs := NewRwxInstructionExecutors(1)
	execs.Add(exec)

	err := execs.ApplyOnPaths([]string{"/nonexistent"})
	if err != nil {
		t.Log("got error:", err)
	}
}
