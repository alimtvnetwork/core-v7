package bytetype

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	basicEnumImpl = enumimpl.NewBasicByte(
		&[]byte{},
		&[]string{},
		constants.Zero,
		constants.MaxUnit8)
)
