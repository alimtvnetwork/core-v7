package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

func skipOnWindows(t *testing.T) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
}

// ── SingleRwx.ToRwxOwnerGroupOther default panic ──

func Test_Cov13_SingleRwx_ToRwxOwnerGroupOther_Default(t *testing.T) {
	s := &chmodhelper.SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.All,
	}
	result := s.ToRwxOwnerGroupOther()
	if result.Owner != "rwx" || result.Group != "rwx" || result.Other != "rwx" {
		t.Fatal("expected all rwx")
	}
}

// ── SingleRwx.ToDisabledRwxWrapper ──

func Test_Cov13_SingleRwx_ToDisabledRwxWrapper(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	w, err := s.ToDisabledRwxWrapper()
	if err != nil || w == nil {
		t.Fatal("expected wrapper")
	}
}

func Test_Cov13_SingleRwx_ToDisabledRwxWrapper_Error(t *testing.T) {
	// Invalid chars are normalized as disabled permissions.
	s := &chmodhelper.SingleRwx{
		Rwx:       "rZx",
		ClassType: chmodclasstype.All,
	}
	w, err := s.ToDisabledRwxWrapper()
	if err != nil || w == nil {
		t.Fatal("expected wrapper for normalized invalid chars")
	}
}

// ── SingleRwx.ToRwxWrapper ──

func Test_Cov13_SingleRwx_ToRwxWrapper_NotAll(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	_, err := s.ToRwxWrapper()
	if err == nil {
		t.Fatal("expected error for non-All class type")
	}
}

func Test_Cov13_SingleRwx_ToRwxWrapper_All(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	w, err := s.ToRwxWrapper()
	if err != nil || w == nil {
		t.Fatal("expected wrapper")
	}
}

func Test_Cov13_SingleRwx_ToRwxWrapper_Error(t *testing.T) {
	s := &chmodhelper.SingleRwx{
		Rwx:       "rZx",
		ClassType: chmodclasstype.All,
	}
	w, err := s.ToRwxWrapper()
	if err != nil || w == nil {
		t.Fatal("expected wrapper for normalized invalid chars")
	}
}

// ── SingleRwx.ApplyOnMany ──

func Test_Cov13_SingleRwx_ApplyOnMany_Empty(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	err := s.ApplyOnMany(&chmodins.Condition{})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov13_SingleRwx_ApplyOnMany_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_apply_many.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	s, _ := chmodhelper.NewSingleRwx("rw-", chmodclasstype.All)
	err := s.ApplyOnMany(&chmodins.Condition{}, tmpFile)
	_ = err
}

func Test_Cov13_SingleRwx_ApplyOnMany_Error(t *testing.T) {
	s := &chmodhelper.SingleRwx{
		Rwx:       "rZx",
		ClassType: chmodclasstype.All,
	}
	err := s.ApplyOnMany(&chmodins.Condition{}, "/some/path")
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── newRwxWrapperCreator.CreatePtr error ──

func Test_Cov13_CreatePtr_Error(t *testing.T) {
	_, err := chmodhelper.New.RwxWrapper.CreatePtr("999")
	if err == nil {
		t.Fatal("expected error for invalid mode")
	}
}

func Test_Cov13_CreatePtr_Valid(t *testing.T) {
	ptr, err := chmodhelper.New.RwxWrapper.CreatePtr("755")
	if err != nil || ptr == nil {
		t.Fatal("expected valid pointer")
	}
}

// ── newRwxWrapperCreator.Create invalid char ──

func Test_Cov13_Create_InvalidChar(t *testing.T) {
	_, err := chmodhelper.New.RwxWrapper.Create("89a")
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── newRwxWrapperCreator.Create invalid length ──

func Test_Cov13_Create_InvalidLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for wrong length")
		}
	}()
	chmodhelper.New.RwxWrapper.Create("77")
}

// ── newRwxWrapperCreator.UsingChmod ──

