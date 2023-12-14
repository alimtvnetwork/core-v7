package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxPartialChmodLocations_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := pathInstructionsV2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		createPathInstructions,
	)
	for caseIndex, testCase := range chmodhelpertestwrappers.VerifyRwxPartialChmodLocationsTestCases {
		header := testCase.Header
		expectationMessage := testCase.ExpectationErrorMessage

		// Act
		err := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(
			testCase.IsContinueOnError,
			testCase.IsSkipOnInvalid,
			testCase.ExpectedPartialRwx,
			testCase.Locations...,
		)

		expectation := &errcore.ExpectationMessageDef{
			CaseIndex:      caseIndex,
			FuncName:       "Test_VerifyRwxPartialChmodLocations_Unix",
			TestCaseName:   "VerifyRwxPartialChmodLocationsTestCases",
			When:           testCase.Header,
			Expected:       expectationMessage,
			IsNonWhiteSort: true,
		}

		// Assert
		Convey(
			header, t, func() {
				isEqual := coretests.IsErrorNonWhiteSortedEqual(
					true,
					err,
					expectation,
				)

				So(isEqual, ShouldBeTrue)
			},
		)
	}
}
