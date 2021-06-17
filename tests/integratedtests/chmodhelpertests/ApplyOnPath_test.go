package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_ApplyOnPath(t *testing.T) {
	coretests.SkipOnWindows(t)

	for _, testCase := range chmodhelpertestwrappers.RwxInstructionsApplyTestCases {
		// Arrange
		caseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(
			caseMessenger)
		// expected := testCase.ExpectedAsRwxOwnerGroupOtherInstruction()
		createDefaultPaths(&testCase.CreatePaths)

		// Act
		actualErr := applyPathInstructions(&testCase)
		testCase.SetActual(actualErr)

		// Assert
		Convey(testHeader, t, func() {
			So(actualErr, ShouldBeNil)
		})

		assertTestCaseChmodAsExpected(
			t,
			&testCase,
			testHeader)
	}
}