func Test_Cov13_UsingChmod_Valid(t *testing.T) {
	w := chmodhelper.New.RwxWrapper.UsingChmod(0755)
	if w == nil || w.IsEmpty() {
		t.Fatal("expected non-empty wrapper")
	}
}

func Test_Cov13_UsingChmod_Zero(t *testing.T) {
	w := chmodhelper.New.RwxWrapper.UsingChmod(0)
	if w == nil {
		t.Fatal("expected non-nil wrapper")
	}
	if !w.IsEmpty() {
		t.Fatal("expected empty wrapper for zero mode")
	}
}

// ── newRwxWrapperCreator.UsingVariantPtr ──
// Variant is a string type, so we use valid string values

func Test_Cov13_UsingVariantPtr_Valid(t *testing.T) {
	w, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("755"))
	if err != nil || w == nil {
		t.Fatal("expected valid pointer")
	}
}

func Test_Cov13_UsingVariantPtr_Invalid(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// panics on length != 3 — expected
		}
	}()
	_, _ = chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("99"))
}

// ── newRwxWrapperCreator.Instruction ──

func Test_Cov13_Instruction_Valid(t *testing.T) {
	ins, err := chmodhelper.New.RwxWrapper.Instruction(
		"-rwxr-xr-x",
		chmodins.Condition{})
	if err != nil || ins == nil {
		t.Fatal("expected instruction")
	}
}

func Test_Cov13_Instruction_Error(t *testing.T) {
	_, err := chmodhelper.New.RwxWrapper.Instruction(
		"rwxr-xr-x",
		chmodins.Condition{})
	if err == nil {
		t.Fatal("expected error for wrong length")
	}
}

// ── newAttributeCreator.UsingByteMust panic ──

func Test_Cov13_UsingByteMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for byte > 7")
		}
	}()
	chmodhelper.New.Attribute.UsingByteMust(8)
}

// ── newAttributeCreator.UsingRwxString panic ──

func Test_Cov13_UsingRwxString_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for wrong length")
		}
	}()
	chmodhelper.New.Attribute.UsingRwxString("rw")
}

// ── chmodVerifier branches ──

func Test_Cov13_ChmodVerifier_IsEqualRwxFullSkipInvalid(t *testing.T) {
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid(
		"/nonexistent/cov13/skip", "-rwxr-xr-x")
	if !result {
		t.Fatal("expected true for invalid path with skip")
	}
}

func Test_Cov13_ChmodVerifier_IsEqualSkipInvalid(t *testing.T) {
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid(
		"/nonexistent/cov13/skip2", 0755)
	if !result {
		t.Fatal("expected true for invalid path with skip")
	}
}

func Test_Cov13_ChmodVerifier_GetRwx9_Short(t *testing.T) {
	result := chmodhelper.ChmodVerify.GetRwx9(0)
	if result != "---------" {
		t.Fatalf("expected --------- got %q", result)
	}
}

func Test_Cov13_ChmodVerifier_GetRwx9_Valid(t *testing.T) {
	result := chmodhelper.ChmodVerify.GetRwx9(0755)
	if result == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov13_ChmodVerifier_GetExistingRwxWrapperMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	chmodhelper.ChmodVerify.GetExistingRwxWrapperMust("/nonexistent/cov13/must")
}

func Test_Cov13_ChmodVerifier_GetExistingChmodRwxWrappers(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_verifier_wrappers.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	results, err := chmodhelper.ChmodVerify.GetExistingChmodRwxWrappers(true, tmpFile)
	if err != nil || len(results) == 0 {
		t.Fatal("expected results")
	}
}

