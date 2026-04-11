package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── ErrorWithRef ──

func Test_ErrorWithRef(t *testing.T) {
	// Arrange
	err := errors.New("test")

	// Act
	actual := args.Map{
		"withRef":    errcore.ErrorWithRef(err, "ref") != "",
		"nilErr":     errcore.ErrorWithRef(nil, "ref"),
		"nilRef":     errcore.ErrorWithRef(err, nil) != "",
		"emptyRef":   errcore.ErrorWithRef(err, "") != "",
	}

	// Assert
	expected := args.Map{
		"withRef": true,
		"nilErr": "",
		"nilRef": true,
		"emptyRef": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns formatted -- with error and ref", actual)
}

// ── ErrorWithCompiledTraceRef ──

func Test_ErrorWithCompiledTraceRef(t *testing.T) {
	// Arrange
	err := errors.New("test")

	// Act
	actual := args.Map{
		"full":        errcore.ErrorWithCompiledTraceRef(err, "trace", "ref") != "",
		"nilErr":      errcore.ErrorWithCompiledTraceRef(nil, "trace", "ref"),
		"emptyTrace":  errcore.ErrorWithCompiledTraceRef(err, "", "ref") != "",
		"nilRef":      errcore.ErrorWithCompiledTraceRef(err, "trace", nil) != "",
	}

	// Assert
	expected := args.Map{
		"full": true,
		"nilErr": "",
		"emptyTrace": true,
		"nilRef": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns non-empty -- with all args", actual)
}

// ── ErrorWithRefToError / ErrorWithCompiledTraceRefToError / ErrorWithTracesRefToError ──

func Test_ErrorWithRefToError(t *testing.T) {
	// Arrange
	err := errors.New("test")
	result := errcore.ErrorWithRefToError(err, "ref")
	nilResult := errcore.ErrorWithRefToError(nil, "ref")

	// Act
	actual := args.Map{
		"hasErr": result != nil,
		"nilResult": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"nilResult": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError returns error -- with error", actual)
}

func Test_ErrorWithCompiledTraceRefToError(t *testing.T) {
	// Arrange
	err := errors.New("test")
	result := errcore.ErrorWithCompiledTraceRefToError(err, "trace", "ref")
	nilResult := errcore.ErrorWithCompiledTraceRefToError(nil, "trace", "ref")

	// Act
	actual := args.Map{
		"hasErr": result != nil,
		"nilResult": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"nilResult": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError returns error -- with args", actual)
}

// ── HandleErr ──

func Test_HandleErr_Nil_FromErrorWithRef(t *testing.T) {
	// Arrange
	errcore.HandleErr(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr completes safely -- nil error", actual)
}

func Test_HandleErr_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "HandleErr panics -- with error", actual)
	}()
	errcore.HandleErr(errors.New("test"))
}

// ── HandleErrMessage ──

func Test_HandleErrMessage_Nil(t *testing.T) {
	// Arrange
	errcore.HandleErrMessage("")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage completes safely -- nil message", actual)
}

// ── SimpleHandleErr ──

func Test_SimpleHandleErr_Nil_FromErrorWithRef(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErr(nil, "msg")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr completes safely -- nil error", actual)
}

// ── SimpleHandleErrMany ──

func Test_SimpleHandleErrMany_AllNil_FromErrorWithRef(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErrMany("msg", nil, nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany completes safely -- all nil", actual)
}

// ── PrintError ──

func Test_PrintError(t *testing.T) {
	// Arrange
	errcore.PrintError(errors.New("test"))
	errcore.PrintError(nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintError completes safely -- with error", actual)
}

// ── PrintErrorWithTestIndex ──

func Test_PrintErrorWithTestIndex(t *testing.T) {
	// Arrange
	errcore.PrintErrorWithTestIndex(0, "header", errors.New("test"))
	errcore.PrintErrorWithTestIndex(0, "header", nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex completes safely -- with error", actual)
}

// ── LineDiff ──

func Test_LineDiff(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "c"})

	// Act
	actual := args.Map{
		"len": len(diffs),
		"firstMatch": diffs[0].Status,
		"secondMismatch": diffs[1].Status,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"firstMatch": "  ",
		"secondMismatch": "!!",
	}
	expected.ShouldBeEqual(t, 0, "LineDiff returns diffs -- with lines", actual)
}

func Test_LineDiff_ExtraActual(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a"})

	// Act
	actual := args.Map{
		"len": len(diffs),
		"status": diffs[1].Status,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"status": "+",
	}
	expected.ShouldBeEqual(t, 0, "LineDiff returns extra-actual -- longer actual", actual)
}

func Test_LineDiff_MissingExpected(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(diffs),
		"status": diffs[1].Status,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"status": "-",
	}
	expected.ShouldBeEqual(t, 0, "LineDiff returns missing-expected -- shorter actual", actual)
}

func Test_LineDiffToString(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "test", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns formatted -- with diffs", actual)
}

func Test_LineDiffToString_Empty(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "test", []string{}, []string{})

	// Act
	actual := args.Map{"empty": result}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns empty -- both empty", actual)
}

func Test_HasAnyMismatchOnLines(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"}),
		"noMatch": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"}),
		"diffLen": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"}),
	}

	// Assert
	expected := args.Map{
		"match": false,
		"noMatch": true,
		"diffLen": true,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns correct -- with lines", actual)
}

func Test_SliceDiffSummary(t *testing.T) {
	// Arrange
	match := errcore.SliceDiffSummary([]string{"a"}, []string{"a"})
	noMatch := errcore.SliceDiffSummary([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{
		"match": match,
		"noMatchNotEmpty": noMatch != "",
	}

	// Assert
	expected := args.Map{
		"match": "all lines match",
		"noMatchNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns correct -- with lines", actual)
}

func Test_ErrorToLinesLineDiff(t *testing.T) {
	// Arrange
	result := errcore.ErrorToLinesLineDiff(0, "test", errors.New("a"), []string{"a"})
	nilResult := errcore.ErrorToLinesLineDiff(0, "test", nil, []string{"a"})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"nilNotEmpty": nilResult != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff returns non-empty -- with error", actual)
}

// ── GetActualAndExpectProcessedMessage / GetActualAndExpectSortedMessage ──

func Test_GetActualAndExpectProcessedMessage_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.GetActualAndExpectProcessedMessage(0, "actual", "expected", "actualProc", "expectedProc")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectProcessedMessage returns non-empty -- with args", actual)
}

func Test_GetSearchTermExpectationMessage(t *testing.T) {
	// Arrange
	result := errcore.GetSearchTermExpectationMessage(0, "header", "expectMsg", 1, "actual", "expected", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage returns non-empty -- with args", actual)
}

func Test_GetSearchTermExpectationSimpleMessage_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.GetSearchTermExpectationSimpleMessage(0, "expectMsg", 1, "content", "search")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationSimpleMessage returns non-empty -- with args", actual)
}

func Test_GetSearchLineNumberExpectationMessage_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.GetSearchLineNumberExpectationMessage(0, 5, 3, "content", "search", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage returns non-empty -- with args", actual)
}

// ── MessageVarTwo / MessageVarThree / MessageVarMap ──

func Test_MessageVarTwo_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo returns formatted -- with args", actual)
}

func Test_MessageVarThree_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree returns formatted -- with args", actual)
}

