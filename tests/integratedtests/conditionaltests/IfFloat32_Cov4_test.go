package conditionaltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// typed_float32.go — all 11 functions
// ============================================================================

func Test_IfFloat32_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  fmt.Sprintf("%.1f", conditional.IfFloat32(true, 1.5, 2.5)),
		"false": fmt.Sprintf("%.1f", conditional.IfFloat32(false, 1.5, 2.5)),
	}

	// Assert
	expected := args.Map{
		"true": "1.5",
		"false": "2.5",
	}
	expected.ShouldBeEqual(t, 0, "IfFloat32 returns correct value -- with args", actual)
}

func Test_IfFuncFloat32_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  fmt.Sprintf("%.1f", conditional.IfFuncFloat32(true, func() float32 { return 1.0 }, func() float32 { return 2.0 })),
		"false": fmt.Sprintf("%.1f", conditional.IfFuncFloat32(false, func() float32 { return 1.0 }, func() float32 { return 2.0 })),
	}

	// Assert
	expected := args.Map{
		"true": "1.0",
		"false": "2.0",
	}
	expected.ShouldBeEqual(t, 0, "IfFuncFloat32 returns correct value -- with args", actual)
}

func Test_IfTrueFuncFloat32_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  fmt.Sprintf("%.2f", conditional.IfTrueFuncFloat32(true, func() float32 { return 3.14 })),
		"false": fmt.Sprintf("%.0f", conditional.IfTrueFuncFloat32(false, func() float32 { return 3.14 })),
	}

	// Assert
	expected := args.Map{
		"true": "3.14",
		"false": "0",
	}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncFloat32 returns non-empty -- with args", actual)
}

func Test_IfSliceFloat32_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"trueLen":  len(conditional.IfSliceFloat32(true, []float32{1.0}, []float32{2.0, 3.0})),
		"falseLen": len(conditional.IfSliceFloat32(false, []float32{1.0}, []float32{2.0, 3.0})),
	}

	// Assert
	expected := args.Map{
		"trueLen": 1,
		"falseLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "IfSliceFloat32 returns correct value -- with args", actual)
}

func Test_IfPtrFloat32_Cov4(t *testing.T) {
	// Arrange
	a, b := float32(1.0), float32(2.0)

	// Act
	actual := args.Map{
		"true":  fmt.Sprintf("%.1f", *conditional.IfPtrFloat32(true, &a, &b)),
		"false": fmt.Sprintf("%.1f", *conditional.IfPtrFloat32(false, &a, &b)),
	}

	// Assert
	expected := args.Map{
		"true": "1.0",
		"false": "2.0",
	}
	expected.ShouldBeEqual(t, 0, "IfPtrFloat32 returns correct value -- with args", actual)
}

func Test_NilDefFloat32_Cov4(t *testing.T) {
	// Arrange
	v := float32(3.3)

	// Act
	actual := args.Map{
		"nil":    fmt.Sprintf("%.1f", conditional.NilDefFloat32(nil, 5.5)),
		"nonNil": fmt.Sprintf("%.1f", conditional.NilDefFloat32(&v, 5.5)),
	}

	// Assert
	expected := args.Map{
		"nil": "5.5",
		"nonNil": "3.3",
	}
	expected.ShouldBeEqual(t, 0, "NilDefFloat32 returns nil -- with args", actual)
}

func Test_NilDefPtrFloat32_Cov4(t *testing.T) {
	// Arrange
	v := float32(1.1)

	// Act
	actual := args.Map{
		"nil":    fmt.Sprintf("%.1f", *conditional.NilDefPtrFloat32(nil, 9.9)),
		"nonNil": fmt.Sprintf("%.1f", *conditional.NilDefPtrFloat32(&v, 9.9)),
	}

	// Assert
	expected := args.Map{
		"nil": "9.9",
		"nonNil": "1.1",
	}
	expected.ShouldBeEqual(t, 0, "NilDefPtrFloat32 returns nil -- with args", actual)
}

