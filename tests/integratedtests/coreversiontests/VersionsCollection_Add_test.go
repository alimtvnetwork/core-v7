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

func Test_VersionsCollection_Add(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")

	// Act
	actual := args.Map{"len": vc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Add returns correct value -- version to collection", actual)
}

func Test_VersionsCollection_AddSkipInvalid_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.AddSkipInvalid("1.0.0")
	vc.AddSkipInvalid("")       // skipped
	vc.AddSkipInvalid("v")      // skipped

	// Act
	actual := args.Map{"len": vc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddSkipInvalid returns empty -- skips empty", actual)
}

func Test_VersionsCollection_AddVersionsRaw_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.AddVersionsRaw("1.0", "2.0")

	// Act
	actual := args.Map{"len": vc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddVersionsRaw returns correct value -- adds multiple", actual)
}

func Test_VersionsCollection_AddVersions_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	v := coreversion.New.Create("1.0.0")
	vc.AddVersions(v)

	// Act
	actual := args.Map{"len": vc.Count()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddVersions returns correct value -- adds version struct", actual)
}

func Test_VersionsCollection_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}

	// Act
	actual := args.Map{
		"empty": vc.IsEmpty(),
		"hasAny": vc.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"hasAny": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- and HasAnyItem on empty", actual)
}

func Test_VersionsCollection_LastIndex_HasIndex(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0").Add("2.0")

	// Act
	actual := args.Map{
		"lastIdx": vc.LastIndex(),
		"hasIdx0": vc.HasIndex(0),
		"hasIdx5": vc.HasIndex(5),
	}

	// Assert
	expected := args.Map{
		"lastIdx": 1,
		"hasIdx0": true,
		"hasIdx5": false,
	}
	expected.ShouldBeEqual(t, 0, "LastIndex returns correct value -- and HasIndex", actual)
}

func Test_VersionsCollection_VersionCompactStrings(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	strs := vc.VersionCompactStrings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VersionCompactStrings returns correct value -- returns strings", actual)
}

func Test_VersionsCollection_VersionCompactStrings_Empty(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	strs := vc.VersionCompactStrings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VersionCompactStrings returns empty -- empty", actual)
}

func Test_VersionsCollection_VersionsStrings(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	strs := vc.VersionsStrings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VersionsStrings returns correct value -- returns display strings", actual)
}

func Test_VersionsCollection_VersionsStrings_Empty(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	strs := vc.VersionsStrings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VersionsStrings returns empty -- empty", actual)
}

func Test_VersionsCollection_IndexOf_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0").Add("2.0.0")

	// Act
	actual := args.Map{
		"found": vc.IndexOf("1.0.0") >= 0,
		"notFound": vc.IndexOf("3.0.0") < 0,
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": true,
	}
	expected.ShouldBeEqual(t, 0, "IndexOf returns correct value -- finds version", actual)
}

func Test_VersionsCollection_IsContainsVersion_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")

	// Act
	actual := args.Map{
		"contains": vc.IsContainsVersion("1.0.0"),
		"missing": vc.IsContainsVersion("9.9"),
	}

	// Assert
	expected := args.Map{
		"contains": true,
		"missing": false,
	}
	expected.ShouldBeEqual(t, 0, "IsContainsVersion returns correct value -- with args", actual)
}

func Test_VersionsCollection_IsEqual_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("1.0.0")
	vc3 := &coreversion.VersionsCollection{}
	vc3.Add("2.0.0")

	// Act
	actual := args.Map{
		"eq":      vc1.IsEqual(vc2),
		"neq":     vc1.IsEqual(vc3),
		"nilNil":  (*coreversion.VersionsCollection)(nil).IsEqual(nil),
		"nilR":    vc1.IsEqual(nil),
	}

	// Assert
	expected := args.Map{
		"eq":      true,
		"neq":     false,
		"nilNil":  true,
		"nilR":    false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- comparisons", actual)
}

