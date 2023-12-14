package corepayload

import (
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	Empty              = emptyCreator{}
	New                = newCreator{}
	attributesTypeName = reflectinternal.TypeName(Attributes{})
)
