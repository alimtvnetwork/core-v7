package errcoretests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ==========================================================================
// shouldBe — all methods
// ==========================================================================

func Test_ShouldBe_StrEqMsg_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.ShouldBe.StrEqMsg("actual", "expect")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg returns non-empty -- different strings", actual)
}

func Test_ShouldBe_StrEqErr_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.StrEqErr("actual", "expect")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr returns error -- different strings", actual)
}

func Test_ShouldBe_AnyEqMsg_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.ShouldBe.AnyEqMsg(1, 2)

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg returns non-empty -- different values", actual)
}

func Test_ShouldBe_AnyEqErr_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.AnyEqErr(1, 2)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr returns error -- different values", actual)
}

func Test_ShouldBe_JsonEqMsg_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.ShouldBe.JsonEqMsg("a", "b")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg returns non-empty -- different json", actual)
}

func Test_ShouldBe_JsonEqErr_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.JsonEqErr("a", "b")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr returns error -- different json", actual)
}

// ==========================================================================
// expected — all methods
// ==========================================================================

func Test_Expected_But_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.Expected.But("title", "expect", "actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.But returns error -- with args", actual)
}

func Test_Expected_ButFoundAsMsg_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.Expected.ButFoundAsMsg("title", "expect", "actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButFoundWithTypeAsMsg_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.Expected.ButFoundWithTypeAsMsg("title", "expect", "actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButUsingType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.Expected.ButUsingType("title", "expect", "actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType returns error -- with args", actual)
}

func Test_Expected_ReflectButFound_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.Expected.ReflectButFound(reflect.String, reflect.Int)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ReflectButFound returns error -- different kinds", actual)
}

func Test_Expected_PrimitiveButFound_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.Expected.PrimitiveButFound(reflect.Struct)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.PrimitiveButFound returns error -- non-primitive kind", actual)
}

func Test_Expected_ValueHasNoElements_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.Expected.ValueHasNoElements(reflect.Slice)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ValueHasNoElements returns error -- with kind", actual)
}

// ==========================================================================
// CountStateChangeTracker — deeper coverage
// ==========================================================================

func Test_CountStateChangeTracker_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	tracker := errcore.NewCountStateChangeTracker(rec)

	// Act
	actual := args.Map{
		"sameState":     tracker.IsSameState(),
		"isValid":       tracker.IsValid(),
		"isSuccess":     tracker.IsSuccess(),
		"hasChanges":    tracker.HasChanges(),
		"isFailed":      tracker.IsFailed(),
		"sameUsingCount": tracker.IsSameStateUsingCount(1),
	}

	// Assert
	expected := args.Map{
		"sameState":     true,
		"isValid":       true,
		"isSuccess":     true,
		"hasChanges":    false,
		"isFailed":      false,
		"sameUsingCount": true,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker returns same -- no changes", actual)

	rec.Add(errors.New("b"))
	actual2 := args.Map{
		"sameState":  tracker.IsSameState(),
		"hasChanges": tracker.HasChanges(),
		"isFailed":   tracker.IsFailed(),
	}
	expected2 := args.Map{
		"sameState":  false,
		"hasChanges": true,
		"isFailed":   true,
	}
	expected2.ShouldBeEqual(t, 1, "CountStateChangeTracker returns changed -- length increased", actual2)
}

// ==========================================================================
// StackTracesCompiled
// ==========================================================================

func Test_StackTracesCompiled_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.StackTracesCompiled([]string{"trace1", "trace2"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled returns non-empty -- with traces", actual)
}

// ==========================================================================
// stackTraceEnhance — all methods
// ==========================================================================

func Test_StackEnhance_Error(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.Error(errors.New("e"))
	errNil := errcore.StackEnhance.Error(nil)

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns error -- with error", actual)
}

func Test_StackEnhance_ErrorSkip(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.ErrorSkip(0, errors.New("e"))
	errNil := errcore.StackEnhance.ErrorSkip(0, nil)

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "StackEnhance.ErrorSkip returns error -- with error", actual)
}

