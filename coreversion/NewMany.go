package coreversion

import "gitlab.com/evatix-go/core/constants"

func NewMany(versions ...string) *VersionsCollection {
	versionsCollection := NewVersionsCollection(len(versions) + constants.Capacity2)

	return versionsCollection.AddVersionsRaw(versions...)
}
