package codegentype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Simple:           "Simple",
		MultipleArranges: "MultipleArranges",
	}

	basicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Simple,
		ranges[:],
	)
)
