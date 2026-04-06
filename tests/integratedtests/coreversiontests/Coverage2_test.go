package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

// =============================================================================
// newCreator — uncovered methods
// =============================================================================

func Test_New_Version(t *testing.T) {
	v := coreversion.New.Version("1.2.3")
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_Major(t *testing.T) {
	v := coreversion.New.Major("5")
	if !v.HasMajor() {
		t.Error("should have major")
	}
}

func Test_New_DefaultPtr(t *testing.T) {
	v := coreversion.New.DefaultPtr("1.2.3")
	if v == nil || v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_SpreadStrings(t *testing.T) {
	v := coreversion.New.SpreadStrings("1", "2", "3")
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_SpreadIntegers(t *testing.T) {
	v := coreversion.New.SpreadIntegers(1, 2, 3)
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_SpreadUnsignedIntegers(t *testing.T) {
	v := coreversion.New.SpreadUnsignedIntegers(1, 2, 3)
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_SpreadBytes(t *testing.T) {
	v := coreversion.New.SpreadBytes(1, 2, 3)
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_MajorMinor(t *testing.T) {
	v := coreversion.New.MajorMinor("1", "2")
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_MajorMinorPatch(t *testing.T) {
	v := coreversion.New.MajorMinorPatch("1", "2", "3")
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_MajorMinorPatchBuild(t *testing.T) {
	v := coreversion.New.MajorMinorPatchBuild("1", "2", "3", "4")
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_All(t *testing.T) {
	v := coreversion.New.All("1", "2", "3", "4")
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_AllInt(t *testing.T) {
	v := coreversion.New.AllInt(1, 2, 3, 4)
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_AllByte(t *testing.T) {
	v := coreversion.New.AllByte(1, 2, 3, 4)
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_MajorMinorInt(t *testing.T) {
	v := coreversion.New.MajorMinorInt(1, 2)
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_MajorMinorPatchInt(t *testing.T) {
	v := coreversion.New.MajorMinorPatchInt(1, 2, 3)
	if v.IsEmptyOrInvalid() {
		t.Error("should be valid")
	}
}

func Test_New_MajorBuildInt(t *testing.T) {
	v := coreversion.New.MajorBuildInt(1, 5)
	if !v.HasMajor() {
		t.Error("should have major")
	}
}

func Test_New_MajorBuild(t *testing.T) {
	v := coreversion.New.MajorBuild("1", "5")
	if !v.HasMajor() {
		t.Error("should have major")
	}
}

func Test_New_MajorMinorBuild(t *testing.T) {
	v := coreversion.New.MajorMinorBuild("1", "2", "5")
	if !v.HasMajor() {
		t.Error("should have major")
	}
}

func Test_New_MajorPatch(t *testing.T) {
	v := coreversion.New.MajorPatch("1", "3")
	if !v.HasMajor() {
		t.Error("should have major")
	}
}

func Test_New_MajorPatchInt(t *testing.T) {
	v := coreversion.New.MajorPatchInt(1, 3)
	if !v.HasMajor() {
		t.Error("should have major")
	}
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
	vc := coreversion.New.Many("1.0.0", "2.0.0")
	if vc.Length() != 2 {
		t.Error("should have 2 versions")
	}
}

func Test_New_Collection(t *testing.T) {
	vc := coreversion.New.Collection("1.0.0")
	if vc.Length() != 1 {
		t.Error("should have 1")
	}
}

func Test_New_CollectionUsingCap(t *testing.T) {
	vc := coreversion.New.CollectionUsingCap(10)
	if vc == nil {
		t.Error("should not be nil")
	}
}

func Test_New_EmptyCollection(t *testing.T) {
	vc := coreversion.New.EmptyCollection()
	if vc.Length() != 0 {
		t.Error("should be empty")
	}
}

func Test_New_Invalid(t *testing.T) {
	v := coreversion.New.Invalid()
	if !v.IsEmptyOrInvalid() {
		t.Error("should be invalid")
	}
}

func Test_New_Empty(t *testing.T) {
	v := coreversion.New.Empty()
	if !v.IsEmptyOrInvalid() {
		t.Error("should be empty/invalid")
	}
}

// =============================================================================
// Version — uncovered methods
// =============================================================================

func Test_Version_IsLeftLessThan(t *testing.T) {
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("2.0.0")
	if !v1.IsLeftLessThan(&v2) {
		t.Error("1.0.0 < 2.0.0")
	}
}

func Test_Version_IsLeftGreaterThan(t *testing.T) {
	v1 := coreversion.New.Create("2.0.0")
	v2 := coreversion.New.Create("1.0.0")
	if !v1.IsLeftGreaterThan(&v2) {
		t.Error("2.0.0 > 1.0.0")
	}
}

func Test_Version_IsLeftLessThanOrEqual(t *testing.T) {
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("1.0.0")
	if !v1.IsLeftLessThanOrEqual(&v2) {
		t.Error("1.0.0 <= 1.0.0")
	}
}

func Test_Version_IsLeftGreaterThanOrEqual(t *testing.T) {
	v1 := coreversion.New.Create("2.0.0")
	v2 := coreversion.New.Create("1.0.0")
	if !v1.IsLeftGreaterThanOrEqual(&v2) {
		t.Error("2.0.0 >= 1.0.0")
	}
}

func Test_Version_IsExpectedComparison(t *testing.T) {
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("1.0.0")
	if !v1.IsExpectedComparison(corecomparator.Equal, &v2) {
		t.Error("should be equal")
	}
}

func Test_Version_IsExpectedComparisonRawVersion(t *testing.T) {
	v := coreversion.New.Create("1.0.0")
	if !v.IsExpectedComparisonRawVersion(corecomparator.Equal, "1.0.0") {
		t.Error("should be equal")
	}
}

func Test_Version_IsAtLeast(t *testing.T) {
	v := coreversion.New.Create("2.0.0")
	if !v.IsAtLeast("1.0.0") {
		t.Error("2.0.0 >= 1.0.0")
	}
}

func Test_Version_IsEqualVersionString(t *testing.T) {
	v := coreversion.New.Create("1.0.0")
	if !v.IsEqualVersionString("1.0.0") {
		t.Error("should be equal")
	}
}

func Test_Version_IsLowerVersionString(t *testing.T) {
	v := coreversion.New.Create("1.0.0")
	if !v.IsLowerVersionString("2.0.0") {
		t.Error("1.0.0 < 2.0.0")
	}
}

func Test_Version_IsLowerEqualVersionString(t *testing.T) {
	v := coreversion.New.Create("1.0.0")
	if !v.IsLowerEqualVersionString("1.0.0") {
		t.Error("1.0.0 <= 1.0.0")
	}
}

func Test_Version_ComparisonValueIndexes(t *testing.T) {
	v1 := coreversion.New.Create("1.2.3")
	v2 := coreversion.New.Create("1.2.3")
	cmp := v1.ComparisonValueIndexes(&v2, versionindexes.Major, versionindexes.Minor)
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_ComparisonValueIndexes_NilRight(t *testing.T) {
	v1 := coreversion.New.Create("1.2.3")
	cmp := v1.ComparisonValueIndexes(nil, versionindexes.Major)
	if !cmp.IsLeftGreater() {
		t.Error("non-nil > nil")
	}
}

func Test_Version_Clone(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	c := v.Clone()
	if c.VersionCompact != "1.2.3" {
		t.Error("should clone")
	}
}

func Test_Version_ClonePtr(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	c := v.ClonePtr()
	if c == nil || c.VersionCompact != "1.2.3" {
		t.Error("should clone ptr")
	}
}

func Test_Version_ClonePtr_Nil(t *testing.T) {
	var v *coreversion.Version
	if v.ClonePtr() != nil {
		t.Error("nil should return nil")
	}
}

func Test_Version_NonPtr(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	np := v.NonPtr()
	if np.VersionCompact != "1.2.3" {
		t.Error("should return value")
	}
}

func Test_Version_Ptr(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	p := v.Ptr()
	if p == nil {
		t.Error("should return pointer")
	}
}

func Test_Version_Json(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	j := v.Json()
	if j.HasError() {
		t.Error("should not have error")
	}
}

func Test_Version_JsonPtr(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	j := v.JsonPtr()
	if j == nil {
		t.Error("should not be nil")
	}
}

func Test_Version_JsonParseSelfInject(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	j := v.JsonPtr()
	v2 := &coreversion.Version{}
	err := v2.JsonParseSelfInject(j)
	if err != nil {
		t.Errorf("should not error: %v", err)
	}
}

func Test_Version_AsJsonContractsBinder(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if v.AsJsonContractsBinder() == nil {
		t.Error("should return self")
	}
}

func Test_Version_ValueByIndex_Invalid(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	val := v.ValueByIndex(versionindexes.Index(99))
	if val != coreversion.InvalidVersionValue {
		t.Error("invalid index should return invalid value")
	}
}

func Test_Version_ValueByIndexes(t *testing.T) {
	v := coreversion.New.Create("1.2.3.4")
	vals := v.ValueByIndexes(versionindexes.Major, versionindexes.Build)
	if len(vals) != 2 {
		t.Error("should return 2 values")
	}
}

func Test_Version_IsSafeInvalidCheck(t *testing.T) {
	v := coreversion.New.Create("")
	if !v.IsSafeInvalidCheck() {
		t.Error("empty should be invalid")
	}
}

func Test_Version_IsInvalidOrEmpty(t *testing.T) {
	v := coreversion.New.Create("")
	if !v.IsInvalidOrEmpty() {
		t.Error("empty should be invalid")
	}
}

func Test_Version_VersionDisplayMajor_Invalid(t *testing.T) {
	v := coreversion.New.Create("")
	if v.VersionDisplayMajor() != "" {
		t.Error("invalid should return empty")
	}
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
	v := coreversion.Empty()
	if !v.IsInvalid {
		t.Error("should be invalid")
	}
}

func Test_EmptyUsingCompactVersion(t *testing.T) {
	v := coreversion.EmptyUsingCompactVersion("1.0.0")
	if v.IsInvalid || v.VersionCompact != "1.0.0" {
		t.Error("should not be invalid")
	}
}

func Test_InvalidCompactVersion(t *testing.T) {
	v := coreversion.InvalidCompactVersion("bad")
	if !v.IsInvalid || v.VersionCompact != "bad" {
		t.Error("should be invalid with compact")
	}
}

// =============================================================================
// VersionsCollection — uncovered methods
// =============================================================================

func Test_VersionsCollection_AddVersions(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	v1 := coreversion.New.Create("1.0.0")
	v2 := coreversion.New.Create("2.0.0")
	vc.AddVersions(v1, v2)
	if vc.Length() != 2 {
		t.Error("should have 2")
	}
}

func Test_VersionsCollection_IsEqual_DifferentVersions(t *testing.T) {
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("2.0.0")
	if vc1.IsEqual(vc2) {
		t.Error("different versions should not be equal")
	}
}

func Test_VersionsCollection_Json(t *testing.T) {
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.Json()
	if j.HasError() {
		t.Error("should not error")
	}
}

func Test_VersionsCollection_JsonPtr(t *testing.T) {
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.JsonPtr()
	if j == nil {
		t.Error("should not be nil")
	}
}

func Test_VersionsCollection_JsonParseSelfInject(t *testing.T) {
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.JsonPtr()
	vc2 := &coreversion.VersionsCollection{}
	err := vc2.JsonParseSelfInject(j)
	if err != nil {
		t.Errorf("should not error: %v", err)
	}
}

func Test_VersionsCollection_AsJsonContractsBinder(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	if vc.AsJsonContractsBinder() == nil {
		t.Error("should return self")
	}
}

func Test_VersionsCollection_AsBasicSliceContractsBinder(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	if vc.AsBasicSliceContractsBinder() == nil {
		t.Error("should return self")
	}
}

// =============================================================================
// Default() edge cases
// =============================================================================

func Test_New_Default_Empty(t *testing.T) {
	v := coreversion.New.Default("")
	if !v.IsEmptyOrInvalid() {
		t.Error("empty should be invalid")
	}
}

func Test_New_Default_Whitespace(t *testing.T) {
	v := coreversion.New.Default("   ")
	if !v.IsEmptyOrInvalid() {
		t.Error("whitespace should be invalid")
	}
}

func Test_New_Default_JustV(t *testing.T) {
	v := coreversion.New.Default("v")
	if !v.IsEmptyOrInvalid() {
		t.Error("just v should be invalid")
	}
}

func Test_New_Default_WithBuild(t *testing.T) {
	v := coreversion.New.Default("1.2.3.4")
	if v.VersionBuild != 4 {
		t.Errorf("expected build 4, got %d", v.VersionBuild)
	}
}

// =============================================================================
// hasDeductUsingNilNess — all branches
// =============================================================================

func Test_Compare_BothNil(t *testing.T) {
	cmp := coreversion.Compare(nil, nil)
	if !cmp.IsEqual() {
		t.Error("both nil should be equal")
	}
}

func Test_Compare_LeftNil(t *testing.T) {
	v := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(nil, v)
	if !cmp.IsLeftLess() {
		t.Error("left nil should be less")
	}
}

func Test_Compare_RightNil(t *testing.T) {
	v := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(v, nil)
	if !cmp.IsLeftGreater() {
		t.Error("right nil should make left greater")
	}
}

func Test_Compare_SamePtr(t *testing.T) {
	v := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(v, v)
	if !cmp.IsEqual() {
		t.Error("same ptr should be equal")
	}
}

func Test_Compare_SameCompact(t *testing.T) {
	v1 := coreversion.New.DefaultPtr("1.0.0")
	v2 := coreversion.New.DefaultPtr("1.0.0")
	cmp := coreversion.Compare(v1, v2)
	if !cmp.IsEqual() {
		t.Error("same compact should be equal")
	}
}

func Test_Compare_DifferentMinor(t *testing.T) {
	v1 := coreversion.New.DefaultPtr("1.1.0")
	v2 := coreversion.New.DefaultPtr("1.2.0")
	cmp := coreversion.Compare(v1, v2)
	if !cmp.IsLeftLess() {
		t.Error("1.1.0 < 1.2.0")
	}
}

func Test_Compare_DifferentPatch(t *testing.T) {
	v1 := coreversion.New.DefaultPtr("1.2.1")
	v2 := coreversion.New.DefaultPtr("1.2.3")
	cmp := coreversion.Compare(v1, v2)
	if !cmp.IsLeftLess() {
		t.Error("1.2.1 < 1.2.3")
	}
}

func Test_Compare_DifferentBuild(t *testing.T) {
	v1 := coreversion.New.DefaultPtr("1.2.3.1")
	v2 := coreversion.New.DefaultPtr("1.2.3.5")
	cmp := coreversion.Compare(v1, v2)
	if !cmp.IsLeftLess() {
		t.Error("1.2.3.1 < 1.2.3.5")
	}
}
