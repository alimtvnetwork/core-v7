package corecmp

import "gitlab.com/auk-go/core/corecomparator"

func Integer32Ptr(left, right *int32) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	return Integer32(*left, *right)
}
