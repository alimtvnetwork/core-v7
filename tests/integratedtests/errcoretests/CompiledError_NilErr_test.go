package errcoretests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── CompiledError ──

func Test_CompiledError_NilErr(t *testing.T) {
	// Arrange
	err := errcore.CompiledError(nil, "msg")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns nil -- nil error", actual)
}

func Test_CompiledError_EmptyMsg(t *testing.T) {
	// Arrange
	inner := errors.New("inner")
	err := errcore.CompiledError(inner, "")

	// Act
	actual := args.Map{"same": err == inner}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns same error -- empty message", actual)
}

func Test_CompiledError_WithMsg(t *testing.T) {
	// Arrange
	err := errcore.CompiledError(errors.New("inner"), "prefix")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns error -- with message", actual)
}

// ── CompiledErrorString ──

func Test_CompiledErrorString_NilErr(t *testing.T) {
	// Arrange
	result := errcore.CompiledErrorString(nil, "msg")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "CompiledErrorString returns empty -- nil error", actual)
}

func Test_CompiledErrorString_WithMsg(t *testing.T) {
	// Arrange
	result := errcore.CompiledErrorString(errors.New("inner"), "prefix")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompiledErrorString returns non-empty -- with message", actual)
}

// ── JoinErrors ──

