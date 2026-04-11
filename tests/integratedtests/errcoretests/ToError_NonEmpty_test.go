package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── ToError ──

func Test_ToError_NonEmpty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ToError("some error")

	// Act
	actual := args.Map{
		"notNil": err != nil,
		"msg": err.Error(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"msg": "some error",
	}
	expected.ShouldBeEqual(t, 0, "ToError returns error -- non-empty string", actual)
}

func Test_ToError_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ToError("")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError returns nil -- empty string", actual)
}

// ── ToString ──

func Test_ToString_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ToString(nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToString returns empty -- nil error", actual)
}

func Test_ToString_WithError(t *testing.T) {
	// Arrange
	result := errcore.ToString(errors.New("hello"))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ToString returns msg -- with error", actual)
}

// ── ToStringPtr ──

func Test_ToStringPtr_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(nil)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"val": *result,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": "",
	}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns empty ptr -- nil error", actual)
}

func Test_ToStringPtr_WithError(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(errors.New("err msg"))

	// Act
	actual := args.Map{"val": *result}

	// Assert
	expected := args.Map{"val": "err msg"}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns ptr -- with error", actual)
}

// ── ToValueString ──

func Test_ToValueString_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ToValueString("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToValueString returns non-empty -- with value", actual)
}

// ── Ref ──

func Test_Ref_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.Ref(nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Ref returns empty -- nil input", actual)
}

func Test_Ref_WithValue(t *testing.T) {
	// Arrange
	result := errcore.Ref("some-ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref returns non-empty -- with value", actual)
}

// ── RefToError ──

func Test_RefToError_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.RefToError(nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError returns nil -- nil input", actual)
}

func Test_RefToError_WithValue(t *testing.T) {
	// Arrange
	err := errcore.RefToError("ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError returns error -- with value", actual)
}

// ── SliceError ──

func Test_SliceError_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceError(",", []string{})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns nil -- empty slice", actual)
}

func Test_SliceError_NonEmpty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceError(",", []string{"a", "b"})

	// Act
	actual := args.Map{"msg": err.Error()}

	// Assert
	expected := args.Map{"msg": "a,b"}
	expected.ShouldBeEqual(t, 0, "SliceError returns error -- non-empty slice", actual)
}

// ── SliceErrorDefault ──

func Test_SliceErrorDefault_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceErrorDefault([]string{})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault returns nil -- empty slice", actual)
}

// ── SliceToError ──

func Test_SliceToError_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceToError([]string{})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns nil -- empty slice", actual)
}

func Test_SliceToError_NonEmpty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceToError([]string{"x", "y"})

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns error -- non-empty slice", actual)
}

// ── SliceToErrorPtr ──

func Test_SliceToErrorPtr_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceToErrorPtr([]string{})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns nil -- empty slice", actual)
}

// ── SliceErrorsToStrings ──

func Test_SliceErrorsToStrings_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.SliceErrorsToStrings(nil...)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns empty -- no errors", actual)
}

func Test_SliceErrorsToStrings_WithNils_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b"))

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns filtered -- with nils", actual)
}

// ── MergeErrors ──

func Test_MergeErrors_AllNil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.MergeErrors(nil, nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns nil -- all nil", actual)
}

func Test_MergeErrors_Mixed(t *testing.T) {
	// Arrange
	err := errcore.MergeErrors(errors.New("e1"), nil, errors.New("e2"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns error -- mixed nil and non-nil", actual)
}

// ── ManyErrorToSingle ──

func Test_ManyErrorToSingle_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingle([]error{})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle returns nil -- empty slice", actual)
}

// ── ManyErrorToSingleDirect ──

func Test_ManyErrorToSingleDirect_Empty(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingleDirect()

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect returns nil -- no args", actual)
}

// ── MergeErrorsToString ──

func Test_MergeErrorsToString_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToString(",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns empty -- no errors", actual)
}

func Test_MergeErrorsToString_WithErrors_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns non-empty -- with errors", actual)
}

// ── MergeErrorsToStringDefault ──

func Test_MergeErrorsToStringDefault_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToStringDefault()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault returns empty -- no errors", actual)
}

// ── ErrorToSplitLines ──

func Test_ErrorToSplitLines_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitLines(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines returns empty -- nil error", actual)
}

func Test_ErrorToSplitLines_WithLines(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitLines(errors.New("line1\nline2"))

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "line1",
		"second": "line2",
	}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines returns lines -- with lines", actual)
}

// ── ErrorToSplitNonEmptyLines ──

func Test_ErrorToSplitNonEmptyLines_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitNonEmptyLines(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines returns empty -- nil error", actual)
}

// ── VarTwo ──

