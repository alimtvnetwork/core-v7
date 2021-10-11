package messages

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/errcore"
)

var (
	ComparatorOutOfRangeMessage = errcore.RangeNotMeet(
		errcore.ComparatorShouldBeWithinRange.String(),
		corecomparator.Min(),
		corecomparator.Max(),
		corecomparator.Ranges())
)
