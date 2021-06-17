package chmodhelpertests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_LinuxApplyRecursiveOnPath(t *testing.T) {
	coretests.SkipOnWindows(t)

	for _, testCase := range chmodhelpertestwrappers.RwxInstructionsUnixApplyRecursivelyTestCases {
		// Arrange
		caseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(
			caseMessenger)
		// expected := testCase.ExpectedAsRwxOwnerGroupOtherInstruction()
		createDefaultPaths(&testCase.CreatePaths)

		// Act
		actualErr := linuxApplyRecursivePathInstructions(&testCase)
		testCase.SetActual(actualErr)

		// Assert
		convey.Convey(testHeader, t, func() {
			convey.So(actualErr, convey.ShouldBeNil)
		})

		assertTestCaseChmodAsExpected(
			t,
			&testCase,
			testHeader)
	}
}
