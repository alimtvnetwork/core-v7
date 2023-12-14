package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/isany"
)

func Test_SliceValidator(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorTestCases {
		// Arrange
		inputs := testCase.
			Case.
			ArrangeInput.([]args.Two)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				"%d : %t (%s, %s)",
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
			)
		}

		actLines := actualSlice.Strings()
		actualError := testCase.Case.VerifyAllEqual(
			caseIndex,
			actLines...,
		)
		validator := testCase.Validator
		errLines := errcore.ErrorToSplitLines(actualError)

		// Assert
		validator.AssertAllQuick(
			t,
			caseIndex,
			testCase.Case.Title,
			errLines...,
		)
	}
}

func Test_SliceValidator_FirstError(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorFirstErrorTestCases {
		// Arrange
		inputs := testCase.
			Case.
			ArrangeInput.([]args.Two)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				"%d : %t (%s, %s)",
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
			)
		}

		actLines := actualSlice.Strings()
		actualError := testCase.Case.VerifyFirst(
			caseIndex,
			stringcompareas.Equal,
			actLines,
		)
		validator := testCase.Validator
		errLines := errcore.ErrorToSplitLines(actualError)

		// Assert
		validator.AssertAllQuick(
			t,
			caseIndex,
			testCase.Case.Title,
			errLines...,
		)
	}
}
