package coremathtests

import (
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── MaxInt / MinInt equal values ──

func Test_Cov2_MaxInt_Equal(t *testing.T) {
	actual := args.Map{"result": coremath.MaxInt(5, 5)}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "MaxInt returns correct value -- equal", actual)
}

func Test_Cov2_MinInt_Equal(t *testing.T) {
	actual := args.Map{"result": coremath.MinInt(5, 5)}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "MinInt returns correct value -- equal", actual)
}

// ── MaxByte / MinByte equal ──

func Test_Cov2_MaxByte_Equal(t *testing.T) {
	actual := args.Map{"result": coremath.MaxByte(5, 5)}
	expected := args.Map{"result": byte(5)}
	expected.ShouldBeEqual(t, 0, "MaxByte returns correct value -- equal", actual)
}

func Test_Cov2_MinByte_Equal(t *testing.T) {
	actual := args.Map{"result": coremath.MinByte(5, 5)}
	expected := args.Map{"result": byte(5)}
	expected.ShouldBeEqual(t, 0, "MinByte returns correct value -- equal", actual)
}

// ── MaxFloat32 / MinFloat32 equal ──

func Test_Cov2_MaxFloat32_Equal(t *testing.T) {
	actual := args.Map{"result": coremath.MaxFloat32(3.14, 3.14)}
	expected := args.Map{"result": float32(3.14)}
	expected.ShouldBeEqual(t, 0, "MaxFloat32 returns correct value -- equal", actual)
}

func Test_Cov2_MinFloat32_Equal(t *testing.T) {
	actual := args.Map{"result": coremath.MinFloat32(3.14, 3.14)}
	expected := args.Map{"result": float32(3.14)}
	expected.ShouldBeEqual(t, 0, "MinFloat32 returns correct value -- equal", actual)
}

// ── Integer boundary edge cases ──

func Test_Cov2_IntegerWithin_ToByte_Exact(t *testing.T) {
	actual := args.Map{
		"min": coremath.IsRangeWithin.Integer.ToByte(0),
		"max": coremath.IsRangeWithin.Integer.ToByte(255),
	}
	expected := args.Map{"min": true, "max": true}
	expected.ShouldBeEqual(t, 0, "IntegerWithin returns non-empty -- ToByte exact boundaries", actual)
}

func Test_Cov2_IntegerOutOfRange_ToInt_Negative(t *testing.T) {
	actual := args.Map{
		"negative": coremath.IsOutOfRange.Integer.ToInt(-1),
	}
	expected := args.Map{"negative": false}
	expected.ShouldBeEqual(t, 0, "IntegerOutOfRange returns correct value -- ToInt negative", actual)
}

// ── Integer64 boundary edges ──

func Test_Cov2_Integer64OutOfRange_Int32_ExactMax(t *testing.T) {
	actual := args.Map{
		"exactMax": coremath.IsOutOfRange.Integer64.Int32(int64(math.MaxInt32)),
	}
	expected := args.Map{"exactMax": false}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- Int32 exact max", actual)
}

func Test_Cov2_Integer64OutOfRange_Int_Large(t *testing.T) {
	actual := args.Map{
		"normal": coremath.IsOutOfRange.Integer64.Int(1000),
	}
	expected := args.Map{"normal": false}
	expected.ShouldBeEqual(t, 0, "Integer64OutOfRange returns correct value -- Int normal", actual)
}

// ── Integer32 boundary edges ──

func Test_Cov2_Integer32Within_ToByte_Negative(t *testing.T) {
	actual := args.Map{"result": coremath.IsRangeWithin.Integer32.ToByte(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Integer32Within returns non-empty -- ToByte negative", actual)
}

// ── UnsignedInteger16 ──

func Test_Cov2_UnsignedInt16Within_ToInt8_Boundary(t *testing.T) {
	actual := args.Map{
		"exact127": coremath.IsRangeWithin.UnsignedInteger16.ToInt8(127),
		"exact128": coremath.IsRangeWithin.UnsignedInteger16.ToInt8(128),
	}
	expected := args.Map{"exact127": true, "exact128": false}
	expected.ShouldBeEqual(t, 0, "UnsignedInt16Within returns non-empty -- ToInt8 boundary", actual)
}