func Test_StackEnhance_MsgToErrSkip(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.MsgToErrSkip(0, "msg")
	errEmpty := errcore.StackEnhance.MsgToErrSkip(0, "")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errEmpty == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgToErrSkip returns error -- with message", actual)
}

func Test_StackEnhance_FmtSkip(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.FmtSkip(0, "hello %d", 1)
	errEmpty := errcore.StackEnhance.FmtSkip(0, "")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errEmpty == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns error -- with format", actual)
}

func Test_StackEnhance_Msg(t *testing.T) {
	// Arrange
	msg := errcore.StackEnhance.Msg("hello")
	msgEmpty := errcore.StackEnhance.Msg("")

	// Act
	actual := args.Map{
		"notEmpty": msg != "",
		"empty": msgEmpty == "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns non-empty -- with message", actual)
}

func Test_StackEnhance_MsgErrorSkip(t *testing.T) {
	// Arrange
	msg := errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("e"))
	msgNil := errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)

	// Act
	actual := args.Map{
		"notEmpty": msg != "",
		"empty": msgNil == "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip returns non-empty -- with error", actual)
}

func Test_StackEnhance_MsgErrorToErrSkip(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e"))
	errNil := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil)

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns error -- with error", actual)
}

// ==========================================================================
// ExpectingFuture / ExpectingRecord
// ==========================================================================

func Test_ExpectingFuture_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	er := errcore.ExpectingFuture("title", "expect")

	// Act
	actual := args.Map{
		"title": er.ExpectingTitle,
		"was":   er.WasExpecting,
	}

	// Assert
	expected := args.Map{
		"title": "title",
		"was":   "expect",
	}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture returns record -- with title", actual)
}

func Test_ExpectingRecord_Message_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.Message("actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Message returns non-empty -- with actual", actual)
}

func Test_ExpectingRecord_MessageSimple_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.MessageSimple("actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimple returns non-empty -- with actual", actual)
}

func Test_ExpectingRecord_MessageSimpleNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.MessageSimpleNoType("actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimpleNoType returns non-empty -- with actual", actual)
}

func Test_ExpectingRecord_Error_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	er := errcore.ExpectingFuture("title", "expect")
	err := er.Error("actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Error returns error -- with actual", actual)
}

func Test_ExpectingRecord_ErrorSimple_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	er := errcore.ExpectingFuture("title", "expect")
	err := er.ErrorSimple("actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimple returns error -- with actual", actual)
}

func Test_ExpectingRecord_ErrorSimpleNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	er := errcore.ExpectingFuture("title", "expect")
	err := er.ErrorSimpleNoType("actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimpleNoType returns error -- with actual", actual)
}

// ==========================================================================
// Expecting / ExpectingSimple / ExpectingSimpleNoType / ExpectingError*
// ==========================================================================

func Test_Expecting_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.Expecting("title", "expect", "actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting returns formatted -- with args", actual)
}

func Test_ExpectingSimple_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.ExpectingSimple("title", "expect", "actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple returns formatted -- with args", actual)
}

func Test_ExpectingSimpleNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.ExpectingSimpleNoType("title", "expect", "actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType returns formatted -- with args", actual)
}

func Test_ExpectingErrorSimpleNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ExpectingErrorSimpleNoType("title", "expect", "actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- with args", actual)
}

func Test_ExpectingNotEqualSimpleNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	msg := errcore.ExpectingNotEqualSimpleNoType("title", "expect", "actual")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType returns non-empty -- with args", actual)
}

func Test_ExpectingError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ExpectingErrorSimpleNoType("title", "expect", "actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- with args", actual)
}

// ==========================================================================
// RawErrorType — remaining uncovered methods
// ==========================================================================

func Test_RawErrorType_CombineWithAnother_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.CombineWithAnother(errcore.NotFound, "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": string(result) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.CombineWithAnother returns non-empty -- with another type", actual)
}

func Test_RawErrorType_TypesAttach_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.TypesAttach("msg", "str", 42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttach returns non-empty -- with types", actual)
}

func Test_RawErrorType_TypesAttachErr_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.TypesAttachErr("msg", "str", 42)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttachErr returns error -- with types", actual)
}

