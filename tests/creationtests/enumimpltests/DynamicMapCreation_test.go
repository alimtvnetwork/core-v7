package enumimpltests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
)

func Test_DynamicMapCreationDiff(t *testing.T) {
	for caseIndex, testCase := range dynamicMapDiffTestCases {
		// Arrange
		arrangeInput := testCase.ArrangeAsLeftRightDynamicMap()
		diffMap := arrangeInput.Left.DiffRaw(
			true,
			arrangeInput.Right)
		mapAnyDiffer := coredynamic.MapAnyItemDiff(
			arrangeInput.Left)

		// Act
		anotherDiff := mapAnyDiffer.
			DiffRaw(
				true,
				arrangeInput.Right)

		// Assert
		testCase.ShouldBe(
			caseIndex,
			t,
			ShouldResemble,
			diffMap)
		testCase.ShouldBeExplicit(
			false,
			caseIndex,
			t,
			"both diff should be equal",
			diffMap.Raw(),
			ShouldResemble,
			anotherDiff)
	}
}

func Test_DynamicMapCreationDiffMessage(t *testing.T) {
	for caseIndex, testCase := range dynamicMapDiffMessageTestCases {
		// Arrange
		arrangeInput := testCase.ArrangeAsLeftRightDynamicMap()

		// Act
		diffJsonMessage := arrangeInput.Left.ShouldDiffMessage(
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
