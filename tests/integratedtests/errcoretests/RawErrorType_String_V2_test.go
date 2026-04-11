package errcoretests

import (
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── RawErrorType methods ──

func Test_RawErrorType_String_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.InvalidRequestType.String()}

	// Assert
	expected := args.Map{"result": "Invalid : request, cannot process it."}
	expected.ShouldBeEqual(t, 0, "RawErrorType.String returns non-empty -- with type", actual)
}

func Test_RawErrorType_Combine_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.Combine("details", "ref")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"containsType": strings.Contains(result, "Invalid"),
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"containsType": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Combine returns formatted -- with msg and ref", actual)
}

func Test_RawErrorType_CombineWithAnother_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.CombineWithAnother(errcore.NotFound, "msg", "ref")

	// Act
	actual := args.Map{"notEmpty": string(result) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.CombineWithAnother returns non-empty -- with another type", actual)
}

func Test_RawErrorType_TypesAttach_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.TypesAttach("msg", "hello", 42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttach returns non-empty -- with types", actual)
}

func Test_RawErrorType_TypesAttachErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.TypesAttachErr("msg", "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttachErr returns error -- with types", actual)
}

func Test_RawErrorType_SrcDestination_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.SrcDestination("msg", "src", "srcVal", "dst", "dstVal")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestination returns formatted -- with args", actual)
}

func Test_RawErrorType_SrcDestinationErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.SrcDestinationErr("msg", "src", "srcVal", "dst", "dstVal")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestinationErr returns error -- with args", actual)
}

func Test_RawErrorType_Error_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.Error("details", "ref")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Error returns error -- with msg and ref", actual)
}

func Test_RawErrorType_ErrorSkip_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.ErrorSkip(0, "details", "ref")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorSkip returns error -- with skip", actual)
}

func Test_RawErrorType_Fmt(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.Fmt("value: %d", 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Fmt returns error -- with format", actual)
}

func Test_RawErrorType_Fmt_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.Fmt("")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Fmt returns error -- empty format", actual)
}

func Test_RawErrorType_FmtIf_True_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.FmtIf(true, "val: %d", 1)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf returns error -- condition true", actual)
}

func Test_RawErrorType_FmtIf_False_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.FmtIf(false, "val: %d", 1)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf returns nil -- condition false", actual)
}

func Test_RawErrorType_MergeError_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeError(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeError_WithErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeError(errors.New("inner"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithMessage(nil, "msg") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_WithErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithMessage(errors.New("inner"), "msg")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithMessageRef(nil, "msg", "ref") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_WithErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(errors.New("inner"), "msg", "ref")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithRef_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithRef(nil, "ref") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithRef_WithErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithRef(errors.New("inner"), "ref")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef returns error -- with error", actual)
}

func Test_RawErrorType_MsgCsvRef_NoItems(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.MsgCsvRef("msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef returns non-empty -- no items", actual)
}

func Test_RawErrorType_MsgCsvRef_WithItems_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.MsgCsvRef("msg", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef returns non-empty -- with items", actual)
}

func Test_RawErrorType_MsgCsvRef_EmptyMsg_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.MsgCsvRef("", "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef returns non-empty -- empty msg", actual)
}

func Test_RawErrorType_MsgCsvRefError_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MsgCsvRefError("msg", "a")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRefError returns error -- with items", actual)
}

func Test_RawErrorType_ErrorRefOnly_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.ErrorRefOnly("ref")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorRefOnly returns error -- with ref", actual)
}

func Test_RawErrorType_ErrorRefOnly_Nil(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.ErrorRefOnly(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorRefOnly returns error -- nil ref", actual)
}

func Test_RawErrorType_Expecting_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.Expecting("expected", "actual")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Expecting returns error -- with args", actual)
}

func Test_RawErrorType_NoRef_EmptyMsg(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.NoRef("")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.NoRef returns non-empty -- empty msg", actual)
}

func Test_RawErrorType_NoRef_WithMsg_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.NoRef("msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.NoRef returns non-empty -- with msg", actual)
}

func Test_RawErrorType_ErrorNoRefs_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.ErrorNoRefs("msg")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefs returns error -- with msg", actual)
}

func Test_RawErrorType_ErrorNoRefs_EmptyMsg(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.ErrorNoRefs("")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefs returns error -- empty msg", actual)
}

