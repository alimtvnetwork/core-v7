package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Src_NonEmptySlice_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptySliceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptySlice(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_NonEmptySlicePtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptySlicePtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptySlicePtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_NonWhitespace_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespaceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespace(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_NonWhitespacePtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespacePtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespacePtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_NonNullStrings_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonNullStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonNullStrings(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_NonEmptyStrings_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptyStrings(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_NonEmptyIf_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyIfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)
		isFilter := input["isFilter"].(bool)

		// Act
		result := stringslice.NonEmptyIf(isFilter, slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_NonEmptyJoin_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyJoinTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptyJoin(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_NonEmptyJoinPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyJoinPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptyJoinPtr(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_NonWhitespaceJoin_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespaceJoinTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespaceJoin(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_NonWhitespaceJoinPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespaceJoinPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespaceJoinPtr(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_JoinWith_Verification(t *testing.T) {
	for caseIndex, tc := range srcJoinWithTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		separator := input["separator"].(string)
		items := input["items"].([]string)

		// Act
		result := stringslice.JoinWith(separator, items...)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_Joins_Verification(t *testing.T) {
	for caseIndex, tc := range srcJoinsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		separator := input["separator"].(string)
		items := input["items"].([]string)

		// Act
		result := stringslice.Joins(separator, items...)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_Src_TrimmedEachWords_Verification(t *testing.T) {
	for caseIndex, tc := range srcTrimmedEachWordsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.TrimmedEachWords(slice)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["isNil"]; has {
			actual["isNil"] = result == nil
		}
		if _, has := expected["length"]; has {
			actual["length"] = len(result)
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_TrimmedEachWordsPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcTrimmedEachWordsPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.TrimmedEachWordsPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_TrimmedEachWordsIf_Verification(t *testing.T) {
	for caseIndex, tc := range srcTrimmedEachWordsIfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)
		isTrim := input["isTrim"].(bool)

		// Act
		result := stringslice.TrimmedEachWordsIf(isTrim, slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_SplitContentsByWhitespace_Verification(t *testing.T) {
	for caseIndex, tc := range srcSplitContentsByWhitespaceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		str := input["input"].(string)

		// Act
		result := stringslice.SplitContentsByWhitespace(str)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_Src_SplitTrimmedNonEmptyAll_Verification(t *testing.T) {
	for caseIndex, tc := range srcSplitTrimmedNonEmptyAllTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		str := input["input"].(string)
		separator := input["separator"].(string)

		// Act
		result := stringslice.SplitTrimmedNonEmptyAll(str, separator)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}
