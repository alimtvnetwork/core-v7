package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/osconsts"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage4 — integerOutOfRange.ToUnsignedInt32 x32 branch
//
// Line 21-23: osconsts.IsX32Architecture == true path
// On 64-bit platforms this branch is normally unreachable.
// We temporarily flip the var to exercise it.
// ══════════════════════════════════════════════════════════════════════════════

func Test_ToUnsignedInt32_X32Branch_InRange(t *testing.T) {
	// Arrange
	original := osconsts.IsX32Architecture
	osconsts.IsX32Architecture = true
	defer func() { osconsts.IsX32Architecture = original }()

	// Act
	result := coremath.IsOutOfRange.Integer.ToUnsignedInt32(100)

	// Assert
	convey.Convey("ToUnsignedInt32 returns false for in-range value on x32", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_ToUnsignedInt32_X32Branch_Negative(t *testing.T) {
	// Arrange
	original := osconsts.IsX32Architecture
	osconsts.IsX32Architecture = true
	defer func() { osconsts.IsX32Architecture = original }()

	// Act
	result := coremath.IsOutOfRange.Integer.ToUnsignedInt32(-1)

	// Assert
	convey.Convey("ToUnsignedInt32 returns true for negative value on x32", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}
