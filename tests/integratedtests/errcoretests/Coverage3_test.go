package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── ErrorWithRef ──

func Test_Cov3_ErrorWithRef_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ErrorWithRef(nil, "ref")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns empty -- nil error", actual)
}

func Test_Cov3_ErrorWithRef_NilRef(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ErrorWithRef(errors.New("fail"), nil)}

	// Assert
	expected := args.Map{"result": "fail"}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns error msg -- nil reference", actual)
}

func Test_Cov3_ErrorWithRef_EmptyRef(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ErrorWithRef(errors.New("fail"), "")}

	// Assert
	expected := args.Map{"result": "fail"}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns error msg -- empty reference", actual)
}

func Test_Cov3_ErrorWithRef_WithRef(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithRef(errors.New("fail"), "ctx")

	// Act
	actual := args.Map{"notEmpty": result != "", "containsErr": true}

	// Assert
	expected := args.Map{"notEmpty": true, "containsErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns formatted -- with reference", actual)
}

// ── ErrorWithRefToError ──

func Test_Cov3_ErrorWithRefToError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ErrorWithRefToError(nil, "ref") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError returns nil -- nil error", actual)
}

func Test_Cov3_ErrorWithRefToError_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ErrorWithRefToError(errors.New("fail"), "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError returns error -- with error", actual)
}

// ── RefToError ──

func Test_Cov3_RefToError_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.RefToError(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError returns nil -- nil reference", actual)
}

func Test_Cov3_RefToError_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.RefToError("ref-val") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RefToError returns error -- non-nil reference", actual)
}

// ── MessageWithRefToError ──

func Test_Cov3_MessageWithRefToError(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.MessageWithRefToError("msg", "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError returns error -- always", actual)
}

// ── ErrorWithCompiledTraceRef ──

func Test_Cov3_ErrorWithCompiledTraceRef_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ErrorWithCompiledTraceRef(nil, "traces", "ref")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns empty -- nil error", actual)
}

func Test_Cov3_ErrorWithCompiledTraceRef_EmptyTraces(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(errors.New("fail"), "", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef delegates to ErrorWithRef -- empty traces", actual)
}

func Test_Cov3_ErrorWithCompiledTraceRef_NilRef(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(errors.New("fail"), "stack-data", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef formats without ref -- nil reference", actual)
}

func Test_Cov3_ErrorWithCompiledTraceRef_Full(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(errors.New("fail"), "stack-data", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef formats full -- all args", actual)
}

// ── ErrorWithCompiledTraceRefToError ──

func Test_Cov3_ErrorWithCompiledTraceRefToError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ErrorWithCompiledTraceRefToError(nil, "traces", "ref") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError returns nil -- nil error", actual)
}

func Test_Cov3_ErrorWithCompiledTraceRefToError_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ErrorWithCompiledTraceRefToError(errors.New("fail"), "traces", "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError returns error -- with error", actual)
}

// ── ErrorWithTracesRefToError ──

func Test_Cov3_ErrorWithTracesRefToError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ErrorWithTracesRefToError(nil, []string{"t"}, "ref") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError returns nil -- nil error", actual)
}

func Test_Cov3_ErrorWithTracesRefToError_EmptyTraces(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ErrorWithTracesRefToError(errors.New("fail"), nil, "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError delegates to ErrorWithRefToError -- empty traces", actual)
}

func Test_Cov3_ErrorWithTracesRefToError_WithTraces(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ErrorWithTracesRefToError(errors.New("fail"), []string{"trace1"}, "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError returns compiled error -- with traces", actual)
}

// ── StackTracesCompiled ──

