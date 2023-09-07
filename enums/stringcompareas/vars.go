package stringcompareas

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	stringRanges = [...]string{
		Equal:         "Equal",
		StartsWith:    "StartsWith",
		EndsWith:      "EndsWith",
		Anywhere:      "Anywhere",
		AnyChars:      "AnyChars",
		Contains:      "Contains",
		Regex:         "Regex",
		NotEqual:      "NotEqual",
		NotStartsWith: "NotStartsWith",
		NotEndsWith:   "NotEndsWith",
		NotContains:   "NotContains",
		NotAnyChars:   "NotAnyChars",
		NotMatchRegex: "NotMatchRegex",
		Invalid:       "Invalid",
	}

	BasicEnumImpl = enumimpl.
			New.
			BasicByte.
			Default(
			Equal,
			stringRanges[:])

	rangesMap = map[Variant]IsLineCompareFunc{
		Equal:         isEqualFunc,
		StartsWith:    isStartsWithFunc,
		EndsWith:      isEndsWithFunc,
		Anywhere:      isAnywhereFunc,
		AnyChars:      isAnyCharsFunc,
		Contains:      isAnywhereFunc, // alias for isAnyWhere
		Regex:         isRegexFunc,
		NotEqual:      isNotEqualFunc,
		NotStartsWith: isNotStartsWithFunc,
		NotEndsWith:   isNotEndsWithFunc,
		NotContains:   isNotContainsFunc,
		NotAnyChars:   isNotAnyCharsFunc,
		NotMatchRegex: isNotMatchRegex,
	}

	negativeCases = []Variant{
		NotEqual,
		NotStartsWith,
		NotEndsWith,
		NotContains,
		NotAnyChars,
		NotMatchRegex,
	}
)
