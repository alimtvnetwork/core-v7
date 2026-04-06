package errcoretests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── RawErrorType methods ──

func Test_Cov2_RawErrorType_String(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType String returns non-empty -- valid type", actual)
}

func Test_Cov2_RawErrorType_Combine(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.Combine("other msg", "ref-value") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType Combine returns non-empty -- with args", actual)
}

func Test_Cov2_RawErrorType_CombineWithAnother(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.CombineWithAnother(errcore.InvalidEmptyValueType, "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": string(result) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType CombineWithAnother returns non-empty -- with another type", actual)
}

func Test_Cov2_RawErrorType_TypesAttach(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.TypeMismatchType.TypesAttach("msg", "string") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType TypesAttach returns non-empty -- with type", actual)
}

func Test_Cov2_RawErrorType_TypesAttachErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.TypeMismatchType.TypesAttachErr("msg", "string") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType TypesAttachErr returns error -- with type", actual)
}

func Test_Cov2_RawErrorType_SrcDestination(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.SrcDestination("msg", "src", "srcVal", "dst", "dstVal") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType SrcDestination returns non-empty -- with args", actual)
}

func Test_Cov2_RawErrorType_SrcDestinationErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.SrcDestinationErr("msg", "src", "srcVal", "dst", "dstVal") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType SrcDestinationErr returns error -- with args", actual)
}

func Test_Cov2_RawErrorType_Error(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Error("msg", "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType Error returns error -- with msg and ref", actual)
}

func Test_Cov2_RawErrorType_ErrorSkip(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorSkip(0, "msg", "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType ErrorSkip returns error -- skip 0", actual)
}

func Test_Cov2_RawErrorType_Fmt(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Fmt("value %d", 42) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType Fmt returns error -- with format", actual)
}

func Test_Cov2_RawErrorType_Fmt_Empty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Fmt("") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType Fmt returns error -- empty format", actual)
}

func Test_Cov2_RawErrorType_FmtIf_True(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.FmtIf(true, "value %d", 42) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType FmtIf returns error -- isError true", actual)
}

func Test_Cov2_RawErrorType_FmtIf_False(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.FmtIf(false, "value %d", 42) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType FmtIf returns nil -- isError false", actual)
}

func Test_Cov2_RawErrorType_MergeError_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeError(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeError returns nil -- nil error", actual)
}

func Test_Cov2_RawErrorType_MergeError_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeError(errors.New("inner")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeError returns error -- with inner error", actual)
}

func Test_Cov2_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithMessage(nil, "msg") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeErrorWithMessage returns nil -- nil error", actual)
}

func Test_Cov2_RawErrorType_MergeErrorWithMessage_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithMessage(errors.New("inner"), "msg") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeErrorWithMessage returns error -- with error", actual)
}

func Test_Cov2_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithMessageRef(nil, "msg", "ref") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeErrorWithMessageRef returns nil -- nil error", actual)
}

func Test_Cov2_RawErrorType_MergeErrorWithMessageRef_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithMessageRef(errors.New("inner"), "msg", "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeErrorWithMessageRef returns error -- with error", actual)
}

func Test_Cov2_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithRef(nil, "ref") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeErrorWithRef returns nil -- nil error", actual)
}

func Test_Cov2_RawErrorType_MergeErrorWithRef_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithRef(errors.New("inner"), "ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MergeErrorWithRef returns error -- with error", actual)
}

func Test_Cov2_RawErrorType_MsgCsvRef_WithItems(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.MsgCsvRef("msg", "a", "b") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MsgCsvRef returns non-empty -- with items", actual)
}

func Test_Cov2_RawErrorType_MsgCsvRef_NoItems(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.MsgCsvRef("msg") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MsgCsvRef returns non-empty -- no items", actual)
}

func Test_Cov2_RawErrorType_MsgCsvRef_EmptyMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.MsgCsvRef("", "a") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MsgCsvRef returns non-empty -- empty msg", actual)
}

func Test_Cov2_RawErrorType_MsgCsvRefError(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MsgCsvRefError("msg", "a") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType MsgCsvRefError returns error -- with items", actual)
}

func Test_Cov2_RawErrorType_ErrorRefOnly(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorRefOnly("ref") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType ErrorRefOnly returns error -- with ref", actual)
}

