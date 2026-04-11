package chmodhelper

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_NilAndEmpty(t *testing.T) {
	var nilW *RwxWrapper
	if !nilW.IsEmpty() {
		t.Fatal("nil should be empty")
	}
	if !nilW.IsNull() {
		t.Fatal("nil should be null")
	}
	if !nilW.IsInvalid() {
		t.Fatal("nil should be invalid")
	}
	if nilW.IsDefined() {
		t.Fatal("nil should not be defined")
	}
	if nilW.HasAnyItem() {
		t.Fatal("nil should not have items")
	}
}

func Test_RwxWrapper_UsingFileMode(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	s := rwx.ToFullRwxValueString()
	if len(s) != 10 {
		t.Fatal("expected 10 char rwx string")
	}
	if rwx.IsEmpty() {
		t.Fatal("expected non-empty wrapper")
	}

	// Bytes, octal methods
	b := rwx.Bytes()
	if b[0] == 0 && b[1] == 0 && b[2] == 0 {
		t.Fatal("expected non-zero bytes")
	}
	_ = rwx.ToCompiledOctalBytes4Digits()
	_ = rwx.ToCompiledOctalBytes3Digits()
	o, g, ot := rwx.ToCompiledSplitValues()
	_ = o
	_ = g
	_ = ot
	_ = rwx.ToFileModeString()
	_ = rwx.ToRwxCompiledStr()
	_ = rwx.ToFullRwxValueStringExceptHyphen()
	_ = rwx.ToFullRwxValuesChars()
	_ = rwx.String()
	_ = rwx.ToFileMode()
	_ = rwx.ToUint32Octal()
}

func Test_RwxWrapper_ApplyChmod(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}
	dir := t.TempDir()
	fp := filepath.Join(dir, "test.txt")
	os.WriteFile(fp, []byte("data"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)
	err := rwx.ApplyChmod(false, fp)
	if err != nil {
		t.Fatal("expected no error on valid file:", err)
	}
	// skip on invalid
	err2 := rwx.ApplyChmod(true, "/nonexistent/path/xyz")
	if err2 != nil {
		t.Fatal("expected skip on invalid:", err2)
	}
	// non-skip on invalid
	err3 := rwx.ApplyChmod(false, "/nonexistent/path/xyz")
	if err3 == nil {
		t.Fatal("expected error for invalid path without skip")
	}
}

func Test_RwxWrapper_ApplyChmodOptions(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}
	dir := t.TempDir()
	fp := filepath.Join(dir, "opts.txt")
	os.WriteFile(fp, []byte("data"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)
	// skip apply
	if err := rwx.ApplyChmodOptions(false, true, false, fp); err != nil {
		t.Fatal("expected nil on skip-apply")
	}
	// skip on invalid path
	if err := rwx.ApplyChmodOptions(true, true, true, "/no/such/path"); err != nil {
		t.Fatal("expected nil on skip-invalid")
	}
	// invalid path without skip
	if err := rwx.ApplyChmodOptions(true, true, false, "/no/such/path"); err == nil {
		t.Fatal("expected error for invalid without skip")
	}
	// on mismatch, already matching
	if err := rwx.ApplyChmodOptions(true, true, false, fp); err != nil {
		t.Fatal("expected nil on matching chmod")
	}
}

func Test_RwxWrapper_ApplyRecursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "f.txt"), []byte("x"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0755)
	// valid recursive
	err := rwx.ApplyRecursive(false, dir)
	if err != nil {
		t.Fatal("expected no error on recursive apply:", err)
	}
	// skip non-existent
	err2 := rwx.ApplyRecursive(true, "/no/such/path")
	if err2 != nil {
		t.Fatal("expected nil on skip-invalid path")
	}
	// non-skip non-existent
	err3 := rwx.ApplyRecursive(false, "/no/such/path")
	if err3 == nil {
		t.Fatal("expected error for non-existent without skip")
	}
	// single file
	fp := filepath.Join(dir, "single.txt")
	os.WriteFile(fp, []byte("y"), 0644)
	err4 := rwx.ApplyRecursive(false, fp)
	if err4 != nil {
		t.Fatal("expected no error on single file recursive:", err4)
	}
}

