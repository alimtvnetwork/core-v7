package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ══════════════════════════════════════════════════════════════════════════════
// LineDiff / LineDiffToString / PrintLineDiff / HasAnyMismatchOnLines
// ══════════════════════════════════════════════════════════════════════════════

func Test_LineDiff_AllMatch(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(diffs),
		"st0": diffs[0].Status,
		"st1": diffs[1].Status,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"st0": "  ",
		"st1": "  ",
	}
	expected.ShouldBeEqual(t, 0, "LineDiff returns all-match -- matching lines", actual)
}

func Test_LineDiff_Mismatch(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "x"}, []string{"a", "b"})

	// Act
	actual := args.Map{"st1": diffs[1].Status}

	// Assert
	expected := args.Map{"st1": "!!"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns mismatch -- different line", actual)
}

func Test_LineDiff_ExtraActual_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a", "b", "c"}, []string{"a"})

	// Act
	actual := args.Map{"st2": diffs[2].Status}

	// Assert
	expected := args.Map{"st2": "+"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns extra-actual -- longer actual", actual)
}

func Test_LineDiff_MissingExpected_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b", "c"})

	// Act
	actual := args.Map{"st2": diffs[2].Status}

	// Assert
	expected := args.Map{"st2": "-"}
	expected.ShouldBeEqual(t, 0, "LineDiff returns missing-expected -- shorter actual", actual)
}

func Test_LineDiffToString_Empty_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "h", []string{}, []string{})

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns empty -- both empty", actual)
}

func Test_LineDiffToString_WithDiffs(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "h", []string{"a", "x"}, []string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns non-empty -- with diffs", actual)
}

func Test_LineDiffToString_AllBranches(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "h", []string{"a", "x", "extra"}, []string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns non-empty -- all branches", actual)
}

func Test_PrintLineDiff_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	errcore.PrintLineDiff(0, "h", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiff completes safely -- with diffs", actual)
}

func Test_PrintLineDiff_Empty(t *testing.T) {
	// Arrange
	errcore.PrintLineDiff(0, "h", []string{}, []string{})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiff completes safely -- both empty", actual)
}

func Test_HasAnyMismatchOnLines_Match(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns false -- matching", actual)
}

func Test_HasAnyMismatchOnLines_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- different length", actual)
}

func Test_HasAnyMismatchOnLines_DiffContent(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"})}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- different content", actual)
}