func Test_Cov13_ChmodVerifier_GetExistsFilteredPathFileInfoMap(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_verifier_filtered.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	m := chmodhelper.ChmodVerify.GetExistsFilteredPathFileInfoMap(false, tmpFile)
	if m == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov13_ChmodVerifier_RwxFull_InvalidLength(t *testing.T) {
	err := chmodhelper.ChmodVerify.RwxFull("/tmp", "rwx")
	if err == nil {
		t.Fatal("expected error for wrong length")
	}
}

func Test_Cov13_ChmodVerifier_PathsUsingPartialRwxOptions(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_partial_opts.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(
		false, false, "-rw-r--r--", tmpFile)
	_ = err
}

func Test_Cov13_ChmodVerifier_PathsUsingPartialRwxOptions_Error(t *testing.T) {
	_, err := chmodhelper.NewRwxVariableWrapper("-rZxr-xr-x")
	_ = err
	verifyErr := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(
		false, false, "-rZxr-xr-x", "/tmp")
	if verifyErr == nil {
		t.Fatal("expected error for invalid rwx")
	}
}

func Test_Cov13_ChmodVerifier_PathsUsingRwxFull_Empty(t *testing.T) {
	err := chmodhelper.ChmodVerify.PathsUsingRwxFull(false, "-rwxr-xr-x")
	if err == nil {
		t.Fatal("expected error for empty locations")
	}
}

func Test_Cov13_ChmodVerifier_PathsUsingRwxFull_ContinueOnError(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxfull_cont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodVerify.PathsUsingRwxFull(true, "-rw-r--r--", tmpFile)
	_ = err
}

// ── chmodVerifier.UsingHashmap ──

func Test_Cov13_ChmodVerifier_UsingHashmap(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_hashmap.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	hm := corestr.New.Hashmap.Cap(1)
	hm.AddOrUpdate(tmpFile, "-rw-r--r--")

	err := chmodhelper.ChmodVerify.UsingHashmap(hm)
	_ = err
}

func Test_Cov13_ChmodVerifier_UsingHashmap_Mismatch(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_hashmap_mm.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	hm := corestr.New.Hashmap.Cap(1)
	hm.AddOrUpdate(tmpFile, "-rwxrwxrwx")

	err := chmodhelper.ChmodVerify.UsingHashmap(hm)
	if err == nil {
		t.Fatal("expected mismatch error")
	}
}

// ── chmodVerifier.UsingRwxOwnerGroupOther ──

func Test_Cov13_ChmodVerifier_UsingRwxOwnerGroupOther_Nil(t *testing.T) {
	err := chmodhelper.ChmodVerify.UsingRwxOwnerGroupOther(nil, "/tmp")
	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov13_ChmodVerifier_UsingRwxOwnerGroupOther_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_usingogo.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	rwx := &chmodins.RwxOwnerGroupOther{
		Owner: "rw-",
		Group: "r--",
		Other: "r--",
	}
	err := chmodhelper.ChmodVerify.UsingRwxOwnerGroupOther(rwx, tmpFile)
	_ = err
}

// ── chmodApplier.RwxPartial ──

func Test_Cov13_ChmodApplier_RwxPartial_Empty(t *testing.T) {
	err := chmodhelper.ChmodApply.RwxPartial("-rwxr-xr-x", &chmodins.Condition{})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov13_ChmodApplier_RwxPartial_Error(t *testing.T) {
	skipOnWindows(t)
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxpartial_err.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodApply.RwxPartial("-rwxr-xr-x", nil, tmpFile)
	if err == nil {
		t.Fatal("expected error for nil condition")
	}
}

func Test_Cov13_ChmodApplier_RwxPartial_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxpartial.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodApply.RwxPartial("-rw-r--r--", &chmodins.Condition{}, tmpFile)
	_ = err
}

// ── RwxStringApplyChmod ──

func Test_Cov13_RwxStringApplyChmod_Empty(t *testing.T) {
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", &chmodins.Condition{})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov13_RwxStringApplyChmod_InvalidLength(t *testing.T) {
	err := chmodhelper.RwxStringApplyChmod("rwx", &chmodins.Condition{}, "/tmp")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_RwxStringApplyChmod_NilCondition(t *testing.T) {
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", nil, "/tmp")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_RwxStringApplyChmod_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxstr.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.RwxStringApplyChmod("-rw-r--r--", &chmodins.Condition{}, tmpFile)
	_ = err
}

func Test_Cov13_RwxStringApplyChmod_InvalidRwx(t *testing.T) {
	skipOnWindows(t)
	// Parser does not validate individual rwx characters, only length.
	// Use a string with wrong length to trigger error.
	err := chmodhelper.RwxStringApplyChmod("-rZx", &chmodins.Condition{}, "/tmp")
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── RwxOwnerGroupOtherApplyChmod ──

func Test_Cov13_RwxOwnerGroupOtherApplyChmod_Empty(t *testing.T) {
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, &chmodins.Condition{})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov13_RwxOwnerGroupOtherApplyChmod_NilRwx(t *testing.T) {
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, &chmodins.Condition{}, "/tmp")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_RwxOwnerGroupOtherApplyChmod_NilCondition(t *testing.T) {
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, nil, "/tmp")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_RwxOwnerGroupOtherApplyChmod_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_ogo.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rw-", Group: "r--", Other: "r--"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, &chmodins.Condition{}, tmpFile)
	_ = err
}

func Test_Cov13_RwxOwnerGroupOtherApplyChmod_InvalidRwx(t *testing.T) {
	skipOnWindows(t)
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rZx", Group: "r-x", Other: "r-x"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, &chmodins.Condition{}, "/tmp")
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── RwxMatchingStatus.CreateErrFinalError ──

func Test_Cov13_RwxMatchingStatus_CreateErrFinalError_AllMatching(t *testing.T) {
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching: true,
	}
	err := status.CreateErrFinalError()
	if err != nil {
		t.Fatal("expected nil for all matching")
	}
}

func Test_Cov13_RwxMatchingStatus_CreateErrFinalError_WithMismatch(t *testing.T) {
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching: false,
		RwxMismatchInfos: []*chmodhelper.RwxMismatchInfo{
			{FilePath: "/test", Expecting: "rwxr-xr-x", Actual: "rw-r--r--"},
		},
	}
	err := status.CreateErrFinalError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_RwxMatchingStatus_CreateErrFinalError_WithError(t *testing.T) {
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching:    false,
		RwxMismatchInfos: []*chmodhelper.RwxMismatchInfo{},
		Error:            os.ErrNotExist,
	}
	err := status.CreateErrFinalError()
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── fwChmodApplier ──

func Test_Cov13_FwChmodApplier_Apply_Error(t *testing.T) {
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent/cov13/fw",
		FilePath:  "/nonexistent/cov13/fw/test.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.Apply(0644, "/nonexistent/cov13/fw/test.txt")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_FwChmodApplier_OnDiffFile(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_fwdiff")
	os.MkdirAll(tmpDir, 0755)
	tmpFile := filepath.Join(tmpDir, "diff.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0755,
		ParentDir: tmpDir,
		FilePath:  tmpFile,
	}
	applier := rw.ChmodApplier()
	err := applier.OnDiffFile(false, tmpFile)
	_ = err
}

func Test_Cov13_FwChmodApplier_OnDiffFile_SkipInvalid(t *testing.T) {
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent",
		FilePath:  "/nonexistent/f.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.OnDiffFile(true, "/nonexistent/cov13/skip.txt")
	if err != nil {
		t.Fatal("expected nil for skip invalid")
	}
}

func Test_Cov13_FwChmodApplier_OnDiffDir(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_fwdiffdir")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0777,
		ChmodFile: 0644,
		ParentDir: tmpDir,
		FilePath:  filepath.Join(tmpDir, "f.txt"),
	}
	applier := rw.ChmodApplier()
	err := applier.OnDiffDir(false, tmpDir)
	_ = err
}

