package corepayloadtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// pagingInfoDiff returns a diff-style string comparing two PagingInfo pointers.
// Prints all fields side-by-side so failures show exactly what differs.
func pagingInfoDiff(label string, left, right *corepayload.PagingInfo) string {
	leftStr := "<nil>"
	rightStr := "<nil>"

	if left != nil {
		leftStr = fmt.Sprintf(
			"{TotalPages:%d, CurrentPageIndex:%d, PerPageItems:%d, TotalItems:%d}",
			left.TotalPages, left.CurrentPageIndex, left.PerPageItems, left.TotalItems,
		)
	}

	if right != nil {
		rightStr = fmt.Sprintf(
			"{TotalPages:%d, CurrentPageIndex:%d, PerPageItems:%d, TotalItems:%d}",
			right.TotalPages, right.CurrentPageIndex, right.PerPageItems, right.TotalItems,
		)
	}

	return fmt.Sprintf(
		"\n[%s]\n  Left:  %s\n  Right: %s",
		label, leftStr, rightStr,
	)
}

func Test_PagingInfo_IsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isLeftNil := input.GetAsBoolDefault("isLeftNil", false)
		isRightNil := input.GetAsBoolDefault("isRightNil", false)

		var left, right *corepayload.PagingInfo
		if !isLeftNil {
			left = buildPagingInfoPrefixed(input, "left")
		}
		if !isRightNil {
			right = buildPagingInfoPrefixed(input, "right")
		}

		// Act
		result := left.IsEqual(right)

		// Assert
		actual := args.Map{
			"isEqual": result,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PagingInfo_State_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var info *corepayload.PagingInfo
		if !isNil {
			info = buildPagingInfoFromMap(input)
		}

		// Act
		actual := args.Map{
			"isEmpty":                   info.IsEmpty(),
			"hasTotalPages":             info.HasTotalPages(),
			"hasCurrentPageIndex":       info.HasCurrentPageIndex(),
			"hasPerPageItems":           info.HasPerPageItems(),
			"hasTotalItems":             info.HasTotalItems(),
			"isInvalidTotalPages":       info.IsInvalidTotalPages(),
			"isInvalidCurrentPageIndex": info.IsInvalidCurrentPageIndex(),
			"isInvalidPerPageItems":     info.IsInvalidPerPageItems(),
			"isInvalidTotalItems":       info.IsInvalidTotalItems(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PagingInfo_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		info := buildPagingInfoFromMap(input)

		// Act
		clone := info.Clone()

		// Assert
		actual := args.Map{
			"totalPages":       clone.TotalPages,
			"currentPageIndex": clone.CurrentPageIndex,
			"perPageItems":     clone.PerPageItems,
			"totalItems":       clone.TotalItems,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PagingInfo_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoClonePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var info *corepayload.PagingInfo
		if !isNil {
			info = buildPagingInfoFromMap(input)
		}

		// Act
		result := info.ClonePtr()

		// Assert
		if isNil {
			actual := args.Map{
				"isNil": result == nil,
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			actual := args.Map{
				"isNil":            result == nil,
				"totalPages":       result.TotalPages,
				"currentPageIndex": result.CurrentPageIndex,
				"perPageItems":     result.PerPageItems,
				"totalItems":       result.TotalItems,
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

// === Independence tests ===

func Test_PagingInfo_ClonePtr_Independence(t *testing.T) {
	tc := pagingInfoClonePtrIndependenceTestCase
	info := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	clone := info.ClonePtr()
	clone.TotalPages = 99
	clone.CurrentPageIndex = 99

	actual := args.Map{
		"originalTotalPages":  info.TotalPages,
		"originalCurrentPage": info.CurrentPageIndex,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PagingInfo_Clone_Independence(t *testing.T) {
	tc := pagingInfoCloneIndependenceTestCase
	info := corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	clone := info.Clone()
	clone.TotalPages = 99

	actual := args.Map{"originalTotalPages": info.TotalPages}

	tc.ShouldBeEqualMapFirst(t, actual)
}