func Test_Cov2_RawErrorType_Expecting(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Expecting("expected", "actual") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType Expecting returns error -- with values", actual)
}

func Test_Cov2_RawErrorType_NoRef_WithMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.NoRef("other msg") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType NoRef returns non-empty -- with msg", actual)
}

func Test_Cov2_RawErrorType_NoRef_EmptyMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.NoRef("") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType NoRef returns non-empty -- empty msg", actual)
}

func Test_Cov2_RawErrorType_ErrorNoRefs(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefs("msg") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType ErrorNoRefs returns error -- with msg", actual)
}

func Test_Cov2_RawErrorType_ErrorNoRefs_Empty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefs("") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType ErrorNoRefs returns error -- empty msg", actual)
}

func Test_Cov2_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefsSkip(0, "msg") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType ErrorNoRefsSkip returns error -- with msg", actual)
}

func Test_Cov2_RawErrorType_ErrorNoRefsSkip_Empty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefsSkip(0, "") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType ErrorNoRefsSkip returns error -- empty msg", actual)
}

// ── GetSet / GetSetVariant ──

func Test_Cov2_GetSet_True(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", errcore.GetSet(true, errcore.InvalidRequestType, errcore.InvalidEmptyValueType))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", errcore.InvalidRequestType)}
	expected.ShouldBeEqual(t, 0, "GetSet returns trueValue -- condition true", actual)
}

func Test_Cov2_GetSet_False(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", errcore.GetSet(false, errcore.InvalidRequestType, errcore.InvalidEmptyValueType))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", errcore.InvalidEmptyValueType)}
	expected.ShouldBeEqual(t, 0, "GetSet returns falseValue -- condition false", actual)
}

func Test_Cov2_GetSetVariant_True(t *testing.T) {
	// Act
	actual := args.Map{"result": string(errcore.GetSetVariant(true, "trueVal", "falseVal"))}

	// Assert
	expected := args.Map{"result": "trueVal"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns trueVal -- condition true", actual)
}

func Test_Cov2_GetSetVariant_False(t *testing.T) {
	// Act
	actual := args.Map{"result": string(errcore.GetSetVariant(false, "trueVal", "falseVal"))}

	// Assert
	expected := args.Map{"result": "falseVal"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns falseVal -- condition false", actual)
}

// ── HandleErr / SimpleHandleErr (nil paths only) ──

func Test_Cov2_HandleErr_NilError(t *testing.T) {
	// Arrange
	errcore.HandleErr(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr returns safely -- nil error", actual)
}

func Test_Cov2_SimpleHandleErr_NilError(t *testing.T) {
	// Arrange
	errcore.SimpleHandleErr(nil, "msg") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr returns safely -- nil error", actual)
}

// ── MustBeEmpty (nil path only — panics on non-nil) ──

func Test_Cov2_MustBeEmpty_NilError(t *testing.T) {
	// Arrange
	errcore.MustBeEmpty(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty returns safely -- nil error", actual)
}

// ── MeaningfulError ──

func Test_Cov2_MeaningfulError_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.MeaningfulError(errcore.InvalidRequestType, "funcName", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns nil -- nil error", actual)
}

func Test_Cov2_MeaningfulError_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.MeaningfulError(errcore.InvalidRequestType, "funcName", errors.New("fail")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- with error", actual)
}

// ── PathMeaningfulMessage ──

func Test_Cov2_PathMeaningfulMessage_NoMessages(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.PathMeaningfulMessage(errcore.InvalidRequestType, "fn", "loc") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage returns nil -- no messages", actual)
}

func Test_Cov2_PathMeaningfulMessage_WithMessages(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.PathMeaningfulMessage(errcore.InvalidRequestType, "fn", "loc", "msg1", "msg2") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage returns error -- with messages", actual)
}

// ── MergeErrorsToString / MergeErrorsToStringDefault ──

func Test_Cov2_MergeErrorsToString_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isEmpty": errcore.MergeErrorsToString(",") == ""}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns empty -- nil errors", actual)
}

func Test_Cov2_MergeErrorsToString_WithErrors(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b")) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns joined -- with errors", actual)
}

func Test_Cov2_MergeErrorsToStringDefault(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MergeErrorsToStringDefault(errors.New("a")) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault returns joined -- with error", actual)
}

// ── MergeErrors / ManyErrorToSingle / ManyErrorToSingleDirect ──

func Test_Cov2_MergeErrors_NilSlice(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.MergeErrors() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns nil -- no errors", actual)
}

func Test_Cov2_MergeErrors_WithErrors(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.MergeErrors(errors.New("a"), errors.New("b")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns error -- with errors", actual)
}

func Test_Cov2_ManyErrorToSingle_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ManyErrorToSingle(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle returns nil -- nil slice", actual)
}

func Test_Cov2_ManyErrorToSingleDirect_WithErrors(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ManyErrorToSingleDirect(errors.New("a")) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect returns error -- with error", actual)
}

// ── MessageNameValues / VarNameValues / VarNameValuesStrings ──

func Test_Cov2_MessageNameValues_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MessageNameValues("msg")}

	// Assert
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns message only -- no name values", actual)
}

