package corecmp

import "gitlab.com/auk-go/core/corecomparator"

func IntegerPtr(left, right *int) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	return Integer(*left, *right)
}
