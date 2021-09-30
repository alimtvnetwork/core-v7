package corecomparator

import "time"

func Time(left, right time.Time) Compare {
	if left.Before(right) {
		return LeftLess
	} else if left.After(right) {
		return LeftGreater
	} else if left.Equal(right) {
		return Equal
	}

	return NotEqual
}
