package anycmp

import (
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/isany"
)

func Cmp(left, right interface{}) corecomparator.Compare {
	if left == right {
		return corecomparator.Equal
	}

	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	isLeftNull, isRightNull := isany.NullLeftRight(left, right)
	isBothEqual := isLeftNull == isRightNull

	if isLeftNull && isBothEqual {
		// both null
		return corecomparator.Equal
	} else if !isBothEqual && isLeftNull || isRightNull {
		// any null but the other is not
		return corecomparator.NotEqual
	}

	return corecomparator.Inconclusive
}
