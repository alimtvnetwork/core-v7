package errcoretests

import (
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── SliceToError / SliceToErrorPtr ──

func Test_SliceToError_Empty_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceToError(nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError nil -- nil", actual)
}

func Test_SliceToError_NonEmpty_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceToError([]string{"err1", "err2"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToError non-empty -- error", actual)
}

func Test_SliceToErrorPtr_Nil(t *testing.T) {
	// Arrange
	err := errcore.SliceToErrorPtr(nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr nil -- nil", actual)
}

func Test_SliceToErrorPtr_NonEmpty_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	s := []string{"e1"}
	err := errcore.SliceToErrorPtr(s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr non-empty -- error", actual)
}

// ── MergeErrors / MergeErrorsToString ──

func Test_MergeErrors_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.MergeErrors(nil, nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors both nil -- nil", actual)
}

func Test_MergeErrors_OneNil(t *testing.T) {
	// Arrange
	e := errors.New("e")

	// Act
	actual := args.Map{
		"first":  errcore.MergeErrors(e, nil) != nil,
		"second": errcore.MergeErrors(nil, e) != nil,
	}

	// Assert
	expected := args.Map{
		"first": true,
		"second": true,
	}
	expected.ShouldBeEqual(t, 0, "MergeErrors one nil -- non-nil", actual)
}

func Test_MergeErrors_Both(t *testing.T) {
	// Arrange
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors both -- merged", actual)
}

func Test_MergeErrorsToString_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MergeErrorsToString(",")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString no errors -- empty", actual)
}

func Test_MergeErrorsToString_NonNil_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString non-nil -- merged", actual)
}

func Test_MergeErrorsToStringDefault_Nil_FromSliceToErrorEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.MergeErrorsToStringDefault(nil, nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault nil -- empty", actual)
}

// ── ConcatMessageWithErr ──

func Test_ConcatMessageWithErr_NilErr_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErr("msg", nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr nil err -- nil", actual)
}

func Test_ConcatMessageWithErr_WithErr_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"containsPrefix": strings.Contains(err.Error(), "prefix"),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"containsPrefix": true,
	}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr with err -- prefixed", actual)
}

// ── ManyErrorToSingle ──

func Test_ManyErrorToSingle_Empty_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingle([]error{})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle empty -- nil", actual)
}

func Test_ManyErrorToSingle_AllNil(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingle([]error{nil, nil})

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle all nil -- nil", actual)
}

func Test_ManyErrorToSingle_WithErrors(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingle([]error{errors.New("a"), errors.New("b")})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle with errors -- joined", actual)
}

// ── ToString / ToStringPtr ──

func Test_ToString_Nil_FromSliceToErrorEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ToString(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToString nil -- empty", actual)
}

func Test_ToString_NonNil_FromSliceToErrorEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.ToString(errors.New("err"))}

	// Assert
	expected := args.Map{"result": "err"}
	expected.ShouldBeEqual(t, 0, "ToString non-nil -- err", actual)
}

func Test_ToStringPtr_Nil_FromSliceToErrorEmpty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ToStringPtr(nil) == nil}

	// Assert
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "ToStringPtr nil -- returns pointer to empty string", actual)
}

func Test_ToStringPtr_NonNil_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(errors.New("err"))

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"val": *result,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": "err",
	}
	expected.ShouldBeEqual(t, 0, "ToStringPtr non-nil -- err", actual)
}

// ── ToError ──

func Test_ToError_Empty_FromSliceToErrorEmpty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ToError("") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError empty -- nil", actual)
}

func Test_ToError_NonEmpty_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	err := errcore.ToError("msg")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"msg": err.Error(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"msg": "msg",
	}
	expected.ShouldBeEqual(t, 0, "ToError non-empty -- msg", actual)
}

// ── LineDiff ──

func Test_LineDiff_Same(t *testing.T) {
	// Arrange
	result := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})

	// Act
	actual := args.Map{"noResults": len(result) == 0}

	// Assert
	expected := args.Map{"noResults": false}
	expected.ShouldBeEqual(t, 0, "LineDiff same -- returns all lines including matches", actual)
}

func Test_LineDiff_Different(t *testing.T) {
	// Arrange
	result := errcore.LineDiff([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"hasResults": len(result) > 0}

	// Assert
	expected := args.Map{"hasResults": true}
	expected.ShouldBeEqual(t, 0, "LineDiff different -- has results", actual)
}

// ── GherkinsString ──

func Test_GherkinsString_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	result := errcore.GherkinsString(0, "feature", "given", "when", "then")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString -- formatted", actual)
}

// ── StringLinesToQuoteLines ──

func Test_StringLinesToQuoteLines_Empty_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines(nil)

	// Act
	actual := args.Map{"isEmpty": len(result) == 0}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines nil -- empty", actual)
}

func Test_StringLinesToQuoteLines_NonEmpty_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines non-empty -- formatted", actual)
}

// ── MessageWithRef / ErrorWithRef ──

func Test_MessageWithRef_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	result := errcore.MessageWithRef("msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef -- formatted", actual)
}

func Test_MessageWithRefToError_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	err := errcore.MessageWithRefToError("msg", "ref")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError -- error", actual)
}

// ── HandleErr ──

func Test_HandleErr_Nil_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	// Should not panic
	errcore.HandleErr(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErr nil -- no panic", actual)
}

// ── MustBeEmpty ──

func Test_MustBeEmpty_Nil_FromSliceToErrorEmpty(t *testing.T) {
	// Arrange
	// Should not panic
	errcore.MustBeEmpty(nil)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty nil -- no panic", actual)
}
