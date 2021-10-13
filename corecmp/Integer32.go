package corecmp

import "gitlab.com/evatix-go/core/corecomparator"

func Integer32(left, right int32) corecomparator.Compare {
	if left == right {
		return corecomparator.Equal
	} else if left < right {
		return corecomparator.LeftLess
	} else if left > right {
		return corecomparator.LeftGreater
	}

	return corecomparator.NotEqual
}
