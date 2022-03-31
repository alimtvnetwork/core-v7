package corepayload

import "gitlab.com/evatix-go/core/coredata/coredynamic"

var (
	Empty              = emptyCreator{}
	New                = newCreator{}
	attributesTypeName = coredynamic.TypeName(Attributes{})
)
