package issetter

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/internal/messages"
)

// IsCompareResult Here left is v, and right is `n`
func (it Value) IsCompareResult(n byte, compare corecomparator.Compare) bool {
	switch compare {
	case corecomparator.Equal:
		return it.IsEqual(n)
	case corecomparator.LeftGreater:
		return it.IsGreater(n)
	case corecomparator.LeftGreaterEqual:
		return it.IsGreaterEqual(n)
	case corecomparator.LeftLess:
		return it.IsLess(n)
	case corecomparator.LeftLessEqual:
		return it.IsLessEqual(n)
	case corecomparator.NotEqual:
		return !it.IsEqual(n)
	default:
		panic(messages.ComparatorOutOfRangeMessage)
	}
}
