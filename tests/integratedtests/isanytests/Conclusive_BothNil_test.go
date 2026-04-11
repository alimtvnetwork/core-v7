package isanytests

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── Conclusive all branches ──

func Test_Conclusive_BothNil_FromConclusiveBothNil(t *testing.T) {
	// Arrange
	isEqual, isConcl := isany.Conclusive(nil, nil)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": true,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_BothNil returns nil -- with args", actual)
}

func Test_Conclusive_SamePointer(t *testing.T) {
	// Arrange
	v := 42
	isEqual, isConcl := isany.Conclusive(&v, &v)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": true,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_SamePointer returns correct value -- with args", actual)
}

func Test_Conclusive_LeftNil_FromConclusiveBothNil(t *testing.T) {
	// Arrange
	isEqual, isConcl := isany.Conclusive(nil, 42)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": false,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_LeftNil returns nil -- with args", actual)
}

func Test_Conclusive_RightNil_FromConclusiveBothNil(t *testing.T) {
	// Arrange
	isEqual, isConcl := isany.Conclusive(42, nil)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": false,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_RightNil returns nil -- with args", actual)
}

func Test_Conclusive_BothTypedNilSameType(t *testing.T) {
	// Arrange
	var a, b *int
	isEqual, isConcl := isany.Conclusive(a, b)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": true,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_BothTypedNilSameType returns nil -- with args", actual)
}

func Test_Conclusive_OneTypedNilOtherNot(t *testing.T) {
	// Arrange
	var a *int
	v := 42
	isEqual, isConcl := isany.Conclusive(a, &v)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": false,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_OneTypedNilOtherNot returns nil -- with args", actual)
}

func Test_Conclusive_DiffTypes(t *testing.T) {
	// Arrange
	isEqual, isConcl := isany.Conclusive(42, "hello")

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": false,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_DiffTypes returns correct value -- with args", actual)
}

func Test_Conclusive_Inconclusive(t *testing.T) {
	// Arrange
	isEqual, isConcl := isany.Conclusive(42, 43)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": false,
		"isConcl": false,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive_Inconclusive returns correct value -- with args", actual)
}

// ── ReflectValueNull ──

func Test_ReflectValueNull_NilPtr(t *testing.T) {
	// Arrange
	var ptr *int

	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(ptr))}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilPtr returns nil -- with args", actual)
}

func Test_ReflectValueNull_NonNilPtr(t *testing.T) {
	// Arrange
	v := 42

	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(&v))}

	// Assert
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NonNilPtr returns nil -- with args", actual)
}

func Test_ReflectValueNull_NilSlice(t *testing.T) {
	// Arrange
	var s []string

	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(s))}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilSlice returns nil -- with args", actual)
}

func Test_ReflectValueNull_NilMap(t *testing.T) {
	// Arrange
	var m map[string]int

	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(m))}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilMap returns nil -- with args", actual)
}

func Test_ReflectValueNull_NilChan(t *testing.T) {
	// Arrange
	var ch chan int

	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(ch))}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilChan returns nil -- with args", actual)
}

func Test_ReflectValueNull_NilFunc(t *testing.T) {
	// Arrange
	var fn func()

	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(fn))}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilFunc returns nil -- with args", actual)
}

func Test_ReflectValueNull_NilUnsafePtr(t *testing.T) {
	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(unsafe.Pointer(nil)))}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilUnsafePtr returns nil -- with args", actual)
}

func Test_ReflectValueNull_IntKind(t *testing.T) {
	// Act
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(42))}

	// Assert
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_IntKind returns correct value -- with args", actual)
}

// ── ReflectNull extended kinds ──

func Test_ReflectNull_NilMap(t *testing.T) {
	// Arrange
	var m map[string]int

	// Act
	actual := args.Map{"isNull": isany.ReflectNull(m)}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull_NilMap returns nil -- with args", actual)
}

func Test_ReflectNull_NonNilMap(t *testing.T) {
	// Act
	actual := args.Map{"isNull": isany.ReflectNull(map[string]int{"a": 1})}

	// Assert
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull_NonNilMap returns nil -- with args", actual)
}

func Test_ReflectNull_Int(t *testing.T) {
	// Act
	actual := args.Map{"isNull": isany.ReflectNull(42)}

	// Assert
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull_Int returns correct value -- with args", actual)
}

// ── FloatingPointTypeRv ──

func Test_FloatingPointTypeRv_Float32_FromConclusiveBothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(float32(3.14)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv_Float32 returns correct value -- with args", actual)
}

func Test_FloatingPointTypeRv_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(42))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv_Int returns correct value -- with args", actual)
}

// ── NumberTypeRv ──

func Test_NumberTypeRv_Int8(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int8(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv_Int8 returns correct value -- with args", actual)
}

func Test_NumberTypeRv_Uint32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint32(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv_Uint32 returns correct value -- with args", actual)
}

func Test_NumberTypeRv_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf("hello"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv_String returns correct value -- with args", actual)
}

// ── PositiveIntegerTypeRv ──

func Test_PositiveIntegerTypeRv_Uint8(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint8(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint8 returns correct value -- with args", actual)
}

func Test_PositiveIntegerTypeRv_Uint16_FromConclusiveBothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint16(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint16 returns correct value -- with args", actual)
}

func Test_PositiveIntegerTypeRv_Uint32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint32(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint32 returns correct value -- with args", actual)
}

func Test_PositiveIntegerTypeRv_Uint64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint64(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint64 returns correct value -- with args", actual)
}

func Test_PositiveIntegerTypeRv_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(42))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Int returns correct value -- with args", actual)
}

