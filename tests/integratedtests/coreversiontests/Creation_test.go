package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreversion"
)

func Test_Creation_Verification(t *testing.T) {
	for caseIndex, testCase := range versionCreationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coreversion.Version)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			if input.IsSafeInvalidCheck() {
				actualSlice.AppendFmt(
					defaultInvalidV2CreationFmt,
					i,
					input.String(),
				)
			} else {
				actualSlice.AppendFmt(
					defaultCreationFmt,
					i,
					input.String(),
					input.VersionCompact,
					input.VersionDisplay(),
				)
			}
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeTrimEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_Creation_UsingString_Verification(t *testing.T) {
	for caseIndex, testCase := range versionCreationUsingStringTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]string)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))
		creatorFunc := coreversion.New.Default

		// Act
		for i, input := range inputs {
			toVersion := creatorFunc(input)

			if toVersion.IsSafeInvalidCheck() {
				actualSlice.AppendFmt(
					defaultInvalidV1CreationFmt,
					i,
					toVersion.String(),
					input,
				)
			} else {
				actualSlice.AppendFmt(
					defaultCreationFmt,
					i,
					toVersion.String(),
					toVersion.VersionCompact,
					toVersion.VersionDisplay(),
				)
			}
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