func Test_RawErrorType_SrcDestination_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.SrcDestination("msg", "src", "sv", "dst", "dv")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestination returns formatted -- with args", actual)
}

func Test_RawErrorType_SrcDestinationErr_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.SrcDestinationErr("msg", "src", "sv", "dst", "dv")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestinationErr returns error -- with args", actual)
}

func Test_RawErrorType_Error_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.Error("msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Error returns error -- with msg and ref", actual)
}

func Test_RawErrorType_ErrorSkip_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorSkip(0, "msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorSkip returns error -- with skip", actual)
}

func Test_RawErrorType_Fmt_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.Fmt("hello %d", 42)
	errEmpty := errcore.InvalidType.Fmt("")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"emptyNotNil": errEmpty != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"emptyNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Fmt returns error -- with format", actual)
}

func Test_RawErrorType_FmtIf(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.FmtIf(true, "hello %d", 42)
	errNil := errcore.InvalidType.FmtIf(false, "hello %d", 42)

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf returns correct value -- with condition", actual)
}

func Test_RawErrorType_MergeError(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeError(errors.New("e"))
	errNil := errcore.InvalidType.MergeError(nil)

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessage(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithMessage(errors.New("e"), "msg")
	errNil := errcore.InvalidType.MergeErrorWithMessage(nil, "msg")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref")
	errNil := errcore.InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithRef(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MergeErrorWithRef(errors.New("e"), "ref")
	errNil := errcore.InvalidType.MergeErrorWithRef(nil, "ref")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef returns error -- with error", actual)
}

func Test_RawErrorType_MsgCsvRef(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.MsgCsvRef("msg", "a", "b")
	resultEmpty := errcore.InvalidType.MsgCsvRef("", "a")
	resultNoRef := errcore.InvalidType.MsgCsvRef("msg")

	// Act
	actual := args.Map{
		"notEmpty":      result != "",
		"emptyMsg":      resultEmpty != "",
		"noRef":         resultNoRef != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty":      true,
		"emptyMsg":      true,
		"noRef":         true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef returns non-empty -- with items", actual)
}

func Test_RawErrorType_MsgCsvRefError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.MsgCsvRefError("msg", "a", "b")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRefError returns error -- with items", actual)
}

func Test_RawErrorType_ErrorRefOnly_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorRefOnly("ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorRefOnly returns error -- with ref", actual)
}

func Test_RawErrorType_Expecting_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.Expecting("expect", "actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Expecting returns error -- with args", actual)
}

func Test_RawErrorType_NoRef(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.NoRef("msg")
	resultEmpty := errcore.InvalidType.NoRef("")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"emptyNotEmpty": resultEmpty != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"emptyNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.NoRef returns non-empty -- with msg", actual)
}

func Test_RawErrorType_ErrorNoRefs_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorNoRefs("msg")
	errEmpty := errcore.InvalidType.ErrorNoRefs("")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"emptyNotNil": errEmpty != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"emptyNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefs returns error -- with msg", actual)
}

func Test_RawErrorType_HandleUsingPanic_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.InvalidType.HandleUsingPanic("msg", "ref")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.HandleUsingPanic panics -- with error", actual)
}

func Test_GetSet(t *testing.T) {
	// Arrange
	r1 := errcore.GetSet(true, errcore.InvalidType, errcore.NotFound)
	r2 := errcore.GetSet(false, errcore.InvalidType, errcore.NotFound)

	// Act
	actual := args.Map{
		"r1": r1,
		"r2": r2,
	}

	// Assert
	expected := args.Map{
		"r1": errcore.InvalidType,
		"r2": errcore.NotFound,
	}
	expected.ShouldBeEqual(t, 0, "GetSet returns correct value -- with condition", actual)
}

func Test_GetSetVariant(t *testing.T) {
	// Arrange
	r1 := errcore.GetSetVariant(true, "a", "b")
	r2 := errcore.GetSetVariant(false, "a", "b")

	// Act
	actual := args.Map{
		"r1": string(r1),
		"r2": string(r2),
	}

	// Assert
	expected := args.Map{
		"r1": "a",
		"r2": "b",
	}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns correct value -- with condition", actual)
}

