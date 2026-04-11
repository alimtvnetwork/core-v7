package isanytests

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── ReflectNull all kinds ──

func Test_ReflectNull_NilFunc(t *testing.T) {
	// Arrange
	var fn func()

	// Act
	actual := args.Map{"result": isany.ReflectNull(fn)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil func", actual)
}

func Test_ReflectNull_NilChan(t *testing.T) {
	// Arrange
	var ch chan int

	// Act
	actual := args.Map{"result": isany.ReflectNull(ch)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil chan", actual)
}

func Test_ReflectNull_NilSlice(t *testing.T) {
	// Arrange
	var s []string

	// Act
	actual := args.Map{"result": isany.ReflectNull(s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil slice", actual)
}

func Test_ReflectNull_NilUnsafePtr(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNull(unsafe.Pointer(nil))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil unsafe pointer", actual)
}

func Test_ReflectNull_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNull("hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns correct value -- string default kind", actual)
}

// ── ReflectNotNull ──

func Test_ReflectNotNull_NilPtr(t *testing.T) {
	// Arrange
	var ptr *int

	// Act
	actual := args.Map{"result": isany.ReflectNotNull(ptr)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNotNull returns nil -- nil ptr", actual)
}

// ── ReflectValueNull all kinds ──

func Test_ReflectValueNull_NilInterface(t *testing.T) {
	// Arrange
	var iface interface{}
	rv := reflect.ValueOf(&iface).Elem()

	// Act
	actual := args.Map{"result": isany.ReflectValueNull(rv)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns nil -- nil interface", actual)
}

func Test_ReflectValueNull_NonNilSlice(t *testing.T) {
	// Arrange
	s := []string{"a"}

	// Act
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(s))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns nil -- non-nil slice", actual)
}

func Test_ReflectValueNull_NonNilMap(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1}

	// Act
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(m))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns nil -- non-nil map", actual)
}

func Test_ReflectValueNull_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf("hello"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns correct value -- string kind", actual)
}

// ── AllZero / AnyZero edge cases ──

func Test_AllZero_SingleZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- single zero", actual)
}

func Test_AnyZero_SingleNonZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero returns non-empty -- single non-zero", actual)
}

// ── AllNull edge cases ──

func Test_AllNull_SingleNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns nil -- single nil", actual)
}

func Test_AllNull_SingleNonNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull returns nil -- single non-nil", actual)
}

// ── AnyNull edge cases ──

func Test_AnyNull_SingleNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- single nil", actual)
}

func Test_AnyNull_SingleNonNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- single non-nil", actual)
}

// ── Zero with various types ──

func Test_Zero_Bool(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Zero(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero returns non-empty -- false bool", actual)
}

func Test_Zero_TrueBool(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Zero(true)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Zero returns non-empty -- true bool", actual)
}

func Test_Zero_Float64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Zero(0.0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero returns correct value -- float64 0.0", actual)
}

// ── StringEqual edge cases ──

func Test_StringEqual_DiffTypes(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.StringEqual(42, "42")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual returns correct value -- different types", actual)
}

// ── Conclusive with two typed nils of different types ──

func Test_Conclusive_BothReflectable(t *testing.T) {
	// Arrange
	a := 42
	b := 42
	isEq, isConcl := isany.Conclusive(&a, &b)

	// Act
	actual := args.Map{
		"isEqual": isEq,
		"isConclusive": isConcl,
	}
	// Different pointers, same type, so inconclusive (needs deep equal)

	// Assert
	expected := args.Map{
		"isEqual": false,
		"isConclusive": false,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- diff pointers same type", actual)
}

// ── PositiveIntegerType with signed ──

func Test_PositiveIntegerType_Int64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerType(int64(42))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType returns correct value -- signed int64", actual)
}

// ── NumberType byte ──

func Test_NumberType_Byte(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(byte(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- byte", actual)
}

// ── PrimitiveType byte ──

func Test_PrimitiveType_Byte(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveType(byte(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveType returns correct value -- byte", actual)
}

// ── NullBoth edge ──

func Test_NullBoth_BothNonNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth(42, "hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- both non-nil", actual)
}

// ── DefinedBoth edge ──

func Test_DefinedBoth_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth(nil, 42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- left nil", actual)
}

// ── NullLeftRight ──

func Test_NullLeftRight_BothNull(t *testing.T) {
	// Arrange
	l, r := isany.NullLeftRight(nil, nil)

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": true,
		"right": true,
	}
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns correct value -- both null", actual)
}

func Test_NullLeftRight_BothDefined_FromReflectNullNilFunc(t *testing.T) {
	// Arrange
	l, r := isany.NullLeftRight(42, "hello")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": false,
		"right": false,
	}
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns correct value -- both defined", actual)
}

// ── DefinedLeftRight ──

func Test_DefinedLeftRight_BothDefined_FromReflectNullNilFunc(t *testing.T) {
	// Arrange
	l, r := isany.DefinedLeftRight(42, "hello")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": true,
		"right": true,
	}
	expected.ShouldBeEqual(t, 0, "DefinedLeftRight returns correct value -- both defined", actual)
}

func Test_DefinedLeftRight_BothNil(t *testing.T) {
	// Arrange
	l, r := isany.DefinedLeftRight(nil, nil)

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": false,
		"right": false,
	}
	expected.ShouldBeEqual(t, 0, "DefinedLeftRight returns nil -- both nil", actual)
}

// ── Function with method ref ──

func Test_Function_NilFunc_FromReflectNullNilFunc(t *testing.T) {
	// Arrange
	var fn func()
	isFunc, _ := isany.Function(fn)

	// Act
	actual := args.Map{"isFunc": isFunc}

	// Assert
	expected := args.Map{"isFunc": true}
	expected.ShouldBeEqual(t, 0, "Function returns nil -- nil func still detects kind", actual)
}
