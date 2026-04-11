package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage17 — errcore remaining 11 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── CompiledError: nil mainErr (line 30-32) ──

func Test_CompiledError_NilMainErr(t *testing.T) {
	// Arrange
	var nilErr error

	// Act
	compiled := errcore.CompiledError(nilErr, "extra message")

	// Assert
	actual := args.Map{"isNil": compiled == nil}
	expected := args.Map{"isNil": true}
	actual.ShouldBeEqual(t, 1, "CompiledError nil main error", expected)
}

// ── RawErrCollection.CompiledJsonErrorWithStackTraces (line 237) ──

func Test_RawErrCollection_CompiledJsonErrorWithStackTraces(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("err1"), errors.New("err2")},
	}

	// Act
	compiledErr := coll.CompiledJsonErrorWithStackTraces()

	// Assert
	actual := args.Map{"hasError": compiledErr != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "RawErrCollection CompiledJsonErrorWithStackTraces", expected)
}

// ── RawErrCollection.CompiledJsonStringWithStackTraces (line 243-245) ──

func Test_RawErrCollection_CompiledJsonStringWithStackTraces(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("err1")},
	}

	// Act
	result := coll.CompiledJsonStringWithStackTraces()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "RawErrCollection CompiledJsonStringWithStackTraces", expected)
}

// ── RawErrCollection.SerializeWithoutTraces (line 449-455) ──

func Test_RawErrCollection_SerializeWithoutTraces_FromCompiledErrorNilMain(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("test error")},
	}

	// Act
	rawBytes, err := coll.SerializeWithoutTraces()

	// Assert
	actual := args.Map{
		"hasBytes": len(rawBytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "RawErrCollection SerializeWithoutTraces", expected)
}

// ── RawErrCollection.Serialize (line 458-464) ──

func Test_RawErrCollection_Serialize(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("test error")},
	}

	// Act
	rawBytes, err := coll.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(rawBytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "RawErrCollection Serialize", expected)
}

// ── RawErrCollection.SerializeMust panic path (line 468-470) ──

func Test_RawErrCollection_SerializeMust(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("test")},
	}

	// Act
	rawBytes := coll.SerializeMust()

	// Assert
	actual := args.Map{"hasBytes": len(rawBytes) > 0}
	expected := args.Map{"hasBytes": true}
	actual.ShouldBeEqual(t, 1, "RawErrCollection SerializeMust", expected)
}

// ── stackTraceEnhance: message without stack trace (line 115-121) ──
// This is an internal unexported method — coverage is achieved indirectly
// through ErrorWithCompiledTraceRef and similar public functions.