func Test_JoinErrors(t *testing.T) {
	// Arrange
	err := errcore.JoinErrors(errors.New("a"), nil, errors.New("b"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JoinErrors returns joined error -- mixed nil and non-nil", actual)
}

// ── ConcatMessageWithErrWithStackTrace ──

func Test_ConcatMessageWithErrWithStackTrace_Nil(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErrWithStackTrace("msg", nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErrWithStackTrace returns nil -- nil error", actual)
}

func Test_ConcatMessageWithErrWithStackTrace_WithErr(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErrWithStackTrace("prefix", errors.New("e"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErrWithStackTrace returns error -- with error", actual)
}

// ── CombineWithMsgTypeNoStack ──

func Test_CombineWithMsgTypeNoStack_EmptyOtherMsg(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- empty otherMsg", actual)
}

func Test_CombineWithMsgTypeNoStack_WithOtherMsg(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "extra", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- with otherMsg", actual)
}

// ── CombineWithMsgTypeStackTrace ──

func Test_CombineWithMsgTypeStackTrace(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeStackTrace(errcore.InvalidType, "msg", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeStackTrace returns non-empty -- with stack trace", actual)
}

// ── CountStateChangeTracker ──

type mockLengthGetter struct{ length int }

func (m *mockLengthGetter) Length() int { return m.length }

func Test_CountStateChangeTracker_SameState(t *testing.T) {
	// Arrange
	lg := &mockLengthGetter{length: 5}
	tracker := errcore.NewCountStateChangeTracker(lg)

	// Act
	actual := args.Map{
		"same":    tracker.IsSameState(),
		"valid":   tracker.IsValid(),
		"success": tracker.IsSuccess(),
		"changes": tracker.HasChanges(),
		"failed":  tracker.IsFailed(),
		"sameC":   tracker.IsSameStateUsingCount(5),
	}

	// Assert
	expected := args.Map{
		"same": true, "valid": true, "success": true,
		"changes": false, "failed": false, "sameC": true,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker returns same state -- no changes", actual)
}

func Test_CountStateChangeTracker_Changed(t *testing.T) {
	// Arrange
	lg := &mockLengthGetter{length: 5}
	tracker := errcore.NewCountStateChangeTracker(lg)
	lg.length = 6

	// Act
	actual := args.Map{
		"same": tracker.IsSameState(),
		"changes": tracker.HasChanges(),
	}

	// Assert
	expected := args.Map{
		"same": false,
		"changes": true,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker returns changed state -- length increased", actual)
}

// ── EnumRangeNotMeet ──

func Test_EnumRangeNotMeet_WithRange(t *testing.T) {
	// Arrange
	result := errcore.EnumRangeNotMeet(1, 10, "1-10")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns non-empty -- with range string", actual)
}

func Test_EnumRangeNotMeet_WithoutRange(t *testing.T) {
	// Arrange
	result := errcore.EnumRangeNotMeet(1, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns non-empty -- nil range", actual)
}

// ── ErrorWithCompiledTraceRef ──

func Test_ErrorWithCompiledTraceRef_NilErr(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(nil, "trace", "ref")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns empty -- nil error", actual)
}

func Test_ErrorWithCompiledTraceRef_EmptyTraces(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns non-empty -- empty traces", actual)
}

func Test_ErrorWithCompiledTraceRef_NilRef(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "trace", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns non-empty -- nil ref", actual)
}

func Test_ErrorWithCompiledTraceRef_All(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "trace", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns non-empty -- all args", actual)
}

// ── ErrorWithCompiledTraceRefToError ──

func Test_ErrorWithCompiledTraceRefToError_Nil(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithCompiledTraceRefToError(nil, "t", "r")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError returns nil -- nil error", actual)
}

// ── ErrorWithRefToError ──

func Test_ErrorWithRefToError_Nil(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithRefToError(nil, "ref")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError returns nil -- nil error", actual)
}

func Test_ErrorWithRefToError_WithErr(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithRefToError(errors.New("e"), "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError returns error -- with error", actual)
}

// ── ErrorWithTracesRefToError ──

func Test_ErrorWithTracesRefToError_Nil(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithTracesRefToError(nil, []string{"t"}, "r")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError returns nil -- nil error", actual)
}

func Test_ErrorWithTracesRefToError_EmptyTraces(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{}, "r")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError returns non-nil -- empty traces", actual)
}

func Test_ErrorWithTracesRefToError_WithTraces(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{"t1", "t2"}, "r")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError returns non-nil -- with traces", actual)
}

// ── ExpectationMessageDef ──

func Test_ExpectationMessageDef_ExpectedSafeString(t *testing.T) {
	// Arrange
	emd := errcore.ExpectationMessageDef{Expected: "hello"}
	s1 := emd.ExpectedSafeString()
	s2 := emd.ExpectedSafeString() // cached

	// Act
	actual := args.Map{
		"notEmpty": s1 != "",
		"same": s1 == s2,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "ExpectedSafeString returns cached non-empty -- valid expected", actual)
}

func Test_ExpectationMessageDef_ExpectedSafeString_Nil(t *testing.T) {
	// Arrange
	emd := errcore.ExpectationMessageDef{}
	s := emd.ExpectedSafeString()

	// Act
	actual := args.Map{"empty": s == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ExpectedSafeString returns empty -- nil expected", actual)
}

func Test_ExpectationMessageDef_ExpectedStringTrim(t *testing.T) {
	// Arrange
	emd := errcore.ExpectationMessageDef{Expected: "  hello  "}
	s := emd.ExpectedStringTrim()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectedStringTrim returns non-empty -- padded input", actual)
}

func Test_ExpectationMessageDef_ExpectedString_Panic(t *testing.T) {
	// Arrange
	emd := errcore.ExpectationMessageDef{}
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		emd.ExpectedString()
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "ExpectedString panics -- nil expected", actual)
}

func Test_ExpectationMessageDef_ToString(t *testing.T) {
	// Arrange
	emd := errcore.ExpectationMessageDef{When: "w", FuncName: "f", Expected: "e", CaseIndex: 0}
	result := emd.ToString("actual")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToString returns formatted -- all args", actual)
}

func Test_ExpectationMessageDef_PrintIf_False(t *testing.T) {
	// Arrange
	emd := errcore.ExpectationMessageDef{When: "w", Expected: "e"}
	emd.PrintIf(false, "actual")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintIf completes safely -- condition false", actual)
}

func Test_ExpectationMessageDef_PrintIfFailed_NotFailed(t *testing.T) {
	// Arrange
	emd := errcore.ExpectationMessageDef{When: "w", Expected: "e"}
	emd.PrintIfFailed(true, false, "actual")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintIfFailed completes safely -- not failed", actual)
}

// ── ExpectingFuture / ExpectingRecord ──

func Test_ExpectingFuture_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	r := errcore.ExpectingFuture("title", "expected")

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"title": r.ExpectingTitle,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"title": "title",
	}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture returns record -- with title and expected", actual)
}

func Test_ExpectingRecord_Message(t *testing.T) {
	// Arrange
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}

	// Act
	actual := args.Map{"notEmpty": r.Message("a") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Message returns non-empty -- with actual", actual)
}

func Test_ExpectingRecord_MessageSimple(t *testing.T) {
	// Arrange
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}

	// Act
	actual := args.Map{"notEmpty": r.MessageSimple("a") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimple returns non-empty -- with actual", actual)
}

func Test_ExpectingRecord_MessageSimpleNoType(t *testing.T) {
	// Arrange
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}

	// Act
	actual := args.Map{"notEmpty": r.MessageSimpleNoType("a") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimpleNoType returns non-empty -- with actual", actual)
}

func Test_ExpectingRecord_Error(t *testing.T) {
	// Arrange
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}

	// Act
	actual := args.Map{"notNil": r.Error("a") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Error returns error -- with actual", actual)
}

func Test_ExpectingRecord_ErrorSimple(t *testing.T) {
	// Arrange
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}

	// Act
	actual := args.Map{"notNil": r.ErrorSimple("a") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimple returns error -- with actual", actual)
}

func Test_ExpectingRecord_ErrorSimpleNoType(t *testing.T) {
	// Arrange
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}

	// Act
	actual := args.Map{"notNil": r.ErrorSimpleNoType("a") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimpleNoType returns error -- with actual", actual)
}

// ── ExpectingNotEqualSimpleNoType ──

func Test_ExpectingNotEqualSimpleNoType_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.ExpectingNotEqualSimpleNoType("t", "e", "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType returns non-empty -- with args", actual)
}

// ── ExpectingSimpleNoTypeError ──

func Test_ExpectingSimpleNoTypeError(t *testing.T) {
	// Arrange
	err := errcore.ExpectingSimpleNoTypeError("t", "e", "a")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoTypeError returns error -- with args", actual)
}

// ── ExpectingErrorSimpleNoTypeNewLineEnds ──

func Test_ExpectingErrorSimpleNoTypeNewLineEnds(t *testing.T) {
	// Arrange
	err := errcore.ExpectingErrorSimpleNoTypeNewLineEnds("t", "e", "a")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoTypeNewLineEnds returns error -- with args", actual)
}

// ── WasExpectingErrorF ──

func Test_WasExpectingErrorF(t *testing.T) {
	// Arrange
	err := errcore.WasExpectingErrorF("e", "a", "title %d", 1)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "WasExpectingErrorF returns error -- with format", actual)
}

// ── FmtDebug / FmtDebugIf ──

func Test_FmtDebug_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	errcore.FmtDebug("test %d", 1)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug completes safely -- with format", actual)
}

func Test_FmtDebugIf_False(t *testing.T) {
	// Arrange
	errcore.FmtDebugIf(false, "test %d", 1)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf completes safely -- condition false", actual)
}

func Test_FmtDebugIf_True(t *testing.T) {
	// Arrange
	errcore.FmtDebugIf(true, "test %d", 1)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf completes safely -- condition true", actual)
}

// ── ValidPrint / FailedPrint ──

func Test_ValidPrint_True(t *testing.T) {
	// Arrange
	errcore.ValidPrint(true, "val")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "ValidPrint completes safely -- isValid true", actual)
}

func Test_ValidPrint_False(t *testing.T) {
	// Arrange
	errcore.ValidPrint(false, "val")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "ValidPrint completes safely -- isValid false", actual)
}

