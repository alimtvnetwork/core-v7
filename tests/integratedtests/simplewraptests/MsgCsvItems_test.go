package simplewraptests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_MsgCsvItems_Verification(t *testing.T) {
	for caseIndex, testCase := range msgCsvItemsTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		title := inputs[0].(string)
		csvItems := inputs[1].([]interface{})

		// Act
		actualSlice.Add(
			simplewrap.MsgCsvItems(
				title,
				csvItems...,
			),
		)

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
