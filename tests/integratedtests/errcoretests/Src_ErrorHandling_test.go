package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func callPanicsErrcore(fn func()) bool {
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		fn()
	}()
	return didPanic
}

// ── CompiledError ──

func Test_Src01_CompiledError_NilErr(t *testing.T) {
	// Arrange
	tc := compiledErrorTestCases[0]
	// Act
	result := errcore.CompiledError(nil, "msg")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": result == nil})
}

func Test_Src01_CompiledError_EmptyMsg(t *testing.T) {
	// Arrange
	tc := compiledErrorTestCases[1]
	err := errors.New("base")
	// Act
	result := errcore.CompiledError(err, "")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isSame": result == err})
}

func Test_Src01_CompiledError_WithMsg(t *testing.T) {
	// Arrange
	tc := compiledErrorTestCases[2]
	// Act
	result := errcore.CompiledError(errors.New("base"), "context")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": result != nil && result.Error() != ""})
}

func Test_Src01_CompiledErrorString_NilErr(t *testing.T) {
	// Arrange
	tc := compiledErrorTestCases[3]
	// Act
	result := errcore.CompiledErrorString(nil, "msg")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": result == ""})
}

func Test_Src01_CompiledErrorString_WithErr(t *testing.T) {
	// Arrange
	tc := compiledErrorTestCases[4]
	// Act
	result := errcore.CompiledErrorString(errors.New("base"), "ctx")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": result != ""})
}

// ── JoinErrors ──

func Test_Src01_JoinErrors_WithErrors(t *testing.T) {
	// Arrange
	tc := joinErrorsTestCases[0]
	// Act
	result := errcore.JoinErrors(errors.New("a"), errors.New("b"))
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": result != nil})
}

func Test_Src01_JoinErrors_Empty(t *testing.T) {
	// Arrange
	tc := joinErrorsTestCases[1]
	// Act
	result := errcore.JoinErrors()
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": result == nil})
}

// ── ErrorWithRef ──

func Test_Src01_ErrorWithRef_NilErr(t *testing.T) {
	// Arrange
	tc := errorWithRefTestCases[0]
	// Act
	result := errcore.ErrorWithRef(nil, "ref")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": result == ""})
}

func Test_Src01_ErrorWithRef_NilRef(t *testing.T) {
	// Arrange
	tc := errorWithRefTestCases[1]
	// Act
	result := errcore.ErrorWithRef(errors.New("e"), nil)
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": result})
}

func Test_Src01_ErrorWithRef_EmptyRef(t *testing.T) {
	// Arrange
	tc := errorWithRefTestCases[2]
	// Act
	result := errcore.ErrorWithRef(errors.New("e"), "")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": result})
}

func Test_Src01_ErrorWithRef_WithRef(t *testing.T) {
	// Arrange
	tc := errorWithRefTestCases[3]
	// Act
	result := errcore.ErrorWithRef(errors.New("e"), "ref")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": result != ""})
}

func Test_Src01_ErrorWithRefToError_NilErr(t *testing.T) {
	// Arrange
	tc := errorWithRefTestCases[4]
	// Act
	result := errcore.ErrorWithRefToError(nil, "ref")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": result == nil})
}

func Test_Src01_ErrorWithRefToError_WithErr(t *testing.T) {
	// Arrange
	tc := errorWithRefTestCases[5]
	// Act
	result := errcore.ErrorWithRefToError(errors.New("e"), "ref")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": result != nil})
}

// ── ErrorWithCompiledTraceRef ──

func Test_Src01_ErrorWithCompiledTraceRef_NilErr(t *testing.T) {
	// Arrange
	tc := errorWithCompiledTraceRefTestCases[0]
	// Act
	result := errcore.ErrorWithCompiledTraceRef(nil, "traces", "ref")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": result == ""})
}

func Test_Src01_ErrorWithCompiledTraceRef_EmptyTraces(t *testing.T) {
	// Arrange
	tc := errorWithCompiledTraceRefTestCases[1]
	// Act
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "", "ref")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": result != ""})
}

