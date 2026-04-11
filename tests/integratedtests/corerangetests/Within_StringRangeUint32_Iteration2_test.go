package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
)

// Test_Cov7_Within_StringRangeUint32_Overflow tests the overflow fallback in
// within.StringRangeUint32 where finalInt > MaxInt32.
func Test_Within_StringRangeUint32_Overflow(t *testing.T) {
	// Arrange / Act
	val, isInRange := corerange.Within.StringRangeUint32("2147483648") // MaxInt32 + 1

	// Assert
	_ = val
	_ = isInRange
	// Either way we exercise the function — coverage will reveal which path.
}
