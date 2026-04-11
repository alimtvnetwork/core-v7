package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Null / NotNull / Defined ──

func Test_Null_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Null(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
	actual = args.Map{"result": isany.Null(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "42 should not be null", actual)
	actual = args.Map{"result": isany.Null("hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be null", actual)

	var s []string
	actual = args.Map{"result": isany.Null(s)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil slice should be null", actual)

	var m map[string]int
	actual = args.Map{"result": isany.Null(m)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil map should be null", actual)

	var fn func()
	actual = args.Map{"result": isany.Null(fn)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil func should be null", actual)

	var ch chan int
	actual = args.Map{"result": isany.Null(ch)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil chan should be null", actual)

	var ptr *int
	actual = args.Map{"result": isany.Null(ptr)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil pointer should be null", actual)

	val := 42
	actual = args.Map{"result": isany.Null(&val)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil pointer should not be null", actual)
}

func Test_NotNull_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotNull(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "42 should be not null", actual)
	actual = args.Map{"result": isany.NotNull(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be not null", actual)
}

func Test_Defined_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Defined(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "42 should be defined", actual)
	actual = args.Map{"result": isany.Defined(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be defined", actual)
}

// ── AllNull / AnyNull ──

func Test_AllNull_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllNull(nil, nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all nils should be AllNull", actual)
	actual = args.Map{"result": isany.AllNull(nil, 42, nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mixed should not be AllNull", actual)
	actual = args.Map{"result": isany.AllNull()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be AllNull (vacuous truth)", actual)
}

func Test_AnyNull_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyNull(nil, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be AnyNull with nil present", actual)
	actual = args.Map{"result": isany.AnyNull()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not be AnyNull", actual)
	actual = args.Map{"result": isany.AnyNull(42, "hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no nils should not be AnyNull", actual)
}

// ── Zero / AllZero / AnyZero ──

func Test_Zero_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Zero(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be zero", actual)
	actual = args.Map{"result": isany.Zero("")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty string should be zero", actual)
	actual = args.Map{"result": isany.Zero(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be zero", actual)
	actual = args.Map{"result": isany.Zero(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "42 should not be zero", actual)
}

func Test_AllZero_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero(0, "", nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all zeros should be AllZero", actual)
	actual = args.Map{"result": isany.AllZero(0, 1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mixed should not be AllZero", actual)
	actual = args.Map{"result": isany.AllZero()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be AllZero", actual)
}

func Test_AnyZero_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero(0, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be AnyZero", actual)
	actual = args.Map{"result": isany.AnyZero(42, "hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no zeros should not be AnyZero", actual)
	actual = args.Map{"result": isany.AnyZero()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be AnyZero", actual)
}

// ── DeepEqual / NotDeepEqual ──

func Test_DeepEqual_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqual(42, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "42 == 42", actual)
	actual = args.Map{"result": isany.DeepEqual(42, 43)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "42 != 43", actual)
	actual = args.Map{"result": isany.DeepEqual(nil, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil == nil", actual)
}

func Test_NotDeepEqual_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotDeepEqual(42, 43)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "42 != 43", actual)
	actual = args.Map{"result": isany.NotDeepEqual(42, 42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "42 == 42", actual)
}

// ── DeepEqualAllItems ──

func Test_DeepEqualAllItems_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(42, 42, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all 42 should be equal", actual)
	actual = args.Map{"result": isany.DeepEqualAllItems(42, 43, 42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mixed should not be equal", actual)
}

// ── DefinedBoth / NullBoth / DefinedAllOf / DefinedAnyOf ──

func Test_DefinedBoth_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedBoth(42, "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both defined should be true", actual)
	actual = args.Map{"result": isany.DefinedBoth(nil, "hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be false", actual)
	actual = args.Map{"result": isany.DefinedBoth(nil, nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be false", actual)
}

func Test_NullBoth_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NullBoth(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be true", actual)
	actual = args.Map{"result": isany.NullBoth(nil, 42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one non-nil should be false", actual)
}

func Test_DefinedAllOf_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAllOf(42, "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all defined should be true", actual)
	actual = args.Map{"result": isany.DefinedAllOf(42, nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be false", actual)
}

func Test_DefinedAnyOf_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAnyOf(nil, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "one defined should be true", actual)
	actual = args.Map{"result": isany.DefinedAnyOf(nil, nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all nil should be false", actual)
}

// ── DefinedItems ──

func Test_DefinedItems_Coverage(t *testing.T) {
	// Arrange
	_, items := isany.DefinedItems(nil, 42, nil, "hello")

	// Act
	actual := args.Map{"result": len(items) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 defined items", actual)
}

// ── DefinedLeftRight ──

func Test_DefinedLeftRight_Coverage(t *testing.T) {
	// Arrange
	leftDef, rightDef := isany.DefinedLeftRight(42, nil)

	// Act
	actual := args.Map{"result": leftDef}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "left should be defined", actual)
	actual = args.Map{"result": rightDef}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right should not be defined", actual)
}

// ── NullLeftRight ──

func Test_NullLeftRight_Coverage(t *testing.T) {
	// Arrange
	leftNull, rightNull := isany.NullLeftRight(nil, 42)

	// Act
	actual := args.Map{"result": leftNull}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "left should be null", actual)
	actual = args.Map{"result": rightNull}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right should not be null", actual)
}

// ── StringEqual ──

func Test_StringEqual_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.StringEqual(42, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same values should be string equal", actual)
	actual = args.Map{"result": isany.StringEqual(42, 43)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different values should not be string equal", actual)
}

// ── JsonEqual / JsonMismatch ──

func Test_JsonEqual_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonEqual("hello", "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same strings should be json equal", actual)
	actual = args.Map{"result": isany.JsonEqual("hello", "world")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different strings should not be json equal", actual)
	actual = args.Map{"result": isany.JsonEqual(42, 42)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same ints should be json equal", actual)
	actual = args.Map{"result": isany.JsonEqual(42, 43)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different ints should not be json equal", actual)

	type s struct{ A int }
	actual = args.Map{"result": isany.JsonEqual(s{1}, s{1})}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same structs should be json equal", actual)
	actual = args.Map{"result": isany.JsonEqual(s{1}, s{2})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different structs should not be json equal", actual)
}

func Test_JsonMismatch_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonMismatch("hello", "world")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "different should mismatch", actual)
	actual = args.Map{"result": isany.JsonMismatch("hello", "hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should not mismatch", actual)
}

// ── TypeSame ──

func Test_TypeSame_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.TypeSame(42, 43)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both int should be same type", actual)
	actual = args.Map{"result": isany.TypeSame(42, "hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "int vs string should not be same type", actual)
}

// ── Pointer / Function / FuncOnly ──

func Test_Pointer_Coverage(t *testing.T) {
	// Arrange
	val := 42

	// Act
	actual := args.Map{"result": isany.Pointer(&val)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "pointer should be pointer", actual)
	actual = args.Map{"result": isany.Pointer(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-pointer should not be pointer", actual)
}

func Test_Function_Coverage(t *testing.T) {
	// Arrange
	fn := func() {}
	isFunc, name := isany.Function(fn)

	// Act
	actual := args.Map{"result": isFunc}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "func should be function", actual)
	actual = args.Map{"result": name == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "func name should not be empty", actual)

	isFunc2, _ := isany.Function(nil)
	actual = args.Map{"result": isFunc2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be function", actual)

	isFunc3, _ := isany.Function(42)
	actual = args.Map{"result": isFunc3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "int should not be function", actual)
}

func Test_FuncOnly_Coverage(t *testing.T) {
	// Arrange
	fn := func() {}

	// Act
	actual := args.Map{"result": isany.FuncOnly(fn)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "func should return true", actual)
	actual = args.Map{"result": isany.FuncOnly(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "int should return false", actual)
}

// ── PrimitiveType / NumberType / FloatingPointType / PositiveIntegerType ──

func Test_PrimitiveType_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveType(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "int should be primitive", actual)
	actual = args.Map{"result": isany.PrimitiveType("hello")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "string should be primitive", actual)
	actual = args.Map{"result": isany.PrimitiveType(true)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "bool should be primitive", actual)
	actual = args.Map{"result": isany.PrimitiveType(3.14)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "float should be primitive", actual)

	type s struct{}
	actual = args.Map{"result": isany.PrimitiveType(s{})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "struct should not be primitive", actual)
}

func Test_NumberType_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "int should be number", actual)
	actual = args.Map{"result": isany.NumberType(3.14)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "float should be number", actual)
	actual = args.Map{"result": isany.NumberType("hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be number", actual)
}

func Test_FloatingPointType_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointType(3.14)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "float64 should be floating point", actual)
	actual = args.Map{"result": isany.FloatingPointType(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "int should not be floating point", actual)
}

func Test_PositiveIntegerType_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerType(uint(42))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "uint should be positive integer", actual)
	actual = args.Map{"result": isany.PositiveIntegerType(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "signed int should not be positive integer type", actual)
}

// ── Conclusive ──

func Test_Conclusive_Coverage(t *testing.T) {
	// Arrange — tests with two ints
	_, isConclusive := isany.Conclusive(42, 43)
	actual := args.Map{"result": isConclusive}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "two same-type non-nil ints should be inconclusive", actual)
}

// ── ReflectNull / ReflectNotNull / ReflectValueNull ──

func Test_ReflectNull_Coverage(t *testing.T) {
	// Arrange
	var ptr *int

	// Act
	actual := args.Map{"result": isany.ReflectNull(ptr)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil ptr should be reflect null", actual)

	val := 42
	actual = args.Map{"result": isany.ReflectNull(&val)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil ptr should not be reflect null", actual)
}

func Test_ReflectNotNull_Coverage(t *testing.T) {
	// Arrange
	val := 42

	// Act
	actual := args.Map{"result": isany.ReflectNotNull(&val)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "non-nil should be reflect not null", actual)
}
