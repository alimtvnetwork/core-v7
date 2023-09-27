package enumimpltests

import (
	"testing"
	
	. "github.com/smartystreets/goconvey/convey"
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
			arrangeInput.Right)
		
		// Assert
		testCase.ShouldBe(
			caseIndex,
			t,
			ShouldResemble,
			diffJsonMessage)
	}
}
