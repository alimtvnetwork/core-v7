package coremathtests

import (
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestMaxByte verifies MaxByte returns the larger byte.
func TestMaxByte(t *testing.T) {
	for _, tc := range maxByteCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := coremath.MaxByte(tc.left, tc.right)

			// Assert
			actual := args.Map{"result": result != tc.expected}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestMinByte verifies MinByte returns the smaller byte.
func TestMinByte(t *testing.T) {
	for _, tc := range minByteCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := coremath.MinByte(tc.left, tc.right)

			// Assert
			actual := args.Map{"result": result != tc.expected}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestMaxFloat32 verifies MaxFloat32.
func TestMaxFloat32(t *testing.T) {
	actual := args.Map{"result": coremath.MaxFloat32(1.5, 2.5) != 2.5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	actual := args.Map{"result": coremath.MaxFloat32(3.0, 1.0) != 3.0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3.0", actual)
}

// TestMinFloat32 verifies MinFloat32.
func TestMinFloat32(t *testing.T) {
	actual := args.Map{"result": coremath.MinFloat32(1.5, 2.5) != 1.5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	actual := args.Map{"result": coremath.MinFloat32(3.0, 1.0) != 1.0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1.0", actual)
}

// TestIsRangeWithin_Integer verifies integerWithin checks.
func TestIsRangeWithin_Integer(t *testing.T) {
	w := coremath.IsRangeWithin.Integer
	actual := args.Map{"result": w.ToByte(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within byte", actual)
	actual := args.Map{"result": w.ToByte(255)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "255 within byte", actual)
	actual := args.Map{"result": w.ToByte(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "-1 not within byte", actual)
	actual := args.Map{"result": w.ToByte(256)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "256 not within byte", actual)
	actual := args.Map{"result": w.ToInt8(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int8", actual)
	actual := args.Map{"result": w.ToInt8(128)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "128 not within int8", actual)
	actual := args.Map{"result": w.ToInt16(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int16", actual)
	actual := args.Map{"result": w.ToInt32(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int32", actual)
	actual := args.Map{"result": w.ToUnsignedInt16(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint16", actual)
	actual := args.Map{"result": w.ToUnsignedInt32(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint32", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint64", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "-1 not within uint64", actual)
}

// TestIsRangeWithin_Integer16 verifies integer16Within.
func TestIsRangeWithin_Integer16(t *testing.T) {
	w := coremath.IsRangeWithin.Integer16
	actual := args.Map{"result": w.ToByte(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within byte", actual)
	actual := args.Map{"result": w.ToByte(255)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "255 within byte", actual)
	actual := args.Map{"result": w.ToByte(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "-1 not within byte", actual)
	actual := args.Map{"result": w.ToByte(256)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "256 not within byte", actual)
	actual := args.Map{"result": w.ToUnsignedInt16(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint16", actual)
	actual := args.Map{"result": w.ToUnsignedInt16(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "-1 not within uint16", actual)
	actual := args.Map{"result": w.ToUnsignedInt32(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint32", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint64", actual)
	actual := args.Map{"result": w.ToInt8(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int8", actual)
	actual := args.Map{"result": w.ToInt8(int16(math.MaxInt8 + 1))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "129 not within int8", actual)
}

// TestIsRangeWithin_Integer32 verifies integer32Within.
func TestIsRangeWithin_Integer32(t *testing.T) {
	w := coremath.IsRangeWithin.Integer32
	actual := args.Map{"result": w.ToByte(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within byte", actual)
	actual := args.Map{"result": w.ToByte(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "-1 not within byte", actual)
	actual := args.Map{"result": w.ToUnsignedInt16(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint16", actual)
	actual := args.Map{"result": w.ToUnsignedInt32(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint32", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint64", actual)
	actual := args.Map{"result": w.ToInt8(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int8", actual)
	actual := args.Map{"result": w.ToInt16(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int16", actual)
	actual := args.Map{"result": w.ToInt(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int", actual)
}

// TestIsRangeWithin_Integer64 verifies integer64Within.
func TestIsRangeWithin_Integer64(t *testing.T) {
	w := coremath.IsRangeWithin.Integer64
	actual := args.Map{"result": w.ToByte(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within byte", actual)
	actual := args.Map{"result": w.ToByte(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "-1 not within byte", actual)
	actual := args.Map{"result": w.ToByte(256)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "256 not within byte", actual)
	actual := args.Map{"result": w.ToUnsignedInt16(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint16", actual)
	actual := args.Map{"result": w.ToUnsignedInt32(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint32", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within uint64", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "-1 not within uint64", actual)
	actual := args.Map{"result": w.ToInt8(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int8", actual)
	actual := args.Map{"result": w.ToInt16(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int16", actual)
	actual := args.Map{"result": w.ToInt32(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int32", actual)
	actual := args.Map{"result": w.ToInt(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int", actual)
}

// TestIsRangeWithin_UnsignedInteger16 verifies unsignedInteger16Within.
func TestIsRangeWithin_UnsignedInteger16(t *testing.T) {
	w := coremath.IsRangeWithin.UnsignedInteger16
	actual := args.Map{"result": w.ToByte(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within byte", actual)
	actual := args.Map{"result": w.ToByte(255)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "255 within byte", actual)
	actual := args.Map{"result": w.ToByte(256)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "256 not within byte", actual)
	actual := args.Map{"result": w.ToInt8(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 within int8", actual)
	actual := args.Map{"result": w.ToInt8(128)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "128 not within int8", actual)
}

// TestIsOutOfRange_Integer verifies integerOutOfRange.
func TestIsOutOfRange_Integer(t *testing.T) {
	w := coremath.IsOutOfRange.Integer
	actual := args.Map{"result": w.ToByte(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 is in range for byte", actual)
	actual := args.Map{"result": w.ToByte(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "-1 is out of range for byte", actual)
	actual := args.Map{"result": w.ToByte(256)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "256 is out of range for byte", actual)
	actual := args.Map{"result": w.ToInt8(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int8", actual)
	actual := args.Map{"result": w.ToInt8(128)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "128 out of range for int8", actual)
	actual := args.Map{"result": w.ToInt16(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int16", actual)
	actual := args.Map{"result": w.ToInt32(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int32", actual)
	actual := args.Map{"result": w.ToUnsignedInt16(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for uint16", actual)
	actual := args.Map{"result": w.ToUnsignedInt32(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for uint32", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for uint64", actual)
	actual := args.Map{"result": w.ToUnsignedInt64(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "-1 out of range for uint64", actual)
	actual := args.Map{"result": w.ToInt(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int", actual)
}

// TestIsOutOfRange_Integer64 verifies integer64OutOfRange.
func TestIsOutOfRange_Integer64(t *testing.T) {
	w := coremath.IsOutOfRange.Integer64
	actual := args.Map{"result": w.Byte(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for byte", actual)
	actual := args.Map{"result": w.Byte(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "-1 out of range for byte", actual)
	actual := args.Map{"result": w.Byte(256)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "256 out of range for byte", actual)
	actual := args.Map{"result": w.Int8(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int8", actual)
	actual := args.Map{"result": w.Int16(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int16", actual)
	actual := args.Map{"result": w.Int32(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int32", actual)
	actual := args.Map{"result": w.Int(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for int", actual)
	actual := args.Map{"result": w.UnsignedInt16(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for uint16", actual)
	actual := args.Map{"result": w.UnsignedInt32(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for uint32", actual)
	actual := args.Map{"result": w.UnsignedInt64(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 in range for uint64", actual)
	actual := args.Map{"result": w.UnsignedInt64(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "-1 out of range for uint64", actual)
}
