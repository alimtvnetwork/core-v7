package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage15 — errcore remaining gaps
// ══════════════════════════════════════════════════════════════════════════════

// --- CompiledErrorString line 30-32: compiled==nil branch ---
// This is logically unreachable: if mainErr != nil, CompiledError returns
// either mainErr or a wrapped error, never nil. Dead code.

// --- HandleCompiledErrorGetter/HandleErrorGetter/etc line 10-12: err==nil ---

func Test_HandleCompiledErrorGetter_NilError(t *testing.T) {
	// Arrange — a getter that returns nil compiled error
	getter := &mockCompiledErrorGetter{err: nil}

	// Act & Assert — should not panic
	convey.Convey("HandleCompiledErrorGetter does not panic when error is nil", t, func() {
		convey.So(func() {
			errcore.HandleCompiledErrorGetter(getter)
		}, convey.ShouldNotPanic)
	})
}

func Test_HandleCompiledErrorWithTracesGetter_NilError(t *testing.T) {
	// Arrange
	getter := &mockCompiledErrorWithTracesGetter{err: nil}

	// Act & Assert
	convey.Convey("HandleCompiledErrorWithTracesGetter does not panic when error is nil", t, func() {
		convey.So(func() {
			errcore.HandleCompiledErrorWithTracesGetter(getter)
		}, convey.ShouldNotPanic)
	})
}

func Test_HandleErrorGetter_NilError(t *testing.T) {
	// Arrange
	getter := &mockErrorGetter{err: nil}

	// Act & Assert
	convey.Convey("HandleErrorGetter does not panic when error is nil", t, func() {
		convey.So(func() {
			errcore.HandleErrorGetter(getter)
		}, convey.ShouldNotPanic)
	})
}

func Test_HandleFullStringsWithTracesGetter_NilError(t *testing.T) {
	// Arrange
	getter := &mockFullStringWithTracesGetter{err: nil}

	// Act & Assert
	convey.Convey("HandleFullStringsWithTracesGetter does not panic when error is nil", t, func() {
		convey.So(func() {
			errcore.HandleFullStringsWithTracesGetter(getter)
		}, convey.ShouldNotPanic)
	})
}

// --- RawErrCollection gaps ---

func Test_RawErrCollection_CompiledJsonStringWithStackTraces_Empty_FromHandleCompiledErrorG(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{}

	// Act
	result := coll.CompiledJsonStringWithStackTraces()

	// Assert
	convey.Convey("CompiledJsonStringWithStackTraces returns empty for empty collection", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_RawErrCollection_AddCompiledErrorGetters_NilGetter(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{}

	// Act — pass nil getter
	coll.AddCompiledErrorGetters(nil)

	// Assert
	convey.Convey("AddCompiledErrorGetters skips nil getters", t, func() {
		convey.So(coll.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_RawErrCollection_AddCompiledErrorGetters_NoErrorGetter(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{}
	getter := &mockCompiledErrorGetter{err: nil}

	// Act
	coll.AddCompiledErrorGetters(getter)

	// Assert
	convey.Convey("AddCompiledErrorGetters skips getters that return nil error", t, func() {
		convey.So(coll.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_RawErrCollection_AddErrorGetters_NilGetter(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{}

	// Act
	coll.AddErrorGetters(nil)

	// Assert
	convey.Convey("AddErrorGetters skips nil getters", t, func() {
		convey.So(coll.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_RawErrCollection_AddErrorGetters_NoErrorGetter(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{}
	getter := &mockErrorGetter{err: nil}

	// Act
	coll.AddErrorGetters(getter)

	// Assert
	convey.Convey("AddErrorGetters skips getters returning nil error", t, func() {
		convey.So(coll.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_RawErrCollection_LogIf_False(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{}
	coll.Add(errors.New("test error"))

	// Act & Assert — LogIf(false) should NOT call LogFatal
	convey.Convey("LogIf(false) does not log/exit", t, func() {
		convey.So(func() {
			coll.LogIf(false)
		}, convey.ShouldNotPanic)
	})
}

// Coverage note: RawErrCollection.LogFatal/LogFatalWithTraces call os.Exit(1) —
// untestable without process isolation. Documented as accepted gaps.
//
// CompiledErrorString line 30-32: logically unreachable — CompiledError(nonNilErr, msg)
// never returns nil. Dead code.
//
// stackTraceEnhance.go:115 — empty trace fallback, requires runtime.Caller failure. Dead code.
