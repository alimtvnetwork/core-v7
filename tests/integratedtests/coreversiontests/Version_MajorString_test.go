package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Version field accessors
// ==========================================

func Test_Version_MajorString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.MajorString() != "1"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '1', got ''", actual)
}

func Test_Version_MinorString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.MinorString() != "2"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '2', got ''", actual)
}

func Test_Version_PatchString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.PatchString() != "3"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '3', got ''", actual)
}

func Test_Version_BuildString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")

	// Act
	actual := args.Map{"result": v.BuildString() != "4"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '4', got ''", actual)
}

// ==========================================
// Version nil receivers
// ==========================================

func Test_Version_Nil_MajorString(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.MajorString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Version_Nil_MinorString(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.MinorString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Version_Nil_PatchString(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.PatchString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Version_Nil_BuildString(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.BuildString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// ==========================================
// Has / IsInvalid checks
// ==========================================

func Test_Version_HasMajor(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v.HasMajor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
}

func Test_Version_HasMinor(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.0")

	// Act
	actual := args.Map{"result": v.HasMinor()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have minor", actual)
}

func Test_Version_HasPatch(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.HasPatch()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have patch", actual)
}

func Test_Version_HasBuild(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")

	// Act
	actual := args.Map{"result": v.HasBuild()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have build", actual)
}

func Test_Version_IsMajorInvalid_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.IsMajorInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Version_IsMinorInvalid_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.IsMinorInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Version_IsPatchInvalid_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.IsPatchInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Version_IsBuildInvalid_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.IsBuildInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

// ==========================================
// InvalidOrZero checks
// ==========================================

func Test_Version_IsMajorInvalidOrZero(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("0.1.0")

	// Act
	actual := args.Map{"result": v.IsMajorInvalidOrZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "major 0 should be invalid or zero", actual)
}

func Test_Version_IsMinorInvalidOrZero(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.0.0")

	// Act
	actual := args.Map{"result": v.IsMinorInvalidOrZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "minor 0 should be invalid or zero", actual)
}

func Test_Version_IsPatchInvalidOrZero(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.1.0")

	// Act
	actual := args.Map{"result": v.IsPatchInvalidOrZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "patch 0 should be invalid or zero", actual)
}

func Test_Version_IsBuildInvalidOrZero(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.1.1")

	// Act
	actual := args.Map{"result": v.IsBuildInvalidOrZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "no build should be invalid or zero", actual)
}

// ==========================================
// Display methods
// ==========================================

func Test_Version_VersionDisplay(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	d := v.VersionDisplay()

	// Act
	actual := args.Map{"result": d == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty display", actual)
}

func Test_Version_VersionDisplay_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.VersionDisplay() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Version_VersionDisplayMajor(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("5.2.3")
	d := v.VersionDisplayMajor()

	// Act
	actual := args.Map{"result": d == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Version_VersionDisplayMajorMinor(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("5.2.3")
	d := v.VersionDisplayMajorMinor()

	// Act
	actual := args.Map{"result": d == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Version_VersionDisplayMajorMinorPatch(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("5.2.3")
	d := v.VersionDisplayMajorMinorPatch()

	// Act
	actual := args.Map{"result": d == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Version_CompiledVersion(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.CompiledVersion() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Version_CompiledVersion_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.CompiledVersion() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Version_String(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// Validity checks
// ==========================================

func Test_Version_IsEmptyOrInvalid_Valid(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "valid version should not be empty/invalid", actual)
}

func Test_Version_IsEmptyOrInvalid_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.IsEmptyOrInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty/invalid", actual)
}

func Test_Version_HasAnyItem(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have item", actual)
}

func Test_Version_IsDefined(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.IsDefined()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be defined", actual)
}

// ==========================================
// Version comparison
// ==========================================

func Test_Version_IsVersionCompareEqual(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.IsVersionCompareEqual("1.2.3")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_IsVersionCompareNotEqual(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v.IsVersionCompareNotEqual("2.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should not be equal", actual)
}

func Test_Version_IsVersionCompareEqual_Nil(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{"result": v.IsVersionCompareEqual("")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil with empty should be equal", actual)
	actual = args.Map{"result": v.IsVersionCompareEqual("1.0.0")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil with value should not be equal", actual)
}

// ==========================================
// ValueByIndex / AllVersionValues
// ==========================================

func Test_Version_AllVersionValues(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")
	vals := v.AllVersionValues()

	// Act
	actual := args.Map{"result": len(vals) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Version_AllValidVersionValues(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	vals := v.AllValidVersionValues()

	// Act
	actual := args.Map{"result": len(vals) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// IsMajorAtLeast / IsMajorMinorAtLeast etc.
// ==========================================

func Test_Version_IsMajorAtLeast(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.0.0")

	// Act
	actual := args.Map{"result": v.IsMajorAtLeast(2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3 should be at least 2", actual)
	actual = args.Map{"result": v.IsMajorAtLeast(5)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3 should not be at least 5", actual)
}

func Test_Version_IsMajorMinorAtLeast(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.2.0")

	// Act
	actual := args.Map{"result": v.IsMajorMinorAtLeast(3, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3.2 should be at least 3.1", actual)
}

func Test_Version_IsMajorMinorPatchAtLeast(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.2.1")

	// Act
	actual := args.Map{"result": v.IsMajorMinorPatchAtLeast(3, 2, 0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3.2.1 should be at least 3.2.0", actual)
}

func Test_Version_IsMajorStringAtLeast(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.0.0")

	// Act
	actual := args.Map{"result": v.IsMajorStringAtLeast("2")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3 should be at least 2", actual)
}

// ==========================================
// Compare functions
// ==========================================

func Test_CompareVersionString(t *testing.T) {
	// Arrange
	cmp := coreversion.CompareVersionString("2.0.0", "1.0.0")

	// Act
	actual := args.Map{"result": cmp.IsLeftGreaterEqualLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2.0.0 should be greater than 1.0.0", actual)
}

func Test_IsAtLeast(t *testing.T) {
	// Act
	actual := args.Map{"result": coreversion.IsAtLeast("2.0.0", "1.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2.0.0 should be at least 1.0.0", actual)
	actual = args.Map{"result": coreversion.IsAtLeast("1.0.0", "2.0.0")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1.0.0 should not be at least 2.0.0", actual)
}

func Test_IsLower(t *testing.T) {
	// Act
	actual := args.Map{"result": coreversion.IsLower("1.0.0", "2.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.0.0 should be lower than 2.0.0", actual)
}

func Test_IsLowerOrEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": coreversion.IsLowerOrEqual("1.0.0", "1.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "equal should be lower or equal", actual)
	actual = args.Map{"result": coreversion.IsLowerOrEqual("1.0.0", "2.0.0")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1.0.0 should be lower or equal to 2.0.0", actual)
}

func Test_IsExpectedVersion(t *testing.T) {
	// just exercise the function
	_ = coreversion.IsExpectedVersion(
		coreversion.CompareVersionString("1.0.0", "1.0.0"),
		"1.0.0",
		"1.0.0",
	)
}

// ==========================================
// VersionsCollection
// ==========================================

func Test_VersionsCollection_Basic(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	vc.Add("2.0.0")

	// Act
	actual := args.Map{"result": vc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": vc.Count() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Count should equal Length", actual)
	actual = args.Map{"result": vc.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	actual = args.Map{"result": vc.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
	actual = args.Map{"result": vc.LastIndex() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected last index 1", actual)
	actual = args.Map{"result": vc.HasIndex(1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have index 1", actual)
	actual = args.Map{"result": vc.HasIndex(5)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have index 5", actual)
}

func Test_VersionsCollection_AddSkipInvalid(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.AddSkipInvalid("1.0.0")
	vc.AddSkipInvalid("")

	// Act
	actual := args.Map{"result": vc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_VersionsCollection_AddVersionsRaw(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.AddVersionsRaw("1.0.0", "2.0.0")

	// Act
	actual := args.Map{"result": vc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_VersionsCollection_Strings(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	compactStrs := vc.VersionCompactStrings()

	// Act
	actual := args.Map{"result": len(compactStrs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	verStrs := vc.VersionsStrings()
	actual = args.Map{"result": len(verStrs) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_VersionsCollection_Strings_Empty(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}

	// Act
	actual := args.Map{"result": len(vc.VersionCompactStrings()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty", actual)
	actual = args.Map{"result": len(vc.VersionsStrings()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty", actual)
}

func Test_VersionsCollection_IndexOf(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	vc.Add("2.0.0")

	// Act
	actual := args.Map{"result": vc.IndexOf("2.0.0") < 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should find version", actual)
	actual = args.Map{"result": vc.IndexOf("3.0.0") >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not find version", actual)
}

func Test_VersionsCollection_IsContainsVersion(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")

	// Act
	actual := args.Map{"result": vc.IsContainsVersion("1.0.0")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain version", actual)
}

func Test_VersionsCollection_IsEqual(t *testing.T) {
	// Arrange
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("1.0.0")

	// Act
	actual := args.Map{"result": vc1.IsEqual(vc2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_VersionsCollection_IsEqual_DifferentLength(t *testing.T) {
	// Arrange
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}

	// Act
	actual := args.Map{"result": vc1.IsEqual(vc2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different lengths should not be equal", actual)
}

func Test_VersionsCollection_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var vc1 *coreversion.VersionsCollection
	var vc2 *coreversion.VersionsCollection

	// Act
	actual := args.Map{"result": vc1.IsEqual(vc2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_VersionsCollection_IsEqual_OneNil(t *testing.T) {
	// Arrange
	var vc1 *coreversion.VersionsCollection
	vc2 := &coreversion.VersionsCollection{}

	// Act
	actual := args.Map{"result": vc1.IsEqual(vc2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should not be equal", actual)
}

func Test_VersionsCollection_String(t *testing.T) {
	// Arrange
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")

	// Act
	actual := args.Map{"result": vc.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_VersionsCollection_Length_Nil(t *testing.T) {
	// Arrange
	var vc *coreversion.VersionsCollection

	// Act
	actual := args.Map{"result": vc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

// ==========================================
// Version comparison methods
// ==========================================

func Test_Version_MajorBuild(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.0.0.5")
	cmp := v.MajorBuild(3, 5)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_IsMajorBuildAtLeast(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.0.0.5")

	// Act
	actual := args.Map{"result": v.IsMajorBuildAtLeast(3, 4)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3.0.0.5 should be at least 3.0.0.4", actual)
}

func Test_Version_MajorBuildString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.0.0.5")
	cmp := v.MajorBuildString("3", "5")

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_MajorMinorPatchBuildString(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("3.2.1.5")
	cmp := v.MajorMinorPatchBuildString("3", "2", "5", "1")

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_Patch(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	cmp := v.Patch(3)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_MajorPatch(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3")
	cmp := v.MajorPatch(1, 3)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_Build(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")
	cmp := v.Build(4)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_MajorMinorPatchBuild(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")
	cmp := v.MajorMinorPatchBuild(1, 2, 3, 4)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_Compare(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("1.2.3")
	v2 := coreversion.New.Create("1.2.3")
	cmp := v1.Compare(&v2)

	// Act
	actual := args.Map{"result": cmp.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Version_IsEqual(t *testing.T) {
	// Arrange
	v1 := coreversion.New.Create("1.2.3")
	v2 := coreversion.New.Create("1.2.3")

	// Act
	actual := args.Map{"result": v1.IsEqual(&v2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}
