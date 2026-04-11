package coreappendtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreappend"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_AppendAnyItemsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range appendAnyItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")

		// Act
		result := coreappend.AppendAnyItemsToStringSkipOnNil(
			joiner,
			"suffix",
			"item1", "item2",
		)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAnyItemsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAnyItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")

		// Act
		result := coreappend.PrependAnyItemsToStringSkipOnNil(
			joiner,
			"prefix",
			"item1", "item2",
		)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendAnyItemsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")

		// Act
		result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
			joiner,
			"prefix",
			"suffix",
			"item1", "item2",
		)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendSkipNil_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendSkipNilTestCases {
		// Arrange & Act
		result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
			"prefix",
			"suffix",
			"item1", nil, "item3",
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendNilPrepend_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendNilPrependTestCases {
		// Arrange & Act
		result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
			nil,
			"suffix",
			"item1",
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendNilAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendNilAppendTestCases {
		// Arrange & Act
		result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
			"prefix",
			nil,
			"item1",
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAppendTestCases {
		// Arrange
		mainMap := map[string]string{"key1": "val1"}
		appendMap := map[string]any{"key2": "val2", "key3": 42}

		// Act
		result := coreappend.MapStringStringAppendMapStringToAnyItems(
			false,
			mainMap,
			appendMap,
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringAppend_SkipEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAppendSkipEmptyTestCases {
		// Arrange
		mainMap := map[string]string{"key1": "val1"}
		appendMap := map[string]any{"key2": "val2", "empty": ""}

		// Act
		result := coreappend.MapStringStringAppendMapStringToAnyItems(
			true,
			mainMap,
			appendMap,
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringAppend_EmptyAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAppendEmptyTestCases {
		// Arrange
		mainMap := map[string]string{"key1": "val1"}
		appendMap := map[string]any{}

		// Act
		result := coreappend.MapStringStringAppendMapStringToAnyItems(
			false,
			mainMap,
			appendMap,
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