// ==========================================================================
// GherkinsString / GherkinsStringWithExpectation
// ==========================================================================

func Test_GherkinsString_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.GherkinsString(0, "feature", "given", "when", "then")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString returns non-empty -- with args", actual)
}

func Test_GherkinsStringWithExpectation_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.GherkinsStringWithExpectation(0, "feature", "given", "when", "then", "actual", "expect")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation returns non-empty -- with args", actual)
}

// ==========================================================================
// Message formatting functions
// ==========================================================================

func Test_MessageNameValues_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MessageNameValues("msg", namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns formatted -- with name-values", actual)
}

func Test_MessageVarTwo_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MessageVarTwo("msg", "n1", "v1", "n2", "v2")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo returns formatted -- with args", actual)
}

func Test_MessageVarThree_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MessageVarThree("msg", "n1", "v1", "n2", "v2", "n3", "v3")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree returns formatted -- with args", actual)
}

func Test_MessageVarMap_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MessageVarMap("msg", map[string]any{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns formatted -- with map", actual)
}

func Test_MessageWithRef_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MessageWithRef("msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef returns non-empty -- with args", actual)
}

func Test_MessageWithRefToError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.MessageWithRefToError("msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError returns error -- with args", actual)
}

func Test_VarTwo_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.VarTwo(true, "n1", "v1", "n2", "v2")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- with args", actual)
}

func Test_VarTwoNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.VarTwo(false, "n1", "v1", "n2", "v2")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType returns formatted -- with args", actual)
}

func Test_VarThree_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.VarThree(true, "n1", "v1", "n2", "v2", "n3", "v3")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- with args", actual)
}

func Test_VarThreeNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.VarThree(false, "n1", "v1", "n2", "v2", "n3", "v3")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType returns formatted -- with args", actual)
}

func Test_VarNameValues_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.VarNameValues(namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns formatted -- with args", actual)
}

func Test_VarNameValuesJoiner_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.VarNameValuesJoiner(",", namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner returns joined -- with name-values", actual)
}

func Test_VarNameValuesStrings_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.VarNameValuesStrings(namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings returns entries -- with name-values", actual)
}

// ==========================================================================
// ErrorWith* and Handle* functions
// ==========================================================================

func Test_ErrorWithRef_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithRef(errors.New("e"), "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns formatted -- with error and ref", actual)
}

func Test_ErrorWithRefToError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithRefToError(errors.New("e"), "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError returns error -- with error", actual)
}

func Test_ErrorWithCompiledTraceRef_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "trace", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns non-empty -- with all args", actual)
}

func Test_ErrorWithCompiledTraceRefToError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithCompiledTraceRefToError(errors.New("e"), "trace", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError returns error -- with args", actual)
}

func Test_ErrorWithTracesRefToError(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{"t"}, "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError returns error -- with traces", actual)
}

func Test_HandleErr(t *testing.T) {
	// Arrange
	errcore.HandleErr(nil) // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr completes safely -- nil error", actual)
}

func Test_HandleErr_Panic_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.HandleErr(errors.New("e"))
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErr panics -- with error", actual)
}

func Test_HandleErrMessage(t *testing.T) {
	// Arrange
	errcore.HandleErrMessage("") // no panic for empty string

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage completes safely -- empty message", actual)
}

func Test_SimpleHandleErr(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErr(nil, "msg") // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr completes safely -- nil error", actual)
}

func Test_SimpleHandleErrMany(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErrMany("msg") // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany completes safely -- nil errors", actual)
}

func Test_HandleCompiledErrorGetter(t *testing.T) {
	// Arrange
	errcore.HandleCompiledErrorGetter(nil) // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorGetter completes safely -- nil getter", actual)
}

func Test_HandleErrorGetter(t *testing.T) {
	// Arrange
	errcore.HandleErrorGetter(nil) // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorGetter completes safely -- nil getter", actual)
}

// ==========================================================================
// CombineWithMsgType variants
// ==========================================================================

