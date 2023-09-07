package coreversion

import (
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/enums/versionindexes"
)

type newCreator struct{}

// Default
//
// CreateUsingAliasMap new Version from given "v0.1.0" or "v0.0" or "v1" even "0.0.0" or empty string
//
// Examples for valid input
//  - "v0.0.0.0" or "0.0.0.0" represents "v{MajorInt}.{MinorInt}.{PatchInt}.{BuildInt}"
//  - "v0.0.0"   or "0.0.0"   represents "v{MajorInt}.{MinorInt}.{PatchInt}"
//  - "v0.0"     or "0.0"     represents "v{MajorInt}.{MinorInt}"
//  - "v0"       or "0"       represents "v{MajorInt}"
//  - "v"        or ""        represents "" Empty or Invalid Result but don't panic
//  - ""                      represents "" Empty or Invalid Result but don't panic
func (it newCreator) Default(version string) *Version {
	if version == "" {
		return Empty()
	}

	trimmed := strings.TrimSpace(version)

	if trimmed == "" {
		return Empty()
	}

	trimmedVersion := strings.TrimPrefix(trimmed, VSymbol)

	if trimmedVersion == "" {
		return Empty()
	}

	slice := strings.Split(trimmedVersion, constants.Dot)
	length := len(slice)

	if length >= versionindexes.BuildLength {
		versionsValuesSlice := converters.StringsTo.IntegersSkipErrors(
			slice...)

		return &Version{
			VersionCompact: trimmedVersion,
			VersionMajor:   versionsValuesSlice[versionindexes.Major],
			VersionMinor:   versionsValuesSlice[versionindexes.Minor],
			VersionPatch:   versionsValuesSlice[versionindexes.Patch],
			VersionBuild:   versionsValuesSlice[versionindexes.Build],
		}
	}

	if length >= versionindexes.PatchLength {
		versionsValuesSlice := converters.StringsTo.IntegersSkipErrors(
			slice...)

		return &Version{
			VersionCompact: trimmedVersion,
			VersionMajor:   versionsValuesSlice[versionindexes.Major],
			VersionMinor:   versionsValuesSlice[versionindexes.Minor],
			VersionPatch:   versionsValuesSlice[versionindexes.Patch],
			VersionBuild:   InvalidVersionValue,
		}
	}

	if length >= versionindexes.MinorLength {
		versionsValuesSlice := converters.StringsTo.IntegersSkipErrors(
			slice...)

		return &Version{
			VersionCompact: trimmedVersion,
			VersionMajor:   versionsValuesSlice[versionindexes.Major],
			VersionMinor:   versionsValuesSlice[versionindexes.Minor],
			VersionPatch:   InvalidVersionValue,
			VersionBuild:   InvalidVersionValue,
		}
	}

	if length >= versionindexes.MajorLength {
		majorVersion, _ := converters.StringToIntegerWithDefault(
			slice[versionindexes.Major],
			InvalidVersionValue)

		return &Version{
			VersionCompact: trimmedVersion,
			VersionMajor:   majorVersion,
			VersionMinor:   InvalidVersionValue,
			VersionPatch:   InvalidVersionValue,
			VersionBuild:   InvalidVersionValue,
		}
	}

	return EmptyUsingCompactVersion(trimmedVersion)
}

// Version
//
// CreateUsingAliasMap new Version from given "v0.1.0" or "v0.0" or "v1" even "0.0.0" or empty string
//
// Examples for valid input
//  - "v0.0.0.0" or "0.0.0.0" represents "v{MajorInt}.{MinorInt}.{PatchInt}.{BuildInt}"
//  - "v0.0.0"   or "0.0.0"   represents "v{MajorInt}.{MinorInt}.{PatchInt}"
//  - "v0.0"     or "0.0"     represents "v{MajorInt}.{MinorInt}"
//  - "v0"       or "0"       represents "v{MajorInt}"
//  - "v"        or ""        represents "" Empty or Invalid Result but don't panic
//  - ""                      represents "" Empty or Invalid Result but don't panic
func (it newCreator) Version(version string) *Version {
	return it.Default(version)
}

// Create
//
// CreateUsingAliasMap new Version from given "v0.1.0" or "v0.0" or "v1" even "0.0.0" or empty string
//
// Examples for valid input
//  - "v0.0.0.0" or "0.0.0.0" represents "v{MajorInt}.{MinorInt}.{PatchInt}.{BuildInt}"
//  - "v0.0.0"   or "0.0.0"   represents "v{MajorInt}.{MinorInt}.{PatchInt}"
//  - "v0.0"     or "0.0"     represents "v{MajorInt}.{MinorInt}"
//  - "v0"       or "0"       represents "v{MajorInt}"
//  - "v"        or ""        represents "" Empty or Invalid Result but don't panic
//  - ""                      represents "" Empty or Invalid Result but don't panic
func (it newCreator) Create(version string) *Version {
	return it.Default(version)
}

func (it newCreator) Major(majorString string) *Version {
	return it.Default(majorString)
}

