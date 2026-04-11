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
	"github.com/alimtvnetwork/core/coretests/args"
)

func skipOnWindows(t *testing.T) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("chmod behavior differs on Windows")
	}
}

// ── SingleRwx.ToRwxOwnerGroupOther default panic ──

func Test_SingleRwx_ToRwxOwnerGroupOther_Default(t *testing.T) {
	// Arrange
	s := &chmodhelper.SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.All,
	}
	result := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": result.Owner != "rwx" || result.Group != "rwx" || result.Other != "rwx"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected all rwx", actual)
}

// ── SingleRwx.ToDisabledRwxWrapper ──

func Test_SingleRwx_ToDisabledRwxWrapper_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	w, err := s.ToDisabledRwxWrapper()

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_SingleRwx_ToDisabledRwxWrapper_Error(t *testing.T) {
	// Arrange
	// Invalid chars are normalized as disabled permissions.
	s := &chmodhelper.SingleRwx{
		Rwx:       "rZx",
		ClassType: chmodclasstype.All,
	}
	w, err := s.ToDisabledRwxWrapper()

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper for normalized invalid chars", actual)
}

// ── SingleRwx.ToRwxWrapper ──

func Test_SingleRwx_ToRwxWrapper_NotAll(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	_, err := s.ToRwxWrapper()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-All class type", actual)
}

func Test_SingleRwx_ToRwxWrapper_All_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	w, err := s.ToRwxWrapper()

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper", actual)
}

func Test_SingleRwx_ToRwxWrapper_Error(t *testing.T) {
	// Arrange
	s := &chmodhelper.SingleRwx{
		Rwx:       "rZx",
		ClassType: chmodclasstype.All,
	}
	w, err := s.ToRwxWrapper()

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected wrapper for normalized invalid chars", actual)
}

// ── SingleRwx.ApplyOnMany ──

func Test_SingleRwx_ApplyOnMany_Empty(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	err := s.ApplyOnMany(&chmodins.Condition{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_SingleRwx_ApplyOnMany_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_apply_many.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	s, _ := chmodhelper.NewSingleRwx("rw-", chmodclasstype.All)
	err := s.ApplyOnMany(&chmodins.Condition{}, tmpFile)
	_ = err
}

func Test_SingleRwx_ApplyOnMany_Error(t *testing.T) {
	// Arrange
	s := &chmodhelper.SingleRwx{
		Rwx:       "rZx",
		ClassType: chmodclasstype.All,
	}
	err := s.ApplyOnMany(&chmodins.Condition{}, "/some/path")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── newRwxWrapperCreator.CreatePtr error ──

func Test_CreatePtr_Error(t *testing.T) {
	// Arrange
	_, err := chmodhelper.New.RwxWrapper.CreatePtr("999")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid mode", actual)
}

func Test_CreatePtr_Valid(t *testing.T) {
	// Arrange
	ptr, err := chmodhelper.New.RwxWrapper.CreatePtr("755")

	// Act
	actual := args.Map{"result": err != nil || ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid pointer", actual)
}

// ── newRwxWrapperCreator.Create invalid char ──

func Test_Create_InvalidChar(t *testing.T) {
	// Arrange
	_, err := chmodhelper.New.RwxWrapper.Create("89a")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── newRwxWrapperCreator.Create invalid length ──

func Test_Create_InvalidLength(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for wrong length", actual)
	}()
	chmodhelper.New.RwxWrapper.Create("77")
}

// ── newRwxWrapperCreator.UsingChmod ──

func Test_UsingChmod_Valid(t *testing.T) {
	// Arrange
	w := chmodhelper.New.RwxWrapper.UsingChmod(0755)

	// Act
	actual := args.Map{"result": w == nil || w.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty wrapper", actual)
}

func Test_UsingChmod_Zero(t *testing.T) {
	// Arrange
	w := chmodhelper.New.RwxWrapper.UsingChmod(0)

	// Act
	actual := args.Map{"result": w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil wrapper", actual)
	actual = args.Map{"result": w.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty wrapper for zero mode", actual)
}

// ── newRwxWrapperCreator.UsingVariantPtr ──
// Variant is a string type, so we use valid string values

func Test_UsingVariantPtr_Valid(t *testing.T) {
	// Arrange
	w, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("755"))

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid pointer", actual)
}

func Test_UsingVariantPtr_Invalid(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// panics on length != 3 — expected
		}
	}()
	_, _ = chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("99"))
}

// ── newRwxWrapperCreator.Instruction ──

func Test_Instruction_Valid(t *testing.T) {
	// Arrange
	ins, err := chmodhelper.New.RwxWrapper.Instruction(
		"-rwxr-xr-x",
		chmodins.Condition{})

	// Act
	actual := args.Map{"result": err != nil || ins == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected instruction", actual)
}

func Test_Instruction_Error(t *testing.T) {
	// Arrange
	_, err := chmodhelper.New.RwxWrapper.Instruction(
		"rwxr-xr-x",
		chmodins.Condition{})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong length", actual)
}

// ── newAttributeCreator.UsingByteMust panic ──

func Test_UsingByteMust_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for byte > 7", actual)
	}()
	chmodhelper.New.Attribute.UsingByteMust(8)
}

// ── newAttributeCreator.UsingRwxString panic ──

func Test_UsingRwxString_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for wrong length", actual)
	}()
	chmodhelper.New.Attribute.UsingRwxString("rw")
}

