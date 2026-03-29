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

// ── Attribute — uncovered branches ──

func Test_Cov5_Attribute_AllMethods(t *testing.T) {
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	emptyAttr := chmodhelper.Attribute{}
	var nilAttr *chmodhelper.Attribute

	actual := args.Map{
		"isNull":    nilAttr.IsNull(),
		"isAnyNull": nilAttr.IsAnyNull(),
		"isEmpty":   emptyAttr.IsEmpty(),
		"isZero":    emptyAttr.IsZero(),
		"isInvalid": emptyAttr.IsInvalid(),
		"isDefined": attr.IsDefined(),
		"hasAny":    attr.HasAnyItem(),
		"toByte":    attr.ToByte(),
		"toSum":     attr.ToSum(),
		"strByte":   attr.ToStringByte(),
		"rwxLen":    len(attr.ToRwxString()),
	}
	expected := args.Map{
		"isNull": true, "isAnyNull": true,
		"isEmpty": true, "isZero": true, "isInvalid": true,
		"isDefined": true, "hasAny": true,
		"toByte": byte(7), "toSum": byte(7),
		"strByte": byte('7'), "rwxLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "Attribute returns correct value -- all methods", actual)
}

func Test_Cov5_Attribute_ToSpecificBytes(t *testing.T) {
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: false, IsExecute: true}
	r, w, e, sum := attr.ToSpecificBytes()
	actual := args.Map{"r": r, "w": w, "e": e, "sum": sum}
	expected := args.Map{"r": byte(4), "w": byte(0), "e": byte(1), "sum": byte(5)}
	expected.ShouldBeEqual(t, 0, "ToSpecificBytes returns correct value -- with args", actual)
}

func Test_Cov5_Attribute_ToAttributeValue(t *testing.T) {
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: false}
	av := attr.ToAttributeValue()
	actual := args.Map{"sum": av.Sum}
	expected := args.Map{"sum": byte(6)}
	expected.ShouldBeEqual(t, 0, "ToAttributeValue returns correct value -- with args", actual)
}

func Test_Cov5_Attribute_ToRwx(t *testing.T) {
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: false, IsExecute: false}
	rwx := attr.ToRwx()
	actual := args.Map{"r": rwx[0], "w": rwx[1], "x": rwx[2]}
	expected := args.Map{"r": byte('r'), "w": byte('-'), "x": byte('-')}
	expected.ShouldBeEqual(t, 0, "ToRwx returns correct value -- with args", actual)
}

func Test_Cov5_Attribute_ToVariant(t *testing.T) {
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	v := attr.ToVariant()
	actual := args.Map{"val": v.Value()}
	expected := args.Map{"val": byte(7)}
	expected.ShouldBeEqual(t, 0, "ToVariant returns correct value -- with args", actual)
}