func Test_ValueOrZeroFloat32_Cov4(t *testing.T) {
	// Arrange
	v := float32(7.7)

	// Act
	actual := args.Map{
		"nil":    fmt.Sprintf("%.0f", conditional.ValueOrZeroFloat32(nil)),
		"nonNil": fmt.Sprintf("%.1f", conditional.ValueOrZeroFloat32(&v)),
	}

	// Assert
	expected := args.Map{
		"nil": "0",
		"nonNil": "7.7",
	}
	expected.ShouldBeEqual(t, 0, "ValueOrZeroFloat32 returns correct value -- with args", actual)
}

func Test_PtrOrZeroFloat32_Cov4(t *testing.T) {
	// Arrange
	v := float32(4.4)

	// Act
	actual := args.Map{
		"nil":    fmt.Sprintf("%.0f", *conditional.PtrOrZeroFloat32(nil)),
		"nonNil": fmt.Sprintf("%.1f", *conditional.PtrOrZeroFloat32(&v)),
	}

	// Assert
	expected := args.Map{
		"nil": "0",
		"nonNil": "4.4",
	}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroFloat32 returns correct value -- with args", actual)
}

func Test_NilValFloat32_Cov4(t *testing.T) {
	// Arrange
	v := float32(5.0)

	// Act
	actual := args.Map{
		"nil":    fmt.Sprintf("%.1f", conditional.NilValFloat32(nil, 1.0, 2.0)),
		"nonNil": fmt.Sprintf("%.1f", conditional.NilValFloat32(&v, 1.0, 2.0)),
	}

	// Assert
	expected := args.Map{
		"nil": "1.0",
		"nonNil": "2.0",
	}
	expected.ShouldBeEqual(t, 0, "NilValFloat32 returns nil -- with args", actual)
}

func Test_NilValPtrFloat32_Cov4(t *testing.T) {
	// Arrange
	v := float32(5.0)

	// Act
	actual := args.Map{
		"nil":    fmt.Sprintf("%.1f", *conditional.NilValPtrFloat32(nil, 1.0, 2.0)),
		"nonNil": fmt.Sprintf("%.1f", *conditional.NilValPtrFloat32(&v, 1.0, 2.0)),
	}

	// Assert
	expected := args.Map{
		"nil": "1.0",
		"nonNil": "2.0",
	}
	expected.ShouldBeEqual(t, 0, "NilValPtrFloat32 returns nil -- with args", actual)
}

// ============================================================================
// typed_float64.go
// ============================================================================

func Test_Float64_Functions_Cov4(t *testing.T) {
	// Arrange
	v := 3.3

	// Act
	actual := args.Map{
		"nilDef":         fmt.Sprintf("%.1f", conditional.NilDefFloat64(nil, 5.5)),
		"nilDefNonNil":   fmt.Sprintf("%.1f", conditional.NilDefFloat64(&v, 5.5)),
		"nilDefPtr":      fmt.Sprintf("%.1f", *conditional.NilDefPtrFloat64(nil, 9.9)),
		"valueOrZero":    fmt.Sprintf("%.0f", conditional.ValueOrZeroFloat64(nil)),
		"valueOrZeroVal": fmt.Sprintf("%.1f", conditional.ValueOrZeroFloat64(&v)),
		"ptrOrZero":      fmt.Sprintf("%.0f", *conditional.PtrOrZeroFloat64(nil)),
		"nilVal":         fmt.Sprintf("%.1f", conditional.NilValFloat64(nil, 1.0, 2.0)),
		"nilValNonNil":   fmt.Sprintf("%.1f", conditional.NilValFloat64(&v, 1.0, 2.0)),
	}

	// Assert
	expected := args.Map{
		"nilDef": "5.5", "nilDefNonNil": "3.3",
		"nilDefPtr": "9.9", "valueOrZero": "0", "valueOrZeroVal": "3.3",
		"ptrOrZero": "0", "nilVal": "1.0", "nilValNonNil": "2.0",
	}
	expected.ShouldBeEqual(t, 0, "Float64_Functions returns correct value -- with args", actual)
}