// ── PrimitiveTypeRv ──

func Test_PrimitiveTypeRv_Uintptr(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Uintptr)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv_Uintptr returns correct value -- with args", actual)
}

func Test_PrimitiveTypeRv_Struct(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Struct)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv_Struct returns correct value -- with args", actual)
}

// ── DeepEqualAllItems edge cases ──

func Test_DeepEqualAllItems(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":      isany.DeepEqualAllItems(),
		"single":     isany.DeepEqualAllItems(42),
		"twoEqual":   isany.DeepEqualAllItems(42, 42),
		"twoDiff":    isany.DeepEqualAllItems(42, 43),
		"threeMixed": isany.DeepEqualAllItems(42, 42, 43),
		"threeEqual": isany.DeepEqualAllItems(42, 42, 42),
	}

	// Assert
	expected := args.Map{
		"empty":      true,
		"single":     true,
		"twoEqual":   true,
		"twoDiff":    false,
		"threeMixed": false,
		"threeEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- with args", actual)
}

// ── DefinedItems full coverage ──

func Test_DefinedItems_Empty_FromConclusiveBothNil(t *testing.T) {
	// Arrange
	isAll, items := isany.DefinedItems()

	// Act
	actual := args.Map{
		"isAll": isAll,
		"isNil": items == nil,
	}

	// Assert
	expected := args.Map{
		"isAll": false,
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DefinedItems_Empty returns empty -- with args", actual)
}

func Test_DefinedItems_AllDefined_FromConclusiveBothNil(t *testing.T) {
	// Arrange
	isAll, items := isany.DefinedItems(1, "hello", 3.14)

	// Act
	actual := args.Map{
		"isAll": isAll,
		"len": len(items),
	}

	// Assert
	expected := args.Map{
		"isAll": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "DefinedItems_AllDefined returns correct value -- with args", actual)
}

func Test_DefinedItems_SomeNil_FromConclusiveBothNil(t *testing.T) {
	// Arrange
	isAll, items := isany.DefinedItems(nil, 42, nil, "hello")

	// Act
	actual := args.Map{
		"isAll": isAll,
		"len": len(items),
	}

	// Assert
	expected := args.Map{
		"isAll": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DefinedItems_SomeNil returns nil -- with args", actual)
}

// ── DefinedAllOf / DefinedAnyOf ──

func Test_DefinedAllOf_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.DefinedAllOf()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf_Empty returns empty -- with args", actual)
}

func Test_DefinedAnyOf(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":      isany.DefinedAnyOf(),
		"allDefined": isany.DefinedAnyOf(42, "hello"),
		"allNil":     isany.DefinedAnyOf(nil, nil),
	}

	// Assert
	expected := args.Map{
		"empty":      false,
		"allDefined": true,
		"allNil":     false,
	}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf returns correct value -- with args", actual)
}

// ── Null / NumberType / PositiveIntegerType / FloatingPointType / PrimitiveType ──

func Test_Null_UnsafePointer(t *testing.T) {
	// Act
	actual := args.Map{"isNull": isany.Null(unsafe.Pointer(nil))}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "Null_UnsafePointer returns correct value -- with args", actual)
}

func Test_NumberType_Extended(t *testing.T) {
	// Act
	actual := args.Map{
		"uint":    isany.NumberType(uint(1)),
		"int8":    isany.NumberType(int8(1)),
		"float32": isany.NumberType(float32(1.0)),
	}

	// Assert
	expected := args.Map{
		"uint":    true,
		"int8":    true,
		"float32": true,
	}
	expected.ShouldBeEqual(t, 0, "NumberType_Extended returns correct value -- with args", actual)
}

func Test_PositiveIntegerType_Extended(t *testing.T) {
	// Act
	actual := args.Map{
		"uint8":  isany.PositiveIntegerType(uint8(1)),
		"uint16": isany.PositiveIntegerType(uint16(1)),
		"uint32": isany.PositiveIntegerType(uint32(1)),
		"uint64": isany.PositiveIntegerType(uint64(1)),
	}

	// Assert
	expected := args.Map{
		"uint8":  true,
		"uint16": true,
		"uint32": true,
		"uint64": true,
	}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType_Extended returns correct value -- with args", actual)
}

func Test_FloatingPointType_Float32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointType(float32(1.0))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointType_Float32 returns correct value -- with args", actual)
}

func Test_PrimitiveType_Extended(t *testing.T) {
	// Act
	actual := args.Map{
		"uint":  isany.PrimitiveType(uint(1)),
		"int8":  isany.PrimitiveType(int8(1)),
		"slice": isany.PrimitiveType([]int{}),
	}

	// Assert
	expected := args.Map{
		"uint":  true,
		"int8":  true,
		"slice": false,
	}
	expected.ShouldBeEqual(t, 0, "PrimitiveType_Extended returns correct value -- with args", actual)
}

// keep fmt imported for any future use
var _ = fmt.Sprintf