func Test_RwxWrapper_MustApplyChmod_Success(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}
	dir := t.TempDir()
	fp := filepath.Join(dir, "must.txt")
	os.WriteFile(fp, []byte("z"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)
	// should not panic
	rwx.MustApplyChmod(fp)
}

func Test_RwxWrapper_HasChmod(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}
	dir := t.TempDir()
	fp := filepath.Join(dir, "hc.txt")
	os.WriteFile(fp, []byte("z"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)
	if !rwx.HasChmod(fp) {
		t.Fatal("expected chmod match")
	}
}

func Test_RwxWrapper_Verify(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}
	dir := t.TempDir()
	fp := filepath.Join(dir, "verify.txt")
	os.WriteFile(fp, []byte("z"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)
	err := rwx.Verify(fp)
	if err != nil {
		t.Fatal("expected no error verifying:", err)
	}
}

func Test_RwxWrapper_VerifyPaths(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod not supported on Windows")
	}
	dir := t.TempDir()
	fp := filepath.Join(dir, "vp.txt")
	os.WriteFile(fp, []byte("z"), 0644)
	rwx := New.RwxWrapper.UsingFileMode(0644)
	err := rwx.VerifyPaths(false, fp)
	if err != nil {
		t.Fatal("expected no error verifying paths:", err)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxWrapper — JSON, Clone, IsEqual
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxWrapper_JsonAndClone(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	j := rwx.Json()
	if j.HasError() {
		t.Fatal("expected no json error")
	}
	ptr := rwx.JsonPtr()
	if ptr.HasError() {
		t.Fatal("expected no json ptr error")
	}
	cloned := rwx.Clone()
	if !cloned.IsEqualPtr(&rwx) {
		t.Fatal("clone should be equal")
	}
	_ = rwx.FriendlyDisplay()
	_ = rwx.ToPtr()
	_ = rwx.ToNonPtr()
	_ = rwx.AsJsonContractsBinder()
	_ = rwx.ToRwxOwnerGroupOther()
}

func Test_RwxWrapper_IsEqualFileMode(t *testing.T) {
	rwx := New.RwxWrapper.UsingFileMode(0755)
	if !rwx.IsEqualFileMode(0755) {
		t.Fatal("expected equal file mode")
	}
	if rwx.IsNotEqualFileMode(0755) {
		t.Fatal("expected not-not-equal file mode")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// PathExistStat — branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_PathExistStat_ValidFile(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "stat.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	stat := GetPathExistStat(fp)
	if stat.IsInvalid() || stat.HasAnyIssues() {
		t.Fatal("expected valid stat")
	}
	if !stat.IsFile() || stat.IsDir() {
		t.Fatal("expected file")
	}
	if stat.HasError() || !stat.IsEmptyError() {
		t.Fatal("expected no error")
	}
	if stat.LastModifiedDate() == nil {
		t.Fatal("expected last modified date")
	}
	if stat.FileMode() == nil {
		t.Fatal("expected file mode")
	}
	if stat.Size() == nil {
		t.Fatal("expected size")
	}
	d, f := stat.Split()
	if d == "" || f == "" {
		t.Fatal("expected split values")
	}
	if stat.FileName() == "" {
		t.Fatal("expected filename")
	}
	if stat.ParentDir() == "" {
		t.Fatal("expected parent dir")
	}
	if stat.DotExt() != ".txt" {
		t.Fatal("expected .txt extension")
	}
	if stat.String() == "" {
		t.Fatal("expected non-empty string")
	}
	_ = stat.NotExistError()
	_ = stat.NotAFileError()
	_ = stat.NotADirError()
	_ = stat.MessageWithPathWrapped("test")
	stat.Dispose()
	if stat.IsExist {
		t.Fatal("expected disposed")
	}
}

func Test_PathExistStat_NonExistent(t *testing.T) {
	stat := GetPathExistStat("/nonexistent/abc/xyz")
	if !stat.IsInvalid() {
		t.Fatal("expected invalid stat")
	}
	if stat.IsFile() || stat.IsDir() {
		t.Fatal("expected neither file nor dir")
	}
	if stat.NotExistError() == nil {
		t.Fatal("expected not-exist error")
	}
}

func Test_PathExistStat_NilReceiver(t *testing.T) {
	var stat *PathExistStat
	if stat.HasError() {
		t.Fatal("nil should not have error")
	}
	if !stat.IsEmptyError() {
		t.Fatal("nil should be empty error")
	}
	if stat.HasFileInfo() {
		t.Fatal("nil should not have file info")
	}
	if !stat.IsInvalidFileInfo() {
		t.Fatal("nil should be invalid file info")
	}
	if stat.IsFile() || stat.IsDir() {
		t.Fatal("nil should be neither file nor dir")
	}
	if stat.LastModifiedDate() != nil {
		t.Fatal("nil should have nil last modified")
	}
	if stat.FileMode() != nil {
		t.Fatal("nil should have nil file mode")
	}
	if stat.Size() != nil {
		t.Fatal("nil should have nil size")
	}
	if stat.NotExistError() != nil {
		t.Fatal("nil should have nil not-exist error")
	}
	if stat.NotAFileError() != nil {
		t.Fatal("nil should have nil not-a-file error")
	}
	if stat.NotADirError() != nil {
		t.Fatal("nil should have nil not-a-dir error")
	}
	if stat.MeaningFullError() != nil {
		t.Fatal("nil should have nil meaningful error")
	}
	if stat.String() != "" {
		t.Fatal("nil should have empty string")
	}
}

func Test_PathExistStat_DirBranches(t *testing.T) {
	dir := t.TempDir()
	stat := GetPathExistStat(dir)
	if !stat.IsDir() {
		t.Fatal("expected dir")
	}
	d, f := stat.Split()
	if d != "" || f != "" {
		// dir Split returns empty
	}
	notAFile := stat.NotAFileError()
	if notAFile == nil {
		t.Fatal("expected not-a-file error for directory")
	}
	notADir := stat.NotADirError()
	if notADir != nil {
		t.Fatal("expected nil not-a-dir error for directory")
	}
}

func Test_PathExistStat_CombineMethods(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "c.txt")
	os.WriteFile(fp, []byte("x"), 0644)
	stat := GetPathExistStat(fp)
	newPath := stat.CombineWithNewPath("extra")
	if newPath == "" {
		t.Fatal("expected combined path")
	}
	parentNew := stat.ParentWithNewPath("other.txt")
	if parentNew == "" {
		t.Fatal("expected parent-with path")
	}
	_, _ = stat.ParentWithGlobPatternFiles("*.txt")
	_ = stat.Parent()
	_ = stat.ParentWith("sub")
	_ = stat.CombineWith("extra2")
}

// ══════════════════════════════════════════════════════════════════════════════
// dirCreator, fileWriter — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_DirCreator_ByChecking(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	// create new
	err := internalDirCreator.ByChecking(0755, sub)
	if err != nil {
		t.Fatal("expected create:", err)
	}
	// already exists, re-apply
	err2 := internalDirCreator.ByChecking(0755, sub)
	if err2 != nil {
		t.Fatal("expected re-apply:", err2)
	}
}