func Test_VarTwo_NoType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarTwo(false, "a", 1, "b", 2)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(a, b) = (1, 2)"}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- no type", actual)
}

func Test_VarTwo_WithType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarTwo(true, "a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- with type", actual)
}

// ── VarTwoNoType ──

func Test_VarTwoNoType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarTwoNoType("x", 10, "y", 20)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(x, y) = (10, 20)"}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType returns formatted -- with args", actual)
}

// ── VarThree ──

func Test_VarThree_NoType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(a, b, c) = (1, 2, 3)"}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- no type", actual)
}

func Test_VarThree_WithType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- with type", actual)
}

// ── VarThreeNoType ──

func Test_VarThreeNoType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarThreeNoType("x", 1, "y", 2, "z", 3)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(x, y, z) = (1, 2, 3)"}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType returns formatted -- with args", actual)
}

// ── MessageVarTwo ──

func Test_MessageVarTwo_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo returns formatted -- with args", actual)
}

// ── MessageVarThree ──

func Test_MessageVarThree_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree returns formatted -- with args", actual)
}

// ── MessageWithRef ──

func Test_MessageWithRef_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageWithRef("msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef returns non-empty -- with args", actual)
}

// ── MessageWithRefToError ──

func Test_MessageWithRefToError_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.MessageWithRefToError("msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError returns error -- with args", actual)
}

// ── ErrorWithRef ──

func Test_ErrorWithRef_NilErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithRef(nil, "ref")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns empty -- nil error", actual)
}

func Test_ErrorWithRef_NilRef_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithRef(errors.New("e"), nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "e"}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns error msg -- nil ref", actual)
}

func Test_ErrorWithRef_EmptyRef_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithRef(errors.New("e"), "")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "e"}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns error msg -- empty ref", actual)
}

func Test_ErrorWithRef_Both(t *testing.T) {
	// Arrange
	result := errcore.ErrorWithRef(errors.New("e"), "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns formatted -- with error and ref", actual)
}

// ── Combine ──

func Test_Combine_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.Combine("generic", "other", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine returns formatted -- with args", actual)
}

// ── ConcatMessageWithErr ──

func Test_ConcatMessageWithErr_NilErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErr("msg", nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr returns nil -- nil error", actual)
}

func Test_ConcatMessageWithErr_WithErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr returns error -- with error", actual)
}

// ── MustBeEmpty ──

func Test_MustBeEmpty_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.MustBeEmpty(nil)
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty nil -- no panic", actual)
}

func Test_MustBeEmpty_WithError(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.MustBeEmpty(errors.New("e"))
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty with error -- panic", actual)
}

// ── VarMap ──

func Test_VarMap_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarMap(map[string]any{})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarMap returns empty -- nil map", actual)
}

func Test_VarMap_NonEmpty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarMap(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap returns formatted -- with entries", actual)
}

// ── VarMapStrings ──

func Test_VarMapStrings_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarMapStrings(map[string]any{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarMapStrings returns empty -- nil map", actual)
}

// ── MessageVarMap ──

func Test_MessageVarMap_Empty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageVarMap("msg", map[string]any{})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns msg only -- empty map", actual)
}

func Test_MessageVarMap_NonEmpty_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageVarMap("msg", map[string]any{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns formatted -- with map", actual)
}

// ── Expecting ──

func Test_Expecting_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.Expecting("title", "expected", "actual")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting returns formatted -- with args", actual)
}

// ── ExpectingSimple ──

func Test_ExpectingSimple_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ExpectingSimple("title", "expected", "actual")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple returns formatted -- with args", actual)
}

// ── ExpectingSimpleNoType ──

func Test_ExpectingSimpleNoType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ExpectingSimpleNoType("title", "expected", "actual")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType returns formatted -- with args", actual)
}

// ── ExpectingErrorSimpleNoType ──

func Test_ExpectingErrorSimpleNoType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ExpectingErrorSimpleNoType("title", "expected", "actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- with args", actual)
}

// ── ExpectingErrorSimpleNoTypeNewLineEnds ──

func Test_ExpectingErrorSimpleNoTypeNewLineEnds_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ExpectingErrorSimpleNoTypeNewLineEnds("title", "expected", "actual")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoTypeNewLineEnds returns error -- with args", actual)
}

// ── ShouldBe ──

func Test_ShouldBe_StrEqMsg_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.StrEqMsg("actual", "expected")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg returns non-empty -- different strings", actual)
}

func Test_ShouldBe_StrEqErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.StrEqErr("actual", "expected")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr returns error -- different strings", actual)
}

func Test_ShouldBe_AnyEqMsg_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.AnyEqMsg(1, 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg returns non-empty -- different values", actual)
}

