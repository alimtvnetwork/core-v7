package coreversion

import "gitlab.com/evatix-go/core/constants"

func Empty() *Version {
	return &Version{
		VersionCompact: constants.EmptyString,
		VersionMajor:   InvalidVersionValue,
		VersionMinor:   InvalidVersionValue,
		VersionPatch:   InvalidVersionValue,
		VersionBuild:   InvalidVersionValue,
	}
}
