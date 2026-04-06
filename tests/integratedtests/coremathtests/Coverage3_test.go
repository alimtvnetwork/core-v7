package coremathtests

import (
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// MaxByte / MinByte — branch coverage
// ═══════════════════════════════════════════

func Test_Cov3_MaxByte_LeftGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxByte(10, 5)}
	expected := args.Map{"result": byte(10)}
	expected.ShouldBeEqual(t, 0, "MaxByte returns left -- left greater", actual)
}

func Test_Cov3_MaxByte_RightGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxByte(5, 10)}
	expected := args.Map{"result": byte(10)}
	expected.ShouldBeEqual(t, 0, "MaxByte returns right -- right greater", actual)
}

func Test_Cov3_MinByte_LeftSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinByte(3, 10)}
	expected := args.Map{"result": byte(3)}
	expected.ShouldBeEqual(t, 0, "MinByte returns left -- left smaller", actual)
}

func Test_Cov3_MinByte_RightSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinByte(10, 3)}
	expected := args.Map{"result": byte(3)}
	expected.ShouldBeEqual(t, 0, "MinByte returns right -- right smaller", actual)
}

// ═══════════════════════════════════════════
// MaxFloat32 / MinFloat32 — branch coverage
// ═══════════════════════════════════════════

func Test_Cov3_MaxFloat32_LeftGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxFloat32(5.5, 3.3)}
	expected := args.Map{"result": float32(5.5)}
	expected.ShouldBeEqual(t, 0, "MaxFloat32 returns left -- left greater", actual)
}

func Test_Cov3_MaxFloat32_RightGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxFloat32(3.3, 5.5)}
	expected := args.Map{"result": float32(5.5)}
	expected.ShouldBeEqual(t, 0, "MaxFloat32 returns right -- right greater", actual)
}

func Test_Cov3_MinFloat32_LeftSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinFloat32(1.1, 9.9)}
	expected := args.Map{"result": float32(1.1)}
	expected.ShouldBeEqual(t, 0, "MinFloat32 returns left -- left smaller", actual)
}

func Test_Cov3_MinFloat32_RightSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinFloat32(9.9, 1.1)}
	expected := args.Map{"result": float32(1.1)}
	expected.ShouldBeEqual(t, 0, "MinFloat32 returns right -- right smaller", actual)
}

// ═══════════════════════════════════════════
// MaxInt / MinInt — branch coverage
// ═══════════════════════════════════════════

func Test_Cov3_MaxInt_LeftGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxInt(10, 5)}
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "MaxInt returns left -- left greater", actual)
}

func Test_Cov3_MaxInt_RightGreater(t *testing.T) {
	actual := args.Map{"result": coremath.MaxInt(5, 10)}
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "MaxInt returns right -- right greater", actual)
}

func Test_Cov3_MinInt_LeftSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinInt(3, 10)}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "MinInt returns left -- left smaller", actual)
}

func Test_Cov3_MinInt_RightSmaller(t *testing.T) {
	actual := args.Map{"result": coremath.MinInt(10, 3)}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "MinInt returns right -- right smaller", actual)
}

// ═══════════════════════════════════════════
// Integer16Within
// ═══════════════════════════════════════════

func Test_Cov3_Integer16Within_ToByte(t *testing.T) {
	actual := args.Map{
		"zero":     coremath.IsRangeWithin.Integer16.ToByte(0),
		"max255":   coremath.IsRangeWithin.Integer16.ToByte(255),
		"over255":  coremath.IsRangeWithin.Integer16.ToByte(256),
		"negative": coremath.IsRangeWithin.Integer16.ToByte(-1),
	}
	expected := args.Map{"zero": true, "max255": true, "over255": false, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToByte returns bool -- boundary", actual)
}

func Test_Cov3_Integer16Within_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"zero":     coremath.IsRangeWithin.Integer16.ToUnsignedInt16(0),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt16(-1),
	}
	expected := args.Map{"zero": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToUnsignedInt16 returns bool -- boundary", actual)
}

