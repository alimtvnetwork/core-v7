package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/errcore"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxChmodUsingRwxInstructions_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Setup
	createPathInstructions := pathInstructionsV2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		createPathInstructions,
	)

	for caseIndex, testCase := range chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsTestCases {
		// Arrange
		expectationMessage := testCase.ExpectedErrorMessage
		executor, err := chmodhelper.ParseRwxInstructionToExecutor(&testCase.RwxInstruction)

		errcore.SimpleHandleErr(err, "")

		// Act
		actualErr := executor.VerifyRwxModifiersDirect(
			false,
			testCase.Locations...,
		)

		expectation := &errcore.ExpectationMessageDef{
			CaseIndex:      caseIndex,
			FuncName:       "Test_VerifyRwxChmodUsingRwxInstructions_Unix",
			TestCaseName:   "VerifyRwxChmodUsingRwxInstructionsTestCases",
			When:           testCase.Header,
			Expected:       expectationMessage,
			IsNonWhiteSort: true,
		}

		// Assert
		Convey(
			testCase.Header, t, func() {
				isEqual := coretests.IsErrorNonWhiteSortedEqual(
					true,
					actualErr,
					expectation,
				)

				So(isEqual, ShouldBeTrue)
			},
		)
	}
}
