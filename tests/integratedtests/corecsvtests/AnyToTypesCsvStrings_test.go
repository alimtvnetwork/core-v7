package corecsvtests

import (
	"testing"

	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_AnyToTypesCsvStrings_SingleQuote_Verification(t *testing.T) {
	for caseIndex, testCase := range anyToTypesCsvStringsSingleQuoteTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})

		// Act
		finalActLines := corecsv.AnyToTypesCsvStrings(
			true,
			true,
			inputs...,
		)

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

func Test_AnyToTypesCsvStrings_DoubleQuote_Verification(t *testing.T) {
	for caseIndex, testCase := range anyToTypesCsvStringsDoubleQuoteTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})

		// Act
		finalActLines := corecsv.AnyToTypesCsvStrings(
			true,
			false,
			inputs...,
		)

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

func Test_AnyToTypesCsvStrings_NoQuote_Verification(t *testing.T) {
	for caseIndex, testCase := range anyToTypesCsvStringsNoQuoteTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]interface{})

		// Act
		finalActLines := corecsv.AnyToTypesCsvStrings(
			false,
			false,
			inputs...,
		)

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
