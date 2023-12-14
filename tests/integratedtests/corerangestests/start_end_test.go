package corerangestests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/errcore"
)

func Test_StartEndRanges_ValidCases(t *testing.T) {
	for _, testCase := range validStartEndRangesTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.StartEndInt)
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				convey.So(
					actualRanges,
					should.Equal,
					testCase.ExpectedInput,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}

func Test_StartEndString_Functions_Result_Verification(t *testing.T) {
	for caseIndex, testCase := range startEndRangesStringFunctionsVerificationTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.StartEndInt)
		sliceValidator := testCase.Validator
		slice := corestr.New.SimpleSlice.Cap(100)

		for i, input := range arrangeInputs {
			slice.AppendFmt("StartEnd : %s", input.String())

			slice.AppendFmt(
				"    [%d] func : %s        | result : %s",
				i,
				"String",
				input.String(),
			)

			slice.AppendFmt(
				"    [%d] func : %s   | result : %s",
				i,
				"StringColon",
				input.StringColon(),
			)

			slice.AppendFmt(
				"    [%d] func : %s  | result : %s",
				i,
				"StringHyphen",
				input.StringHyphen(),
			)

			slice.AppendFmt(
				"    [%d] func : %s   | result : %s",
				i,
				"StringSpace",
				input.StringSpace(),
			)
		}

		// Act
		actual := slice.Strings()
		testCase.SetActual(actual)
		sliceValidator.SetActual(actual)

		nextBaseParam := corevalidator.Parameter{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam,
		)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				errcore.PrintErrorWithTestIndex(
					caseIndex,
					testCase.Title,
					validationFinalError,
				)

				convey.So(
					validationFinalError,
					should.BeNil,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}
