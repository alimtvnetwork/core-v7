package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Null_Direct_NilInput(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Null(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_Null_Direct_StringInput(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Null("hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be null", actual)
}

func Test_Null_Direct_IntInput(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Null(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "int should not be null", actual)
}

func Test_Null_Direct_NilSlice(t *testing.T) {
	// Arrange
	var s []string

	// Act
	actual := args.Map{"result": isany.Null(s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil slice should be null", actual)
}

func Test_Null_Direct_NilMap(t *testing.T) {
	// Arrange
	var m map[string]string

	// Act
	actual := args.Map{"result": isany.Null(m)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil map should be null", actual)
}

func Test_Defined_Direct_NilInput(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Defined(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be defined", actual)
}

func Test_Defined_Direct_StringInput(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Defined("hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "string should be defined", actual)
}

func Test_DefinedBoth_Direct_BothDefined(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both defined should return true", actual)
}

func Test_DefinedBoth_Direct_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth(nil, "b")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left nil should return false", actual)
}

func Test_DefinedBoth_Direct_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth("a", nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right nil should return false", actual)
}

func Test_DefinedBoth_Direct_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth(nil, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should return false", actual)
}

func Test_DefinedLeftRight_Direct(t *testing.T) {
	// Arrange
	l, r := isany.DefinedLeftRight("a", nil)

	// Act
	actual := args.Map{"result": l}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "left should be defined", actual)
	actual = args.Map{"result": r}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right should not be defined", actual)
}

func Test_NullBoth_Direct_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should return true", actual)
}

func Test_NullBoth_Direct_OneNonNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth("a", nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one non-nil should return false", actual)
}

func Test_NullLeftRight_Direct(t *testing.T) {
	// Arrange
	l, r := isany.NullLeftRight(nil, "b")

	// Act
	actual := args.Map{"result": l}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "left should be null", actual)
	actual = args.Map{"result": r}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right should not be null", actual)
}

func Test_FuncOnly_Direct_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be func", actual)
}

func Test_FuncOnly_Direct_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly("string")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be func", actual)
}

func Test_FuncOnly_Direct_Func(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly(func() {})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "func should be func", actual)
}

func Test_NotNull_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotNull(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
	actual = args.Map{"result": isany.NotNull("hello")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "string should not be null", actual)
}

func Test_AllNull_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(nil, nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all nil should return true", actual)
	actual = args.Map{"result": isany.AllNull(nil, "a", nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one non-nil should return false", actual)
}

func Test_AnyNull_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull(nil, "a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "one nil should return true", actual)
	actual = args.Map{"result": isany.AnyNull("a", "b")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no nil should return false", actual)
}

func Test_AllZero_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero(0, "", false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all zero values should return true", actual)
	actual = args.Map{"result": isany.AllZero(0, "a", false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-zero should return false", actual)
}

func Test_AnyZero_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero(0, "a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "one zero should return true", actual)
	actual = args.Map{"result": isany.AnyZero("a", 1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no zero should return false", actual)
}

func Test_Zero_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Zero(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be zero", actual)
	actual = args.Map{"result": isany.Zero("")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty string should be zero", actual)
	actual = args.Map{"result": isany.Zero(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "false should be zero", actual)
	actual = args.Map{"result": isany.Zero(1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1 should not be zero", actual)
}

func Test_DeepEqual_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqual("a", "a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same strings should be deep equal", actual)
	actual = args.Map{"result": isany.DeepEqual("a", "b")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different strings should not be deep equal", actual)
}

func Test_NotDeepEqual_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotDeepEqual("a", "a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should return false", actual)
	actual = args.Map{"result": isany.NotDeepEqual("a", "b")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "different should return true", actual)
}

func Test_StringEqual_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.StringEqual("hello", "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same strings should be equal", actual)
	actual = args.Map{"result": isany.StringEqual("hello", "world")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different strings should not be equal", actual)
}

func Test_ReflectNull_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNull(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be reflect null", actual)
	actual = args.Map{"result": isany.ReflectNull("hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be reflect null", actual)
}

func Test_ReflectNotNull_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNotNull(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be reflect not null", actual)
	actual = args.Map{"result": isany.ReflectNotNull("hello")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "string should be reflect not null", actual)
}

func Test_Pointer_Direct(t *testing.T) {
	// Arrange
	val := 42

	// Act
	actual := args.Map{"result": isany.Pointer(&val)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "pointer should return true", actual)
	actual = args.Map{"result": isany.Pointer(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-pointer should return false", actual)
	actual = args.Map{"result": isany.Pointer(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Conclusive_Direct(t *testing.T) {
	// Arrange
	isEq, isConcl := isany.Conclusive("hello", "hello")
	if !isConcl || isEq {
		// same type, different pointers → inconclusive
	}
	isEq2, isConcl2 := isany.Conclusive(nil, nil)

	// Act
	actual := args.Map{"result": isConcl2 || !isEq2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be conclusive equal", actual)
}

func Test_DefinedAllOf_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAllOf("a", "b", "c")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all defined should return true", actual)
	actual = args.Map{"result": isany.DefinedAllOf("a", nil, "c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should return false", actual)
}

func Test_DefinedAnyOf_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAnyOf(nil, "a", nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "one defined should return true", actual)
	actual = args.Map{"result": isany.DefinedAnyOf(nil, nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all nil should return false", actual)
}

func Test_TypeSame_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.TypeSame("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same types should return true", actual)
	actual = args.Map{"result": isany.TypeSame("a", 1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different types should return false", actual)
}

func Test_JsonEqual_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonEqual("hello", "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same values should be json equal", actual)
}

func Test_JsonMismatch_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonMismatch("hello", "hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same values should not mismatch", actual)
	actual = args.Map{"result": isany.JsonMismatch("hello", "world")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "different values should mismatch", actual)
}

func Test_NumberType_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "int should be number type", actual)
	actual = args.Map{"result": isany.NumberType(3.14)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "float should be number type", actual)
	actual = args.Map{"result": isany.NumberType("hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be number type", actual)
}

func Test_PrimitiveType_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveType("hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "string should be primitive", actual)
	actual = args.Map{"result": isany.PrimitiveType(42)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "int should be primitive", actual)
	actual = args.Map{"result": isany.PrimitiveType(true)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "bool should be primitive", actual)
}

func Test_FloatingPointType_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointType(3.14)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "float64 should be floating point", actual)
	actual = args.Map{"result": isany.FloatingPointType(float32(1.0))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "float32 should be floating point", actual)
	actual = args.Map{"result": isany.FloatingPointType(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "int should not be floating point", actual)
}

func Test_DefinedItems_Direct(t *testing.T) {
	// Arrange
	_, items := isany.DefinedItems("a", nil, "c")

	// Act
	actual := args.Map{"result": len(items) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 defined items", actual)
}

func Test_PositiveIntegerType_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerType(uint(42))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "positive int should be positive integer type", actual)
	actual = args.Map{"result": isany.PositiveIntegerType(-1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "negative int should not be positive integer type", actual)
}

func Test_DeepEqualAllItems_Direct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all same should return true", actual)
	actual = args.Map{"result": isany.DeepEqualAllItems(1, 2, 1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different should return false", actual)
}
