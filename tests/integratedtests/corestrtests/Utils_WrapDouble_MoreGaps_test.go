package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// utils — WrapDouble, WrapSingle, WrapTilda, WrapDoubleIfMissing, WrapSingleIfMissing
// =============================================================================

func Test_Utils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapDouble", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapDouble("hello")}

		// Assert
		expected := args.Map{"r": "\"hello\""}
		expected.ShouldBeEqual(t, 0, "WrapDouble", actual)
	})
}

func Test_Utils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapSingle", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapSingle("hello")}

		// Assert
		expected := args.Map{"r": "'hello'"}
		expected.ShouldBeEqual(t, 0, "WrapSingle", actual)
	})
}

func Test_Utils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapTilda", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapTilda("hello")}

		// Assert
		expected := args.Map{"r": "`hello`"}
		expected.ShouldBeEqual(t, 0, "WrapTilda", actual)
	})
}

func Test_Utils_WrapDoubleIfMissing_NoWrap(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapDoubleIfMissing_NoWrap", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapDoubleIfMissing("\"already\"")}

		// Assert
		expected := args.Map{"r": "\"already\""}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing already wrapped", actual)
	})
}

func Test_Utils_WrapDoubleIfMissing_Wrap(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapDoubleIfMissing_Wrap", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapDoubleIfMissing("need")}

		// Assert
		expected := args.Map{"r": "\"need\""}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing wraps", actual)
	})
}

func Test_Utils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapDoubleIfMissing_Empty", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapDoubleIfMissing("")}

		// Assert
		expected := args.Map{"r": "\"\""}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing empty", actual)
	})
}

func Test_Utils_WrapSingleIfMissing_NoWrap(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapSingleIfMissing_NoWrap", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapSingleIfMissing("'already'")}

		// Assert
		expected := args.Map{"r": "'already'"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing already wrapped", actual)
	})
}

func Test_Utils_WrapSingleIfMissing_Wrap(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapSingleIfMissing_Wrap", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapSingleIfMissing("need")}

		// Assert
		expected := args.Map{"r": "'need'"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing wraps", actual)
	})
}

func Test_Utils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov66_Utils_WrapSingleIfMissing_Empty", func() {
		// Act
		actual := args.Map{"r": corestr.StringUtils.WrapSingleIfMissing("")}

		// Assert
		expected := args.Map{"r": "''"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing empty", actual)
	})
}

// =============================================================================
// LeftMiddleRight — 15 uncovered methods
// =============================================================================

func Test_LMR_LeftBytes_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_LeftBytes", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"len": len(lmr.LeftBytes())}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LMR LeftBytes", actual)
	})
}

func Test_LMR_RightBytes_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_RightBytes", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"len": len(lmr.RightBytes())}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LMR RightBytes", actual)
	})
}

func Test_LMR_MiddleBytes_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_MiddleBytes", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"len": len(lmr.MiddleBytes())}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LMR MiddleBytes", actual)
	})
}

func Test_LMR_LeftTrim_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_LeftTrim", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("  abc  ", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.LeftTrim()}

		// Assert
		expected := args.Map{"r": "abc"}
		expected.ShouldBeEqual(t, 0, "LMR LeftTrim", actual)
	})
}

func Test_LMR_RightTrim_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_RightTrim", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "  xyz  ")

		// Act
		actual := args.Map{"r": lmr.RightTrim()}

		// Assert
		expected := args.Map{"r": "xyz"}
		expected.ShouldBeEqual(t, 0, "LMR RightTrim", actual)
	})
}

func Test_LMR_MiddleTrim_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_MiddleTrim", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "  mid  ", "xyz")

		// Act
		actual := args.Map{"r": lmr.MiddleTrim()}

		// Assert
		expected := args.Map{"r": "mid"}
		expected.ShouldBeEqual(t, 0, "LMR MiddleTrim", actual)
	})
}

func Test_LMR_IsLeftEmpty_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_IsLeftEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.IsLeftEmpty()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR IsLeftEmpty", actual)
	})
}