func Test_VersionsCollection_IsEqual_DiffLen(t *testing.T) {
	// Arrange
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("1.0.0").Add("2.0.0")

	// Act
	actual := args.Map{"eq": vc1.IsEqual(vc2)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- diff lengths", actual)
}

func Test_VersionsCollection_String_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	s := vc.String()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- returns display", actual)
}

func Test_VersionsCollection_Json_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.Json()
	jp := vc.JsonPtr()

	// Act
	actual := args.Map{
		"hasResult": j.HasSafeItems(),
		"ptrNotNil": jp != nil,
	}

	// Assert
	expected := args.Map{
		"hasResult": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- and JsonPtr", actual)
}

func Test_VersionsCollection_Length_Nil_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	var vc *coreversion.VersionsCollection

	// Act
	actual := args.Map{"len": vc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- on nil returns 0", actual)
}

// ==========================================================================
// EmptyUsingCompactVersion + InvalidCompactVersion
// ==========================================================================

func Test_EmptyUsingCompactVersion_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.EmptyUsingCompactVersion("1.2.3")

	// Act
	actual := args.Map{
		"compact": v.VersionCompact,
		"invalid": v.IsInvalid,
	}

	// Assert
	expected := args.Map{
		"compact": "1.2.3",
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "EmptyUsingCompactVersion returns empty -- with args", actual)
}

func Test_InvalidCompactVersion_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.InvalidCompactVersion("bad")

	// Act
	actual := args.Map{
		"compact": v.VersionCompact,
		"invalid": v.IsInvalid,
	}

	// Assert
	expected := args.Map{
		"compact": "bad",
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidCompactVersion returns error -- with args", actual)
}

// ==========================================================================
// Version — nil receiver methods
// ==========================================================================

func Test_Version_NilReceiver_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
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

	// Assert
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

func Test_Version_IsMajorBuildAtLeast_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v2.0.0.5")

	// Act
	actual := args.Map{
		"atLeast": v.IsMajorBuildAtLeast(2, 5),
		"below":  v.IsMajorBuildAtLeast(2, 10),
	}

	// Assert
	expected := args.Map{
		"atLeast": true,
		"below":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsMajorBuildAtLeast returns correct value -- with args", actual)
}

func Test_Version_IsMajorMinorPatchAtLeast_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v3.2.1")

	// Act
	actual := args.Map{
		"atLeast": v.IsMajorMinorPatchAtLeast(3, 2, 1),
		"below":  v.IsMajorMinorPatchAtLeast(3, 2, 5),
	}

	// Assert
	expected := args.Map{
		"atLeast": true,
		"below":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsMajorMinorPatchAtLeast returns correct value -- with args", actual)
}

func Test_Version_MajorMinorPatchBuildString_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.2.3.4")
	cmp := v.MajorMinorPatchBuildString("1", "2", "4", "3")

	// Act
	actual := args.Map{"eq": cmp.IsEqual()}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MajorMinorPatchBuildString returns correct value -- with args", actual)
}

func Test_Version_MajorBuildString_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.0.0.5")
	cmp := v.MajorBuildString("1", "5")

	// Act
	actual := args.Map{"eq": cmp.IsEqual()}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MajorBuildString returns correct value -- with args", actual)
}

func Test_Version_ComparisonValueIndexes_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("v1.2.3")
	v2 := coreversion.New.Create("v1.2.3")
	cmp := v1.ComparisonValueIndexes(&v2, versionindexes.Major, versionindexes.Minor)

	// Act
	actual := args.Map{"eq": cmp.IsEqual()}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "ComparisonValueIndexes returns correct value -- equal", actual)
}