func Test_Cov5_Attribute_Clone_Nil(t *testing.T) {
	var nilAttr *chmodhelper.Attribute
	actual := args.Map{"nil": nilAttr.Clone() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
}

func Test_Cov5_Attribute_Clone_Valid(t *testing.T) {
	attr := &chmodhelper.Attribute{IsRead: true}
	c := attr.Clone()
	actual := args.Map{"read": c.IsRead}
	expected := args.Map{"read": true}
	expected.ShouldBeEqual(t, 0, "Clone returns non-empty -- valid", actual)
}

func Test_Cov5_Attribute_IsEqualPtr(t *testing.T) {
	a := &chmodhelper.Attribute{IsRead: true}
	b := &chmodhelper.Attribute{IsRead: true}
	c := &chmodhelper.Attribute{IsRead: false}
	var nilA *chmodhelper.Attribute
	actual := args.Map{
		"equal": a.IsEqualPtr(b), "notEqual": a.IsEqualPtr(c),
		"bothNil": nilA.IsEqualPtr(nil), "oneNil": a.IsEqualPtr(nil),
	}
	expected := args.Map{"equal": true, "notEqual": false, "bothNil": true, "oneNil": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- with args", actual)
}

func Test_Cov5_Attribute_IsEqual(t *testing.T) {
	a := chmodhelper.Attribute{IsRead: true}
	b := chmodhelper.Attribute{IsRead: true}
	c := chmodhelper.Attribute{IsRead: false}
	actual := args.Map{"equal": a.IsEqual(b), "notEqual": a.IsEqual(c)}
	expected := args.Map{"equal": true, "notEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

// ── AttrVariant ──

func Test_Cov5_AttrVariant_All(t *testing.T) {
	v := chmodhelper.ReadWriteExecute
	actual := args.Map{
		"str":    v.String() != "",
		"val":    v.Value(),
		"gt":     v.IsGreaterThan(8),
		"notGt":  v.IsGreaterThan(3),
		"attrOk": func() byte { a := v.ToAttribute(); return a.ToByte() }(),
	}
	expected := args.Map{"str": true, "val": byte(7), "gt": true, "notGt": false, "attrOk": byte(7)}
	expected.ShouldBeEqual(t, 0, "AttrVariant returns correct value -- all", actual)
}

// ── Variant ──

func Test_Cov5_Variant_String(t *testing.T) {
	v := chmodhelper.X755
	actual := args.Map{"val": v.String()}
	expected := args.Map{"val": "755"}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- String", actual)
}

func Test_Cov5_Variant_ExpandOctalByte(t *testing.T) {
	r, w, x := chmodhelper.X755.ExpandOctalByte()
	actual := args.Map{"r": r, "w": w, "x": x}
	expected := args.Map{"r": byte('7'), "w": byte('5'), "x": byte('5')}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ExpandOctalByte", actual)
}

func Test_Cov5_Variant_ToWrapper(t *testing.T) {
	rwx, err := chmodhelper.X755.ToWrapper()
	actual := args.Map{"noErr": err == nil, "notEmpty": !rwx.IsEmpty()}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToWrapper", actual)
}

func Test_Cov5_Variant_ToWrapperPtr(t *testing.T) {
	rwx, err := chmodhelper.X755.ToWrapperPtr()
	actual := args.Map{"noErr": err == nil, "notNil": rwx != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToWrapperPtr", actual)
}

// ── ExpandCharRwx ──

func Test_Cov5_ExpandCharRwx(t *testing.T) {
	r, w, x := chmodhelper.ExpandCharRwx("755")
	actual := args.Map{"r": r, "w": w, "x": x}
	expected := args.Map{"r": byte('7'), "w": byte('5'), "x": byte('5')}
	expected.ShouldBeEqual(t, 0, "ExpandCharRwx returns correct value -- with args", actual)
}

// ── IsChmod ──

func Test_Cov5_IsChmod_ShortString(t *testing.T) {
	actual := args.Map{"val": chmodhelper.IsChmod("/tmp", "rwx")}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsChmod returns correct value -- short", actual)
}

func Test_Cov5_IsChmod_EmptyLoc(t *testing.T) {
	actual := args.Map{"val": chmodhelper.IsChmod("", "-rwxrwxrwx")}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsChmod returns empty -- empty loc", actual)
}

func Test_Cov5_IsChmod_NonExistent(t *testing.T) {
	actual := args.Map{"val": chmodhelper.IsChmod("/nonexistent_cov5_xyz", "-rwxrwxrwx")}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsChmod returns non-empty -- non-existent", actual)
}

// ── FileModeFriendlyString ──

func Test_Cov5_FileModeFriendlyString(t *testing.T) {
	result := chmodhelper.FileModeFriendlyString(0755)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileModeFriendlyString returns correct value -- with args", actual)
}

// ── RwxWrapper — uncovered branches ──

func Test_Cov5_RwxWrapper_IsEqualPtr(t *testing.T) {
	a, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	b, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	c, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	var nilRwx *chmodhelper.RwxWrapper
	actual := args.Map{
		"equal":    a.ToPtr().IsEqualPtr(b.ToPtr()),
		"notEqual": a.ToPtr().IsEqualPtr(c.ToPtr()),
		"bothNil":  nilRwx.IsEqualPtr(nil),
		"oneNil":   a.ToPtr().IsEqualPtr(nil),
	}
	expected := args.Map{"equal": true, "notEqual": false, "bothNil": true, "oneNil": false}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- IsEqualPtr", actual)
}

func Test_Cov5_RwxWrapper_IsEqualFileMode(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{
		"match":    rwx.IsEqualFileMode(0755),
		"notMatch": rwx.IsNotEqualFileMode(0644),
	}
	expected := args.Map{"match": true, "notMatch": true}
	expected.ShouldBeEqual(t, 0, "IsEqualFileMode returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_IsRwxFullEqual(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{
		"match": rwx.IsRwxFullEqual("-rwxr-xr-x"),
		"short": rwx.IsRwxFullEqual("rwx"),
	}
	expected := args.Map{"match": true, "short": false}
	expected.ShouldBeEqual(t, 0, "IsRwxFullEqual returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_IsRwxEqualLocation(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{"nonExist": rwx.IsRwxEqualLocation("/nonexistent_cov5")}
	expected := args.Map{"nonExist": false}
	expected.ShouldBeEqual(t, 0, "IsRwxEqualLocation returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_IsRwxEqualFileInfo(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{"nil": rwx.IsRwxEqualFileInfo(nil)}
	expected := args.Map{"nil": false}
	expected.ShouldBeEqual(t, 0, "IsRwxEqualFileInfo returns nil -- nil", actual)
}

func Test_Cov5_RwxWrapper_IsEqualVarWrapper_Nil(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{"nil": rwx.IsEqualVarWrapper(nil)}
	expected := args.Map{"nil": false}
	expected.ShouldBeEqual(t, 0, "IsEqualVarWrapper returns nil -- nil", actual)
}

func Test_Cov5_RwxWrapper_Clone_Nil(t *testing.T) {
	var nilRwx *chmodhelper.RwxWrapper
	actual := args.Map{"nil": nilRwx.Clone() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
}

func Test_Cov5_RwxWrapper_Clone_Valid(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	c := rwx.Clone()
	actual := args.Map{"notNil": c != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns non-empty -- valid", actual)
}

func Test_Cov5_RwxWrapper_ToPtr_ToNonPtr(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	p := rwx.ToPtr()
	np := p.ToNonPtr()
	actual := args.Map{"ptrNotNil": p != nil, "npNotEmpty": !np.IsEmpty()}
	expected := args.Map{"ptrNotNil": true, "npNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToPtr/ToNonPtr returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_ToRwxOwnerGroupOther(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	ogo := rwx.ToRwxOwnerGroupOther()
	actual := args.Map{"notNil": ogo != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToRwxOwnerGroupOther returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_ToRwxInstruction(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	cond := &chmodins.Condition{}
	ins := rwx.ToRwxInstruction(cond)
	actual := args.Map{"notNil": ins != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToRwxInstruction returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_MarshalUnmarshalJSON(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	b, err := rwx.MarshalJSON()
	var rwx2 chmodhelper.RwxWrapper
	err2 := rwx2.UnmarshalJSON(b)
	actual := args.Map{"noErr": err == nil, "noErr2": err2 == nil, "match": rwx2.ToFullRwxValueString() == rwx.ToFullRwxValueString()}
	expected := args.Map{"noErr": true, "noErr2": true, "match": true}
	expected.ShouldBeEqual(t, 0, "MarshalUnmarshalJSON returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_FriendlyDisplay(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{"notEmpty": rwx.FriendlyDisplay() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FriendlyDisplay returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_Json(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	r := rwx.Json()
	rp := rwx.JsonPtr()
	actual := args.Map{"noErr": !r.HasError(), "ptrNotNil": rp != nil}
	expected := args.Map{"noErr": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Json/JsonPtr returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_AsJsonContractsBinder(t *testing.T) {
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	actual := args.Map{"notNil": rwx.AsJsonContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_VerifyPaths(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	dir := t.TempDir()
	_ = os.Chmod(dir, 0755)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.VerifyPaths(true, dir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyPaths returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_ApplyLinuxChmodOnMany(t *testing.T) {
	dir := t.TempDir()
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	cond := &chmodins.Condition{IsContinueOnError: true}
	err := rwx.ApplyLinuxChmodOnMany(cond, dir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyLinuxChmodOnMany returns correct value -- with args", actual)
}

func Test_Cov5_RwxWrapper_ApplyLinuxChmodOnMany_Recursive(t *testing.T) {
	dir := t.TempDir()
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	cond := &chmodins.Condition{IsRecursive: true, IsContinueOnError: true}
	err := rwx.ApplyLinuxChmodOnMany(cond, dir)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyLinuxChmodOnMany returns correct value -- recursive", actual)
}

// ── newRwxWrapperCreator — uncovered branches ──

func Test_Cov5_NewRwxWrapper_CreatePtr(t *testing.T) {
	rwx, err := chmodhelper.New.RwxWrapper.CreatePtr("755")
	actual := args.Map{"noErr": err == nil, "notNil": rwx != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper returns correct value -- CreatePtr", actual)
}

func Test_Cov5_NewRwxWrapper_UsingBytes(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingBytes([3]byte{7, 5, 5})
	actual := args.Map{"notEmpty": !rwx.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "UsingBytes returns correct value -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_Invalid(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.Invalid()
	actual := args.Map{"empty": rwx.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Invalid returns error -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_InvalidPtr(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.InvalidPtr()
	actual := args.Map{"empty": rwx.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "InvalidPtr returns error -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_Empty(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.Empty()
	actual := args.Map{"empty": rwx.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_UsingFileMode_Zero(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0)
	actual := args.Map{"empty": rwx.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "UsingFileMode returns correct value -- zero", actual)
}

func Test_Cov5_NewRwxWrapper_UsingFileModePtr_Zero(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingFileModePtr(0)
	actual := args.Map{"empty": rwx.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "UsingFileModePtr returns correct value -- zero", actual)
}

func Test_Cov5_NewRwxWrapper_UsingAttrVariants(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingAttrVariants(chmodhelper.ReadWriteExecute, chmodhelper.ReadExecute, chmodhelper.ReadExecute)
	actual := args.Map{"notEmpty": !rwx.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "UsingAttrVariants returns correct value -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_UsingAttrs(t *testing.T) {
	rwx := chmodhelper.New.RwxWrapper.UsingAttrs(
		chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true},
		chmodhelper.Attribute{IsRead: true, IsExecute: true},
		chmodhelper.Attribute{IsRead: true, IsExecute: true},
	)
	actual := args.Map{"notEmpty": !rwx.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "UsingAttrs returns correct value -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_Rwx10(t *testing.T) {
	rwx, err := chmodhelper.New.RwxWrapper.Rwx10("-rwxr-xr-x")
	actual := args.Map{"noErr": err == nil, "notEmpty": !rwx.IsEmpty()}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Rwx10 returns correct value -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_Rwx9(t *testing.T) {
	rwx, err := chmodhelper.New.RwxWrapper.Rwx9("rwxr-xr-x")
	actual := args.Map{"noErr": err == nil, "notEmpty": !rwx.IsEmpty()}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Rwx9 returns correct value -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_RwxFullString_BadLen(t *testing.T) {
	_, err := chmodhelper.New.RwxWrapper.RwxFullString("rwx")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxFullString returns correct value -- bad len", actual)
}

func Test_Cov5_NewRwxWrapper_RwxFullStringWtHyphen_BadLen(t *testing.T) {
	_, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rw")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxFullStringWtHyphen returns correct value -- bad len", actual)
}

func Test_Cov5_NewRwxWrapper_UsingExistingFile(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "exist.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	rwx, err := chmodhelper.New.RwxWrapper.UsingExistingFile(fp)
	actual := args.Map{"noErr": err == nil, "notNil": rwx != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "UsingExistingFile returns correct value -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_UsingExistingFileSkipInvalidFile(t *testing.T) {
	rwx, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileSkipInvalidFile("/nonexistent_cov5")
	actual := args.Map{"notNil": rwx != nil, "isInvalid": isInvalid}
	expected := args.Map{"notNil": true, "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "UsingExistingFileSkipInvalidFile returns error -- with args", actual)
}

func Test_Cov5_NewRwxWrapper_UsingExistingFileOption_Skip(t *testing.T) {
	rwx, err, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileOption(true, "/nonexistent_cov5")
	actual := args.Map{"notNil": rwx != nil, "noErr": err == nil, "isInvalid": isInvalid}
	expected := args.Map{"notNil": true, "noErr": true, "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "UsingExistingFileOption returns correct value -- skip", actual)
}

func Test_Cov5_NewRwxWrapper_UsingExistingFileOption_NoSkip(t *testing.T) {
	rwx, err, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileOption(false, "/nonexistent_cov5")
	actual := args.Map{"notNil": rwx != nil, "hasErr": err != nil, "isInvalid": isInvalid}
	expected := args.Map{"notNil": true, "hasErr": true, "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "UsingExistingFileOption returns correct value -- no-skip", actual)
}

func Test_Cov5_NewRwxWrapper_Instruction(t *testing.T) {
	ins, err := chmodhelper.New.RwxWrapper.Instruction("-rwxr-xr-x", chmodins.Condition{})
	actual := args.Map{"noErr": err == nil, "notNil": ins != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Instruction returns correct value -- with args", actual)
}

// ── SingleRwx ──

func Test_Cov5_SingleRwx_AllClassTypes(t *testing.T) {
	classTypes := []chmodclasstype.Variant{
		chmodclasstype.All, chmodclasstype.Owner, chmodclasstype.Group,
		chmodclasstype.Other, chmodclasstype.OwnerGroup,
		chmodclasstype.GroupOther, chmodclasstype.OwnerOther,
	}
	for _, ct := range classTypes {
		s, err := chmodhelper.NewSingleRwx("rwx", ct)
		if err != nil {
			continue
		}
		ogo := s.ToRwxOwnerGroupOther()
		actual := args.Map{"notNil": ogo != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- "+ct.Name(), actual)
	}
}

func Test_Cov5_SingleRwx_BadLength(t *testing.T) {
	_, err := chmodhelper.NewSingleRwx("rw", chmodclasstype.All)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- bad length", actual)
}

func Test_Cov5_SingleRwx_ToRwxInstruction(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	ins := s.ToRwxInstruction(cond)
	actual := args.Map{"notNil": ins != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToRwxInstruction", actual)
}

func Test_Cov5_SingleRwx_ToVarRwxWrapper(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	vw, err := s.ToVarRwxWrapper()
	actual := args.Map{"noErr": err == nil, "notNil": vw != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToVarRwxWrapper", actual)
}

func Test_Cov5_SingleRwx_ToDisabledRwxWrapper(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	rwx, err := s.ToDisabledRwxWrapper()
	actual := args.Map{"noErr": err == nil, "notNil": rwx != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToDisabledRwxWrapper", actual)
}

func Test_Cov5_SingleRwx_ToRwxWrapper_All(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	rwx, err := s.ToRwxWrapper()
	actual := args.Map{"noErr": err == nil, "notNil": rwx != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToRwxWrapper all", actual)
}

func Test_Cov5_SingleRwx_ToRwxWrapper_NonAll(t *testing.T) {
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	_, err := s.ToRwxWrapper()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns non-empty -- ToRwxWrapper non-all", actual)
}

// ── RwxVariableWrapper ──

func Test_Cov5_RwxVariableWrapper_Clone(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	c := vw.Clone()
	var nilVw *chmodhelper.RwxVariableWrapper
	actual := args.Map{"notNil": c != nil, "nilClone": nilVw.Clone() == nil}
	expected := args.Map{"notNil": true, "nilClone": true}
	expected.ShouldBeEqual(t, 0, "RwxVariableWrapper returns correct value -- Clone", actual)
}

func Test_Cov5_RwxVariableWrapper_IsEqualPtr(t *testing.T) {
	a, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	b, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	var nilVw *chmodhelper.RwxVariableWrapper
	actual := args.Map{
		"equal":   a.IsEqualPtr(b),
		"bothNil": nilVw.IsEqualPtr(nil),
		"oneNil":  a.IsEqualPtr(nil),
	}
	expected := args.Map{"equal": true, "bothNil": true, "oneNil": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- with args", actual)
}

func Test_Cov5_RwxVariableWrapper_ToString(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	withH := vw.ToString(true)
	withoutH := vw.ToString(false)
	actual := args.Map{"withH": len(withH), "withoutH": len(withoutH)}
	expected := args.Map{"withH": 10, "withoutH": 9}
	expected.ShouldBeEqual(t, 0, "ToString returns correct value -- with args", actual)
}

// ── RwxMatchingStatus ──

func Test_Cov5_RwxMatchingStatus_Invalid(t *testing.T) {
	s := chmodhelper.InvalidRwxMatchingStatus(nil)
	actual := args.Map{"notNil": s != nil, "notAll": !s.IsAllMatching}
	expected := args.Map{"notNil": true, "notAll": true}
	expected.ShouldBeEqual(t, 0, "InvalidRwxMatchingStatus returns error -- with args", actual)
}

func Test_Cov5_RwxMatchingStatus_Empty(t *testing.T) {
	s := chmodhelper.EmptyRwxMatchingStatus()
	actual := args.Map{"notNil": s != nil, "missingStr": s.MissingFilesToString()}
	expected := args.Map{"notNil": true, "missingStr": ""}
	expected.ShouldBeEqual(t, 0, "EmptyRwxMatchingStatus returns empty -- with args", actual)
}

func Test_Cov5_RwxMatchingStatus_CreateErrFinalError_AllMatching(t *testing.T) {
	s := &chmodhelper.RwxMatchingStatus{IsAllMatching: true}
	actual := args.Map{"nil": s.CreateErrFinalError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "CreateErrFinalError returns error -- all matching", actual)
}

// ── FilteredPathFileInfoMap ──

func Test_Cov5_FilteredPathFileInfoMap_Invalid(t *testing.T) {
	m := chmodhelper.InvalidFilteredPathFileInfoMap()
	actual := args.Map{
		"notNil":   m != nil,
		"emptyV":   m.IsEmptyValidFileInfos(),
		"emptyI":   m.IsEmptyIssues(),
		"noIssues": !m.HasAnyIssues(),
		"noErr":    !m.HasError(),
		"noMiss":   !m.HasAnyMissingPaths(),
		"missStr":  m.MissingPathsToString(),
	}
	expected := args.Map{
		"notNil": true, "emptyV": true, "emptyI": true,
		"noIssues": true, "noErr": true, "noMiss": true, "missStr": "",
	}
	expected.ShouldBeEqual(t, 0, "InvalidFilteredPathFileInfoMap returns error -- with args", actual)
}

// ── PathExistStat ──

func Test_Cov5_PathExistStat_AllMethods(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "stat.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	stat := chmodhelper.GetPathExistStat(fp)
	actual := args.Map{
		"hasError":   stat.HasError(),
		"emptyErr":   stat.IsEmptyError(),
		"hasFileInfo": stat.HasFileInfo(),
		"isFile":     stat.IsFile(),
		"isDir":      stat.IsDir(),
		"isInvalid":  stat.IsInvalid(),
		"hasIssues":  stat.HasAnyIssues(),
		"mode":       stat.FileMode() != nil,
		"size":       stat.Size() != nil,
		"lastMod":    stat.LastModifiedDate() != nil,
		"fileName":   stat.FileName() != "",
		"parentDir":  stat.ParentDir() != "",
		"dotExt":     stat.DotExt(),
		"str":        stat.String() != "",
	}
	expected := args.Map{
		"hasError": false, "emptyErr": true, "hasFileInfo": true,
		"isFile": true, "isDir": false, "isInvalid": false,
		"hasIssues": false, "mode": true, "size": true,
		"lastMod": true, "fileName": true, "parentDir": true,
		"dotExt": ".txt", "str": true,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns correct value -- all", actual)
}

func Test_Cov5_PathExistStat_Nil(t *testing.T) {
	var stat *chmodhelper.PathExistStat
	actual := args.Map{
		"isInvalid": stat.HasAnyIssues(),
		"str":       stat.String(),
		"dispose":   true,
		"notExist":  stat.NotExistError() == nil,
		"notFile":   stat.NotAFileError() == nil,
		"notDir":    stat.NotADirError() == nil,
		"msgPath":   stat.MessageWithPathWrapped("x"),
		"meaningful": stat.MeaningFullError() == nil,
	}
	stat.Dispose()
	expected := args.Map{
		"isInvalid": true, "str": "", "dispose": true,
		"notExist": true, "notFile": true, "notDir": true,
		"msgPath": "", "meaningful": true,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns nil -- nil", actual)
}

func Test_Cov5_PathExistStat_NonExist(t *testing.T) {
	stat := chmodhelper.GetPathExistStat("/nonexistent_cov5_xyz")
	actual := args.Map{
		"fileMode": stat.FileMode() == nil,
		"size":     stat.Size() == nil,
		"lastMod":  stat.LastModifiedDate() == nil,
	}
	expected := args.Map{"fileMode": true, "size": true, "lastMod": true}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns non-empty -- non-exist", actual)
}

func Test_Cov5_PathExistStat_Dispose(t *testing.T) {
	stat := chmodhelper.GetPathExistStat(t.TempDir())
	stat.Dispose()
	actual := args.Map{"loc": stat.Location, "exist": stat.IsExist}
	expected := args.Map{"loc": "", "exist": false}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns correct value -- Dispose", actual)
}

func Test_Cov5_PathExistStat_Parent(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "sub.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	stat := chmodhelper.GetPathExistStat(fp)
	parent := stat.Parent()
	combine := stat.CombineWithNewPath("extra")
	combineWith := stat.CombineWith("extra")
	parentWith := stat.ParentWith("extra")
	parentNewPath := stat.ParentWithNewPath("extra")
	actual := args.Map{
		"parentExist":  parent.IsExist,
		"combine":      combine != "",
		"combineWith":  combineWith != nil,
		"parentWith":   parentWith != nil,
		"parentNew":    parentNewPath != "",
	}
	expected := args.Map{
		"parentExist": true, "combine": true, "combineWith": true,
		"parentWith": true, "parentNew": true,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns correct value -- Parent", actual)
}

func Test_Cov5_PathExistStat_NotAFileError_Dir(t *testing.T) {
	stat := chmodhelper.GetPathExistStat(t.TempDir())
	err := stat.NotAFileError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAFileError returns error -- dir", actual)
}

func Test_Cov5_PathExistStat_NotADirError_File(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	stat := chmodhelper.GetPathExistStat(fp)
	err := stat.NotADirError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotADirError returns error -- file", actual)
}

// ── SimpleFileReaderWriter — uncovered branches ──

func Test_Cov5_SimpleFileRW_MarshalUnmarshalJSON(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: "/tmp", FilePath: "/tmp/test.txt"}
	b, err := rw.MarshalJSON()
	var rw2 chmodhelper.SimpleFileReaderWriter
	err2 := rw2.UnmarshalJSON(b)
	actual := args.Map{"noErr": err == nil, "noErr2": err2 == nil, "path": rw2.FilePath}
	expected := args.Map{"noErr": true, "noErr2": true, "path": "/tmp/test.txt"}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- MarshalUnmarshalJSON", actual)
}

func Test_Cov5_SimpleFileRW_Clone(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	c := rw.Clone()
	actual := args.Map{"path": c.FilePath}
	expected := args.Map{"path": "/tmp/test.txt"}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- Clone", actual)
}

func Test_Cov5_SimpleFileRW_ClonePtr_Nil(t *testing.T) {
	var rw *chmodhelper.SimpleFileReaderWriter
	actual := args.Map{"nil": rw.ClonePtr() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns nil -- ClonePtr nil", actual)
}

func Test_Cov5_SimpleFileRW_Json(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	r := rw.Json()
	rp := rw.JsonPtr()
	actual := args.Map{"noErr": !r.HasError(), "ptrNotNil": rp != nil}
	expected := args.Map{"noErr": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- Json", actual)
}

func Test_Cov5_SimpleFileRW_AsJsonContractsBinder(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644}
	actual := args.Map{"notNil": rw.AsJsonContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_Cov5_SimpleFileRW_String(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	actual := args.Map{"notEmpty": rw.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- String", actual)
}

func Test_Cov5_SimpleFileRW_Expire(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "expire.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: fp}
	err := rw.Expire()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Expire returns correct value -- with args", actual)
}

func Test_Cov5_SimpleFileRW_Expire_NonExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent_cov5"}
	err := rw.Expire()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Expire returns non-empty -- non-exist", actual)
}

func Test_Cov5_SimpleFileRW_RemoveOnExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent_cov5"}
	err := rw.RemoveOnExist()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveOnExist returns correct value -- with args", actual)
}

func Test_Cov5_SimpleFileRW_RemoveDirOnExist(t *testing.T) {
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/nonexistent_cov5_dir"}
	err := rw.RemoveDirOnExist()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirOnExist returns correct value -- with args", actual)
}

// ── newAttributeCreator — uncovered branches ──

func Test_Cov5_NewAttribute_Create(t *testing.T) {
	attr := chmodhelper.New.Attribute.Create(true, false, true)
	actual := args.Map{"read": attr.IsRead, "write": attr.IsWrite, "exe": attr.IsExecute}
	expected := args.Map{"read": true, "write": false, "exe": true}
	expected.ShouldBeEqual(t, 0, "NewAttribute returns correct value -- Create", actual)
}

func Test_Cov5_NewAttribute_Default(t *testing.T) {
	attr := chmodhelper.New.Attribute.Default(false, true, false)
	actual := args.Map{"write": attr.IsWrite}
	expected := args.Map{"write": true}
	expected.ShouldBeEqual(t, 0, "NewAttribute returns correct value -- Default", actual)
}

func Test_Cov5_NewAttribute_UsingByte(t *testing.T) {
	attr, err := chmodhelper.New.Attribute.UsingByte(5)
	actual := args.Map{"noErr": err == nil, "read": attr.IsRead, "exe": attr.IsExecute}
	expected := args.Map{"noErr": true, "read": true, "exe": true}
	expected.ShouldBeEqual(t, 0, "UsingByte returns correct value -- with args", actual)
}

func Test_Cov5_NewAttribute_UsingVariant(t *testing.T) {
	attr, err := chmodhelper.New.Attribute.UsingVariant(chmodhelper.ReadWrite)
	actual := args.Map{"noErr": err == nil, "read": attr.IsRead, "write": attr.IsWrite}
	expected := args.Map{"noErr": true, "read": true, "write": true}
	expected.ShouldBeEqual(t, 0, "UsingVariant returns correct value -- with args", actual)
}

// ── VarAttribute — uncovered branches ──

func Test_Cov5_VarAttribute_ToCompileFixAttr_NonFixed(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwx***r-x")
	result := vw.Group.ToCompileFixAttr()
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ToCompileFixAttr returns non-empty -- non-fixed", actual)
}

func Test_Cov5_VarAttribute_ToCompileAttr_Wildcard(t *testing.T) {
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwx***r-x")
	fixed := &chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	attr := vw.Group.ToCompileAttr(fixed)
	actual := args.Map{"read": attr.IsRead, "write": attr.IsWrite, "exe": attr.IsExecute}
	expected := args.Map{"read": true, "write": true, "exe": true}
	expected.ShouldBeEqual(t, 0, "ToCompileAttr returns correct value -- wildcard", actual)
}

// ── newSimpleFileReaderWriterCreator ──

func Test_Cov5_NewSimpleFileRW_Default(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW returns correct value -- Default", actual)
}

func Test_Cov5_NewSimpleFileRW_Path(t *testing.T) {
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/test.txt")
	actual := args.Map{"notNil": rw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW returns correct value -- Path", actual)
}

// ── RwxInstructionExecutors ──

func Test_Cov5_RwxInstructionExecutors_Basic(t *testing.T) {
	execs := chmodhelper.NewRwxInstructionExecutors(2)
	actual := args.Map{
		"empty":    execs.IsEmpty(),
		"hasAny":   execs.HasAnyItem(),
		"len":      execs.Length(),
		"count":    execs.Count(),
		"lastIdx":  execs.LastIndex(),
		"hasIdx0":  execs.HasIndex(0),
	}
	expected := args.Map{"empty": true, "hasAny": false, "len": 0, "count": 0, "lastIdx": -1, "hasIdx0": false}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns correct value -- basic", actual)
}

func Test_Cov5_RwxInstructionExecutors_Add_Nil(t *testing.T) {
	execs := chmodhelper.NewRwxInstructionExecutors(2)
	execs.Add(nil)
	actual := args.Map{"len": execs.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Add returns nil -- nil", actual)
}

func Test_Cov5_RwxInstructionExecutors_Items(t *testing.T) {
	execs := chmodhelper.NewRwxInstructionExecutors(2)
	actual := args.Map{"notNil": execs.Items() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Items returns correct value -- with args", actual)
}
