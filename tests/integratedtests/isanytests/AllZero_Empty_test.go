package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── AllZero ──

func Test_AllZero_Empty_FromAllZeroEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero empty -- true", actual)
}

func Test_AllZero_AllZeros(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero(0, "", false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero all zeros -- true", actual)
}

func Test_AllZero_OneNonZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AllZero(0, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllZero one non-zero -- false", actual)
}

// ── AnyZero ──

func Test_AnyZero_Empty_FromAllZeroEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero empty -- true", actual)
}

func Test_AnyZero_OneZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero(1, 0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero one zero -- true", actual)
}

func Test_AnyZero_NoneZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.AnyZero(1, "a", true)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero none zero -- false", actual)
}

// ── Conclusive ──

func Test_Conclusive_BothNil(t *testing.T) {
	// Arrange
	isEqual, isConclusive := isany.Conclusive(nil, nil)

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"conclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- both nil", actual)
}

func Test_Conclusive_LeftNil(t *testing.T) {
	// Arrange
	isEqual, isConclusive := isany.Conclusive(nil, "x")

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": false,
		"conclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- left nil", actual)
}

func Test_Conclusive_RightNil(t *testing.T) {
	// Arrange
	isEqual, isConclusive := isany.Conclusive("x", nil)

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": false,
		"conclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- right nil", actual)
}

func Test_Conclusive_SameRef(t *testing.T) {
	// Arrange
	s := "hello"
	isEqual, isConclusive := isany.Conclusive(s, s)

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"conclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- same ref", actual)
}

func Test_Conclusive_DiffType(t *testing.T) {
	// Arrange
	isEqual, isConclusive := isany.Conclusive(1, "1")

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": false,
		"conclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- diff type", actual)
}

func Test_Conclusive_BothNilPtr(t *testing.T) {
	// Arrange
	var p1, p2 *int
	isEqual, isConclusive := isany.Conclusive(p1, p2)

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"conclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- both nil ptr", actual)
}

func Test_Conclusive_OneNilPtr(t *testing.T) {
	// Arrange
	var p1 *int
	v := 5
	isEqual, isConclusive := isany.Conclusive(p1, &v)

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": false,
		"conclusive": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- one nil ptr", actual)
}

func Test_Conclusive_SameTypeDiffValue(t *testing.T) {
	// Arrange
	isEqual, isConclusive := isany.Conclusive(1, 2)

	// Act
	actual := args.Map{
		"equal": isEqual,
		"conclusive": isConclusive,
	}

	// Assert
	expected := args.Map{
		"equal": false,
		"conclusive": false,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive same type diff value -- inconclusive", actual)
}

// ── Defined ──

func Test_Defined_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Defined(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Defined nil -- false", actual)
}

func Test_Defined_Value(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Defined(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Defined value -- true", actual)
}

// ── DefinedItems ──

func Test_DefinedItems_Empty(t *testing.T) {
	// Arrange
	allDefined, items := isany.DefinedItems()

	// Act
	actual := args.Map{
		"allDefined": allDefined,
		"len": len(items),
	}

	// Assert
	expected := args.Map{
		"allDefined": false,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "DefinedItems returns empty -- empty", actual)
}

func Test_DefinedItems_AllDefined(t *testing.T) {
	// Arrange
	allDefined, items := isany.DefinedItems("a", 1, true)

	// Act
	actual := args.Map{
		"allDefined": allDefined,
		"len": len(items),
	}

	// Assert
	expected := args.Map{
		"allDefined": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "DefinedItems returns correct value -- all defined", actual)
}

func Test_DefinedItems_SomeNil(t *testing.T) {
	// Arrange
	allDefined, items := isany.DefinedItems("a", nil, "b")

	// Act
	actual := args.Map{
		"allDefined": allDefined,
		"len": len(items),
	}

	// Assert
	expected := args.Map{
		"allDefined": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DefinedItems returns nil -- some nil", actual)
}

// ── DefinedLeftRight ──

func Test_DefinedLeftRight_BothDefined(t *testing.T) {
	// Arrange
	l, r := isany.DefinedLeftRight("a", "b")

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

func Test_DefinedLeftRight_LeftNil(t *testing.T) {
	// Arrange
	l, r := isany.DefinedLeftRight(nil, "b")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": false,
		"right": true,
	}
	expected.ShouldBeEqual(t, 0, "DefinedLeftRight returns nil -- left nil", actual)
}

// ── NullLeftRight ──

func Test_NullLeftRight_BothDefined(t *testing.T) {
	// Arrange
	l, r := isany.NullLeftRight("a", "b")

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

// ── FloatingPointType / FloatingPointTypeRv ──

func Test_FloatingPointType_Float64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointType(3.14)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointType float64 -- true", actual)
}

func Test_FloatingPointType_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointType(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointType int -- false", actual)
}

func Test_FloatingPointTypeRv_Float32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(float32(1.0)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv float32 -- true", actual)
}

// ── NumberType / NumberTypeRv ──

func Test_NumberType_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType int -- true", actual)
}

func Test_NumberType_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType("hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberType string -- false", actual)
}

func Test_NumberTypeRv_Uint(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint(5)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv uint -- true", actual)
}

// ── PositiveIntegerType ──

func Test_PositiveIntegerType_Uint(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerType(uint(5))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType uint -- true", actual)
}

func Test_PositiveIntegerType_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerType(5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType int -- false", actual)
}

func Test_PositiveIntegerTypeRv_Uint16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint16(5)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv uint16 -- true", actual)
}

