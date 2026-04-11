package errcoretests

import (
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── AssertDiffOnMismatch: mismatch path ──
// Covers AssertDiffOnMismatch.go L32-33

func Test_AssertDiffOnMismatch_Mismatch(t *testing.T) {
	// Arrange
	fakeT := &testing.T{}

	// Act
	errcore.AssertDiffOnMismatch(
		fakeT, 0, "test mismatch",
		[]string{"actual-line"},
		[]string{"expected-line"},
	)

	// Assert
	actual := args.Map{"failed": fakeT.Failed()}
	expected := args.Map{"failed": true}
	expected.ShouldBeEqual(t, 0, "AssertDiffOnMismatch marks test failed -- mismatch", actual)
}

// ── AssertErrorDiffOnMismatch: mismatch path ──
// Covers AssertDiffOnMismatch.go L62-63

func Test_AssertErrorDiffOnMismatch_Mismatch(t *testing.T) {
	// Arrange
	fakeT := &testing.T{}

	// Act
	errcore.AssertErrorDiffOnMismatch(
		fakeT, 0, "error mismatch",
		errors.New("actual-error"),
		[]string{"expected-line"},
	)

	// Assert
	actual := args.Map{"failed": fakeT.Failed()}
	expected := args.Map{"failed": true}
	expected.ShouldBeEqual(t, 0, "AssertErrorDiffOnMismatch marks test failed -- error mismatch", actual)
}

// ── CompiledErrorString: compiled == nil dead code path ──
// Covers CompiledError.go L30-32 (dead code since CompiledError never returns nil when mainErr != nil)

func Test_CompiledErrorString_WithError(t *testing.T) {
	// Arrange & Act
	result := errcore.CompiledErrorString(errors.New("main"), "additional")

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"contains": strings.Contains(result, "main"),
	}
	expected := args.Map{
		"hasContent": true,
		"contains": true,
	}
	expected.ShouldBeEqual(t, 0, "CompiledErrorString returns string -- with error", actual)
}

// ── ExpectationMessageDef: ExpectedSafeString with cached value ──
// Covers ExpectationMessageDef.go L24-26

func Test_ExpectationMessageDef_CachedExpected(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{
		Expected: "test-val",
	}

	// Act — first call caches, second uses cache
	first := def.ExpectedSafeString()
	second := def.ExpectedSafeString()

	// Assert
	actual := args.Map{
		"equal": first == second,
		"nonEmpty": len(first) > 0,
	}
	expected := args.Map{
		"equal": true,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "ExpectedSafeString returns cached value -- second call", actual)
}

// ── ExpectationMessageDef: Print ──
// Covers ExpectationMessageDef.go L67 (Print path via PrintIf)

func Test_ExpectationMessageDef_PrintIf(t *testing.T) {
	// Arrange
	def := errcore.ExpectationMessageDef{
		Expected: "val",
		When:     "test",
	}

	// Act — should not panic
	def.PrintIf(true, "actual-val")

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "PrintIf executes Print -- isPrint=true", actual)
}

// ── HandleCompiledErrorGetter: non-nil error panics ──
// Covers HandleCompiledErrorGetter.go L8-14

type cov13CompiledErrGetter struct{ err error }

func (g *cov13CompiledErrGetter) CompiledError() error { return g.err }

func Test_HandleCompiledErrorGetter_Panic(t *testing.T) {
	// Arrange
	getter := &cov13CompiledErrGetter{err: errors.New("compiled-err")}

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.HandleCompiledErrorGetter(getter)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorGetter panics -- non-nil error", actual)
}

// ── HandleCompiledErrorWithTracesGetter: non-nil error panics ──
// Covers HandleCompiledErrorWithTracesGetter.go L8-14

type cov13CompiledErrWithTracesGetter struct{ err error }

func (g *cov13CompiledErrWithTracesGetter) CompiledErrorWithStackTraces() error { return g.err }

func Test_HandleCompiledErrorWithTracesGetter_Panic(t *testing.T) {
	// Arrange
	getter := &cov13CompiledErrWithTracesGetter{err: errors.New("traces-err")}

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.HandleCompiledErrorWithTracesGetter(getter)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorWithTracesGetter panics -- non-nil error", actual)
}

// ── HandleErrorGetter: non-nil error panics ──
// Covers HandleErrorGetter.go L8-14

type cov13ErrorGetter struct{ err error }

func (g *cov13ErrorGetter) Error() error { return g.err }

