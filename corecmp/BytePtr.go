package corecmp

import "gitlab.com/evatix-go/core/corecomparator"

func BytePtr(left, right *byte) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	return Byte(*left, *right)
}
