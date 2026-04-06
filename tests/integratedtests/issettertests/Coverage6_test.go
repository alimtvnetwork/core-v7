package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ── RangesByte panics ──

func Test_Cov6_RangesByte_Panics(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.RangesByte()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RangesByte panics -- by design", actual)
}

// ── PanicOnOutOfRange ──

func Test_Cov6_PanicOnOutOfRange_InRange(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.PanicOnOutOfRange(1, "out of range")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "PanicOnOutOfRange in range -- no panic", actual)
}

func Test_Cov6_PanicOnOutOfRange_OutOfRange(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.PanicOnOutOfRange(255, "out of range")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicOnOutOfRange out of range -- panic", actual)
}

// ── IsCompareResult panic on default ──

func Test_Cov6_IsCompareResult_Panic(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		issetter.True.IsCompareResult(1, 99) // invalid Compare value
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult invalid compare -- panic", actual)
}

// ── OnlySupportedErr with unsupported names ──

func Test_Cov6_OnlySupportedErr_Unsupported(t *testing.T) {
	err := issetter.True.OnlySupportedErr("True")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr not all names -- error", actual)
}

func Test_Cov6_OnlySupportedMsgErr_WithError(t *testing.T) {
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr with unsupported -- error", actual)
}

// ── NewBooleans / CombinedBooleans ──

func Test_Cov6_NewBooleans_AllTrue(t *testing.T) {
	result := issetter.NewBooleans(true, true, true)
	actual := args.Map{"isTrue": result == issetter.True}
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "NewBooleans all true -- True", actual)
}

func Test_Cov6_NewBooleans_OneFalse(t *testing.T) {
	result := issetter.NewBooleans(true, false, true)
	actual := args.Map{"isFalse": result == issetter.False}
	expected := args.Map{"isFalse": true}
	expected.ShouldBeEqual(t, 0, "NewBooleans one false -- False", actual)
}

// ── GetSet / GetSetByte / GetSetUnset ──

func Test_Cov6_GetSet(t *testing.T) {
	actual := args.Map{
		"true":  issetter.GetSet(true, issetter.True, issetter.False) == issetter.True,
		"false": issetter.GetSet(false, issetter.True, issetter.False) == issetter.False,
	}
	expected := args.Map{"true": true, "false": true}
	expected.ShouldBeEqual(t, 0, "GetSet returns correct value -- with args", actual)
}

func Test_Cov6_GetSetByte(t *testing.T) {
	actual := args.Map{
		"true":  issetter.GetSetByte(true, 1, 2) == issetter.True,
		"false": issetter.GetSetByte(false, 1, 2) == issetter.False,
	}
	expected := args.Map{"true": true, "false": true}
	expected.ShouldBeEqual(t, 0, "GetSetByte returns correct value -- with args", actual)
}

func Test_Cov6_GetSetUnset(t *testing.T) {
	actual := args.Map{
		"set":   issetter.GetSetUnset(true) == issetter.Set,
		"unset": issetter.GetSetUnset(false) == issetter.Unset,
	}
	expected := args.Map{"set": true, "unset": true}
	expected.ShouldBeEqual(t, 0, "GetSetUnset returns correct value -- with args", actual)
}

// ── GetSetterByComparing ──

func Test_Cov6_GetSetterByComparing_Match(t *testing.T) {
	result := issetter.GetSetterByComparing(issetter.True, issetter.False, "hello", "world", "hello")
	actual := args.Map{"isTrue": result == issetter.True}
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing match -- True", actual)
}

func Test_Cov6_GetSetterByComparing_NoMatch(t *testing.T) {
	result := issetter.GetSetterByComparing(issetter.True, issetter.False, "hello", "world")
	actual := args.Map{"isFalse": result == issetter.False}
	expected := args.Map{"isFalse": true}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing no match -- False", actual)
}

// ── New / NewMust / NewBool ──

func Test_Cov6_New_Valid(t *testing.T) {
	v, err := issetter.New("True")
	actual := args.Map{"noErr": err == nil, "isTrue": v == issetter.True}
	expected := args.Map{"noErr": true, "isTrue": true}
	expected.ShouldBeEqual(t, 0, "New valid -- True", actual)
}

func Test_Cov6_New_Invalid(t *testing.T) {
	_, err := issetter.New("garbage")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New invalid -- error", actual)
}

func Test_Cov6_NewBool(t *testing.T) {
	actual := args.Map{
		"true":  issetter.NewBool(true) == issetter.True,
		"false": issetter.NewBool(false) == issetter.False,
	}
	expected := args.Map{"true": true, "false": true}
	expected.ShouldBeEqual(t, 0, "NewBool returns correct value -- with args", actual)
}

func Test_Cov6_NewMust_Valid(t *testing.T) {
	result := issetter.NewMust("True")
	actual := args.Map{"isTrue": result == issetter.True}
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "NewMust valid -- True", actual)
}

// ── Min / Max / MinByte / MaxByte / RangeNamesCsv ──

func Test_Cov6_MinMax(t *testing.T) {
	actual := args.Map{
		"min":     issetter.Min() == issetter.Uninitialized,
		"max":     issetter.Max() == issetter.Wildcard,
		"minByte": int(issetter.MinByte()),
		"maxByte": int(issetter.MaxByte()),
		"csv":     issetter.RangeNamesCsv() != "",
	}
	expected := args.Map{
		"min": true, "max": true,
		"minByte": 0, "maxByte": 4,
		"csv": true,
	}
	expected.ShouldBeEqual(t, 0, "Min/Max/MinByte/MaxByte/RangeNamesCsv returns correct value -- with args", actual)
}

// ── IsUnSetOrUninitialized ──

func Test_Cov6_IsUnSetOrUninitialized(t *testing.T) {
	actual := args.Map{
		"uninit": issetter.Uninitialized.IsUnSetOrUninitialized(),
		"unset":  issetter.Unset.IsUnSetOrUninitialized(),
		"true":   issetter.True.IsUnSetOrUninitialized(),
	}
	expected := args.Map{"uninit": true, "unset": true, "true": false}
	expected.ShouldBeEqual(t, 0, "IsUnSetOrUninitialized returns correct value -- with args", actual)
}

// ── GetBool ──

func Test_Cov6_GetBool(t *testing.T) {
	actual := args.Map{
		"true":  issetter.GetBool(true) == issetter.True,
		"false": issetter.GetBool(false) == issetter.False,
	}
	expected := args.Map{"true": true, "false": true}
	expected.ShouldBeEqual(t, 0, "GetBool returns correct value -- with args", actual)
}

// ── YesNoMappedValue -- uninit ──

func Test_Cov6_YesNoMappedValue_Uninit(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.YesNoMappedValue()}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "YesNoMappedValue uninit -- empty", actual)
}

// ── Comparison methods (byte) additional values ──

func Test_Cov6_ByteComparisons(t *testing.T) {
	v := issetter.True // value 1
	actual := args.Map{
		"equal":      v.IsEqual(1),
		"greater":    v.IsGreater(0),
		"greaterEq":  v.IsGreaterEqual(1),
		"less":       v.IsLess(2),
		"lessEq":     v.IsLessEqual(1),
		"between":    v.IsBetween(0, 5),
		"notBetween": v.IsBetween(2, 5),
	}
	expected := args.Map{
		"equal": true, "greater": true, "greaterEq": true,
		"less": true, "lessEq": true,
		"between": true, "notBetween": false,
	}
	expected.ShouldBeEqual(t, 0, "Value byte comparisons -- True(1)", actual)
}
