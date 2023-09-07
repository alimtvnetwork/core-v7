package ostype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

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
		"InvalidGroup",
	}

	basicEnumImplOsType = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Any),
		osTypesStrings[:])

	basicEnumImplOsGroup = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(WindowsGroup),
		osGroups)
)
