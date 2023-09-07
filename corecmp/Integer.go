package corecmp

import "gitlab.com/auk-go/core/corecomparator"

func Integer(left, right int) corecomparator.Compare {
	if left == right {
		return corecomparator.Equal
	} else if left < right {
		return corecomparator.LeftLess
	} else if left > right {
		return corecomparator.LeftGreater
	}

	return corecomparator.NotEqual
}
