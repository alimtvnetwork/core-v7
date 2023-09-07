package messages

import (
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/errcore"
)

var (
	ComparatorOutOfRangeMessage = errcore.RangeNotMeet(
		errcore.ComparatorShouldBeWithinRangeType.String(),
		corecomparator.Min(),
		corecomparator.Max(),
		corecomparator.Ranges())
)