func Test_RawErrorType_ErrorNoRefsSkip_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.ErrorNoRefsSkip(0, "msg")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefsSkip returns error -- with skip", actual)
}

func Test_RawErrorType_HandleUsingPanic_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "RawErrorType.HandleUsingPanic panics -- with error", actual)
	}()
	errcore.InvalidRequestType.HandleUsingPanic("msg", "ref")
}

// ── GetSet / GetSetVariant ──

func Test_GetSet_True_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.GetSet(true, errcore.InvalidRequestType, errcore.NotFound)

	// Act
	actual := args.Map{"result": result.String()}

	// Assert
	expected := args.Map{"result": errcore.InvalidRequestType.String()}
	expected.ShouldBeEqual(t, 0, "GetSet returns trueValue -- condition true", actual)
}

func Test_GetSet_False_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.GetSet(false, errcore.InvalidRequestType, errcore.NotFound)

	// Act
	actual := args.Map{"result": result.String()}

	// Assert
	expected := args.Map{"result": errcore.NotFound.String()}
	expected.ShouldBeEqual(t, 0, "GetSet returns falseValue -- condition false", actual)
}

func Test_GetSetVariant_True_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.GetSetVariant(true, "yes", "no")

	// Act
	actual := args.Map{"result": result.String()}

	// Assert
	expected := args.Map{"result": "yes"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns trueValue -- condition true", actual)
}

func Test_GetSetVariant_False_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.GetSetVariant(false, "yes", "no")

	// Act
	actual := args.Map{"result": result.String()}

	// Assert
	expected := args.Map{"result": "no"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns falseValue -- condition false", actual)
}

// ── VarTwo / VarThree ──

func Test_VarTwo_WithType(t *testing.T) {
	// Arrange
	result := errcore.VarTwo(true, "a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- with type", actual)
}

func Test_VarTwo_NoType(t *testing.T) {
	// Arrange
	result := errcore.VarTwo(false, "a", 1, "b", 2)

	// Act
	actual := args.Map{"contains": strings.Contains(result, "a")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- no type", actual)
}

func Test_VarThree_WithType(t *testing.T) {
	// Arrange
	result := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- with type", actual)
}

func Test_VarThree_NoType(t *testing.T) {
	// Arrange
	result := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"contains": strings.Contains(result, "a")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- no type", actual)
}

func Test_VarTwoNoType_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.VarTwoNoType("x", 10, "y", 20)

	// Act
	actual := args.Map{"contains": strings.Contains(result, "x")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType returns formatted -- with args", actual)
}

func Test_VarThreeNoType_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.VarThreeNoType("x", 1, "y", 2, "z", 3)

	// Act
	actual := args.Map{"contains": strings.Contains(result, "x")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType returns formatted -- with args", actual)
}

// ── VarMap / MessageVarMap / MessageVarTwo / MessageVarThree ──

func Test_VarMap_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.VarMap(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarMap returns empty -- nil map", actual)
}

func Test_VarMap_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.VarMap(map[string]any{"key": "val"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap returns formatted -- with entries", actual)
}

func Test_MessageVarMap_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MessageVarMap("msg", nil)}

	// Assert
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns msg only -- empty map", actual)
}

func Test_MessageVarMap_NonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageVarMap("msg", map[string]any{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns formatted -- with map", actual)
}

func Test_MessageVarTwo_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo returns formatted -- with args", actual)
}

func Test_MessageVarThree_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree returns formatted -- with args", actual)
}

// ── VarNameValues / MessageNameValues ──

func Test_VarNameValues_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.VarNameValues()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns empty -- no args", actual)
}

func Test_VarNameValues_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "k", Value: "v"}

	// Act
	actual := args.Map{"notEmpty": errcore.VarNameValues(nv) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns formatted -- with args", actual)
}

func Test_MessageNameValues_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MessageNameValues("msg")}

	// Assert
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns msg only -- no name-values", actual)
}

func Test_MessageNameValues_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "k", Value: "v"}
	result := errcore.MessageNameValues("msg", nv)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns formatted -- with name-values", actual)
}

// ── Expecting functions ──

func Test_Expecting_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.Expecting("title", "expected", "actual")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting returns formatted -- with args", actual)
}