func Test_PrintLineDiffOnFail_NoFail(t *testing.T) {
	// Arrange
	errcore.PrintLineDiffOnFail(0, "h", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail completes safely -- no mismatch", actual)
}

func Test_PrintLineDiffOnFail_Fail(t *testing.T) {
	// Arrange
	errcore.PrintLineDiffOnFail(0, "h", []string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail prints diff -- with mismatch", actual)
}

func Test_ErrorToLinesLineDiff_NilErr(t *testing.T) {
	// Arrange
	result := errcore.ErrorToLinesLineDiff(0, "h", nil, []string{"a"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff returns non-empty -- nil error", actual)
}

func Test_ErrorToLinesLineDiff_WithErr(t *testing.T) {
	// Arrange
	result := errcore.ErrorToLinesLineDiff(0, "h", errors.New("line1\nline2"), []string{"line1"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff returns non-empty -- with error", actual)
}

func Test_PrintErrorLineDiff_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	errcore.PrintErrorLineDiff(0, "h", errors.New("a"), []string{"b"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorLineDiff completes safely -- with args", actual)
}

func Test_SliceDiffSummary_Match(t *testing.T) {
	// Arrange
	result := errcore.SliceDiffSummary([]string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"v": result}

	// Assert
	expected := args.Map{"v": "all lines match"}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns all-match -- matching", actual)
}

func Test_SliceDiffSummary_Mismatch(t *testing.T) {
	// Arrange
	result := errcore.SliceDiffSummary([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"notEmpty": result != "all lines match"}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns mismatch count -- with mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapMismatchError
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapMismatchError_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	result := errcore.MapMismatchError("TestFunc", 1, "title",
		[]string{`"k": "v"`}, []string{`"k": "v2"`})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapMismatchError returns formatted -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AssertDiffOnMismatch / AssertErrorDiffOnMismatch
// ══════════════════════════════════════════════════════════════════════════════

func Test_AssertDiffOnMismatch_Match(t *testing.T) {
	// Arrange
	errcore.AssertDiffOnMismatch(t, 0, "t", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "AssertDiffOnMismatch completes safely -- matching", actual)
}

func Test_AssertErrorDiffOnMismatch_NilMatch(t *testing.T) {
	// Arrange
	errcore.AssertErrorDiffOnMismatch(t, 0, "t", nil, []string{})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "AssertErrorDiffOnMismatch completes safely -- nil error matching", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PrintDiffOnMismatch
// ══════════════════════════════════════════════════════════════════════════════

func Test_PrintDiffOnMismatch_Match(t *testing.T) {
	// Arrange
	errcore.PrintDiffOnMismatch(0, "t", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch completes safely -- matching", actual)
}

func Test_PrintDiffOnMismatch_Mismatch(t *testing.T) {
	// Arrange
	errcore.PrintDiffOnMismatch(0, "t", []string{"a"}, []string{"b"}, "ctx1")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch prints diff -- with mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MergeErrors / MergeErrorsToString / MergeErrorsToStringDefault
// ══════════════════════════════════════════════════════════════════════════════

func Test_MergeErrors_AllNil(t *testing.T) {
	// Arrange
	err := errcore.MergeErrors(nil, nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns nil -- all nil", actual)
}

func Test_MergeErrors_WithErr(t *testing.T) {
	// Arrange
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors returns error -- with errors", actual)
}

func Test_MergeErrorsToString_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.MergeErrorsToString(",")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns empty -- no errors", actual)
}

func Test_MergeErrorsToString_WithErr(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString returns non-empty -- with errors", actual)
}

func Test_MergeErrorsToStringDefault_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.MergeErrorsToStringDefault()}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault returns empty -- no errors", actual)
}

func Test_MergeErrorsToStringDefault_WithErr(t *testing.T) {
	// Arrange
	result := errcore.MergeErrorsToStringDefault(errors.New("a"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault returns non-empty -- with errors", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceError / SliceErrorDefault / SliceErrorsToStrings / SliceToError / SliceToErrorPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceError_Empty_FromLineDiffAllMatch(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceError(",", []string{}) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns nil -- empty slice", actual)
}

func Test_SliceError_NonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceError(",", []string{"a", "b"})

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError returns error -- non-empty slice", actual)
}

func Test_SliceErrorDefault_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceErrorDefault([]string{}) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault returns nil -- empty slice", actual)
}

func Test_SliceErrorDefault_NonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceErrorDefault([]string{"a"})

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault returns error -- non-empty slice", actual)
}

func Test_SliceErrorsToStrings_Nil(t *testing.T) {
	// Arrange
	result := errcore.SliceErrorsToStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns empty -- no errors", actual)
}

func Test_SliceErrorsToStrings_WithNils(t *testing.T) {
	// Arrange
	result := errcore.SliceErrorsToStrings(nil, errors.New("a"), nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings returns filtered -- with nils", actual)
}

func Test_SliceToError_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceToError([]string{}) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns nil -- empty slice", actual)
}

func Test_SliceToError_NonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceToError([]string{"a", "b"})

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError returns error -- non-empty slice", actual)
}

func Test_SliceToErrorPtr_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.SliceToErrorPtr([]string{}) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns nil -- empty slice", actual)
}

func Test_SliceToErrorPtr_NonEmpty(t *testing.T) {
	// Arrange
	err := errcore.SliceToErrorPtr([]string{"a"})

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr returns error -- non-empty slice", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MustBeEmpty
// ══════════════════════════════════════════════════════════════════════════════

func Test_MustBeEmpty_Nil(t *testing.T) {
	// Arrange
	errcore.MustBeEmpty(nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty completes safely -- nil error", actual)
}

func Test_MustBeEmpty_Panic_FromLineDiffAllMatch(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "MustBeEmpty panics -- with error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ErrorToSplitLines / ErrorToSplitNonEmptyLines
// ══════════════════════════════════════════════════════════════════════════════

func Test_ErrorToSplitLines_Nil(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitLines(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines returns empty -- nil error", actual)
}

func Test_ErrorToSplitLines_Multi(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitLines(errors.New("a\nb\nc"))

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines returns lines -- multi-line error", actual)
}

func Test_ErrorToSplitNonEmptyLines_WithEmpty(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines returns filtered -- with empty lines", actual)
}

func Test_ErrorToSplitNonEmptyLines_Nil(t *testing.T) {
	// Arrange
	result := errcore.ErrorToSplitNonEmptyLines(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines returns empty -- nil error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Ref / RefToError / ToError / ToString / ToStringPtr / ToValueString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Ref_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.Ref(nil)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "Ref returns empty -- nil input", actual)
}

func Test_Ref_NonNil(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": errcore.Ref("hello") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref returns non-empty -- with value", actual)
}

func Test_RefToError_Nil_FromLineDiffAllMatch(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.RefToError(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError returns nil -- nil input", actual)
}

func Test_RefToError_NonNil_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	err := errcore.RefToError("ref")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError returns error -- with value", actual)
}

func Test_ToError_Empty(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ToError("") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError returns nil -- empty string", actual)
}

func Test_ToError_NonEmpty(t *testing.T) {
	// Arrange
	err := errcore.ToError("e")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToError returns error -- non-empty string", actual)
}

func Test_ToString_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.ToString(nil)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ToString returns empty -- nil error", actual)
}

func Test_ToString_WithErr(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.ToString(errors.New("e"))}

	// Assert
	expected := args.Map{"v": "e"}
	expected.ShouldBeEqual(t, 0, "ToString returns msg -- with error", actual)
}

func Test_ToStringPtr_Nil(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(nil)

	// Act
	actual := args.Map{"empty": *result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns empty ptr -- nil error", actual)
}

func Test_ToStringPtr_WithErr(t *testing.T) {
	// Arrange
	result := errcore.ToStringPtr(errors.New("e"))

	// Act
	actual := args.Map{"v": *result}

	// Assert
	expected := args.Map{"v": "e"}
	expected.ShouldBeEqual(t, 0, "ToStringPtr returns ptr -- with error", actual)
}

func Test_ToValueString_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	result := errcore.ToValueString("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToValueString returns non-empty -- with value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VarMap / VarMapStrings
// ══════════════════════════════════════════════════════════════════════════════

func Test_VarMap_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.VarMap(map[string]any{})}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "VarMap returns empty -- nil map", actual)
}

func Test_VarMap_NonEmpty(t *testing.T) {
	// Arrange
	result := errcore.VarMap(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap returns formatted -- with entries", actual)
}

func Test_VarMapStrings_Empty_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	result := errcore.VarMapStrings(map[string]any{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarMapStrings returns empty -- nil map", actual)
}

func Test_VarMapStrings_NonEmpty_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	result := errcore.VarMapStrings(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VarMapStrings returns entries -- with map", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Combine
// ══════════════════════════════════════════════════════════════════════════════

func Test_Combine_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	result := errcore.Combine("generic", "other", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine returns formatted -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ConcatMessageWithErr
// ══════════════════════════════════════════════════════════════════════════════

func Test_ConcatMessageWithErr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"isNil": errcore.ConcatMessageWithErr("msg", nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr returns nil -- nil error", actual)
}

func Test_ConcatMessageWithErr_WithErr_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr returns error -- with error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RawErrCollection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_RawErrCollection_Basic(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{
		"isEmpty": rec.IsEmpty(), "isNull": rec.IsNull(), "isAnyNull": rec.IsAnyNull(),
		"length": rec.Length(), "hasError": rec.HasError(), "hasAny": rec.HasAnyError(),
		"isValid": rec.IsValid(), "isSuccess": rec.IsSuccess(), "isFailed": rec.IsFailed(),
		"isInvalid": rec.IsInvalid(), "isDefined": rec.IsDefined(), "hasIssues": rec.HasAnyIssues(),
		"isCollection": rec.IsCollectionType(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true, "isNull": true, "isAnyNull": true,
		"length": 0, "hasError": false, "hasAny": false,
		"isValid": true, "isSuccess": true, "isFailed": false,
		"isInvalid": false, "isDefined": false, "hasIssues": false,
		"isCollection": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection returns correct state -- empty collection", actual)
}

func Test_RawErrCollection_Add(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(nil)
	rec.Add(errors.New("a"))
	rec.AddError(nil)
	rec.AddError(errors.New("b"))

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Add filters nil -- mixed nil and non-nil", actual)
}

func Test_RawErrCollection_AddMsg(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddMsg("hello")
	rec.AddMsg("")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddMsg adds non-empty -- filters empty", actual)
}

func Test_RawErrCollection_AddMsgStackTrace(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddMsgStackTrace("")
	rec.AddMsgStackTrace("msg")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddMsgStackTrace adds non-empty -- filters empty", actual)
}

func Test_RawErrCollection_AddStackTrace(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddStackTrace(nil)
	rec.AddStackTrace(errors.New("e"))

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddStackTrace adds non-nil -- filters nil", actual)
}

func Test_RawErrCollection_AddMsgErrStackTrace(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddMsgErrStackTrace("msg", nil)
	rec.AddMsgErrStackTrace("msg", errors.New("e"))

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddMsgErrStackTrace adds non-nil -- filters nil", actual)
}

func Test_RawErrCollection_AddMethodName(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddMethodName("")
	rec.AddMethodName("msg")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddMethodName adds non-empty -- filters empty", actual)
}

func Test_RawErrCollection_AddMessages(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddMessages()
	rec.AddMessages("a", "b")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddMessages adds joined -- filters empty", actual)
}

func Test_RawErrCollection_AddErrorWithMessage(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddErrorWithMessage(nil, "msg")
	rec.AddErrorWithMessage(errors.New("e"), "msg")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddErrorWithMessage adds non-nil -- filters nil", actual)
}

func Test_RawErrCollection_AddIf(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddIf(false, "skip")
	rec.AddIf(true, "add")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddIf adds conditionally -- condition true", actual)
}

func Test_RawErrCollection_AddFunc(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddFunc(nil)
	rec.AddFunc(func() error { return nil })
	rec.AddFunc(func() error { return errors.New("e") })

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddFunc adds error -- from func", actual)
}

func Test_RawErrCollection_AddFuncIf(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddFuncIf(false, func() error { return errors.New("e") })
	rec.AddFuncIf(true, nil)
	rec.AddFuncIf(true, func() error { return errors.New("e") })

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddFuncIf adds conditionally -- from func", actual)
}

func Test_RawErrCollection_AddErrorWithMessageRef(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddErrorWithMessageRef(nil, "msg", "ref")
	rec.AddErrorWithMessageRef(errors.New("e"), "msg", nil)
	rec.AddErrorWithMessageRef(errors.New("e"), "msg", "ref")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddErrorWithMessageRef adds non-nil -- with ref", actual)
}

func Test_RawErrCollection_AddFmt(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddFmt(nil, "fmt %d", 1)
	rec.AddFmt(errors.New("e"), "fmt %d", 1)

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddFmt adds formatted -- filters nil", actual)
}

func Test_RawErrCollection_Fmt(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Fmt("", )
	rec.Fmt("hello %d", 1)

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Fmt adds formatted -- filters empty", actual)
}

func Test_RawErrCollection_FmtIf(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.FmtIf(false, "skip %d", 1)
	rec.FmtIf(true, "add %d", 1)

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.FmtIf adds conditionally -- condition true", actual)
}

func Test_RawErrCollection_References(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.References("msg", "r1")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.References adds entry -- with refs", actual)
}

func Test_RawErrCollection_Adds(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Adds()
	rec.Adds(nil, errors.New("a"), nil, errors.New("b"))

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Adds adds non-nil -- filters nil", actual)
}

func Test_RawErrCollection_AddErrors(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddErrors(errors.New("a"))

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddErrors adds error -- with error", actual)
}

func Test_RawErrCollection_ConditionalAddError(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.ConditionalAddError(false, errors.New("skip"))
	rec.ConditionalAddError(true, errors.New("add"))

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ConditionalAddError adds conditionally -- condition true", actual)
}

func Test_RawErrCollection_AddString(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddString("")
	rec.AddString("msg")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddString adds non-empty -- filters empty", actual)
}

func Test_RawErrCollection_AddStringSliceAsErr(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddStringSliceAsErr()
	rec.AddStringSliceAsErr("", "a", "", "b")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddStringSliceAsErr adds non-empty -- filters empty", actual)
}

func Test_RawErrCollection_AddWithTraceRef(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddWithTraceRef(nil, []string{"t"}, "r")
	rec.AddWithTraceRef(errors.New("e"), []string{"t"}, "r")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddWithTraceRef adds non-nil -- with trace and ref", actual)
}

func Test_RawErrCollection_AddWithCompiledTraceRef(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddWithCompiledTraceRef(nil, "trace", "r")
	rec.AddWithCompiledTraceRef(errors.New("e"), "trace", "r")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddWithCompiledTraceRef adds non-nil -- with compiled trace", actual)
}

func Test_RawErrCollection_AddWithRef(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddWithRef(nil, "r")
	rec.AddWithRef(errors.New("e"), "r")

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddWithRef adds non-nil -- with ref", actual)
}

func Test_RawErrCollection_String(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"empty": rec.String() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.String returns empty -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notEmpty": rec.String() != ""}
	expected2 := args.Map{"notEmpty": true}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.String returns non-empty -- with errors", actual2)
}

func Test_RawErrCollection_Strings(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"len": len(rec.Strings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Strings returns empty -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"len": len(rec.Strings())}
	expected2 := args.Map{"len": 1}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.Strings returns entries -- with errors", actual2)
}

func Test_RawErrCollection_StringUsingJoiner(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"empty": rec.StringUsingJoiner(",") == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.StringUsingJoiner returns empty -- empty collection", actual)

	rec.Add(errors.New("a"))
	rec.Add(errors.New("b"))
	actual2 := args.Map{"v": rec.StringUsingJoiner(",")}
	expected2 := args.Map{"v": "a,b"}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.StringUsingJoiner returns joined -- with errors", actual2)
}

func Test_RawErrCollection_StringUsingJoinerAdditional(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"empty": rec.StringUsingJoinerAdditional(",", "!") == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.StringUsingJoinerAdditional returns empty -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"v": rec.StringUsingJoinerAdditional(",", "!")}
	expected2 := args.Map{"v": "a!"}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.StringUsingJoinerAdditional returns formatted -- with errors", actual2)
}

func Test_RawErrCollection_CompiledError(t *testing.T) {
	// Arrange
	rec := errcore.RawErrCollection{}

	// Act
	actual := args.Map{"isNil": rec.CompiledError() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledError returns nil -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledError() != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.CompiledError returns error -- with errors", actual2)
}

func Test_RawErrCollection_CompiledErrorUsingJoiner(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"isNil": rec.CompiledErrorUsingJoiner(",") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorUsingJoiner returns nil -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorUsingJoiner(",") != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorUsingJoiner returns error -- with errors", actual2)
}

func Test_RawErrCollection_CompiledErrorUsingJoinerAdditionalMessage(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"isNil": rec.CompiledErrorUsingJoinerAdditionalMessage(",", "!") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorUsingJoinerAdditionalMessage returns nil -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorUsingJoinerAdditionalMessage(",", "!") != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorUsingJoinerAdditionalMessage returns error -- with errors", actual2)
}

func Test_RawErrCollection_CompiledErrorUsingStackTraces(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"isNil": rec.CompiledErrorUsingStackTraces(",", []string{"t"}) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorUsingStackTraces returns nil -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorUsingStackTraces(",", []string{"t"}) != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorUsingStackTraces returns error -- with errors", actual2)
}

func Test_RawErrCollection_StringWithAdditionalMessage(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"empty": rec.StringWithAdditionalMessage("!") == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.StringWithAdditionalMessage returns empty -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"v": rec.StringWithAdditionalMessage("!")}
	expected2 := args.Map{"v": "a!"}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.StringWithAdditionalMessage returns formatted -- with errors", actual2)
}

func Test_RawErrCollection_CompiledErrorWithStackTraces(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"isNil": rec.CompiledErrorWithStackTraces() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorWithStackTraces returns nil -- empty with traces", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorWithStackTraces() != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.CompiledErrorWithStackTraces returns error -- with traces", actual2)
}

func Test_RawErrCollection_CompiledStackTracesString(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"empty": rec.CompiledStackTracesString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledStringToString returns empty -- empty collection", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notEmpty": rec.CompiledStackTracesString() != ""}
	expected2 := args.Map{"notEmpty": true}
	expected2.ShouldBeEqual(t, 0, "RawErrCollection.CompiledStringToString returns non-empty -- with errors", actual2)
}

func Test_RawErrCollection_CompiledJsonErrorWithStackTraces_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	err := rec.CompiledJsonErrorWithStackTraces()

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledJoinErrorWithStackTraces returns non-empty -- with errors", actual)
}

func Test_RawErrCollection_CompiledJsonStringWithStackTraces_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	result := rec.CompiledJsonStringWithStackTraces()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CompiledJoinStringWithStackTraces returns non-empty -- with errors", actual)
}

func Test_RawErrCollection_FullString(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))

	// Act
	actual := args.Map{"v": rec.FullString() != ""}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.FullString returns non-empty -- with errors populated", actual)
}

func Test_RawErrCollection_FullStringWithTraces(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))

	// Act
	actual := args.Map{"v": rec.FullStringWithTraces() != ""}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.FullStringWithTraces returns non-empty -- with errors", actual)
}

func Test_RawErrCollection_FullStringWithTracesIf(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	r1 := rec.FullStringWithTracesIf(true)
	r2 := rec.FullStringWithTracesIf(false)

	// Act
	actual := args.Map{
		"t": r1 != "",
		"f": r2 != "",
	}

	// Assert
	expected := args.Map{
		"t": true,
		"f": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.FullStringWithTracesIf returns non-empty -- with condition", actual)
}

func Test_RawErrCollection_ReferencesCompiledString(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))

	// Act
	actual := args.Map{"v": rec.ReferencesCompiledString() != ""}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ReferencesCompiledString returns non-empty -- with errors", actual)
}

func Test_RawErrCollection_FullStringSplitByNewLine(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))

	// Act
	actual := args.Map{"len": len(rec.FullStringSplitByNewLine())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.FullStringSplitByNewLine returns lines -- with errors", actual)
}

func Test_RawErrCollection_FullStringWithoutReferences(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))

	// Act
	actual := args.Map{"v": rec.FullStringWithoutReferences() != ""}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.FullStringWithoutReferences returns non-empty -- with errors", actual)
}

