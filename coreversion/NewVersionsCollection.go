package coreversion

func NewVersionsCollection(cap int) *VersionsCollection {
	return &VersionsCollection{
		Versions: make([]*Version, 0, cap),
	}
}
