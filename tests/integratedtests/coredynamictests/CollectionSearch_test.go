package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ==========================================
// Test: Contains
// ==========================================

func Test_Collection_Contains_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionContainsTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.Contains(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IndexOf
// ==========================================

func Test_Collection_IndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIndexOfTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.IndexOf(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: HasAll
// ==========================================

func Test_Collection_HasAll_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionHasAllTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsStrings("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.HasAll(col, search...)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: LastIndexOf
// ==========================================

func Test_Collection_LastIndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionLastIndexOfTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.LastIndexOf(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Count
// ==========================================

func Test_Collection_Count_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCountTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.Count(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
