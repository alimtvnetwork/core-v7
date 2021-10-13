package corecmp

import "gitlab.com/evatix-go/core/corecomparator"

func Integer16(left, right int16) corecomparator.Compare {
	if left == right {
		return corecomparator.Equal
	} else if left < right {
		return corecomparator.LeftLess
	} else if left > right {
		return corecomparator.LeftGreater
	}

	return corecomparator.NotEqual
}
