package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxPartialChmodLocations_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := chmodhelpertestwrappers.CreatePathInstruction2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		&createPathInstructions)
	for caseIndex, testCase := range chmodhelpertestwrappers.VerifyRwxPartialChmodLocationsTestCases {
		header := testCase.Header
		expectationMessage := testCase.ExpectationErrorMessage

		// Act
		err := chmodhelper.VerifyChmodLocationsUsingPartialRwx(
			testCase.IsContinueOnError,
			testCase.IsSkipOnInvalid,
			testCase.ExpectedPartialRwx,
			testCase.Locations)

		expectation := &errcore.ExpectationMessageDef{
			CaseIndex:      caseIndex,
			FuncName:       "Test_VerifyRwxPartialChmodLocations_Unix",
			TestCaseName:   "VerifyRwxPartialChmodLocationsTestCases",
			When:           testCase.Header,
			Expected:       expectationMessage,
			IsNonWhiteSort: true,
		}

		// Assert
		Convey(header, t, func() {
			isEqual := coretests.IsErrorNonWhiteSortedEqual(
				true,
				err,
				expectation)

			So(isEqual, ShouldBeTrue)
		})
	}
}