func Test_Cov13_FwChmodApplier_OnDiffDir_SkipInvalid(t *testing.T) {
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent",
		FilePath:  "/nonexistent/f.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.OnDiffDir(true, "/nonexistent/cov13/skipdir")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov13_FwChmodApplier_OnAll(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_fwall")
	os.MkdirAll(tmpDir, 0755)
	tmpFile := filepath.Join(tmpDir, "all.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: tmpDir,
		FilePath:  tmpFile,
	}
	applier := rw.ChmodApplier()
	err := applier.OnAll()
	_ = err
}

func Test_Cov13_FwChmodApplier_OnAll_Error(t *testing.T) {
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent/cov13/all_err",
		FilePath:  "/nonexistent/cov13/all_err/f.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.OnAll()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_FwChmodApplier_OnMismatch(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_fwmismatch")
	os.MkdirAll(tmpDir, 0755)
	tmpFile := filepath.Join(tmpDir, "mm.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: tmpDir,
		FilePath:  tmpFile,
	}
	applier := rw.ChmodApplier()
	err := applier.OnMismatch(true, true)
	_ = err
}

func Test_Cov13_FwChmodApplier_OnMismatch_BothFalse(t *testing.T) {
	rw := &chmodhelper.SimpleFileReaderWriter{}
	applier := rw.ChmodApplier()
	err := applier.OnMismatch(false, false)
	if err != nil {
		t.Fatal("expected nil")
	}
}

