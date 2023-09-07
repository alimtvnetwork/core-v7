package coreversion

import (
	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
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

// CompareVersionString
//
// See New.Default for more details
func CompareVersionString(
	leftVersion,
	rightVersion string,
) corecomparator.Compare {
	left := New.Default(leftVersion)
	right := New.Default(rightVersion)

	return Compare(left, right)
}

// IsExpectedVersion
//
// See New.Default for more details
func IsExpectedVersion(
	expectedCompare corecomparator.Compare,
	leftVersion,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftVersion, rightVersion)

	return cmp.IsCompareEqualLogically(expectedCompare)
}

// IsAtLeast
//
//  returns true if left version is equal or greater than the right
func IsAtLeast(
	leftGreaterOrEqual,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftGreaterOrEqual, rightVersion)

	return cmp.IsLeftGreaterEqualLogically()
}

// IsLower
//
//  returns true if left version is less than the right version
func IsLower(
	leftGreaterOrEqual,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftGreaterOrEqual, rightVersion)

	return cmp.IsLeftLess()
}

// IsLowerOrEqual
//
//  returns true if left version is less or equal than the right version
func IsLowerOrEqual(
	leftGreaterOrEqual,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftGreaterOrEqual, rightVersion)

	return cmp.IsLeftLessEqualLogically()
}
