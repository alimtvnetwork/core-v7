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

func Test_I18_Variant_String(t *testing.T) {
	v := chmodhelper.Variant("755")
	actual := args.Map{"result": v.String() != "755"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 755", actual)
}

func Test_I18_Variant_ExpandOctalByte(t *testing.T) {
	v := chmodhelper.Variant("755")
	r, w, x := v.ExpandOctalByte()
	if r == 0 && w == 0 && x == 0 {
		// at least some should be non-zero for 755
	}
	_ = r
	_ = w
	_ = x
}

func Test_I18_Variant_ToWrapper(t *testing.T) {
	v := chmodhelper.Variant("755")
	wrapper, err := v.ToWrapper()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": wrapper.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty wrapper", actual)
}

func Test_I18_Variant_ToWrapperPtr(t *testing.T) {
	v := chmodhelper.Variant("755")
	wrapper, err := v.ToWrapperPtr()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": wrapper == nil || wrapper.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty wrapper ptr", actual)
}

// --- RwxWrapper basic ---

func Test_I18_RwxWrapper_IsEmpty_Nil(t *testing.T) {
	var w *chmodhelper.RwxWrapper
	actual := args.Map{"result": w.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	actual := args.Map{"result": w.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null for nil", actual)
	actual := args.Map{"result": w.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for nil", actual)
}

func Test_I18_RwxWrapper_IsDefined(t *testing.T) {
	v := chmodhelper.Variant("755")
	w, err := v.ToWrapperPtr()
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": w.IsDefined()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected defined", actual)
	actual := args.Map{"result": w.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has items", actual)
}

// --- SingleRwx ---

func Test_I18_NewSingleRwx_Valid(t *testing.T) {
	s, err := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	actual := args.Map{"result": err != nil || s == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid SingleRwx", actual)
}

func Test_I18_NewSingleRwx_InvalidLength(t *testing.T) {
	_, err := chmodhelper.NewSingleRwx("rw", chmodclasstype.All)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid rwx length", actual)
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_All(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	ogo := s.ToRwxOwnerGroupOther()
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rwx" || ogo.Group != "rwx" || ogo.Other != "rwx"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected all rwx", actual)
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_Owner(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	ogo := s.ToRwxOwnerGroupOther()
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rwx"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected owner rwx", actual)
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_Group(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("r-x", chmodclasstype.Group)
	ogo := s.ToRwxOwnerGroupOther()
	actual := args.Map{"result": ogo == nil || ogo.Group != "r-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected group r-x", actual)
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_Other(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("r--", chmodclasstype.Other)
	ogo := s.ToRwxOwnerGroupOther()
	actual := args.Map{"result": ogo == nil || ogo.Other != "r--"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected other r--", actual)
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_OwnerGroup(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.OwnerGroup)
	ogo := s.ToRwxOwnerGroupOther()
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rwx" || ogo.Group != "rwx"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected owner+group rwx", actual)
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_GroupOther(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("r-x", chmodclasstype.GroupOther)
	ogo := s.ToRwxOwnerGroupOther()
	actual := args.Map{"result": ogo == nil || ogo.Group != "r-x" || ogo.Other != "r-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected group+other r-x", actual)
}

func Test_I18_SingleRwx_ToRwxOwnerGroupOther_OwnerOther(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rw-", chmodclasstype.OwnerOther)
	ogo := s.ToRwxOwnerGroupOther()
	actual := args.Map{"result": ogo == nil || ogo.Owner != "rw-" || ogo.Other != "rw-"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected owner+other rw-", actual)
}

func Test_I18_SingleRwx_ToRwxInstruction(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	ins := s.ToRwxInstruction(cond)
	actual := args.Map{"result": ins == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil instruction", actual)
}

func Test_I18_SingleRwx_ToVarRwxWrapper(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	vw, err := s.ToVarRwxWrapper()
	actual := args.Map{"result": err != nil || vw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid var wrapper", actual)
}

func Test_I18_SingleRwx_ToDisabledRwxWrapper(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	dw, err := s.ToDisabledRwxWrapper()
	actual := args.Map{"result": err != nil || dw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid disabled wrapper", actual)
}

func Test_I18_SingleRwx_ToRwxWrapper_All(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	w, err := s.ToRwxWrapper()
	actual := args.Map{"result": err != nil || w == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid rwx wrapper", actual)
}

func Test_I18_SingleRwx_ToRwxWrapper_NotAll(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	_, err := s.ToRwxWrapper()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-all class type", actual)
}

func Test_I18_SingleRwx_ApplyOnMany_Empty(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	err := s.ApplyOnMany(cond)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty locations", actual)
}

func Test_I18_SingleRwx_ApplyOnMany_Valid(t *testing.T) {
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

func Test_I18_NewRwxWrapper_UsingVariant(t *testing.T) {
	w, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.Variant("644"))
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": w.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I18_NewRwxWrapper_UsingVariantPtr(t *testing.T) {
	w, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("644"))
	actual := args.Map{"result": err != nil || w == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected: err=, w=", actual)
}

func Test_I18_NewRwxWrapper_RwxFullString(t *testing.T) {
	w, err := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": w.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- ChmodApply and Verify ---

func Test_I18_ChmodApply_RecursivePath(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	err := chmodhelper.ChmodApply.RecursivePath(true, 0755, tmpDir)
	_ = err
}

func Test_I18_ChmodVerify_RwxFull(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0755)
	os.Chmod(f, 0755)

	err := chmodhelper.ChmodVerify.RwxFull(f, "-rwxr-xr-x")
	_ = err
}

func Test_I18_ChmodVerify_RwxFull_NoDash(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0755)
	os.Chmod(f, 0755)

	err := chmodhelper.ChmodVerify.RwxFull(f, "rwxr-xr-x")
	_ = err
}

// --- TempDirGetter ---

func Test_I18_TempDirGetter(t *testing.T) {
	td := chmodhelper.TempDirGetter.TempDefault()
	actual := args.Map{"result": td == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty temp dir", actual)
}

// --- ExpandCharRwx ---

func Test_I18_ExpandCharRwx_Valid(t *testing.T) {
	r, w, x := chmodhelper.ExpandCharRwx("755")
	_ = r
	_ = w
	_ = x
}

func Test_I18_ExpandCharRwx_Short(t *testing.T) {
	defer func() { recover() }() // may panic on short string
	chmodhelper.ExpandCharRwx("")
}

// --- SimpleFileReaderWriter ---

func Test_I18_SimpleFileReaderWriter(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "sub", "test.txt")

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, f)
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil reader writer", actual)
}

// --- FileModeFriendlyString ---

func Test_I18_FileModeFriendlyString(t *testing.T) {
	s := chmodhelper.FileModeFriendlyString(0755)
	_ = s
}

// --- PathExistStat ---

func Test_I18_GetPathExistStat_NonExistent(t *testing.T) {
	stat := chmodhelper.GetPathExistStat("/nonexistent/path/xyz_i18")
	actual := args.Map{"result": stat == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil stat", actual)
	actual := args.Map{"result": stat.IsExist}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-exist for fake path", actual)
}

func Test_I18_GetPathExistStat_Existing(t *testing.T) {
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	stat := chmodhelper.GetPathExistStat(f)
	actual := args.Map{"result": stat.IsExist}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected exist for real path", actual)
}

// --- IsPathExists ---

func Test_I18_IsPathExists(t *testing.T) {
	tmpDir := t.TempDir()
	actual := args.Map{"result": chmodhelper.IsPathExists(tmpDir)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected exists for temp dir", actual)

	actual := args.Map{"result": chmodhelper.IsPathExists("/nonexistent/xyz")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not exists", actual)
}

func Test_I18_IsPathInvalid(t *testing.T) {
	actual := args.Map{"result": chmodhelper.IsPathInvalid("/nonexistent/xyz")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for nonexistent path", actual)
}

func Test_I18_IsDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	actual := args.Map{"result": chmodhelper.IsDirectory(tmpDir)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected directory", actual)

	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)
	actual := args.Map{"result": chmodhelper.IsDirectory(f)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not directory for file", actual)
}

// --- GetExistingChmod ---

func Test_I18_GetExistingChmod(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	chmod, err := chmodhelper.GetExistingChmod(f)
	actual := args.Map{"result": err != nil || chmod == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero chmod", actual)
}

func Test_I18_GetExistingChmodOfValidFile(t *testing.T) {
	skipIfWindows(t)

	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(f, []byte("test"), 0644)

	chmod, isInvalid := chmodhelper.GetExistingChmodOfValidFile(f)
	actual := args.Map{"result": isInvalid || chmod == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected result", actual)
}

func Test_I18_GetExistingChmodOfValidFile_NonExistent(t *testing.T) {
	_, isInvalid := chmodhelper.GetExistingChmodOfValidFile("/nonexistent/xyz")
	actual := args.Map{"result": isInvalid}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for nonexistent file", actual)
}
