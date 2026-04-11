package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ── RangesByte panics ──

func Test_RangesByte_Panics(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.RangesByte()
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RangesByte panics -- by design", actual)
}

// ── PanicOnOutOfRange ──

func Test_PanicOnOutOfRange_InRange(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.PanicOnOutOfRange(1, "out of range")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "PanicOnOutOfRange in range -- no panic", actual)
}

func Test_PanicOnOutOfRange_OutOfRange(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.PanicOnOutOfRange(255, "out of range")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicOnOutOfRange out of range -- panic", actual)
}

// ── IsCompareResult panic on default ──

func Test_IsCompareResult_Panic(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.IsCompareResult(1, 99) // invalid Compare value
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult invalid compare -- panic", actual)
}

// ── OnlySupportedErr with unsupported names ──

func Test_OnlySupportedErr_Unsupported(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr("True")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr not all names -- error", actual)
}

func Test_OnlySupportedMsgErr_WithError(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr with unsupported -- error", actual)
}

// ── NewBooleans / CombinedBooleans ──

func Test_NewBooleans_AllTrue_FromRangesBytePanics(t *testing.T) {
	// Arrange
	result := issetter.NewBooleans(true, true, true)

	// Act
	actual := args.Map{"isTrue": result == issetter.True}

	// Assert
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "NewBooleans all true -- True", actual)
}

func Test_NewBooleans_OneFalse(t *testing.T) {
	// Arrange
	result := issetter.NewBooleans(true, false, true)

	// Act
	actual := args.Map{"isFalse": result == issetter.False}

	// Assert
	expected := args.Map{"isFalse": true}
	expected.ShouldBeEqual(t, 0, "NewBooleans one false -- False", actual)
}

// ── GetSet / GetSetByte / GetSetUnset ──

func Test_GetSet(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  issetter.GetSet(true, issetter.True, issetter.False) == issetter.True,
		"false": issetter.GetSet(false, issetter.True, issetter.False) == issetter.False,
	}

	// Assert
	expected := args.Map{
		"true": true,
		"false": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSet returns correct value -- with args", actual)
}

func Test_GetSetByte(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  issetter.GetSetByte(true, 1, 2) == issetter.True,
		"false": issetter.GetSetByte(false, 1, 2) == issetter.False,
	}

	// Assert
	expected := args.Map{
		"true": true,
		"false": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSetByte returns correct value -- with args", actual)
}

func Test_GetSetUnset(t *testing.T) {
	// Act
	actual := args.Map{
		"set":   issetter.GetSetUnset(true) == issetter.Set,
		"unset": issetter.GetSetUnset(false) == issetter.Unset,
	}

	// Assert
	expected := args.Map{
		"set": true,
		"unset": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSetUnset returns correct value -- with args", actual)
}

// ── GetSetterByComparing ──

func Test_GetSetterByComparing_Match_FromRangesBytePanics(t *testing.T) {
	// Arrange
	result := issetter.GetSetterByComparing(issetter.True, issetter.False, "hello", "world", "hello")

	// Act
	actual := args.Map{"isTrue": result == issetter.True}

	// Assert
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing match -- True", actual)
}

func Test_GetSetterByComparing_NoMatch_FromRangesBytePanics(t *testing.T) {
	// Arrange
	result := issetter.GetSetterByComparing(issetter.True, issetter.False, "hello", "world")

	// Act
	actual := args.Map{"isFalse": result == issetter.False}

	// Assert
	expected := args.Map{"isFalse": true}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing no match -- False", actual)
}

// ── New / NewMust / NewBool ──

func Test_New_Valid(t *testing.T) {
	// Arrange
	v, err := issetter.New("True")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"isTrue": v == issetter.True,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"isTrue": true,
	}
	expected.ShouldBeEqual(t, 0, "New valid -- True", actual)
}

func Test_New_Invalid(t *testing.T) {
	// Arrange
	_, err := issetter.New("garbage")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New invalid -- error", actual)
}

func Test_NewBool(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  issetter.NewBool(true) == issetter.True,
		"false": issetter.NewBool(false) == issetter.False,
	}

	// Assert
	expected := args.Map{
		"true": true,
		"false": true,
	}
	expected.ShouldBeEqual(t, 0, "NewBool returns correct value -- with args", actual)
}

func Test_NewMust_Valid(t *testing.T) {
	// Arrange
	result := issetter.NewMust("True")

	// Act
	actual := args.Map{"isTrue": result == issetter.True}

	// Assert
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "NewMust valid -- True", actual)
}

// ── Min / Max / MinByte / MaxByte / RangeNamesCsv ──

func Test_MinMax(t *testing.T) {
	// Act
	actual := args.Map{
		"min":     issetter.Min() == issetter.Uninitialized,
		"max":     issetter.Max() == issetter.Wildcard,
		"minByte": int(issetter.MinByte()),
		"maxByte": int(issetter.MaxByte()),
		"csv":     issetter.RangeNamesCsv() != "",
	}

	// Assert
	expected := args.Map{
		"min": true, "max": true,
		"minByte": 0, "maxByte": 4,
		"csv": true,
	}
	expected.ShouldBeEqual(t, 0, "Min/Max/MinByte/MaxByte/RangeNamesCsv returns correct value -- with args", actual)
}

// ── IsUnSetOrUninitialized ──

func Test_IsUnSetOrUninitialized(t *testing.T) {
	// Act
	actual := args.Map{
		"uninit": issetter.Uninitialized.IsUnSetOrUninitialized(),
		"unset":  issetter.Unset.IsUnSetOrUninitialized(),
		"true":   issetter.True.IsUnSetOrUninitialized(),
	}

	// Assert
	expected := args.Map{
		"uninit": true,
		"unset": true,
		"true": false,
	}
	expected.ShouldBeEqual(t, 0, "IsUnSetOrUninitialized returns correct value -- with args", actual)
}

// ── GetBool ──

func Test_GetBool(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  issetter.GetBool(true) == issetter.True,
		"false": issetter.GetBool(false) == issetter.False,
	}

	// Assert
	expected := args.Map{
		"true": true,
		"false": true,
	}
	expected.ShouldBeEqual(t, 0, "GetBool returns correct value -- with args", actual)
}

// ── YesNoMappedValue -- uninit ──

func Test_YesNoMappedValue_Uninit(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.YesNoMappedValue()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "YesNoMappedValue uninit -- empty", actual)
}

// ── Comparison methods (byte) additional values ──

func Test_ByteComparisons(t *testing.T) {
	// Arrange
	v := issetter.True // value 1

	// Act
	actual := args.Map{
		"equal":      v.IsEqual(1),
		"greater":    v.IsGreater(0),
		"greaterEq":  v.IsGreaterEqual(1),
		"less":       v.IsLess(2),
		"lessEq":     v.IsLessEqual(1),
		"between":    v.IsBetween(0, 5),
		"notBetween": v.IsBetween(2, 5),
	}

	// Assert
	expected := args.Map{
		"equal": true, "greater": true, "greaterEq": true,
		"less": true, "lessEq": true,
		"between": true, "notBetween": false,
	}
	expected.ShouldBeEqual(t, 0, "Value byte comparisons -- True(1)", actual)
}
