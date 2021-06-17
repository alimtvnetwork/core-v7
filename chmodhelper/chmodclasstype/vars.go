package chmodclasstype

import "gitlab.com/evatix-go/core/coreimpl/enumimpl"

var (
	Ranges = [...]string{
		UnInitialized: "UnInitialized",
		All:           "All",
		Owner:         "Owner",
		Group:         "Group",
		Other:         "Other",
		OwnerGroup:    "OwnerGroup",
		GroupOther:    "GroupOther",
		OwnerOther:    "OwnerOther",
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(Ranges[:])
)
