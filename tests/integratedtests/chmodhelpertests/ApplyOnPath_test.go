package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_ApplyOnPath_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	for _, testCase := range chmodhelpertestwrappers.RwxInstructionsApplyTestCases {
		// Arrange
		caseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(
			caseMessenger,
		)
		chmodhelper.CreateDirFilesWithRwxPermissionsMust(
			true,
			testCase.CreatePaths,
		)

		// Act
		actualErr := applyPathInstructions(&testCase)
		testCase.SetActual(actualErr)

		// Assert
		Convey(
			testHeader, t, func() {
				So(actualErr, ShouldBeNil)
			},
		)

		assertTestCaseChmodAsExpected(
			t,
			&testCase,
			testHeader,
		)
	}
}
