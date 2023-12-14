package coreversion

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/enums/versionindexes"
)

type Version struct {
	VersionCompact string `json:"Compact,omitempty"`   // ex : 1.0.1
	Compiled       string `json:"Compiled,omitempty"`  // ex : v1.0.1
	IsInvalid      bool   `json:"IsInvalid,omitempty"` // JSON export field for serialize
	VersionMajor   int    `json:"Major,omitempty"`
	VersionMinor   int    `json:"Minor,omitempty"`
	VersionPatch   int    `json:"Patch,omitempty"`
	VersionBuild   int    `json:"Build,omitempty"`
}

func (it Version) String() string {
	return it.CompiledVersion()
}

// VersionDisplay
//
// Display with a prefix of `v`
func (it *Version) VersionDisplay() string {
	if it == nil || it.VersionCompact == "" {
		return constants.EmptyString
	}

	return VSymbol + it.VersionCompact
}

// CompiledVersion
//
// It is similar to DisplayVersion, however,
// it gets generated during the creation time
// from the parsed major, minor, patch, build versions
func (it *Version) CompiledVersion() string {
	if it == nil || it.Compiled == "" {
		return constants.EmptyString
	}

	return it.Compiled
}

func (it *Version) VersionDisplayMajor() string {
	if it == nil ||
		it.VersionCompact == "" ||
		it.IsMajorInvalid() ||
		it.IsSafeInvalidCheck() {
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

func (it *Version) IsMajorInvalidOrZero() bool {
	return it == nil ||
		it.VersionMajor == InvalidVersionValue ||
		it.VersionMajor == constants.Zero
}

func (it *Version) IsMinorInvalidOrZero() bool {
	return it == nil ||
		it.VersionMinor == InvalidVersionValue ||
		it.VersionMinor == constants.Zero
}

func (it *Version) IsPatchInvalidOrZero() bool {
	return it == nil ||
		it.VersionPatch == InvalidVersionValue ||
		it.VersionPatch == constants.Zero
}

func (it *Version) IsBuildInvalidOrZero() bool {
	return it == nil ||
		it.VersionBuild == InvalidVersionValue ||
		it.VersionBuild == constants.Zero
}

func (it *Version) isInvalidOrEmptyAll() bool {
	return it.IsInvalid == true ||
		it.IsMajorInvalidOrZero() &&
			it.IsMinorInvalidOrZero() &&
			it.IsPatchInvalidOrZero() &&
			it.IsBuildInvalidOrZero()
}

func (it *Version) IsEmptyOrInvalid() bool {
	return it == nil ||
		it.IsInvalid == true ||
		it.VersionDisplay() == "" ||
		it.isInvalidOrEmptyAll()
}

func (it *Version) HasAnyItem() bool {
	return !it.IsEmptyOrInvalid()
}

func (it *Version) IsDefined() bool {
	return !it.IsEmptyOrInvalid()
}

func (it *Version) IsSafeInvalidCheck() bool {
	return it.IsEmptyOrInvalid()
}

func (it *Version) IsInvalidOrEmpty() bool {
	return it.IsEmptyOrInvalid()
}

func (it *Version) IsVersionCompareNotEqual(
	versionCompact string,
) bool {
	return !it.IsVersionCompareEqual(
		versionCompact)
}

func (it *Version) IsVersionCompareEqual(
	versionCompact string,
) bool {
	if it == nil && versionCompact == "" {
		return true
	}

	if it == nil && versionCompact != "" {
		return false
	}

	return it.VersionCompact == versionCompact
}

func (it *Version) ValueByIndex(
	index versionindexes.Index,
) int {
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

func (it *Version) ValueByIndexes(
	indexes ...versionindexes.Index,
) []int {
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

func (it *Version) IsMajorStringAtLeast(comparingMajor string) bool {
	// fine to swallow error
	majorInt, _ := strconv.Atoi(comparingMajor)

	return corecmp.Integer(it.VersionMajor, majorInt).
		IsLeftGreaterOrGreaterEqualOrEqual()
}

// IsMajorMinorAtLeast
//
// Current major version and minor is greater or equal to the given ones.
func (it *Version) IsMajorMinorAtLeast(
	major, minor int,
) bool {
	return it.MajorMinor(major, minor).
		IsLeftGreaterOrGreaterEqualOrEqual()
}

// IsMajorBuildAtLeast
//
// Current major version and build is greater or equal to the given ones.
func (it *Version) IsMajorBuildAtLeast(
	major, build int,
) bool {
	return it.MajorBuild(major, build).
		IsLeftGreaterEqualLogically()
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
	majorCmp := corecmp.Integer(
		it.VersionMajor, major)

	if majorCmp.IsNotEqualLogically() {
		return majorCmp
	}

	minorCmp := corecmp.Integer(
		it.VersionMinor, minor)

	if minorCmp.IsNotEqualLogically() {
		return minorCmp
	}

	return corecomparator.Equal
}

func (it *Version) MajorMinorPatchBuildString(
	major,
	minor,
	build,
	patch string,
) corecomparator.Compare {
	// fine to swallow error
	majorInt, _ := strconv.Atoi(major)
	minorInt, _ := strconv.Atoi(minor)
	patchInt, _ := strconv.Atoi(patch)
	buildInt, _ := strconv.Atoi(build)

	return it.MajorMinorPatchBuild(
		majorInt,
		minorInt,
		patchInt,
		buildInt)
}

func (it *Version) MajorBuildString(
	major,
	build string,
) corecomparator.Compare {
	// fine to swallow error
	majorInt, _ := strconv.Atoi(major)
	buildInt, _ := strconv.Atoi(build)

	return it.MajorBuild(
		majorInt, buildInt)
}

func (it *Version) MajorBuild(
	major,
	build int,
) corecomparator.Compare {
	majorCmp := corecmp.Integer(
		it.VersionMajor, major)

	if majorCmp.IsNotEqualLogically() {
		return majorCmp
	}

	buildCmp := corecmp.Integer(
		it.VersionBuild, build)

	if buildCmp.IsNotEqualLogically() {
		return buildCmp
	}

	return corecomparator.Equal
}

func (it *Version) Patch(
	patch int,
) corecomparator.Compare {
	patchCmp := corecmp.Integer(
		it.VersionPatch, patch)

	if patchCmp.IsNotEqualLogically() {
		return patchCmp
	}

	return corecomparator.Equal
}

func (it *Version) MajorPatch(
	major,
	patch int,
) corecomparator.Compare {
	majorCmp := corecmp.Integer(
		it.VersionMajor, major)

	if majorCmp.IsNotEqualLogically() {
		return majorCmp
	}

	patchCmp := corecmp.Integer(
		it.VersionPatch, patch)

	if patchCmp.IsNotEqualLogically() {
		return patchCmp
	}

	return corecomparator.Equal
}

func (it *Version) Build(
	build int,
) corecomparator.Compare {
	buildCmp := corecmp.Integer(
		it.VersionBuild, build)

	if buildCmp.IsNotEqualLogically() {
		return buildCmp
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

func (it *Version) MajorMinorPatchBuild(
	major,
	minor,
	patch,
	build int,
) corecomparator.Compare {
	majorMinorPatch := it.MajorMinorPatch(
		major,
		minor,
		patch)

	if majorMinorPatch.IsNotEqualLogically() {
		return majorMinorPatch
	}

	// everything equal before
	buildCmp := it.Build(build)

	if buildCmp.IsNotEqualLogically() {
		return buildCmp
	}

	return corecomparator.Equal
}

func (it *Version) Compare(
	right *Version,
) corecomparator.Compare {
	return Compare(it, right)
}

func (it *Version) IsEqual(
	right *Version,
) bool {
	return Compare(it, right).IsEqual()
}

// IsLeftLessThan it < right
func (it *Version) IsLeftLessThan(
	right *Version,
) bool {
	return Compare(it, right).IsLeftLess()
}

// IsLeftGreaterThan it > right
func (it *Version) IsLeftGreaterThan(
	right *Version,
) bool {
	return Compare(it, right).IsLeftGreater()
}

// IsLeftLessThanOrEqual it <= right
func (it *Version) IsLeftLessThanOrEqual(
	right *Version,
) bool {
	return Compare(it, right).IsLeftLessOrLessEqualOrEqual()
}

// IsLeftGreaterThanOrEqual it >= right
func (it *Version) IsLeftGreaterThanOrEqual(
	right *Version,
) bool {
	return Compare(it, right).
		IsLeftGreaterOrGreaterEqualOrEqual()
}

func (it *Version) IsExpectedComparison(
	expectedComparison corecomparator.Compare,
	right *Version,
) bool {
	c := Compare(it, right)

	return c.
		IsCompareEqualLogically(expectedComparison)
}

// IsExpectedComparisonRawVersion
//
//	@Description: it returns the expected comparison result
//	@param expectedComparison
//	@param rightVersion : can have "v0.0.0" or "0.0.0" or "v0.0.0.0" or "v0" or "v0.1"
//
//	@return bool
func (it *Version) IsExpectedComparisonRawVersion(
	expectedComparison corecomparator.Compare,
	rightVersion string,
) bool {
	return it.IsExpectedComparison(
		expectedComparison,
		New.DefaultPtr(rightVersion),
	)
}

// IsAtLeast
//
//	@Description: it returns the true if current version is at same or above as the given one or more
//	@param rightVersion : can have "v0.0.0" or "0.0.0" or "v0.0.0.0" or "v0" or "v0.1"
//
//	@return bool
func (it *Version) IsAtLeast(
	rightVersion string,
) bool {
	return it.IsExpectedComparison(
		corecomparator.LeftGreaterEqual,
		New.DefaultPtr(rightVersion),
	)
}

// IsEqualVersionString
//
//	@Description: it returns the true if current version same by deduction
//	@param rightVersion : can have "v0.0.0" or "0.0.0" or "v0.0.0.0" or "v0" or "v0.1"
//
//	@return bool
func (it *Version) IsEqualVersionString(
	rightVersion string,
) bool {
	return it.IsExpectedComparison(
		corecomparator.Equal,
		New.DefaultPtr(rightVersion),
	)
}

// IsLowerVersionString
//
//	@Description: it returns the true if current version less than the given version.
//	@param rightVersion : can have "v0.0.0" or "0.0.0" or "v0.0.0.0" or "v0" or "v0.1"
//
//	@return bool
func (it *Version) IsLowerVersionString(
	rightVersion string,
) bool {
	return it.IsExpectedComparison(
		corecomparator.LeftLess,
		New.DefaultPtr(rightVersion),
	)
}

// IsLowerEqualVersionString
//
//	@Description: it returns the true if current version less or equal than the given version.
//	@param rightVersion : can have "v0.0.0" or "0.0.0" or "v0.0.0.0" or "v0" or "v0.1"
//
//	@return bool
func (it *Version) IsLowerEqualVersionString(
	rightVersion string,
) bool {
	return it.IsExpectedComparison(
		corecomparator.LeftLessEqual,
		New.DefaultPtr(rightVersion),
	)
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
		Compiled:       it.Compiled,
		IsInvalid:      it.IsInvalid,
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

	toVersion := it.Clone()

	return &toVersion
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
	return jsonResult.Deserialize(it)
}

func (it *Version) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}
