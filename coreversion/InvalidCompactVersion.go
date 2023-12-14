package coreversion

func InvalidCompactVersion(compactVersion string) *Version {
	return &Version{
		VersionCompact: compactVersion,
		Compiled:       compactVersion,
		IsInvalid:      true,
		VersionMajor:   InvalidVersionValue,
		VersionMinor:   InvalidVersionValue,
		VersionPatch:   InvalidVersionValue,
		VersionBuild:   InvalidVersionValue,
	}
}
