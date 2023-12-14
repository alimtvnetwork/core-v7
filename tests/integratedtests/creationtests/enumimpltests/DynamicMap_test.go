package enumimpltests

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
)

func Test_DynamicMapDiff1(t *testing.T) {
	for caseIndex, testCase := range dynamicMapSimpleDiffTestCases {
		// Arrange
		arrangeInput := testCase.ArrangeAsLeftRightDynamicMapWithDefaultChecker()

		// Act
		diffJsonMessage := arrangeInput.Left.ShouldDiffLeftRightMessageUsingDifferChecker(
			arrangeInput.DifferChecker,
			true,
			testCase.CaseTitle(),
			arrangeInput.Right,
		)

		actualLines := strings.Split(
			diffJsonMessage,
			constants.NewLineUnix,
		)

		// Assert
		testCase.ShouldBe(
			caseIndex,
			t,
			ShouldResemble,
			actualLines,
		)
	}
}
