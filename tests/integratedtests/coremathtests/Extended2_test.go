package coremathtests

import (
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coremath"
)

// TestMaxByte verifies MaxByte returns the larger byte.
func TestMaxByte(t *testing.T) {
	for _, tc := range maxByteCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := coremath.MaxByte(tc.left, tc.right)

			// Assert
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
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
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

// TestMaxFloat32 verifies MaxFloat32.
func TestMaxFloat32(t *testing.T) {
	if coremath.MaxFloat32(1.5, 2.5) != 2.5 {
		t.Error("expected 2.5")
	}
	if coremath.MaxFloat32(3.0, 1.0) != 3.0 {
		t.Error("expected 3.0")
	}
}

// TestMinFloat32 verifies MinFloat32.
func TestMinFloat32(t *testing.T) {
	if coremath.MinFloat32(1.5, 2.5) != 1.5 {
		t.Error("expected 1.5")
	}
	if coremath.MinFloat32(3.0, 1.0) != 1.0 {
		t.Error("expected 1.0")
	}
}

// TestIsRangeWithin_Integer verifies integerWithin checks.
func TestIsRangeWithin_Integer(t *testing.T) {
	w := coremath.IsRangeWithin.Integer
	if !w.ToByte(0) { t.Error("0 within byte") }
	if !w.ToByte(255) { t.Error("255 within byte") }
	if w.ToByte(-1) { t.Error("-1 not within byte") }
	if w.ToByte(256) { t.Error("256 not within byte") }
	if !w.ToInt8(0) { t.Error("0 within int8") }
	if w.ToInt8(128) { t.Error("128 not within int8") }
	if !w.ToInt16(0) { t.Error("0 within int16") }
	if !w.ToInt32(0) { t.Error("0 within int32") }
	if !w.ToUnsignedInt16(0) { t.Error("0 within uint16") }
	if !w.ToUnsignedInt32(0) { t.Error("0 within uint32") }
	if !w.ToUnsignedInt64(0) { t.Error("0 within uint64") }
	if w.ToUnsignedInt64(-1) { t.Error("-1 not within uint64") }
}

// TestIsRangeWithin_Integer16 verifies integer16Within.
func TestIsRangeWithin_Integer16(t *testing.T) {
	w := coremath.IsRangeWithin.Integer16
	if !w.ToByte(0) { t.Error("0 within byte") }
	if !w.ToByte(255) { t.Error("255 within byte") }
	if w.ToByte(-1) { t.Error("-1 not within byte") }
	if w.ToByte(256) { t.Error("256 not within byte") }
	if !w.ToUnsignedInt16(0) { t.Error("0 within uint16") }
	if w.ToUnsignedInt16(-1) { t.Error("-1 not within uint16") }
	if !w.ToUnsignedInt32(0) { t.Error("0 within uint32") }
	if !w.ToUnsignedInt64(0) { t.Error("0 within uint64") }
	if !w.ToInt8(0) { t.Error("0 within int8") }
	if w.ToInt8(int16(math.MaxInt8 + 1)) { t.Error("129 not within int8") }
}

// TestIsRangeWithin_Integer32 verifies integer32Within.
func TestIsRangeWithin_Integer32(t *testing.T) {
	w := coremath.IsRangeWithin.Integer32
	if !w.ToByte(0) { t.Error("0 within byte") }
	if w.ToByte(-1) { t.Error("-1 not within byte") }
	if !w.ToUnsignedInt16(0) { t.Error("0 within uint16") }
	if !w.ToUnsignedInt32(0) { t.Error("0 within uint32") }
	if !w.ToUnsignedInt64(0) { t.Error("0 within uint64") }
	if !w.ToInt8(0) { t.Error("0 within int8") }
	if !w.ToInt16(0) { t.Error("0 within int16") }
	if !w.ToInt(0) { t.Error("0 within int") }
}

// TestIsRangeWithin_Integer64 verifies integer64Within.
func TestIsRangeWithin_Integer64(t *testing.T) {
	w := coremath.IsRangeWithin.Integer64
	if !w.ToByte(0) { t.Error("0 within byte") }
	if w.ToByte(-1) { t.Error("-1 not within byte") }
	if w.ToByte(256) { t.Error("256 not within byte") }
	if !w.ToUnsignedInt16(0) { t.Error("0 within uint16") }
	if !w.ToUnsignedInt32(0) { t.Error("0 within uint32") }
	if !w.ToUnsignedInt64(0) { t.Error("0 within uint64") }
	if w.ToUnsignedInt64(-1) { t.Error("-1 not within uint64") }
	if !w.ToInt8(0) { t.Error("0 within int8") }
	if !w.ToInt16(0) { t.Error("0 within int16") }
	if !w.ToInt32(0) { t.Error("0 within int32") }
	if !w.ToInt(0) { t.Error("0 within int") }
}

// TestIsRangeWithin_UnsignedInteger16 verifies unsignedInteger16Within.
func TestIsRangeWithin_UnsignedInteger16(t *testing.T) {
	w := coremath.IsRangeWithin.UnsignedInteger16
	if !w.ToByte(0) { t.Error("0 within byte") }
	if !w.ToByte(255) { t.Error("255 within byte") }
	if w.ToByte(256) { t.Error("256 not within byte") }
	if !w.ToInt8(0) { t.Error("0 within int8") }
	if w.ToInt8(128) { t.Error("128 not within int8") }
}

// TestIsOutOfRange_Integer verifies integerOutOfRange.
func TestIsOutOfRange_Integer(t *testing.T) {
	w := coremath.IsOutOfRange.Integer
	if w.ToByte(0) { t.Error("0 is in range for byte") }
	if !w.ToByte(-1) { t.Error("-1 is out of range for byte") }
	if !w.ToByte(256) { t.Error("256 is out of range for byte") }
	if w.ToInt8(0) { t.Error("0 in range for int8") }
	if !w.ToInt8(128) { t.Error("128 out of range for int8") }
	if w.ToInt16(0) { t.Error("0 in range for int16") }
	if w.ToInt32(0) { t.Error("0 in range for int32") }
	if w.ToUnsignedInt16(0) { t.Error("0 in range for uint16") }
	if w.ToUnsignedInt32(0) { t.Error("0 in range for uint32") }
	if w.ToUnsignedInt64(0) { t.Error("0 in range for uint64") }
	if !w.ToUnsignedInt64(-1) { t.Error("-1 out of range for uint64") }
	if w.ToInt(0) { t.Error("0 in range for int") }
}

// TestIsOutOfRange_Integer64 verifies integer64OutOfRange.
func TestIsOutOfRange_Integer64(t *testing.T) {
	w := coremath.IsOutOfRange.Integer64
	if w.Byte(0) { t.Error("0 in range for byte") }
	if !w.Byte(-1) { t.Error("-1 out of range for byte") }
	if !w.Byte(256) { t.Error("256 out of range for byte") }
	if w.Int8(0) { t.Error("0 in range for int8") }
	if w.Int16(0) { t.Error("0 in range for int16") }
	if w.Int32(0) { t.Error("0 in range for int32") }
	if w.Int(0) { t.Error("0 in range for int") }
	if w.UnsignedInt16(0) { t.Error("0 in range for uint16") }
	if w.UnsignedInt32(0) { t.Error("0 in range for uint32") }
	if w.UnsignedInt64(0) { t.Error("0 in range for uint64") }
	if !w.UnsignedInt64(-1) { t.Error("-1 out of range for uint64") }
}
