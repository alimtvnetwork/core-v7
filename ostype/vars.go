package ostype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
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
		reflectinternal.TypeName(Any),
		osTypesStrings[:],
	)

	basicEnumImplOsGroup = enumimpl.New.BasicByte.UsingTypeSlice(
		reflectinternal.TypeName(WindowsGroup),
		osGroups,
	)
)