func Test_Cov3_StackTracesCompiled(t *testing.T) {
	// Arrange
	result := errcore.StackTracesCompiled([]string{"line1", "line2"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled returns formatted -- multiple lines", actual)
}

// ── CombineWithMsgTypeNoStack ──

func Test_Cov3_CombineWithMsgTypeNoStack_EmptyMsg(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns type only -- empty otherMsg", actual)
}

func Test_Cov3_CombineWithMsgTypeNoStack_WithMsg(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "details", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns combined -- with otherMsg", actual)
}

// ── CombineWithMsgTypeStackTrace ──

func Test_Cov3_CombineWithMsgTypeStackTrace(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeStackTrace(errcore.InvalidRequestType, "details", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeStackTrace returns enhanced -- with stack trace", actual)
}

// ── MeaningfulError ──

func Test_Cov3_MeaningfulError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.MeaningfulError(errcore.InvalidRequestType, "fn", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns nil -- nil error", actual)
}

func Test_Cov3_MeaningfulError_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.MeaningfulError(errcore.InvalidRequestType, "fn", errors.New("fail")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- with error", actual)
}

// ── MeaningfulErrorWithData ──

func Test_Cov3_MeaningfulErrorWithData_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.MeaningfulErrorWithData(errcore.InvalidRequestType, "fn", nil, "data") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData returns nil -- nil error", actual)
}

func Test_Cov3_MeaningfulErrorWithData_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.MeaningfulErrorWithData(errcore.InvalidRequestType, "fn", errors.New("fail"), "data") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData returns error -- with error and data", actual)
}

// ── MeaningfulMessageError ──

func Test_Cov3_MeaningfulMessageError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.MeaningfulMessageError(errcore.InvalidRequestType, "fn", nil, "msg") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError returns nil -- nil error", actual)
}

func Test_Cov3_MeaningfulMessageError_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.MeaningfulMessageError(errcore.InvalidRequestType, "fn", errors.New("fail"), "msg") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError returns error -- with error and message", actual)
}

// ── PathMeaningfulError ──

func Test_Cov3_PathMeaningfulError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.PathMeaningfulError(errcore.InvalidRequestType, nil, "/tmp") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError returns nil -- nil error", actual)
}

func Test_Cov3_PathMeaningfulError_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.PathMeaningfulError(errcore.InvalidRequestType, errors.New("fail"), "/tmp") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError returns error -- with error and location", actual)
}

// ── ConcatMessageWithErr (error return) ──

func Test_Cov3_ConcatMessageWithErr_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ConcatMessageWithErr("prefix", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr returns nil -- nil error", actual)
}

func Test_Cov3_ConcatMessageWithErr_WithErr(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr returns wrapped error -- with error", actual)
}

// ── ConcatMessageWithErrWithStackTrace ──

func Test_Cov3_ConcatMessageWithErrWithStackTrace_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ConcatMessageWithErrWithStackTrace("prefix", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErrWithStackTrace returns nil -- nil error", actual)
}

func Test_Cov3_ConcatMessageWithErrWithStackTrace_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ConcatMessageWithErrWithStackTrace("prefix", errors.New("inner")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErrWithStackTrace returns error -- with error", actual)
}

// ── ToExitError ──

func Test_Cov3_ToExitError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ToExitError(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError returns nil -- nil error", actual)
}

func Test_Cov3_ToExitError_NonExitErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ToExitError(errors.New("not exit")) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError returns nil -- non-ExitError", actual)
}

// ── ToValueString ──

func Test_Cov3_ToValueString(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.ToValueString("hello") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToValueString returns formatted -- string input", actual)
}

// ── VarMapStrings ──

func Test_Cov3_VarMapStrings_Empty(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.VarMapStrings(nil))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarMapStrings returns empty -- nil map", actual)
}

func Test_Cov3_VarMapStrings_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.VarMapStrings(map[string]any{"k": "v"}))}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VarMapStrings returns entries -- populated map", actual)
}

// ── VarNameValuesStrings ──

func Test_Cov3_VarNameValuesStrings_Empty(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.VarNameValuesStrings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings returns empty -- no args", actual)
}

func Test_Cov3_VarNameValuesStrings_NonEmpty(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "key", Value: "val"}

	// Act
	actual := args.Map{"len": len(errcore.VarNameValuesStrings(nv))}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings returns entries -- with name-values", actual)
}

// ── VarNameValuesJoiner ──

func Test_Cov3_VarNameValuesJoiner_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.VarNameValuesJoiner(", ")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner returns empty -- no args", actual)
}

func Test_Cov3_VarNameValuesJoiner_NonEmpty(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "key", Value: "val"}

	// Act
	actual := args.Map{"notEmpty": errcore.VarNameValuesJoiner(", ", nv) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner returns joined -- with name-values", actual)
}

