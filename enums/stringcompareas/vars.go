package stringcompareas

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
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

	basicEnumImpl = enumimpl.
			NewBasicByteUsingIndexedSlice(
			reflectinternal.TypeName(Equal), stringRanges[:])

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