func Test_CombineWithMsgTypeNoStack_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- with args", actual)
}

func Test_CombineWithMsgTypeStackTrace_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeStackTrace(errcore.InvalidType, "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeStackTrace returns non-empty -- with stack trace", actual)
}

func Test_Combine_Func(t *testing.T) {
	// Arrange
	result := errcore.Combine("errType", "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine returns formatted -- with args", actual)
}

// ==========================================================================
// MeaningFulError / MeaningfulMessageError
// ==========================================================================

func Test_MeaningfulError(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulError(errcore.InvalidType, "fn", errors.New("e"))
	errNil := errcore.MeaningfulError(errcore.InvalidType, "fn", nil)

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"nilNil": errNil == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- with error", actual)
}

func Test_MeaningfulErrorHandle(t *testing.T) {
	// Arrange
	errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", nil) // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorHandle completes safely -- nil error", actual)
}

func Test_MeaningfulErrorWithData(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", errors.New("e"), "data")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData returns error -- with error", actual)
}

func Test_MeaningfulMessageError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulMessageError(errcore.InvalidType, "fn", errors.New("e"), "msg")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError returns error -- with error", actual)
}

// ==========================================================================
// MsgHeader / MsgHeaderIf / MsgHeaderPlusEnding
// ==========================================================================

func Test_MsgHeader_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MsgHeader("msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader returns non-empty -- with args", actual)
}

func Test_MsgHeaderIf_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MsgHeaderIf(true, "msg")
	resultFalse := errcore.MsgHeaderIf(false, "msg")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"empty": resultFalse == "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": false,
	}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf returns correct value -- with condition", actual)
}

func Test_MsgHeaderPlusEnding_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.MsgHeaderPlusEnding("msg", "ending")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding returns non-empty -- with args", actual)
}

// ==========================================================================
// StringLines functions
// ==========================================================================

func Test_StringLinesToQuoteLines_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns formatted -- with input", actual)
}

func Test_StringLinesToQuoteLines_Integrated(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines([]string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns formatted -- integrated test", actual)
}

func Test_StringLinesToQuoteLinesToSingle_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle returns non-empty -- with input", actual)
}

// ==========================================================================
// Print / FmtDebug / FmtDebugIf
// ==========================================================================

func Test_PrintError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	errcore.PrintError(nil)
	errcore.PrintError(errors.New("e"))

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintError completes safely -- with error", actual)
}

func Test_PrintErrorWithTestIndex_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	errcore.PrintErrorWithTestIndex(0, "test", errors.New("e"))

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex completes safely -- with error", actual)
}

func Test_FmtDebug_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	errcore.FmtDebug("msg")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug completes safely -- with format", actual)
}

func Test_FmtDebugIf_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	errcore.FmtDebugIf(true, "msg")
	errcore.FmtDebugIf(false, "msg")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf completes safely -- with condition", actual)
}

// ==========================================================================
// GetActualAndExpect / GetSearchLine / PathMeaningful
// ==========================================================================

func Test_GetActualAndExpectProcessedMessage_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.GetActualAndExpectProcessedMessage(1, "act", "exp", "act-processed", "exp-processed")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectProcessedMessage returns non-empty -- with args", actual)
}

func Test_GetActualAndExpectSortedMessage(t *testing.T) {
	// Arrange
	result := errcore.GetActualAndExpectProcessedMessage(2, []string{"b", "a"}, []string{"a", "b"}, []string{"a", "b"}, []string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectSortedMessage returns non-empty -- with args", actual)
}

func Test_GetSearchLineNumberExpectationMessage_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.GetSearchLineNumberExpectationMessage(1, 10, 9, "line-content", "term", "extra")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage returns non-empty -- with args", actual)
}

func Test_GetSearchTermExpectationMessage_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.GetSearchTermExpectationMessage(1, "header", "expectation", 0, "act", "exp", nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage returns non-empty -- with args", actual)
}

func Test_GetSearchTermExpectationSimpleMessage_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.GetSearchTermExpectationSimpleMessage(1, "expectation", 0, "act", "exp")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationSimpleMessage returns non-empty -- with args", actual)
}