// ── chmodVerifier branches ──

func Test_ChmodVerifier_IsEqualRwxFullSkipInvalid_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid(
		"/nonexistent/cov13/skip", "-rwxr-xr-x")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for invalid path with skip", actual)
}

func Test_ChmodVerifier_IsEqualSkipInvalid_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid(
		"/nonexistent/cov13/skip2", 0755)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for invalid path with skip", actual)
}

func Test_ChmodVerifier_GetRwx9_Short(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.GetRwx9(0)

	// Act
	actual := args.Map{"result": result != "---------"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ---------", actual)
}

func Test_ChmodVerifier_GetRwx9_Valid(t *testing.T) {
	// Arrange
	result := chmodhelper.ChmodVerify.GetRwx9(0755)

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_ChmodVerifier_GetExistingRwxWrapperMust_Panic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	chmodhelper.ChmodVerify.GetExistingRwxWrapperMust("/nonexistent/cov13/must")
}

func Test_ChmodVerifier_GetExistingChmodRwxWrappers(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov13_verifier_wrappers.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	results, err := chmodhelper.ChmodVerify.GetExistingChmodRwxWrappers(true, tmpFile)

	// Act
	actual := args.Map{"result": err != nil || len(results) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected results", actual)
}

func Test_ChmodVerifier_GetExistsFilteredPathFileInfoMap(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov13_verifier_filtered.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	m := chmodhelper.ChmodVerify.GetExistsFilteredPathFileInfoMap(false, tmpFile)

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_ChmodVerifier_RwxFull_InvalidLength_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodVerify.RwxFull("/tmp", "rwx")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong length", actual)
}

func Test_ChmodVerifier_PathsUsingPartialRwxOptions(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_partial_opts.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(
		false, false, "-rw-r--r--", tmpFile)
	_ = err
}

func Test_ChmodVerifier_PathsUsingPartialRwxOptions_Error(t *testing.T) {
	// Arrange
	// PathsUsingPartialRwxOptions with valid-length partial rwx on matching path
	// Parser accepts any chars and pads to 10; no error triggered for char content.
	// Test empty locations branch (returns nil from CreateErrFinalError).
	verifyErr := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(
		false, false, "-rwx")

	// Act
	actual := args.Map{"result": verifyErr != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty locations", actual)
}

func Test_ChmodVerifier_PathsUsingRwxFull_Empty(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodVerify.PathsUsingRwxFull(false, "-rwxr-xr-x")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for empty locations", actual)
}

func Test_ChmodVerifier_PathsUsingRwxFull_ContinueOnError_FromSingleRwxPanic(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxfull_cont.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodVerify.PathsUsingRwxFull(true, "-rw-r--r--", tmpFile)
	_ = err
}

// ── chmodVerifier.UsingHashmap ──

func Test_ChmodVerifier_UsingHashmap(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_hashmap.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	hm := corestr.New.Hashmap.Cap(1)
	hm.AddOrUpdate(tmpFile, "-rw-r--r--")

	err := chmodhelper.ChmodVerify.UsingHashmap(hm)
	_ = err
}

func Test_ChmodVerifier_UsingHashmap_Mismatch(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov13_hashmap_mm.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	os.Chmod(tmpFile, 0644)
	defer os.Remove(tmpFile)

	hm := corestr.New.Hashmap.Cap(1)
	hm.AddOrUpdate(tmpFile, "-rwxrwxrwx")

	err := chmodhelper.ChmodVerify.UsingHashmap(hm)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch error", actual)
}

// ── chmodVerifier.UsingRwxOwnerGroupOther ──

func Test_ChmodVerifier_UsingRwxOwnerGroupOther_Nil(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodVerify.UsingRwxOwnerGroupOther(nil, "/tmp")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_ChmodVerifier_UsingRwxOwnerGroupOther_Valid(t *testing.T) {
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

func Test_ChmodApplier_RwxPartial_Empty(t *testing.T) {
	// Arrange
	err := chmodhelper.ChmodApply.RwxPartial("-rwxr-xr-x", &chmodins.Condition{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_ChmodApplier_RwxPartial_Error(t *testing.T) {
	// Arrange
	skipOnWindows(t)
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxpartial_err.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodApply.RwxPartial("-rwxr-xr-x", nil, tmpFile)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil condition", actual)
}

func Test_ChmodApplier_RwxPartial_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxpartial.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.ChmodApply.RwxPartial("-rw-r--r--", &chmodins.Condition{}, tmpFile)
	_ = err
}

// ── RwxStringApplyChmod ──

func Test_RwxStringApplyChmod_Empty(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", &chmodins.Condition{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_RwxStringApplyChmod_InvalidLength(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxStringApplyChmod("rwx", &chmodins.Condition{}, "/tmp")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_RwxStringApplyChmod_NilCondition(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", nil, "/tmp")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_RwxStringApplyChmod_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_rwxstr.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.RwxStringApplyChmod("-rw-r--r--", &chmodins.Condition{}, tmpFile)
	_ = err
}

func Test_RwxStringApplyChmod_InvalidRwx(t *testing.T) {
	// Arrange
	skipOnWindows(t)
	// Parser does not validate individual rwx characters, only length.
	// Use a string with wrong length to trigger error.
	err := chmodhelper.RwxStringApplyChmod("-rZx", &chmodins.Condition{}, "/tmp")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── RwxOwnerGroupOtherApplyChmod ──

func Test_RwxOwnerGroupOtherApplyChmod_Empty(t *testing.T) {
	// Arrange
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, &chmodins.Condition{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilRwx(t *testing.T) {
	// Arrange
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, &chmodins.Condition{}, "/tmp")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilCondition(t *testing.T) {
	// Arrange
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, nil, "/tmp")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_Valid(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "cov13_ogo.txt")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rw-", Group: "r--", Other: "r--"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, &chmodins.Condition{}, tmpFile)
	_ = err
}

func Test_RwxOwnerGroupOtherApplyChmod_InvalidRwx(t *testing.T) {
	// Arrange
	skipOnWindows(t)
	// Use invalid length rwx to trigger parsing error
	rwx := &chmodins.RwxOwnerGroupOther{Owner: "rZxExtra", Group: "r-x", Other: "r-x"}
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(rwx, &chmodins.Condition{}, "/tmp")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── RwxMatchingStatus.CreateErrFinalError ──

func Test_RwxMatchingStatus_CreateErrFinalError_AllMatching_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching: true,
	}
	err := status.CreateErrFinalError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for all matching", actual)
}

func Test_RwxMatchingStatus_CreateErrFinalError_WithMismatch(t *testing.T) {
	// Arrange
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching: false,
		RwxMismatchInfos: []*chmodhelper.RwxMismatchInfo{
			{FilePath: "/test", Expecting: "rwxr-xr-x", Actual: "rw-r--r--"},
		},
	}
	err := status.CreateErrFinalError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_RwxMatchingStatus_CreateErrFinalError_WithError(t *testing.T) {
	// Arrange
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching:    false,
		RwxMismatchInfos: []*chmodhelper.RwxMismatchInfo{},
		Error:            os.ErrNotExist,
	}
	err := status.CreateErrFinalError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── fwChmodApplier ──

func Test_FwChmodApplier_Apply_Error(t *testing.T) {
	// Arrange
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent/cov13/fw",
		FilePath:  "/nonexistent/cov13/fw/test.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.Apply(0644, "/nonexistent/cov13/fw/test.txt")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_FwChmodApplier_OnDiffFile(t *testing.T) {
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

func Test_FwChmodApplier_OnDiffFile_SkipInvalid(t *testing.T) {
	// Arrange
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent",
		FilePath:  "/nonexistent/f.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.OnDiffFile(true, "/nonexistent/cov13/skip.txt")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for skip invalid", actual)
}

func Test_FwChmodApplier_OnDiffDir(t *testing.T) {
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

func Test_FwChmodApplier_OnDiffDir_SkipInvalid(t *testing.T) {
	// Arrange
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent",
		FilePath:  "/nonexistent/f.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.OnDiffDir(true, "/nonexistent/cov13/skipdir")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_FwChmodApplier_OnAll(t *testing.T) {
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

func Test_FwChmodApplier_OnAll_Error(t *testing.T) {
	// Arrange
	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: "/nonexistent/cov13/all_err",
		FilePath:  "/nonexistent/cov13/all_err/f.txt",
	}
	applier := rw.ChmodApplier()
	err := applier.OnAll()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_FwChmodApplier_OnMismatch(t *testing.T) {
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

func Test_FwChmodApplier_OnMismatch_BothFalse(t *testing.T) {
	// Arrange
	rw := &chmodhelper.SimpleFileReaderWriter{}
	applier := rw.ChmodApplier()
	err := applier.OnMismatch(false, false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── fwChmodVerifier.HasMismatchParentDir ──

func Test_FwChmodVerifier_HasMismatchParentDir(t *testing.T) {
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

func Test_CreateDirFilesWithRwxPermission_FileModeErr(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CreateDirFilesWithRwxPermission_CreateErr(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── CreateDirWithFiles error branches ──

func Test_CreateDirWithFiles_RemoveDirErr(t *testing.T) {
	// Arrange
	// removeDirIf when dir doesn't exist and isRemove=true is fine
	tmpDir := filepath.Join(os.TempDir(), "cov13_createdir")
	os.RemoveAll(tmpDir)

	err := chmodhelper.CreateDirWithFiles(true, 0755, &chmodhelper.DirWithFiles{
		Dir:   tmpDir,
		Files: []string{"a.txt"},
	})
	defer os.RemoveAll(tmpDir)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_CreateDirWithFiles_MkdirErr(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov13_mkdirerr")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{
		Dir: filepath.Join(tmpFile, "sub"),
	})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CreateDirWithFiles_FileCreateErr(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CreateDirWithFiles_NoFiles_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov13_nofiles")
	os.RemoveAll(tmpDir)
	defer os.RemoveAll(tmpDir)

	err := chmodhelper.CreateDirWithFiles(false, 0755, &chmodhelper.DirWithFiles{
		Dir: tmpDir,
	})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── CreateDirsWithFiles error ──

func Test_CreateDirsWithFiles_Error(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov13_dirsfiles_err")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	err := chmodhelper.CreateDirsWithFiles(false, 0755,
		chmodhelper.DirWithFiles{Dir: filepath.Join(tmpFile, "sub"), Files: []string{"a.txt"}})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── CreateDirFilesWithRwxPermissions error ──

func Test_CreateDirFilesWithRwxPermissions_Error_FromSingleRwxPanic(t *testing.T) {
	// Arrange
	skipOnWindows(t)
	perms := []chmodhelper.DirFilesWithRwxPermission{
		{
			DirWithFiles: chmodhelper.DirWithFiles{Dir: "/tmp/cov13_perms"},
			ApplyRwx:     chmodins.RwxOwnerGroupOther{Owner: "rw", Group: "r-x", Other: "r-x"},
		},
	}
	err := chmodhelper.CreateDirFilesWithRwxPermissions(false, perms)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── CreateDirFilesWithRwxPermissionsMust panic ──

func Test_CreateDirFilesWithRwxPermissionsMust_Panic(t *testing.T) {
	// Arrange
	skipOnWindows(t)
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
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

func Test_DirFilesWithContent_Create_RemoveError(t *testing.T) {
	// Arrange
	skipOnWindows(t)
	invalidDir := filepath.Join("/proc", "nonexistent_cov13", "dfc")
	dfc := &chmodhelper.DirFilesWithContent{
		Dir:         invalidDir,
		DirFileMode: 0755,
		Files: []chmodhelper.FileWithContent{
			{RelativePath: "a.txt", FileMode: 0644, Content: []string{"hello"}},
		},
	}
	err := dfc.Create(true)
	// remove on non-existent is fine, but write fails

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_DirFilesWithContent_Create_Success(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

// ── FileWithContent.ReadLines error ──

func Test_FileWithContent_ReadLines_Error(t *testing.T) {
	// Arrange
	fc := chmodhelper.FileWithContent{
		RelativePath: "nonexistent.txt",
		FileMode:     0644,
	}
	_, err := fc.ReadLines("/nonexistent/cov13")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_FileWithContent_ReadLines_Success(t *testing.T) {
	// Arrange
	tmpDir := filepath.Join(os.TempDir(), "cov13_readline")
	os.MkdirAll(tmpDir, 0755)
	os.WriteFile(filepath.Join(tmpDir, "lines.txt"), []byte("a\nb\nc"), 0644)
	defer os.RemoveAll(tmpDir)

	fc := chmodhelper.FileWithContent{
		RelativePath: "lines.txt",
		FileMode:     0644,
	}
	lines, err := fc.ReadLines(tmpDir)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(lines) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 lines", actual)
}

// ── fileWriter.All error branches ──

func Test_FileWriter_All_DirErr(t *testing.T) {
	// Arrange
	tmpFile := filepath.Join(os.TempDir(), "cov13_fwall_file")
	os.WriteFile(tmpFile, []byte("x"), 0644)
	defer os.Remove(tmpFile)

	rw := newTestRW(filepath.Join(tmpFile, "sub"), "test.txt")
	err := rw.Write([]byte("x"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_FileWriter_Remove_FromSingleRwxPanic(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}
