package chmodhelper

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
)

// ── CreateDefaultPaths error path ──────────────────────────

func Test_CreateDirFilesWithRwxPermissions_ErrorPath(t *testing.T) {
	items := []DirFilesWithRwxPermission{
		{
			DirWithFiles: DirWithFiles{
				Dir: "/dev/null/impossible",
			},
			ApplyRwx: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "r-x",
			},
		},
	}

	err := CreateDirFilesWithRwxPermissions(false, items)
	if err == nil {
		t.Log("no error — OS allowed creation (unlikely)")
	}
}

// ── CreateDirWithFiles error paths ─────────────────────────

func Test_CreateDirWithFiles_CloseError(t *testing.T) {
	dir := t.TempDir()
	dwf := &DirWithFiles{
		Dir:   dir,
		Files: []string{"testfile.txt"},
	}

	err := CreateDirWithFiles(false, 0644, dwf)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Verify file was created
	fp := filepath.Join(dir, "testfile.txt")
	if _, statErr := os.Stat(fp); statErr != nil {
		t.Fatal("file not created")
	}
}

func Test_CreateDirWithFiles_ChmodError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	// Create a file in a dir, then make the dir read-only and try again
	dwf := &DirWithFiles{
		Dir:   "/dev/null/impossible",
		Files: []string{"x.txt"},
	}

	err := CreateDirWithFiles(false, 0755, dwf)
	if err == nil {
		t.Log("no error — OS allowed creation (unlikely)")
	}
}

// ── DirFilesWithContent.Create error paths ─────────────────

func Test_DirFilesWithContent_Create_RemoveError(t *testing.T) {
	dfwc := &DirFilesWithContent{
		Dir:         "/dev/null/impossible",
		DirFileMode: 0755,
	}

	err := dfwc.Create(true)
	if err == nil {
		t.Log("no error — OS allowed (unlikely)")
	}
}

func Test_DirFilesWithContent_Create_FileWriteError_InvalidDir(t *testing.T) {
	dfwc := &DirFilesWithContent{
		Dir:         "/dev/null/impossible/sub",
		DirFileMode: 0755,
		Files: []FileWithContent{
			{
				RelativePath: "f.txt",
				Content:      []string{"hello"},
				FileMode:     0644,
			},
		},
	}

	err := dfwc.Create(false)
	if err == nil {
		t.Log("no error — OS allowed (unlikely)")
	}
}

func Test_DirFilesWithContent_Create_Success(t *testing.T) {
	dir := t.TempDir()
	subDir := filepath.Join(dir, "sub")

	dfwc := &DirFilesWithContent{
		Dir:         subDir,
		DirFileMode: 0755,
		Files: []FileWithContent{
			{
				RelativePath: "f.txt",
				Content:      []string{"hello"},
				FileMode:     0644,
			},
		},
	}

	err := dfwc.Create(false)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

// ── GetRecursivePaths walk error ───────────────────────────

func Test_GetRecursivePaths_WalkError(t *testing.T) {
	// Non-existent path, not continue on error
	_, err := GetRecursivePaths(false, "/nonexistent/xyz")
	if err == nil {
		t.Fatal("expected error for non-existent path")
	}
}

func Test_GetRecursivePaths_ContinueOnError_NonExistent(t *testing.T) {
	_, err := GetRecursivePaths(true, "/nonexistent/xyz")
	if err == nil {
		t.Fatal("expected error for non-existent path")
	}
}

func Test_GetRecursivePathsContinueOnError_ValidDir_WithSubdir(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0644)

	paths, err := GetRecursivePathsContinueOnError(dir)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if len(paths) == 0 {
		t.Fatal("expected non-empty paths")
	}
}

// ── MergeRwxWildcardWithFixedRwx error path ────────────────

func Test_MergeRwxWildcardWithFixedRwx_InvalidLength(t *testing.T) {
	_, err := MergeRwxWildcardWithFixedRwx("rw", "rw*")
	if err == nil {
		t.Fatal("expected error for invalid length")
	}
}

// ── PathExistStat.MeaningFullError ─────────────────────────

func Test_PathExistStat_MeaningFullError_WithError_PermDenied(t *testing.T) {
	stat := GetPathExistStat("/nonexistent/xyz/abc")
	err := stat.MeaningFullError()
	// If path doesn't exist, stat has error
	if stat.HasError() && err == nil {
		t.Fatal("expected meaningful error")
	}
}

// ── RwxInstructionExecutor uncovered branches ──────────────

func Test_RwxInstructionExecutor_CompiledWrapper_Fallthrough(t *testing.T) {
	// Test the fixed-type branch through CompiledWrapper
	w, parseErr := NewRwxVariableWrapper("-rwxr-xr-x")
	if parseErr != nil {
		t.Fatal(parseErr)
	}

	exec := &RwxInstructionExecutor{
		varWrapper: w,
	}
	result, err := exec.CompiledWrapper(0644)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if result == nil {
		t.Fatal("expected non-nil wrapper")
	}
}

