package coremathtests

import (
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── MaxByte ──

func Test_MaxByte_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MaxByte(10, 5)}

	// Assert
	expected := args.Map{"result": byte(10)}
	expected.ShouldBeEqual(t, 0, "MaxByte returns correct value -- left greater", actual)
}

func Test_MaxByte_RightGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MaxByte(3, 8)}

	// Assert
	expected := args.Map{"result": byte(8)}
	expected.ShouldBeEqual(t, 0, "MaxByte returns correct value -- right greater", actual)
}

// ── MinByte ──

func Test_MinByte_LeftSmaller(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MinByte(2, 9)}

	// Assert
	expected := args.Map{"result": byte(2)}
	expected.ShouldBeEqual(t, 0, "MinByte returns correct value -- left smaller", actual)
}

func Test_MinByte_RightSmaller(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MinByte(9, 3)}

	// Assert
	expected := args.Map{"result": byte(3)}
	expected.ShouldBeEqual(t, 0, "MinByte returns correct value -- right smaller", actual)
}

// ── MaxFloat32 ──

func Test_MaxFloat32_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MaxFloat32(3.5, 1.2)}

	// Assert
	expected := args.Map{"result": float32(3.5)}
	expected.ShouldBeEqual(t, 0, "MaxFloat32 returns correct value -- left greater", actual)
}

func Test_MaxFloat32_RightGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MaxFloat32(1.0, 2.5)}

	// Assert
	expected := args.Map{"result": float32(2.5)}
	expected.ShouldBeEqual(t, 0, "MaxFloat32 returns correct value -- right greater", actual)
}

// ── MinFloat32 ──

func Test_MinFloat32_LeftSmaller(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MinFloat32(1.0, 5.0)}

	// Assert
	expected := args.Map{"result": float32(1.0)}
	expected.ShouldBeEqual(t, 0, "MinFloat32 returns correct value -- left smaller", actual)
}

func Test_MinFloat32_RightSmaller(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MinFloat32(5.0, 2.0)}

	// Assert
	expected := args.Map{"result": float32(2.0)}
	expected.ShouldBeEqual(t, 0, "MinFloat32 returns correct value -- right smaller", actual)
}

// ── MaxInt ──

func Test_MaxInt_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MaxInt(10, 5)}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "MaxInt returns correct value -- left greater", actual)
}

func Test_MaxInt_RightGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MaxInt(3, 8)}

	// Assert
	expected := args.Map{"result": 8}
	expected.ShouldBeEqual(t, 0, "MaxInt returns correct value -- right greater", actual)
}

// ── MinInt ──

func Test_MinInt_LeftSmaller(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MinInt(2, 9)}

	// Assert
	expected := args.Map{"result": 2}
	expected.ShouldBeEqual(t, 0, "MinInt returns correct value -- left smaller", actual)
}

func Test_MinInt_RightSmaller(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.MinInt(9, 3)}

	// Assert
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "MinInt returns correct value -- right smaller", actual)
}

// ── IsRangeWithin.Integer ──

func Test_IntegerWithin_ToByte(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":    coremath.IsRangeWithin.Integer.ToByte(100),
		"outOfRange": coremath.IsRangeWithin.Integer.ToByte(300),
		"negative":   coremath.IsRangeWithin.Integer.ToByte(-1),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outOfRange": false,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToByte", actual)
}

func Test_IntegerWithin_ToUnsignedInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToUnsignedInt16(100),
		"outRange": coremath.IsRangeWithin.Integer.ToUnsignedInt16(70000),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToUnsignedInt16", actual)
}

func Test_IntegerWithin_ToUnsignedInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer.ToUnsignedInt32(-1),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToUnsignedInt32", actual)
}

func Test_IntegerWithin_ToUnsignedInt64(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer.ToUnsignedInt64(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToUnsignedInt64", actual)
}

func Test_IntegerWithin_ToInt8(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer.ToInt8(200),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToInt8", actual)
}

func Test_IntegerWithin_ToInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer.ToInt16(40000),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToInt16", actual)
}