func Test_MessageVarMap_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.MessageVarMap("msg", map[string]any{"a": 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns formatted -- with map", actual)
}

// ── MergeErrors ──

func Test_MergeErrors(t *testing.T) {
	// Arrange
	result := errcore.MergeErrors(errors.New("a"), nil, errors.New("b"))
	nilResult := errcore.MergeErrors(nil, nil)

	// Act
	actual := args.Map{
		"hasErr": result != nil,
		"nilNil": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"nilNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns correct -- with errors", actual)
}

func Test_MergeErrorsToString(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToString(", ", errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns correct -- with errors", actual)
}

// ── SliceToError / SliceToErrorPtr ──

func Test_SliceToError(t *testing.T) {
	// Arrange
	result := errcore.SliceToError([]string{"a"})
	nilResult := errcore.SliceToError(nil)

	// Act
	actual := args.Map{
		"hasErr": result != nil,
		"nilNil": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"nilNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceToError returns correct -- with slice", actual)
}

func Test_SliceToErrorPtr(t *testing.T) {
	// Arrange
	result := errcore.SliceToErrorPtr([]string{"a"})
	empty := errcore.SliceToErrorPtr([]string{})

	// Act
	actual := args.Map{
		"hasErr": result != nil,
		"emptyNil": empty == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"emptyNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns correct -- with slice", actual)
}

// ── ShouldBe / Expected ──

func Test_ShouldBe(t *testing.T) {
	// Arrange
	msg := errcore.ShouldBe.AnyEqMsg("a", "b")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe exists -- singleton check", actual)
}

// ── PrintLineDiff / PrintLineDiffOnFail ──

func Test_PrintLineDiff(t *testing.T) {
	// Arrange
	errcore.PrintLineDiff(0, "test", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiff completes safely -- with diffs", actual)
}

func Test_PrintLineDiffOnFail_Match(t *testing.T) {
	// Arrange
	errcore.PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail completes safely -- matching", actual)
}

func Test_PrintLineDiffOnFail_Mismatch(t *testing.T) {
	// Arrange
	errcore.PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail prints diff -- with mismatch", actual)
}

// ── PrintErrorLineDiff ──

func Test_PrintErrorLineDiff(t *testing.T) {
	// Arrange
	errcore.PrintErrorLineDiff(0, "test", errors.New("a"), []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorLineDiff completes safely -- with args", actual)
}

// ── AssertDiffOnMismatch / PrintDiffOnMismatch ──

func Test_AssertDiffOnMismatch(t *testing.T) {
	// Arrange
	errcore.AssertDiffOnMismatch(t, 0, "test", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "AssertDiffOnMismatch completes safely -- matching", actual)
}

func Test_PrintDiffOnMismatch(t *testing.T) {
	// Arrange
	errcore.PrintDiffOnMismatch(0, "test", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch completes safely -- with args", actual)
}

// ── StackTracesCompiled ──

func Test_StackTracesCompiled_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.StackTracesCompiled([]string{"trace1", "trace2"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled returns non-empty -- with traces", actual)
}

// ── CombineWithMsgType ──

func Test_CombineWithMsgTypeNoStack_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack("type", "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- with args", actual)
}

// ── CompiledError ──

func Test_CompiledError(t *testing.T) {
	// Arrange
	result := errcore.CompiledError(errors.New("inner"), "additional")

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns error -- with message", actual)
}

// ── PathMeaningFulMessage / PathMeaningfulError ──

func Test_PathMeaningfulMessage(t *testing.T) {
	// Arrange
	result := errcore.PathMeaningfulMessage("type", "funcName", "/path", "msg")

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage returns error -- with messages", actual)
}

func Test_PathMeaningfulError(t *testing.T) {
	// Arrange
	result := errcore.PathMeaningfulError("type", errors.New("inner"), "/path")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError returns error -- with error", actual)
}

// ── MeaningFulError / MeaningFulErrorHandle / MeaningFulErrorWithData ──

func Test_MeaningFulError(t *testing.T) {
	// Arrange
	result := errcore.MeaningfulError("type", "msg", errors.New("inner"))

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- with error", actual)
}

func Test_MeaningfulMessageError(t *testing.T) {
	// Arrange
	result := errcore.MeaningfulMessageError("type", "funcName", errors.New("inner"), "msg")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError returns error -- with error", actual)
}

func Test_MeaningFulErrorWithData(t *testing.T) {
	// Arrange
	result := errcore.MeaningfulErrorWithData("type", "msg", errors.New("inner"), "data")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData returns error -- with error", actual)
}

// ── ToExitError ──

func Test_ToExitError(t *testing.T) {
	// Arrange
	result := errcore.ToExitError(errors.New("test"))
	nilResult := errcore.ToExitError(nil)

	// Act
	actual := args.Map{
		"isNil": result == nil,
		"nilNil": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"nilNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ToExitError returns correct value -- with error type", actual)
}

// ── RangeNotMeet / EnumRangeNotMeet ──

func Test_RangeNotMeet_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.RangeNotMeet("type", 5, 1, 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns non-empty -- with range", actual)
}

func Test_EnumRangeNotMeet_FromErrorWithRef(t *testing.T) {
	// Arrange
	result := errcore.EnumRangeNotMeet(1, 3, 5)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns non-empty -- with range", actual)
}

// ── StackEnhance ──

func Test_StackEnhance(t *testing.T) {
	// Arrange
	result := errcore.StackEnhance.MsgSkip(0, "test msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgSkip returns non-empty -- with message", actual)
}
