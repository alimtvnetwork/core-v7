package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov5_VersionDisplayMajorMinorPatch_InvalidPatch tests the IsPatchInvalid branch.
func Test_VersionDisplayMajorMinorPatch_InvalidPatch(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 1,
		VersionMinor: 2,
		VersionPatch: -1,
	}

	// Act
	actual := v.VersionDisplayMajorMinorPatch()

	// Assert
	expected := "v1.2"
	actual := args.Map{"result": actual != expected}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "VersionDisplayMajorMinorPatch with invalid patch: got, want", actual)
}

// Test_Cov5_Major_Compare tests the Major() comparison method.
func Test_Major_Compare(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 3,
		VersionMinor: 0,
	}

	// Act
	result := v.Major(3)

	// Assert
	actual := args.Map{"result": result.IsEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Major(3) on version with major=3 should be Equal", actual)
}
