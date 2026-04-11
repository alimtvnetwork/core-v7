package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── NumberTypeRv — extended int types ──

func Test_NumberTypeRv_Int16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int16(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- int16", actual)
}

func Test_NumberTypeRv_Int32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int32(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- int32", actual)
}

func Test_NumberTypeRv_Int64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int64(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- int64", actual)
}

func Test_NumberTypeRv_Uint8(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint8(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- uint8", actual)
}

func Test_NumberTypeRv_Uint16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint16(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- uint16", actual)
}

func Test_NumberTypeRv_Uint64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint64(1)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- uint64", actual)
}

func Test_NumberTypeRv_Float32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(float32(1.0)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- float32", actual)
}

func Test_NumberTypeRv_Float64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(float64(1.0)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns correct value -- float64", actual)
}

func Test_NumberTypeRv_Bool(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(true))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv returns non-empty -- bool false", actual)
}

// ── FloatingPointTypeRv — float64 ──

func Test_FloatingPointTypeRv_Float64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(float64(1.0)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv returns correct value -- float64", actual)
}

func Test_FloatingPointTypeRv_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf("nope"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv returns correct value -- string", actual)
}

// ── PrimitiveTypeRv — more kinds ──

func Test_PrimitiveTypeRv_Int(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Int)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv returns correct value -- int", actual)
}

func Test_PrimitiveTypeRv_Bool_FromNumberTypeRvInt16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Bool)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv returns correct value -- bool", actual)
}

func Test_PrimitiveTypeRv_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.String)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv returns correct value -- string", actual)
}

func Test_PrimitiveTypeRv_Float32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Float32)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv returns correct value -- float32", actual)
}

func Test_PrimitiveTypeRv_Float64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Float64)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv returns correct value -- float64", actual)
}

func Test_PrimitiveTypeRv_Map_FromNumberTypeRvInt16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Map)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv returns correct value -- map", actual)
}

func Test_PrimitiveTypeRv_Slice(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Slice)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv returns correct value -- slice", actual)
}

// ── PositiveIntegerTypeRv — Uint ──

func Test_PositiveIntegerTypeRv_Uint(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint(42)))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv returns correct value -- uint", actual)
}

func Test_PositiveIntegerTypeRv_String(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf("nope"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv returns correct value -- string", actual)
}

// ── NumberType — int16, int32, int64, uint8/16/32/64 ──

func Test_NumberType_Int16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(int16(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- int16", actual)
}

func Test_NumberType_Int32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(int32(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- int32", actual)
}

func Test_NumberType_Int64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(int64(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- int64", actual)
}

func Test_NumberType_Uint8(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(uint8(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- uint8", actual)
}

func Test_NumberType_Uint16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(uint16(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- uint16", actual)
}

func Test_NumberType_Uint32(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(uint32(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- uint32", actual)
}

func Test_NumberType_Uint64(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(uint64(1))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- uint64", actual)
}

func Test_NumberType_Bool(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.NumberType(true)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- bool", actual)
}

// ── Pointer with nil ──

func Test_Pointer_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.Pointer(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer returns nil -- nil", actual)
}

// ── FuncOnly with nil ──

func Test_FuncOnly_Nil_FromNumberTypeRvInt16(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.FuncOnly(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncOnly returns nil -- nil", actual)
}

// ── TypeSame with nil ──

func Test_TypeSame_NilBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.TypeSame(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TypeSame returns nil -- nil nil", actual)
}

func Test_TypeSame_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.TypeSame(nil, 42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeSame returns nil -- nil vs int", actual)
}

// ── Conclusive same values ──

func Test_Conclusive_SameValues(t *testing.T) {
	// Arrange
	isEq, isConcl := isany.Conclusive(42, 42)

	// Act
	actual := args.Map{
		"isEqual": isEq,
		"isConcl": isConcl,
	}

	// Assert
	expected := args.Map{
		"isEqual": true,
		"isConcl": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns non-empty -- same int values equal", actual)
}

// ── Zero struct ──

func Test_Zero_Struct(t *testing.T) {
	// Arrange
	type s struct{}

	// Act
	actual := args.Map{"result": isany.Zero(s{})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero returns empty -- empty struct", actual)
}

// ── DeepEqual structs ──

func Test_DeepEqual_Structs(t *testing.T) {
	// Arrange
	type s struct{ A int }

	// Act
	actual := args.Map{
		"same": isany.DeepEqual(s{1}, s{1}),
		"diff": isany.DeepEqual(s{1}, s{2}),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "DeepEqual returns correct value -- structs", actual)
}