func Test_ExpectingSimple_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ExpectingSimple("title", "expected", "actual")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple returns formatted -- with args", actual)
}

func Test_ExpectingSimpleNoType_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ExpectingSimpleNoType("title", "expected", "actual")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType returns formatted -- with args", actual)
}

func Test_ExpectingErrorSimpleNoType(t *testing.T) {
	// Arrange
	err := errcore.ExpectingErrorSimpleNoType("title", "exp", "act")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- with args", actual)
}

func Test_ExpectingErrorSimpleNoTypeNewLineEnds_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.ExpectingErrorSimpleNoTypeNewLineEnds("title", "exp", "act")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"endsNewLine": strings.HasSuffix(err.Error(), "\n"),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"endsNewLine": true,
	}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoTypeNewLineEnds returns error -- with args", actual)
}

func Test_WasExpectingErrorF_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.WasExpectingErrorF("exp", "act", "title %d", 1)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "WasExpectingErrorF returns error -- with format", actual)
}

func Test_ExpectingSimpleNoTypeError_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.ExpectingSimpleNoTypeError("title", "exp", "act")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoTypeError returns error -- with args", actual)
}

func Test_ExpectingNotEqualSimpleNoType_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ExpectingNotEqualSimpleNoType("title", "exp", "act")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType returns non-empty -- with args", actual)
}

// ── ExpectingFuture / ExpectingRecord ──

func Test_ExpectingFuture_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	rec := errcore.ExpectingFuture("title", "expected")

	// Act
	actual := args.Map{
		"msg":            rec.Message("actual") != "",
		"msgSimple":      rec.MessageSimple("actual") != "",
		"msgSimpleNoType": rec.MessageSimpleNoType("actual") != "",
		"err":            rec.Error("actual") != nil,
		"errSimple":      rec.ErrorSimple("actual") != nil,
		"errSimpleNoType": rec.ErrorSimpleNoType("actual") != nil,
	}

	// Assert
	expected := args.Map{
		"msg":            true,
		"msgSimple":      true,
		"msgSimpleNoType": true,
		"err":            true,
		"errSimple":      true,
		"errSimpleNoType": true,
	}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture/ExpectingRecord returns record -- with title and actual", actual)
}

// ── expected struct methods ──

func Test_Expected_But_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.Expected.But("title", "exp", "act")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.But returns error -- with args", actual)
}

func Test_Expected_ButFoundAsMsg_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.Expected.ButFoundAsMsg("title", "exp", "act")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButFoundWithTypeAsMsg_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.Expected.ButFoundWithTypeAsMsg("title", "exp", "act")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButUsingType_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.Expected.ButUsingType("title", "exp", "act")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType returns error -- with args", actual)
}

func Test_Expected_ReflectButFound_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.Expected.ReflectButFound(1, 2) // reflect.Kind values

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.ReflectButFound returns error -- different kinds", actual)
}

func Test_Expected_PrimitiveButFound_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.Expected.PrimitiveButFound(20) // reflect.Kind

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.PrimitiveButFound returns error -- non-primitive kind", actual)
}

func Test_Expected_ValueHasNoElements_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.Expected.ValueHasNoElements(23) // reflect.Kind

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.ValueHasNoElements returns error -- with kind", actual)
}

// ── shouldBe struct methods ──

func Test_ShouldBe_StrEqMsg_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.StrEqMsg("actual", "expecting")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg returns non-empty -- different strings", actual)
}

func Test_ShouldBe_StrEqErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.StrEqErr("actual", "expecting")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr returns error -- different strings", actual)
}

func Test_ShouldBe_AnyEqMsg_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.AnyEqMsg(1, 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg returns non-empty -- different values", actual)
}

func Test_ShouldBe_AnyEqErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.AnyEqErr(1, 2)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr returns error -- different values", actual)
}

func Test_ShouldBe_JsonEqMsg_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.JsonEqMsg("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg returns non-empty -- different json", actual)
}

func Test_ShouldBe_JsonEqErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.JsonEqErr("a", "b")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr returns error -- different json", actual)
}

// ── Slice functions ──

func Test_SliceToError_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceToError(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns nil -- empty slice", actual)
}

