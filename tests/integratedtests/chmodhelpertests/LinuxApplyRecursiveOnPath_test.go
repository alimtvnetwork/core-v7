package chmodhelpertests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
)

func Test_LinuxApplyRecursiveOnPath_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	for _, testCase := range rwxInstructionsUnixApplyRecursivelyTestCases {
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
		actualErr := linuxApplyRecursivePathInstructions(&testCase)
		testCase.SetActual(actualErr)

		// Assert
		convey.Convey(
			testHeader, t, func() {
				convey.So(actualErr, convey.ShouldBeNil)
			},
		)

		assertTestCaseChmodAsExpected(
			t,
			&testCase,
			testHeader,
		)
	}
}
