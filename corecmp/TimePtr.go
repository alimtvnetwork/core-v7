package corecmp

import (
	"time"

	"gitlab.com/auk-go/core/corecomparator"
)

func TimePtr(left, right *time.Time) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	return Time(*left, *right)
}