// SpreadStrings
//
//  versionindexes.Major = v[0]
//  versionindexes.Minor = v[1]
//   ...
func (it newCreator) SpreadStrings(
	v ...string,
) *Version {
	actualCompiledVersionString := strings.Join(
		v[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

// SpreadIntegers
//
//  versionindexes.Major = v[0]
//  versionindexes.Minor = v[1]
//   ...
func (it newCreator) SpreadIntegers(
	v ...int,
) *Version {
	slice := make([]string, len(v))

	for i, indexedVersionValue := range v {
		slice[i] = strconv.Itoa(indexedVersionValue)
	}

	return it.SpreadStrings(slice...)
}

// SpreadUnsignedIntegers
//
//  versionindexes.Major = v[0]
//  versionindexes.Minor = v[1]
//   ...
func (it newCreator) SpreadUnsignedIntegers(
	v ...uint,
) *Version {
	slice := make([]string, len(v))

	for i, indexedVersionValue := range v {
		slice[i] = strconv.Itoa(int(indexedVersionValue))
	}

	return it.SpreadStrings(slice...)
}

// SpreadBytes
//
//  versionindexes.Major = v[0]
//  versionindexes.Minor = v[1]
//   ...
func (it newCreator) SpreadBytes(
	v ...byte,
) *Version {
	slice := make([]string, len(v))

	for i, indexedVersionValue := range v {
		slice[i] = strconv.Itoa(int(indexedVersionValue))
	}

	return it.SpreadStrings(slice...)
}

func (it newCreator) MajorMinor(
	major, minor string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: major,
		versionindexes.Minor: minor,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorMinorPatch(
	major,
	minor,
	patch string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: major,
		versionindexes.Minor: minor,
		versionindexes.Patch: patch,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorMinorPatchBuild(
	major,
	minor,
	patch,
	build string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: major,
		versionindexes.Minor: minor,
		versionindexes.Patch: patch,
		versionindexes.Build: build,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) All(
	major,
	minor,
	patch,
	build string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: major,
		versionindexes.Minor: minor,
		versionindexes.Patch: patch,
		versionindexes.Build: build,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) AllInt(
	major,
	minor,
	patch,
	build int,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: strconv.Itoa(major),
		versionindexes.Minor: strconv.Itoa(minor),
		versionindexes.Patch: strconv.Itoa(patch),
		versionindexes.Build: strconv.Itoa(build),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) AllByte(
	major,
	minor,
	patch,
	build byte,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: strconv.Itoa(int(major)),
		versionindexes.Minor: strconv.Itoa(int(minor)),
		versionindexes.Patch: strconv.Itoa(int(patch)),
		versionindexes.Build: strconv.Itoa(int(build)),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorMinorInt(
	major,
	minor int,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: strconv.Itoa(major),
		versionindexes.Minor: strconv.Itoa(minor),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorMinorPatchInt(
	major,
	minor,
	patch int,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: strconv.Itoa(major),
		versionindexes.Minor: strconv.Itoa(minor),
		versionindexes.Patch: strconv.Itoa(patch),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorBuildInt(
	major,
	build int,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: strconv.Itoa(major),
		versionindexes.Build: strconv.Itoa(build),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorBuild(
	major,
	build string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: major,
		versionindexes.Build: build,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorMinorBuild(
	major,
	minor,
	build string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: major,
		versionindexes.Minor: minor,
		versionindexes.Build: build,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorPatch(
	major,
	patch string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: major,
		versionindexes.Patch: patch,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MajorPatchInt(
	major,
	patch int,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Major: strconv.Itoa(major),
		versionindexes.Patch: strconv.Itoa(patch),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MinorBuildInt(
	minor,
	build int,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Minor: strconv.Itoa(minor),
		versionindexes.Build: strconv.Itoa(build),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) PatchBuildInt(
	patch,
	build int,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Patch: strconv.Itoa(patch),
		versionindexes.Build: strconv.Itoa(build),
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) MinorBuild(
	minor,
	build string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Minor: minor,
		versionindexes.Build: build,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) PatchBuild(
	patch,
	build string,
) *Version {
	actualVersionSlice := [...]string{
		versionindexes.Patch: patch,
		versionindexes.Build: build,
	}

	actualCompiledVersionString := strings.Join(
		actualVersionSlice[:],
		constants.Dot)

	return it.Default(actualCompiledVersionString)
}

func (it newCreator) Many(
	versions ...string,
) *VersionsCollection {
	versionsCollection := it.CollectionUsingCap(len(versions) + constants.Capacity2)

	return versionsCollection.AddVersionsRaw(versions...)
}

func (it newCreator) Collection(
	versions ...string,
) *VersionsCollection {
	versionsCollection := it.CollectionUsingCap(len(versions) + constants.Capacity2)

	return versionsCollection.AddVersionsRaw(versions...)
}

func (it newCreator) CollectionUsingCap(
	capacity int,
) *VersionsCollection {
	return &VersionsCollection{
		Versions: make([]*Version, 0, capacity),
	}
}

func (it newCreator) EmptyCollection() *VersionsCollection {
	return it.CollectionUsingCap(0)
}

func (it newCreator) Invalid() *Version {
	return Empty()
}

func (it newCreator) Empty() *Version {
	return Empty()
}