func Test_FailedPrint_True(t *testing.T) {
	// Arrange
	errcore.FailedPrint(true, "val")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FailedPrint completes safely -- isFailed true", actual)
}

func Test_FailedPrint_False(t *testing.T) {
	// Arrange
	errcore.FailedPrint(false, "val")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FailedPrint completes safely -- isFailed false", actual)
}

// ── GetActualAndExpectProcessedMessage ──

func Test_GetActualAndExpectProcessedMessage(t *testing.T) {
	// Arrange
	result := errcore.GetActualAndExpectProcessedMessage(0, "a", "e", "ap", "ep")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectProcessedMessage returns non-empty -- with args", actual)
}

// ── GetSearchLineNumberExpectationMessage ──

func Test_GetSearchLineNumberExpectationMessage(t *testing.T) {
	// Arrange
	result := errcore.GetSearchLineNumberExpectationMessage(0, 1, 2, "c", "s", "info")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage returns non-empty -- with args", actual)
}

// ── GetSearchTermExpectationMessage ──

func Test_GetSearchTermExpectationMessage_WithInfo(t *testing.T) {
	// Arrange
	result := errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", "info")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage returns non-empty -- with info", actual)
}

func Test_GetSearchTermExpectationMessage_NilInfo(t *testing.T) {
	// Arrange
	result := errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage returns non-empty -- nil info", actual)
}

