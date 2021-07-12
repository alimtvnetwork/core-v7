package stringcomparator

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	stringRanges = [...]string{
		Equal:      "Equal",
		StartsWith: "StartsWith",
		EndsWith:   "EndsWith",
		AnyWhere:   "AnyWhere",
		NotEqual:   "NotEqual",
		Regex:      "Regex",
	}

	typeRanges = [...]Variant{
		Equal:      Equal,
		StartsWith: StartsWith,
		EndsWith:   EndsWith,
		AnyWhere:   AnyWhere,
		NotEqual:   NotEqual,
		Regex:      Regex,
	}

	basicEnumImpl = enumimpl.
		NewBasicByteUsingIndexedSlice(stringRanges[:])
)
