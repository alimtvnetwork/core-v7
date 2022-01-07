package coreversion

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecmp"
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/enums/versionindexes"
)

type Version struct {
	VersionCompact string // ex : 1.0.1
	VersionMajor   int
	VersionMinor   int
	VersionPatch   int
	VersionBuild   int
}

func (it *Version) String() string {
	return it.VersionDisplay()
}

func (it *Version) VersionDisplay() string {
	if it == nil || it.VersionCompact == "" {
		return constants.EmptyString
	}

	return VSymbol + it.VersionCompact
}

func (it *Version) VersionDisplayMajor() string {
	if it == nil || it.VersionCompact == "" || it.IsMajorInvalid() {
		return constants.EmptyString
	}

	return VSymbol + strconv.Itoa(it.VersionMajor)
}

func (it *Version) VersionDisplayMajorMinor() string {
	if it.IsMinorInvalid() {
		return it.VersionDisplayMajor()
	}

	return VSymbol +
		strconv.Itoa(it.VersionMajor) +
		constants.Dot +
		strconv.Itoa(it.VersionMinor)
}

func (it *Version) VersionDisplayMajorMinorPatch() string {
	if it.IsPatchInvalid() {
		return it.VersionDisplayMajorMinor()
	}

	return VSymbol +
		strconv.Itoa(it.VersionMajor) +
		constants.Dot +
		strconv.Itoa(it.VersionMinor) +
		constants.Dot +
		strconv.Itoa(it.VersionPatch)
}

func (it *Version) MajorString() string {
	if it == nil {
		return constants.EmptyString
	}

	return strconv.Itoa(it.VersionMajor)
}

func (it *Version) MinorString() string {
	if it == nil {
		return constants.EmptyString
	}

	return strconv.Itoa(it.VersionMinor)
}

func (it *Version) PatchString() string {
	if it == nil {
		return constants.EmptyString
	}

	return strconv.Itoa(it.VersionPatch)
}

func (it *Version) BuildString() string {
	if it == nil {
		return constants.EmptyString
	}

	return strconv.Itoa(it.VersionBuild)
}

func (it *Version) HasMajor() bool {
	return it != nil && it.VersionMajor > InvalidVersionValue
}

func (it *Version) HasMinor() bool {
	return it != nil && it.VersionMinor > InvalidVersionValue
}

func (it *Version) HasPatch() bool {
	return it != nil && it.VersionPatch > InvalidVersionValue
}

func (it *Version) HasBuild() bool {
	return it != nil && it.VersionBuild > InvalidVersionValue
}

func (it *Version) IsMajorInvalid() bool {
	return it == nil || it.VersionMajor == InvalidVersionValue
}

func (it *Version) IsMinorInvalid() bool {
	return it == nil || it.VersionMinor == InvalidVersionValue
}

func (it *Version) IsPatchInvalid() bool {
	return it == nil || it.VersionPatch == InvalidVersionValue
}

func (it *Version) IsBuildInvalid() bool {
	return it == nil || it.VersionBuild == InvalidVersionValue
}

func (it *Version) IsEmptyOrInvalid() bool {
	return it.VersionDisplay() == ""
}

func (it *Version) ValueByIndex(index versionindexes.Index) int {
	switch index {
	case versionindexes.Major:
		return it.VersionMajor
	case versionindexes.Minor:
		return it.VersionMinor
	case versionindexes.Patch:
		return it.VersionPatch
	case versionindexes.Build:
		return it.VersionBuild
	}

	return InvalidVersionValue
}

func (it *Version) ValueByIndexes(indexes ...versionindexes.Index) []int {
	slice := make([]int, len(indexes))

	for i, index := range indexes {
		slice[i] = it.ValueByIndex(index)
	}

	return slice
}

func (it *Version) AllVersionValues() []int {
	return it.ValueByIndexes(versionindexes.AllVersionIndexes...)
}