func Test_IntegerWithin_ToInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer.ToInt32(1000),
		"outRange": coremath.IsRangeWithin.Integer.ToInt32(math.MaxInt32 + 1),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToInt32", actual)
}

// ── IsRangeWithin.Integer16 ──

func Test_Integer16Within_ToByte(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer16.ToByte(100),
		"outRange": coremath.IsRangeWithin.Integer16.ToByte(300),
		"negative": coremath.IsRangeWithin.Integer16.ToByte(-1),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer16Within returns non-empty -- ToByte", actual)
}

func Test_Integer16Within_ToUnsignedInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt16(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt16(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer16Within returns non-empty -- ToUnsignedInt16", actual)
}

func Test_Integer16Within_ToUnsignedInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt32(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer16Within returns non-empty -- ToUnsignedInt32", actual)
}

func Test_Integer16Within_ToUnsignedInt64(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt64(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer16Within returns non-empty -- ToUnsignedInt64", actual)
}

func Test_Integer16Within_ToInt8(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer16.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer16.ToInt8(200),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer16Within returns non-empty -- ToInt8", actual)
}

// ── IsRangeWithin.Integer32 ──

func Test_Integer32Within_ToByte(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToByte(200),
		"outRange": coremath.IsRangeWithin.Integer32.ToByte(300),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToByte", actual)
}

func Test_Integer32Within_ToUnsignedInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToUnsignedInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer32.ToUnsignedInt16(70000),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToUnsignedInt16", actual)
}

func Test_Integer32Within_ToUnsignedInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer32.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer32.ToUnsignedInt32(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToUnsignedInt32", actual)
}

func Test_Integer32Within_ToUnsignedInt64(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer32.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer32.ToUnsignedInt64(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToUnsignedInt64", actual)
}

func Test_Integer32Within_ToInt8(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer32.ToInt8(200),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToInt8", actual)
}

func Test_Integer32Within_ToInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer32.ToInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer32.ToInt16(40000),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToInt16", actual)
}

func Test_Integer32Within_ToInt(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange": coremath.IsRangeWithin.Integer32.ToInt(1000),
	}

	// Assert
	expected := args.Map{"inRange": true}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToInt", actual)
}

// ── IsRangeWithin.Integer64 ──

func Test_Integer64Within_ToByte(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToByte(200),
		"outRange": coremath.IsRangeWithin.Integer64.ToByte(300),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToByte", actual)
}

func Test_Integer64Within_ToUnsignedInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToUnsignedInt16(100),
		"outRange": coremath.IsRangeWithin.Integer64.ToUnsignedInt16(70000),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToUnsignedInt16", actual)
}

func Test_Integer64Within_ToUnsignedInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToUnsignedInt32(100),
		"outRange": coremath.IsRangeWithin.Integer64.ToUnsignedInt32(-1),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToUnsignedInt32", actual)
}

func Test_Integer64Within_ToUnsignedInt64(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer64.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer64.ToUnsignedInt64(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToUnsignedInt64", actual)
}

func Test_Integer64Within_ToInt8(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToInt8(50),
		"outRange": coremath.IsRangeWithin.Integer64.ToInt8(200),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToInt8", actual)
}

func Test_Integer64Within_ToInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToInt16(1000),
		"outRange": coremath.IsRangeWithin.Integer64.ToInt16(40000),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToInt16", actual)
}

func Test_Integer64Within_ToInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.Integer64.ToInt32(1000),
		"outRange": coremath.IsRangeWithin.Integer64.ToInt32(int64(math.MaxInt32) + 1),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToInt32", actual)
}

func Test_Integer64Within_ToInt(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange": coremath.IsRangeWithin.Integer64.ToInt(1000),
	}

	// Assert
	expected := args.Map{"inRange": true}
	expected.ShouldBeEqual(t, 0, "Integer64Within returns non-empty -- ToInt", actual)
}

// ── IsRangeWithin.UnsignedInteger16 ──

func Test_UnsignedInt16Within_ToByte(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.UnsignedInteger16.ToByte(200),
		"outRange": coremath.IsRangeWithin.UnsignedInteger16.ToByte(300),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "UnsignedInt16Within returns non-empty -- ToByte", actual)
}