func Test_Cov3_Integer16Within_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt32(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToUnsignedInt32 returns bool -- boundary", actual)
}

func Test_Cov3_Integer16Within_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer16.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer16.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToUnsignedInt64 returns bool -- boundary", actual)
}

func Test_Cov3_Integer16Within_ToInt8(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer16.ToInt8(127),
		"outside": coremath.IsRangeWithin.Integer16.ToInt8(128),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer16Within ToInt8 returns bool -- boundary", actual)
}

// ═══════════════════════════════════════════
// Integer32Within
// ═══════════════════════════════════════════

func Test_Cov3_Integer32Within_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer32.ToUnsignedInt16(int32(math.MaxUint16)),
		"outside": coremath.IsRangeWithin.Integer32.ToUnsignedInt16(int32(math.MaxUint16) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToUnsignedInt16 returns bool -- boundary", actual)
}

func Test_Cov3_Integer32Within_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer32.ToUnsignedInt32(100),
		"negative": coremath.IsRangeWithin.Integer32.ToUnsignedInt32(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToUnsignedInt32 returns bool -- boundary", actual)
}

func Test_Cov3_Integer32Within_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer32.ToUnsignedInt64(100),
		"negative": coremath.IsRangeWithin.Integer32.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToUnsignedInt64 returns bool -- boundary", actual)
}

