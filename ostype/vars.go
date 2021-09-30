package ostype

var (
	CurrentGroupVariant = GetGroupVariant()
	// CurrentGroup Current os group
	CurrentGroup = CurrentGroupVariant.Group
	// Type Current Os Type
	Type = CurrentGroupVariant.Variation

	osGroups = []string{
		"WindowsGroup",
		"UnixGroup",
		"AndroidGroup",
		"JavaScriptGroup",
		"UnknownGroup",
	}
)
