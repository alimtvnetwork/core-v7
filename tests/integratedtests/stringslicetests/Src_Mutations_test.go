package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Src_Clone_Verification(t *testing.T) {
	for caseIndex, tc := range srcCloneTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.Clone(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_ClonePtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcClonePtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.ClonePtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_CloneUsingCap_Verification(t *testing.T) {
	for caseIndex, tc := range srcCloneUsingCapTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		capVal := input["cap"].(int)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.CloneUsingCap(capVal, slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_MergeNew_Verification(t *testing.T) {
	for caseIndex, tc := range srcMergeNewTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		extras := input["extras"].([]string)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.MergeNew(slice, extras...)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_MergeNewSimple_Verification(t *testing.T) {
	for caseIndex, tc := range srcMergeNewSimpleTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left := input["left"].([]string)
		right := input["right"].([]string)

		// Act
		result := stringslice.MergeNewSimple(left, right)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_InPlaceReverse_Verification(t *testing.T) {
	for caseIndex, tc := range srcInPlaceReverseTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			orig := sliceRaw.([]string)
			slice = make([]string, len(orig))
			copy(slice, orig)
		}

		// Act
		result := stringslice.InPlaceReverse(&slice)

		// Assert
		actual := args.Map{}
		expected := tc.ExpectedInput.(args.Map)

		if _, has := expected["length"]; has {
			actual["length"] = len(*result)
		}
		if _, has := expected["first"]; has {
			actual["first"] = (*result)[0]
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_PrependNew_Verification(t *testing.T) {
	for caseIndex, tc := range srcPrependNewTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		prepend := input["prepend"].([]string)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.PrependNew(slice, prepend...)

		// Assert
		actual := args.Map{
			"length": len(*result),
		}
		expected := tc.ExpectedInput.(args.Map)
		if _, has := expected["first"]; has {
			actual["first"] = (*result)[0]
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_AppendLineNew_Verification(t *testing.T) {
	for caseIndex, tc := range srcAppendLineNewTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)
		appendVal := input["append"].(string)

		// Act
		result := stringslice.AppendLineNew(slice, appendVal)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_PrependLineNew_Verification(t *testing.T) {
	for caseIndex, tc := range srcPrependLineNewTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		line := input["line"].(string)
		slice := input["input"].([]string)

		// Act
		result := stringslice.PrependLineNew(line, slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_SortIf_Verification(t *testing.T) {
	for caseIndex, tc := range srcSortIfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		origSlice := input["input"].([]string)
		isSort := input["isSort"].(bool)

		slice := make([]string, len(origSlice))
		copy(slice, origSlice)

		// Act
		stringslice.SortIf(isSort, slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"first": slice[0],
		})
	}
}

func Test_Src_ExpandBySplit_Verification(t *testing.T) {
	for caseIndex, tc := range srcExpandBySplitTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.ExpandBySplit(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}

		if _, has := expected["minLength"]; has {
			if len(result) >= expected["minLength"].(int) {
				actual["minLength"] = expected["minLength"]
			} else {
				actual["minLength"] = len(result)
			}
		}
		if _, has := expected["length"]; has {
			actual["length"] = len(result)
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
