package coreversion

func EmptyUsingCompactVersion(compactVersion string) *Version {
	return &Version{
		VersionCompact: compactVersion,
		Compiled:       "v" + compactVersion,
		IsInvalid:      false,
		VersionMajor:   InvalidVersionValue,
		VersionMinor:   InvalidVersionValue,
		VersionPatch:   InvalidVersionValue,
		VersionBuild:   InvalidVersionValue,
	}
}