func Test_LMR_IsRightEmpty_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_IsRightEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "")

		// Act
		actual := args.Map{"r": lmr.IsRightEmpty()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR IsRightEmpty", actual)
	})
}

func Test_LMR_IsMiddleEmpty_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_IsMiddleEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "", "xyz")

		// Act
		actual := args.Map{"r": lmr.IsMiddleEmpty()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR IsMiddleEmpty", actual)
	})
}

func Test_LMR_IsMiddleWhitespace_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_IsMiddleWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "   ", "xyz")

		// Act
		actual := args.Map{"r": lmr.IsMiddleWhitespace()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR IsMiddleWhitespace", actual)
	})
}

func Test_LMR_IsLeftWhitespace_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_IsLeftWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("   ", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.IsLeftWhitespace()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR IsLeftWhitespace", actual)
	})
}

func Test_LMR_IsRightWhitespace_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_IsRightWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "   ")

		// Act
		actual := args.Map{"r": lmr.IsRightWhitespace()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR IsRightWhitespace", actual)
	})
}

func Test_LMR_HasValidNonEmptyLeft_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasValidNonEmptyLeft", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasValidNonEmptyLeft()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmptyLeft", actual)
	})
}

func Test_LMR_HasValidNonEmptyRight_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasValidNonEmptyRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasValidNonEmptyRight()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmptyRight", actual)
	})
}

func Test_LMR_HasValidNonEmptyMiddle_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasValidNonEmptyMiddle", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasValidNonEmptyMiddle()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmptyMiddle", actual)
	})
}

func Test_LMR_HasValidNonWhitespaceLeft_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasValidNonWhitespaceLeft", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasValidNonWhitespaceLeft()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonWhitespaceLeft", actual)
	})
}

func Test_LMR_HasValidNonWhitespaceRight_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasValidNonWhitespaceRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasValidNonWhitespaceRight()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonWhitespaceRight", actual)
	})
}

func Test_LMR_HasValidNonWhitespaceMiddle_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasValidNonWhitespaceMiddle", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasValidNonWhitespaceMiddle()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonWhitespaceMiddle", actual)
	})
}

func Test_LMR_HasSafeNonEmpty_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasSafeNonEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR HasSafeNonEmpty", actual)
	})
}

func Test_LMR_HasSafeNonEmpty_False(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_HasSafeNonEmpty_False", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "mid", "xyz")

		// Act
		actual := args.Map{"r": lmr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"r": false}
		expected.ShouldBeEqual(t, 0, "LMR HasSafeNonEmpty false", actual)
	})
}

func Test_LMR_IsAll_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_IsAll", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"r": lmr.IsAll("a", "b", "c")}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR IsAll", actual)
	})
}

func Test_LMR_Is_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_Is", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"r": lmr.Is("a", "c")}

		// Assert
		expected := args.Map{"r": true}
		expected.ShouldBeEqual(t, 0, "LMR Is", actual)
	})
}

func Test_LMR_Clone_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		clone := lmr.Clone()

		// Act
		actual := args.Map{
			"left": clone.Left,
			"mid": clone.Middle,
			"right": clone.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LMR Clone", actual)
	})
}

func Test_LMR_ToLeftRight_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LMR ToLeftRight", actual)
	})
}

func Test_LMR_Clear_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_Clear", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"mid": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LMR Clear", actual)
	})
}

func Test_LMR_Dispose_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_Dispose", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Dispose()

		// Act
		actual := args.Map{"left": lmr.Left}

		// Assert
		expected := args.Map{"left": ""}
		expected.ShouldBeEqual(t, 0, "LMR Dispose", actual)
	})
}

func Test_LMR_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_InvalidNoMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"valid": lmr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LMR InvalidNoMessage", actual)
	})
}

