package simplewraptests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_MsgWrapMsg_Wraps_Verification(t *testing.T) {
	for caseIndex, testCase := range msgWrapsMsgTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		firstMsg := inputs[0]
		secondMsg := inputs[1]

		// Act
		actualSlice.Add(
			simplewrap.MsgWrapMsg(
				firstMsg,
				secondMsg,
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