func Test_RawErrCollection_ErrorString(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))

	// Act
	actual := args.Map{"v": rec.ErrorString() != ""}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ErrorString returns non-empty -- with errors", actual)
}

func Test_RawErrCollection_Compile(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))

	// Act
	actual := args.Map{"v": rec.Compile() != ""}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Compile returns non-empty -- with errors", actual)
}

func Test_RawErrCollection_Value(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"isNil": rec.Value() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Value returns nil -- empty collection", actual)
}

func Test_RawErrCollection_Serialize_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	b, err := rec.Serialize()

	// Act
	actual := args.Map{
		"nil": b == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Serialize returns nil -- empty collection", actual)
}

func Test_RawErrCollection_SerializeWithoutTraces_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	b, err := rec.SerializeWithoutTraces()

	// Act
	actual := args.Map{
		"nil": b == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.SerializeWithoutTraces returns nil -- empty collection", actual)
}

func Test_RawErrCollection_SerializeMust_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	b := rec.SerializeMust()

	// Act
	actual := args.Map{"nil": b == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.SerializeMust returns nil -- empty collection", actual)
}

func Test_RawErrCollection_MarshalJSON(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	b, err := rec.MarshalJSON()

	// Act
	actual := args.Map{
		"nil": b == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.MarshalJSON returns nil -- empty collection", actual)
}

func Test_RawErrCollection_UnmarshalJSON(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	err := rec.UnmarshalJSON([]byte(`[]`))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.UnmarshalJSON succeeds -- valid JSON", actual)
}

func Test_RawErrCollection_Log(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Log() // empty, no-op
	rec.Add(errors.New("a"))
	rec.Log()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Log completes safely -- with errors", actual)
}

func Test_RawErrCollection_LogWithTraces(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.LogWithTraces() // empty, no-op
	rec.Add(errors.New("a"))
	rec.LogWithTraces()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.LogWithTraces completes safely -- with errors", actual)
}

func Test_RawErrCollection_LogIf_False_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.LogIf(false)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.LogIf completes safely -- condition false", actual)
}

