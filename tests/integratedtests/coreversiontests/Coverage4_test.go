package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

// ==========================================================================
// VersionsCollection — full coverage
// ==========================================================================

func Test_Cov4_VersionsCollection_Add(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	actual := args.Map{"len": vc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Add returns correct value -- version to collection", actual)
}

func Test_Cov4_VersionsCollection_AddSkipInvalid(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.AddSkipInvalid("1.0.0")
	vc.AddSkipInvalid("")       // skipped
	vc.AddSkipInvalid("v")      // skipped
	actual := args.Map{"len": vc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddSkipInvalid returns empty -- skips empty", actual)
}

func Test_Cov4_VersionsCollection_AddVersionsRaw(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.AddVersionsRaw("1.0", "2.0")
	actual := args.Map{"len": vc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddVersionsRaw returns correct value -- adds multiple", actual)
}

func Test_Cov4_VersionsCollection_AddVersions(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	v := coreversion.New.Create("1.0.0")
	vc.AddVersions(v)
	actual := args.Map{"len": vc.Count()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddVersions returns correct value -- adds version struct", actual)
}

func Test_Cov4_VersionsCollection_IsEmpty_HasAnyItem(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	actual := args.Map{"empty": vc.IsEmpty(), "hasAny": vc.HasAnyItem()}
	expected := args.Map{"empty": true, "hasAny": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- and HasAnyItem on empty", actual)
}

func Test_Cov4_VersionsCollection_LastIndex_HasIndex(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0").Add("2.0")
	actual := args.Map{"lastIdx": vc.LastIndex(), "hasIdx0": vc.HasIndex(0), "hasIdx5": vc.HasIndex(5)}
	expected := args.Map{"lastIdx": 1, "hasIdx0": true, "hasIdx5": false}
	expected.ShouldBeEqual(t, 0, "LastIndex returns correct value -- and HasIndex", actual)
}

func Test_Cov4_VersionsCollection_VersionCompactStrings(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	strs := vc.VersionCompactStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VersionCompactStrings returns correct value -- returns strings", actual)
}

func Test_Cov4_VersionsCollection_VersionCompactStrings_Empty(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	strs := vc.VersionCompactStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VersionCompactStrings returns empty -- empty", actual)
}

func Test_Cov4_VersionsCollection_VersionsStrings(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	strs := vc.VersionsStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VersionsStrings returns correct value -- returns display strings", actual)
}

func Test_Cov4_VersionsCollection_VersionsStrings_Empty(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	strs := vc.VersionsStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VersionsStrings returns empty -- empty", actual)
}

func Test_Cov4_VersionsCollection_IndexOf(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0").Add("2.0.0")
	actual := args.Map{"found": vc.IndexOf("1.0.0") >= 0, "notFound": vc.IndexOf("3.0.0") < 0}
	expected := args.Map{"found": true, "notFound": true}
	expected.ShouldBeEqual(t, 0, "IndexOf returns correct value -- finds version", actual)
}

func Test_Cov4_VersionsCollection_IsContainsVersion(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	actual := args.Map{"contains": vc.IsContainsVersion("1.0.0"), "missing": vc.IsContainsVersion("9.9")}
	expected := args.Map{"contains": true, "missing": false}
	expected.ShouldBeEqual(t, 0, "IsContainsVersion returns correct value -- with args", actual)
}

func Test_Cov4_VersionsCollection_IsEqual(t *testing.T) {
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("1.0.0")
	vc3 := &coreversion.VersionsCollection{}
	vc3.Add("2.0.0")
	actual := args.Map{
		"eq":      vc1.IsEqual(vc2),
		"neq":     vc1.IsEqual(vc3),
		"nilNil":  (*coreversion.VersionsCollection)(nil).IsEqual(nil),
		"nilR":    vc1.IsEqual(nil),
	}
	expected := args.Map{
		"eq":      true,
		"neq":     false,
		"nilNil":  true,
		"nilR":    false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- comparisons", actual)
}

func Test_Cov4_VersionsCollection_IsEqual_DiffLen(t *testing.T) {
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("1.0.0").Add("2.0.0")
	actual := args.Map{"eq": vc1.IsEqual(vc2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- diff lengths", actual)
}

func Test_Cov4_VersionsCollection_String(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	s := vc.String()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- returns display", actual)
}

func Test_Cov4_VersionsCollection_Json(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.Json()
	jp := vc.JsonPtr()
	actual := args.Map{"hasResult": j.HasSafeItems(), "ptrNotNil": jp != nil}
	expected := args.Map{"hasResult": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- and JsonPtr", actual)
}

func Test_Cov4_VersionsCollection_Length_Nil(t *testing.T) {
	var vc *coreversion.VersionsCollection
	actual := args.Map{"len": vc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- on nil returns 0", actual)
}

// ==========================================================================
// EmptyUsingCompactVersion + InvalidCompactVersion
// ==========================================================================

func Test_Cov4_EmptyUsingCompactVersion(t *testing.T) {
	v := coreversion.EmptyUsingCompactVersion("1.2.3")
	actual := args.Map{"compact": v.VersionCompact, "invalid": v.IsInvalid}
	expected := args.Map{"compact": "1.2.3", "invalid": false}
	expected.ShouldBeEqual(t, 0, "EmptyUsingCompactVersion returns empty -- with args", actual)
}

func Test_Cov4_InvalidCompactVersion(t *testing.T) {
	v := coreversion.InvalidCompactVersion("bad")
	actual := args.Map{"compact": v.VersionCompact, "invalid": v.IsInvalid}
	expected := args.Map{"compact": "bad", "invalid": true}
	expected.ShouldBeEqual(t, 0, "InvalidCompactVersion returns error -- with args", actual)
}

// ==========================================================================
// Version — nil receiver methods
// ==========================================================================

func Test_Cov4_Version_NilReceiver(t *testing.T) {
	var v *coreversion.Version
	actual := args.Map{
		"display":   v.VersionDisplay(),
		"compiled":  v.CompiledVersion(),
		"major":     v.MajorString(),
		"minor":     v.MinorString(),
		"patch":     v.PatchString(),
		"build":     v.BuildString(),
		"hasMajor":  v.HasMajor(),
		"hasMinor":  v.HasMinor(),
		"hasPatch":  v.HasPatch(),
		"hasBuild":  v.HasBuild(),
		"emptyOrInv": v.IsEmptyOrInvalid(),
		"cloneNil":  v.ClonePtr() == nil,
	}
	expected := args.Map{
		"display":   "",
		"compiled":  "",
		"major":     "",
		"minor":     "",
		"patch":     "",
		"build":     "",
		"hasMajor":  false,
		"hasMinor":  false,
		"hasPatch":  false,
		"hasBuild":  false,
		"emptyOrInv": true,
		"cloneNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "Version returns nil -- nil receiver methods", actual)
}

// ==========================================================================
// Version — comparison methods
// ==========================================================================

func Test_Cov4_Version_IsMajorBuildAtLeast(t *testing.T) {
	v := coreversion.New.Create("v2.0.0.5")
	actual := args.Map{
		"atLeast": v.IsMajorBuildAtLeast(2, 5),
		"below":  v.IsMajorBuildAtLeast(2, 10),
	}
	expected := args.Map{
		"atLeast": true,
		"below":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsMajorBuildAtLeast returns correct value -- with args", actual)
}

func Test_Cov4_Version_IsMajorMinorPatchAtLeast(t *testing.T) {
	v := coreversion.New.Create("v3.2.1")
	actual := args.Map{
		"atLeast": v.IsMajorMinorPatchAtLeast(3, 2, 1),
		"below":  v.IsMajorMinorPatchAtLeast(3, 2, 5),
	}
	expected := args.Map{
		"atLeast": true,
		"below":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsMajorMinorPatchAtLeast returns correct value -- with args", actual)
}

func Test_Cov4_Version_MajorMinorPatchBuildString(t *testing.T) {
	v := coreversion.New.Create("v1.2.3.4")
	cmp := v.MajorMinorPatchBuildString("1", "2", "4", "3")
	actual := args.Map{"eq": cmp.IsEqual()}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MajorMinorPatchBuildString returns correct value -- with args", actual)
}

func Test_Cov4_Version_MajorBuildString(t *testing.T) {
	v := coreversion.New.Create("v1.0.0.5")
	cmp := v.MajorBuildString("1", "5")
	actual := args.Map{"eq": cmp.IsEqual()}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MajorBuildString returns correct value -- with args", actual)
}

func Test_Cov4_Version_ComparisonValueIndexes(t *testing.T) {
	v1 := coreversion.New.Create("v1.2.3")
	v2 := coreversion.New.Create("v1.2.3")
	cmp := v1.ComparisonValueIndexes(&v2, versionindexes.Major, versionindexes.Minor)
	actual := args.Map{"eq": cmp.IsEqual()}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "ComparisonValueIndexes returns correct value -- equal", actual)
}

func Test_Cov4_Version_ComparisonValueIndexes_NilRight(t *testing.T) {
	v1 := coreversion.New.Create("v1.0")
	cmp := v1.ComparisonValueIndexes(nil, versionindexes.Major)
	actual := args.Map{"greater": cmp.IsLeftGreater()}
	expected := args.Map{"greater": true}
	expected.ShouldBeEqual(t, 0, "ComparisonValueIndexes returns nil -- nil right", actual)
}

func Test_Cov4_Version_IsVersionCompareEqual_NilBothEmpty(t *testing.T) {
	var v *coreversion.Version
	actual := args.Map{
		"nilEmpty": v.IsVersionCompareEqual(""),
		"nilNonEmpty": v.IsVersionCompareEqual("1.0"),
	}
	expected := args.Map{
		"nilEmpty": true,
		"nilNonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsVersionCompareEqual returns nil -- nil receiver", actual)
}

func Test_Cov4_Version_Clone(t *testing.T) {
	v := coreversion.New.Create("v1.2.3")
	c := v.Clone()
	actual := args.Map{"compact": c.VersionCompact}
	expected := args.Map{"compact": "1.2.3"}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- copies version", actual)
}

func Test_Cov4_Version_NonPtr_Ptr(t *testing.T) {
	v := coreversion.New.Create("v1.0")
	np := v.NonPtr()
	p := v.Ptr()
	actual := args.Map{"nonPtrCompact": np.VersionCompact, "ptrNotNil": p != nil}
	expected := args.Map{"nonPtrCompact": "1.0", "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "NonPtr returns correct value -- and Ptr", actual)
}

// ==========================================================================
// Package-level comparison functions
// ==========================================================================

func Test_Cov4_CompareVersionString(t *testing.T) {
	cmp := coreversion.CompareVersionString("1.0.0", "2.0.0")
	actual := args.Map{"less": cmp.IsLeftLess()}
	expected := args.Map{"less": true}
	expected.ShouldBeEqual(t, 0, "CompareVersionString returns correct value -- with args", actual)
}

func Test_Cov4_IsExpectedVersion(t *testing.T) {
	actual := args.Map{
		"eq": coreversion.IsExpectedVersion(corecomparator.Equal, "1.0", "1.0"),
	}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsExpectedVersion returns correct value -- with args", actual)
}

func Test_Cov4_IsAtLeast(t *testing.T) {
	actual := args.Map{
		"atLeast": coreversion.IsAtLeast("2.0", "1.0"),
		"below":   coreversion.IsAtLeast("1.0", "2.0"),
	}
	expected := args.Map{"atLeast": true, "below": false}
	expected.ShouldBeEqual(t, 0, "IsAtLeast returns correct value -- with args", actual)
}

func Test_Cov4_IsLower(t *testing.T) {
	actual := args.Map{
		"lower": coreversion.IsLower("1.0", "2.0"),
	}
	expected := args.Map{"lower": true}
	expected.ShouldBeEqual(t, 0, "IsLower returns correct value -- with args", actual)
}

func Test_Cov4_IsLowerOrEqual(t *testing.T) {
	actual := args.Map{
		"lowerEq": coreversion.IsLowerOrEqual("1.0", "1.0"),
	}
	expected := args.Map{"lowerEq": true}
	expected.ShouldBeEqual(t, 0, "IsLowerOrEqual returns correct value -- with args", actual)
}

// ==========================================================================
// Version — string comparison methods
// ==========================================================================

func Test_Cov4_Version_IsEqualVersionString(t *testing.T) {
	v := coreversion.New.Create("v1.0.0")
	actual := args.Map{"eq": v.IsEqualVersionString("1.0.0")}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqualVersionString returns correct value -- with args", actual)
}

func Test_Cov4_Version_IsLowerVersionString(t *testing.T) {
	v := coreversion.New.Create("v1.0.0")
	actual := args.Map{"lower": v.IsLowerVersionString("2.0.0")}
	expected := args.Map{"lower": true}
	expected.ShouldBeEqual(t, 0, "IsLowerVersionString returns correct value -- with args", actual)
}

func Test_Cov4_Version_IsLowerEqualVersionString(t *testing.T) {
	v := coreversion.New.Create("v1.0.0")
	actual := args.Map{"lowerEq": v.IsLowerEqualVersionString("1.0.0")}
	expected := args.Map{"lowerEq": true}
	expected.ShouldBeEqual(t, 0, "IsLowerEqualVersionString returns correct value -- with args", actual)
}

func Test_Cov4_Version_IsAtLeast(t *testing.T) {
	v := coreversion.New.Create("v2.0.0")
	actual := args.Map{"atLeast": v.IsAtLeast("1.0.0")}
	expected := args.Map{"atLeast": true}
	expected.ShouldBeEqual(t, 0, "Version.IsAtLeast returns correct value -- with args", actual)
}

func Test_Cov4_Version_IsExpectedComparisonRawVersion(t *testing.T) {
	v := coreversion.New.Create("v1.0.0")
	actual := args.Map{"eq": v.IsExpectedComparisonRawVersion(corecomparator.Equal, "1.0.0")}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsExpectedComparisonRawVersion returns correct value -- with args", actual)
}

// ==========================================================================
// Version — IsMajorStringAtLeast + IsMajorMinorAtLeast
// ==========================================================================

func Test_Cov4_Version_IsMajorStringAtLeast(t *testing.T) {
	v := coreversion.New.Create("v3.0.0")
	actual := args.Map{"atLeast": v.IsMajorStringAtLeast("2")}
	expected := args.Map{"atLeast": true}
	expected.ShouldBeEqual(t, 0, "IsMajorStringAtLeast returns correct value -- with args", actual)
}

func Test_Cov4_Version_IsMajorMinorAtLeast(t *testing.T) {
	v := coreversion.New.Create("v2.5.0")
	actual := args.Map{
		"atLeast": v.IsMajorMinorAtLeast(2, 5),
		"below":   v.IsMajorMinorAtLeast(2, 6),
	}
	expected := args.Map{"atLeast": true, "below": false}
	expected.ShouldBeEqual(t, 0, "IsMajorMinorAtLeast returns correct value -- with args", actual)
}

// ==========================================================================
// newCreator — SpreadIntegers, SpreadUnsignedIntegers, SpreadBytes
// ==========================================================================

func Test_Cov4_New_SpreadIntegers(t *testing.T) {
	v := coreversion.New.SpreadIntegers(1, 2, 3)
	actual := args.Map{"major": v.VersionMajor, "minor": v.VersionMinor, "patch": v.VersionPatch}
	expected := args.Map{"major": 1, "minor": 2, "patch": 3}
	expected.ShouldBeEqual(t, 0, "SpreadIntegers returns correct value -- creates version", actual)
}

func Test_Cov4_New_SpreadUnsignedIntegers(t *testing.T) {
	v := coreversion.New.SpreadUnsignedIntegers(1, 2)
	actual := args.Map{"major": v.VersionMajor, "minor": v.VersionMinor}
	expected := args.Map{"major": 1, "minor": 2}
	expected.ShouldBeEqual(t, 0, "SpreadUnsignedIntegers returns correct value -- creates version", actual)
}

func Test_Cov4_New_SpreadBytes(t *testing.T) {
	v := coreversion.New.SpreadBytes(1, 2, 3, 4)
	actual := args.Map{"major": v.VersionMajor}
	expected := args.Map{"major": 1}
	expected.ShouldBeEqual(t, 0, "SpreadBytes returns correct value -- creates version", actual)
}

func Test_Cov4_New_AllByte(t *testing.T) {
	v := coreversion.New.AllByte(1, 2, 3, 4)
	actual := args.Map{"major": v.VersionMajor, "build": v.VersionBuild}
	expected := args.Map{"major": 1, "build": 4}
	expected.ShouldBeEqual(t, 0, "AllByte returns correct value -- creates version", actual)
}

func Test_Cov4_New_MajorBuildInt(t *testing.T) {
	v := coreversion.New.MajorBuildInt(1, 5)
	actual := args.Map{"major": v.VersionMajor}
	expected := args.Map{"major": 1}
	expected.ShouldBeEqual(t, 0, "MajorBuildInt returns correct value -- creates version", actual)
}

func Test_Cov4_New_MajorMinorBuild(t *testing.T) {
	v := coreversion.New.MajorMinorBuild("1", "2", "5")
	actual := args.Map{"major": v.VersionMajor, "minor": v.VersionMinor}
	expected := args.Map{"major": 1, "minor": 2}
	expected.ShouldBeEqual(t, 0, "MajorMinorBuild returns correct value -- creates version", actual)
}

func Test_Cov4_New_MajorPatch(t *testing.T) {
	v := coreversion.New.MajorPatch("3", "7")
	actual := args.Map{"major": v.VersionMajor}
	expected := args.Map{"major": 3}
	expected.ShouldBeEqual(t, 0, "MajorPatch returns correct value -- creates version", actual)
}

func Test_Cov4_New_MajorPatchInt(t *testing.T) {
	v := coreversion.New.MajorPatchInt(3, 7)
	actual := args.Map{"major": v.VersionMajor}
	expected := args.Map{"major": 3}
	expected.ShouldBeEqual(t, 0, "MajorPatchInt returns correct value -- creates version", actual)
}

func Test_Cov4_New_MajorBuild(t *testing.T) {
	v := coreversion.New.MajorBuild("2", "9")
	actual := args.Map{"major": v.VersionMajor}
	expected := args.Map{"major": 2}
	expected.ShouldBeEqual(t, 0, "MajorBuild returns correct value -- creates version", actual)
}

// ==========================================================================
// Version — JSON + AsJsonContractsBinder
// ==========================================================================

func Test_Cov4_Version_Json(t *testing.T) {
	v := coreversion.New.Create("v1.0.0")
	j := v.Json()
	jp := v.JsonPtr()
	actual := args.Map{"hasResult": j.HasSafeItems(), "ptrNotNil": jp != nil}
	expected := args.Map{"hasResult": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Version returns correct value -- Json and JsonPtr", actual)
}

func Test_Cov4_Version_AsJsonContractsBinder(t *testing.T) {
	v := coreversion.New.DefaultPtr("v1.0.0")
	binder := v.AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_Cov4_VersionsCollection_AsJsonContractsBinder(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0")
	binder := vc.AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "VersionsCollection returns correct value -- AsJsonContractsBinder", actual)
}

func Test_Cov4_VersionsCollection_AsBasicSliceContractsBinder(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	binder := vc.AsBasicSliceContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsBasicSliceContractsBinder returns correct value -- with args", actual)
}

// ==========================================================================
// hasDeductUsingNilNess — all branches
// ==========================================================================

func Test_Cov4_Compare_BothNil(t *testing.T) {
	cmp := coreversion.Compare(nil, nil)
	actual := args.Map{"eq": cmp.IsEqual()}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Compare returns nil -- both nil", actual)
}

func Test_Cov4_Compare_LeftNilRightNonNil(t *testing.T) {
	v := coreversion.New.DefaultPtr("1.0")
	cmp := coreversion.Compare(nil, v)
	actual := args.Map{"leftLess": cmp.IsLeftLess()}
	expected := args.Map{"leftLess": true}
	expected.ShouldBeEqual(t, 0, "Compare returns nil -- left nil", actual)
}

func Test_Cov4_Compare_SamePtr(t *testing.T) {
	v := coreversion.New.DefaultPtr("1.0")
	cmp := coreversion.Compare(v, v)
	actual := args.Map{"eq": cmp.IsEqual()}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- same ptr", actual)
}

func Test_Cov4_Compare_SameCompact(t *testing.T) {
	v1 := coreversion.New.DefaultPtr("1.0")
	v2 := coreversion.New.DefaultPtr("1.0")
	cmp := coreversion.Compare(v1, v2)
	actual := args.Map{"eq": cmp.IsEqual()}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- same compact", actual)
}

// ==========================================================================
// Version — ValueByIndex default branch
// ==========================================================================

func Test_Cov4_Version_ValueByIndex_Invalid(t *testing.T) {
	v := coreversion.New.Create("v1.2.3")
	val := v.ValueByIndex(versionindexes.Index(99))
	actual := args.Map{"val": val}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "ValueByIndex returns error -- returns -1 for invalid index", actual)
}
