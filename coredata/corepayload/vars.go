package corepayload

import "gitlab.com/auk-go/core/coredata/coredynamic"

var (
	Empty              = emptyCreator{}
	New                = newCreator{}
	attributesTypeName = coredynamic.TypeName(Attributes{})
)
