package corecomparator

import "time"

func TimePtr(left, right *time.Time) Compare {
	if left == nil && right == nil {
		return Equal
	}

	if left == nil || right == nil {
		return NotEqual
	}

	return Time(*left, *right)
}
