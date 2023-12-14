package corecsvtests

import (
	"testing"

	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_RangeNamesWithValuesIndexes_Verification(t *testing.T) {
	for caseIndex, testCase := range rangeNamesWithValuesIndexesTestCases {
		// Arrange
		inputs := testCase.Arrange()

		// Act
		finalActLines := corecsv.RangeNamesWithValuesIndexes(
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

func Test_RangeNamesWithValuesIndexesCsvString_Verification(t *testing.T) {
	for caseIndex, testCase := range rangeNamesWithValuesIndexesStringTestCases {
		// Arrange
		inputs := testCase.Arrange()

		// Act
		finalActLines := corecsv.RangeNamesWithValuesIndexesCsvString(
			inputs...,
		)

		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines,
		)
	}
}
