package corecmp

import "gitlab.com/auk-go/core/corecomparator"

func AnyItem(left, right interface{}) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	if left == right {
		return corecomparator.Equal
	}

	return corecomparator.Inconclusive
}
