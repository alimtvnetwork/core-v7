package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/errcore"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxChmodUsingRwxInstructions_Unix(t *testing.T) {
	// coretests.SkipOnWindows(t)
	// Setup
	createPathInstructions := chmodhelpertestwrappers.CreatePathInstruction2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		&createPathInstructions)

	for caseIndex, testCase := range chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsTestCases {
		// Arrange
		expectationMessage := testCase.ExpectedErrorMessage
		executor, err := chmodhelper.ParseRwxInstructionToExecutor(&testCase.RwxInstruction)

		errcore.SimpleHandleErr(err, "")

		// Act
		actualErr := executor.VerifyRwxModifiersDirect(
			false,
			testCase.Locations...)

		expectation := &errcore.ExpectationMessageDef{
			CaseIndex:      caseIndex,
			FuncName:       "Test_VerifyRwxChmodUsingRwxInstructions_Unix",
			TestCaseName:   "VerifyRwxChmodUsingRwxInstructionsTestCases",
			When:           testCase.Header,
			Expected:       expectationMessage,
			IsNonWhiteSort: true,
		}

		// Assert
		Convey(testCase.Header, t, func() {
			isEqual := coretests.IsErrorNonWhiteSortedEqual(
				true,
				actualErr,
				expectation)

			So(isEqual, ShouldBeTrue)
		})
	}
}
