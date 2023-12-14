package corecsvtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_StringersToString_Verification(t *testing.T) {
	for caseIndex, testCase := range stringersTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]fmt.Stringer)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		actualSlice.Add(
			corecsv.StringersToString(
				constants.CommaSpace,
				true,
				true,
				inputs...,
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
