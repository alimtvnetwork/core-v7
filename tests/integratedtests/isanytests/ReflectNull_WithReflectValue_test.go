package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── ReflectNull — reflect.Value input ──

func Test_ReflectNull_WithReflectValue(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)

	// Act
	actual := args.Map{"result": isany.ReflectNull(rv)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull with reflect.Value -- false", actual)
}

func Test_ReflectNull_WithInvalidReflectValue(t *testing.T) {
	// Arrange
	rv := reflect.Value{}

	// Act
	actual := args.Map{"result": isany.ReflectNull(rv)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull with invalid reflect.Value -- true", actual)
}

func Test_ReflectNull_NilMap_FromReflectNullWithRefle(t *testing.T) {
	// Arrange
	var m map[string]string

	// Act
	actual := args.Map{"result": isany.ReflectNull(m)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull nil map -- true", actual)
}

func Test_ReflectNull_NonNilValue(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNull(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull non-nil value -- false", actual)
}

// ── ReflectValueNull ──

func Test_ReflectValueNull_InvalidValue(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectValueNull(reflect.Value{})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull invalid -- true", actual)
}

func Test_ReflectValueNull_NilSlice_FromReflectNullWithRefle(t *testing.T) {
	// Arrange
	var s []int

	// Act
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(s))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull nil slice -- true", actual)
}

func Test_ReflectValueNull_NonNilInt(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(42))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull non-nil int -- false", actual)
}

// ── NullBoth ──

func Test_NullBoth_BothNil_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NullBoth both nil -- true", actual)
}

func Test_NullBoth_OneDefined(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth(nil, "x")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth one defined -- false", actual)
}

func Test_NullBoth_BothDefined_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth("a", "b")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth both defined -- false", actual)
}

// ── DefinedBoth ──

func Test_DefinedBoth_BothDefined_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedBoth both defined -- true", actual)
}

func Test_DefinedBoth_OneNil_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth("a", nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth one nil -- false", actual)
}

func Test_DefinedBoth_BothNil_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth(nil, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth both nil -- false", actual)
}

// ── DefinedAllOf ──

func Test_DefinedAllOf_Empty_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAllOf()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf empty -- false", actual)
}

func Test_DefinedAllOf_AllDefined_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAllOf("a", 1, true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf all defined -- true", actual)
}

func Test_DefinedAllOf_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAllOf("a", nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf one nil -- false", actual)
}

// ── DefinedAnyOf ──

func Test_DefinedAnyOf_Empty_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAnyOf()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf empty -- false", actual)
}

func Test_DefinedAnyOf_OneDefined(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAnyOf(nil, "a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf one defined -- true", actual)
}

func Test_DefinedAnyOf_AllNil_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAnyOf(nil, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf all nil -- false", actual)
}

// ── AllNull ──

func Test_AllNull_AllNil_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull all nil -- true", actual)
}

func Test_AllNull_OneDefined(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(nil, "a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull one defined -- false", actual)
}

// ── AnyNull ──

func Test_AnyNull_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull("a", nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull one nil -- true", actual)
}

func Test_AnyNull_NoneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull("a", "b")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull none nil -- false", actual)
}

// ── NotDeepEqual ──

func Test_NotDeepEqual_FromReflectNullWithRefle(t *testing.T) {
	// Act
	actual := args.Map{
		"diff": isany.NotDeepEqual(1, 2),
		"same": isany.NotDeepEqual(1, 1),
	}

	// Assert
	expected := args.Map{
		"diff": true,
		"same": false,
	}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual returns correct value -- with args", actual)
}

// ── NullLeftRight ──

func Test_NullLeftRight_BothNil(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns nil -- both nil", actual)
}

// ── Null with typed nil variants ──

func Test_Null_NilChannel(t *testing.T) {
	// Arrange
	var ch chan int

	// Act
	actual := args.Map{"result": isany.Null(ch)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil channel -- true", actual)
}

func Test_Null_NilFunc_FromReflectNullWithRefle(t *testing.T) {
	// Arrange
	var fn func()

	// Act
	actual := args.Map{"result": isany.Null(fn)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil func -- true", actual)
}