func Test_PathMeaningfulMessage_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulMessage(errcore.InvalidType, "fn", "path", "msg")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage returns error -- with messages", actual)
}

func Test_PathMeaningfulError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulError(errcore.InvalidType, errors.New("boom"), "path")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError returns error -- with error", actual)
}

// ==========================================================================
// Panic / Range functions
// ==========================================================================

func Test_PanicOnIndexOutOfRange_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.PanicOnIndexOutOfRange(-1, []int{10})
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange panics -- index out of range", actual)
}

func Test_PanicOnIndexOutOfRange_Valid(t *testing.T) {
	// Arrange
	errcore.PanicOnIndexOutOfRange(10, []int{0}) // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange completes safely -- index in range", actual)
}

func Test_RangeNotMeet_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.RangeNotMeet("test", 0, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns non-empty -- with range", actual)
}

func Test_EnumRangeNotMeet_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.EnumRangeNotMeet(0, 10, "1,2,3")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns non-empty -- with range", actual)
}

func Test_PanicRangeNotMeet(t *testing.T) {
	// Arrange
	result := errcore.PanicRangeNotMeet("test", 0, 10, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PanicRangeNotMeet returns message -- with range", actual)
}

// ==========================================================================
// ManyErrorToSingle / ManyErrorToSingleDirect
// ==========================================================================

func Test_ManyErrorToSingle_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingle([]error{errors.New("a"), errors.New("b")})

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle returns error -- with errors", actual)
}

func Test_ManyErrorToSingleDirect_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingleDirect(errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect returns error -- with errors", actual)
}

// ==========================================================================
// SourceDestination / SourceDestinationErr / SourceDestinationNoType
// ==========================================================================

func Test_SourceDestination_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.SourceDestination(true, "sv", "dv")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- with args", actual)
}

func Test_SourceDestinationErr_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.SourceDestinationErr(true, "sv", "dv")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr returns error -- with args", actual)
}

func Test_SourceDestinationNoType_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.SourceDestinationNoType("sv", "dv")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType returns formatted -- with args", actual)
}

// ==========================================================================
// CompiledError
// ==========================================================================

func Test_CompiledError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	result := errcore.CompiledError(errors.New("main"), "additional")

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns error -- with message", actual)
}

// ==========================================================================
// ToExitError
// ==========================================================================

func Test_ToExitError_FromShouldBeStrEqMsgIter(t *testing.T) {
	// Arrange
	err := errcore.ToExitError(errors.New("e"))

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true} // non-ExitError returns nil
	expected.ShouldBeEqual(t, 0, "ToExitError returns correct value -- with error type", actual)
}

// ==========================================================================
// ExpectationMessageDef
// ==========================================================================

func Test_ExpectationMessageDef(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{
		CaseIndex: 1,
		FuncName:  "TestFunc",
		When:      "when testing",
		Expected:  "expected-value",
	}
	result := def.ToString("actual-value")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef returns non-empty -- with struct", actual)
}

// ==========================================================================
// HandleCompiledErrorWithTracesGetter / HandleFullStringsWithTracesGetter
// ==========================================================================

func Test_HandleCompiledErrorWithTracesGetter(t *testing.T) {
	// Arrange
	errcore.HandleCompiledErrorWithTracesGetter(nil) // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorWithTracesGetter completes safely -- nil getter", actual)
}

func Test_HandleFullStringsWithTracesGetter(t *testing.T) {
	// Arrange
	errcore.HandleFullStringsWithTracesGetter(nil) // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleFullStringsWithTracesGetter completes safely -- nil getter", actual)
}

// ==========================================================================
// ReferenceStart / ReferenceEnd constants
// ==========================================================================

func Test_ReferenceConstants(t *testing.T) {
	// Act
	actual := args.Map{
		"startNotEmpty": errcore.ReferenceStart != "",
		"endNotEmpty":   errcore.ReferenceEnd != "",
	}

	// Assert
	expected := args.Map{
		"startNotEmpty": true,
		"endNotEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "ReferenceConstants returns non-empty -- defined constants", actual)
}