// ── GetSearchTermExpectationSimpleMessage ──

func Test_GetSearchTermExpectationSimpleMessage(t *testing.T) {
	// Arrange
	result := errcore.GetSearchTermExpectationSimpleMessage(0, "e", 1, "c", "s")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationSimpleMessage returns non-empty -- with args", actual)
}

// ── GherkinsString / GherkinsStringWithExpectation ──

func Test_GherkinsString_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.GherkinsString(0, "f", "g", "w", "th")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString returns non-empty -- with args", actual)
}

func Test_GherkinsStringWithExpectation_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.GherkinsStringWithExpectation(0, "f", "g", "w", "th", "a", "e")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation returns non-empty -- with args", actual)
}

// ── Handle functions (panic paths) ──

func Test_HandleErr_Nil(t *testing.T) {
	// Arrange
	errcore.HandleErr(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErr completes safely -- nil error", actual)
}

func Test_HandleErr_WithErr(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.HandleErr(errors.New("e"))
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErr panics -- with error", actual)
}

func Test_HandleErrMessage_Empty(t *testing.T) {
	// Arrange
	errcore.HandleErrMessage("")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage completes safely -- empty message", actual)
}

func Test_HandleErrMessage_WithMsg(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.HandleErrMessage("e")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage panics -- with message", actual)
}

func Test_HandleErrorGetter_Nil(t *testing.T) {
	// Arrange
	errcore.HandleErrorGetter(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorGetter completes safely -- nil getter", actual)
}

func Test_HandleCompiledErrorGetter_Nil(t *testing.T) {
	// Arrange
	errcore.HandleCompiledErrorGetter(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorGetter completes safely -- nil getter", actual)
}

func Test_HandleCompiledErrorWithTracesGetter_Nil(t *testing.T) {
	// Arrange
	errcore.HandleCompiledErrorWithTracesGetter(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorWithTracesGetter completes safely -- nil getter", actual)
}

func Test_HandleFullStringsWithTracesGetter_Nil(t *testing.T) {
	// Arrange
	errcore.HandleFullStringsWithTracesGetter(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleFullStringsWithTracesGetter completes safely -- nil getter", actual)
}

// ── SimpleHandleErr ──

func Test_SimpleHandleErr_Nil(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErr(nil, "msg")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr completes safely -- nil error", actual)
}

func Test_SimpleHandleErr_WithErr(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.SimpleHandleErr(errors.New("e"), "msg")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr panics -- with error", actual)
}

// ── SimpleHandleErrMany ──

func Test_SimpleHandleErrMany_Nil(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErrMany("msg")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany completes safely -- nil errors", actual)
}

func Test_SimpleHandleErrMany_AllNil(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErrMany("msg", nil, nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany completes safely -- all nil errors", actual)
}

func Test_SimpleHandleErrMany_WithErr(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.SimpleHandleErrMany("msg", errors.New("e"))
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany panics -- with error", actual)
}

// ── MsgHeader / MsgHeaderIf / MsgHeaderPlusEnding ──

func Test_MsgHeader_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.MsgHeader("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader returns non-empty -- with args", actual)
}

func Test_MsgHeaderIf_True(t *testing.T) {
	// Arrange
	result := errcore.MsgHeaderIf(true, "hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf returns header -- condition true", actual)
}

func Test_MsgHeaderIf_False(t *testing.T) {
	// Arrange
	result := errcore.MsgHeaderIf(false, "hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf returns sprint -- condition false", actual)
}

func Test_MsgHeaderPlusEnding_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.MsgHeaderPlusEnding("h", "e")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding returns non-empty -- with args", actual)
}

// ── PanicOnIndexOutOfRange ──

func Test_PanicOnIndexOutOfRange_InRange(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.PanicOnIndexOutOfRange(5, []int{0, 1, 4})
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange completes safely -- index in range", actual)
}

func Test_PanicOnIndexOutOfRange_OutOfRange(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.PanicOnIndexOutOfRange(3, []int{5})
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange panics -- index out of range", actual)
}

// ── PanicRangeNotMeet / RangeNotMeet ──

func Test_PanicRangeNotMeet_WithRange(t *testing.T) {
	// Arrange
	result := errcore.PanicRangeNotMeet("msg", 1, 10, "1-10")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PanicRangeNotMeet panics -- with range string", actual)
}

func Test_PanicRangeNotMeet_WithoutRange(t *testing.T) {
	// Arrange
	result := errcore.PanicRangeNotMeet("msg", 1, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PanicRangeNotMeet panics -- nil range", actual)
}

func Test_RangeNotMeet_WithRange(t *testing.T) {
	// Arrange
	result := errcore.RangeNotMeet("msg", 1, 10, "1-10")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns non-empty -- with range string", actual)
}

func Test_RangeNotMeet_WithoutRange(t *testing.T) {
	// Arrange
	result := errcore.RangeNotMeet("msg", 1, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns non-empty -- nil range", actual)
}

// ── PathMeaningfulMessage ──

func Test_PathMeaningfulMessage_Empty(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulMessage(errcore.PathErrorType, "fn", "/path")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage returns nil -- no messages", actual)
}

func Test_PathMeaningfulMessage_WithMsgs(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulMessage(errcore.PathErrorType, "fn", "/path", "msg1", "msg2")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage returns error -- with messages", actual)
}

// ── PathMeaningfulError ──

func Test_PathMeaningfulError_Nil(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulError(errcore.PathErrorType, nil, "/path")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError returns nil -- nil error", actual)
}

func Test_PathMeaningfulError_WithErr(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulError(errcore.PathErrorType, errors.New("e"), "/path")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError returns error -- with error", actual)
}

// ── MeaningfulError / MeaningfulErrorWithData / MeaningfulMessageError ──

func Test_MeaningfulError_Nil(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulError(errcore.InvalidType, "fn", nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns nil -- nil error", actual)
}

func Test_MeaningfulError_WithErr(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulError(errcore.InvalidType, "fn", errors.New("e"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- with error", actual)
}

func Test_MeaningfulErrorWithData_Nil(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", nil, "data")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData returns nil -- nil error", actual)
}

func Test_MeaningfulErrorWithData_WithErr(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", errors.New("e"), "data")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData returns error -- with error", actual)
}

func Test_MeaningfulMessageError_Nil(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulMessageError(errcore.InvalidType, "fn", nil, "msg")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError returns nil -- nil error", actual)
}

func Test_MeaningfulMessageError_WithErr(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulMessageError(errcore.InvalidType, "fn", errors.New("e"), "msg")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError returns error -- with error", actual)
}

// ── MeaningfulErrorHandle ──

func Test_MeaningfulErrorHandle_Nil(t *testing.T) {
	// Arrange
	errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorHandle completes safely -- nil error", actual)
}

func Test_MeaningfulErrorHandle_WithErr(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", errors.New("e"))
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorHandle panics -- with error", actual)
}

// ── PrintError / PrintErrorWithTestIndex ──

func Test_PrintError_Nil(t *testing.T) {
	// Arrange
	errcore.PrintError(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintError completes safely -- nil error", actual)
}

func Test_PrintError_WithErr(t *testing.T) {
	// Arrange
	errcore.PrintError(errors.New("e"))

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintError logs error -- with error", actual)
}

func Test_PrintErrorWithTestIndex_Nil(t *testing.T) {
	// Arrange
	errcore.PrintErrorWithTestIndex(0, "h", nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex completes safely -- nil error", actual)
}

func Test_PrintErrorWithTestIndex_WithErr(t *testing.T) {
	// Arrange
	errcore.PrintErrorWithTestIndex(0, "h", errors.New("e"))

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex logs error -- with error", actual)
}

// ── SourceDestination / SourceDestinationErr / SourceDestinationNoType ──

func Test_SourceDestination_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.SourceDestination(true, "s", "d")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- with args", actual)
}

func Test_SourceDestinationErr_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	err := errcore.SourceDestinationErr(false, "s", "d")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr returns error -- with args", actual)
}

func Test_SourceDestinationNoType_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.SourceDestinationNoType("s", "d")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType returns formatted -- with args", actual)
}

// ── StackTracesCompiled ──

func Test_StackTracesCompiled(t *testing.T) {
	// Arrange
	result := errcore.StackTracesCompiled([]string{"t1", "t2"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled returns non-empty -- with traces", actual)
}

// ── StringLinesToQuoteLines / StringLinesToQuoteLinesToSingle / LinesToDoubleQuoteLinesWithTabs ──

func Test_StringLinesToQuoteLines_Empty(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns empty -- nil input", actual)
}

func Test_StringLinesToQuoteLines_NonEmpty(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns formatted -- non-empty input", actual)
}

func Test_StringLinesToQuoteLinesToSingle_FromCompiledErrorNilErr(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle returns non-empty -- with input", actual)
}

func Test_LinesToDoubleQuoteLinesWithTabs_Empty(t *testing.T) {
	// Arrange
	result := errcore.LinesToDoubleQuoteLinesWithTabs(2, []string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesToDoubleQuoteLinesWithTabs returns empty -- empty input", actual)
}

func Test_LinesToDoubleQuoteLinesWithTabs_WithTabs(t *testing.T) {
	// Arrange
	result := errcore.LinesToDoubleQuoteLinesWithTabs(4, []string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesToDoubleQuoteLinesWithTabs returns formatted -- with input", actual)
}

// ── ToExitError ──

func Test_ToExitError_Nil(t *testing.T) {
	// Arrange
	result := errcore.ToExitError(nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError returns nil -- nil error", actual)
}

func Test_ToExitError_NonExitError(t *testing.T) {
	// Arrange
	result := errcore.ToExitError(errors.New("e"))

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError returns nil -- non-ExitError", actual)
}

// ── getReferenceMessage (indirect via CombineWithMsgTypeNoStack) ──

func Test_getReferenceMessage_EmptyString(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "msg", "")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "getReferenceMessage returns empty -- empty string", actual)
}

// ── RawErrorType methods ──

func Test_RawErrorType_CombineWithAnother(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.CombineWithAnother(errcore.NotFound, "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": string(result) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithAnother returns non-empty -- with another type", actual)
}

func Test_RawErrorType_TypesAttach(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.TypesAttach("msg", 42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypesAttach returns non-empty -- with types", actual)
}

func Test_RawErrorType_TypesAttachErr(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.TypesAttachErr("msg", 42)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypesAttachErr returns error -- with types", actual)
}

func Test_RawErrorType_SrcDestination(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.SrcDestination("msg", "src", 1, "dst", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SrcDestination returns non-empty -- with args", actual)
}

func Test_RawErrorType_SrcDestinationErr(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.SrcDestinationErr("msg", "src", 1, "dst", 2)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SrcDestinationErr returns error -- with args", actual)
}

func Test_RawErrorType_Error(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.Error("msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Error returns error -- with msg and ref", actual)
}

func Test_RawErrorType_ErrorSkip(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorSkip(0, "msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorSkip returns error -- with skip", actual)
}

func Test_RawErrorType_Fmt_Empty(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.Fmt("", )

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Fmt returns error -- empty format", actual)
}

func Test_RawErrorType_Fmt_WithFormat(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.Fmt("format %d", 1)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Fmt returns error -- with format", actual)
}

func Test_RawErrorType_FmtIf_False(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.FmtIf(false, "format %d", 1)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FmtIf returns nil -- condition false", actual)
}

func Test_RawErrorType_FmtIf_True(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.FmtIf(true, "format %d", 1)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FmtIf returns error -- condition true", actual)
}

func Test_RawErrorType_MergeError_Nil(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeError(nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeError returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeError_WithErr(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeError(errors.New("e"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeError returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithMessage(nil, "msg")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessage returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_WithErr(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithMessage(errors.New("e"), "msg")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessage returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessageRef returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_WithErr(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessageRef returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithRef(nil, "ref")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithRef returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithRef_WithErr(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithRef(errors.New("e"), "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithRef returns error -- with error", actual)
}

func Test_RawErrorType_MsgCsvRef_Empty(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.MsgCsvRef("msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRef returns non-empty -- empty msg", actual)
}

func Test_RawErrorType_MsgCsvRef_WithItems(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.MsgCsvRef("msg", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRef returns non-empty -- with items", actual)
}

func Test_RawErrorType_MsgCsvRef_EmptyMsg(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.MsgCsvRef("", "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRef returns non-empty -- empty msg", actual)
}

func Test_RawErrorType_MsgCsvRefError(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MsgCsvRefError("msg", "a")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRefError returns error -- with items", actual)
}

func Test_RawErrorType_ErrorRefOnly(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorRefOnly("ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorRefOnly returns error -- with ref", actual)
}

func Test_RawErrorType_Expecting(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.Expecting("e", "a")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expecting returns formatted -- with args", actual)
}

func Test_RawErrorType_NoRef_Empty(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.NoRef("")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NoRef returns non-empty -- empty msg", actual)
}

func Test_RawErrorType_NoRef_WithMsg(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.NoRef("msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NoRef returns non-empty -- with msg", actual)
}

func Test_RawErrorType_ErrorNoRefs(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorNoRefs("msg")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorNoRefs returns error -- with msg", actual)
}

func Test_RawErrorType_ErrorNoRefs_Empty(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorNoRefs("")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorNoRefs returns error -- empty msg", actual)
}

func Test_RawErrorType_HandleUsingPanic(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.InvalidType.HandleUsingPanic("msg", "ref")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleUsingPanic panics -- with error", actual)
}

// ── GetSet / GetSetVariant ──

func Test_GetSet_True(t *testing.T) {
	// Arrange
	result := errcore.GetSet(true, errcore.InvalidType, errcore.NotFound)

	// Act
	actual := args.Map{"val": string(result)}

	// Assert
	expected := args.Map{"val": string(errcore.InvalidType)}
	expected.ShouldBeEqual(t, 0, "GetSet returns trueValue -- condition true", actual)
}

func Test_GetSet_False(t *testing.T) {
	// Arrange
	result := errcore.GetSet(false, errcore.InvalidType, errcore.NotFound)

	// Act
	actual := args.Map{"val": string(result)}

	// Assert
	expected := args.Map{"val": string(errcore.NotFound)}
	expected.ShouldBeEqual(t, 0, "GetSet returns falseValue -- condition false", actual)
}

func Test_GetSetVariant_True(t *testing.T) {
	// Arrange
	result := errcore.GetSetVariant(true, "a", "b")

	// Act
	actual := args.Map{"val": string(result)}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns trueValue -- condition true", actual)
}

func Test_GetSetVariant_False(t *testing.T) {
	// Arrange
	result := errcore.GetSetVariant(false, "a", "b")

	// Act
	actual := args.Map{"val": string(result)}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns falseValue -- condition false", actual)
}

// ── ShouldBe ──

func Test_ShouldBe_StrEqMsg(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.StrEqMsg("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg returns non-empty -- different strings", actual)
}

func Test_ShouldBe_StrEqErr(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.StrEqErr("a", "b")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr returns error -- different strings", actual)
}

func Test_ShouldBe_AnyEqMsg(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.AnyEqMsg(1, 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg returns non-empty -- different values", actual)
}

func Test_ShouldBe_AnyEqErr(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.AnyEqErr(1, 2)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr returns error -- different values", actual)
}

func Test_ShouldBe_JsonEqMsg(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.JsonEqMsg("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg returns non-empty -- different json", actual)
}

func Test_ShouldBe_JsonEqErr(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.JsonEqErr("a", "b")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr returns error -- different json", actual)
}

// ── Expected ──

func Test_Expected_But(t *testing.T) {
	// Arrange
	err := errcore.Expected.But("title", "exp", "act")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.But returns error -- with args", actual)
}

func Test_Expected_ButFoundAsMsg(t *testing.T) {
	// Arrange
	result := errcore.Expected.ButFoundAsMsg("title", "exp", "act")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	// Arrange
	result := errcore.Expected.ButFoundWithTypeAsMsg("title", "exp", "act")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButUsingType(t *testing.T) {
	// Arrange
	err := errcore.Expected.ButUsingType("title", "exp", "act")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType returns error -- with args", actual)
}

func Test_Expected_ReflectButFound(t *testing.T) {
	// Arrange
	err := errcore.Expected.ReflectButFound(reflect.String, reflect.Int)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ReflectButFound returns error -- different kinds", actual)
}

func Test_Expected_PrimitiveButFound(t *testing.T) {
	// Arrange
	err := errcore.Expected.PrimitiveButFound(reflect.Slice)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.PrimitiveButFound returns error -- non-primitive kind", actual)
}

func Test_Expected_ValueHasNoElements(t *testing.T) {
	// Arrange
	err := errcore.Expected.ValueHasNoElements(reflect.Slice)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ValueHasNoElements returns error -- with kind", actual)
}

// ── StackEnhance ──

func Test_StackEnhance_Error_Nil(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.Error(nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns nil -- nil error", actual)
}

func Test_StackEnhance_Error_WithErr(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.Error(errors.New("e"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns error -- with error", actual)
}

func Test_StackEnhance_Msg_Empty(t *testing.T) {
	// Arrange
	result := errcore.StackEnhance.Msg("")

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns empty -- empty message", actual)
}

func Test_StackEnhance_Msg_NonEmpty(t *testing.T) {
	// Arrange
	result := errcore.StackEnhance.Msg("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns non-empty -- non-empty message", actual)
}

func Test_StackEnhance_MsgToErrSkip_Empty(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.MsgToErrSkip(0, "")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgToErrSkip returns nil -- empty message", actual)
}

func Test_StackEnhance_FmtSkip_Empty(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.FmtSkip(0, "")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns nil -- empty format", actual)
}

func Test_StackEnhance_FmtSkip_NonEmpty(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.FmtSkip(0, "hello %d", 1)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns error -- non-empty format", actual)
}

func Test_StackEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	// Arrange
	result := errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip returns empty -- nil error input", actual)
}

func Test_StackEnhance_MsgErrorSkip_WithErr(t *testing.T) {
	// Arrange
	result := errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("e"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip returns non-empty -- with error", actual)
}

func Test_StackEnhance_MsgErrorToErrSkip_Nil(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns nil -- nil error", actual)
}

func Test_StackEnhance_MsgErrorToErrSkip_WithErr(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns error -- with error", actual)
}

// ── VarNameValues / VarNameValuesJoiner / VarNameValuesStrings ──

func Test_VarNameValues_Empty(t *testing.T) {
	// Arrange
	result := errcore.VarNameValues()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns empty -- no args", actual)
}

func Test_VarNameValues_NonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarNameValues(namevalue.StringAny{Name: "k", Value: "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns formatted -- with args", actual)
}

func Test_VarNameValuesJoiner_Empty(t *testing.T) {
	// Arrange
	result := errcore.VarNameValuesJoiner(",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner returns empty -- no args", actual)
}

func Test_VarNameValuesJoiner_NonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarNameValuesJoiner(",", namevalue.StringAny{Name: "k", Value: "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner returns joined -- with name-values", actual)
}

func Test_VarNameValuesStrings_Empty(t *testing.T) {
	// Arrange
	result := errcore.VarNameValuesStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings returns empty -- no args", actual)
}

// ── MessageNameValues ──

func Test_MessageNameValues_Empty(t *testing.T) {
	// Arrange
	result := errcore.MessageNameValues("msg")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns msg only -- no name-values", actual)
}

func Test_MessageNameValues_NonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageNameValues("msg", namevalue.StringAny{Name: "k", Value: "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns formatted -- with name-values", actual)
}