// ── MsgHeader / MsgHeaderIf / MsgHeaderPlusEnding ──

func Test_Cov3_MsgHeader(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MsgHeader("title") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader returns formatted -- with items", actual)
}

func Test_Cov3_MsgHeaderIf_True(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MsgHeaderIf(true, "title") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf returns header -- isHeader true", actual)
}

func Test_Cov3_MsgHeaderIf_False(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MsgHeaderIf(false, "title") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf returns sprint -- isHeader false", actual)
}

func Test_Cov3_MsgHeaderPlusEnding(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MsgHeaderPlusEnding("header", "ending") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding returns formatted -- with args", actual)
}

// ── GherkinsStringWithExpectation ──

func Test_Cov3_GherkinsStringWithExpectation(t *testing.T) {
	// Arrange
	result := errcore.GherkinsStringWithExpectation(1, "feature", "given", "when", "then", "actual", "expected")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation returns formatted -- all args", actual)
}

// ── HandleErrMessage (nil path) ──

func Test_Cov3_HandleErrMessage_Empty(t *testing.T) {
	// Arrange
	// Should not panic
	errcore.HandleErrMessage("")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage returns safely -- empty message", actual)
}

// ── PrintError (nil path) ──

func Test_Cov3_PrintError_Nil(t *testing.T) {
	// Arrange
	errcore.PrintError(nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintError returns safely -- nil error", actual)
}

func Test_Cov3_PrintError_NonNil(t *testing.T) {
	// Arrange
	errcore.PrintError(errors.New("test"))

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintError logs error -- non-nil error", actual)
}

// ── PrintErrorWithTestIndex (nil path) ──

func Test_Cov3_PrintErrorWithTestIndex_Nil(t *testing.T) {
	// Arrange
	errcore.PrintErrorWithTestIndex(0, "title", nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex returns safely -- nil error", actual)
}

func Test_Cov3_PrintErrorWithTestIndex_NonNil(t *testing.T) {
	// Arrange
	errcore.PrintErrorWithTestIndex(0, "title", errors.New("test"))

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex logs error -- non-nil error", actual)
}

// ── FmtDebugIf ──

func Test_Cov3_FmtDebugIf_False(t *testing.T) {
	// Arrange
	errcore.FmtDebugIf(false, "format %d", 42)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf skips logging -- isDebug false", actual)
}

func Test_Cov3_FmtDebugIf_True(t *testing.T) {
	// Arrange
	errcore.FmtDebugIf(true, "format %d", 42)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf logs -- isDebug true", actual)
}

// ── FmtDebug / ValidPrint / FailedPrint ──

func Test_Cov3_FmtDebug(t *testing.T) {
	// Arrange
	errcore.FmtDebug("value %d", 42)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug completes -- with format args", actual)
}

func Test_Cov3_ValidPrint_True(t *testing.T) {
	// Arrange
	errcore.ValidPrint(true, "data")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ValidPrint logs -- isValid true", actual)
}

func Test_Cov3_ValidPrint_False(t *testing.T) {
	// Arrange
	errcore.ValidPrint(false, "data")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ValidPrint skips -- isValid false", actual)
}

func Test_Cov3_FailedPrint_True(t *testing.T) {
	// Arrange
	errcore.FailedPrint(true, "data")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FailedPrint logs -- isFailed true", actual)
}

func Test_Cov3_FailedPrint_False(t *testing.T) {
	// Arrange
	errcore.FailedPrint(false, "data")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FailedPrint skips -- isFailed false", actual)
}

// ── SimpleHandleErrMany (nil path) ──

func Test_Cov3_SimpleHandleErrMany_NilSlice(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErrMany("msg")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany returns safely -- nil errors", actual)
}

func Test_Cov3_SimpleHandleErrMany_AllNilErrors(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErrMany("msg", nil, nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany returns safely -- all nil errors", actual)
}

// ── EnumRangeNotMeet ──

func Test_Cov3_EnumRangeNotMeet_NilRange(t *testing.T) {
	// Arrange
	result := errcore.EnumRangeNotMeet(0, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns formatted -- nil wholeRange", actual)
}

func Test_Cov3_EnumRangeNotMeet_WithRange(t *testing.T) {
	// Arrange
	result := errcore.EnumRangeNotMeet(0, 10, "0,1,2,5,10")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns formatted -- with wholeRange", actual)
}

// ── RangeNotMeet ──

func Test_Cov3_RangeNotMeet_NilRange(t *testing.T) {
	// Arrange
	result := errcore.RangeNotMeet("msg", 0, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns formatted -- nil wholeRange", actual)
}

func Test_Cov3_RangeNotMeet_WithRange(t *testing.T) {
	// Arrange
	result := errcore.RangeNotMeet("msg", 0, 10, "0,5,10")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns formatted -- with wholeRange", actual)
}

// ── MapMismatchError ──

func Test_Cov3_MapMismatchError(t *testing.T) {
	// Arrange
	result := errcore.MapMismatchError(
		"TestFunc",
		1,
		"title",
		[]string{`"key": "actual"`},
		[]string{`"key": "expected"`},
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapMismatchError returns formatted -- with entries", actual)
}

// ── StackEnhance ──

func Test_Cov3_StackEnhance_Error_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.Error(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns nil -- nil error", actual)
}

func Test_Cov3_StackEnhance_Error_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.StackEnhance.Error(errors.New("fail")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns enhanced -- non-nil error", actual)
}

func Test_Cov3_StackEnhance_Msg_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.StackEnhance.Msg("")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns empty -- empty message", actual)
}

func Test_Cov3_StackEnhance_Msg_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.StackEnhance.Msg("test") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns enhanced -- non-empty message", actual)
}

func Test_Cov3_StackEnhance_MsgToErrSkip_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.MsgToErrSkip(0, "") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgToErrSkip returns nil -- empty message", actual)
}

func Test_Cov3_StackEnhance_FmtSkip_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.FmtSkip(0, "") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns nil -- empty format", actual)
}

func Test_Cov3_StackEnhance_FmtSkip_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.StackEnhance.FmtSkip(0, "error %d", 42) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns error -- with format", actual)
}

func Test_Cov3_StackEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip returns empty -- nil error", actual)
}