func Test_Cov2_MessageNameValues_WithValues(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "key", Value: "val"}

	// Act
	actual := args.Map{"notEmpty": errcore.MessageNameValues("msg", nv) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns formatted -- with name value", actual)
}

func Test_Cov2_VarNameValues_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.VarNameValues()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns empty -- no args", actual)
}

func Test_Cov2_VarNameValues_WithValues(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "key", Value: "val"}

	// Act
	actual := args.Map{"notEmpty": errcore.VarNameValues(nv) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns formatted -- with name value", actual)
}

// ── SourceDestination / SourceDestinationErr / SourceDestinationNoType ──

func Test_Cov2_SourceDestination_WithType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.SourceDestination(true, "srcVal", "dstVal") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- with type", actual)
}

func Test_Cov2_SourceDestination_NoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.SourceDestination(false, "srcVal", "dstVal") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- no type", actual)
}

func Test_Cov2_SourceDestinationErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.SourceDestinationErr(true, "srcVal", "dstVal") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr returns error -- with type", actual)
}

func Test_Cov2_SourceDestinationNoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.SourceDestinationNoType("srcVal", "dstVal") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType returns formatted -- values only", actual)
}

// ── StringLinesToQuoteLines / StringLinesToQuoteLinesToSingle ──

func Test_Cov2_StringLinesToQuoteLines(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.StringLinesToQuoteLines([]string{"a", "b"}))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns entries -- 2 lines", actual)
}

func Test_Cov2_StringLinesToQuoteLines_Empty(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.StringLinesToQuoteLines(nil))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns empty -- nil lines", actual)
}

func Test_Cov2_StringLinesToQuoteLinesToSingle(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle returns non-empty -- 2 lines", actual)
}

// ── LineDiff ──

func Test_Cov2_LineDiff_Match(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(diffs),
		"allMatch": diffs[0].Status == "  ",
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"allMatch": true,
	}
	expected.ShouldBeEqual(t, 0, "LineDiff returns matching -- same lines", actual)
}

func Test_Cov2_LineDiff_Mismatch(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"status": diffs[0].Status}

	// Assert
	expected := args.Map{"status": "!!"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns mismatch -- different lines", actual)
}