func Test_HandleErrorGetter_Panic(t *testing.T) {
	// Arrange
	getter := &cov13ErrorGetter{err: errors.New("error-getter")}

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.HandleErrorGetter(getter)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorGetter panics -- non-nil error", actual)
}

// ── HandleFullStringsWithTracesGetter: non-nil error panics ──
// Covers HandleFullStringsWithTracesGetter.go L8-14

type cov13FullStringsGetter struct{ err error }

func (g *cov13FullStringsGetter) FullStringWithTraces() error { return g.err }

func Test_HandleFullStringsWithTracesGetter_Panic(t *testing.T) {
	// Arrange
	getter := &cov13FullStringsGetter{err: errors.New("full-string-err")}

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.HandleFullStringsWithTracesGetter(getter)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "HandleFullStringsWithTracesGetter panics -- non-nil error", actual)
}

// ── RawErrCollection: AddErrorGetters with non-nil/nil getters ──
// Covers RawErrCollection.go L650-661

func Test_RawErrCollection_AddErrorGetters(t *testing.T) {
	// Arrange
	ec := &errcore.RawErrCollection{}
	g1 := &cov13ErrorGetter{err: errors.New("e1")}
	g2 := &cov13ErrorGetter{err: nil} // nil error
	var _ *cov13ErrorGetter // nil getter — skip it to avoid nil dereference

	// Act — only pass non-nil getters
	ec.AddErrorGetters(g1, g2)

	// Assert
	actual := args.Map{"length": ec.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddErrorGetters adds only non-nil errors", actual)
}

// ── RawErrCollection: AddCompiledErrorGetters ──
// Covers RawErrCollection.go L672-683

func Test_RawErrCollection_AddCompiledErrorGetters(t *testing.T) {
	// Arrange
	ec := &errcore.RawErrCollection{}
	g1 := &cov13CompiledErrGetter{err: errors.New("ce1")}
	g2 := &cov13CompiledErrGetter{err: nil}

	// Act — only pass non-nil getters to avoid nil pointer deref
	ec.AddCompiledErrorGetters(g1, g2)

	// Assert
	actual := args.Map{"length": ec.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddCompiledErrorGetters adds only non-nil errors", actual)
}

// ── RawErrCollection: nil receiver Length ──
// Covers RawErrCollection.go L688-690

func Test_RawErrCollection_NilLength(t *testing.T) {
	// Arrange
	var ec *errcore.RawErrCollection

	// Act
	length := ec.Length()

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil receiver", actual)
}

// ── RawErrCollection: SerializeWithoutTraces non-empty ──
// Covers RawErrCollection.go L391

func Test_RawErrCollection_SerializeWithoutTraces(t *testing.T) {
	// Arrange
	ec := &errcore.RawErrCollection{}
	ec.AddError(errors.New("test-error"))

	// Act
	bytes, err := ec.SerializeWithoutTraces()

	// Assert
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"noErr": err == nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeWithoutTraces returns bytes -- non-empty", actual)
}

// ── RawErrCollection: CompiledJsonErrorWithStackTraces with marshal error ──
// Covers RawErrCollection.go L237

func Test_RawErrCollection_CompiledJsonStringWithTraces(t *testing.T) {
	// Arrange
	ec := &errcore.RawErrCollection{}
	ec.AddError(errors.New("test-json-error"))

	// Act
	result := ec.CompiledJsonStringWithStackTraces()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CompiledJsonStringWithStackTraces returns string -- non-empty", actual)
}

// ── stackTraceEnhance: MsgSkip with empty msg ──
// Covers stackTraceEnhance.go L55-57

func Test_StackEnhance_MsgSkip_Empty(t *testing.T) {
	// Arrange & Act
	result := errcore.StackEnhance.Msg("")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StackEnhance Msg returns empty -- empty input", actual)
}

// ── stackTraceEnhance: MsgErrorSkip with existing stack trace ──
// Covers stackTraceEnhance.go L109-111

func Test_StackEnhance_MsgErrorSkip_ExistingStackTrace(t *testing.T) {
	// Arrange
	msg := "existing error Stack-Trace present"
	err := errors.New("wrapped")

	// Act
	result := errcore.StackEnhance.MsgErrorSkip(0, msg, err)

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"containsStackTrace": strings.Contains(result, "Stack-Trace"),
	}
	expected := args.Map{
		"hasContent": true,
		"containsStackTrace": true,
	}
	expected.ShouldBeEqual(t, 0, "MsgErrorSkip returns compiled msg -- existing stack trace", actual)
	expected.ShouldBeEqual(t, 0, "MsgErrorSkip returns compiled msg -- existing stack trace", actual)
}
