package chmodclasstype

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		All:        "All",
		Owner:      "Owner",
		Group:      "Group",
		Other:      "Other",
		OwnerGroup: "OwnerGroup",
		GroupOther: "GroupOther",
		OwnerOther: "OwnerOther",
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		reflectinternal.TypeName(Invalid),
		Ranges[:])
)
