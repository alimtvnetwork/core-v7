package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestVersion_Creation verifies version creation from various formats.
func TestVersion_Creation(t *testing.T) {
	for _, tc := range versionCreationCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			v := coreversion.New.Default(tc.input)

			// Assert
			actual := args.Map{"result": v.VersionMajor != tc.expectedMajor}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "major: expected", actual)
			actual := args.Map{"result": v.VersionMinor != tc.expectedMinor}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "minor: expected", actual)
			actual := args.Map{"result": v.VersionPatch != tc.expectedPatch}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "patch: expected", actual)
		})
	}
}

// TestVersion_Display verifies display methods.
func TestVersion_Display(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	actual := args.Map{"result": v.VersionDisplay() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "VersionDisplay should not be empty", actual)
	actual := args.Map{"result": v.CompiledVersion() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CompiledVersion should not be empty", actual)
	actual := args.Map{"result": v.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)
	actual := args.Map{"result": v.VersionDisplayMajor() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "VersionDisplayMajor should not be empty", actual)
	actual := args.Map{"result": v.VersionDisplayMajorMinor() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "VersionDisplayMajorMinor should not be empty", actual)
	actual := args.Map{"result": v.VersionDisplayMajorMinorPatch() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "VersionDisplayMajorMinorPatch should not be empty", actual)
}

// TestVersion_NilDisplay verifies nil pointer display.
func TestVersion_NilDisplay(t *testing.T) {
	var v *coreversion.Version
	actual := args.Map{"result": v.VersionDisplay() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil VersionDisplay should be empty", actual)
	actual := args.Map{"result": v.CompiledVersion() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil CompiledVersion should be empty", actual)
	actual := args.Map{"result": v.MajorString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil MajorString should be empty", actual)
	actual := args.Map{"result": v.MinorString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil MinorString should be empty", actual)
	actual := args.Map{"result": v.PatchString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil PatchString should be empty", actual)
	actual := args.Map{"result": v.BuildString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil BuildString should be empty", actual)
}

// TestVersion_HasMethods verifies Has/Invalid checks.
func TestVersion_HasMethods(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	actual := args.Map{"result": v.HasMajor()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have major", actual)
	actual := args.Map{"result": v.HasMinor()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have minor", actual)
	actual := args.Map{"result": v.HasPatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have patch", actual)
	actual := args.Map{"result": v.IsMajorInvalid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "major should be valid", actual)
	actual := args.Map{"result": v.IsMinorInvalid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "minor should be valid", actual)
	actual := args.Map{"result": v.IsPatchInvalid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "patch should be valid", actual)
}

// TestVersion_NilHas verifies nil receiver for Has methods.
func TestVersion_NilHas(t *testing.T) {
	var v *coreversion.Version
	actual := args.Map{"result": v.HasMajor()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have major", actual)
	actual := args.Map{"result": v.HasMinor()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have minor", actual)
	actual := args.Map{"result": v.HasPatch()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have patch", actual)
	actual := args.Map{"result": v.HasBuild()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have build", actual)
	actual := args.Map{"result": v.IsMajorInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil major should be invalid", actual)
}

// TestVersion_Empty verifies empty version.
func TestVersion_Empty(t *testing.T) {
	v := coreversion.New.Default("")
	actual := args.Map{"result": v.IsEmptyOrInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be invalid", actual)
	actual := args.Map{"result": v.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have items", actual)
	actual := args.Map{"result": v.IsDefined()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not be defined", actual)
}

// TestVersion_Comparison verifies comparison methods.
func TestVersion_Comparison(t *testing.T) {
	v1 := coreversion.New.Default("v1.2.3")
	v2 := coreversion.New.Default("v1.2.3")
	v3 := coreversion.New.Default("v2.0.0")

	actual := args.Map{"result": v1.IsEqual(&v2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v1 should equal v2", actual)
	actual := args.Map{"result": v1.IsEqual(&v3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "v1 should not equal v3", actual)
	actual := args.Map{"result": v1.IsLeftLessThan(&v3)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v1 should be less than v3", actual)
	actual := args.Map{"result": v3.IsLeftGreaterThan(&v1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v3 should be greater than v1", actual)
	actual := args.Map{"result": v1.IsLeftLessThanOrEqual(&v2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v1 should be <= v2", actual)
	actual := args.Map{"result": v1.IsLeftGreaterThanOrEqual(&v2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v1 should be >= v2", actual)
}

// TestVersion_AtLeast verifies AtLeast.
func TestVersion_AtLeast(t *testing.T) {
	v := coreversion.New.Default("v2.1.0")
	actual := args.Map{"result": v.IsAtLeast("v2.0.0")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v2.1.0 should be at least v2.0.0", actual)
	actual := args.Map{"result": v.IsAtLeast("v3.0.0")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "v2.1.0 should not be at least v3.0.0", actual)
}

// TestVersion_IsEqualVersionString verifies string comparison.
func TestVersion_IsEqualVersionString(t *testing.T) {
	v := coreversion.New.Default("v1.0.0")
	actual := args.Map{"result": v.IsEqualVersionString("v1.0.0")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
	actual := args.Map{"result": v.IsEqualVersionString("1.0.0")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal without v prefix", actual)
}

// TestVersion_IsLowerVersionString verifies lower version string.
func TestVersion_IsLowerVersionString(t *testing.T) {
	v := coreversion.New.Default("v1.0.0")
	actual := args.Map{"result": v.IsLowerVersionString("v2.0.0")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v1 should be lower than v2", actual)
}

// TestVersion_IsLowerEqualVersionString verifies lower or equal.
func TestVersion_IsLowerEqualVersionString(t *testing.T) {
	v := coreversion.New.Default("v1.0.0")
	actual := args.Map{"result": v.IsLowerEqualVersionString("v1.0.0")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "v1 should be <= v1", actual)
}

// TestVersion_IsVersionCompareEqual verifies compact comparison.
func TestVersion_IsVersionCompareEqual(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	actual := args.Map{"result": v.IsVersionCompareEqual("1.2.3")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
	actual := args.Map{"result": v.IsVersionCompareNotEqual("2.0.0")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should not be equal", actual)

	var nilV *coreversion.Version
	actual := args.Map{"result": nilV.IsVersionCompareEqual("")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil with empty should be equal", actual)
	actual := args.Map{"result": nilV.IsVersionCompareEqual("1.0")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil with non-empty should not be equal", actual)
}

// TestVersion_Clone verifies clone.
func TestVersion_Clone(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	cloned := v.Clone()
	actual := args.Map{"result": cloned.VersionMajor != v.VersionMajor}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should have same major", actual)
}

// TestVersion_ClonePtr verifies pointer clone.
func TestVersion_ClonePtr(t *testing.T) {
	var nilV *coreversion.Version
	actual := args.Map{"result": nilV.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should be nil", actual)

	v := coreversion.New.Default("v1.0")
	ptr := v.ClonePtr()
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should not be nil", actual)
}

// TestVersion_AllVersionValues verifies all version values.
func TestVersion_AllVersionValues(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	vals := v.AllVersionValues()
	actual := args.Map{"result": len(vals) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have values", actual)
}

// TestVersion_AllValidVersionValues verifies valid version values.
func TestVersion_AllValidVersionValues(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	vals := v.AllValidVersionValues()
	actual := args.Map{"result": len(vals) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have valid values", actual)
}

// TestVersion_IsMajorAtLeast verifies major at least.
func TestVersion_IsMajorAtLeast(t *testing.T) {
	v := coreversion.New.Default("v2.0.0")
	actual := args.Map{"result": v.IsMajorAtLeast(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2 should be at least 1", actual)
	actual := args.Map{"result": v.IsMajorAtLeast(2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2 should be at least 2", actual)
	actual := args.Map{"result": v.IsMajorAtLeast(3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "2 should not be at least 3", actual)
}

// TestVersion_IsMajorMinorAtLeast verifies major.minor at least.
func TestVersion_IsMajorMinorAtLeast(t *testing.T) {
	v := coreversion.New.Default("v2.3.0")
	actual := args.Map{"result": v.IsMajorMinorAtLeast(2, 3)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2.3 should be at least 2.3", actual)
	actual := args.Map{"result": v.IsMajorMinorAtLeast(2, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2.3 should be at least 2.2", actual)
	actual := args.Map{"result": v.IsMajorMinorAtLeast(2, 4)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "2.3 should not be at least 2.4", actual)
}

// TestNewCreator_Spread verifies spread constructors.
func TestNewCreator_Spread(t *testing.T) {
	v := coreversion.New.SpreadIntegers(1, 2, 3)
	actual := args.Map{"result": v.VersionMajor != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "major should be 1", actual)

	v = coreversion.New.SpreadStrings("1", "2")
	actual := args.Map{"result": v.VersionMajor != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "major should be 1", actual)

	v = coreversion.New.SpreadBytes(1, 2, 3)
	actual := args.Map{"result": v.VersionMajor != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "major should be 1", actual)

	v = coreversion.New.SpreadUnsignedIntegers(1, 2)
	actual := args.Map{"result": v.VersionMajor != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "major should be 1", actual)
}

// TestNewCreator_AllVariants verifies all constructors.
func TestNewCreator_AllVariants(t *testing.T) {
	v := coreversion.New.MajorMinor("1", "2")
	actual := args.Map{"result": v.VersionMinor != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "minor should be 2", actual)

	v = coreversion.New.MajorMinorPatch("1", "2", "3")
	actual := args.Map{"result": v.VersionPatch != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "patch should be 3", actual)

	v = coreversion.New.MajorMinorPatchBuild("1", "2", "3", "4")
	actual := args.Map{"result": v.VersionBuild != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "build should be 4", actual)

	v = coreversion.New.AllInt(1, 2, 3, 4)
	actual := args.Map{"result": v.VersionBuild != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllInt build should be 4", actual)

	v = coreversion.New.AllByte(1, 2, 3, 4)
	actual := args.Map{"result": v.VersionBuild != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllByte build should be 4", actual)

	v = coreversion.New.MajorMinorInt(1, 2)
	actual := args.Map{"result": v.VersionMinor != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MajorMinorInt minor should be 2", actual)

	v = coreversion.New.MajorMinorPatchInt(1, 2, 3)
	actual := args.Map{"result": v.VersionPatch != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MajorMinorPatchInt patch should be 3", actual)
}

// TestVersion_Json verifies JSON serialization.
func TestVersion_Json(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	j := v.Json()
	actual := args.Map{"result": j.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "json should not have error", actual)

	jp := v.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "json ptr should not be nil", actual)
}

// TestVersion_IsMajorStringAtLeast verifies string-based comparison.
func TestVersion_IsMajorStringAtLeast(t *testing.T) {
	v := coreversion.New.Default("v3.0.0")
	actual := args.Map{"result": v.IsMajorStringAtLeast("2")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3 should be at least 2", actual)
}

// TestVersion_InvalidOrZero verifies invalid-or-zero checks.
func TestVersion_InvalidOrZero(t *testing.T) {
	v := coreversion.New.Default("v0.0.0")
	actual := args.Map{"result": v.IsMajorInvalidOrZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be invalid or zero", actual)
	actual := args.Map{"result": v.IsMinorInvalidOrZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be invalid or zero", actual)
	actual := args.Map{"result": v.IsPatchInvalidOrZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be invalid or zero", actual)
	actual := args.Map{"result": v.IsBuildInvalidOrZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be invalid or zero", actual)
}

// TestVersion_NonPtrAndPtr verifies NonPtr/Ptr.
func TestVersion_NonPtrAndPtr(t *testing.T) {
	v := coreversion.New.Default("v1.0")
	np := v.NonPtr()
	actual := args.Map{"result": np.VersionMajor != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NonPtr major should be 1", actual)
	p := v.Ptr()
	actual := args.Map{"result": p == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Ptr should not be nil", actual)
}
