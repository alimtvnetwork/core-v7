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

func Test_Attribute_AllMethods(t *testing.T) {
	// Arrange
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	emptyAttr := chmodhelper.Attribute{}
	var nilAttr *chmodhelper.Attribute

	// Act
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

	// Assert
	expected := args.Map{
		"isNull": true, "isAnyNull": true,
		"isEmpty": true, "isZero": true, "isInvalid": true,
		"isDefined": true, "hasAny": true,
		"toByte": byte(7), "toSum": byte(7),
		"strByte": byte('7'), "rwxLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "Attribute returns correct value -- all methods", actual)
}

func Test_Attribute_ToSpecificBytes(t *testing.T) {
	// Arrange
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: false, IsExecute: true}
	r, w, e, sum := attr.ToSpecificBytes()

	// Act
	actual := args.Map{
		"r": r,
		"w": w,
		"e": e,
		"sum": sum,
	}

	// Assert
	expected := args.Map{
		"r": byte(4),
		"w": byte(0),
		"e": byte(1),
		"sum": byte(5),
	}
	expected.ShouldBeEqual(t, 0, "ToSpecificBytes returns correct value -- with args", actual)
}

func Test_Attribute_ToAttributeValue(t *testing.T) {
	// Arrange
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: false}
	av := attr.ToAttributeValue()

	// Act
	actual := args.Map{"sum": av.Sum}

	// Assert
	expected := args.Map{"sum": byte(6)}
	expected.ShouldBeEqual(t, 0, "ToAttributeValue returns correct value -- with args", actual)
}

func Test_Attribute_ToRwx(t *testing.T) {
	// Arrange
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: false, IsExecute: false}
	rwx := attr.ToRwx()

	// Act
	actual := args.Map{
		"r": rwx[0],
		"w": rwx[1],
		"x": rwx[2],
	}

	// Assert
	expected := args.Map{
		"r": byte('r'),
		"w": byte('-'),
		"x": byte('-'),
	}
	expected.ShouldBeEqual(t, 0, "ToRwx returns correct value -- with args", actual)
}

func Test_Attribute_ToVariant(t *testing.T) {
	// Arrange
	attr := chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	v := attr.ToVariant()

	// Act
	actual := args.Map{"val": v.Value()}

	// Assert
	expected := args.Map{"val": byte(7)}
	expected.ShouldBeEqual(t, 0, "ToVariant returns correct value -- with args", actual)
}