func Test_Cov2_LineDiff_ExtraActual(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a"})

	// Act
	actual := args.Map{"status": diffs[1].Status}

	// Assert
	expected := args.Map{"status": "+"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns extra actual -- longer actual", actual)
}

func Test_Cov2_LineDiff_MissingExpected(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b"})

	// Act
	actual := args.Map{"status": diffs[1].Status}

	// Assert
	expected := args.Map{"status": "-"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns missing expected -- longer expected", actual)
}

// ── LineDiffToString / HasAnyMismatchOnLines / SliceDiffSummary ──

func Test_Cov2_LineDiffToString_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.LineDiffToString(0, "h", nil, nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns empty -- both nil", actual)
}

func Test_Cov2_LineDiffToString_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.LineDiffToString(0, "h", []string{"a"}, []string{"b"}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns formatted -- mismatched", actual)
}

func Test_Cov2_HasAnyMismatchOnLines_Same(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns false -- same lines", actual)
}

func Test_Cov2_HasAnyMismatchOnLines_Different(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- different lines", actual)
}

func Test_Cov2_HasAnyMismatchOnLines_DiffLength(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- different lengths", actual)
}

func Test_Cov2_SliceDiffSummary_Match(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.SliceDiffSummary([]string{"a"}, []string{"a"})}

	// Assert
	expected := args.Map{"result": "all lines match"}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns all match -- same lines", actual)
}

func Test_Cov2_SliceDiffSummary_Mismatch(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.SliceDiffSummary([]string{"a"}, []string{"b"}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns mismatch info -- different lines", actual)
}

// ── ErrorToLinesLineDiff / PrintLineDiff / PrintLineDiffOnFail ──

func Test_Cov2_ErrorToLinesLineDiff_NilErr(t *testing.T) {
	// Arrange
	result := errcore.ErrorToLinesLineDiff(0, "h", nil, []string{"a"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff returns diff -- nil error vs expected", actual)
}

func Test_Cov2_PrintLineDiff_NoMismatch(t *testing.T) {
	// Arrange
	errcore.PrintLineDiff(0, "h", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiff completes -- matching lines", actual)
}

func Test_Cov2_PrintLineDiffOnFail_NoMismatch(t *testing.T) {
	// Arrange
	errcore.PrintLineDiffOnFail(0, "h", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail skips -- matching lines", actual)
}

func Test_Cov2_PrintLineDiffOnFail_WithMismatch(t *testing.T) {
	// Arrange
	errcore.PrintLineDiffOnFail(0, "h", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail prints -- mismatched lines", actual)
}

// ── PrintDiffOnMismatch ──

func Test_Cov2_PrintDiffOnMismatch_NoMismatch(t *testing.T) {
	// Arrange
	errcore.PrintDiffOnMismatch(0, "h", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch skips -- no mismatch", actual)
}

func Test_Cov2_PrintDiffOnMismatch_WithMismatch(t *testing.T) {
	// Arrange
	errcore.PrintDiffOnMismatch(0, "h", []string{"a"}, []string{"b"}, "context: test")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch prints -- with mismatch and context", actual)
}

// ── GherkinsString ──

func Test_Cov2_GherkinsString(t *testing.T) {
	// Arrange
	result := errcore.GherkinsString(1, "feature", "given", "when", "then")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString returns formatted -- all args", actual)
}

// ── Expecting / ExpectingSimple / ExpectingSimpleNoType / ExpectingNotEqualSimpleNoType ──

func Test_Cov2_Expecting(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.Expecting("title", "expected", "actual") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting returns formatted -- all args", actual)
}

func Test_Cov2_ExpectingSimple(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.ExpectingSimple("title", "expected", "actual") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple returns formatted -- all args", actual)
}

func Test_Cov2_ExpectingSimpleNoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.ExpectingSimpleNoType("title", "expected", "actual") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType returns formatted -- all args", actual)
}

func Test_Cov2_ExpectingNotEqualSimpleNoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.ExpectingNotEqualSimpleNoType("title", "expected", "actual") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType returns formatted -- all args", actual)
}

// ── ExpectingFuture (returns *ExpectingRecord) ──

func Test_Cov2_ExpectingFuture(t *testing.T) {
	// Arrange
	record := errcore.ExpectingFuture("title", "expected")

	// Act
	actual := args.Map{
		"notNil": record != nil,
		"title": record.ExpectingTitle,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"title": "title",
	}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture returns record -- with title", actual)
}

// ── ExpectingRecord struct methods ──

func Test_Cov2_ExpectingRecord_Message(t *testing.T) {
	// Arrange
	record := &errcore.ExpectingRecord{ExpectingTitle: "title", WasExpecting: "expected"}

	// Act
	actual := args.Map{"notEmpty": record.Message("actual") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord Message returns formatted -- all args", actual)
}

func Test_Cov2_ExpectingRecord_MessageSimple(t *testing.T) {
	// Arrange
	record := &errcore.ExpectingRecord{ExpectingTitle: "title", WasExpecting: "expected"}

	// Act
	actual := args.Map{"notEmpty": record.MessageSimple("actual") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord MessageSimple returns formatted -- all args", actual)
}

func Test_Cov2_ExpectingRecord_MessageSimpleNoType(t *testing.T) {
	// Arrange
	record := &errcore.ExpectingRecord{ExpectingTitle: "title", WasExpecting: "expected"}

	// Act
	actual := args.Map{"notEmpty": record.MessageSimpleNoType("actual") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord MessageSimpleNoType returns formatted -- all args", actual)
}

func Test_Cov2_ExpectingRecord_Error(t *testing.T) {
	// Arrange
	record := &errcore.ExpectingRecord{ExpectingTitle: "title", WasExpecting: "expected"}

	// Act
	actual := args.Map{"hasErr": record.Error("actual") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord Error returns error -- with actual", actual)
}

func Test_Cov2_ExpectingRecord_ErrorSimple(t *testing.T) {
	// Arrange
	record := &errcore.ExpectingRecord{ExpectingTitle: "title", WasExpecting: "expected"}

	// Act
	actual := args.Map{"hasErr": record.ErrorSimple("actual") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord ErrorSimple returns error -- with actual", actual)
}

func Test_Cov2_ExpectingRecord_ErrorSimpleNoType(t *testing.T) {
	// Arrange
	record := &errcore.ExpectingRecord{ExpectingTitle: "title", WasExpecting: "expected"}

	// Act
	actual := args.Map{"hasErr": record.ErrorSimpleNoType("actual") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord ErrorSimpleNoType returns error -- with actual", actual)
}

// ── ExpectingError functions ──

func Test_Cov2_ExpectingErrorSimpleNoType(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ExpectingErrorSimpleNoType("t", "e", "a") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- all args", actual)
}

func Test_Cov2_ExpectingErrorSimpleNoTypeNewLineEnds(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ExpectingErrorSimpleNoTypeNewLineEnds("t", "e", "a") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoTypeNewLineEnds returns error -- all args", actual)
}

func Test_Cov2_WasExpectingErrorF(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.WasExpectingErrorF("e", "a", "title %d", 1) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "WasExpectingErrorF returns error -- with format", actual)
}

// ── VarTwo / VarThree / VarTwoNoType / VarThreeNoType ──

func Test_Cov2_VarTwo_WithType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.VarTwo(true, "a", 1, "b", 2) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- with type", actual)
}

func Test_Cov2_VarTwo_NoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.VarTwo(false, "a", 1, "b", 2) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- no type", actual)
}

func Test_Cov2_VarThree_WithType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.VarThree(true, "a", 1, "b", 2, "c", 3) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- with type", actual)
}

func Test_Cov2_VarThree_NoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.VarThree(false, "a", 1, "b", 2, "c", 3) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- no type", actual)
}

func Test_Cov2_VarTwoNoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.VarTwoNoType("a", 1, "b", 2) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType returns formatted -- values only", actual)
}

func Test_Cov2_VarThreeNoType(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.VarThreeNoType("a", 1, "b", 2, "c", 3) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType returns formatted -- values only", actual)
}

// ── VarMap / MessageVarMap / MessageVarTwo / MessageVarThree ──

func Test_Cov2_VarMap_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.VarMap(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarMap returns empty -- nil map", actual)
}

func Test_Cov2_VarMap_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.VarMap(map[string]any{"k": "v"}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap returns formatted -- with entries", actual)
}

func Test_Cov2_MessageVarMap_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MessageVarMap("msg", nil)}

	// Assert
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns msg only -- empty map", actual)
}

func Test_Cov2_MessageVarMap_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MessageVarMap("msg", map[string]any{"k": "v"}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns formatted -- with entries", actual)
}

func Test_Cov2_MessageVarTwo(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MessageVarTwo("msg", "a", 1, "b", 2) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo returns formatted -- with values", actual)
}

func Test_Cov2_MessageVarThree(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree returns formatted -- with values", actual)
}

// ── ShouldBe methods ──

func Test_Cov2_ShouldBe_StrEqMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.ShouldBe.StrEqMsg("a", "b") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe StrEqMsg returns formatted -- different strings", actual)
}

func Test_Cov2_ShouldBe_StrEqErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ShouldBe.StrEqErr("a", "b") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe StrEqErr returns error -- different strings", actual)
}

func Test_Cov2_ShouldBe_AnyEqMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.ShouldBe.AnyEqMsg(1, 2) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe AnyEqMsg returns formatted -- different values", actual)
}

func Test_Cov2_ShouldBe_AnyEqErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ShouldBe.AnyEqErr(1, 2) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe AnyEqErr returns error -- different values", actual)
}

