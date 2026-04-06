package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)
	for caseIndex, testCase := range compareIsMethodTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		otherVal, _ := input.GetAsInt("other")
		compare := corecomparator.Compare(val)
		other := corecomparator.Compare(otherVal)

		// Act
		didPanic := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()
			compare.Format("test")
		}()

		actual := args.Map{
			"is":                    compare.Is(other),
			"isInvalid":            compare.IsInvalid(),
			"isValueEqual":         compare.IsValueEqual(byte(otherVal)),
			"isLeftGreater":        compare.IsLeftGreater(),
			"isLeftGreaterEqual":   compare.IsLeftGreaterEqual(),
			"isLeftLessEqual":      compare.IsLeftLessEqual(),
			"isLeftLessOrLeOrEq":   compare.IsLeftLessOrLessEqualOrEqual(),
			"isDefinedPlus":        compare.IsDefinedPlus(other),
			"isNotInconclusive":    compare.IsNotInconclusive(),
			"rangeNamesCsvNotEmpty": compare.RangeNamesCsv() != "",
			"sqlOpNotEmpty":        compare.SqlOperatorSymbol() != "",
			"stringValueNotEmpty":  compare.StringValue() != "",
			"valueInt8":            compare.ValueInt8(),
			"valueInt16":           compare.ValueInt16(),
			"valueInt32":           compare.ValueInt32(),
			"valueString":         compare.ValueString(),
			"formatPanic":         didPanic,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
	for caseIndex, testCase := range compareUnmarshalJsonTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilData := input.GetAsBoolDefault("isNilData", false)

		var compare corecomparator.Compare
		var err error

		if isNilData {
			// Act
			err = compare.UnmarshalJSON(nil)
		} else {
			data, _ := input.GetAsString("data")
			// Act
			err = compare.UnmarshalJSON([]byte(data))
		}

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RangeNamesCsv_Verification(t *testing.T) {
	// Arrange - no input needed

	// Act
	result := corecomparator.RangeNamesCsv()

	// Assert
	if result == "" {
		t.Error("RangeNamesCsv should not be empty")
	}
}
