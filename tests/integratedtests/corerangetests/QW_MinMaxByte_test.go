package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
)

func Test_QW_MinMaxByte_DifferenceAbsolute(t *testing.T) {
	mm := corerange.MinMaxByte{Min: 200, Max: 100}
	_ = mm.DifferenceAbsolute()
}

func Test_QW_Within_StringRangeUint32_LargeValue(t *testing.T) {
	// StringRangeUint32 takes only 1 arg: string input
	val, inRange := corerange.Within.StringRangeUint32("3000000000")
	_, _ = val, inRange
}
