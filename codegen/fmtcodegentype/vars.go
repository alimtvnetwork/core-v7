package fmtcodegentype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Default:       "Default",
		WithFunction:  "WithFunction",
		WithFuncError: "WithFuncError",
	}

	rangesFmt = map[Variant]string{
		Default:       "%d : %s -> %s",
		WithFunction:  "%d : %s(%s) -> %s | %s",
		WithFuncError: "%d : %s - %s",
	}

	basicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Default,
		ranges[:],
	)
)