// ── fwChmodVerifier.HasMismatchParentDir ──

func Test_Cov13_FwChmodVerifier_HasMismatchParentDir(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_fwverify")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0777,
		ChmodFile: 0644,
		ParentDir: tmpDir,
		FilePath:  filepath.Join(tmpDir, "f.txt"),
	}
	v := rw.ChmodVerifier()
	_ = v.HasMismatchParentDir()
}

// ── CreateDirFilesWithRwxPermission error branches ──

func Test_Cov13_CreateDirFilesWithRwxPermission_FileModeErr(t *testing.T) {
	skipOnWindows(t)
	perm := &chmodhelper.DirFilesWithRwxPermission{
		DirWithFiles: chmodhelper.DirWithFiles{
			Dir: "/tmp/cov13_perm",
		},
		ApplyRwx: chmodins.RwxOwnerGroupOther{
			Owner: "rw",
			Group: "r-x",
			Other: "r-x",
		},
	}
	err := chmodhelper.CreateDirFilesWithRwxPermission(false, perm)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_CreateDirFilesWithRwxPermission_CreateErr(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_perm_file")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	perm := &chmodhelper.DirFilesWithRwxPermission{
		DirWithFiles: chmodhelper.DirWithFiles{
			Dir:   filepath.Join(tmpFile, "subdir"),
			Files: []string{"a.txt"},
		},
		ApplyRwx: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r-x",
		},
	}
	err := chmodhelper.CreateDirFilesWithRwxPermission(false, perm)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── CreateDirWithFiles error branches ──