func Test_Version_ComparisonValueIndexes_NilRight_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("v1.0")
	cmp := v1.ComparisonValueIndexes(nil, versionindexes.Major)

	// Act
	actual := args.Map{"greater": cmp.IsLeftGreater()}

	// Assert
	expected := args.Map{"greater": true}
	expected.ShouldBeEqual(t, 0, "ComparisonValueIndexes returns nil -- nil right", actual)
}

func Test_Version_IsVersionCompareEqual_NilBothEmpty(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{
		"nilEmpty": v.IsVersionCompareEqual(""),
		"nilNonEmpty": v.IsVersionCompareEqual("1.0"),
	}

	// Assert
	expected := args.Map{
		"nilEmpty": true,
		"nilNonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsVersionCompareEqual returns nil -- nil receiver", actual)
}

func Test_Version_Clone_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.2.3")
	c := v.Clone()

	// Act
	actual := args.Map{"compact": c.VersionCompact}

	// Assert
	expected := args.Map{"compact": "1.2.3"}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- copies version", actual)
}

func Test_Version_NonPtr_Ptr(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.0")
	np := v.NonPtr()
	p := v.Ptr()

	// Act
	actual := args.Map{
		"nonPtrCompact": np.VersionCompact,
		"ptrNotNil": p != nil,
	}

	// Assert
	expected := args.Map{
		"nonPtrCompact": "1.0",
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "NonPtr returns correct value -- and Ptr", actual)
}

// ==========================================================================
// Package-level comparison functions
// ==========================================================================

func Test_CompareVersionString_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	cmp := coreversion.CompareVersionString("1.0.0", "2.0.0")

	// Act
	actual := args.Map{"less": cmp.IsLeftLess()}

	// Assert
	expected := args.Map{"less": true}
	expected.ShouldBeEqual(t, 0, "CompareVersionString returns correct value -- with args", actual)
}

func Test_IsExpectedVersion_FromVersionsCollectionAd(t *testing.T) {
	// Act
	actual := args.Map{
		"eq": coreversion.IsExpectedVersion(corecomparator.Equal, "1.0", "1.0"),
	}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsExpectedVersion returns correct value -- with args", actual)
}

func Test_IsAtLeast_FromVersionsCollectionAd(t *testing.T) {
	// Act
	actual := args.Map{
		"atLeast": coreversion.IsAtLeast("2.0", "1.0"),
		"below":   coreversion.IsAtLeast("1.0", "2.0"),
	}

	// Assert
	expected := args.Map{
		"atLeast": true,
		"below": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAtLeast returns correct value -- with args", actual)
}

func Test_IsLower_FromVersionsCollectionAd(t *testing.T) {
	// Act
	actual := args.Map{
		"lower": coreversion.IsLower("1.0", "2.0"),
	}

	// Assert
	expected := args.Map{"lower": true}
	expected.ShouldBeEqual(t, 0, "IsLower returns correct value -- with args", actual)
}

func Test_IsLowerOrEqual_FromVersionsCollectionAd(t *testing.T) {
	// Act
	actual := args.Map{
		"lowerEq": coreversion.IsLowerOrEqual("1.0", "1.0"),
	}

	// Assert
	expected := args.Map{"lowerEq": true}
	expected.ShouldBeEqual(t, 0, "IsLowerOrEqual returns correct value -- with args", actual)
}

// ==========================================================================
// Version — string comparison methods
// ==========================================================================

func Test_Version_IsEqualVersionString_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.0.0")

	// Act
	actual := args.Map{"eq": v.IsEqualVersionString("1.0.0")}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqualVersionString returns correct value -- with args", actual)
}

func Test_Version_IsLowerVersionString_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.0.0")

	// Act
	actual := args.Map{"lower": v.IsLowerVersionString("2.0.0")}

	// Assert
	expected := args.Map{"lower": true}
	expected.ShouldBeEqual(t, 0, "IsLowerVersionString returns correct value -- with args", actual)
}