func Test_Float64_NilValPtr_Cov4(t *testing.T) {
	// Arrange
	v := 5.0

	// Act
	actual := args.Map{
		"nil":    fmt.Sprintf("%.1f", *conditional.NilValPtrFloat64(nil, 1.0, 2.0)),
		"nonNil": fmt.Sprintf("%.1f", *conditional.NilValPtrFloat64(&v, 1.0, 2.0)),
	}

	// Assert
	expected := args.Map{
		"nil": "1.0",
		"nonNil": "2.0",
	}
	expected.ShouldBeEqual(t, 0, "NilValPtrFloat64 returns nil -- with args", actual)
}

func Test_Float64_SlicePtrFunc_Cov4(t *testing.T) {
	// Arrange
	a, b := 1.0, 2.0

	// Act
	actual := args.Map{
		"sliceLen":     len(conditional.IfSliceFloat64(true, []float64{1.0}, []float64{2.0, 3.0})),
		"ptrTrue":      fmt.Sprintf("%.1f", *conditional.IfPtrFloat64(true, &a, &b)),
		"ifFunc":       fmt.Sprintf("%.1f", conditional.IfFuncFloat64(true, func() float64 { return 1.0 }, func() float64 { return 2.0 })),
		"trueFuncTrue": fmt.Sprintf("%.2f", conditional.IfTrueFuncFloat64(true, func() float64 { return 3.14 })),
		"trueFuncFalse": fmt.Sprintf("%.0f", conditional.IfTrueFuncFloat64(false, func() float64 { return 3.14 })),
	}

	// Assert
	expected := args.Map{
		"sliceLen": 1, "ptrTrue": "1.0", "ifFunc": "1.0",
		"trueFuncTrue": "3.14", "trueFuncFalse": "0",
	}
	expected.ShouldBeEqual(t, 0, "Float64_SlicePtrFunc returns correct value -- with args", actual)
}

// ============================================================================
// Helper for typed int tests — generates args.Map for all 11 functions of a type
// ============================================================================

