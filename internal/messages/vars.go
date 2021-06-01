package messages

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/msgtype"
)

var (
	ComparatorOutOfRangeMessage = msgtype.RangeNotMeet(
		msgtype.ComparatorShouldBeWithinRange.String(),
		corecomparator.Min(),
		corecomparator.Max(),
		corecomparator.Ranges())
)