func Test_Cov2_ShouldBe_JsonEqMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.ShouldBe.JsonEqMsg("a", "b") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe JsonEqMsg returns formatted -- different values", actual)
}

func Test_Cov2_ShouldBe_JsonEqErr(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ShouldBe.JsonEqErr("a", "b") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe JsonEqErr returns error -- different values", actual)
}

// ── Expected methods ──

func Test_Cov2_Expected_But(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.Expected.But("t", "e", "a") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected But returns error -- all args", actual)
}

func Test_Cov2_Expected_ButFoundAsMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.Expected.ButFoundAsMsg("t", "e", "a") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected ButFoundAsMsg returns message -- all args", actual)
}

func Test_Cov2_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.Expected.ButFoundWithTypeAsMsg("t", "e", "a") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected ButFoundWithTypeAsMsg returns message -- all args", actual)
}

func Test_Cov2_Expected_ButUsingType(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.Expected.ButUsingType("t", "e", "a") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected ButUsingType returns error -- all args", actual)
}

// ── ToString / ToStringPtr / ToError / ToValueString ──

func Test_Cov2_ToString_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ToString(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToString returns empty -- nil error", actual)
}

func Test_Cov2_ToString_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ToString(errors.New("fail"))}

	// Assert
	expected := args.Map{"result": "fail"}
	expected.ShouldBeEqual(t, 0, "ToString returns message -- with error", actual)
}

