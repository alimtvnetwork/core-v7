package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreversion"
)

// ── Version nil receiver methods ──

func Test_Version_NilReceiver(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{
		"versionDisplay":  v.VersionDisplay(),
		"compiledVersion": v.CompiledVersion(),
		"majorString":     v.MajorString(),
		"minorString":     v.MinorString(),
		"patchString":     v.PatchString(),
		"buildString":     v.BuildString(),
		"hasMajor":        v.HasMajor(),
		"hasMinor":        v.HasMinor(),
		"hasPatch":        v.HasPatch(),
		"hasBuild":        v.HasBuild(),
		"isMajorInvalid":  v.IsMajorInvalid(),
		"isMinorInvalid":  v.IsMinorInvalid(),
		"isPatchInvalid":  v.IsPatchInvalid(),
		"isBuildInvalid":  v.IsBuildInvalid(),
		"isEmptyOrInvalid": v.IsEmptyOrInvalid(),
		"hasAnyItem":      v.HasAnyItem(),
		"isDefined":       v.IsDefined(),
		"displayMajor":    v.VersionDisplayMajor(),
	}

	// Assert
	expected := args.Map{
		"versionDisplay":  "",
		"compiledVersion": "",
		"majorString":     "",
		"minorString":     "",
		"patchString":     "",
		"buildString":     "",
		"hasMajor":        false,
		"hasMinor":        false,
		"hasPatch":        false,
		"hasBuild":        false,
		"isMajorInvalid":  true,
		"isMinorInvalid":  true,
		"isPatchInvalid":  true,
		"isBuildInvalid":  true,
		"isEmptyOrInvalid": true,
		"hasAnyItem":      false,
		"isDefined":       false,
		"displayMajor":    "",
	}
	expected.ShouldBeEqual(t, 0, "Version returns nil -- nil receiver", actual)
}

// ── Version nil VersionCompareEqual ──

func Test_Version_IsVersionCompareEqual_Nil_FromVersionNilReceiver(t *testing.T) {
	// Arrange
	var v *coreversion.Version

	// Act
	actual := args.Map{
		"bothEmpty":   v.IsVersionCompareEqual(""),
		"nilNonEmpty": v.IsVersionCompareEqual("1.0.0"),
	}

	// Assert
	expected := args.Map{
		"bothEmpty": true,
		"nilNonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsVersionCompareEqual returns nil -- nil", actual)
}

// ── Version IsMajorInvalidOrZero / IsMinorInvalidOrZero / IsPatchInvalidOrZero / IsBuildInvalidOrZero ──

func Test_Version_InvalidOrZero(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")
	zeroV := coreversion.New.Create("0.0.0.0")

	// Act
	actual := args.Map{
		"majorNotZero": v.IsMajorInvalidOrZero(),
		"minorNotZero": v.IsMinorInvalidOrZero(),
		"patchNotZero": v.IsPatchInvalidOrZero(),
		"buildNotZero": v.IsBuildInvalidOrZero(),
		"zeroMajor":    zeroV.IsMajorInvalidOrZero(),
		"zeroMinor":    zeroV.IsMinorInvalidOrZero(),
		"zeroPatch":    zeroV.IsPatchInvalidOrZero(),
		"zeroBuild":    zeroV.IsBuildInvalidOrZero(),
	}

	// Assert
	expected := args.Map{
		"majorNotZero": false,
		"minorNotZero": false,
		"patchNotZero": false,
		"buildNotZero": false,
		"zeroMajor":    true,
		"zeroMinor":    true,
		"zeroPatch":    true,
		"zeroBuild":    true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidOrZero returns error -- with args", actual)
}

// ── Version String / AllVersionValues / AllValidVersionValues ──

func Test_Version_String_And_AllValues(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")

	// Act
	actual := args.Map{
		"stringNotEmpty":   v.String() != "",
		"allValuesLen":     len(v.AllVersionValues()),
		"allValidLen":      len(v.AllValidVersionValues()),
	}

	// Assert
	expected := args.Map{
		"stringNotEmpty":   true,
		"allValuesLen":     4,
		"allValidLen":      4,
	}
	expected.ShouldBeEqual(t, 0, "String returns non-empty -- and AllValues", actual)
}

// ── Version comparison methods ──

func Test_Version_MajorMinorPatchBuild_Comparisons(t *testing.T) {
	// Arrange
	v := coreversion.New.Create("1.2.3.4")

	// Act
	actual := args.Map{
		"majorAtLeast":        v.IsMajorAtLeast(1),
		"majorStringAtLeast":  v.IsMajorStringAtLeast("1"),
		"majorMinorAtLeast":   v.IsMajorMinorAtLeast(1, 2),
		"majorBuildAtLeast":   v.IsMajorBuildAtLeast(1, 4),
		"majorMinorPatchAtLeast": v.IsMajorMinorPatchAtLeast(1, 2, 3),
		"isNotEqual":          v.IsVersionCompareNotEqual("2.0.0"),
		"patch":               v.Patch(3).IsEqual(),
		"majorPatch":          v.MajorPatch(1, 3).IsEqual(),
		"majorBuild":          v.MajorBuild(1, 4).IsEqual(),
		"majorBuildString":    v.MajorBuildString("1", "4").IsEqual(),
		"majorMinorPatchBuildString": v.MajorMinorPatchBuildString("1", "2", "4", "3").IsEqual(),
	}

	// Assert
	expected := args.Map{
		"majorAtLeast":        true,
		"majorStringAtLeast":  true,
		"majorMinorAtLeast":   true,
		"majorBuildAtLeast":   true,
		"majorMinorPatchAtLeast": true,
		"isNotEqual":          true,
		"patch":               true,
		"majorPatch":          true,
		"majorBuild":          true,
		"majorBuildString":    true,
		"majorMinorPatchBuildString": true,
	}
	expected.ShouldBeEqual(t, 0, "Comparison returns correct value -- methods", actual)
}

// ── Package-level compare functions ──

func Test_CompareVersionString_FromVersionNilReceiver(t *testing.T) {
	// Act
	actual := args.Map{
		"equal": coreversion.CompareVersionString("1.0.0", "1.0.0").IsEqual(),
		"less":  coreversion.IsLower("1.0.0", "2.0.0"),
		"lessOrEq": coreversion.IsLowerOrEqual("1.0.0", "1.0.0"),
		"atLeast": coreversion.IsAtLeast("2.0.0", "1.0.0"),
		"expected": coreversion.IsExpectedVersion(1, "1.0.0", "1.0.0"),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"less":  true,
		"lessOrEq": true,
		"atLeast": true,
		"expected": true,
	}
	expected.ShouldBeEqual(t, 0, "Package returns correct value -- compare functions", actual)
}

// ── VersionsCollection — uncovered methods ──

func Test_VersionsCollection_AddSkipInvalid_FromVersionNilReceiver(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}

	// Act
	vc.AddSkipInvalid("1.0.0")
	vc.AddSkipInvalid("")
	vc.AddSkipInvalid("2.0.0")

	// Assert
	actual := args.Map{"length": vc.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "AddSkipInvalid returns error -- with args", actual)
}