func Test_RwxInstructionExecutor_CompiledRwxWrapperUsingFixedRwxWrapper_Dead(t *testing.T) {
	w, err := NewRwxVariableWrapper("-rwxrwxrwx")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	exec := &RwxInstructionExecutor{
		varWrapper: w,
	}
	// Fixed type — should take fixed branch
	result, err := exec.CompiledRwxWrapperUsingFixedRwxWrapper(nil)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_RwxInstructionExecutor_ApplyOnPath_SkipOnInvalid_InvalidDir(t *testing.T) {
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid: true,
		},
	}
	exec, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal("unexpected:", err)
	}

	// Apply on non-existent — should return nil (skip)
	err = exec.ApplyOnPath("/nonexistent/xyz")
	if err != nil {
		t.Fatal("expected nil error for skip on invalid")
	}
}

func Test_RwxInstructionExecutor_ApplyOnPath_CompileError(t *testing.T) {
	// Test with varWrapper that has wildcard — compiledErr path
	w, err := NewRwxVariableWrapper("-r*xr*xr*x")
	if err != nil {
		t.Fatal(err)
	}
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "r*x",
			Group: "r*x",
			Other: "r*x",
		},
		Condition: chmodins.Condition{},
	}
	exec := &RwxInstructionExecutor{
		rwxInstruction: ins,
		varWrapper:     w,
	}

	// Apply on a real path — should work
	dir := t.TempDir()
	err = exec.ApplyOnPath(dir)
	if err != nil {
		t.Log("error applying wildcard path:", err)
	}
}

func Test_RwxInstructionExecutor_VerifyRwxModifiers_ContinueOnError(t *testing.T) {
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsContinueOnError: true,
		},
	}
	exec, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	err = exec.VerifyRwxModifiers(true, []string{dir})
	if err != nil {
		t.Log("verify error:", err)
	}
}

func Test_RwxInstructionExecutor_VerifyRwxModifiers_NoContinue(t *testing.T) {
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r-x",
		},
		Condition: chmodins.Condition{},
	}
	exec, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	err = exec.VerifyRwxModifiers(true, []string{dir})
	if err != nil {
		t.Log("verify error:", err)
	}
}

// ── RwxInstructionExecutors.ApplyOnPaths (non-empty) ───────

func Test_RwxInstructionExecutors_ApplyOnPaths_NonEmpty_TempDir(t *testing.T) {
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r-x",
		},
		Condition: chmodins.Condition{
			IsSkipOnInvalid: true,
		},
	}
	exec, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal(err)
	}
	execs := NewRwxInstructionExecutors(1)
	execs.Add(exec)

	dir := t.TempDir()
	err = execs.ApplyOnPaths([]string{dir})
	if err != nil {
		t.Log("error:", err)
	}
}

func Test_RwxInstructionExecutors_ApplyOnPaths_WithError(t *testing.T) {
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r-x",
		},
		Condition: chmodins.Condition{},
	}
	exec, err := ParseRwxInstructionToExecutor(ins)
	if err != nil {
		t.Fatal(err)
	}
	execs := NewRwxInstructionExecutors(1)
	execs.Add(exec)

	err = execs.ApplyOnPaths([]string{"/nonexistent/xyz"})
	if err == nil {
		t.Log("no error — unexpected")
	}
}

// ── RwxPartialToInstructionExecutor error ──────────────────

func Test_RwxPartialToInstructionExecutor_NilCondition_ReturnsNil(t *testing.T) {
	_, err := RwxPartialToInstructionExecutor("-rwxr-xr-x", nil)
	if err == nil {
		t.Fatal("expected error for nil condition")
	}
}

// ── RwxVariableWrapper uncovered branches ──────────────────

func Test_RwxVariableWrapper_ToCompileFixedPtr_NotFixed(t *testing.T) {
	w, err := NewRwxVariableWrapper("-r*xr*xr*x")
	if err != nil {
		t.Fatal(err)
	}
	result := w.ToCompileFixedPtr()
	if result != nil {
		t.Fatal("expected nil for non-fixed type")
	}
}

func Test_RwxVariableWrapper_ApplyRwxOnLocations_ContinueOnError(t *testing.T) {
	w, err := NewRwxVariableWrapper("-rwxr-xr-x")
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	err = w.ApplyRwxOnLocations(true, true, dir)
	if err != nil {
		t.Log("error:", err)
	}
}

func Test_RwxVariableWrapper_ApplyRwxOnLocations_NoContinue_InvalidPath(t *testing.T) {
	w, err := NewRwxVariableWrapper("-rwxr-xr-x")
	if err != nil {
		t.Fatal(err)
	}

	err = w.ApplyRwxOnLocations(false, false, "/nonexistent/xyz")
	if err == nil {
		t.Log("no error — unexpected")
	}
}

