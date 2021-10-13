package corecmp

import "gitlab.com/evatix-go/core/corecomparator"

func Integer16Ptr(left, right *int16) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	return Integer16(*left, *right)
}