func Test_Src01_ErrorWithCompiledTraceRef_NilRef(t *testing.T) {
	// Arrange
	tc := errorWithCompiledTraceRefTestCases[2]
	// Act
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "traces", nil)
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": result != ""})
}

func Test_Src01_ErrorWithCompiledTraceRef_All(t *testing.T) {
	// Arrange
	tc := errorWithCompiledTraceRefTestCases[3]
	// Act
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "traces", "ref")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": result != ""})
}

func Test_Src01_ErrorWithCompiledTraceRefToError_NilErr(t *testing.T) {
	// Arrange
	tc := errorWithCompiledTraceRefTestCases[4]
	// Act
	result := errcore.ErrorWithCompiledTraceRefToError(nil, "t", "r")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": result == nil})
}

func Test_Src01_ErrorWithCompiledTraceRefToError_WithErr(t *testing.T) {
	// Arrange
	tc := errorWithCompiledTraceRefTestCases[5]
	// Act
	result := errcore.ErrorWithCompiledTraceRefToError(errors.New("e"), "t", "r")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": result != nil})
}

// ── ErrorWithTracesRefToError ──

func Test_Src01_ErrorWithTracesRefToError_NilErr(t *testing.T) {
	// Arrange
	tc := errorWithTracesRefToErrorTestCases[0]
	// Act
	result := errcore.ErrorWithTracesRefToError(nil, []string{"t"}, "r")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": result == nil})
}

func Test_Src01_ErrorWithTracesRefToError_EmptyTraces(t *testing.T) {
	// Arrange
	tc := errorWithTracesRefToErrorTestCases[1]
	// Act
	result := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{}, "r")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": result != nil})
}

func Test_Src01_ErrorWithTracesRefToError_WithTraces(t *testing.T) {
	// Arrange
	tc := errorWithTracesRefToErrorTestCases[2]
	// Act
	result := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{"t1", "t2"}, "r")
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": result != nil})
}

// ── ConcatMessageWithErr ──

func Test_Src01_ConcatMessageWithErr_NilErr(t *testing.T) {
	// Arrange
	tc := concatMessageTestCases[0]
	// Act
	result := errcore.ConcatMessageWithErr("msg", nil)
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": result == nil})
}

func Test_Src01_ConcatMessageWithErr_WithErr(t *testing.T) {
	// Arrange
	tc := concatMessageTestCases[1]
	// Act
	result := errcore.ConcatMessageWithErr("msg", errors.New("e"))
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": result != nil})
}

func Test_Src01_ConcatMessageWithErrWithStackTrace_NilErr(t *testing.T) {
	// Arrange
	tc := concatMessageTestCases[2]
	// Act
	result := errcore.ConcatMessageWithErrWithStackTrace("msg", nil)
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": result == nil})
}

func Test_Src01_ConcatMessageWithErrWithStackTrace_WithErr(t *testing.T) {
	// Arrange
	tc := concatMessageTestCases[3]
	// Act
	result := errcore.ConcatMessageWithErrWithStackTrace("msg", errors.New("e"))
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": result != nil})
}

// ── ErrorToSplitLines ──

func Test_Src01_ErrorToSplitLines_NilErr(t *testing.T) {
	// Arrange
	tc := errorToSplitLinesTestCases[0]
	// Act
	result := errcore.ErrorToSplitLines(nil)
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(result)})
}

func Test_Src01_ErrorToSplitLines_WithErr(t *testing.T) {
	// Arrange
	tc := errorToSplitLinesTestCases[1]
	// Act
	result := errcore.ErrorToSplitLines(errors.New("a\nb"))
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(result)})
}

func Test_Src01_ErrorToSplitNonEmptyLines_WithErr(t *testing.T) {
	// Arrange
	tc := errorToSplitLinesTestCases[2]
	// Act
	result := errcore.ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"ge2": len(result) >= 2})
}

