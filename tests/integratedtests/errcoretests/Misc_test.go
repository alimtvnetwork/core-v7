package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestSliceToError verifies SliceToError.
func TestSliceToError(t *testing.T) {
	// Empty returns nil
	actual := args.Map{"result": errcore.SliceToError(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
	actual := args.Map{"result": errcore.SliceToError([]string{}) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)

	// Non-empty returns error
	err := errcore.SliceToError([]string{"err1", "err2"})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// TestSliceToErrorPtr verifies SliceToErrorPtr.
func TestSliceToErrorPtr(t *testing.T) {
	actual := args.Map{"result": errcore.SliceToErrorPtr(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
	err := errcore.SliceToErrorPtr([]string{"e1"})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// TestToError verifies ToError.
func TestToError(t *testing.T) {
	actual := args.Map{"result": errcore.ToError("") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
	err := errcore.ToError("fail")
	actual := args.Map{"result": err == nil || err.Error() != "fail"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error with message", actual)
}

// TestToString verifies ToString.
func TestToString(t *testing.T) {
	actual := args.Map{"result": errcore.ToString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	actual := args.Map{"result": errcore.ToString(errors.New("test")) != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error string", actual)
}

// TestToStringPtr verifies ToStringPtr.
func TestToStringPtr(t *testing.T) {
	r := errcore.ToStringPtr(nil)
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return pointer", actual)
	actual := args.Map{"result": *r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil error should give empty string", actual)
}

// TestMergeErrors verifies MergeErrors.
func TestMergeErrors(t *testing.T) {
	actual := args.Map{"result": errcore.MergeErrors() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no errors should return nil", actual)
	actual := args.Map{"result": errcore.MergeErrors(nil, nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all nil should return nil", actual)
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return combined error", actual)
}

// TestCombine verifies Combine.
func TestCombine(t *testing.T) {
	result := errcore.Combine("generic", "other", "ref")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestSliceError verifies SliceError.
func TestSliceError(t *testing.T) {
	actual := args.Map{"result": errcore.SliceError(",", nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
	actual := args.Map{"result": errcore.SliceError(",", []string{}) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
	err := errcore.SliceError(",", []string{"a", "b"})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// TestSliceErrorDefault verifies SliceErrorDefault.
func TestSliceErrorDefault(t *testing.T) {
	actual := args.Map{"result": errcore.SliceErrorDefault(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

// TestManyErrorToSingle verifies ManyErrorToSingle.
func TestManyErrorToSingle(t *testing.T) {
	r := errcore.ManyErrorToSingle(nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
	r = errcore.ManyErrorToSingle([]error{errors.New("x")})
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "single error should return it", actual)
}

// TestManyErrorToSingleDirect verifies ManyErrorToSingleDirect.
func TestManyErrorToSingleDirect(t *testing.T) {
	r := errcore.ManyErrorToSingleDirect()
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
	r = errcore.ManyErrorToSingleDirect(errors.New("a"))
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// TestConcatMessageWithErr verifies ConcatMessageWithErr.
func TestConcatMessageWithErr(t *testing.T) {
	r := errcore.ConcatMessageWithErr("prefix", nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil error should return nil", actual)
	r = errcore.ConcatMessageWithErr("prefix", errors.New("err"))
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return concatenated error", actual)
}

// TestExpecting verifies Expecting error message.
func TestExpecting(t *testing.T) {
	r := errcore.Expecting("header", "expected", "actual")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestExpectingSimple verifies ExpectingSimple.
func TestExpectingSimple(t *testing.T) {
	r := errcore.ExpectingSimple("header", "expected", "actual")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestExpectingSimpleNoType verifies ExpectingSimpleNoType.
func TestExpectingSimpleNoType(t *testing.T) {
	r := errcore.ExpectingSimpleNoType("header", "expected", "actual")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestExpectingError verifies ExpectingError.
func TestExpectingError(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("header", "expected", "actual")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// TestExpectingErrorSimpleNoType verifies ExpectingErrorSimpleNoType.
func TestExpectingErrorSimpleNoType(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("header", "expected", "actual")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// TestExpectingNotEqualSimpleNoType verifies ExpectingNotEqualSimpleNoType.
func TestExpectingNotEqualSimpleNoType(t *testing.T) {
	r := errcore.ExpectingNotEqualSimpleNoType("header", "a", "b")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestVarTwo verifies VarTwo.
func TestVarTwo(t *testing.T) {
	r := errcore.VarTwo(false, "a", 1, "b", 2)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestVarThree verifies VarThree.
func TestVarThree(t *testing.T) {
	r := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestMessageVarTwo verifies MessageVarTwo.
func TestMessageVarTwo(t *testing.T) {
	r := errcore.MessageVarTwo("msg", "a", 1, "b", 2)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestMessageVarThree verifies MessageVarThree.
func TestMessageVarThree(t *testing.T) {
	r := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestMessageVarMap verifies MessageVarMap.
func TestMessageVarMap(t *testing.T) {
	r := errcore.MessageVarMap("msg", map[string]any{"k": "v"})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestVarMap verifies VarMap.
func TestVarMap(t *testing.T) {
	r := errcore.VarMap(map[string]any{"k": "v"})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestShouldBe verifies ShouldBe.
func TestShouldBe(t *testing.T) {
	r := errcore.ShouldBe.StrEqMsg("actual", "expected")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
	err := errcore.ShouldBe.StrEqErr("actual", "expected")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

// TestRawErrCollection verifies RawErrCollection.
func TestRawErrCollection(t *testing.T) {
	c := errcore.RawErrCollection{}
	actual := args.Map{"result": c.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have error", actual)
	c.Add(errors.New("err1"))
	actual := args.Map{"result": c.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have error", actual)
	actual := args.Map{"result": c.CompiledError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return compiled error", actual)
}

// TestSliceErrorsToStrings verifies SliceErrorsToStrings.
func TestSliceErrorsToStrings(t *testing.T) {
	r := errcore.SliceErrorsToStrings(nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	r = errcore.SliceErrorsToStrings(errors.New("a"), errors.New("b"))
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestErrorToSplitLines verifies ErrorToSplitLines.
func TestErrorToSplitLines(t *testing.T) {
	r := errcore.ErrorToSplitLines(nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice", actual)
	r = errcore.ErrorToSplitLines(errors.New("a\nb"))
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 lines", actual)
}

// TestErrorToSplitNonEmptyLines verifies ErrorToSplitNonEmptyLines.
func TestErrorToSplitNonEmptyLines(t *testing.T) {
	r := errcore.ErrorToSplitNonEmptyLines(nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice", actual)
}

// TestRef verifies Ref.
func TestRef(t *testing.T) {
	r := errcore.Ref("ref")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestMessageWithRef verifies MessageWithRef.
func TestMessageWithRef(t *testing.T) {
	r := errcore.MessageWithRef("msg", "ref")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestVarTwoNoType verifies VarTwoNoType.
func TestVarTwoNoType(t *testing.T) {
	r := errcore.VarTwoNoType("a", 1, "b", 2)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestVarThreeNoType verifies VarThreeNoType.
func TestVarThreeNoType(t *testing.T) {
	r := errcore.VarThreeNoType("a", 1, "b", 2, "c", 3)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestGetSearchTermExpectationMessage verifies search term message.
func TestGetSearchTermExpectationMessage(t *testing.T) {
	r := errcore.GetSearchTermExpectationMessage(1, "header", "expectation", 0, "actual", "expected", nil)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// TestGetSearchTermExpectationSimpleMessage verifies simple search term message.
func TestGetSearchTermExpectationSimpleMessage(t *testing.T) {
	r := errcore.GetSearchTermExpectationSimpleMessage(1, "expectation", 0, "content", "search")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}
