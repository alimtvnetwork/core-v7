package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/enums/versionindexes"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// newCreator — uncovered methods
// =============================================================================

func Test_New_Version(t *testing.T) {
	// Arrange
	v := coreversion.New.Version("1.2.3")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_Major(t *testing.T) {
	// Arrange
	v := coreversion.New.Major("5")

	// Act
	actual := args.Map{"result": v.HasMajor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
}

func Test_New_DefaultPtr(t *testing.T) {
	// Arrange
	v := coreversion.New.DefaultPtr("1.2.3")

	// Act
	actual := args.Map{"result": v == nil || v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_SpreadStrings(t *testing.T) {
	// Arrange
	v := coreversion.New.SpreadStrings("1", "2", "3")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_SpreadIntegers(t *testing.T) {
	// Arrange
	v := coreversion.New.SpreadIntegers(1, 2, 3)

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_SpreadUnsignedIntegers(t *testing.T) {
	// Arrange
	v := coreversion.New.SpreadUnsignedIntegers(1, 2, 3)

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_SpreadBytes(t *testing.T) {
	// Arrange
	v := coreversion.New.SpreadBytes(1, 2, 3)

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_MajorMinor(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorMinor("1", "2")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_MajorMinorPatch(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorMinorPatch("1", "2", "3")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_MajorMinorPatchBuild(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorMinorPatchBuild("1", "2", "3", "4")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_All(t *testing.T) {
	// Arrange
	v := coreversion.New.All("1", "2", "3", "4")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_AllInt(t *testing.T) {
	// Arrange
	v := coreversion.New.AllInt(1, 2, 3, 4)

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_AllByte(t *testing.T) {
	// Arrange
	v := coreversion.New.AllByte(1, 2, 3, 4)

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_MajorMinorInt(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorMinorInt(1, 2)

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_MajorMinorPatchInt(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorMinorPatchInt(1, 2, 3)

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_New_MajorBuildInt(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorBuildInt(1, 5)

	// Act
	actual := args.Map{"result": v.HasMajor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
}

func Test_New_MajorBuild(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorBuild("1", "5")

	// Act
	actual := args.Map{"result": v.HasMajor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
}

func Test_New_MajorMinorBuild(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorMinorBuild("1", "2", "5")

	// Act
	actual := args.Map{"result": v.HasMajor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
}

func Test_New_MajorPatch(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorPatch("1", "3")

	// Act
	actual := args.Map{"result": v.HasMajor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
}

func Test_New_MajorPatchInt(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorPatchInt(1, 3)

	// Act
	actual := args.Map{"result": v.HasMajor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
}

func Test_New_MinorBuildInt(t *testing.T) {
	v := coreversion.New.MinorBuildInt(2, 5)
	_ = v
}

func Test_New_PatchBuildInt(t *testing.T) {
	v := coreversion.New.PatchBuildInt(3, 5)
	_ = v
}

func Test_New_MinorBuild(t *testing.T) {
	v := coreversion.New.MinorBuild("2", "5")
	_ = v
}

func Test_New_PatchBuild(t *testing.T) {
	v := coreversion.New.PatchBuild("3", "5")
	_ = v
}

func Test_New_Many(t *testing.T) {
	// Arrange
	vc := coreversion.New.Many("1.0.0", "2.0.0")

	// Act
	actual := args.Map{"result": vc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2 versions", actual)
}

func Test_New_Collection(t *testing.T) {
	// Arrange
	vc := coreversion.New.Collection("1.0.0")

	// Act
	actual := args.Map{"result": vc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1", actual)
}

func Test_New_CollectionUsingCap(t *testing.T) {
	// Arrange
	vc := coreversion.New.CollectionUsingCap(10)

	// Act
	actual := args.Map{"result": vc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_New_EmptyCollection(t *testing.T) {
	// Arrange
	vc := coreversion.New.EmptyCollection()

	// Act
	actual := args.Map{"result": vc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_New_Invalid(t *testing.T) {
	// Arrange
	v := coreversion.New.Invalid()

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_New_Empty(t *testing.T) {
	// Arrange
	v := coreversion.New.Empty()

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty/invalid", actual)
}

// =============================================================================
// Version — uncovered methods
// =============================================================================

func Test_Version_IsLeftLessThan(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("2.0.0")

	// Act
	actual := args.Map{"result": v1.IsLeftLessThan(&v2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.0.0 < 2.0.0", actual)
}

func Test_Version_IsLeftGreaterThan(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("2.0.0")
	v2 := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v1.IsLeftGreaterThan(&v2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2.0.0 > 1.0.0", actual)
}

func Test_Version_IsLeftLessThanOrEqual(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v1.IsLeftLessThanOrEqual(&v2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.0.0 <= 1.0.0", actual)
}

func Test_Version_IsLeftGreaterThanOrEqual(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("2.0.0")
	v2 := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v1.IsLeftGreaterThanOrEqual(&v2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2.0.0 >= 1.0.0", actual)
}

func Test_Version_IsExpectedComparison(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v1.IsExpectedComparison(corecomparator.Equal, &v2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_IsExpectedComparisonRawVersion(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v.IsExpectedComparisonRawVersion(corecomparator.Equal, "1.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_IsAtLeast(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("2.0.0")

	// Act
	actual := args.Map{"result": v.IsAtLeast("1.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2.0.0 >= 1.0.0", actual)
}

func Test_Version_IsEqualVersionString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v.IsEqualVersionString("1.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_IsLowerVersionString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v.IsLowerVersionString("2.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.0.0 < 2.0.0", actual)
}

func Test_Version_IsLowerEqualVersionString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v.IsLowerEqualVersionString("1.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.0.0 <= 1.0.0", actual)
}

func Test_Version_ComparisonValueIndexes(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("1.2.3")
	v2 := coreversion.New.Create("1.2.3")
	cmp := v1.ComparisonValueIndexes(&v2, versionindexes.Major, versionindexes.Minor)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_ComparisonValueIndexes_NilRight(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("1.2.3")
	cmp := v1.ComparisonValueIndexes(nil, versionindexes.Major)

	// Act
	actual := args.Map{"result": cmp.IsLeftGreater()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "non-nil > nil", actual)
}

func Test_Version_Clone(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	c := v.Clone()

	// Act
	actual := args.Map{"result": c.VersionCompact != "1.2.3"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_Version_ClonePtr(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	c := v.ClonePtr()

	// Act
	actual := args.Map{"result": c == nil || c.VersionCompact != "1.2.3"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone ptr", actual)
}

func Test_Version_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.ClonePtr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Version_NonPtr(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	np := v.NonPtr()

	// Act
	actual := args.Map{"result": np.VersionCompact != "1.2.3"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return value", actual)
}

func Test_Version_Ptr(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	p := v.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return pointer", actual)
}

func Test_Version_Json(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	j := v.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error", actual)
}

func Test_Version_JsonPtr(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	j := v.JsonPtr()

	// Act
	actual := args.Map{"result": j == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Version_JsonParseSelfInject(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	j := v.JsonPtr()
	v2 := &coreversion.Version{}
	err := v2.JsonParseSelfInject(j)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not error:", actual)
}

func Test_Version_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.AsJsonContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return self", actual)
}

func Test_Version_ValueByIndex_Invalid(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	val := v.ValueByIndex(versionindexes.Index(99))

	// Act
	actual := args.Map{"result": val != coreversion.InvalidVersionValue}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "invalid index should return invalid value", actual)
}

func Test_Version_ValueByIndexes(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")
	vals := v.ValueByIndexes(versionindexes.Major, versionindexes.Build)

	// Act
	actual := args.Map{"result": len(vals) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 2 values", actual)
}

func Test_Version_IsSafeInvalidCheck(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("")

	// Act
	actual := args.Map{"result": v.IsSafeInvalidCheck()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be invalid", actual)
}

func Test_Version_IsInvalidOrEmpty(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("")

	// Act
	actual := args.Map{"result": v.IsInvalidOrEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be invalid", actual)
}

func Test_Version_VersionDisplayMajor_Invalid(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("")

	// Act
	actual := args.Map{"result": v.VersionDisplayMajor() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "invalid should return empty", actual)
}

func Test_Version_VersionDisplayMajorMinor_MinorInvalid(t *testing.T) {
	// Create a version with only major
	v := coreversion.New.Create("1")
	// VersionMinor should be 0, which is <= InvalidVersionValue
	d := v.VersionDisplayMajorMinor()
	_ = d // just exercise the path
}

func Test_Version_VersionDisplayMajorMinorPatch_PatchInvalid(t *testing.T) {
	v := coreversion.New.Create("1.2")
	d := v.VersionDisplayMajorMinorPatch()
	_ = d // exercise the fallback to MajorMinor
}

// =============================================================================
// Empty / EmptyUsingCompactVersion / InvalidCompactVersion
// =============================================================================

func Test_Empty(t *testing.T) {
	// Arrange
	v := coreversion.Empty()

	// Act
	actual := args.Map{"result": v.IsInvalid}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_EmptyUsingCompactVersion(t *testing.T) {
	// Arrange
	v := coreversion.EmptyUsingCompactVersion("1.0.0")

	// Act
	actual := args.Map{"result": v.IsInvalid || v.VersionCompact != "1.0.0"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be invalid", actual)
}

func Test_InvalidCompactVersion(t *testing.T) {
	// Arrange
	v := coreversion.InvalidCompactVersion("bad")

	// Act
	actual := args.Map{"result": v.IsInvalid || v.VersionCompact != "bad"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid with compact", actual)
}

// =============================================================================
// VersionsCollection — uncovered methods
// =============================================================================

func Test_VersionsCollection_AddVersions(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("2.0.0")
	vc.AddVersions(v1, v2)

	// Act
	actual := args.Map{"result": vc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2", actual)
}

func Test_VersionsCollection_IsEqual_DifferentVersions(t *testing.T) {
	// Arrange
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("2.0.0")

	// Act
	actual := args.Map{"result": vc1.IsEqual(vc2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different versions should not be equal", actual)
}

func Test_VersionsCollection_Json(t *testing.T) {
	// Arrange
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not error", actual)
}

func Test_VersionsCollection_JsonPtr(t *testing.T) {
	// Arrange
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.JsonPtr()

	// Act
	actual := args.Map{"result": j == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_VersionsCollection_JsonParseSelfInject(t *testing.T) {
	// Arrange
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.JsonPtr()
	vc2 := &coreversion.VersionsCollection{}
	err := vc2.JsonParseSelfInject(j)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not error:", actual)
}

func Test_VersionsCollection_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}

	// Act
	actual := args.Map{"result": vc.AsJsonContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return self", actual)
}

func Test_VersionsCollection_AsBasicSliceContractsBinder(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}

	// Act
	actual := args.Map{"result": vc.AsBasicSliceContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return self", actual)
}

// =============================================================================
// Default() edge cases
// =============================================================================

func Test_New_Default_Empty(t *testing.T) {
	// Arrange
	v := coreversion.New.Default("")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be invalid", actual)
}

func Test_New_Default_Whitespace(t *testing.T) {
	// Arrange
	v := coreversion.New.Default("   ")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "whitespace should be invalid", actual)
}

func Test_New_Default_JustV(t *testing.T) {
	// Arrange
	v := coreversion.New.Default("v")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "just v should be invalid", actual)
}

func Test_New_Default_WithBuild(t *testing.T) {
	// Arrange
	v := coreversion.New.Default("1.2.3.4")

	// Act
	actual := args.Map{"result": v.VersionBuild != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected build 4", actual)
}

// =============================================================================
// hasDeductUsingNilNess — all branches
// =============================================================================

func Test_Compare_BothNil(t *testing.T) {
	// Arrange
	cmp := coreversion.Compare(nil, nil)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_Compare_LeftNil(t *testing.T) {
	// Arrange
	v := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(nil, v)

	// Act
	actual := args.Map{"result": cmp.IsLeftLess()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "left nil should be less", actual)
}

func Test_Compare_RightNil(t *testing.T) {
	// Arrange
	v := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(v, nil)

	// Act
	actual := args.Map{"result": cmp.IsLeftGreater()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "right nil should make left greater", actual)
}

func Test_Compare_SamePtr(t *testing.T) {
	// Arrange
	v := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(v, v)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same ptr should be equal", actual)
}

func Test_Compare_SameCompact(t *testing.T) {
	// Arrange
	v1 := coreversion.New.DefaultPtr("1.0.0")
	v2 := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(v1, v2)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same compact should be equal", actual)
}

func Test_Compare_DifferentMinor(t *testing.T) {
	// Arrange
	v1 := coreversion.New.DefaultPtr("1.1.0")
	v2 := coreversion.New.DefaultPtr("1.2.0")
	cmp := coreversion.Compare(v1, v2)

	// Act
	actual := args.Map{"result": cmp.IsLeftLess()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.1.0 < 1.2.0", actual)
}

func Test_Compare_DifferentPatch(t *testing.T) {
	// Arrange
	v1 := coreversion.New.DefaultPtr("1.2.1")
	v2 := coreversion.New.DefaultPtr("1.2.3")
	cmp := coreversion.Compare(v1, v2)

	// Act
	actual := args.Map{"result": cmp.IsLeftLess()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.2.1 < 1.2.3", actual)
}

func Test_Compare_DifferentBuild(t *testing.T) {
	// Arrange
	v1 := coreversion.New.DefaultPtr("1.2.3.1")
	v2 := coreversion.New.DefaultPtr("1.2.3.5")
	cmp := coreversion.Compare(v1, v2)

	// Act
	actual := args.Map{"result": cmp.IsLeftLess()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.2.3.1 < 1.2.3.5", actual)
}
