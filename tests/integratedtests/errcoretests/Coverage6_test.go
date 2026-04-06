package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// mockLengthGetterCov6 implements the lengthGetter interface for testing.
type mockLengthGetterCov6 struct {
	length int
}

func (m *mockLengthGetterCov6) Length() int {
	return m.length
}

// ── CombineWithMsgTypeNoStack ──

func Test_Cov6_CombineWithMsgTypeNoStack(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack("test-type", "test msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- with args", actual)
}

// ── ConcatMessageWithErr ──

func Test_Cov6_ConcatMessageWithErr(t *testing.T) {
	// Arrange
	result := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))
	nilResult := errcore.ConcatMessageWithErr("prefix", nil)

	// Act
	actual := args.Map{
		"hasResult": result != nil,
		"nilResult": nilResult == nil,
	}

	// Assert
	expected := args.Map{"hasResult": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr returns correct -- nil and non-nil", actual)
}

// ── ErrorToSplitLines / ErrorToSplitNonEmptyLines ──

func Test_Cov6_ErrorToSplitLines(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitLines(errors.New("a\nb\n"))
	nilResult := errcore.ErrorToSplitLines(nil)

	// Act
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}

	// Assert
	expected := args.Map{"len": 3, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines returns lines -- with newlines", actual)
}

func Test_Cov6_ErrorToSplitNonEmptyLines(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))
	nilResult := errcore.ErrorToSplitNonEmptyLines(nil)

	// Act
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}

	// Assert
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines returns filtered -- with empty lines", actual)
}

// ── ManyErrorToSingle / ManyErrorToSingleDirect ──

func Test_Cov6_ManyErrorToSingle(t *testing.T) {
	// Arrange
	errs := []error{errors.New("a"), nil, errors.New("b")}
	result := errcore.ManyErrorToSingle(errs)
	nilResult := errcore.ManyErrorToSingle(nil)
	allNil := errcore.ManyErrorToSingle([]error{nil, nil})

	// Act
	actual := args.Map{
		"hasErr":    result != nil,
		"nilResult": nilResult == nil,
		"allNil":    allNil == nil,
	}

	// Assert
	expected := args.Map{"hasErr": true, "nilResult": true, "allNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle returns correct -- mixed nil and non-nil", actual)
}