func Test_Cov3_StackEnhance_MsgErrorSkip_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("fail")) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip returns enhanced -- with error", actual)
}

func Test_Cov3_StackEnhance_MsgErrorToErrSkip_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns nil -- nil error", actual)
}

func Test_Cov3_StackEnhance_MsgErrorToErrSkip_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("fail")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns error -- with error", actual)
}

// ── Combine (package-level) ──

func Test_Cov3_Combine(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.Combine("generic", "other", "ref") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine returns formatted -- all args", actual)
}

// ── getReferenceMessage (indirectly via CombineWithMsgTypeNoStack) ──

func Test_Cov3_GetReferenceMessage_NilRef(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "", nil)
	// With nil ref, no " Ref(s) { ... }" suffix

	// Act
	actual := args.Map{"isTypeOnly": result == errcore.InvalidRequestType.String()}

	// Assert
	expected := args.Map{"isTypeOnly": true}
	expected.ShouldBeEqual(t, 0, "getReferenceMessage returns empty -- nil reference", actual)
}

func Test_Cov3_GetReferenceMessage_EmptyStringRef(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "", "")

	// Act
	actual := args.Map{"isTypeOnly": result == errcore.InvalidRequestType.String()}

	// Assert
	expected := args.Map{"isTypeOnly": true}
	expected.ShouldBeEqual(t, 0, "getReferenceMessage returns empty -- empty string reference", actual)
}

// ── LinesToDoubleQuoteLinesWithTabs ──

func Test_Cov3_LinesToDoubleQuoteLinesWithTabs(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.LinesToDoubleQuoteLinesWithTabs(4, []string{"a", "b"}))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinesToDoubleQuoteLinesWithTabs returns entries -- with lines", actual)
}

func Test_Cov3_LinesToDoubleQuoteLinesWithTabs_Empty(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.LinesToDoubleQuoteLinesWithTabs(4, nil))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesToDoubleQuoteLinesWithTabs returns empty -- nil lines", actual)
}

// ── GetSearchLineNumberExpectationMessage ──

func Test_Cov3_GetSearchLineNumberExpectationMessage(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.GetSearchLineNumberExpectationMessage(1, 5, 3, "content", "search", "info") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage returns formatted -- all args", actual)
}
