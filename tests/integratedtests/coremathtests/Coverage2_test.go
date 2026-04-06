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
