package corevalidatortests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/corevalidatortestwrappers"
)

func Test_TestValidators(t *testing.T) {
	for caseIndex, testCase := range corevalidatortestwrappers.TextValidatorsTestCases {
		// Arrange
		paramsBase := corevalidator.ValidatorParamsBase{
			CaseIndex:                         constants.Zero, // fixing test case number here as it is fixed data
			IsIgnoreCompareOnActualInputEmpty: testCase.IsSkipOnContentsEmpty,
			IsAttachUserInputs:                true,
			IsCaseSensitive:                   testCase.IsCaseSensitive,
		}

		err := testCase.Validators.AllVerifyErrorMany(
			&paramsBase,
			testCase.ComparingLines...)

		errorLines := errcore.ErrorToSplitLines(
			err)

		sliceValidator := corevalidator.SliceValidator{
			ActualLines:   errorLines,
			ExpectedLines: testCase.ExpectationLines,
			ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
				IsTrimCompare:        false,
				IsNonEmptyWhitespace: false,
				IsSortStringsBySpace: false,
			},
			CompareAs: stringcompareas.Equal,
		}

		paramsBase2 := corevalidator.ValidatorParamsBase{
			CaseIndex:                         caseIndex,
			IsIgnoreCompareOnActualInputEmpty: false,
			IsAttachUserInputs:                true,
			IsCaseSensitive:                   testCase.IsCaseSensitive,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&paramsBase2)

		isValid := validationFinalError == nil

		// Assert
		convey.Convey(testCase.Header, t, func() {
			errcore.ErrPrintWithTestIndex(caseIndex, validationFinalError)

			convey.So(isValid, convey.ShouldBeTrue)
		})
	}
}
