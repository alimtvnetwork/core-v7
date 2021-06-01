package inttype

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/internal/messages"
)

// IsCompareResult Here left is v, and right is `n`
func (v Variant) IsCompareResult(n int, compare corecomparator.Compare) bool {
	switch compare {
	case corecomparator.Equal:
		return v.IsEqual(n)
	case corecomparator.LeftGreater:
		return v.IsGreater(n)
	case corecomparator.LeftGreaterEqual:
		return v.IsGreaterEqual(n)
	case corecomparator.LeftLess:
		return v.IsLess(n)
	case corecomparator.LeftLessEqual:
		return v.IsLessEqual(n)
	case corecomparator.NotEqual:
		return !v.IsEqual(n)
	default:
		panic(messages.ComparatorOutOfRangeMessage)
	}
}
