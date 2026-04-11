package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage7 — corerange remaining gaps
//
// Target 1: MinMaxByte.go:46-48 — DifferenceAbsolute negative diff branch
//   byte is unsigned (0-255), so diff = Max - Min is always >= 0.
//   Dead code.
//
// Target 2: within.go:89 — StringRangeUint32 finalInt > MaxInt32 branch
//   The range is already clamped to [0, MaxInt32], so finalInt is always
//   <= MaxInt32. Dead code.
// ══════════════════════════════════════════════════════════════════════════════

func Test_MinMaxByte_DifferenceAbsolute(t *testing.T) {
	// Arrange
	mmb := corerange.MinMaxByte{
		Min: 10,
		Max: 50,
	}

	// Act
	result := mmb.DifferenceAbsolute()

	// Assert
	convey.Convey("DifferenceAbsolute returns correct positive difference", t, func() {
		convey.So(result, convey.ShouldEqual, byte(40))
	})
}

func Test_Within_StringRangeUint32_Valid(t *testing.T) {
	// Arrange & Act
	val, isInRange := corerange.Within.StringRangeUint32("100")

	// Assert
	convey.Convey("StringRangeUint32 parses valid uint32 string", t, func() {
		convey.So(val, convey.ShouldEqual, uint32(100))
		convey.So(isInRange, convey.ShouldBeTrue)
	})
}

// Coverage note: Both uncovered branches are dead code:
// MinMaxByte.DifferenceAbsolute line 46: byte subtraction cannot be negative
// within.StringRangeUint32 line 89: input already clamped to [0, MaxInt32]