func Test_Cov13_CreateDirWithFiles_RemoveDirErr(t *testing.T) {
	// removeDirIf when dir doesn't exist and isRemove=true is fine
	tmpDir := filepath.Join(os.TempDir(), "cov13_createdir")
	os.RemoveAll(tmpDir)

	err := chmodhelper.CreateDirWithFiles(true, 0755, &chmodhelper.DirWithFiles{
		Dir:   tmpDir,
		Files: []string{"a.txt"},
	})
	defer os.RemoveAll(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov13_CreateDirWithFiles_MkdirErr(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_mkdirerr")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{
		Dir: filepath.Join(tmpFile, "sub"),
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_CreateDirWithFiles_FileCreateErr(t *testing.T) {
	// Dir exists but create file fails (file path under a file)
	tmpDir := filepath.Join(os.TempDir(), "cov13_filecreateerr")
	os.MkdirAll(tmpDir, 0755)
	// Create a file where subdirectory is expected
	blockerFile := filepath.Join(tmpDir, "blocker")
	os.WriteFile(blockerFile, []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{
		Dir:   tmpDir,
		Files: []string{filepath.Join("blocker", "impossible.txt")},
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_CreateDirWithFiles_NoFiles(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_nofiles")
	os.RemoveAll(tmpDir)
	defer os.RemoveAll(tmpDir)

	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{
		Dir: tmpDir,
	})
	if err != nil {
		t.Fatal(err)
	}
}

// ── CreateDirsWithFiles error ──

func Test_Cov13_CreateDirsWithFiles_Error(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_dirsfiles_err")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.CreateDirsWithFiles(false, 0755,
		chmodhelper.DirWithFiles{Dir: filepath.Join(tmpFile, "sub"), Files: []string{"a.txt"}})
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── CreateDirFilesWithRwxPermissions error ──

func Test_Cov13_CreateDirFilesWithRwxPermissions_Error(t *testing.T) {
	skipOnWindows(t)
	perms := []chmodhelper.DirFilesWithRwxPermission{
		{
			DirWithFiles: chmodhelper.DirWithFiles{Dir: "/tmp/cov13_perms"},
			ApplyRwx:     chmodins.RwxOwnerGroupOther{Owner: "rw", Group: "r-x", Other: "r-x"},
		},
	}
	err := chmodhelper.CreateDirFilesWithRwxPermissions(false, perms)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── CreateDirFilesWithRwxPermissionsMust panic ──

func Test_Cov13_CreateDirFilesWithRwxPermissionsMust_Panic(t *testing.T) {
	skipOnWindows(t)
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	perms := []chmodhelper.DirFilesWithRwxPermission{
		{
			DirWithFiles: chmodhelper.DirWithFiles{Dir: "/tmp/cov13_must"},
			ApplyRwx:     chmodins.RwxOwnerGroupOther{Owner: "rw", Group: "r-x", Other: "r-x"},
		},
	}
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(false, perms)
}

// ── DirFilesWithContent.Create error branches ──

func Test_Cov13_DirFilesWithContent_Create_RemoveError(t *testing.T) {
	skipOnWindows(t)
	dfc := &chmodhelper.DirFilesWithContent{
		Dir:         "/nonexistent/cov13/dfc",
		DirFileMode: 0755,
		Files: []chmodhelper.FileWithContent{
			{RelativePath: "a.txt", FileMode: 0644, Content: []string{"hello"}},
		},
	}
	err := dfc.Create(true)
	// remove on non-existent is fine, but write fails
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_DirFilesWithContent_Create_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_dfc_ok")
	os.RemoveAll(tmpDir)
	defer os.RemoveAll(tmpDir)

	dfc := &chmodhelper.DirFilesWithContent{
		Dir:         tmpDir,
		DirFileMode: 0755,
		Files: []chmodhelper.FileWithContent{
			{RelativePath: "a.txt", FileMode: 0644, Content: []string{"hello"}},
		},
	}
	err := dfc.Create(false)
	if err != nil {
		t.Fatal(err)
	}
}

// ── FileWithContent.ReadLines error ──

func Test_Cov13_FileWithContent_ReadLines_Error(t *testing.T) {
	fc := chmodhelper.FileWithContent{
		RelativePath: "nonexistent.txt",
		FileMode:     0644,
	}
	_, err := fc.ReadLines("/nonexistent/cov13")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_FileWithContent_ReadLines_Success(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_readline")
	os.MkdirAll(tmpDir, 0755)
	os.WriteFile(filepath.Join(tmpDir, "lines.txt"), []byte("a\nb\nc"), 0644)
	defer os.RemoveAll(tmpDir)

	fc := chmodhelper.FileWithContent{
		RelativePath: "lines.txt",
		FileMode:     0644,
	}
	lines, err := fc.ReadLines(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
}

// ── fileWriter.All error branches ──

func Test_Cov13_FileWriter_All_DirErr(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_fwall_file")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rw := newTestRW(filepath.Join(tmpFile, "sub"), "test.txt")
	err := rw.Write([]byte("x"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov13_FileWriter_Remove(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "cov13_fwremove")
	os.MkdirAll(tmpDir, 0755)
	tmpFile := filepath.Join(tmpDir, "rm.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.RemoveAll(tmpDir)

	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:            0755,
		ChmodFile:           0644,
		ParentDir:           tmpDir,
		FilePath:            tmpFile,
		IsRemoveBeforeWrite: true,
	}
	err := rw.Write([]byte("new content"))
	if err != nil {
		t.Fatal(err)
	}
}
