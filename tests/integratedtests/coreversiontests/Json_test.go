package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreversion"
)

func Test_Json_Verification(t *testing.T) {
	for caseIndex, testCase := range jsonTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]string)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		creatorFunc := coreversion.New.Default

		// Act
		for index, input := range inputs {
			v := creatorFunc(input)
			toJson := v.JsonPtr().JsonString()

			actualSlice.AppendFmt(
				jsonFmt,
				index,
				input,
				toJson,
			)
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
