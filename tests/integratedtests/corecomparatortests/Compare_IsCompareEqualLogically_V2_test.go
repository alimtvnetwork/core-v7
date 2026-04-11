package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Compare_IsCompareEqualLogically_Verification(t *testing.T) {
	for caseIndex, testCase := range compareLogicallyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftVal, _ := input.GetAsInt("left")
		rightVal, _ := input.GetAsInt("right")
		left := corecomparator.Compare(leftVal)
		right := corecomparator.Compare(rightVal)

		// Act
		result := left.IsCompareEqualLogically(right)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_IsAnyOf_Verification(t *testing.T) {
	for caseIndex, testCase := range compareIsAnyOfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(val)

		// Act
		result := compare.IsAnyOf(corecomparator.Equal, corecomparator.LeftLess)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_IsAnyOf_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range compareIsAnyOfEmptyTestCases {
		// Arrange
		compare := corecomparator.LeftGreater

		// Act
		result := compare.IsAnyOf()

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_NameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range compareNameValueTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(val)

		// Act
		result := compare.NameValue()

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_CsvStrings_Verification(t *testing.T) {
	for caseIndex, testCase := range compareCsvStringsTestCases {
		// Arrange
		compare := corecomparator.Equal

		// Act
		result := compare.CsvStrings(corecomparator.Equal, corecomparator.LeftGreater)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_CsvStrings_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range compareCsvStringsEmptyTestCases {
		// Arrange
		compare := corecomparator.Equal

		// Act
		result := compare.CsvStrings()

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_ValueConversions_Verification(t *testing.T) {
	for caseIndex, testCase := range compareValueConversionsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(val)

		// Act
		actual := args.Map{
			"valueByte":         int(compare.ValueByte()),
			"valueInt":          compare.ValueInt(),
			"toNumberString":    compare.ToNumberString(),
			"numberString":      compare.NumberString(),
			"numberJsonString":  compare.NumberJsonString(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_MarshalJSON_Verification(t *testing.T) {
	for caseIndex, testCase := range compareMarshalJsonTestCases {
		// Arrange
		compare := corecomparator.Equal

		// Act
		bytes, err := compare.MarshalJSON()

		actual := args.Map{
			"hasError": err != nil,
			"notEmpty": len(bytes) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_OnlySupportedErr_Verification(t *testing.T) {
	for caseIndex, testCase := range compareOnlySupportedErrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(val)

		// Act
		err := compare.OnlySupportedErr("test message", corecomparator.Equal, corecomparator.LeftLess)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_OnlySupportedDirectErr_Verification(t *testing.T) {
	for caseIndex, testCase := range compareOnlySupportedDirectErrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(val)

		// Act
		err := compare.OnlySupportedDirectErr(corecomparator.Equal, corecomparator.LeftLess)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_OnlySupportedErr_EmptyMessage_Verification(t *testing.T) {
	for caseIndex, testCase := range compareOnlySupportedEmptyMsgTestCases {
		// Arrange
		compare := corecomparator.LeftGreater

		// Act
		err := compare.OnlySupportedErr("", corecomparator.Equal)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinLength_Verification(t *testing.T) {
	for caseIndex, testCase := range minLengthTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsInt("left")
		right, _ := input.GetAsInt("right")

		// Act
		result := corecomparator.MinLength(left, right)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_IsAnyNamesOf_Verification(t *testing.T) {
	for caseIndex, testCase := range compareIsAnyNamesOfTestCases {
		// Arrange
		compare := corecomparator.Equal

		// Act
		result := compare.IsAnyNamesOf("Equal", "LeftGreater")

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_IsInconclusiveOrNotEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range compareIsInconclusiveOrNotEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(val)

		// Act
		result := compare.IsInconclusiveOrNotEqual()

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
