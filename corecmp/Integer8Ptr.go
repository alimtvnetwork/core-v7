package corecmp

import "gitlab.com/auk-go/core/corecomparator"

func Integer8Ptr(left, right *int8) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	return Integer8(*left, *right)
}
