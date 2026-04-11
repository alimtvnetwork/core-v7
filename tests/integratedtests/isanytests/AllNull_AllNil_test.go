package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ============================================================================
// AllNull
// ============================================================================

func Test_AllNull_AllNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns nil -- all nil", actual)
}

func Test_AllNull_Mixed(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(nil, "hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull returns correct value -- mixed", actual)
}

func Test_AllNull_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns empty -- empty", actual)
}

// ============================================================================
// AnyNull
// ============================================================================

func Test_AnyNull_HasNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull("a", nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- has nil", actual)
}

func Test_AnyNull_NoNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull("a", "b")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- no nil", actual)
}

func Test_AnyNull_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns empty -- empty", actual)
}

// ============================================================================
// AllZero
// ============================================================================

func Test_AllZero_AllZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero(0, "", false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- all zero", actual)
}

func Test_AllZero_Mixed(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero(0, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- mixed", actual)
}

func Test_AllZero_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns empty -- empty", actual)
}

// ============================================================================
// AnyZero
// ============================================================================

func Test_AnyZero_HasZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero(1, 0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero returns correct value -- has zero", actual)
}

func Test_AnyZero_NoZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero(1, 2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero returns empty -- no zero", actual)
}

func Test_AnyZero_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero returns empty -- empty", actual)
}

// ============================================================================
// DefinedBoth
// ============================================================================

func Test_DefinedBoth_BothDefined(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns correct value -- both defined", actual)
}

func Test_DefinedBoth_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth("a", nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- one nil", actual)
}

func Test_DefinedBoth_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth(nil, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- both nil", actual)
}

// ============================================================================
// NotDeepEqual
// ============================================================================

func Test_NotDeepEqual_Same(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotDeepEqual(42, 42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual returns correct value -- same", actual)
}

func Test_NotDeepEqual_Diff(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotDeepEqual(42, 43)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual returns correct value -- diff", actual)
}

// ============================================================================
// DeepEqualAllItems
// ============================================================================

func Test_DeepEqualAllItems_AllSame(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- all same", actual)
}

func Test_DeepEqualAllItems_OneDiff(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- one diff", actual)
}

func Test_DeepEqualAllItems_Single(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- single", actual)
}

func Test_DeepEqualAllItems_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns empty -- empty", actual)
}

func Test_DeepEqualAllItems_Two(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems("a", "a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- two same", actual)
}

// ============================================================================
// NullBoth
// ============================================================================

func Test_NullBoth_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- both nil", actual)
}

func Test_NullBoth_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth(nil, "a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- one nil", actual)
}

// ============================================================================
// NullLeftRight
// ============================================================================

func Test_NullLeftRight(t *testing.T) {
	// Arrange
	l, r := isany.NullLeftRight(nil, "a")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": true,
		"right": false,
	}
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns nil -- nil and string", actual)
}

// ============================================================================
// NotNull
// ============================================================================

func Test_NotNull_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotNull(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotNull returns nil -- nil", actual)
}

func Test_NotNull_Value(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotNull("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotNull returns correct value -- value", actual)
}

// ============================================================================
// StringEqual
// ============================================================================

func Test_StringEqual_Same(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.StringEqual(42, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual returns non-empty -- same values", actual)
}

func Test_StringEqual_Diff(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.StringEqual(42, 43)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringEqual returns non-empty -- diff values", actual)
}

// ============================================================================
// Function
// ============================================================================

func Test_Function_Func(t *testing.T) {
	// Arrange
	isFunc, name := isany.Function(func() {})

	// Act
	actual := args.Map{
		"isFunc": isFunc,
		"hasName": name != "",
	}

	// Assert
	expected := args.Map{
		"isFunc": true,
		"hasName": true,
	}
	expected.ShouldBeEqual(t, 0, "Function returns correct value -- func", actual)
}

func Test_Function_NotFunc(t *testing.T) {
	// Arrange
	isFunc, name := isany.Function("hello")

	// Act
	actual := args.Map{
		"isFunc": isFunc,
		"name": name,
	}

	// Assert
	expected := args.Map{
		"isFunc": false,
		"name": "",
	}
	expected.ShouldBeEqual(t, 0, "Function returns correct value -- not func", actual)
}

func Test_Function_Nil(t *testing.T) {
	// Arrange
	isFunc, name := isany.Function(nil)

	// Act
	actual := args.Map{
		"isFunc": isFunc,
		"name": name,
	}

	// Assert
	expected := args.Map{
		"isFunc": false,
		"name": "",
	}
	expected.ShouldBeEqual(t, 0, "Function returns nil -- nil", actual)
}

// ============================================================================
// Pointer
// ============================================================================

func Test_Pointer_Ptr(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	actual := args.Map{"result": isany.Pointer(&s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Pointer returns correct value -- ptr", actual)
}

func Test_Pointer_NotPtr(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Pointer("hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer returns correct value -- not ptr", actual)
}

// ============================================================================
// Null — typed nil channels, maps, slices
// ============================================================================

func Test_Null_NilSlice(t *testing.T) {
	// Arrange
	var s []string

	// Act
	actual := args.Map{"result": isany.Null(s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil slice", actual)
}

func Test_Null_NilMap(t *testing.T) {
	// Arrange
	var m map[string]int

	// Act
	actual := args.Map{"result": isany.Null(m)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil map", actual)
}

func Test_Null_NilFunc(t *testing.T) {
	// Arrange
	var f func()

	// Act
	actual := args.Map{"result": isany.Null(f)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil func", actual)
}

func Test_Null_NonNilSlice(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Null([]string{})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- non-nil slice", actual)
}
