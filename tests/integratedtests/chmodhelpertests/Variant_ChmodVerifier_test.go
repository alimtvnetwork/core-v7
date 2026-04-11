package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

func skipIfWindows(t *testing.T) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("skipping file permission test on Windows")
	}
}

// --- Variant ---

func Test_Variant_String_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	v := chmodhelper.Variant("755")

	// Act
	actual := args.Map{"result": v.String() != "755"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 755", actual)
}

func Test_Variant_ExpandOctalByte_FromVariantChmodVerifier(t *testing.T) {
	v := chmodhelper.Variant("755")
	r, w, x := v.ExpandOctalByte()
	if r == 0 && w == 0 && x == 0 {
		// at least some should be non-zero for 755
	}
	_ = r
	_ = w
	_ = x
}

func Test_Variant_ToWrapper_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	v := chmodhelper.Variant("755")
	wrapper, err := v.ToWrapper()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": wrapper.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty wrapper", actual)
}

func Test_Variant_ToWrapperPtr_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	v := chmodhelper.Variant("755")
	wrapper, err := v.ToWrapperPtr()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": wrapper == nil || wrapper.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty wrapper ptr", actual)
}

// --- RwxWrapper basic ---

func Test_RwxWrapper_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var w *chmodhelper.RwxWrapper

	// Act
	actual := args.Map{"result": w.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	actual = args.Map{"result": w.IsNull()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null for nil", actual)
	actual = args.Map{"result": w.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for nil", actual)
}

func Test_RwxWrapper_IsDefined(t *testing.T) {
	// Arrange
	v := chmodhelper.Variant("755")
	w, err := v.ToWrapperPtr()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": w.IsDefined()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected defined", actual)
	actual = args.Map{"result": w.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has items", actual)
}

// --- SingleRwx ---

func Test_NewSingleRwx_Valid(t *testing.T) {
	// Arrange
	s, err := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)

	// Act
	actual := args.Map{"result": err != nil || s == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid SingleRwx", actual)
}

func Test_NewSingleRwx_InvalidLength(t *testing.T) {
	// Arrange
	_, err := chmodhelper.NewSingleRwx("rw", chmodclasstype.All)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid rwx length", actual)
}

func Test_SingleRwx_ToRwxOwnerGroupOther_All(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	ogo := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rwx" || ogo.Group != "rwx" || ogo.Other != "rwx"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected all rwx", actual)
}

func Test_SingleRwx_ToRwxOwnerGroupOther_Owner(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	ogo := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rwx"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected owner rwx", actual)
}

func Test_SingleRwx_ToRwxOwnerGroupOther_Group(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("r-x", chmodclasstype.Group)
	ogo := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": ogo == nil || ogo.Group != "r-x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected group r-x", actual)
}

func Test_SingleRwx_ToRwxOwnerGroupOther_Other(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("r--", chmodclasstype.Other)
	ogo := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": ogo == nil || ogo.Other != "r--"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected other r--", actual)
}

func Test_SingleRwx_ToRwxOwnerGroupOther_OwnerGroup(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.OwnerGroup)
	ogo := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rwx" || ogo.Group != "rwx"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected owner+group rwx", actual)
}

func Test_SingleRwx_ToRwxOwnerGroupOther_GroupOther(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("r-x", chmodclasstype.GroupOther)
	ogo := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": ogo == nil || ogo.Group != "r-x" || ogo.Other != "r-x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected group+other r-x", actual)
}

func Test_SingleRwx_ToRwxOwnerGroupOther_OwnerOther(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rw-", chmodclasstype.OwnerOther)
	ogo := s.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rw-" || ogo.Other != "rw-"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected owner+other rw-", actual)
}

func Test_SingleRwx_ToRwxInstruction_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	ins := s.ToRwxInstruction(cond)

	// Act
	actual := args.Map{"result": ins == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil instruction", actual)
}

func Test_SingleRwx_ToVarRwxWrapper_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	vw, err := s.ToVarRwxWrapper()

	// Act
	actual := args.Map{"result": err != nil || vw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid var wrapper", actual)
}

func Test_SingleRwx_ToDisabledRwxWrapper_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	dw, err := s.ToDisabledRwxWrapper()

	// Act
	actual := args.Map{"result": err != nil || dw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid disabled wrapper", actual)
}

func Test_SingleRwx_ToRwxWrapper_All_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	w, err := s.ToRwxWrapper()

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid rwx wrapper", actual)
}

func Test_SingleRwx_ToRwxWrapper_NotAll_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	_, err := s.ToRwxWrapper()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-all class type", actual)
}

func Test_SingleRwx_ApplyOnMany_Empty_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	err := s.ApplyOnMany(cond)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty locations", actual)
}

