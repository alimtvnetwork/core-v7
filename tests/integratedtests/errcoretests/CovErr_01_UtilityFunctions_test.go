package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// errcore Coverage — Batch 1: Utility functions, formatters, converters
// ══════════════════════════════════════════════════════════════════════════════

// --- CountStateChangeTracker ---

type cov14MockLengthGetter struct {
	length int
}

func (m *cov14MockLengthGetter) Length() int {
	return m.length
}

func Test_CovErr_01_CountStateChangeTracker(t *testing.T) {
	// Arrange
	mg := &cov14MockLengthGetter{length: 5}
	tracker := errcore.NewCountStateChangeTracker(mg)

	// Act
	actual := args.Map{"result": tracker.IsSameState()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same state", actual)
	actual = args.Map{"result": tracker.IsValid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual = args.Map{"result": tracker.IsSuccess()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
	actual = args.Map{"result": tracker.HasChanges()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no changes", actual)
	actual = args.Map{"result": tracker.IsFailed()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not failed", actual)
	actual = args.Map{"result": tracker.IsSameStateUsingCount(5)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same state using count", actual)
	actual = args.Map{"result": tracker.IsSameStateUsingCount(3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not same state", actual)
	// change length
	mg.length = 10
	actual = args.Map{"result": tracker.IsSameState()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected different state", actual)
	actual = args.Map{"result": tracker.HasChanges()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected changes", actual)
	actual = args.Map{"result": tracker.IsFailed()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected failed", actual)
}

// --- CombineWithMsgType ---

func Test_CovErr_02_CombineWithMsgTypeNoStack(t *testing.T) {
	// Arrange
	r := errcore.CombineWithMsgTypeNoStack(errcore.OutOfRangeType, "extra", "ref")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// empty otherMsg
	r2 := errcore.CombineWithMsgTypeNoStack(errcore.OutOfRangeType, "", "ref")
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil reference
	r3 := errcore.CombineWithMsgTypeNoStack(errcore.OutOfRangeType, "msg", nil)
	actual = args.Map{"result": r3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_03_CombineWithMsgTypeStackTrace(t *testing.T) {
	// Arrange
	r := errcore.CombineWithMsgTypeStackTrace(errcore.OutOfRangeType, "msg", "ref")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- ConcatMessageWithErr ---

func Test_CovErr_04_ConcatMessageWithErr(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErr("prefix", errors.New("original"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// nil err
	err2 := errcore.ConcatMessageWithErr("prefix", nil)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovErr_05_ConcatMessageWithErrWithStackTrace(t *testing.T) {
	// Arrange
	err := errcore.ConcatMessageWithErrWithStackTrace("prefix", errors.New("orig"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.ConcatMessageWithErrWithStackTrace("prefix", nil)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- Combine ---

func Test_CovErr_06_Combine(t *testing.T) {
	// Arrange
	r := errcore.Combine("generic", "other", "ref")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- EnumRangeNotMeet ---

func Test_CovErr_07_EnumRangeNotMeet(t *testing.T) {
	// Arrange
	r := errcore.EnumRangeNotMeet(1, 10, []int{1, 2, 3})

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil wholeRange
	r2 := errcore.EnumRangeNotMeet(1, 10, nil)
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- ErrorToSplitLines ---

func Test_CovErr_08_ErrorToSplitLines(t *testing.T) {
	// Arrange
	lines := errcore.ErrorToSplitLines(errors.New("a\nb"))

	// Act
	actual := args.Map{"result": len(lines) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	lines2 := errcore.ErrorToSplitLines(nil)
	actual = args.Map{"result": len(lines2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovErr_09_ErrorToSplitNonEmptyLines(t *testing.T) {
	// Arrange
	lines := errcore.ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))

	// Act
	actual := args.Map{"result": len(lines) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
}

// --- ErrorWithRef ---

func Test_CovErr_10_ErrorWithRef(t *testing.T) {
	// Arrange
	r := errcore.ErrorWithRef(errors.New("err"), "ref")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil err
	r2 := errcore.ErrorWithRef(nil, "ref")
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	// nil ref
	r3 := errcore.ErrorWithRef(errors.New("err"), nil)
	actual = args.Map{"result": r3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// empty ref
	r4 := errcore.ErrorWithRef(errors.New("err"), "")
	actual = args.Map{"result": r4 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_11_ErrorWithRefToError(t *testing.T) {
	// Arrange
	err := errcore.ErrorWithRefToError(errors.New("err"), "ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.ErrorWithRefToError(nil, "ref")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- ErrorWithCompiledTraceRef ---

func Test_CovErr_12_ErrorWithCompiledTraceRef(t *testing.T) {
	// Arrange
	// nil err
	r := errcore.ErrorWithCompiledTraceRef(nil, "trace", "ref")

	// Act
	actual := args.Map{"result": r != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	// empty traces
	r2 := errcore.ErrorWithCompiledTraceRef(errors.New("err"), "", "ref")
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil reference
	r3 := errcore.ErrorWithCompiledTraceRef(errors.New("err"), "trace", nil)
	actual = args.Map{"result": r3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// all present
	r4 := errcore.ErrorWithCompiledTraceRef(errors.New("err"), "trace", "ref")
	actual = args.Map{"result": r4 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- MeaningfulError ---

func Test_CovErr_13_MeaningfulError(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulError(errcore.OutOfRangeType, "func", errors.New("orig"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.MeaningfulError(errcore.OutOfRangeType, "func", nil)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovErr_14_MeaningfulErrorWithData(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulErrorWithData(errcore.OutOfRangeType, "func", errors.New("orig"), "data")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.MeaningfulErrorWithData(errcore.OutOfRangeType, "func", nil, "data")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovErr_15_MeaningfulMessageError(t *testing.T) {
	// Arrange
	err := errcore.MeaningfulMessageError(errcore.OutOfRangeType, "func", errors.New("orig"), " extra")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.MeaningfulMessageError(errcore.OutOfRangeType, "func", nil, "extra")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- MeaningfulErrorHandle ---

func Test_CovErr_16_MeaningfulErrorHandle_NilSafe(t *testing.T) {
	// nil should not panic
	errcore.MeaningfulErrorHandle(errcore.OutOfRangeType, "func", nil)
}

// --- PathMeaningfulMessage ---

func Test_CovErr_17_PathMeaningfulMessage(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulMessage(errcore.OutOfRangeType, "func", "/path", "msg1", "msg2")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.PathMeaningfulMessage(errcore.OutOfRangeType, "func", "/path")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- PathMeaningfulError ---

func Test_CovErr_18_PathMeaningfulError(t *testing.T) {
	// Arrange
	err := errcore.PathMeaningfulError(errcore.OutOfRangeType, errors.New("orig"), "/path")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.PathMeaningfulError(errcore.OutOfRangeType, nil, "/path")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- MergeErrors ---

func Test_CovErr_19_MergeErrors(t *testing.T) {
	// Arrange
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.MergeErrors()
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovErr_20_MergeErrorsToString(t *testing.T) {
	// Arrange
	r := errcore.MergeErrorsToString(", ", errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.MergeErrorsToString(", ")
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovErr_21_MergeErrorsToStringDefault(t *testing.T) {
	// Arrange
	r := errcore.MergeErrorsToStringDefault(errors.New("a"))

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.MergeErrorsToStringDefault()
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// --- ManyErrorToSingle ---

func Test_CovErr_22_ManyErrorToSingle(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingle([]error{errors.New("a")})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.ManyErrorToSingle(nil)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovErr_23_ManyErrorToSingleDirect(t *testing.T) {
	// Arrange
	err := errcore.ManyErrorToSingleDirect(errors.New("a"), errors.New("b"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// --- SliceError ---

func Test_CovErr_24_SliceError(t *testing.T) {
	// Arrange
	err := errcore.SliceError(", ", []string{"a", "b"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.SliceError(", ", []string{})
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovErr_25_SliceErrorDefault(t *testing.T) {
	// Arrange
	err := errcore.SliceErrorDefault([]string{"a"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovErr_26_SliceErrorsToStrings(t *testing.T) {
	// Arrange
	ss := errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b"))

	// Act
	actual := args.Map{"result": len(ss) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	ss2 := errcore.SliceErrorsToStrings()
	actual = args.Map{"result": len(ss2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovErr_27_SliceToError_SliceToErrorPtr(t *testing.T) {
	// Arrange
	err := errcore.SliceToError([]string{"a"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.SliceToError([]string{})
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err3 := errcore.SliceToErrorPtr([]string{"a"})
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err4 := errcore.SliceToErrorPtr([]string{})
	actual = args.Map{"result": err4 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- VarTwo / VarThree ---

func Test_CovErr_28_VarTwo(t *testing.T) {
	// Arrange
	r := errcore.VarTwo(true, "a", 1, "b", 2)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.VarTwo(false, "a", 1, "b", 2)
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_29_VarTwoNoType(t *testing.T) {
	// Arrange
	r := errcore.VarTwoNoType("a", 1, "b", 2)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_30_VarThree(t *testing.T) {
	// Arrange
	r := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_31_VarThreeNoType(t *testing.T) {
	// Arrange
	r := errcore.VarThreeNoType("a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- VarMap / VarMapStrings / VarNameValues ---

func Test_CovErr_32_VarMap(t *testing.T) {
	// Arrange
	r := errcore.VarMap(map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.VarMap(map[string]any{})
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovErr_33_VarMapStrings(t *testing.T) {
	// Arrange
	ss := errcore.VarMapStrings(map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": len(ss) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	ss2 := errcore.VarMapStrings(map[string]any{})
	actual = args.Map{"result": len(ss2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovErr_34_VarNameValues(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "a", Value: 1}
	r := errcore.VarNameValues(nv)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.VarNameValues()
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovErr_35_VarNameValuesJoiner(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "a", Value: 1}
	r := errcore.VarNameValuesJoiner(", ", nv)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.VarNameValuesJoiner(", ")
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovErr_36_VarNameValuesStrings(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "a", Value: 1}
	ss := errcore.VarNameValuesStrings(nv)

	// Act
	actual := args.Map{"result": len(ss) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	ss2 := errcore.VarNameValuesStrings()
	actual = args.Map{"result": len(ss2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// --- MessageVarTwo / MessageVarThree / MessageVarMap / MessageNameValues ---

func Test_CovErr_37_MessageVarTwo(t *testing.T) {
	// Arrange
	r := errcore.MessageVarTwo("msg", "a", 1, "b", 2)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_38_MessageVarThree(t *testing.T) {
	// Arrange
	r := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_39_MessageVarMap(t *testing.T) {
	// Arrange
	r := errcore.MessageVarMap("msg", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.MessageVarMap("msg", map[string]any{})
	actual = args.Map{"result": r2 != "msg"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected just msg", actual)
}

func Test_CovErr_40_MessageNameValues(t *testing.T) {
	// Arrange
	nv := namevalue.StringAny{Name: "a", Value: 1}
	r := errcore.MessageNameValues("msg", nv)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.MessageNameValues("msg")
	actual = args.Map{"result": r2 != "msg"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected just msg", actual)
}

// --- MessageWithRef / MessageWithRefToError ---

func Test_CovErr_41_MessageWithRef(t *testing.T) {
	// Arrange
	r := errcore.MessageWithRef("msg", "ref")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_42_MessageWithRefToError(t *testing.T) {
	// Arrange
	err := errcore.MessageWithRefToError("msg", "ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// --- SourceDestination ---

func Test_CovErr_43_SourceDestination(t *testing.T) {
	// Arrange
	r := errcore.SourceDestination(true, "src", "dst")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.SourceDestination(false, "src", "dst")
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_44_SourceDestinationErr(t *testing.T) {
	// Arrange
	err := errcore.SourceDestinationErr(false, "src", "dst")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovErr_45_SourceDestinationNoType(t *testing.T) {
	// Arrange
	r := errcore.SourceDestinationNoType("src", "dst")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- ToError / ToString / ToStringPtr / ToValueString ---

func Test_CovErr_46_ToError(t *testing.T) {
	// Arrange
	err := errcore.ToError("msg")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.ToError("")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovErr_47_ToString(t *testing.T) {
	// Arrange
	r := errcore.ToString(errors.New("msg"))

	// Act
	actual := args.Map{"result": r != "msg"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected msg", actual)
	r2 := errcore.ToString(nil)
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovErr_48_ToStringPtr(t *testing.T) {
	// Arrange
	r := errcore.ToStringPtr(errors.New("msg"))

	// Act
	actual := args.Map{"result": *r != "msg"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected msg", actual)
	r2 := errcore.ToStringPtr(nil)
	actual = args.Map{"result": *r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovErr_49_ToValueString(t *testing.T) {
	// Arrange
	r := errcore.ToValueString("hello")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_50_ToExitError(t *testing.T) {
	// Arrange
	r := errcore.ToExitError(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	r2 := errcore.ToExitError(errors.New("not exit error"))
	actual = args.Map{"result": r2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for non-exit error", actual)
}

// --- StringLinesToQuoteLines ---

func Test_CovErr_51_StringLinesToQuoteLines(t *testing.T) {
	// Arrange
	ss := errcore.StringLinesToQuoteLines([]string{"a", "b"})

	// Act
	actual := args.Map{"result": len(ss) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	ss2 := errcore.StringLinesToQuoteLines([]string{})
	actual = args.Map{"result": len(ss2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovErr_52_StringLinesToQuoteLinesToSingle(t *testing.T) {
	// Arrange
	r := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_53_LinesToDoubleQuoteLinesWithTabs(t *testing.T) {
	// Arrange
	ss := errcore.LinesToDoubleQuoteLinesWithTabs(4, []string{"a"})

	// Act
	actual := args.Map{"result": len(ss) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	ss2 := errcore.LinesToDoubleQuoteLinesWithTabs(0, []string{})
	actual = args.Map{"result": len(ss2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// --- MsgHeader ---

func Test_CovErr_54_MsgHeader(t *testing.T) {
	// Arrange
	r := errcore.MsgHeader("test")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_55_MsgHeaderIf(t *testing.T) {
	// Arrange
	r := errcore.MsgHeaderIf(true, "test")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.MsgHeaderIf(false, "test")
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_56_MsgHeaderPlusEnding(t *testing.T) {
	// Arrange
	r := errcore.MsgHeaderPlusEnding("header", "message")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- MustBeEmpty ---

func Test_CovErr_57_MustBeEmpty_NilSafe(t *testing.T) {
	errcore.MustBeEmpty(nil)
}

// --- HandleErr / HandleErrMessage ---

func Test_CovErr_58_HandleErr_NilSafe(t *testing.T) {
	errcore.HandleErr(nil)
}

func Test_CovErr_59_HandleErrMessage_NilSafe(t *testing.T) {
	errcore.HandleErrMessage("")
}

// --- SimpleHandleErr ---

func Test_CovErr_60_SimpleHandleErr_NilSafe(t *testing.T) {
	errcore.SimpleHandleErr(nil, "msg")
}

// --- SimpleHandleErrMany ---

func Test_CovErr_61_SimpleHandleErrMany_NilSafe(t *testing.T) {
	errcore.SimpleHandleErrMany("msg")
	errcore.SimpleHandleErrMany("msg", nil)
}

// --- PanicOnIndexOutOfRange ---

func Test_CovErr_62_PanicOnIndexOutOfRange_Valid(t *testing.T) {
	errcore.PanicOnIndexOutOfRange(5, []int{0, 1, 4})
}

// --- RangeNotMeet / PanicRangeNotMeet ---

func Test_CovErr_63_RangeNotMeet(t *testing.T) {
	// Arrange
	r := errcore.RangeNotMeet("msg", 0, 10, nil)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.RangeNotMeet("msg", 0, 10, "range")
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_64_PanicRangeNotMeet(t *testing.T) {
	// Arrange
	r := errcore.PanicRangeNotMeet("msg", 0, 10, nil)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.PanicRangeNotMeet("msg", 0, 10, "range")
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- Ref / RefToError ---

func Test_CovErr_65_Ref(t *testing.T) {
	// Arrange
	r := errcore.Ref("ref")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := errcore.Ref(nil)
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovErr_66_RefToError(t *testing.T) {
	// Arrange
	err := errcore.RefToError("ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := errcore.RefToError(nil)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- GherkinsString ---

func Test_CovErr_67_GherkinsString(t *testing.T) {
	// Arrange
	r := errcore.GherkinsString(1, "feature", "given", "when", "then")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovErr_68_GherkinsStringWithExpectation(t *testing.T) {
	// Arrange
	r := errcore.GherkinsStringWithExpectation(1, "feature", "given", "when", "then", "actual", "expect")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- FmtDebug / FmtDebugIf / ValidPrint / FailedPrint / PrintError ---

func Test_CovErr_69_FmtDebug(t *testing.T) {
	errcore.FmtDebug("test %d", 1)
}

func Test_CovErr_70_FmtDebugIf(t *testing.T) {
	errcore.FmtDebugIf(false, "test %d", 1)
	errcore.FmtDebugIf(true, "test %d", 1)
}

func Test_CovErr_71_ValidPrint(t *testing.T) {
	errcore.ValidPrint(true, "val")
	errcore.ValidPrint(false, "val")
}

func Test_CovErr_72_FailedPrint(t *testing.T) {
	errcore.FailedPrint(true, "val")
	errcore.FailedPrint(false, "val")
}

func Test_CovErr_73_PrintError(t *testing.T) {
	errcore.PrintError(nil)
	errcore.PrintError(errors.New("err"))
}

func Test_CovErr_74_PrintErrorWithTestIndex(t *testing.T) {
	errcore.PrintErrorWithTestIndex(0, "header", nil)
	errcore.PrintErrorWithTestIndex(0, "header", errors.New("err"))
}
