package bytetype

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingMap(
		reflectinternal.TypeName(Variant(0)),
		map[byte]string{
			Zero.Value():  "Zero",
			Min.Value():   "Min",
			One.Value():   "One",
			Two.Value():   "Two",
			Three.Value(): "Three",
			Max.Value():   "Max",
		})
)