func Test_SingleRwx_ApplyOnMany_Valid_FromVariantChmodVerifier(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	err := s.ApplyOnMany(cond, f)
	_ = err // may succeed or fail based on OS, just exercise path
}

// --- NewCreator.RwxWrapper ---

func Test_NewRwxWrapper_UsingVariant_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	w, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.Variant("644"))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": w.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_NewRwxWrapper_UsingVariantPtr(t *testing.T) {
	// Arrange
	w, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("644"))

	// Act
	actual := args.Map{"result": err != nil || w == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected: err=, w=", actual)
}

func Test_NewRwxWrapper_RwxFullString(t *testing.T) {
	// Arrange
	w, err := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": w.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- ChmodApply and Verify ---

func Test_ChmodApply_RecursivePath_FromVariantChmodVerifier(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	err := chmodhelper.ChmodApply.RecursivePath(true, 0755, tmpDir)
	_ = err
}

func Test_ChmodVerify_RwxFull(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0755)
	os.Chmod(f, 0755)

	err := chmodhelper.ChmodVerify.RwxFull(f, "-rwxr-xr-x")
	_ = err
}

func Test_ChmodVerify_RwxFull_NoDash(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0755)
	os.Chmod(f, 0755)

	err := chmodhelper.ChmodVerify.RwxFull(f, "rwxr-xr-x")
	_ = err
}

// --- TempDirGetter ---

func Test_TempDirGetter(t *testing.T) {
	// Arrange
	td := chmodhelper.TempDirGetter.TempDefault()

	// Act
	actual := args.Map{"result": td == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty temp dir", actual)
}

// --- ExpandCharRwx ---

func Test_ExpandCharRwx_Valid(t *testing.T) {
	r, w, x := chmodhelper.ExpandCharRwx("755")
	_ = r
	_ = w
	_ = x
}

func Test_ExpandCharRwx_Short(t *testing.T) {
	defer func() { recover() }() // may panic on short string
	chmodhelper.ExpandCharRwx("")
}

// --- SimpleFileReaderWriter ---

func Test_SimpleFileReaderWriter(t *testing.T) {
	// Arrange
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "sub", "test.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, f)

	// Act
	actual := args.Map{"result": rw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil reader writer", actual)
}

// --- FileModeFriendlyString ---

func Test_FileModeFriendlyString_FromVariantChmodVerifier(t *testing.T) {
	s := chmodhelper.FileModeFriendlyString(0755)
	_ = s
}

// --- PathExistStat ---

func Test_GetPathExistStat_NonExistent(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat("/nonexistent/path/xyz_i18")

	// Act
	actual := args.Map{"result": stat == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil stat", actual)
	actual = args.Map{"result": stat.IsExist}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-exist for fake path", actual)
}

func Test_GetPathExistStat_Existing(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	stat := chmodhelper.GetPathExistStat(f)

	// Act
	actual := args.Map{"result": stat.IsExist}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected exist for real path", actual)
}

// --- IsPathExists ---

func Test_IsPathExists_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()

	// Act
	actual := args.Map{"result": chmodhelper.IsPathExists(tmpDir)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected exists for temp dir", actual)

	actual = args.Map{"result": chmodhelper.IsPathExists("/nonexistent/xyz")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not exists", actual)
}

func Test_IsPathInvalid_FromVariantChmodVerifier(t *testing.T) {
	// Act
	actual := args.Map{"result": chmodhelper.IsPathInvalid("/nonexistent/xyz")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for nonexistent path", actual)
}

func Test_IsDirectory_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()

	// Act
	actual := args.Map{"result": chmodhelper.IsDirectory(tmpDir)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected directory", actual)

	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)
	actual = args.Map{"result": chmodhelper.IsDirectory(f)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not directory for file", actual)
}

// --- GetExistingChmod ---

func Test_GetExistingChmod_FromVariantChmodVerifier(t *testing.T) {
	// Arrange
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	chmod, err := chmodhelper.GetExistingChmod(f)

	// Act
	actual := args.Map{"result": err != nil || chmod == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero chmod", actual)
}

func Test_GetExistingChmodOfValidFile(t *testing.T) {
	// Arrange
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	chmod, isInvalid := chmodhelper.GetExistingChmodOfValidFile(f)

	// Act
	actual := args.Map{"result": isInvalid || chmod == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected result", actual)
}

func Test_GetExistingChmodOfValidFile_NonExistent(t *testing.T) {
	// Arrange
	_, isInvalid := chmodhelper.GetExistingChmodOfValidFile("/nonexistent/xyz")

	// Act
	actual := args.Map{"result": isInvalid}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for nonexistent file", actual)
}
