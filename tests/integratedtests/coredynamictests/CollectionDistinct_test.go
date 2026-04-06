package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ==========================================
// Test: Distinct
// ==========================================

func Test_Collection_Distinct_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDistinctTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		result := coredynamic.Distinct(col)

		// Handle mixed ExpectedInput types
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {

		// Act
			actual := args.Map{
				"distinctCount": result.Length(),
			}
			for i := 0; i < result.Length(); i++ {
				actual[fmt.Sprintf("item%d", i)] = result.SafeAt(i)
			}

		// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result.Length()))
		}
	}
}

// ==========================================
// Test: DistinctCount
// ==========================================

func Test_Collection_DistinctCount_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDistinctCountTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", coredynamic.DistinctCount(col)))
	}
}

// ==========================================
// Test: IsDistinct
// ==========================================

func Test_Collection_IsDistinct_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsDistinctTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", coredynamic.IsDistinct(col)))
	}
}