func Test_LMR_InvalidWithMessage(t *testing.T) {
	safeTest(t, "Test_Cov66_LMR_InvalidWithMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRight("bad")

		// Act
		actual := args.Map{
			"valid": lmr.IsValid,
			"msg": lmr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "bad",
		}
		expected.ShouldBeEqual(t, 0, "LMR InvalidWithMessage", actual)
	})
}

// =============================================================================
// NonChainedLinkedCollectionNodes — 6 uncovered
// =============================================================================

func Test_NCLCN_IsEmpty_Empty(t *testing.T) {
	safeTest(t, "Test_Cov66_NCLCN_IsEmpty_Empty", func() {
		// Arrange
		n := corestr.NewNonChainedLinkedCollectionNodes(0)

		// Act
		actual := args.Map{
			"empty": n.IsEmpty(),
			"has": n.HasItems(),
			"len": n.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"has": false,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "NCLCN empty", actual)
	})
}

func Test_NCLCN_IsChainingApplied_False(t *testing.T) {
	safeTest(t, "Test_Cov66_NCLCN_IsChainingApplied_False", func() {
		// Arrange
		n := corestr.NewNonChainedLinkedCollectionNodes(2)

		// Act
		actual := args.Map{"chained": n.IsChainingApplied()}

		// Assert
		expected := args.Map{"chained": false}
		expected.ShouldBeEqual(t, 0, "NCLCN not chained", actual)
	})
}

func Test_NCLCN_FirstLast_WithItems(t *testing.T) {
	safeTest(t, "Test_Cov66_NCLCN_FirstLast_WithItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		head := lc.Head()
		nodes := corestr.NewNonChainedLinkedCollectionNodes(2)
		nodes.Adds(head)
		if next := head.Next(); next != nil {
			nodes.Adds(next)
		}

		// Act
		actual := args.Map{
			"firstNonNil": nodes.First() != nil,
			"lastNonNil":  nodes.Last() != nil,
			"len":         nodes.Length(),
		}

		// Assert
		expected := args.Map{
			"firstNonNil": true,
			"lastNonNil":  true,
			"len":         2,
		}
		expected.ShouldBeEqual(t, 0, "NCLCN First/Last", actual)
	})
}

// =============================================================================
// NonChainedLinkedListNodes — 6 uncovered
// =============================================================================

func Test_NCLLN_IsEmpty_Empty(t *testing.T) {
	safeTest(t, "Test_Cov66_NCLLN_IsEmpty_Empty", func() {
		// Arrange
		n := corestr.NewNonChainedLinkedListNodes(0)

		// Act
		actual := args.Map{
			"empty": n.IsEmpty(),
			"has": n.HasItems(),
			"len": n.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"has": false,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "NCLLN empty", actual)
	})
}

func Test_NCLLN_IsChainingApplied_False(t *testing.T) {
	safeTest(t, "Test_Cov66_NCLLN_IsChainingApplied_False", func() {
		// Arrange
		n := corestr.NewNonChainedLinkedListNodes(2)

		// Act
		actual := args.Map{"chained": n.IsChainingApplied()}

		// Assert
		expected := args.Map{"chained": false}
		expected.ShouldBeEqual(t, 0, "NCLLN not chained", actual)
	})
}

func Test_NCLLN_FirstLast_WithItems(t *testing.T) {
	safeTest(t, "Test_Cov66_NCLLN_FirstLast_WithItems", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("x")
		ll.Add("y")
		head := ll.Head()
		nodes := corestr.NewNonChainedLinkedListNodes(2)
		nodes.Adds(head)
		if next := head.Next(); next != nil {
			nodes.Adds(next)
		}

		// Act
		actual := args.Map{
			"firstNonNil": nodes.First() != nil,
			"lastNonNil":  nodes.Last() != nil,
			"len":         nodes.Length(),
		}

		// Assert
		expected := args.Map{
			"firstNonNil": true,
			"lastNonNil":  true,
			"len":         2,
		}
		expected.ShouldBeEqual(t, 0, "NCLLN First/Last", actual)
	})
}

// =============================================================================
// LinkedCollections — Tail
// =============================================================================