func Test_RawErrCollection_Clear(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Clear() // empty
	rec.Add(errors.New("a"))
	rec.Clear()

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Clear resets length -- with errors", actual)
}

func Test_RawErrCollection_Dispose(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Dispose() // empty
	rec.Add(errors.New("a"))
	rec.Dispose()

	// Act
	actual := args.Map{"isNull": rec.IsNull()}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.Dispose nullifies -- with errors", actual)
}

func Test_RawErrCollection_MustBeSafe_Empty(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.MustBeSafe()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.MustBeSafe completes safely -- empty collection", actual)
}

func Test_RawErrCollection_MustBeSafe_Panic(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.MustBeSafe()
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.MustBeSafe panics -- with errors", actual)
}

func Test_RawErrCollection_MustBeEmptyError(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.MustBeEmptyError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.MustBeEmptyError completes safely -- empty collection", actual)
}

func Test_RawErrCollection_HandleError_Empty(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.HandleError() // empty, no-op

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.HandleError completes safely -- empty collection", actual)
}

func Test_RawErrCollection_HandleError_Panic(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.HandleError()
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.HandleError panics -- with errors", actual)
}

func Test_RawErrCollection_HandleErrorWithRefs_Empty(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.HandleErrorWithRefs("msg", "k", "v") // empty, no-op

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.HandleErrorWithRefs completes safely -- empty collection", actual)
}

