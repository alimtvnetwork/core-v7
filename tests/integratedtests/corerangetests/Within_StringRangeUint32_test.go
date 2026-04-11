package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── within.go — uncovered branches ──

func Test_Within_StringRangeUint32_AboveMax(t *testing.T) {
	// Arrange & Act — input that parses but exceeds MaxInt32 after Atoi
	// Actually Atoi would fail for values > MaxInt. Test the boundary:
	// On 64-bit, Atoi("2147483648") succeeds but > MaxInt32
	val, ok := corerange.Within.StringRangeUint32("2147483648")

	// Assert — should return 0 because finalInt > MaxInt32
	actual := args.Map{
		"val": int(val),
		"ok": ok,
	}
	expected := args.Map{
		"val": 2147483647,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Within StringRangeUint32 returns 0 -- above MaxInt32", actual)
}

func Test_Within_StringRangeUint32_Invalid(t *testing.T) {
	// Arrange & Act
	_, ok := corerange.Within.StringRangeUint32("abc")

	// Assert
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeUint32 returns false -- invalid input", actual)
}

func Test_Within_StringRangeFloat_OutOfRange_NoBoundary(t *testing.T) {
	// Arrange & Act — valid float string, out of range, no boundary usage
	val, ok := corerange.Within.StringRangeFloat(false, 10.0, 20.0, "5.0")

	// Assert — isInRange=false, isUsageMinMaxBoundary=false → return 0
	actual := args.Map{
		"val": val == 0,
		"ok": ok,
	}
	expected := args.Map{
		"val": true,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat returns zero -- no boundary out of range", actual)
}

func Test_Within_StringRangeFloat_OutOfRange_WithBoundary(t *testing.T) {
	// Arrange & Act — valid float, out of range, WITH boundary → returns clamped
	val, ok := corerange.Within.StringRangeFloat(true, 10.0, 20.0, "5.0")

	// Assert — clamped to min
	actual := args.Map{
		"isMin": val == 10.0,
		"ok": ok,
	}
	expected := args.Map{
		"isMin": true,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat returns min -- boundary below min", actual)
}

func Test_Within_StringRangeInt16_Invalid(t *testing.T) {
	// Arrange
	_, ok := corerange.Within.StringRangeInt16("abc")

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeInt16 returns false -- invalid input", actual)
}

func Test_Within_StringRangeInt8_Invalid(t *testing.T) {
	// Arrange
	_, ok := corerange.Within.StringRangeInt8("abc")

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeInt8 returns false -- invalid input", actual)
}

func Test_Within_StringRangeUint16_Invalid(t *testing.T) {
	// Arrange
	_, ok := corerange.Within.StringRangeUint16("abc")

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeUint16 returns false -- invalid input", actual)
}

// ── MinMaxInt.go — uncovered branches ──

func Test_MinMaxInt_DifferenceAbsolute_Negative_FromWithinStringRangeUin(t *testing.T) {
	// Arrange — Min > Max → negative diff
	mm := &corerange.MinMaxInt{Min: 10, Max: 3}

	// Act
	actual := args.Map{
		"diffAbs": mm.DifferenceAbsolute(),
		"diff": mm.Difference(),
	}
	expected := args.Map{
		"diffAbs": 7,
		"diff": -7,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt DifferenceAbsolute returns positive -- negative diff", actual)
}

func Test_MinMaxInt_IsInvalidValue_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 1, Max: 10}

	// Act
	actual := args.Map{
		"invalid0": mm.IsInvalidValue(0),
		"invalid5": mm.IsInvalidValue(5),
	}

	// Assert
	expected := args.Map{
		"invalid0": true,
		"invalid5": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt IsInvalidValue returns expected -- boundary check", actual)
}

func Test_MinMaxInt_CreateRangeInt(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}

	// Act
	actual := args.Map{"notNil": mm.CreateRangeInt("3:7", ":") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt CreateRangeInt returns non-nil -- valid", actual)
}

func Test_MinMaxInt_CreateRangeInt8_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}

	// Act
	actual := args.Map{"notNil": mm.CreateRangeInt8("3:7", ":") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_MinMaxInt_CreateRangeInt16_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}

	// Act
	actual := args.Map{"notNil": mm.CreateRangeInt16("3:7", ":") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt CreateRangeInt16 returns non-nil -- valid", actual)
}

func Test_MinMaxInt_RangeLengthInt(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 3, Max: 7}

	// Act
	actual := args.Map{"rangeLenInt": mm.RangeLengthInt()}

	// Assert
	expected := args.Map{"rangeLenInt": 5}
	expected.ShouldBeEqual(t, 0, "MinMaxInt RangeLengthInt returns 5 -- range 3 to 7", actual)
}

func Test_MinMaxInt_NilComparisons(t *testing.T) {
	// Arrange
	var nilMm *corerange.MinMaxInt

	// Act
	actual := args.Map{
		"isMinEqual":    nilMm.IsMinEqual(0),
		"isMinAboveEq":  nilMm.IsMinAboveEqual(0),
		"isMinAbove":    nilMm.IsMinAbove(0),
		"isMinLess":     nilMm.IsMinLess(0),
		"isMinLessEq":   nilMm.IsMinLessEqual(0),
		"isMaxEqual":    nilMm.IsMaxEqual(0),
		"isMaxAboveEq":  nilMm.IsMaxAboveEqual(0),
		"isMaxAbove":    nilMm.IsMaxAbove(0),
		"isMaxLess":     nilMm.IsMaxLess(0),
		"isMaxLessEq":   nilMm.IsMaxLessEqual(0),
	}

	// Assert
	expected := args.Map{
		"isMinEqual": false, "isMinAboveEq": false, "isMinAbove": false,
		"isMinLess": false, "isMinLessEq": false,
		"isMaxEqual": false, "isMaxAboveEq": false, "isMaxAbove": false,
		"isMaxLess": false, "isMaxLessEq": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt nil comparisons return false -- nil receiver", actual)
}

// ── MinMaxInt16.go — uncovered branches ──

func Test_MinMaxInt16_DifferenceAbsolute_Negative_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt16{Min: 10, Max: 3}

	// Act
	actual := args.Map{
		"diffAbs": int(mm.DifferenceAbsolute()),
		"diff": int(mm.Difference()),
	}

	// Assert
	expected := args.Map{
		"diffAbs": 7,
		"diff": -7,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 DifferenceAbsolute returns positive -- reversed range", actual)
}

func Test_MinMaxInt16_NilComparisons(t *testing.T) {
	// Arrange
	var nilMm *corerange.MinMaxInt16

	// Act
	actual := args.Map{
		"isMinEqual": nilMm.IsMinEqual(0), "isMinAboveEq": nilMm.IsMinAboveEqual(0),
		"isMinAbove": nilMm.IsMinAbove(0), "isMinLess": nilMm.IsMinLess(0),
		"isMinLessEq": nilMm.IsMinLessEqual(0),
		"isMaxEqual": nilMm.IsMaxEqual(0), "isMaxAboveEq": nilMm.IsMaxAboveEqual(0),
		"isMaxAbove": nilMm.IsMaxAbove(0), "isMaxLess": nilMm.IsMaxLess(0),
		"isMaxLessEq": nilMm.IsMaxLessEqual(0),
	}

	// Assert
	expected := args.Map{
		"isMinEqual": false, "isMinAboveEq": false, "isMinAbove": false,
		"isMinLess": false, "isMinLessEq": false,
		"isMaxEqual": false, "isMaxAboveEq": false, "isMaxAbove": false,
		"isMaxLess": false, "isMaxLessEq": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 nil comparisons return false -- nil receiver", actual)
}

// ── MinMaxInt8.go — uncovered branches ──

func Test_MinMaxInt8_DifferenceAbsolute_Negative(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt8{Min: 10, Max: 3}

	// Act
	actual := args.Map{
		"diffAbs": int(mm.DifferenceAbsolute()),
		"diff": int(mm.Difference()),
	}

	// Assert
	expected := args.Map{
		"diffAbs": 7,
		"diff": -7,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 DifferenceAbsolute returns positive -- reversed range", actual)
}

func Test_MinMaxInt8_NilComparisons(t *testing.T) {
	// Arrange
	var nilMm *corerange.MinMaxInt8

	// Act
	actual := args.Map{
		"isMinEqual": nilMm.IsMinEqual(0), "isMinAboveEq": nilMm.IsMinAboveEqual(0),
		"isMinAbove": nilMm.IsMinAbove(0), "isMinLess": nilMm.IsMinLess(0),
		"isMinLessEq": nilMm.IsMinLessEqual(0),
		"isMaxEqual": nilMm.IsMaxEqual(0), "isMaxAboveEq": nilMm.IsMaxAboveEqual(0),
		"isMaxAbove": nilMm.IsMaxAbove(0), "isMaxLess": nilMm.IsMaxLess(0),
		"isMaxLessEq": nilMm.IsMaxLessEqual(0),
	}

	// Assert
	expected := args.Map{
		"isMinEqual": false, "isMinAboveEq": false, "isMinAbove": false,
		"isMinLess": false, "isMinLessEq": false,
		"isMaxEqual": false, "isMaxAboveEq": false, "isMaxAbove": false,
		"isMaxLess": false, "isMaxLessEq": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 nil comparisons return false -- nil receiver", actual)
}

func Test_MinMaxInt8_IsInvalidValue(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt8{Min: 1, Max: 10}

	// Act
	actual := args.Map{
		"invalid0": mm.IsInvalidValue(0),
		"invalid5": mm.IsInvalidValue(5),
	}

	// Assert
	expected := args.Map{
		"invalid0": true,
		"invalid5": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 IsInvalidValue returns expected -- boundary", actual)
}

// ── MinMaxInt64.go — uncovered branches ──

func Test_MinMaxInt64_DifferenceAbsolute_Negative(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt64{Min: 10, Max: 3}

	// Act
	actual := args.Map{
		"diffAbs": int(mm.DifferenceAbsolute()),
		"diff": int(mm.Difference()),
	}

	// Assert
	expected := args.Map{
		"diffAbs": 7,
		"diff": -7,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 DifferenceAbsolute returns positive -- reversed range", actual)
}

func Test_MinMaxInt64_NilComparisons(t *testing.T) {
	// Arrange
	var nilMm *corerange.MinMaxInt64

	// Act
	actual := args.Map{
		"isMinEqual": nilMm.IsMinEqual(0), "isMinAboveEq": nilMm.IsMinAboveEqual(0),
		"isMinAbove": nilMm.IsMinAbove(0), "isMinLess": nilMm.IsMinLess(0),
		"isMinLessEq": nilMm.IsMinLessEqual(0),
		"isMaxEqual": nilMm.IsMaxEqual(0), "isMaxAboveEq": nilMm.IsMaxAboveEqual(0),
		"isMaxAbove": nilMm.IsMaxAbove(0), "isMaxLess": nilMm.IsMaxLess(0),
		"isMaxLessEq": nilMm.IsMaxLessEqual(0),
	}

	// Assert
	expected := args.Map{
		"isMinEqual": false, "isMinAboveEq": false, "isMinAbove": false,
		"isMinLess": false, "isMinLessEq": false,
		"isMaxEqual": false, "isMaxAboveEq": false, "isMaxAbove": false,
		"isMaxLess": false, "isMaxLessEq": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 nil comparisons return false -- nil receiver", actual)
}

func Test_MinMaxInt64_IsInvalidValue(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt64{Min: 1, Max: 10}

	// Act
	actual := args.Map{
		"invalid0": mm.IsInvalidValue(0),
		"invalid5": mm.IsInvalidValue(5),
	}

	// Assert
	expected := args.Map{
		"invalid0": true,
		"invalid5": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 IsInvalidValue returns expected -- boundary", actual)
}

func Test_MinMaxInt64_IsWithinRange_Nil(t *testing.T) {
	// Arrange
	var nilMm *corerange.MinMaxInt64

	// Act
	actual := args.Map{"within": nilMm.IsWithinRange(5)}

	// Assert
	expected := args.Map{"within": false}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 IsWithinRange returns false -- nil receiver", actual)
}

// ── MinMaxByte.go — uncovered branches ──

func Test_MinMaxByte_NilComparisons(t *testing.T) {
	// Arrange
	var nilMm *corerange.MinMaxByte

	// Act
	actual := args.Map{
		"isMinEqual": nilMm.IsMinEqual(0), "isMinAboveEq": nilMm.IsMinAboveEqual(0),
		"isMinAbove": nilMm.IsMinAbove(0), "isMinLess": nilMm.IsMinLess(0),
		"isMinLessEq": nilMm.IsMinLessEqual(0),
		"isMaxEqual": nilMm.IsMaxEqual(0), "isMaxAboveEq": nilMm.IsMaxAboveEqual(0),
		"isMaxAbove": nilMm.IsMaxAbove(0), "isMaxLess": nilMm.IsMaxLess(0),
		"isMaxLessEq": nilMm.IsMaxLessEqual(0),
	}

	// Assert
	expected := args.Map{
		"isMinEqual": false, "isMinAboveEq": false, "isMinAbove": false,
		"isMinLess": false, "isMinLessEq": false,
		"isMaxEqual": false, "isMaxAboveEq": false, "isMaxAbove": false,
		"isMaxLess": false, "isMaxLessEq": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxByte nil comparisons return false -- nil receiver", actual)
}

func Test_MinMaxByte_CreateRangeInt8_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	mb := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	actual := args.Map{"notNil": mb.CreateRangeInt8("2:5", ":") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxByte CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_MinMaxByte_CreateRangeInt16_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	mb := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	actual := args.Map{"notNil": mb.CreateRangeInt16("2:5", ":") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxByte CreateRangeInt16 returns non-nil -- valid", actual)
}

func Test_MinMaxByte_CreateRangeInt_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	mb := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	actual := args.Map{"notNil": mb.CreateRangeInt("2:5", ":") != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxByte CreateRangeInt returns non-nil -- valid", actual)
}

// ── RangeByte.go — uncovered branches ──

func Test_RangeByte_Ranges_Invalid_FromWithinStringRangeUin(t *testing.T) {
	// Arrange — create invalid RangeByte
	rb := corerange.NewRangeByteMinMax("abc", "|", 0, 255)

	// Act
	ranges := rb.Ranges()
	rangesInt := rb.RangesInt()

	// Assert
	actual := args.Map{
		"rangesLen": len(ranges),
		"rangesIntLen": len(rangesInt),
	}
	expected := args.Map{
		"rangesLen": 0,
		"rangesIntLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "RangeByte Ranges returns empty -- invalid input", actual)
}

func Test_RangeByte_IsInvalidValue(t *testing.T) {
	// Arrange
	rb := corerange.NewRangeByteMinMax("3:7", ":", 0, 10)

	// Act
	actual := args.Map{
		"invalidOutside": rb.IsInvalidValue(10),
		"invalidInside":  rb.IsInvalidValue(5),
	}

	// Assert
	expected := args.Map{
		"invalidOutside": true,
		"invalidInside": false,
	}
	expected.ShouldBeEqual(t, 0, "RangeByte IsInvalidValue returns expected -- boundary", actual)
}

func Test_RangeByte_String(t *testing.T) {
	// Arrange
	rb := corerange.NewRangeByteMinMax("3:7", ":", 0, 10)

	// Act
	actual := args.Map{"notEmpty": rb.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeByte String returns non-empty -- valid", actual)
}

func Test_RangeByte_NewRangeByte_WithMinMax(t *testing.T) {
	// Arrange — non-nil minMax
	mb := &corerange.MinMaxByte{Min: 0, Max: 255}
	rb := corerange.NewRangeByte("3:7", ":", mb)

	// Assert
	actual := args.Map{
		"notNil": rb != nil,
		"start": int(rb.Start),
	}
	expected := args.Map{
		"notNil": true,
		"start": 3,
	}
	expected.ShouldBeEqual(t, 0, "NewRangeByte returns valid -- non-nil minMax", actual)
}

// ── RangeInt8.go — uncovered branches ──

func Test_RangeInt8_Ranges_Invalid(t *testing.T) {
	// Arrange
	ri8 := corerange.NewRangeInt8MinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"rangesLen": len(ri8.Ranges())}

	// Assert
	expected := args.Map{"rangesLen": 0}
	expected.ShouldBeEqual(t, 0, "RangeInt8 Ranges returns empty -- invalid input", actual)
}

func Test_RangeInt8_DifferenceAbsolute_Negative(t *testing.T) {
	// Arrange — reversed range
	ri8 := &corerange.RangeInt8{
		BaseRange: &corerange.BaseRange{IsValid: true},
		Start:     7,
		End:       3,
	}
	actual := args.Map{"diffAbs": int(ri8.DifferenceAbsolute())}
	expected := args.Map{"diffAbs": 4}
	expected.ShouldBeEqual(t, 0, "RangeInt8 DifferenceAbsolute returns positive -- reversed", actual)
}

func Test_RangeInt8_IsInvalidValue_NotValid(t *testing.T) {
	// Arrange
	ri8 := corerange.NewRangeInt8MinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"invalid": ri8.IsInvalidValue(5)}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "RangeInt8 IsInvalidValue returns true -- not valid", actual)
}

func Test_RangeInt8_String(t *testing.T) {
	// Arrange
	ri8 := corerange.NewRangeInt8MinMax("3:7", ":", 0, 10)

	// Act
	actual := args.Map{"notEmpty": ri8.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeInt8 String returns non-empty -- valid", actual)
}

func Test_RangeInt8_NewRangeInt8_WithMinMax(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt8{Min: 0, Max: 10}
	ri8 := corerange.NewRangeInt8("3:7", ":", mm)

	// Act
	actual := args.Map{
		"notNil": ri8 != nil,
		"start": int(ri8.Start),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"start": 3,
	}
	expected.ShouldBeEqual(t, 0, "NewRangeInt8 returns valid -- non-nil minMax", actual)
}

func Test_RangeInt8_IsValidPlusWithinRange_NotValid(t *testing.T) {
	// Arrange
	ri8 := corerange.NewRangeInt8MinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"result": ri8.IsValidPlusWithinRange(5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeInt8 IsValidPlusWithinRange returns false -- not valid", actual)
}

// ── RangeInt16.go — uncovered branches ──

func Test_RangeInt16_Ranges_Invalid(t *testing.T) {
	// Arrange
	ri16 := corerange.NewRangeInt16MinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"rangesLen": len(ri16.Ranges())}

	// Assert
	expected := args.Map{"rangesLen": 0}
	expected.ShouldBeEqual(t, 0, "RangeInt16 Ranges returns empty -- invalid input", actual)
}

func Test_RangeInt16_DifferenceAbsolute_Negative_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	ri16 := &corerange.RangeInt16{
		BaseRange: &corerange.BaseRange{IsValid: true},
		Start:     7,
		End:       3,
	}

	// Act
	actual := args.Map{"diffAbs": int(ri16.DifferenceAbsolute())}

	// Assert
	expected := args.Map{"diffAbs": 4}
	expected.ShouldBeEqual(t, 0, "RangeInt16 DifferenceAbsolute returns positive -- reversed", actual)
}

func Test_RangeInt16_IsInvalidValue_NotValid(t *testing.T) {
	// Arrange
	ri16 := corerange.NewRangeInt16MinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"invalid": ri16.IsInvalidValue(5)}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "RangeInt16 IsInvalidValue returns true -- not valid", actual)
}

func Test_RangeInt16_String(t *testing.T) {
	// Arrange
	ri16 := corerange.NewRangeInt16MinMax("3:7", ":", 0, 10)

	// Act
	actual := args.Map{"notEmpty": ri16.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeInt16 String returns non-empty -- valid", actual)
}

func Test_RangeInt16_NewRangeInt16_WithMinMax(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt16{Min: 0, Max: 10}
	ri16 := corerange.NewRangeInt16("3:7", ":", mm)

	// Act
	actual := args.Map{
		"notNil": ri16 != nil,
		"start": int(ri16.Start),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"start": 3,
	}
	expected.ShouldBeEqual(t, 0, "NewRangeInt16 returns valid -- non-nil minMax", actual)
}

func Test_RangeInt16_IsValidPlusWithinRange_NotValid(t *testing.T) {
	// Arrange
	ri16 := corerange.NewRangeInt16MinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"result": ri16.IsValidPlusWithinRange(5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeInt16 IsValidPlusWithinRange returns false -- not valid", actual)
}

func Test_RangeInt16_RangesInt16(t *testing.T) {
	// Arrange
	ri16 := corerange.NewRangeInt16MinMax("3:7", ":", 0, 10)

	// Act
	actual := args.Map{"len": len(ri16.RangesInt16())}

	// Assert
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "RangeInt16 RangesInt16 returns 5 -- valid range", actual)
}

// ── RangeInt — uncovered branches ──

func Test_RangeInt_Ranges_Invalid(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"rangesLen": len(ri.Ranges())}

	// Assert
	expected := args.Map{"rangesLen": 0}
	expected.ShouldBeEqual(t, 0, "RangeInt Ranges returns empty -- invalid input", actual)
}

func Test_RangeInt_IsValidPlusWithinRange_NotValid(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("abc", ":", 0, 10)

	// Act
	actual := args.Map{"result": ri.IsValidPlusWithinRange(5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeInt IsValidPlusWithinRange returns false -- not valid", actual)
}

// ── StartEndInt.go — uncovered branches ──

func Test_StartEndInt_Ranges_Nil(t *testing.T) {
	// Arrange
	var se *corerange.StartEndInt
	ranges := se.Ranges()

	// Act
	actual := args.Map{"rangesLen": len(ranges)}

	// Assert
	expected := args.Map{"rangesLen": 0}
	expected.ShouldBeEqual(t, 0, "StartEndInt Ranges returns empty -- nil receiver", actual)
}

func Test_StartEndInt_ZeroStartEnd(t *testing.T) {
	// Arrange — Start=0, End=0
	se := &corerange.StartEndInt{Start: 0, End: 0}

	// Act
	actual := args.Map{
		"invalidStart": se.IsInvalidStart(),
		"invalidEnd":   se.IsInvalidEnd(),
		"invalidBoth":  se.IsInvalidStartEndBoth(),
		"invalidAny":   se.IsInvalidAnyStartEnd(),
		"bothDefined":  se.IsStartEndBothDefined(),
	}
	expected := args.Map{
		"invalidStart": true, "invalidEnd": true,
		"invalidBoth": true, "invalidAny": true,
		"bothDefined": false,
	}
	expected.ShouldBeEqual(t, 0, "StartEndInt zero values -- both invalid", actual)
}

func Test_StartEndInt_RangeInt16(t *testing.T) {
	// Arrange
	se := &corerange.StartEndInt{Start: 3, End: 7}
	mm := &corerange.MinMaxInt16{Min: 0, Max: 100}

	// Act
	actual := args.Map{"notNil": se.RangeInt16(mm) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndInt RangeInt16 returns non-nil -- valid", actual)
}

func Test_StartEndInt_RangeInt8(t *testing.T) {
	// Arrange
	se := &corerange.StartEndInt{Start: 3, End: 7}
	mm := &corerange.MinMaxInt8{Min: 0, Max: 10}

	// Act
	actual := args.Map{"notNil": se.RangeInt8(mm) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndInt RangeInt8 returns non-nil -- valid", actual)
}

func Test_StartEndInt_RangeInt_FromWithinStringRangeUin(t *testing.T) {
	// Arrange
	se := &corerange.StartEndInt{Start: 3, End: 7}
	mm := &corerange.MinMaxInt{Min: 0, Max: 100}

	// Act
	actual := args.Map{"notNil": se.RangeInt(mm) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndInt RangeInt returns non-nil -- valid minMax", actual)
}

// ── StartEndSimpleString.go — uncovered branches ──

func Test_StartEndSimpleString_EmptyStart(t *testing.T) {
	// Arrange
	se := &corerange.StartEndSimpleString{Start: "", End: "world"}

	// Act
	actual := args.Map{
		"invalidStart": se.IsInvalidStart(),
		"invalidEnd":   se.IsInvalidEnd(),
		"invalidBoth":  se.IsInvalidStartEndBoth(),
		"invalidAny":   se.IsInvalidAnyStartEnd(),
		"bothDefined":  se.IsStartEndBothDefined(),
	}

	// Assert
	expected := args.Map{
		"invalidStart": true, "invalidEnd": false,
		"invalidBoth": false, "invalidAny": true,
		"bothDefined": false,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString empty start -- partial validity", actual)
}

func Test_StartEndSimpleString_RangeInt16(t *testing.T) {
	// Arrange
	se := &corerange.StartEndSimpleString{Start: "3", End: "7"}
	mm := &corerange.MinMaxInt16{Min: 0, Max: 100}

	// Act
	actual := args.Map{"notNil": se.RangeInt16(mm) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString RangeInt16 returns non-nil -- valid", actual)
}

func Test_StartEndSimpleString_RangeInt8(t *testing.T) {
	// Arrange
	se := &corerange.StartEndSimpleString{Start: "3", End: "7"}
	mm := &corerange.MinMaxInt8{Min: 0, Max: 10}

	// Act
	actual := args.Map{"notNil": se.RangeInt8(mm) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString RangeInt8 returns non-nil -- valid", actual)
}

func Test_StartEndSimpleString_RangeInt_WithMinMax(t *testing.T) {
	// Arrange
	se := &corerange.StartEndSimpleString{Start: "3", End: "7"}
	mm := &corerange.MinMaxInt{Min: 0, Max: 100}

	// Act
	actual := args.Map{"notNil": se.RangeInt(mm) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString RangeInt returns non-nil -- with minMax", actual)
}

func Test_StartEndSimpleString_StartEndString_BothEmpty(t *testing.T) {
	// Arrange
	se := &corerange.StartEndSimpleString{Start: "", End: ""}
	result := se.StartEndString()

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"isValid": result.IsValid,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isValid": false,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString StartEndString returns invalid -- both empty", actual)
}

// ── RangeByte — IsValidPlusWithinRange not valid ──

func Test_RangeByte_IsValidPlusWithinRange_NotValid(t *testing.T) {
	// Arrange
	rb := corerange.NewRangeByteMinMax("abc", "|", 0, 255)

	// Act
	actual := args.Map{"result": rb.IsValidPlusWithinRange(5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeByte IsValidPlusWithinRange returns false -- not valid", actual)
}

// ── RangeInt — NewRangeInt invalid start/end parsing ──

func Test_RangeInt_NewRangeInt_SingleElement(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("5", ":", 0, 10)

	// Act
	actual := args.Map{
		"isValid": ri.IsValid,
		"hasStart": ri.HasStart,
		"hasEnd": ri.HasEnd,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"hasStart": true,
		"hasEnd": false,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt NewRangeInt single element -- no end", actual)
}

func Test_RangeInt_NewRangeInt_ReversedRange(t *testing.T) {
	// Arrange
	// end < start → isValid=false
	ri := corerange.NewRangeIntMinMax("7:3", ":", 0, 10)

	// Act
	actual := args.Map{"isValid": ri.IsValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeInt NewRangeInt reversed range -- end less than start", actual)
}

func Test_RangeInt_NewRangeInt_OutOfMinMax(t *testing.T) {
	// Arrange
	// valid parse but outside min/max bounds
	ri := corerange.NewRangeIntMinMax("1:20", ":", 5, 10)

	// Act
	actual := args.Map{"isValid": ri.IsValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeInt NewRangeInt out of bounds -- below min", actual)
}
