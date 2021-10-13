package coreversion

import (
	"gitlab.com/evatix-go/core/corecmp"
	"gitlab.com/evatix-go/core/corecomparator"
)

func Compare(
	left,
	right *Version,
) corecomparator.Compare {
	compare, isApplicable := hasDeductUsingNilNess(left, right)

	if isApplicable {
		return compare
	}

	majorVersionCompare := corecmp.Integer(
		left.VersionMajor,
		right.VersionMajor)

	if majorVersionCompare.IsNotEqualLogically() {
		return majorVersionCompare
	}

	// proceed only on equal
	minorVersionCompare := corecmp.Integer(
		left.VersionMinor,
		right.VersionMinor)

	if minorVersionCompare.IsNotEqualLogically() {
		return minorVersionCompare
	}

	patchVersionCompare := corecmp.Integer(
		left.VersionPatch,
		right.VersionPatch,
	)

	if patchVersionCompare.IsNotEqualLogically() {
		return patchVersionCompare
	}

	return corecomparator.Equal
}
