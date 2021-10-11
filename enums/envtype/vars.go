package envtype

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	Ranges = [...]string{
		Uninitialized: "Uninitialized",
		Development:   "Development",
		Development1:  "Development1",
		Development2:  "Development2",
		Test:          "Test",
		Test1:         "Test1",
		Test2:         "Test2",
		Production:    "Production",
		Production1:   "Production1",
		Production2:   "Production2",
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		reflectinternal.TypeName(Uninitialized),
		Ranges[:])
)