func Test_Cov2_ToStringPtr_NilErr(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(nil)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"empty": *result == "",
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns empty ptr -- nil error", actual)
}

func Test_Cov2_ToStringPtr_WithErr(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(errors.New("fail"))

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"value": *result,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"value": "fail",
	}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns error ptr -- with error", actual)
}

func Test_Cov2_ToError_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ToError("") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError returns nil -- empty message", actual)
}

func Test_Cov2_ToError_WithMsg(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.ToError("fail") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToError returns error -- with message", actual)
}

// ── ErrorToSplitLines / ErrorToSplitNonEmptyLines ──

func Test_Cov2_ErrorToSplitLines_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.ErrorToSplitLines(nil))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines returns empty -- nil error", actual)
}

func Test_Cov2_ErrorToSplitLines_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.ErrorToSplitLines(errors.New("a\nb")))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines returns 2 -- multiline error", actual)
}

func Test_Cov2_ErrorToSplitNonEmptyLines_NilErr(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.ErrorToSplitNonEmptyLines(nil))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines returns empty -- nil error", actual)
}

// ── SliceError / SliceErrorDefault / SliceToError / SliceToErrorPtr / SliceErrorsToStrings ──

func Test_Cov2_SliceError_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceError(",", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns nil -- empty slice", actual)
}

func Test_Cov2_SliceError_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.SliceError(",", []string{"a", "b"}) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns error -- with items", actual)
}

func Test_Cov2_SliceErrorDefault_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.SliceErrorDefault([]string{"a"}) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault returns error -- with items", actual)
}

func Test_Cov2_SliceToError_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceToError(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns nil -- empty slice", actual)
}

func Test_Cov2_SliceToError_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.SliceToError([]string{"a"}) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns error -- with items", actual)
}

func Test_Cov2_SliceToErrorPtr_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceToErrorPtr(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns nil -- empty slice", actual)
}

func Test_Cov2_SliceToErrorPtr_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": errcore.SliceToErrorPtr([]string{"a"}) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns error -- with items", actual)
}

func Test_Cov2_SliceErrorsToStrings_Nil(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.SliceErrorsToStrings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns empty -- nil", actual)
}

func Test_Cov2_SliceErrorsToStrings_WithNilAndErrors(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.SliceErrorsToStrings(nil, errors.New("a"), nil, errors.New("b")))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns 2 -- filtering nils", actual)
}

// ── Ref / MessageWithRef ──

func Test_Cov2_Ref_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.Ref(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Ref returns empty -- nil reference", actual)
}

func Test_Cov2_Ref_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.Ref("val") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref returns formatted -- with reference", actual)
}

func Test_Cov2_MessageWithRef(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.MessageWithRef("msg", "ref") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef returns formatted -- with msg and ref", actual)
}

// ── RawErrCollection ──

