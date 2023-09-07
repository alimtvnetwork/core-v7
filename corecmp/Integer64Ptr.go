package corecmp

import "gitlab.com/auk-go/core/corecomparator"

func Integer64Ptr(left, right *int64) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	return Integer64(*left, *right)
}
