package bytetype

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	BasicEnumImpl = enumimpl.NewBasicByte(
		reflectinternal.TypeName(Variant(0)),
		[]byte{},
		[]string{},
		constants.Zero,
		constants.MaxUnit8)
)