func Test_SliceToError_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.SliceToError([]string{"e1", "e2"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns error -- non-empty slice", actual)
}

func Test_SliceToErrorPtr_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceToErrorPtr(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns nil -- empty slice", actual)
}

func Test_SliceToErrorPtr_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.SliceToErrorPtr([]string{"e1"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns error -- non-empty slice", actual)
}

func Test_SliceError_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceError(",", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns nil -- empty slice", actual)
}

func Test_SliceError_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.SliceError(",", []string{"a", "b"})

	// Act
	actual := args.Map{"result": err.Error()}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "SliceError returns error -- non-empty slice", actual)
}

func Test_SliceErrorDefault_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.SliceErrorDefault([]string{"a"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault returns error -- non-empty slice", actual)
}

func Test_SliceErrorsToStrings_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"len": len(errcore.SliceErrorsToStrings(nil...))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns empty -- no errors", actual)
}

func Test_SliceErrorsToStrings_Mixed(t *testing.T) {
	// Arrange
	result := errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b"))

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns filtered -- mixed nil and non-nil", actual)
}

// ── MergeErrors / MergeErrorsToString ──

func Test_MergeErrors_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.MergeErrors(nil, nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns nil -- all nil", actual)
}

func Test_MergeErrors_NonNil(t *testing.T) {
	// Arrange
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns error -- with errors", actual)
}

func Test_MergeErrorsToString_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MergeErrorsToString(",", nil...)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns empty -- no errors", actual)
}

func Test_MergeErrorsToString_NonNil(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns non-empty -- with errors", actual)
}

func Test_MergeErrorsToStringDefault_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MergeErrorsToStringDefault(nil...)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault returns empty -- no errors", actual)
}

func Test_MergeErrorsToStringDefault_NonNil(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToStringDefault(errors.New("a"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault returns non-empty -- with errors", actual)
}

// ── ToString / ToStringPtr / ToError / Ref ──

func Test_ToString_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ToString(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToString returns empty -- nil error", actual)
}

func Test_ToString_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ToString(errors.New("hello"))}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ToString returns msg -- with error", actual)
}

func Test_ToStringPtr_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(nil)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns empty ptr -- nil error", actual)
}

func Test_ToStringPtr_NonNil(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(errors.New("hi"))

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "hi"}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns ptr -- with error", actual)
}

func Test_ToError_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ToError("") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError returns nil -- empty string", actual)
}

func Test_ToError_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.ToError("msg")

	// Act
	actual := args.Map{"msg": err.Error()}

	// Assert
	expected := args.Map{"msg": "msg"}
	expected.ShouldBeEqual(t, 0, "ToError returns error -- non-empty string", actual)
}

func Test_Ref_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.Ref(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Ref returns empty -- nil input", actual)
}

func Test_Ref_NonNil_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.Ref("val")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref returns non-empty -- with value", actual)
}

// ── MessageWithRef ──

func Test_MessageWithRef_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.MessageWithRef("msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef returns non-empty -- with args", actual)
}

// ── SourceDestination / SourceDestinationErr / SourceDestinationNoType ──

func Test_SourceDestination_WithType(t *testing.T) {
	// Arrange
	result := errcore.SourceDestination(true, "src", "dst")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- with type", actual)
}

func Test_SourceDestination_NoType(t *testing.T) {
	// Arrange
	result := errcore.SourceDestination(false, "src", "dst")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- no type", actual)
}

func Test_SourceDestinationErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.SourceDestinationErr(false, "src", "dst")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr returns error -- with args", actual)
}

func Test_SourceDestinationNoType_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.SourceDestinationNoType("src", "dst")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType returns formatted -- with args", actual)
}

// ── stackTraceEnhance ──

func Test_StackEnhance_Error_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.Error(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns nil -- nil error", actual)
}

func Test_StackEnhance_Error_NonNil_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.Error(errors.New("test"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns error -- with error", actual)
}

func Test_StackEnhance_Msg_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.StackEnhance.Msg("")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns empty -- empty message", actual)
}

func Test_StackEnhance_Msg_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.StackEnhance.Msg("test")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns non-empty -- with message", actual)
}

func Test_StackEnhance_MsgToErrSkip_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.MsgToErrSkip(0, "") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgToErrSkip returns nil -- empty message", actual)
}