func Test_VersionsCollection_Methods(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	vc.Add("2.0.0")

	// Act
	actual := args.Map{
		"count":         vc.Count(),
		"isEmpty":       vc.IsEmpty(),
		"hasAnyItem":    vc.HasAnyItem(),
		"lastIndex":     vc.LastIndex(),
		"hasIndex0":     vc.HasIndex(0),
		"hasIndex5":     vc.HasIndex(5),
		"compactLen":    len(vc.VersionCompactStrings()),
		"versionsLen":   len(vc.VersionsStrings()),
		"indexOf":       vc.IndexOf("1.0.0") >= 0,
		"contains":      vc.IsContainsVersion("1.0.0"),
		"notContains":   vc.IsContainsVersion("9.9.9"),
		"stringNotEmpty": vc.String() != "",
	}

	// Assert
	expected := args.Map{
		"count":         2,
		"isEmpty":       false,
		"hasAnyItem":    true,
		"lastIndex":     1,
		"hasIndex0":     true,
		"hasIndex5":     false,
		"compactLen":    2,
		"versionsLen":   2,
		"indexOf":       true,
		"contains":      true,
		"notContains":   false,
		"stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "VersionsCollection returns correct value -- methods", actual)
}

func Test_VersionsCollection_NilReceiver(t *testing.T) {
	// Arrange
	var vc *coreversion.VersionsCollection

	// Act
	actual := args.Map{
		"length": vc.Length(),
		"isEmpty": vc.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"length": 0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "VersionsCollection returns nil -- nil receiver", actual)
}

func Test_VersionsCollection_IsEqual_FromVersionNilReceiver(t *testing.T) {
	// Arrange
	var nilVc *coreversion.VersionsCollection
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("1.0.0")

	// Act
	actual := args.Map{
		"bothNil":     nilVc.IsEqual(nil),
		"leftNil":     nilVc.IsEqual(vc1),
		"equal":       vc1.IsEqual(vc2),
		"diffLen":     vc1.IsEqual(&coreversion.VersionsCollection{}),
	}

	// Assert
	expected := args.Map{
		"bothNil":     true,
		"leftNil":     false,
		"equal":       true,
		"diffLen":     false,
	}
	expected.ShouldBeEqual(t, 0, "VersionsCollection returns correct value -- IsEqual", actual)
}

func Test_VersionsCollection_JsonParseSelfInject_FromVersionNilReceiver(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	j := vc.JsonPtr()
	vc2 := &coreversion.VersionsCollection{}

	// Act
	err := vc2.JsonParseSelfInject(j)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"len": vc2.Length(),
	}
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "VersionsCollection returns correct value -- JsonParseSelfInject", actual)
}

func Test_VersionsCollection_AsBinderInterfaces(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}

	// Act
	actual := args.Map{
		"jsonBinder":  vc.AsJsonContractsBinder() != nil,
		"sliceBinder": vc.AsBasicSliceContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"jsonBinder": true,
		"sliceBinder": true,
	}
	expected.ShouldBeEqual(t, 0, "VersionsCollection returns correct value -- AsBinderInterfaces", actual)
}

func Test_VersionsCollection_AddVersionsRaw_FromVersionNilReceiver(t *testing.T) {
	// Arrange
	vc := &coreversion.VersionsCollection{}

	// Act
	vc.AddVersionsRaw("1.0.0", "2.0.0", "3.0.0")

	// Assert
	actual := args.Map{"length": vc.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "AddVersionsRaw returns correct value -- with args", actual)
}