func Test_Src01_ErrorToSplitNonEmptyLines_Nil(t *testing.T) {
	// Arrange
	tc := errorToSplitLinesTestCases[3]
	// Act
	result := errcore.ErrorToSplitNonEmptyLines(nil)
	_ = result
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noErr": true})
}

// ── Handlers ──

func Test_Src05_HandleErr_Nil(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[0]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.HandleErr(nil) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_HandleErr_WithErr(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[1]
	// Act
	panics := callPanicsErrcore(func() { errcore.HandleErr(errors.New("e")) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_Src05_HandleErrMessage_Empty(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[2]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.HandleErrMessage("") })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_HandleErrMessage_WithMsg(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[3]
	// Act
	panics := callPanicsErrcore(func() { errcore.HandleErrMessage("msg") })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_Src05_SimpleHandleErr_Nil(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[4]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.SimpleHandleErr(nil, "msg") })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_SimpleHandleErr_WithErr(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[5]
	// Act
	panics := callPanicsErrcore(func() { errcore.SimpleHandleErr(errors.New("e"), "msg") })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_Src05_SimpleHandleErrMany_Nil(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[6]
	// Act
	noPanic := !callPanicsErrcore(func() {
		errcore.SimpleHandleErrMany("msg")
		errcore.SimpleHandleErrMany("msg", nil)
	})
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_SimpleHandleErrMany_WithErr(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[7]
	// Act
	panics := callPanicsErrcore(func() { errcore.SimpleHandleErrMany("msg", errors.New("e")) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_Src05_MustBeEmpty_Nil(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[8]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.MustBeEmpty(nil) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_MustBeEmpty_WithErr(t *testing.T) {
	// Arrange
	tc := handleErrTestCases[9]
	// Act
	panics := callPanicsErrcore(func() { errcore.MustBeEmpty(errors.New("e")) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_Src05_HandleGetters_Nil(t *testing.T) {
	// Arrange
	tc0 := handleGetterTestCases[0]
	tc1 := handleGetterTestCases[1]
	tc2 := handleGetterTestCases[2]
	tc3 := handleGetterTestCases[3]

	// Act & Assert
	noPanic0 := !callPanicsErrcore(func() { errcore.HandleCompiledErrorGetter(nil) })
	tc0.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic0})

	noPanic1 := !callPanicsErrcore(func() { errcore.HandleCompiledErrorWithTracesGetter(nil) })
	tc1.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic1})

	noPanic2 := !callPanicsErrcore(func() { errcore.HandleErrorGetter(nil) })
	tc2.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic2})

	noPanic3 := !callPanicsErrcore(func() { errcore.HandleFullStringsWithTracesGetter(nil) })
	tc3.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic3})
}

func Test_Src05_PrintError_Nil(t *testing.T) {
	// Arrange
	tc := printErrorTestCases[0]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintError(nil) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_PrintError_WithErr(t *testing.T) {
	// Arrange
	tc := printErrorTestCases[1]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintError(errors.New("e")) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_PrintErrorWithTestIndex_Nil(t *testing.T) {
	// Arrange
	tc := printErrorTestCases[2]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintErrorWithTestIndex(0, "test", nil) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_PrintErrorWithTestIndex_WithErr(t *testing.T) {
	// Arrange
	tc := printErrorTestCases[3]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PrintErrorWithTestIndex(0, "test", errors.New("e")) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_PanicOnIndexOutOfRange_Valid(t *testing.T) {
	// Arrange
	tc := panicOnIndexTestCases[0]
	// Act
	noPanic := !callPanicsErrcore(func() { errcore.PanicOnIndexOutOfRange(5, []int{0, 1, 2}) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_Src05_PanicOnIndexOutOfRange_OutOfRange(t *testing.T) {
	// Arrange
	tc := panicOnIndexTestCases[1]
	// Act
	panics := callPanicsErrcore(func() { errcore.PanicOnIndexOutOfRange(2, []int{5}) })
	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}