func (it *Version) AllValidVersionValues() []int {
	slice := it.AllValidVersionValues()

	for i, item := range slice {
		if item == InvalidVersionValue {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func (it *Version) Major(comparingMajor int) corecomparator.Compare {
	return corecmp.Integer(it.VersionMajor, comparingMajor)
}

func (it *Version) IsMajorAtLeast(comparingMajor int) bool {
	return corecmp.Integer(it.VersionMajor, comparingMajor).
		IsLeftGreaterOrGreaterEqualOrEqual()
}

func (it *Version) IsMajorMinorAtLeast(major, minor int) bool {
	return it.MajorMinor(major, minor).
		IsLeftGreaterOrGreaterEqualOrEqual()
}

func (it *Version) IsMajorMinorPatchAtLeast(
	major,
	minor,
	patch int,
) bool {
	cmp := it.MajorMinorPatch(
		major,
		minor,
		patch,
	)

	return cmp.
		IsLeftGreaterOrGreaterEqualOrEqual()
}

func (it *Version) MajorMinor(
	major,
	minor int,
) corecomparator.Compare {
	majorCmp := corecmp.Integer(it.VersionMajor, major)

	if majorCmp.IsNotEqualLogically() {
		return majorCmp
	}

	minorCmp := corecmp.Integer(it.VersionMinor, minor)

	if minorCmp.IsNotEqualLogically() {
		return minorCmp
	}

	return corecomparator.Equal
}

func (it *Version) MajorMinorPatch(
	major,
	minor,
	patch int,
) corecomparator.Compare {
	majorMinor := it.MajorMinor(major, minor)

	if majorMinor.IsNotEqualLogically() {
		return majorMinor
	}

	patchCmp := corecmp.Integer(
		it.VersionPatch,
		patch)

	if patchCmp.IsNotEqualLogically() {
		return patchCmp
	}

	return corecomparator.Equal
}

func (it *Version) Compare(right *Version) corecomparator.Compare {
	return Compare(it, right)
}

func (it *Version) IsEqual(right *Version) bool {
	return Compare(it, right).IsEqual()
}

// IsLeftLessThan it < right
func (it *Version) IsLeftLessThan(right *Version) bool {
	return Compare(it, right).IsLeftLess()
}

// IsLeftGreaterThan it > right
func (it *Version) IsLeftGreaterThan(right *Version) bool {
	return Compare(it, right).IsLeftGreater()
}

// IsLeftLessThanOrEqual it <= right
func (it *Version) IsLeftLessThanOrEqual(right *Version) bool {
	return Compare(it, right).IsLeftLessOrLessEqualOrEqual()
}

// IsLeftGreaterThanOrEqual it >= right
func (it *Version) IsLeftGreaterThanOrEqual(right *Version) bool {
	return Compare(it, right).IsLeftGreaterOrGreaterEqualOrEqual()
}

func (it *Version) IsExpectedComparison(
	right *Version,
	expectedComparison corecomparator.Compare,
) bool {
	return Compare(it, right) == expectedComparison
}

func (it *Version) IsExpectedComparisonUsingVersionString(
	rightVersion string, // can have "v0.0.0" or "0.0.0" or "v0.0.0.0" or "v0" or "v0.1"
	expectedComparison corecomparator.Compare,
) bool {
	return it.IsExpectedComparison(New(rightVersion), expectedComparison)
}

func (it *Version) ComparisonValueIndexes(
	right *Version,
	indexes ...versionindexes.Index,
) corecomparator.Compare {
	r, isApplicable := hasDeductUsingNilNess(it, right)

	if isApplicable {
		return r
	}

	leftVersions := make([]int, len(indexes))
	rightVersions := make([]int, len(indexes))
	for i, index := range indexes {
		leftVersions[i] = it.ValueByIndex(index)
		rightVersions[i] = right.ValueByIndex(index)
	}

	return corecmp.VersionSliceInteger(
		leftVersions,
		rightVersions)
}

func (it Version) Clone() Version {
	return Version{
		VersionCompact: it.VersionCompact,
		VersionMajor:   it.VersionMajor,
		VersionMinor:   it.VersionMinor,
		VersionPatch:   it.VersionPatch,
		VersionBuild:   it.VersionBuild,
	}
}

func (it *Version) ClonePtr() *Version {
	if it == nil {
		return nil
	}

	return &Version{
		VersionCompact: it.VersionCompact,
		VersionMajor:   it.VersionMajor,
		VersionMinor:   it.VersionMinor,
		VersionPatch:   it.VersionPatch,
		VersionBuild:   it.VersionBuild,
	}
}

func (it Version) NonPtr() Version {
	return it
}

func (it *Version) Ptr() *Version {
	return it
}

func (it Version) Json() corejson.Result {
	return corejson.New(it)
}

func (it Version) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Version) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Unmarshal(it)
}

func (it *Version) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}