func Test_LC_Tail(t *testing.T) {
	safeTest(t, "Test_Cov66_LC_Tail", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		tail := lc.Tail()

		// Act
		actual := args.Map{"nonNil": tail != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "LC Tail", actual)
	})
}

func Test_LC_Tail_Empty(t *testing.T) {
	safeTest(t, "Test_Cov66_LC_Tail_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		tail := lc.Tail()

		// Act
		actual := args.Map{"isNil": tail == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "LC Tail empty", actual)
	})
}

// =============================================================================
// CollectionsOfCollection — JSON methods (9 uncovered)
// =============================================================================

func Test_COC_Json(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_Json", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		r := coc.Json()

		// Act
		actual := args.Map{"noErr": r.Error == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "COC Json", actual)
	})
}

func Test_COC_JsonModel_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_JsonModel", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		m := coc.JsonModel()

		// Act
		actual := args.Map{"nonNil": m.Items != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "COC JsonModel", actual)
	})
}

func Test_COC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_JsonModelAny", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"nonNil": coc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "COC JsonModelAny", actual)
	})
}

func Test_COC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_MarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		b, err := coc.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "COC MarshalJSON", actual)
	})
}

func Test_COC_UnmarshalJSON_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_UnmarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		b, _ := coc.MarshalJSON()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err := coc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "COC UnmarshalJSON", actual)
	})
}

func Test_COC_ParseInjectUsingJson_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_ParseInjectUsingJson", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		r, err := coc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonNil": r != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonNil": true,
		}
		expected.ShouldBeEqual(t, 0, "COC ParseInjectUsingJson", actual)
	})
}

func Test_COC_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_ParseInjectUsingJson_Error", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := coc.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "COC ParseInjectUsingJson error", actual)
	})
}

func Test_COC_ParseInjectUsingJsonMust_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_ParseInjectUsingJsonMust", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		r := coc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "COC ParseInjectUsingJsonMust", actual)
	})
}

func Test_COC_JsonParseSelfInject_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_COC_JsonParseSelfInject", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err := coc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "COC JsonParseSelfInject", actual)
	})
}

// =============================================================================
// HashsetsCollection — JSON methods (8 uncovered)
// =============================================================================

func Test_HC_JsonModel_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_JsonModel", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
		m := hc.JsonModel()

		// Act
		actual := args.Map{"nonNil": m != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HC JsonModel", actual)
	})
}

func Test_HC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_JsonModelAny", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"nonNil": hc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HC JsonModelAny", actual)
	})
}

func Test_HC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_MarshalJSON", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
		b, err := hc.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "HC MarshalJSON", actual)
	})
}

func Test_HC_UnmarshalJSON_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_UnmarshalJSON", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
		b, _ := hc.MarshalJSON()
		hc2 := corestr.New.HashsetsCollection.Empty()
		err := hc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "HC UnmarshalJSON", actual)
	})
}

func Test_HC_ParseInjectUsingJson_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_ParseInjectUsingJson", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
		jr := hc.JsonPtr()
		hc2 := corestr.New.HashsetsCollection.Empty()
		r, err := hc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonNil": r != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HC ParseInjectUsingJson", actual)
	})
}

func Test_HC_ParseInjectUsingJsonMust_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_ParseInjectUsingJsonMust", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
		jr := hc.JsonPtr()
		hc2 := corestr.New.HashsetsCollection.Empty()
		r := hc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HC ParseInjectUsingJsonMust", actual)
	})
}

func Test_HC_JsonParseSelfInject_FromUtilsWrapDoubleMoreG(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_JsonParseSelfInject", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
		jr := hc.JsonPtr()
		hc2 := corestr.New.HashsetsCollection.Empty()
		err := hc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "HC JsonParseSelfInject", actual)
	})
}

func Test_HC_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov66_HC_UnmarshalJSON_Error", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		err := hc.UnmarshalJSON([]byte("invalid"))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "HC UnmarshalJSON error", actual)
	})
}