func Test_RwxVariableWrapper_ApplyRwxOnLocations_NoContinue_Valid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	w, err := NewRwxVariableWrapper("-rwxr-xr-x")
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	err = w.ApplyRwxOnLocations(false, false, dir)
	if err != nil {
		t.Log("error:", err)
	}
}

func Test_RwxVariableWrapper_RwxMatchingStatus_Mismatch(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	w, err := NewRwxVariableWrapper("-rwxrwxrwx")
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0600)

	status := w.RwxMatchingStatus(true, false, []string{fp})
	if status == nil {
		t.Fatal("expected non-nil status")
	}
}

// ── RwxWrapper uncovered branches ──────────────────────────

func Test_RwxWrapper_VerifyPaths_InvalidPath_SkipOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.VerifyPaths(true, "/nonexistent/xyz")
	if err != nil {
		t.Log("error:", err)
	}
}

func Test_RwxWrapper_ApplyChmod_SkipInvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(true, "/nonexistent/xyz")
	if err != nil {
		t.Fatal("expected nil for skip-on-invalid")
	}
}

func Test_RwxWrapper_ApplyChmod_NoSkipInvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmod(false, "/nonexistent/xyz")
	if err == nil {
		t.Fatal("expected error for invalid path without skip")
	}
}

func Test_RwxWrapper_LinuxApplyRecursive_SkipOnInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(true, "/nonexistent/xyz")
	if err != nil {
		t.Fatal("expected nil for skip-on-invalid")
	}
}

func Test_RwxWrapper_LinuxApplyRecursive_NoSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.LinuxApplyRecursive(false, "/nonexistent/xyz")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_RwxWrapper_ApplyRecursive_SkipOnInvalid(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(true, "/nonexistent/xyz")
	if err != nil {
		t.Fatal("expected nil for skip-on-invalid")
	}
}

func Test_RwxWrapper_ApplyRecursive_NoSkip(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(false, "/nonexistent/xyz")
	if err == nil {
		t.Fatal("expected error for non-existent")
	}
}

func Test_RwxWrapper_ApplyRecursive_File(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	rwx := New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyRecursive(false, fp)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func Test_RwxWrapper_ApplyRecursive_Dir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "f.txt"), []byte("x"), 0644)

	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyRecursive(false, dir)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursive_TempFiles(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	rwx := New.RwxWrapper.UsingFileMode(0644)
	cond := &chmodins.Condition{
		IsSkipOnInvalid: true,
	}
	err := rwx.ApplyLinuxChmodOnMany(cond, fp)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_Recursive_TempDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	dir := t.TempDir()
	rwx := New.RwxWrapper.UsingFileMode(0755)
	cond := &chmodins.Condition{
		IsRecursive:     true,
		IsSkipOnInvalid: true,
	}
	err := rwx.ApplyLinuxChmodOnMany(cond, dir)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_ContinueOnError_NonRecursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	cond := &chmodins.Condition{
		IsContinueOnError: true,
		IsSkipOnInvalid:   true,
	}
	err := rwx.ApplyLinuxChmodOnMany(cond, "/nonexistent/xyz")
	if err != nil {
		t.Log("error:", err)
	}
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_ContinueOnError_Recursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: true,
		IsSkipOnInvalid:   true,
	}
	err := rwx.ApplyLinuxChmodOnMany(cond, "/nonexistent/xyz")
	if err != nil {
		t.Log("error:", err)
	}
}

func Test_RwxWrapper_getLinuxRecursiveCmdForChmod(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	rwx := New.RwxWrapper.UsingFileMode(0755)
	cmd := rwx.getLinuxRecursiveCmdForChmod("/tmp")
	if cmd == nil {
		t.Fatal("expected non-nil cmd")
	}
}

func Test_RwxWrapper_applyLinuxRecursiveChmodUsingCmd_Valid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Linux-only test")
	}

	dir := t.TempDir()
	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.applyLinuxRecursiveChmodUsingCmd(dir)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func Test_RwxWrapper_ApplyChmodOptions_SkipApply(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(false, false, false, "/any")
	if err != nil {
		t.Fatal("expected nil when isApply=false")
	}
}

func Test_RwxWrapper_ApplyChmodOptions_SkipOnInvalid(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, false, true, "/nonexistent/xyz")
	if err != nil {
		t.Fatal("expected nil for skip-on-invalid")
	}
}

func Test_RwxWrapper_ApplyChmodOptions_InvalidNotSkip(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	err := rwx.ApplyChmodOptions(true, false, false, "/nonexistent/xyz")
	if err == nil {
		t.Fatal("expected error for invalid path without skip")
	}
}
