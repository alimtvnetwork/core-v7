package coreversion

import "gitlab.com/auk-go/core/corecomparator"

func hasDeductUsingNilNess(left *Version, right *Version) (r corecomparator.Compare, isApplicable bool) {
	if left == nil && right == nil {
		return corecomparator.Equal, true
	}

	if left != nil && right == nil {
		return corecomparator.LeftGreater, true
	}

	if left == nil && right != nil {
		return corecomparator.LeftLess, true
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual, true
	}

	if left == right {
		return corecomparator.Equal, true
	}

	if left.VersionCompact == right.VersionCompact {
		return corecomparator.Equal, true
	}

	return corecomparator.Inconclusive, false
}