func Test_Attribute_Clone_Nil(t *testing.T) {
	// Arrange
	var nilAttr *chmodhelper.Attribute

	// Act
	actual := args.Map{"nil": nilAttr.Clone() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
}

func Test_Attribute_Clone_Valid(t *testing.T) {
	// Arrange
	attr := &chmodhelper.Attribute{IsRead: true}
	c := attr.Clone()

	// Act
	actual := args.Map{"read": c.IsRead}

	// Assert
	expected := args.Map{"read": true}
	expected.ShouldBeEqual(t, 0, "Clone returns non-empty -- valid", actual)
}

func Test_Attribute_IsEqualPtr(t *testing.T) {
	// Arrange
	a := &chmodhelper.Attribute{IsRead: true}
	b := &chmodhelper.Attribute{IsRead: true}
	c := &chmodhelper.Attribute{IsRead: false}
	var nilA *chmodhelper.Attribute

	// Act
	actual := args.Map{
		"equal": a.IsEqualPtr(b), "notEqual": a.IsEqualPtr(c),
		"bothNil": nilA.IsEqualPtr(nil), "oneNil": a.IsEqualPtr(nil),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
		"bothNil": true,
		"oneNil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- with args", actual)
}

func Test_Attribute_IsEqual(t *testing.T) {
	// Arrange
	a := chmodhelper.Attribute{IsRead: true}
	b := chmodhelper.Attribute{IsRead: true}
	c := chmodhelper.Attribute{IsRead: false}

	// Act
	actual := args.Map{
		"equal": a.IsEqual(b),
		"notEqual": a.IsEqual(c),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

// ── AttrVariant ──

func Test_AttrVariant_All(t *testing.T) {
	// Arrange
	v := chmodhelper.ReadWriteExecute

	// Act
	actual := args.Map{
		"str":    v.String() != "",
		"val":    v.Value(),
		"gt":     v.IsGreaterThan(8),
		"notGt":  v.IsGreaterThan(3),
		"attrOk": func() byte { a := v.ToAttribute(); return a.ToByte() }(),
	}

	// Assert
	expected := args.Map{
		"str": true,
		"val": byte(7),
		"gt": true,
		"notGt": false,
		"attrOk": byte(7),
	}
	expected.ShouldBeEqual(t, 0, "AttrVariant returns correct value -- all", actual)
}

// ── Variant ──

func Test_Variant_String(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	actual := args.Map{"val": v.String()}

	// Assert
	expected := args.Map{"val": "755"}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- String", actual)
}

func Test_Variant_ExpandOctalByte(t *testing.T) {
	// Arrange
	r, w, x := chmodhelper.X755.ExpandOctalByte()

	// Act
	actual := args.Map{
		"r": r,
		"w": w,
		"x": x,
	}

	// Assert
	expected := args.Map{
		"r": byte('7'),
		"w": byte('5'),
		"x": byte('5'),
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ExpandOctalByte", actual)
}

func Test_Variant_ToWrapper(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.X755.ToWrapper()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": !rwx.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToWrapper", actual)
}

func Test_Variant_ToWrapperPtr(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.X755.ToWrapperPtr()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": rwx != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToWrapperPtr", actual)
}

// ── ExpandCharRwx ──

func Test_ExpandCharRwx(t *testing.T) {
	// Arrange
	r, w, x := chmodhelper.ExpandCharRwx("755")

	// Act
	actual := args.Map{
		"r": r,
		"w": w,
		"x": x,
	}

	// Assert
	expected := args.Map{
		"r": byte('7'),
		"w": byte('5'),
		"x": byte('5'),
	}
	expected.ShouldBeEqual(t, 0, "ExpandCharRwx returns correct value -- with args", actual)
}

// ── IsChmod ──

func Test_IsChmod_ShortString(t *testing.T) {
	// Act
	actual := args.Map{"val": chmodhelper.IsChmod("/tmp", "rwx")}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsChmod returns correct value -- short", actual)
}

func Test_IsChmod_EmptyLoc(t *testing.T) {
	// Act
	actual := args.Map{"val": chmodhelper.IsChmod("", "-rwxrwxrwx")}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsChmod returns empty -- empty loc", actual)
}

func Test_IsChmod_NonExistent(t *testing.T) {
	// Act
	actual := args.Map{"val": chmodhelper.IsChmod("/nonexistent_cov5_xyz", "-rwxrwxrwx")}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsChmod returns non-empty -- non-existent", actual)
}

// ── FileModeFriendlyString ──

func Test_FileModeFriendlyString(t *testing.T) {
	// Arrange
	result := chmodhelper.FileModeFriendlyString(0755)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileModeFriendlyString returns correct value -- with args", actual)
}

// ── RwxWrapper — uncovered branches ──

func Test_RwxWrapper_IsEqualPtr(t *testing.T) {
	// Arrange
	a, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	b, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr--")
	c, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rw-r--r--")
	var nilRwx *chmodhelper.RwxWrapper

	// Act
	actual := args.Map{
		"equal":    a.ToPtr().IsEqualPtr(b.ToPtr()),
		"notEqual": a.ToPtr().IsEqualPtr(c.ToPtr()),
		"bothNil":  nilRwx.IsEqualPtr(nil),
		"oneNil":   a.ToPtr().IsEqualPtr(nil),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
		"bothNil": true,
		"oneNil": false,
	}
	expected.ShouldBeEqual(t, 0, "RwxWrapper returns correct value -- IsEqualPtr", actual)
}

func Test_RwxWrapper_IsEqualFileMode(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{
		"match":    rwx.IsEqualFileMode(0755),
		"notMatch": rwx.IsNotEqualFileMode(0644),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"notMatch": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEqualFileMode returns correct value -- with args", actual)
}

func Test_RwxWrapper_IsRwxFullEqual(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{
		"match": rwx.IsRwxFullEqual("-rwxr-xr-x"),
		"short": rwx.IsRwxFullEqual("rwx"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"short": false,
	}
	expected.ShouldBeEqual(t, 0, "IsRwxFullEqual returns correct value -- with args", actual)
}

func Test_RwxWrapper_IsRwxEqualLocation(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{"nonExist": rwx.IsRwxEqualLocation("/nonexistent_cov5")}

	// Assert
	expected := args.Map{"nonExist": false}
	expected.ShouldBeEqual(t, 0, "IsRwxEqualLocation returns correct value -- with args", actual)
}

func Test_RwxWrapper_IsRwxEqualFileInfo(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{"nil": rwx.IsRwxEqualFileInfo(nil)}

	// Assert
	expected := args.Map{"nil": false}
	expected.ShouldBeEqual(t, 0, "IsRwxEqualFileInfo returns nil -- nil", actual)
}

func Test_RwxWrapper_IsEqualVarWrapper_Nil(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{"nil": rwx.IsEqualVarWrapper(nil)}

	// Assert
	expected := args.Map{"nil": false}
	expected.ShouldBeEqual(t, 0, "IsEqualVarWrapper returns nil -- nil", actual)
}

func Test_RwxWrapper_Clone_Nil(t *testing.T) {
	// Arrange
	var nilRwx *chmodhelper.RwxWrapper

	// Act
	actual := args.Map{"nil": nilRwx.Clone() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
}

func Test_RwxWrapper_Clone_Valid(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	c := rwx.Clone()

	// Act
	actual := args.Map{"notNil": c != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns non-empty -- valid", actual)
}

func Test_RwxWrapper_ToPtr_ToNonPtr(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	p := rwx.ToPtr()
	np := p.ToNonPtr()

	// Act
	actual := args.Map{
		"ptrNotNil": p != nil,
		"npNotEmpty": !np.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"ptrNotNil": true,
		"npNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "ToPtr/ToNonPtr returns correct value -- with args", actual)
}

func Test_RwxWrapper_ToRwxOwnerGroupOther(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	ogo := rwx.ToRwxOwnerGroupOther()

	// Act
	actual := args.Map{"notNil": ogo != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToRwxOwnerGroupOther returns correct value -- with args", actual)
}

func Test_RwxWrapper_ToRwxInstruction(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	cond := &chmodins.Condition{}
	ins := rwx.ToRwxInstruction(cond)

	// Act
	actual := args.Map{"notNil": ins != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToRwxInstruction returns correct value -- with args", actual)
}

func Test_RwxWrapper_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	b, err := rwx.MarshalJSON()
	var rwx2 chmodhelper.RwxWrapper
	err2 := rwx2.UnmarshalJSON(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"noErr2": err2 == nil,
		"match": rwx2.ToFullRwxValueString() == rwx.ToFullRwxValueString(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"noErr2": true,
		"match": true,
	}
	expected.ShouldBeEqual(t, 0, "MarshalUnmarshalJSON returns correct value -- with args", actual)
}

func Test_RwxWrapper_FriendlyDisplay(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{"notEmpty": rwx.FriendlyDisplay() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FriendlyDisplay returns correct value -- with args", actual)
}

func Test_RwxWrapper_Json(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	r := rwx.Json()
	rp := rwx.JsonPtr()

	// Act
	actual := args.Map{
		"noErr": !r.HasError(),
		"ptrNotNil": rp != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Json/JsonPtr returns correct value -- with args", actual)
}

func Test_RwxWrapper_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	actual := args.Map{"notNil": rwx.AsJsonContractsBinder() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_RwxWrapper_VerifyPaths(t *testing.T) {
	// Arrange
	if runtime.GOOS == "windows" {
		t.Skip("chmod verification differs on Windows")
	}
	dir := t.TempDir()
	_ = os.Chmod(dir, 0755)
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	err := rwx.VerifyPaths(true, dir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyPaths returns correct value -- with args", actual)
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	cond := &chmodins.Condition{IsContinueOnError: true}
	err := rwx.ApplyLinuxChmodOnMany(cond, dir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyLinuxChmodOnMany returns correct value -- with args", actual)
}

func Test_RwxWrapper_ApplyLinuxChmodOnMany_Recursive(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	rwx, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")
	cond := &chmodins.Condition{IsRecursive: true, IsContinueOnError: true}
	err := rwx.ApplyLinuxChmodOnMany(cond, dir)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyLinuxChmodOnMany returns correct value -- recursive", actual)
}

// ── newRwxWrapperCreator — uncovered branches ──

func Test_NewRwxWrapper_CreatePtr(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.New.RwxWrapper.CreatePtr("755")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": rwx != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "NewRwxWrapper returns correct value -- CreatePtr", actual)
}

func Test_NewRwxWrapper_UsingBytes(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingBytes([3]byte{7, 5, 5})

	// Act
	actual := args.Map{"notEmpty": !rwx.IsEmpty()}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "UsingBytes returns correct value -- with args", actual)
}

func Test_NewRwxWrapper_Invalid(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.Invalid()

	// Act
	actual := args.Map{"empty": rwx.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Invalid returns error -- with args", actual)
}

func Test_NewRwxWrapper_InvalidPtr(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.InvalidPtr()

	// Act
	actual := args.Map{"empty": rwx.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "InvalidPtr returns error -- with args", actual)
}

func Test_NewRwxWrapper_Empty(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.Empty()

	// Act
	actual := args.Map{"empty": rwx.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- with args", actual)
}

func Test_NewRwxWrapper_UsingFileMode_Zero(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileMode(0)

	// Act
	actual := args.Map{"empty": rwx.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "UsingFileMode returns correct value -- zero", actual)
}

func Test_NewRwxWrapper_UsingFileModePtr_Zero(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingFileModePtr(0)

	// Act
	actual := args.Map{"empty": rwx.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "UsingFileModePtr returns correct value -- zero", actual)
}

func Test_NewRwxWrapper_UsingAttrVariants(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingAttrVariants(chmodhelper.ReadWriteExecute, chmodhelper.ReadExecute, chmodhelper.ReadExecute)

	// Act
	actual := args.Map{"notEmpty": !rwx.IsEmpty()}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "UsingAttrVariants returns correct value -- with args", actual)
}

func Test_NewRwxWrapper_UsingAttrs(t *testing.T) {
	// Arrange
	rwx := chmodhelper.New.RwxWrapper.UsingAttrs(
		chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true},
		chmodhelper.Attribute{IsRead: true, IsExecute: true},
		chmodhelper.Attribute{IsRead: true, IsExecute: true},
	)

	// Act
	actual := args.Map{"notEmpty": !rwx.IsEmpty()}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "UsingAttrs returns correct value -- with args", actual)
}

func Test_NewRwxWrapper_Rwx10(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.New.RwxWrapper.Rwx10("-rwxr-xr-x")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": !rwx.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Rwx10 returns correct value -- with args", actual)
}

func Test_NewRwxWrapper_Rwx9(t *testing.T) {
	// Arrange
	rwx, err := chmodhelper.New.RwxWrapper.Rwx9("rwxr-xr-x")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": !rwx.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Rwx9 returns correct value -- with args", actual)
}

func Test_NewRwxWrapper_RwxFullString_BadLen(t *testing.T) {
	// Arrange
	_, err := chmodhelper.New.RwxWrapper.RwxFullString("rwx")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxFullString returns correct value -- bad len", actual)
}

func Test_NewRwxWrapper_RwxFullStringWtHyphen_BadLen(t *testing.T) {
	// Arrange
	_, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rw")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RwxFullStringWtHyphen returns correct value -- bad len", actual)
}

func Test_NewRwxWrapper_UsingExistingFile(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "exist.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	rwx, err := chmodhelper.New.RwxWrapper.UsingExistingFile(fp)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": rwx != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingExistingFile returns correct value -- with args", actual)
}

func Test_NewRwxWrapper_UsingExistingFileSkipInvalidFile(t *testing.T) {
	// Arrange
	rwx, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileSkipInvalidFile("/nonexistent_cov5")

	// Act
	actual := args.Map{
		"notNil": rwx != nil,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingExistingFileSkipInvalidFile returns error -- with args", actual)
}

func Test_NewRwxWrapper_UsingExistingFileOption_Skip(t *testing.T) {
	// Arrange
	rwx, err, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileOption(true, "/nonexistent_cov5")

	// Act
	actual := args.Map{
		"notNil": rwx != nil,
		"noErr": err == nil,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingExistingFileOption returns correct value -- skip", actual)
}

func Test_NewRwxWrapper_UsingExistingFileOption_NoSkip(t *testing.T) {
	// Arrange
	rwx, err, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileOption(false, "/nonexistent_cov5")

	// Act
	actual := args.Map{
		"notNil": rwx != nil,
		"hasErr": err != nil,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasErr": true,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingExistingFileOption returns correct value -- no-skip", actual)
}

func Test_NewRwxWrapper_Instruction(t *testing.T) {
	// Arrange
	ins, err := chmodhelper.New.RwxWrapper.Instruction("-rwxr-xr-x", chmodins.Condition{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": ins != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Instruction returns correct value -- with args", actual)
}

// ── SingleRwx ──

func Test_SingleRwx_AllClassTypes(t *testing.T) {
	// Arrange
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

	// Act
		actual := args.Map{"notNil": ogo != nil}

	// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- "+ct.Name(), actual)
	}
}

func Test_SingleRwx_BadLength(t *testing.T) {
	// Arrange
	_, err := chmodhelper.NewSingleRwx("rw", chmodclasstype.All)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- bad length", actual)
}

func Test_SingleRwx_ToRwxInstruction(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{}
	ins := s.ToRwxInstruction(cond)

	// Act
	actual := args.Map{"notNil": ins != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToRwxInstruction", actual)
}

func Test_SingleRwx_ToVarRwxWrapper(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	vw, err := s.ToVarRwxWrapper()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": vw != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToVarRwxWrapper", actual)
}

func Test_SingleRwx_ToDisabledRwxWrapper(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	rwx, err := s.ToDisabledRwxWrapper()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": rwx != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToDisabledRwxWrapper", actual)
}

func Test_SingleRwx_ToRwxWrapper_All(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	rwx, err := s.ToRwxWrapper()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": rwx != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns correct value -- ToRwxWrapper all", actual)
}

func Test_SingleRwx_ToRwxWrapper_NonAll(t *testing.T) {
	// Arrange
	s, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)
	_, err := s.ToRwxWrapper()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SingleRwx returns non-empty -- ToRwxWrapper non-all", actual)
}

// ── RwxVariableWrapper ──

func Test_RwxVariableWrapper_Clone(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	c := vw.Clone()
	var nilVw *chmodhelper.RwxVariableWrapper

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"nilClone": nilVw.Clone() == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxVariableWrapper returns correct value -- Clone", actual)
}

func Test_RwxVariableWrapper_IsEqualPtr(t *testing.T) {
	// Arrange
	a, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	b, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	var nilVw *chmodhelper.RwxVariableWrapper

	// Act
	actual := args.Map{
		"equal":   a.IsEqualPtr(b),
		"bothNil": nilVw.IsEqualPtr(nil),
		"oneNil":  a.IsEqualPtr(nil),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"bothNil": true,
		"oneNil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- with args", actual)
}

func Test_RwxVariableWrapper_ToString(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")
	withH := vw.ToString(true)
	withoutH := vw.ToString(false)

	// Act
	actual := args.Map{
		"withH": len(withH),
		"withoutH": len(withoutH),
	}

	// Assert
	expected := args.Map{
		"withH": 10,
		"withoutH": 9,
	}
	expected.ShouldBeEqual(t, 0, "ToString returns correct value -- with args", actual)
}

// ── RwxMatchingStatus ──

func Test_RwxMatchingStatus_Invalid(t *testing.T) {
	// Arrange
	s := chmodhelper.InvalidRwxMatchingStatus(nil)

	// Act
	actual := args.Map{
		"notNil": s != nil,
		"notAll": !s.IsAllMatching,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"notAll": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidRwxMatchingStatus returns error -- with args", actual)
}

func Test_RwxMatchingStatus_Empty(t *testing.T) {
	// Arrange
	s := chmodhelper.EmptyRwxMatchingStatus()

	// Act
	actual := args.Map{
		"notNil": s != nil,
		"missingStr": s.MissingFilesToString(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"missingStr": "",
	}
	expected.ShouldBeEqual(t, 0, "EmptyRwxMatchingStatus returns empty -- with args", actual)
}

func Test_RwxMatchingStatus_CreateErrFinalError_AllMatching(t *testing.T) {
	// Arrange
	s := &chmodhelper.RwxMatchingStatus{IsAllMatching: true}

	// Act
	actual := args.Map{"nil": s.CreateErrFinalError() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "CreateErrFinalError returns error -- all matching", actual)
}

// ── FilteredPathFileInfoMap ──

func Test_FilteredPathFileInfoMap_Invalid(t *testing.T) {
	// Arrange
	m := chmodhelper.InvalidFilteredPathFileInfoMap()

	// Act
	actual := args.Map{
		"notNil":   m != nil,
		"emptyV":   m.IsEmptyValidFileInfos(),
		"emptyI":   m.IsEmptyIssues(),
		"noIssues": !m.HasAnyIssues(),
		"noErr":    !m.HasError(),
		"noMiss":   !m.HasAnyMissingPaths(),
		"missStr":  m.MissingPathsToString(),
	}

	// Assert
	expected := args.Map{
		"notNil": true, "emptyV": true, "emptyI": true,
		"noIssues": true, "noErr": true, "noMiss": true, "missStr": "",
	}
	expected.ShouldBeEqual(t, 0, "InvalidFilteredPathFileInfoMap returns error -- with args", actual)
}

// ── PathExistStat ──

func Test_PathExistStat_AllMethods(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "stat.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	stat := chmodhelper.GetPathExistStat(fp)

	// Act
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

	// Assert
	expected := args.Map{
		"hasError": false, "emptyErr": true, "hasFileInfo": true,
		"isFile": true, "isDir": false, "isInvalid": false,
		"hasIssues": false, "mode": true, "size": true,
		"lastMod": true, "fileName": true, "parentDir": true,
		"dotExt": ".txt", "str": true,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns correct value -- all", actual)
}

func Test_PathExistStat_Nil(t *testing.T) {
	// Arrange
	var stat *chmodhelper.PathExistStat

	// Act
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

	// Assert
	expected := args.Map{
		"isInvalid": true, "str": "", "dispose": true,
		"notExist": true, "notFile": true, "notDir": true,
		"msgPath": "", "meaningful": true,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns nil -- nil", actual)
}

func Test_PathExistStat_NonExist(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat("/nonexistent_cov5_xyz")

	// Act
	actual := args.Map{
		"fileMode": stat.FileMode() == nil,
		"size":     stat.Size() == nil,
		"lastMod":  stat.LastModifiedDate() == nil,
	}

	// Assert
	expected := args.Map{
		"fileMode": true,
		"size": true,
		"lastMod": true,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns non-empty -- non-exist", actual)
}

func Test_PathExistStat_Dispose(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(t.TempDir())
	stat.Dispose()

	// Act
	actual := args.Map{
		"loc": stat.Location,
		"exist": stat.IsExist,
	}

	// Assert
	expected := args.Map{
		"loc": "",
		"exist": false,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns correct value -- Dispose", actual)
}

func Test_PathExistStat_Parent(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "sub.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	stat := chmodhelper.GetPathExistStat(fp)
	parent := stat.Parent()
	combine := stat.CombineWithNewPath("extra")
	combineWith := stat.CombineWith("extra")
	parentWith := stat.ParentWith("extra")
	parentNewPath := stat.ParentWithNewPath("extra")

	// Act
	actual := args.Map{
		"parentExist":  parent.IsExist,
		"combine":      combine != "",
		"combineWith":  combineWith != nil,
		"parentWith":   parentWith != nil,
		"parentNew":    parentNewPath != "",
	}

	// Assert
	expected := args.Map{
		"parentExist": true, "combine": true, "combineWith": true,
		"parentWith": true, "parentNew": true,
	}
	expected.ShouldBeEqual(t, 0, "PathExistStat returns correct value -- Parent", actual)
}

func Test_PathExistStat_NotAFileError_Dir(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat(t.TempDir())
	err := stat.NotAFileError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAFileError returns error -- dir", actual)
}

func Test_PathExistStat_NotADirError_File(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "f.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	stat := chmodhelper.GetPathExistStat(fp)
	err := stat.NotADirError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotADirError returns error -- file", actual)
}

// ── SimpleFileReaderWriter — uncovered branches ──

func Test_SimpleFileRW_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, ParentDir: "/tmp", FilePath: "/tmp/test.txt"}
	b, err := rw.MarshalJSON()
	var rw2 chmodhelper.SimpleFileReaderWriter
	err2 := rw2.UnmarshalJSON(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"noErr2": err2 == nil,
		"path": rw2.FilePath,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"noErr2": true,
		"path": "/tmp/test.txt",
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- MarshalUnmarshalJSON", actual)
}

func Test_SimpleFileRW_Clone(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	c := rw.Clone()

	// Act
	actual := args.Map{"path": c.FilePath}

	// Assert
	expected := args.Map{"path": "/tmp/test.txt"}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- Clone", actual)
}

func Test_SimpleFileRW_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var rw *chmodhelper.SimpleFileReaderWriter

	// Act
	actual := args.Map{"nil": rw.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns nil -- ClonePtr nil", actual)
}

func Test_SimpleFileRW_Json(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}
	r := rw.Json()
	rp := rw.JsonPtr()

	// Act
	actual := args.Map{
		"noErr": !r.HasError(),
		"ptrNotNil": rp != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- Json", actual)
}

func Test_SimpleFileRW_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644}

	// Act
	actual := args.Map{"notNil": rw.AsJsonContractsBinder() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_SimpleFileRW_String(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ChmodDir: 0755, ChmodFile: 0644, FilePath: "/tmp/test.txt"}

	// Act
	actual := args.Map{"notEmpty": rw.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SimpleFileRW returns correct value -- String", actual)
}

func Test_SimpleFileRW_Expire(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "expire.txt")
	_ = os.WriteFile(fp, []byte("x"), 0644)
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: fp}
	err := rw.Expire()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Expire returns correct value -- with args", actual)
}

func Test_SimpleFileRW_Expire_NonExist(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent_cov5"}
	err := rw.Expire()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Expire returns non-empty -- non-exist", actual)
}

func Test_SimpleFileRW_RemoveOnExist(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{FilePath: "/nonexistent_cov5"}
	err := rw.RemoveOnExist()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveOnExist returns correct value -- with args", actual)
}

func Test_SimpleFileRW_RemoveDirOnExist(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{ParentDir: "/nonexistent_cov5_dir"}
	err := rw.RemoveDirOnExist()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirOnExist returns correct value -- with args", actual)
}

// ── newAttributeCreator — uncovered branches ──

func Test_NewAttribute_Create(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, false, true)

	// Act
	actual := args.Map{
		"read": attr.IsRead,
		"write": attr.IsWrite,
		"exe": attr.IsExecute,
	}

	// Assert
	expected := args.Map{
		"read": true,
		"write": false,
		"exe": true,
	}
	expected.ShouldBeEqual(t, 0, "NewAttribute returns correct value -- Create", actual)
}

func Test_NewAttribute_Default(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Default(false, true, false)

	// Act
	actual := args.Map{"write": attr.IsWrite}

	// Assert
	expected := args.Map{"write": true}
	expected.ShouldBeEqual(t, 0, "NewAttribute returns correct value -- Default", actual)
}

func Test_NewAttribute_UsingByte(t *testing.T) {
	// Arrange
	attr, err := chmodhelper.New.Attribute.UsingByte(5)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"read": attr.IsRead,
		"exe": attr.IsExecute,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"read": true,
		"exe": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingByte returns correct value -- with args", actual)
}

func Test_NewAttribute_UsingVariant(t *testing.T) {
	// Arrange
	attr, err := chmodhelper.New.Attribute.UsingVariant(chmodhelper.ReadWrite)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"read": attr.IsRead,
		"write": attr.IsWrite,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"read": true,
		"write": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingVariant returns correct value -- with args", actual)
}

// ── VarAttribute — uncovered branches ──

func Test_VarAttribute_ToCompileFixAttr_NonFixed(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwx***r-x")
	result := vw.Group.ToCompileFixAttr()

	// Act
	actual := args.Map{"nil": result == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ToCompileFixAttr returns non-empty -- non-fixed", actual)
}

func Test_VarAttribute_ToCompileAttr_Wildcard(t *testing.T) {
	// Arrange
	vw, _ := chmodhelper.NewRwxVariableWrapper("-rwx***r-x")
	fixed := &chmodhelper.Attribute{IsRead: true, IsWrite: true, IsExecute: true}
	attr := vw.Group.ToCompileAttr(fixed)

	// Act
	actual := args.Map{
		"read": attr.IsRead,
		"write": attr.IsWrite,
		"exe": attr.IsExecute,
	}

	// Assert
	expected := args.Map{
		"read": true,
		"write": true,
		"exe": true,
	}
	expected.ShouldBeEqual(t, 0, "ToCompileAttr returns correct value -- wildcard", actual)
}

// ── newSimpleFileReaderWriterCreator ──

func Test_NewSimpleFileRW_Default(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(false, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW returns correct value -- Default", actual)
}

func Test_NewSimpleFileRW_Path(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(false, 0755, 0644, "/tmp/test.txt")

	// Act
	actual := args.Map{"notNil": rw != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewSimpleFileRW returns correct value -- Path", actual)
}

// ── RwxInstructionExecutors ──

func Test_RwxInstructionExecutors_Basic(t *testing.T) {
	// Arrange
	execs := chmodhelper.NewRwxInstructionExecutors(2)

	// Act
	actual := args.Map{
		"empty":    execs.IsEmpty(),
		"hasAny":   execs.HasAnyItem(),
		"len":      execs.Length(),
		"count":    execs.Count(),
		"lastIdx":  execs.LastIndex(),
		"hasIdx0":  execs.HasIndex(0),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"hasAny": false,
		"len": 0,
		"count": 0,
		"lastIdx": -1,
		"hasIdx0": false,
	}
	expected.ShouldBeEqual(t, 0, "RwxInstructionExecutors returns correct value -- basic", actual)
}

func Test_RwxInstructionExecutors_Add_Nil(t *testing.T) {
	// Arrange
	execs := chmodhelper.NewRwxInstructionExecutors(2)
	execs.Add(nil)

	// Act
	actual := args.Map{"len": execs.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Add returns nil -- nil", actual)
}

func Test_RwxInstructionExecutors_Items(t *testing.T) {
	// Arrange
	execs := chmodhelper.NewRwxInstructionExecutors(2)

	// Act
	actual := args.Map{"notNil": execs.Items() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Items returns correct value -- with args", actual)
}