func Test_Cov2_RawErrCollection_AddNilAndNonNil(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Add(nil)
	hasErrAfterNil := c.HasError()
	c.Add(errors.New("err"))

	// Act
	actual := args.Map{
		"hasErrAfterNil":   hasErrAfterNil,
		"hasErrAfterError": c.HasError(),
		"lengthAtLeast1":   c.Length() >= 1,
	}

	// Assert
	expected := args.Map{
		"hasErrAfterNil":   false,
		"hasErrAfterError": true,
		"lengthAtLeast1":   true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection Add returns expected -- nil and non-nil", actual)
}

func Test_Cov2_RawErrCollection_AddIf(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.AddIf(false, "skipped")
	c.AddIf(true, "added")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection AddIf returns 1 -- conditional add", actual)
}

func Test_Cov2_RawErrCollection_AddFunc_Nil(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.AddFunc(nil)

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection AddFunc skips -- nil func", actual)
}

func Test_Cov2_RawErrCollection_AddFuncIf(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.AddFuncIf(false, func() error { return errors.New("a") })
	c.AddFuncIf(true, func() error { return errors.New("b") })

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection AddFuncIf returns 1 -- conditional", actual)
}

func Test_Cov2_RawErrCollection_HasAnyIssues(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}

	// Act
	actual := args.Map{"hasIssues": c.HasAnyIssues()}

	// Assert
	expected := args.Map{"hasIssues": false}
	expected.ShouldBeEqual(t, 0, "RawErrCollection HasAnyIssues returns false -- empty", actual)
}

func Test_Cov2_RawErrCollection_IsNull(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}

	// Act
	actual := args.Map{"isNull": c.IsNull()}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection IsNull returns true -- nil items", actual)
}

func Test_Cov2_RawErrCollection_Log_Empty(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Log() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection Log skips -- empty", actual)
}

func Test_Cov2_RawErrCollection_Fmt(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Fmt("error %d", 42)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection Fmt adds entry -- with format", actual)
}

func Test_Cov2_RawErrCollection_Fmt_Empty(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Fmt("")

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection Fmt skips -- empty format", actual)
}

func Test_Cov2_RawErrCollection_FmtIf(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.FmtIf(false, "skipped %d", 1)
	c.FmtIf(true, "added %d", 2)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection FmtIf returns 1 -- conditional", actual)
}

// ── ExpectationMessageDef ──

func Test_Cov2_ExpectationMessageDef_ExpectedSafeString(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{Expected: "hello"}

	// Act
	actual := args.Map{"result": def.ExpectedSafeString()}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef ExpectedSafeString returns value -- with expected", actual)
}

func Test_Cov2_ExpectationMessageDef_ExpectedSafeString_Nil(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{}

	// Act
	actual := args.Map{"result": def.ExpectedSafeString()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef ExpectedSafeString returns empty -- nil expected", actual)
}

func Test_Cov2_ExpectationMessageDef_ExpectedStringTrim(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{Expected: "  hello  "}

	// Act
	actual := args.Map{"result": def.ExpectedStringTrim()}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef ExpectedStringTrim returns trimmed -- with spaces", actual)
}

func Test_Cov2_ExpectationMessageDef_PrintIf_False(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{Expected: "e", When: "w"}
	def.PrintIf(false, "actual")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef PrintIf skips -- isPrint false", actual)
}

func Test_Cov2_ExpectationMessageDef_PrintIfFailed_NotFailed(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{Expected: "e", When: "w"}
	def.PrintIfFailed(true, false, "actual")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef PrintIfFailed skips -- not failed", actual)
}

// ── Combine (package-level) ──

func Test_Cov2_Combine(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.Combine("generic", "other", "ref") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine returns formatted -- all string args", actual)
}

// ── HandleErrMessage (empty path) ──

func Test_Cov2_HandleErrMessage_Empty(t *testing.T) {
	// Arrange
	errcore.HandleErrMessage("")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage skips -- empty message", actual)
}

// ── MapMismatchError ──

func Test_Cov2_MapMismatchError(t *testing.T) {
	// Arrange
	result := errcore.MapMismatchError("TestFunc", 1, "title", []string{`"k":"a"`}, []string{`"k":"e"`})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapMismatchError returns formatted -- with entries", actual)
}