func Test_Int8_Functions_Cov4(t *testing.T) {
	// Arrange
	v := int8(3)

	// Act
	actual := args.Map{
		"ifTrue":       int(conditional.IfInt8(true, 1, 2)),
		"ifFalse":      int(conditional.IfInt8(false, 1, 2)),
		"ifFunc":       int(conditional.IfFuncInt8(true, func() int8 { return 1 }, func() int8 { return 2 })),
		"trueFuncT":    int(conditional.IfTrueFuncInt8(true, func() int8 { return 5 })),
		"trueFuncF":    int(conditional.IfTrueFuncInt8(false, func() int8 { return 5 })),
		"sliceLen":     len(conditional.IfSliceInt8(true, []int8{1}, []int8{2, 3})),
		"ptrTrue":      int(*conditional.IfPtrInt8(true, &v, &v)),
		"nilDef":       int(conditional.NilDefInt8(nil, 5)),
		"nilDefNonNil": int(conditional.NilDefInt8(&v, 5)),
		"nilDefPtr":    int(*conditional.NilDefPtrInt8(nil, 9)),
		"valueOrZero":  int(conditional.ValueOrZeroInt8(nil)),
		"ptrOrZero":    int(*conditional.PtrOrZeroInt8(nil)),
		"nilVal":       int(conditional.NilValInt8(nil, 1, 2)),
		"nilValPtr":    int(*conditional.NilValPtrInt8(nil, 1, 2)),
	}

	// Assert
	expected := args.Map{
		"ifTrue": 1, "ifFalse": 2, "ifFunc": 1,
		"trueFuncT": 5, "trueFuncF": 0, "sliceLen": 1, "ptrTrue": 3,
		"nilDef": 5, "nilDefNonNil": 3, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Int8_Functions returns correct value -- with args", actual)
}

func Test_Int16_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfInt16(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncInt16(true, func() int16 { return 10 }, func() int16 { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncInt16(true, func() int16 { return 50 })),
		"trueFuncF":   int(conditional.IfTrueFuncInt16(false, func() int16 { return 50 })),
		"sliceLen":    len(conditional.IfSliceInt16(true, []int16{1}, []int16{2, 3})),
		"nilDef":      int(conditional.NilDefInt16(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrInt16(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroInt16(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroInt16(nil)),
		"nilVal":      int(conditional.NilValInt16(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrInt16(nil, 1, 2)),
	}

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50, "trueFuncF": 0,
		"sliceLen": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	a, b := int16(1), int16(2)
	actual["ptrTrue"] = int(*conditional.IfPtrInt16(true, &a, &b))
	expected["ptrTrue"] = 1
	expected.ShouldBeEqual(t, 0, "Int16_Functions returns correct value -- with args", actual)
}

func Test_Int32_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfInt32(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncInt32(true, func() int32 { return 10 }, func() int32 { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncInt32(true, func() int32 { return 50 })),
		"trueFuncF":   int(conditional.IfTrueFuncInt32(false, func() int32 { return 50 })),
		"sliceLen":    len(conditional.IfSliceInt32(true, []int32{1}, []int32{2, 3})),
		"nilDef":      int(conditional.NilDefInt32(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrInt32(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroInt32(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroInt32(nil)),
		"nilVal":      int(conditional.NilValInt32(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrInt32(nil, 1, 2)),
	}
	a, b := int32(1), int32(2)
	actual["ptrTrue"] = int(*conditional.IfPtrInt32(true, &a, &b))

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50, "trueFuncF": 0,
		"sliceLen": 1, "ptrTrue": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Int32_Functions returns correct value -- with args", actual)
}

func Test_Int64_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfInt64(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncInt64(true, func() int64 { return 10 }, func() int64 { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncInt64(true, func() int64 { return 50 })),
		"trueFuncF":   int(conditional.IfTrueFuncInt64(false, func() int64 { return 50 })),
		"sliceLen":    len(conditional.IfSliceInt64(true, []int64{1}, []int64{2, 3})),
		"nilDef":      int(conditional.NilDefInt64(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrInt64(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroInt64(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroInt64(nil)),
		"nilVal":      int(conditional.NilValInt64(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrInt64(nil, 1, 2)),
	}
	a, b := int64(1), int64(2)
	actual["ptrTrue"] = int(*conditional.IfPtrInt64(true, &a, &b))

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50, "trueFuncF": 0,
		"sliceLen": 1, "ptrTrue": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Int64_Functions returns correct value -- with args", actual)
}

func Test_Uint_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfUint(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncUint(true, func() uint { return 10 }, func() uint { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncUint(true, func() uint { return 50 })),
		"trueFuncF":   int(conditional.IfTrueFuncUint(false, func() uint { return 50 })),
		"sliceLen":    len(conditional.IfSliceUint(true, []uint{1}, []uint{2, 3})),
		"nilDef":      int(conditional.NilDefUint(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrUint(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroUint(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroUint(nil)),
		"nilVal":      int(conditional.NilValUint(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrUint(nil, 1, 2)),
	}
	a, b := uint(1), uint(2)
	actual["ptrTrue"] = int(*conditional.IfPtrUint(true, &a, &b))

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50, "trueFuncF": 0,
		"sliceLen": 1, "ptrTrue": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Uint_Functions returns correct value -- with args", actual)
}

func Test_Uint8_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfUint8(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncUint8(true, func() uint8 { return 10 }, func() uint8 { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncUint8(true, func() uint8 { return 50 })),
		"trueFuncF":   int(conditional.IfTrueFuncUint8(false, func() uint8 { return 50 })),
		"sliceLen":    len(conditional.IfSliceUint8(true, []uint8{1}, []uint8{2, 3})),
		"nilDef":      int(conditional.NilDefUint8(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrUint8(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroUint8(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroUint8(nil)),
		"nilVal":      int(conditional.NilValUint8(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrUint8(nil, 1, 2)),
	}
	a, b := uint8(1), uint8(2)
	actual["ptrTrue"] = int(*conditional.IfPtrUint8(true, &a, &b))

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50, "trueFuncF": 0,
		"sliceLen": 1, "ptrTrue": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Uint8_Functions returns correct value -- with args", actual)
}

func Test_Uint16_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfUint16(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncUint16(true, func() uint16 { return 10 }, func() uint16 { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncUint16(true, func() uint16 { return 50 })),
		"sliceLen":    len(conditional.IfSliceUint16(true, []uint16{1}, []uint16{2, 3})),
		"nilDef":      int(conditional.NilDefUint16(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrUint16(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroUint16(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroUint16(nil)),
		"nilVal":      int(conditional.NilValUint16(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrUint16(nil, 1, 2)),
	}
	a, b := uint16(1), uint16(2)
	actual["ptrTrue"] = int(*conditional.IfPtrUint16(true, &a, &b))

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50,
		"sliceLen": 1, "ptrTrue": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Uint16_Functions returns correct value -- with args", actual)
}

func Test_Uint32_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfUint32(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncUint32(true, func() uint32 { return 10 }, func() uint32 { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncUint32(true, func() uint32 { return 50 })),
		"sliceLen":    len(conditional.IfSliceUint32(true, []uint32{1}, []uint32{2, 3})),
		"nilDef":      int(conditional.NilDefUint32(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrUint32(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroUint32(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroUint32(nil)),
		"nilVal":      int(conditional.NilValUint32(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrUint32(nil, 1, 2)),
	}
	a, b := uint32(1), uint32(2)
	actual["ptrTrue"] = int(*conditional.IfPtrUint32(true, &a, &b))

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50,
		"sliceLen": 1, "ptrTrue": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Uint32_Functions returns correct value -- with args", actual)
}

func Test_Uint64_Functions_Cov4(t *testing.T) {
	// Act
	actual := args.Map{
		"ifTrue":      int(conditional.IfUint64(true, 10, 20)),
		"ifFunc":      int(conditional.IfFuncUint64(true, func() uint64 { return 10 }, func() uint64 { return 20 })),
		"trueFuncT":   int(conditional.IfTrueFuncUint64(true, func() uint64 { return 50 })),
		"sliceLen":    len(conditional.IfSliceUint64(true, []uint64{1}, []uint64{2, 3})),
		"nilDef":      int(conditional.NilDefUint64(nil, 5)),
		"nilDefPtr":   int(*conditional.NilDefPtrUint64(nil, 9)),
		"valueOrZero": int(conditional.ValueOrZeroUint64(nil)),
		"ptrOrZero":   int(*conditional.PtrOrZeroUint64(nil)),
		"nilVal":      int(conditional.NilValUint64(nil, 1, 2)),
		"nilValPtr":   int(*conditional.NilValPtrUint64(nil, 1, 2)),
	}
	a, b := uint64(1), uint64(2)
	actual["ptrTrue"] = int(*conditional.IfPtrUint64(true, &a, &b))

	// Assert
	expected := args.Map{
		"ifTrue": 10, "ifFunc": 10, "trueFuncT": 50,
		"sliceLen": 1, "ptrTrue": 1, "nilDef": 5, "nilDefPtr": 9,
		"valueOrZero": 0, "ptrOrZero": 0, "nilVal": 1, "nilValPtr": 1,
	}
	expected.ShouldBeEqual(t, 0, "Uint64_Functions returns correct value -- with args", actual)
}

// ============================================================================
// Bool/Byte/Int Ptr functions
// ============================================================================

func Test_NilDefPtrBool_Cov4(t *testing.T) {
	// Arrange
	v := false

	// Act
	actual := args.Map{
		"nil":    *conditional.NilDefPtrBool(nil, true),
		"nonNil": *conditional.NilDefPtrBool(&v, true),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"nonNil": false,
	}
	expected.ShouldBeEqual(t, 0, "NilDefPtrBool returns nil -- with args", actual)
}

func Test_PtrOrZeroBool_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": *conditional.PtrOrZeroBool(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroBool returns correct value -- with args", actual)
}

func Test_NilValPtrBool_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": *conditional.NilValPtrBool(nil, true, false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NilValPtrBool returns nil -- with args", actual)
}

func Test_NilDefPtrByte_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": int(*conditional.NilDefPtrByte(nil, 9))}

	// Assert
	expected := args.Map{"result": 9}
	expected.ShouldBeEqual(t, 0, "NilDefPtrByte returns nil -- with args", actual)
}

func Test_PtrOrZeroByte_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": int(*conditional.PtrOrZeroByte(nil))}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroByte returns correct value -- with args", actual)
}

func Test_NilValPtrByte_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": int(*conditional.NilValPtrByte(nil, 1, 2))}

	// Assert
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "NilValPtrByte returns nil -- with args", actual)
}

func Test_NilDefPtrInt_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": *conditional.NilDefPtrInt(nil, 42)}

	// Assert
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "NilDefPtrInt returns nil -- with args", actual)
}

func Test_PtrOrZeroInt_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": *conditional.PtrOrZeroInt(nil)}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroInt returns correct value -- with args", actual)
}

func Test_NilValPtrInt_Cov4(t *testing.T) {
	// Act
	actual := args.Map{"result": *conditional.NilValPtrInt(nil, -1, 1)}

	// Assert
	expected := args.Map{"result": -1}
	expected.ShouldBeEqual(t, 0, "NilValPtrInt returns nil -- with args", actual)
}

// ============================================================================
// VoidFunctions / ErrorFunctionsExecuteResults / FunctionsExecuteResults
// ============================================================================

func Test_VoidFunctions_WithNilFunc_Cov4(t *testing.T) {
	// Arrange
	count := 0
	funcs := []func(){nil, func() { count++ }}
	conditional.VoidFunctions(true, funcs, []func(){})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "VoidFunctions_WithNilFunc returns nil -- with args", actual)
}

func Test_ErrorFunctionsExecuteResults_WithNilFunc_Cov4(t *testing.T) {
	// Arrange
	funcs := []func() error{nil, func() error { return nil }}
	err := conditional.ErrorFunctionsExecuteResults(true, funcs, []func() error{})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults_WithNilFunc returns nil -- with args", actual)
}

func Test_FunctionsExecuteResults_SkipTake_Cov4(t *testing.T) {
	// Arrange
	funcs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "skip", false, false },
		func() (string, bool, bool) { return "take", true, false },
	}
	results := conditional.FunctionsExecuteResults[string](true, funcs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"first": results[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "take",
	}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults_SkipTake returns correct value -- with args", actual)
}

func Test_FunctionsExecuteResults_NilFunc_Cov4(t *testing.T) {
	// Arrange
	funcs := []func() (string, bool, bool){
		nil,
		func() (string, bool, bool) { return "a", true, false },
	}
	results := conditional.FunctionsExecuteResults[string](true, funcs, nil)

	// Act
	actual := args.Map{"len": len(results)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults_NilFunc returns nil -- with args", actual)
}

func Test_FunctionsExecuteResults_Empty_Cov4(t *testing.T) {
	// Arrange
	results := conditional.FunctionsExecuteResults[string](true, nil, nil)

	// Act
	actual := args.Map{"isNil": results == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults_Empty returns empty -- with args", actual)
}