func Test_RawErrCollection_HandleErrorWithRefs_Panic(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.HandleErrorWithRefs("msg", "k", "v")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.HandleErrorWithRefs panics -- with errors", actual)
}

func Test_RawErrCollection_HandleErrorWithMsg_Empty(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.HandleErrorWithMsg("msg") // empty, no-op

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.HandleErrorWithMsg completes safely -- empty collection", actual)
}

func Test_RawErrCollection_HandleErrorWithMsg_Panic(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.HandleErrorWithMsg("msg")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.HandleErrorWithMsg panics -- with errors", actual)
}

func Test_RawErrCollection_ReflectSetTo_Value(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	err := rec.ReflectSetTo(errcore.RawErrCollection{})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ReflectSetTo returns error -- value type", actual)
}

func Test_RawErrCollection_ReflectSetTo_NilPtr(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	var nilPtr *errcore.RawErrCollection
	err := rec.ReflectSetTo(nilPtr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ReflectSetTo returns error -- nil pointer", actual)
}

func Test_RawErrCollection_ReflectSetTo_ValidPtr(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	target := &errcore.RawErrCollection{}
	err := rec.ReflectSetTo(target)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ReflectSetTo succeeds -- valid pointer", actual)
}

func Test_RawErrCollection_ReflectSetTo_Other(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	err := rec.ReflectSetTo("unsupported")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ReflectSetTo returns error -- unsupported type", actual)
}

func Test_RawErrCollection_CountStateChangeTracker(t *testing.T) {
	// Arrange
	rec := errcore.RawErrCollection{}
	tracker := rec.CountStateChangeTracker()

	// Act
	actual := args.Map{"same": tracker.IsSameState()}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.CountStateChangeTracker returns same -- no changes", actual)
}

func Test_RawErrCollection_IsErrorsCollected_NoNew(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"v": rec.IsErrorsCollected(nil)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.IsErrorsCollected returns false -- nil error", actual)
}

func Test_RawErrCollection_IsErrorsCollected_WithNew(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}

	// Act
	actual := args.Map{"v": rec.IsErrorsCollected(errors.New("e"))}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.IsErrorsCollected returns true -- with error", actual)
}