func Test_ShouldBe_AnyEqErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.AnyEqErr(1, 2)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr returns error -- different values", actual)
}

func Test_ShouldBe_JsonEqMsg_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.ShouldBe.JsonEqMsg("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg returns non-empty -- different json", actual)
}

func Test_ShouldBe_JsonEqErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ShouldBe.JsonEqErr("a", "b")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr returns error -- different json", actual)
}

// ── Expected ──

func Test_Expected_But_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.Expected.But("title", "exp", "act")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.But returns error -- with args", actual)
}

func Test_Expected_ButFoundAsMsg_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.Expected.ButFoundAsMsg("title", "exp", "act")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButFoundWithTypeAsMsg_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.Expected.ButFoundWithTypeAsMsg("title", "exp", "act")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg returns non-empty -- with args", actual)
}

func Test_Expected_ButUsingType_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.Expected.ButUsingType("title", "exp", "act")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType returns error -- with args", actual)
}

// ── RawErrorType ──

func Test_RawErrorType_String_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.String returns non-empty -- with type", actual)
}

func Test_RawErrorType_Combine_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.Combine("other", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Combine returns formatted -- with msg and ref", actual)
}

func Test_RawErrorType_CombineWithAnother_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.CombineWithAnother(errcore.NotFound, "other", "ref")

	// Act
	actual := args.Map{"notEmpty": string(result) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.CombineWithAnother returns non-empty -- with another type", actual)
}

func Test_RawErrorType_MergeError_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeError(nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeError_WithErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeError(errors.New("inner"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithMessage(nil, "msg")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_WithErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithMessage(errors.New("inner"), "msg")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithRef_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithRef(nil, "ref")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithRef_WithErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithRef(errors.New("inner"), "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef returns error -- with error", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_Nil_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(nil, "msg", "ref")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef returns nil -- nil error", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_WithErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(errors.New("inner"), "msg", "ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef returns error -- with error", actual)
}

func Test_RawErrorType_FmtIf_False_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.FmtIf(false, "x=%d", 1)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf returns nil -- condition false", actual)
}

func Test_RawErrorType_FmtIf_True_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.FmtIf(true, "x=%d", 1)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf returns error -- condition true", actual)
}

func Test_RawErrorType_SrcDestination_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.SrcDestination("msg", "src", "sv", "dst", "dv")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestination returns formatted -- with args", actual)
}

func Test_RawErrorType_SrcDestinationErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.SrcDestinationErr("msg", "src", "sv", "dst", "dv")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestinationErr returns error -- with args", actual)
}

func Test_RawErrorType_TypesAttach_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.InvalidRequestType.TypesAttach("msg", "type1")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttach returns non-empty -- with types", actual)
}

func Test_RawErrorType_TypesAttachErr_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.InvalidRequestType.TypesAttachErr("msg", "type1")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttachErr returns error -- with types", actual)
}

func Test_GetSet_True_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.GetSet(true, errcore.InvalidRequestType, errcore.NotFound)

	// Act
	actual := args.Map{"result": string(result)}

	// Assert
	expected := args.Map{"result": string(errcore.InvalidRequestType)}
	expected.ShouldBeEqual(t, 0, "GetSet returns trueValue -- condition true", actual)
}

func Test_GetSet_False_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.GetSet(false, errcore.InvalidRequestType, errcore.NotFound)

	// Act
	actual := args.Map{"result": string(result)}

	// Assert
	expected := args.Map{"result": string(errcore.NotFound)}
	expected.ShouldBeEqual(t, 0, "GetSet returns falseValue -- condition false", actual)
}

func Test_GetSetVariant_True_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.GetSetVariant(true, "yes", "no")

	// Act
	actual := args.Map{"result": string(result)}

	// Assert
	expected := args.Map{"result": "yes"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns trueValue -- condition true", actual)
}

func Test_GetSetVariant_False_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	result := errcore.GetSetVariant(false, "yes", "no")

	// Act
	actual := args.Map{"result": string(result)}

	// Assert
	expected := args.Map{"result": "no"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns falseValue -- condition false", actual)
}

// ── CombineWithMsgTypeNoStack ──

func Test_CombineWithMsgTypeNoStack_EmptyOther(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- empty otherMsg", actual)
}

func Test_CombineWithMsgTypeNoStack_WithOther(t *testing.T) {
	// Arrange
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "other", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- with otherMsg", actual)
}

// ── WasExpectingErrorF ──

func Test_WasExpectingErrorF_FromToErrorNonEmpty(t *testing.T) {
	// Arrange
	err := errcore.WasExpectingErrorF("exp", "act", "title: %d", 1)

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "WasExpectingErrorF returns error -- with format", actual)
}