// ── PrimitiveType / PrimitiveTypeRv ──

func Test_PrimitiveType_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveType("hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveType string -- true", actual)
}

func Test_PrimitiveType_Slice(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveType([]int{1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveType slice -- false", actual)
}

func Test_PrimitiveTypeRv_Bool(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Bool)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv bool -- true", actual)
}

func Test_PrimitiveTypeRv_Map(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Map)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv map -- false", actual)
}

// ── Pointer ──

func Test_Pointer_Ptr_FromAllZeroEmpty(t *testing.T) {
	// Arrange
	v := 5

	// Act
	actual := args.Map{"result": isany.Pointer(&v)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Pointer ptr -- true", actual)
}

func Test_Pointer_NonPtr(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Pointer(5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer non-ptr -- false", actual)
}

// ── FuncOnly / Function ──

func Test_FuncOnly_Func(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly(func() {})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FuncOnly func -- true", actual)
}

func Test_FuncOnly_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncOnly nil -- false", actual)
}

func Test_FuncOnly_NonFunc(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncOnly non-func -- false", actual)
}

func Test_Function_NilFunc(t *testing.T) {
	// Arrange
	var fn func()
	isFunc, name := isany.Function(fn)

	// Act
	actual := args.Map{
		"isFunc": isFunc,
		"name": name,
	}

	// Assert
	expected := args.Map{
		"isFunc": true,
		"name": "",
	}
	expected.ShouldBeEqual(t, 0, "Function nil func -- isFunc true name empty", actual)
}

func Test_Function_ValidFunc(t *testing.T) {
	// Arrange
	isFunc, name := isany.Function(isany.Null)

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
	expected.ShouldBeEqual(t, 0, "Function valid -- isFunc true has name", actual)
}

// ── JsonEqual / JsonMismatch ──

func Test_JsonEqual_Strings(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonEqual("abc", "abc")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonEqual returns correct value -- strings match", actual)
}

func Test_JsonEqual_Ints(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonEqual(1, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonEqual returns correct value -- ints match", actual)
}

func Test_JsonEqual_Structs(t *testing.T) {
	// Arrange
	type s struct{ A int }

	// Act
	actual := args.Map{"result": isany.JsonEqual(s{1}, s{1})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonEqual returns correct value -- structs match", actual)
}

func Test_JsonEqual_Different(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonEqual(1, 2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual different -- false", actual)
}

func Test_JsonMismatch(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonMismatch(1, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonMismatch different -- true", actual)
}

// ── NotNull / ReflectNotNull ──

func Test_NotNull_Value_FromAllZeroEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NotNull(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotNull value -- true", actual)
}

func Test_ReflectNotNull_Value(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.ReflectNotNull(42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNotNull value -- true", actual)
}

// ── StringEqual ──

func Test_StringEqual_Same_FromAllZeroEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.StringEqual("a", "a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual same -- true", actual)
}

func Test_StringEqual_Diff_FromAllZeroEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.StringEqual("a", "b")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringEqual diff -- false", actual)
}

// ── TypeSame ──

func Test_TypeSame_Same(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.TypeSame(1, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TypeSame same type -- true", actual)
}

func Test_TypeSame_Diff(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.TypeSame(1, "2")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeSame diff type -- false", actual)
}

// ── DeepEqualAllItems ──

func Test_DeepEqualAllItems_Empty_FromAllZeroEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems empty -- true", actual)
}

func Test_DeepEqualAllItems_Single_FromAllZeroEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems single -- true", actual)
}

func Test_DeepEqualAllItems_TwoEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems two equal -- true", actual)
}

func Test_DeepEqualAllItems_TwoDiff(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems two diff -- false", actual)
}

func Test_DeepEqualAllItems_ThreeEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems three equal -- true", actual)
}

func Test_DeepEqualAllItems_ThreeOneDiff(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems three one diff -- false", actual)
}

// ── Zero ──

func Test_Zero_ZeroInt(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Zero(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero zero int -- true", actual)
}

func Test_Zero_NonZero(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Zero(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Zero non-zero -- false", actual)
}