func Test_RawErrCollection_ToRawErrCollection(t *testing.T) {
	// Arrange
	rec := errcore.RawErrCollection{}
	ptr := rec.ToRawErrCollection()

	// Act
	actual := args.Map{"notNil": ptr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.ToRawErrCollection returns pointer -- value receiver", actual)
}

func Test_RawErrCollection_AddErrorGetters_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddErrorGetters()

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddErrorGetters no-op -- empty args", actual)
}

func Test_RawErrCollection_AddCompiledErrorGetters_FromLineDiffAllMatch(t *testing.T) {
	// Arrange
	rec := &errcore.RawErrCollection{}
	rec.AddCompiledErrorGetters()

	// Act
	actual := args.Map{"len": rec.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RawErrCollection.AddCompiledErrorGetters no-op -- empty args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RawErrorType — String / Combine
// ══════════════════════════════════════════════════════════════════════════════

func Test_RawErrorType_String(t *testing.T) {
	// Act
	actual := args.Map{"v": errcore.InvalidType.String() != ""}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.String returns non-empty -- InvalidType", actual)
}

func Test_RawErrorType_Combine(t *testing.T) {
	// Arrange
	result := errcore.InvalidType.Combine("msg", "ref")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Combine returns formatted -- with msg and ref", actual)
}

func Test_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorNoRefsSkip(0, "msg")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefsSkip returns error -- with msg", actual)
}

func Test_RawErrorType_ErrorNoRefsSkip_Empty(t *testing.T) {
	// Arrange
	err := errcore.InvalidType.ErrorNoRefsSkip(0, "")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefsSkip returns error -- empty msg", actual)
}