func Test_Version_IsLowerEqualVersionString_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.0.0")

	// Act
	actual := args.Map{"lowerEq": v.IsLowerEqualVersionString("1.0.0")}

	// Assert
	expected := args.Map{"lowerEq": true}
	expected.ShouldBeEqual(t, 0, "IsLowerEqualVersionString returns correct value -- with args", actual)
}

func Test_Version_IsAtLeast_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v2.0.0")

	// Act
	actual := args.Map{"atLeast": v.IsAtLeast("1.0.0")}

	// Assert
	expected := args.Map{"atLeast": true}
	expected.ShouldBeEqual(t, 0, "Version.IsAtLeast returns correct value -- with args", actual)
}

func Test_Version_IsExpectedComparisonRawVersion_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.0.0")

	// Act
	actual := args.Map{"eq": v.IsExpectedComparisonRawVersion(corecomparator.Equal, "1.0.0")}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsExpectedComparisonRawVersion returns correct value -- with args", actual)
}

// ==========================================================================
// Version — IsMajorStringAtLeast + IsMajorMinorAtLeast
// ==========================================================================

func Test_Version_IsMajorStringAtLeast_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v3.0.0")

	// Act
	actual := args.Map{"atLeast": v.IsMajorStringAtLeast("2")}

	// Assert
	expected := args.Map{"atLeast": true}
	expected.ShouldBeEqual(t, 0, "IsMajorStringAtLeast returns correct value -- with args", actual)
}

func Test_Version_IsMajorMinorAtLeast_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v2.5.0")

	// Act
	actual := args.Map{
		"atLeast": v.IsMajorMinorAtLeast(2, 5),
		"below":   v.IsMajorMinorAtLeast(2, 6),
	}

	// Assert
	expected := args.Map{
		"atLeast": true,
		"below": false,
	}
	expected.ShouldBeEqual(t, 0, "IsMajorMinorAtLeast returns correct value -- with args", actual)
}

// ==========================================================================
// newCreator — SpreadIntegers, SpreadUnsignedIntegers, SpreadBytes
// ==========================================================================

func Test_New_SpreadIntegers_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.SpreadIntegers(1, 2, 3)

	// Act
	actual := args.Map{
		"major": v.VersionMajor,
		"minor": v.VersionMinor,
		"patch": v.VersionPatch,
	}

	// Assert
	expected := args.Map{
		"major": 1,
		"minor": 2,
		"patch": 3,
	}
	expected.ShouldBeEqual(t, 0, "SpreadIntegers returns correct value -- creates version", actual)
}

func Test_New_SpreadUnsignedIntegers_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.SpreadUnsignedIntegers(1, 2)

	// Act
	actual := args.Map{
		"major": v.VersionMajor,
		"minor": v.VersionMinor,
	}

	// Assert
	expected := args.Map{
		"major": 1,
		"minor": 2,
	}
	expected.ShouldBeEqual(t, 0, "SpreadUnsignedIntegers returns correct value -- creates version", actual)
}

func Test_New_SpreadBytes_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.SpreadBytes(1, 2, 3, 4)

	// Act
	actual := args.Map{"major": v.VersionMajor}

	// Assert
	expected := args.Map{"major": 1}
	expected.ShouldBeEqual(t, 0, "SpreadBytes returns correct value -- creates version", actual)
}

func Test_New_AllByte_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.AllByte(1, 2, 3, 4)

	// Act
	actual := args.Map{
		"major": v.VersionMajor,
		"build": v.VersionBuild,
	}

	// Assert
	expected := args.Map{
		"major": 1,
		"build": 4,
	}
	expected.ShouldBeEqual(t, 0, "AllByte returns correct value -- creates version", actual)
}

func Test_New_MajorBuildInt_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorBuildInt(1, 5)

	// Act
	actual := args.Map{"major": v.VersionMajor}

	// Assert
	expected := args.Map{"major": 1}
	expected.ShouldBeEqual(t, 0, "MajorBuildInt returns correct value -- creates version", actual)
}