func Test_StackEnhance_FmtSkip_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.FmtSkip(0, "") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns nil -- empty format", actual)
}

func Test_StackEnhance_FmtSkip_NonEmpty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.FmtSkip(0, "val %d", 1)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns error -- with format", actual)
}

func Test_StackEnhance_MsgErrorSkip_NilErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip returns empty -- nil error", actual)
}

func Test_StackEnhance_MsgErrorToErrSkip_NilErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns nil -- nil error", actual)
}

func Test_StackEnhance_MsgErrorToErrSkip_WithErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("inner"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns error -- with error", actual)
}

// ── LineDiff ──

func Test_LineDiff_Matching(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(diffs),
		"status0": diffs[0].Status,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"status0": "  ",
	}
	expected.ShouldBeEqual(t, 0, "LineDiff returns all-match -- matching lines", actual)
}

func Test_LineDiff_Mismatch_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"status": diffs[0].Status}

	// Assert
	expected := args.Map{"status": "!!"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns mismatch -- different line", actual)
}

func Test_LineDiff_ExtraActual_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a"})

	// Act
	actual := args.Map{"status1": diffs[1].Status}

	// Assert
	expected := args.Map{"status1": "+"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns extra-actual -- longer actual", actual)
}

func Test_LineDiff_MissingExpected_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b"})

	// Act
	actual := args.Map{"status1": diffs[1].Status}

	// Assert
	expected := args.Map{"status1": "-"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns missing-expected -- shorter actual", actual)
}

func Test_LineDiffToString_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "test", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns formatted -- with diffs", actual)
}

func Test_LineDiffToString_Empty_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "test", nil, nil)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns empty -- both empty", actual)
}

func Test_HasAnyMismatchOnLines_Same(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns false -- matching", actual)
}

func Test_HasAnyMismatchOnLines_Different(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- different content", actual)
}

func Test_HasAnyMismatchOnLines_DiffLen_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- different length", actual)
}

func Test_SliceDiffSummary_Match_FromRawErrorTypeStringV2(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.SliceDiffSummary([]string{"a"}, []string{"a"})}

	// Assert
	expected := args.Map{"result": "all lines match"}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns all-match -- matching", actual)
}