func Test_UnsignedInt16Within_ToInt8(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsRangeWithin.UnsignedInteger16.ToInt8(50),
		"outRange": coremath.IsRangeWithin.UnsignedInteger16.ToInt8(200),
	}

	// Assert
	expected := args.Map{
		"inRange": true,
		"outRange": false,
	}
	expected.ShouldBeEqual(t, 0, "UnsignedInt16Within returns non-empty -- ToInt8", actual)
}

// ── IsOutOfRange.Integer ──

func Test_IntegerOutOfRange_ToByte(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToByte(100),
		"outRange": coremath.IsOutOfRange.Integer.ToByte(300),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToByte", actual)
}

func Test_IntegerOutOfRange_ToUnsignedInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToUnsignedInt16(100),
		"outRange": coremath.IsOutOfRange.Integer.ToUnsignedInt16(70000),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToUnsignedInt16", actual)
}

func Test_IntegerOutOfRange_ToUnsignedInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer.ToUnsignedInt32(100),
	}

	// Assert
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToUnsignedInt32", actual)
}

func Test_IntegerOutOfRange_ToUnsignedInt64(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsOutOfRange.Integer.ToUnsignedInt64(100),
		"negative": coremath.IsOutOfRange.Integer.ToUnsignedInt64(-1),
	}

	// Assert
	expected := args.Map{
		"positive": false,
		"negative": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToUnsignedInt64", actual)
}

func Test_IntegerOutOfRange_ToInt8(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToInt8(50),
		"outRange": coremath.IsOutOfRange.Integer.ToInt8(200),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToInt8", actual)
}

func Test_IntegerOutOfRange_ToInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer.ToInt16(1000),
		"outRange": coremath.IsOutOfRange.Integer.ToInt16(40000),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToInt16", actual)
}

func Test_IntegerOutOfRange_ToInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer.ToInt32(1000),
	}

	// Assert
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToInt32", actual)
}

func Test_IntegerOutOfRange_ToInt_FromMaxByteLeftGreaterV2(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer.ToInt(1000),
	}

	// Assert
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToInt", actual)
}

// ── IsOutOfRange.Integer64 ──

func Test_Integer64OutOfRange_Byte(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Byte(100),
		"outRange": coremath.IsOutOfRange.Integer64.Byte(300),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- Byte", actual)
}

func Test_Integer64OutOfRange_UnsignedInt16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.UnsignedInt16(100),
		"outRange": coremath.IsOutOfRange.Integer64.UnsignedInt16(70000),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- UnsignedInt16", actual)
}

func Test_Integer64OutOfRange_UnsignedInt32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.UnsignedInt32(100),
		"outRange": coremath.IsOutOfRange.Integer64.UnsignedInt32(-1),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- UnsignedInt32", actual)
}

func Test_Integer64OutOfRange_UnsignedInt64(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coremath.IsOutOfRange.Integer64.UnsignedInt64(100),
		"negative": coremath.IsOutOfRange.Integer64.UnsignedInt64(-1),
	}

	// Assert
	expected := args.Map{
		"positive": false,
		"negative": true,
	}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- UnsignedInt64", actual)
}

func Test_Integer64OutOfRange_Int8(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Int8(50),
		"outRange": coremath.IsOutOfRange.Integer64.Int8(200),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- Int8", actual)
}

func Test_Integer64OutOfRange_Int16(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Int16(1000),
		"outRange": coremath.IsOutOfRange.Integer64.Int16(40000),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- Int16", actual)
}

func Test_Integer64OutOfRange_Int32(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  coremath.IsOutOfRange.Integer64.Int32(1000),
		"outRange": coremath.IsOutOfRange.Integer64.Int32(int64(math.MaxInt32) + 1),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- Int32", actual)
}

func Test_Integer64OutOfRange_Int(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange": coremath.IsOutOfRange.Integer64.Int(1000),
	}

	// Assert
	expected := args.Map{"inRange": false}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- Int", actual)
}