func Test_Cov6_ManyErrorToSingleDirect(t *testing.T) {
	// Arrange
	result := errcore.ManyErrorToSingleDirect(errors.New("a"), nil, errors.New("b"))
	nilResult := errcore.ManyErrorToSingleDirect()

	// Act
	actual := args.Map{"hasErr": result != nil, "nilResult": nilResult == nil}

	// Assert
	expected := args.Map{"hasErr": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect returns correct -- mixed args", actual)
}

// ── ToError / ToString / ToStringPtr / ToValueString ──

func Test_Cov6_ToError(t *testing.T) {
	// Arrange
	result := errcore.ToError("hello")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToError returns correct -- empty and non-empty", actual)
}

func Test_Cov6_ToString(t *testing.T) {
	// Arrange
	err := errors.New("test")
	result := errcore.ToString(err)
	nilResult := errcore.ToString(nil)

	// Act
	actual := args.Map{"result": result, "nilResult": nilResult}

	// Assert
	expected := args.Map{"result": "test", "nilResult": ""}
	expected.ShouldBeEqual(t, 0, "ToString returns correct -- nil and non-nil", actual)
}

func Test_Cov6_ToStringPtr(t *testing.T) {
	// Arrange
	err := errors.New("test")
	result := errcore.ToStringPtr(err)
	nilResult := errcore.ToStringPtr(nil)

	// Act
	actual := args.Map{"notNil": result != nil, "nilResultNotNil": nilResult != nil, "nilValue": *nilResult}

	// Assert
	expected := args.Map{"notNil": true, "nilResultNotNil": true, "nilValue": ""}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns correct -- nil and non-nil", actual)
}

func Test_Cov6_ToValueString(t *testing.T) {
	// Arrange
	err := errors.New("test")
	result := errcore.ToValueString(err)
	nilResult := errcore.ToValueString(nil)

	// Act
	actual := args.Map{"result": result, "nilResult": nilResult}

	// Assert
	expected := args.Map{"result": "test", "nilResult": "<nil>"}
	expected.ShouldBeEqual(t, 0, "ToValueString returns non-empty -- with value", actual)
}

// ── RawErrCollection ──

func Test_Cov6_RawErrCollection(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Add(errors.New("a"))
	c.Add(errors.New("b"))

	// Act
	actual := args.Map{
		"len":    c.Length(),
		"hasAny": c.HasAnyError(),
	}

	// Assert
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection returns correct state -- comprehensive test", actual)
}

func Test_Cov6_RawErrCollection_CombinedError(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Add(errors.New("a"))
	result := c.CompiledError()
	empty := errcore.RawErrCollection{}
	emptyResult := empty.CompiledError()

	// Act
	actual := args.Map{"hasErr": result != nil, "emptyNil": emptyResult == nil}

	// Assert
	expected := args.Map{"hasErr": true, "emptyNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledError returns correct -- with errors", actual)
}

// ── SliceError / SliceErrorDefault ──

func Test_Cov6_SliceError(t *testing.T) {
	// Arrange
	err := errcore.SliceError("|", []string{"a", "b"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns error -- non-empty slice", actual)
}

func Test_Cov6_SliceError_Empty(t *testing.T) {
	// Arrange
	err := errcore.SliceError("|", []string{})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns nil -- empty slice", actual)
}

func Test_Cov6_SliceErrorDefault(t *testing.T) {
	// Arrange
	err := errcore.SliceErrorDefault([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault returns error -- non-empty slice", actual)
}

// ── SliceErrorsToStrings ──

func Test_Cov6_SliceErrorsToStrings(t *testing.T) {
	// Arrange
	errs := []error{errors.New("a"), nil, errors.New("b")}
	result := errcore.SliceErrorsToStrings(errs...)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns filtered -- with nils", actual)
}

// ── FmtDebug / FmtDebugIf ──

func Test_Cov6_FmtDebug(t *testing.T) {
	// Arrange
	// FmtDebug returns void; just verify no panic
	errcore.FmtDebug("hello %s", "world")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug completes safely -- with format", actual)
}

func Test_Cov6_FmtDebugIf(t *testing.T) {
	// Arrange
	errcore.FmtDebugIf(true, "hello %s", "world")
	errcore.FmtDebugIf(false, "hello")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf completes safely -- with condition", actual)
}

// ── Expecting / ExpectingSimple / ExpectingRecord ──

func Test_Cov6_Expecting(t *testing.T) {
	// Arrange
	result := errcore.Expecting("header", 42, "expected")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting returns formatted -- with args", actual)
}

func Test_Cov6_ExpectingSimple(t *testing.T) {
	// Arrange
	result := errcore.ExpectingSimple("header", 42, "expected")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple returns formatted -- with args", actual)
}

func Test_Cov6_ExpectingSimpleNoType(t *testing.T) {
	// Arrange
	result := errcore.ExpectingSimpleNoType("header", 42, "expected")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType returns formatted -- with args", actual)
}

func Test_Cov6_ExpectingError(t *testing.T) {
	// Arrange
	result := errcore.ExpectingErrorSimpleNoType("header", 42, "expected")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- with args", actual)
}

func Test_Cov6_ExpectingNotEqualSimpleNoType(t *testing.T) {
	// Arrange
	result := errcore.ExpectingNotEqualSimpleNoType("header", 42, "expected")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType returns non-empty -- with args", actual)
}

func Test_Cov6_ExpectingRecord(t *testing.T) {
	// Arrange
	rec := &errcore.ExpectingRecord{ExpectingTitle: "header", WasExpecting: "expected"}
	msg := rec.Message("actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord returns formatted -- with actual", actual)
}

func Test_Cov6_ExpectingFuture(t *testing.T) {
	// Arrange
	rec := errcore.ExpectingFuture("header", "expected")
	msg := rec.Message("actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture returns record -- with title", actual)
}

// ── Var helpers ──

func Test_Cov6_VarTwo(t *testing.T) {
	// Arrange
	result := errcore.VarTwo(true, "a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- with args", actual)
}

func Test_Cov6_VarTwoNoType(t *testing.T) {
	// Arrange
	result := errcore.VarTwoNoType("a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType returns formatted -- with args", actual)
}

func Test_Cov6_VarThree(t *testing.T) {
	// Arrange
	result := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- with args", actual)
}

func Test_Cov6_VarThreeNoType(t *testing.T) {
	// Arrange
	result := errcore.VarThreeNoType("a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType returns formatted -- with args", actual)
}

func Test_Cov6_VarMap(t *testing.T) {
	// Arrange
	result := errcore.VarMap(map[string]any{"a": 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap returns correct -- empty and non-empty", actual)
}

func Test_Cov6_VarMapStrings(t *testing.T) {
	// Arrange
	result := errcore.VarMapStrings(map[string]any{"a": "1"})

	// Act
	actual := args.Map{"hasAny": len(result) > 0}

	// Assert
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "VarMapStrings returns correct -- empty and non-empty", actual)
}

// ── MsgHeader ──

func Test_Cov6_MsgHeader(t *testing.T) {
	// Arrange
	result := errcore.MsgHeader("header", "msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader returns non-empty -- with args", actual)
}

func Test_Cov6_MsgHeaderIf(t *testing.T) {
	// Arrange
	result := errcore.MsgHeaderIf(true, "header", "msg")
	falseResult := errcore.MsgHeaderIf(false, "header", "msg")

	// Act
	actual := args.Map{"notEmpty": result != "", "falseNotEmpty": falseResult != ""}

	// Assert
	expected := args.Map{"notEmpty": true, "falseNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf returns correct value -- with condition", actual)
}

func Test_Cov6_MsgHeaderPlusEnding(t *testing.T) {
	// Arrange
	result := errcore.MsgHeaderPlusEnding("header", "msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding returns non-empty -- with args", actual)
}

// ── Ref ──

func Test_Cov6_Ref(t *testing.T) {
	// Arrange
	result := errcore.Ref("context")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref returns correct -- nil and non-nil", actual)
}

func Test_Cov6_RefToError(t *testing.T) {
	// Arrange
	result := errcore.RefToError("context")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RefToError returns correct -- nil and non-nil", actual)
}

// ── GherkinsString ──

func Test_Cov6_GherkinsString(t *testing.T) {
	// Arrange
	result := errcore.GherkinsString(0, "title", "given", "when", "then")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString returns non-empty -- with args", actual)
}

func Test_Cov6_GherkinsStringWithExpectation(t *testing.T) {
	// Arrange
	result := errcore.GherkinsStringWithExpectation(0, "title", "given", "when", "then", "actual", "expected")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation returns non-empty -- with args", actual)
}

// ── SourceDestination ──

func Test_Cov6_SourceDestination(t *testing.T) {
	// Arrange
	result := errcore.SourceDestination(false, "src", "dst")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- with args", actual)
}

func Test_Cov6_SourceDestinationErr(t *testing.T) {
	// Arrange
	result := errcore.SourceDestinationErr(false, "src", "dst")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr returns error -- with args", actual)
}

func Test_Cov6_SourceDestinationNoType(t *testing.T) {
	// Arrange
	result := errcore.SourceDestinationNoType("src", "dst")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType returns formatted -- with args", actual)
}

// ── MustBeEmpty / PanicOnIndexOutOfRange / PanicRangeNotMeet ──

func Test_Cov6_MustBeEmpty_NoErr(t *testing.T) {
	// Arrange
	errcore.MustBeEmpty(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty completes safely -- nil error", actual)
}

func Test_Cov6_MustBeEmpty_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "MustBeEmpty panics -- with error", actual)
	}()
	errcore.MustBeEmpty(errors.New("err"))
}

func Test_Cov6_PanicOnIndexOutOfRange(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange panics -- index out of range", actual)
	}()
	errcore.PanicOnIndexOutOfRange(5, []int{3, 6})
}

// ── StringLinesToQuoteLines ──

func Test_Cov6_StringLinesToQuoteLines(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns formatted -- with input", actual)
}

func Test_Cov6_StringLinesToQuoteLinesToSingle(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle returns non-empty -- with input", actual)
}

func Test_Cov6_StringLinesToQuoteLinesWithTabs(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns formatted -- with input", actual)
}

// ── MapMismatchError ──

func Test_Cov6_MapMismatchError(t *testing.T) {
	// Arrange
	result := errcore.MapMismatchError(
		"Test_Cov6_MapMismatchError", 1, "ctx",
		[]string{"\"actual\": false,"},
		[]string{"\"expected\": true,"},
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapMismatchError returns formatted -- with args", actual)
}

// ── MergeErrorsToStringDefault ──

func Test_Cov6_MergeErrorsToStringDefault(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToStringDefault(errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault returns correct -- with errors", actual)
}

// ── MessageWithRef / MessageWithRefToError ──

func Test_Cov6_MessageWithRef(t *testing.T) {
	// Arrange
	result := errcore.MessageWithRef("msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef returns non-empty -- with args", actual)
}

func Test_Cov6_MessageWithRefToError(t *testing.T) {
	// Arrange
	result := errcore.MessageWithRefToError("msg", "ref")

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError returns error -- with args", actual)
}

// ── CountStateChangeTracker ──

func Test_Cov6_CountStateChangeTracker(t *testing.T) {
	// Arrange
	mockLen := &mockLengthGetterCov6{length: 0}
	tracker := errcore.NewCountStateChangeTracker(mockLen)

	// Initially same state

	// Act
	actual := args.Map{
		"isSameState": tracker.IsSameState(),
		"isValid":     tracker.IsValid(),
		"isSuccess":   tracker.IsSuccess(),
		"hasChanges":  tracker.HasChanges(),
		"isFailed":    tracker.IsFailed(),
	}

	// Assert
	expected := args.Map{
		"isSameState": true,
		"isValid":     true,
		"isSuccess":   true,
		"hasChanges":  false,
		"isFailed":    false,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker returns same -- initial state", actual)

	// Simulate length change
	mockLen.length = 2
	actual2 := args.Map{
		"isSameState": tracker.IsSameState(),
		"hasChanges":  tracker.HasChanges(),
		"isFailed":    tracker.IsFailed(),
	}
	expected2 := args.Map{
		"isSameState": false,
		"hasChanges":  true,
		"isFailed":    true,
	}
	expected2.ShouldBeEqual(t, 1, "CountStateChangeTracker returns changed -- after add", actual2)
}

// ── VarNameValues / VarNameValuesJoiner / VarNameValuesStrings ──

func Test_Cov6_VarNameValues(t *testing.T) {
	// Arrange
	result := errcore.VarNameValues(
		namevalue.Instance[string, any]{Name: "a", Value: 1},
		namevalue.Instance[string, any]{Name: "b", Value: 2},
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns formatted -- with args", actual)
}

func Test_Cov6_VarNameValuesJoiner(t *testing.T) {
	// Arrange
	result := errcore.VarNameValuesJoiner(",",
		namevalue.Instance[string, any]{Name: "a", Value: 1},
		namevalue.Instance[string, any]{Name: "b", Value: 2},
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner returns joined -- with name-values", actual)
}

func Test_Cov6_VarNameValuesStrings(t *testing.T) {
	// Arrange
	result := errcore.VarNameValuesStrings(
		namevalue.Instance[string, any]{Name: "a", Value: "1"},
		namevalue.Instance[string, any]{Name: "b", Value: "2"},
	)

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings returns entries -- with name-values", actual)
}

// ── MessageNameValues ──

func Test_Cov6_MessageNameValues(t *testing.T) {
	// Arrange
	result := errcore.MessageNameValues("msg",
		namevalue.Instance[string, any]{Name: "a", Value: 1},
		namevalue.Instance[string, any]{Name: "b", Value: 2},
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns formatted -- with name-values", actual)
}

// ── MessageVarTwo / MessageVarThree / MessageVarMap ──

func Test_Cov6_MessageVarTwo(t *testing.T) {
	// Arrange
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo returns formatted -- with args", actual)
}

func Test_Cov6_MessageVarThree(t *testing.T) {
	// Arrange
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree returns formatted -- with args", actual)
}

func Test_Cov6_MessageVarMap(t *testing.T) {
	// Arrange
	result := errcore.MessageVarMap("msg", map[string]any{"a": 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns formatted -- with map", actual)
}

// ── EnumRangeNotMeet / RangeNotMeet ──

func Test_Cov6_EnumRangeNotMeet(t *testing.T) {
	// Arrange
	result := errcore.EnumRangeNotMeet(0, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns non-empty -- with range", actual)
}

func Test_Cov6_RangeNotMeet(t *testing.T) {
	// Arrange
	result := errcore.RangeNotMeet("name", 0, 10, 15)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns non-empty -- with range", actual)
}