func Test_Cov3_Integer32Within_ToInt8(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer32.ToInt8(int32(math.MaxInt8)),
		"outside": coremath.IsRangeWithin.Integer32.ToInt8(int32(math.MaxInt8) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToInt8 returns bool -- boundary", actual)
}

func Test_Cov3_Integer32Within_ToInt16(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer32.ToInt16(int32(math.MaxInt16)),
		"outside": coremath.IsRangeWithin.Integer32.ToInt16(int32(math.MaxInt16) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToInt16 returns bool -- boundary", actual)
}

func Test_Cov3_Integer32Within_ToInt(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer32.ToInt(100),
		"outside": coremath.IsRangeWithin.Integer32.ToInt(math.MaxInt32 + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within ToInt returns bool -- boundary", actual)
}

// ═══════════════════════════════════════════
// Integer64Within
// ═══════════════════════════════════════════

func Test_Cov3_Integer64Within_ToByte(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer64.ToByte(255),
		"outside": coremath.IsRangeWithin.Integer64.ToByte(256),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToByte returns bool -- boundary", actual)
}

func Test_Cov3_Integer64Within_ToUnsignedInt16(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer64.ToUnsignedInt16(int64(math.MaxUint16)),
		"outside": coremath.IsRangeWithin.Integer64.ToUnsignedInt16(int64(math.MaxUint16) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToUnsignedInt16 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64Within_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer64.ToUnsignedInt32(int64(math.MaxUint32)),
		"outside": coremath.IsRangeWithin.Integer64.ToUnsignedInt32(int64(math.MaxUint32) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToUnsignedInt32 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64Within_ToUnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsRangeWithin.Integer64.ToUnsignedInt64(0),
		"negative": coremath.IsRangeWithin.Integer64.ToUnsignedInt64(-1),
	}
	expected := args.Map{"positive": true, "negative": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToUnsignedInt64 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64Within_ToInt8(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer64.ToInt8(int64(math.MaxInt8)),
		"outside": coremath.IsRangeWithin.Integer64.ToInt8(int64(math.MaxInt8) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt8 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64Within_ToInt16(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer64.ToInt16(int64(math.MaxInt16)),
		"outside": coremath.IsRangeWithin.Integer64.ToInt16(int64(math.MaxInt16) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt16 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64Within_ToInt32(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer64.ToInt32(int64(math.MaxInt32)),
		"outside": coremath.IsRangeWithin.Integer64.ToInt32(int64(math.MaxInt32) + 1),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt32 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64Within_ToInt(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.Integer64.ToInt(100),
		"minInt64": coremath.IsRangeWithin.Integer64.ToInt(-9223372036854775808),
	}
	expected := args.Map{"within": true, "minInt64": true}
	expected.ShouldBeEqual(t, 0, "Integer64Within ToInt returns bool -- boundary", actual)
}

// ═══════════════════════════════════════════
// Integer64OutOfRange
// ═══════════════════════════════════════════

func Test_Cov3_Integer64OutOfRange_Byte(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsOutOfRange.Integer64.Byte(100),
		"outside": coremath.IsOutOfRange.Integer64.Byte(256),
	}
	expected := args.Map{"within": false, "outside": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Byte returns bool -- boundary", actual)
}

func Test_Cov3_Integer64OutOfRange_UnsignedInt16(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsOutOfRange.Integer64.UnsignedInt16(100),
		"outside": coremath.IsOutOfRange.Integer64.UnsignedInt16(int64(math.MaxUint16) + 1),
	}
	expected := args.Map{"within": false, "outside": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange UnsignedInt16 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64OutOfRange_UnsignedInt32(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsOutOfRange.Integer64.UnsignedInt32(100),
		"outside": coremath.IsOutOfRange.Integer64.UnsignedInt32(int64(math.MaxUint32) + 1),
	}
	expected := args.Map{"within": false, "outside": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange UnsignedInt32 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64OutOfRange_UnsignedInt64(t *testing.T) {
	actual := args.Map{
		"positive": coremath.IsOutOfRange.Integer64.UnsignedInt64(100),
		"negative": coremath.IsOutOfRange.Integer64.UnsignedInt64(-1),
	}
	expected := args.Map{"positive": false, "negative": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange UnsignedInt64 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64OutOfRange_Int8(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsOutOfRange.Integer64.Int8(100),
		"outside": coremath.IsOutOfRange.Integer64.Int8(int64(math.MaxInt8) + 1),
	}
	expected := args.Map{"within": false, "outside": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int8 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64OutOfRange_Int16(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsOutOfRange.Integer64.Int16(100),
		"outside": coremath.IsOutOfRange.Integer64.Int16(int64(math.MaxInt16) + 1),
	}
	expected := args.Map{"within": false, "outside": true}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int16 returns bool -- boundary", actual)
}

func Test_Cov3_Integer64OutOfRange_Int32_Within(t *testing.T) {
	actual := args.Map{"result": coremath.IsOutOfRange.Integer64.Int32(100)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int32 returns false -- within range", actual)
}

func Test_Cov3_Integer64OutOfRange_Int(t *testing.T) {
	actual := args.Map{"result": coremath.IsOutOfRange.Integer64.Int(100)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange Int returns false -- within range", actual)
}

// ═══════════════════════════════════════════
// IntegerOutOfRange
// ═══════════════════════════════════════════

func Test_Cov3_IntegerOutOfRange_ToUnsignedInt32(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsOutOfRange.Integer.ToUnsignedInt32(100),
		"outside": coremath.IsOutOfRange.Integer.ToUnsignedInt32(-1),
	}
	expected := args.Map{"within": false, "outside": true}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToUnsignedInt32 returns bool -- boundary", actual)
}

func Test_Cov3_IntegerOutOfRange_ToInt(t *testing.T) {
	actual := args.Map{"result": coremath.IsOutOfRange.Integer.ToInt(100)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange ToInt returns false -- within range", actual)
}

// ═══════════════════════════════════════════
// UnsignedInteger16Within
// ═══════════════════════════════════════════

func Test_Cov3_UnsignedInt16Within_ToByte(t *testing.T) {
	actual := args.Map{
		"within":  coremath.IsRangeWithin.UnsignedInteger16.ToByte(255),
		"outside": coremath.IsRangeWithin.UnsignedInteger16.ToByte(256),
	}
	expected := args.Map{"within": true, "outside": false}
	expected.ShouldBeEqual(t, 0, "UnsignedInt16Within ToByte returns bool -- boundary", actual)
}