func Test_New_MajorMinorBuild_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorMinorBuild("1", "2", "5")

	// Act
	actual := args.Map{
		"major": v.VersionMajor,
		"minor": v.VersionMinor,
	}

	// Assert
	expected := args.Map{
		"major": 1,
		"minor": 2,
	}
	expected.ShouldBeEqual(t, 0, "MajorMinorBuild returns correct value -- creates version", actual)
}

func Test_New_MajorPatch_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorPatch("3", "7")

	// Act
	actual := args.Map{"major": v.VersionMajor}

	// Assert
	expected := args.Map{"major": 3}
	expected.ShouldBeEqual(t, 0, "MajorPatch returns correct value -- creates version", actual)
}

func Test_New_MajorPatchInt_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorPatchInt(3, 7)

	// Act
	actual := args.Map{"major": v.VersionMajor}

	// Assert
	expected := args.Map{"major": 3}
	expected.ShouldBeEqual(t, 0, "MajorPatchInt returns correct value -- creates version", actual)
}

func Test_New_MajorBuild_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.MajorBuild("2", "9")

	// Act
	actual := args.Map{"major": v.VersionMajor}

	// Assert
	expected := args.Map{"major": 2}
	expected.ShouldBeEqual(t, 0, "MajorBuild returns correct value -- creates version", actual)
}

// ==========================================================================
// Version — JSON + AsJsonContractsBinder
// ==========================================================================

func Test_Version_Json_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.0.0")
	j := v.Json()
	jp := v.JsonPtr()

	// Act
	actual := args.Map{
		"hasResult": j.HasSafeItems(),
		"ptrNotNil": jp != nil,
	}

	// Assert
	expected := args.Map{
		"hasResult": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Version returns correct value -- Json and JsonPtr", actual)
}

func Test_Version_AsJsonContractsBinder_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.DefaultPtr("v1.0.0")
	binder := v.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_VersionsCollection_AsJsonContractsBinder_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0")
	binder := vc.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "VersionsCollection returns correct value -- AsJsonContractsBinder", actual)
}

func Test_VersionsCollection_AsBasicSliceContractsBinder_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	binder := vc.AsBasicSliceContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsBasicSliceContractsBinder returns correct value -- with args", actual)
}

// ==========================================================================
// hasDeductUsingNilNess — all branches
// ==========================================================================

func Test_Compare_BothNil_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	cmp := coreversion.Compare(nil, nil)

	// Act
	actual := args.Map{"eq": cmp.IsEqual()}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Compare returns nil -- both nil", actual)
}

func Test_Compare_LeftNilRightNonNil(t *testing.T) {
	// Arrange
	v := coreversion.New.DefaultPtr("1.0")
	cmp := coreversion.Compare(nil, v)

	// Act
	actual := args.Map{"leftLess": cmp.IsLeftLess()}

	// Assert
	expected := args.Map{"leftLess": true}
	expected.ShouldBeEqual(t, 0, "Compare returns nil -- left nil", actual)
}

func Test_Compare_SamePtr_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.DefaultPtr("1.0")
	cmp := coreversion.Compare(v, v)

	// Act
	actual := args.Map{"eq": cmp.IsEqual()}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- same ptr", actual)
}

func Test_Compare_SameCompact_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v1 := coreversion.New.DefaultPtr("1.0")
	v2 := coreversion.New.DefaultPtr("1.0")
	cmp := coreversion.Compare(v1, v2)

	// Act
	actual := args.Map{"eq": cmp.IsEqual()}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- same compact", actual)
}

// ==========================================================================
// Version — ValueByIndex default branch
// ==========================================================================

func Test_Version_ValueByIndex_Invalid_FromVersionsCollectionAd(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("v1.2.3")
	val := v.ValueByIndex(versionindexes.Index(99))

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "ValueByIndex returns error -- returns -1 for invalid index", actual)
}
