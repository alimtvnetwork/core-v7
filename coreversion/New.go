package coreversion

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/enums/versionindexes"
)

// New
//
// Create new Version from given "v0.1.0" or "v0.0" or "v1" even "0.0.0" or empty string
//
// Examples for valid input
//  - "v0.0.0.0" or "0.0.0.0" represents "v{MajorInt}.{MinorInt}.{PatchInt}.{BuildInt}"
//  - "v0.0.0"   or "0.0.0"   represents "v{MajorInt}.{MinorInt}.{PatchInt}"
//  - "v0.0"     or "0.0"     represents "v{MajorInt}.{MinorInt}"
//  - "v0"       or "0"       represents "v{MajorInt}"
//  - "v"        or ""        represents "" Empty or Invalid Result but don't panic
//  - ""                      represents "" Empty or Invalid Result but don't panic
func New(version string) *Version {
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
		versionsValuesSlice := converters.StringsToIntegers(
			false,
			slice)

		return &Version{
			VersionCompact: trimmedVersion,
			VersionMajor:   versionsValuesSlice[versionindexes.Major],
			VersionMinor:   versionsValuesSlice[versionindexes.Minor],
			VersionPatch:   versionsValuesSlice[versionindexes.Patch],
			VersionBuild:   versionsValuesSlice[versionindexes.Build],
		}
	}

	if length >= versionindexes.PatchLength {
		versionsValuesSlice := converters.StringsToIntegers(
			false,
			slice)

		return &Version{
			VersionCompact: trimmedVersion,
			VersionMajor:   versionsValuesSlice[versionindexes.Major],
			VersionMinor:   versionsValuesSlice[versionindexes.Minor],
			VersionPatch:   versionsValuesSlice[versionindexes.Patch],
			VersionBuild:   InvalidVersionValue,
		}
	}

	if length >= versionindexes.MinorLength {
		versionsValuesSlice := converters.StringsToIntegers(
			false,
			slice)

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