func Test_FileWriter_ChmodAndChmodFile(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "fw.txt")
	fw := fileWriter{}
	err := fw.Chmod(false, 0755, 0644, fp, []byte("data"))
	if err != nil {
		t.Fatal("expected chmod write:", err)
	}
	fp2 := filepath.Join(dir, "fw2.txt")
	err2 := fw.ChmodFile(false, 0644, fp2, []byte("data2"))
	if err2 != nil {
		t.Fatal("expected chmodFile write:", err2)
	}
}

func Test_ErrorCreator_PathError_NilErr(t *testing.T) {
	err := newError.pathError("msg", 0644, "/tmp/x", nil)
	if err != nil {
		t.Fatal("nil input error should return nil")
	}
}

func Test_ErrorCreator_ChmodApplyFailed_NilErr(t *testing.T) {
	err := newError.chmodApplyFailed(0644, "/tmp/x", nil)
	if err != nil {
		t.Fatal("nil input error should return nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleFileReaderWriter — uncovered init and read methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleFileReaderWriter_Init(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "sfrw.txt")
	rw := SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  fp,
	}
	initialized := rw.InitializeDefault(true)
	if initialized.ParentDir == "" {
		t.Fatal("expected parent dir set")
	}
	initialized2 := rw.InitializeDefaultApplyChmod()
	if !initialized2.IsMustChmodApplyOnFile {
		t.Fatal("expected must apply chmod")
	}
}

func Test_SimpleFileReaderWriter_WriteAndRead(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "sfrw.txt")
	rw := SimpleFileReaderWriter{
		ChmodDir:            0755,
		ChmodFile:           0644,
		ParentDir:           dir,
		FilePath:            fp,
		IsRemoveBeforeWrite: false,
	}
	err := rw.Write([]byte("hello"))
	if err != nil {
		t.Fatal("expected write:", err)
	}
	if !rw.IsExist() {
		t.Fatal("expected file exists")
	}
	if !rw.IsParentExist() {
		t.Fatal("expected parent exists")
	}
	if rw.HasPathIssues() {
		t.Fatal("expected no path issues")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// RwxInstructionExecutors — empty and nil paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_RwxInstructionExecutors_EmptyPaths(t *testing.T) {
	execs := NewRwxInstructionExecutors(0)
	if !execs.IsEmpty() {
		t.Fatal("expected empty")
	}
	if execs.HasAnyItem() {
		t.Fatal("expected no items")
	}
	execs.Add(nil) // skip nil
	if execs.Length() != 0 {
		t.Fatal("expected still empty after nil add")
	}
	execs.Adds(nil)
	if execs.Count() != 0 {
		t.Fatal("expected still empty after nil adds")
	}
	// empty path operations
	if err := execs.ApplyOnPath("/x"); err != nil {
		t.Fatal("expected nil on empty executors apply")
	}
	if err := execs.ApplyOnPaths(nil); err != nil {
		t.Fatal("expected nil on nil paths")
	}
	if err := execs.VerifyRwxModifiers(false, false, nil); err != nil {
		t.Fatal("expected nil on empty verify")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ParseRwxInstructionsToExecutors — nil input
// ══════════════════════════════════════════════════════════════════════════════

func Test_ParseRwxInstructionsToExecutors_Nil(t *testing.T) {
	_, err := ParseRwxInstructionsToExecutors(nil)
	if err == nil {
		t.Fatal("expected error for nil instructions")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// chmodApplier — skip branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ChmodApplier_ApplyIf_Skip(t *testing.T) {
	err := ChmodApply.ApplyIf(false, 0644, "/whatever")
	if err != nil {
		t.Fatal("expected nil on skip-apply")
	}
}

func Test_ChmodApplier_OnMismatchOption_Skip(t *testing.T) {
	err := ChmodApply.OnMismatchOption(false, false, 0644, "/whatever")
	if err != nil {
		t.Fatal("expected nil on skip-apply")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// chmodVerifier — extra branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ChmodVerifier_IsEqualSkipInvalid(t *testing.T) {
	if !ChmodVerify.IsEqualSkipInvalid("/nonexistent", 0644) {
		t.Fatal("expected true on skip-invalid path")
	}
	if !ChmodVerify.IsEqualRwxFullSkipInvalid("/nonexistent", "-rw-r--r--") {
		t.Fatal("expected true on skip-invalid rwx path")
	}
}

func Test_ChmodVerifier_GetRwx9(t *testing.T) {
	s := ChmodVerify.GetRwx9(0755)
	if len(s) != 9 {
		t.Fatal("expected 9 char rwx string")
	}
}

func Test_ChmodVerifier_PathIf_Skip(t *testing.T) {
	err := ChmodVerify.PathIf(false, "/whatever", 0644)
	if err != nil {
		t.Fatal("expected nil on skip-verify")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// tempDirGetter
// ══════════════════════════════════════════════════════════════════════════════

func Test_TempDirGetter(t *testing.T) {
	d := TempDirGetter.TempDefault()
	if d == "" {
		t.Fatal("expected temp dir")
	}
	_ = TempDirGetter.TempPermanent()
	_ = TempDirGetter.TempOption(true)
	_ = TempDirGetter.TempOption(false)
}