func Test_SliceDiffSummary_Mismatch_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.SliceDiffSummary([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"contains": strings.Contains(result, "mismatch")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns mismatch count -- with mismatch", actual)
}

func Test_ErrorToLinesLineDiff_NilErr_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	result := errcore.ErrorToLinesLineDiff(0, "test", nil, []string{"a"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff returns non-empty -- nil error", actual)
}

// ── RawErrCollection ──

func Test_RawErrCollection_BasicOps(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.AddError(errors.New("e2"))
	c.Add(nil) // should skip
	c.AddString("e3")
	c.AddString("") // should skip

	// Act
	actual := args.Map{
		"len":          c.Length(),
		"isEmpty":      c.IsEmpty(),
		"hasError":     c.HasError(),
		"hasAnyError":  c.HasAnyError(),
		"hasAnyIssues": c.HasAnyIssues(),
		"isDefined":    c.IsDefined(),
		"isValid":      c.IsValid(),
		"isSuccess":    c.IsSuccess(),
		"isFailed":     c.IsFailed(),
		"isInvalid":    c.IsInvalid(),
		"isNull":       c.IsNull(),
		"isAnyNull":    c.IsAnyNull(),
		"isCollType":   c.IsCollectionType(),
	}

	// Assert
	expected := args.Map{
		"len":          3,
		"isEmpty":      false,
		"hasError":     true,
		"hasAnyError":  true,
		"hasAnyIssues": true,
		"isDefined":    true,
		"isValid":      false,
		"isSuccess":    false,
		"isFailed":     true,
		"isInvalid":    true,
		"isNull":       false,
		"isAnyNull":    false,
		"isCollType":   true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection returns correct state -- basic operations", actual)
}

func Test_RawErrCollection_StringOps(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.Add(errors.New("e2"))

	// Act
	actual := args.Map{
		"stringNotEmpty":    c.String() != "",
		"errorStringEq":    c.ErrorString() == c.String(),
		"compileEq":        c.Compile() == c.String(),
		"fullStringEq":     c.FullString() == c.String(),
		"stringsLen":       len(c.Strings()),
		"splitLen":         len(c.FullStringSplitByNewLine()),
		"joinerNotEmpty":   c.StringUsingJoiner(",") != "",
		"joinerAddlNotEm":  c.StringUsingJoinerAdditional(",", "!") != "",
		"withAddlNotEmpty": c.StringWithAdditionalMessage("!") != "",
		"refCompStr":       c.ReferencesCompiledString() != "",
		"fullWithoutRef":   c.FullStringWithoutReferences() != "",
	}

	// Assert
	expected := args.Map{
		"stringNotEmpty":    true,
		"errorStringEq":    true,
		"compileEq":        true,
		"fullStringEq":     true,
		"stringsLen":       2,
		"splitLen":         2,
		"joinerNotEmpty":   true,
		"joinerAddlNotEm":  true,
		"withAddlNotEmpty": true,
		"refCompStr":       true,
		"fullWithoutRef":   true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.String returns correct -- string operations", actual)
}

func Test_RawErrCollection_CompiledErrors(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))

	// Act
	actual := args.Map{
		"compiledErr":       c.CompiledError() != nil,
		"compiledJoiner":    c.CompiledErrorUsingJoiner(",") != nil,
		"compiledJoinerAdd": c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") != nil,
		"compiledStacks":    c.CompiledErrorWithStackTraces() != nil,
		"value":             c.Value() != nil,
	}

	// Assert
	expected := args.Map{
		"compiledErr":       true,
		"compiledJoiner":    true,
		"compiledJoinerAdd": true,
		"compiledStacks":    true,
		"value":             true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledError returns correct -- with errors", actual)
}

func Test_RawErrCollection_EmptyPaths(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{
		"stringEmpty":     c.String() == "",
		"compiledNil":     c.CompiledError() == nil,
		"compiledJoinNil": c.CompiledErrorUsingJoiner(",") == nil,
		"compiledAddNil":  c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") == nil,
		"compiledStkNil":  c.CompiledErrorWithStackTraces() == nil,
		"stkStrEmpty":     c.CompiledStackTracesString() == "",
		"valueNil":        c.Value() == nil,
		"stringsLen":      len(c.Strings()),
		"joinerEmpty":     c.StringUsingJoiner(",") == "",
		"addlEmpty":       c.StringUsingJoinerAdditional(",", "!") == "",
		"withAddlEmpty":   c.StringWithAdditionalMessage("!") == "",
		"fullTracesIf":    c.FullStringWithTracesIf(false) == "",
	}

	// Assert
	expected := args.Map{
		"stringEmpty":     true,
		"compiledNil":     true,
		"compiledJoinNil": true,
		"compiledAddNil":  true,
		"compiledStkNil":  true,
		"stkStrEmpty":     true,
		"valueNil":        true,
		"stringsLen":      0,
		"joinerEmpty":     true,
		"addlEmpty":       true,
		"withAddlEmpty":   true,
		"fullTracesIf":    true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection returns correct state -- empty collection paths", actual)
}

func Test_RawErrCollection_AddOps(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Adds(errors.New("a"), nil, errors.New("b"))
	c.AddErrors(errors.New("c"))
	c.AddIf(true, "d")
	c.AddIf(false, "skip")
	c.ConditionalAddError(true, errors.New("e"))
	c.ConditionalAddError(false, errors.New("skip"))
	c.AddStringSliceAsErr("f", "", "g")
	c.AddFunc(func() error { return errors.New("h") })
	c.AddFunc(nil)
	c.AddFuncIf(true, func() error { return errors.New("i") })
	c.AddFuncIf(false, func() error { return errors.New("skip") })
	c.AddFuncIf(true, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 9}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Add filters correctly -- mixed operations", actual)
}

func Test_RawErrCollection_AddWithRef_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.AddWithRef(errors.New("e"), "ref")
	c.AddWithRef(nil, "ref")
	c.AddWithCompiledTraceRef(errors.New("e"), "trace", "ref")
	c.AddWithCompiledTraceRef(nil, "trace", "ref")
	c.AddWithTraceRef(errors.New("e"), []string{"t"}, "ref")
	c.AddWithTraceRef(nil, nil, "ref")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddWithRef filters nil -- with ref", actual)
}

func Test_RawErrCollection_ClearDispose(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.Clear()

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Clear resets length -- with errors", actual)
}

func Test_RawErrCollection_Dispose_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.Dispose()

	// Act
	actual := args.Map{"isNull": c.IsNull()}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Dispose nullifies -- with errors", actual)
}

func Test_RawErrCollection_ClearEmpty(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Clear()   // should not panic on empty
	c.Dispose() // should not panic on empty

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Clear/Dispose completes safely -- empty collection", actual)
}

func Test_RawErrCollection_ToRawErrCollection_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Add(errors.New("e"))
	ptr := c.ToRawErrCollection()

	// Act
	actual := args.Map{"notNil": ptr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ToRawErrCollection returns pointer -- value receiver", actual)
}

func Test_RawErrCollection_Serialize_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	emptyBytes, emptyErr := c.Serialize()
	c.Add(errors.New("e"))
	bytes, err := c.Serialize()
	mustBytes := c.SerializeMust()

	// Act
	actual := args.Map{
		"emptyBytesNil": emptyBytes == nil,
		"emptyErrNil":   emptyErr == nil,
		"hasBytes":      len(bytes) > 0,
		"noErr":         err == nil,
		"mustHasBytes":  len(mustBytes) > 0,
	}

	// Assert
	expected := args.Map{
		"emptyBytesNil": true,
		"emptyErrNil":   true,
		"hasBytes":      true,
		"noErr":         true,
		"mustHasBytes":  true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Serialize returns nil -- empty collection", actual)
}

func Test_RawErrCollection_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	bytes, err := c.MarshalJSON()

	// Act
	actual := args.Map{
		"bytesNil": bytes == nil,
		"errNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"bytesNil": true,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.MarshalJSON returns nil -- empty collection", actual)
}

func Test_RawErrCollection_SerializeWithoutTraces_Empty(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	bytes, err := c.SerializeWithoutTraces()

	// Act
	actual := args.Map{
		"bytesNil": bytes == nil,
		"errNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"bytesNil": true,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.SerializeWithoutTraces returns nil -- empty collection", actual)
}

func Test_RawErrCollection_LogOps(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.Log()            // empty - should not log
	c.LogWithTraces()  // empty - should not log
	c.LogIf(false)     // false - should not log
	c.Add(errors.New("e"))
	c.Log()
	c.LogWithTraces()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Log completes safely -- with errors", actual)
}

func Test_RawErrCollection_HandleEmpty(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	c.HandleError()           // should not panic
	c.HandleErrorWithMsg("m") // should not panic
	c.HandleErrorWithRefs("m", "k", "v") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Handle completes safely -- empty collection", actual)
}

func Test_RawErrCollection_IsErrorsCollected(t *testing.T) {
	// Arrange
	c := &errcore.RawErrCollection{}
	changed := c.IsErrorsCollected(errors.New("e"))
	noChange := c.IsErrorsCollected(nil)

	// Act
	actual := args.Map{
		"changed": changed,
		"noChange": noChange,
	}

	// Assert
	expected := args.Map{
		"changed": true,
		"noChange": false,
	}
	expected.ShouldBeEqual(t, 0, "IsErrorsCollected returns correct value -- with args", actual)
}

func Test_RawErrCollection_CountStateChangeTracker_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	c := errcore.RawErrCollection{}
	c.Add(errors.New("e"))
	tracker := c.CountStateChangeTracker()

	// Act
	actual := args.Map{
		"isSameState": tracker.IsSameState(),
		"isValid":     tracker.IsValid(),
		"isSuccess":   tracker.IsSuccess(),
		"hasChanges":  tracker.HasChanges(),
		"isFailed":    tracker.IsFailed(),
		"isSameUsing": tracker.IsSameStateUsingCount(1),
	}

	// Assert
	expected := args.Map{
		"isSameState": true,
		"isValid":     true,
		"isSuccess":   true,
		"hasChanges":  false,
		"isFailed":    false,
		"isSameUsing": true,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker returns same -- no changes", actual)
}

// ── HandleErr (nil path) ──

func Test_HandleErr_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	errcore.HandleErr(nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr completes safely -- nil error", actual)
}

// ── MustBeEmpty (nil path) ──

func Test_MustBeEmpty_Nil_FromRawErrorTypeStringV2(t *testing.T) {
	// Arrange
	errcore.MustBeEmpty(nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty completes safely -- nil error", actual)
}
