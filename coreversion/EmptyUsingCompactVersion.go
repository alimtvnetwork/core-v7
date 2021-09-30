package coreversion

func EmptyUsingCompactVersion(compactVersion string) *Version {
	return &Version{
		VersionCompact: compactVersion,
		VersionMajor:   InvalidVersionValue,
		VersionMinor:   InvalidVersionValue,
		VersionPatch:   InvalidVersionValue,
		VersionBuild:   InvalidVersionValue,
	}
}
